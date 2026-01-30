package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"
	"unicode"

	"golang.org/x/crypto/pbkdf2"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)

const (
	availableCharacters                      = "!#$&()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_abcdefghijklmnopqrstuvwxyz{|}~"
	availableNumberAndLetters                = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	availableNumberAndLowercaseLetters       = "0123456789abcdefghijklmnopqrstuvwxyz"
	availableCharactersLength                = len(availableCharacters)
	availableNumberAndLettersLength          = len(availableNumberAndLetters)
	availableNumberAndLowercaseLettersLength = len(availableNumberAndLowercaseLetters)
)


func SubString(str string, start int, length int) string {
	chars := []rune(str)
	realLength := len(chars)
	end := 0

	if start < 0 {
		start = realLength + start
	}

	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}

	if start > realLength {
		start = realLength
	}

	if end < 0 {
		end = 0
	}

	if end > realLength {
		end = realLength
	}

	return string(chars[start:end])
}


func ContainsAnyString(s string, substrs []string) bool {
	for i := 0; i < len(substrs); i++ {
		if strings.Index(s, substrs[i]) >= 0 {
			return true
		}
	}

	return false
}


func GetFirstLowerCharString(s string) string {
	if s == "" {
		return s
	}

	chars := []rune(s)

	if unicode.IsLower(chars[0]) {
		return s
	}

	chars[0] = unicode.ToLower(chars[0])
	return string(chars)
}


func ContainsOnlyOneRune(s string, r rune) bool {
	if len(s) < 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if rune(s[i]) != r {
			return false
		}
	}

	return true
}


func GetRandomString(n int) (string, error) {
	var result = make([]byte, n)

	for i := 0; i < n; i++ {
		index, err := GetRandomInteger(availableCharactersLength)

		if err != nil {
			return "", err
		}

		result[i] = availableCharacters[index]
	}

	return string(result), nil
}


func GetRandomNumberOrLetter(n int) (string, error) {
	var result = make([]byte, n)

	for i := 0; i < n; i++ {
		index, err := GetRandomInteger(availableNumberAndLettersLength)

		if err != nil {
			return "", err
		}

		result[i] = availableNumberAndLetters[index]
	}

	return string(result), nil
}


func GetRandomNumberOrLowercaseLetter(n int) (string, error) {
	var result = make([]byte, n)

	for i := 0; i < n; i++ {
		index, err := GetRandomInteger(availableNumberAndLowercaseLettersLength)

		if err != nil {
			return "", err
		}

		result[i] = availableNumberAndLowercaseLetters[index]
	}

	return string(result), nil
}


func MD5Encode(data []byte) []byte {
	m := md5.New()
	m.Write(data)
	return m.Sum(nil)
}


func MD5EncodeToString(data []byte) string {
	hash := MD5Encode(data)
	return hex.EncodeToString(hash)
}


func AESGCMEncrypt(key []byte, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())

	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plainText, nil)
	result := append(nonce, ciphertext...)

	return result, nil
}


func AESGCMDecrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()

	if len(ciphertext)-nonceSize <= 0 {
		return nil, errs.ErrCiphertextInvalid
	}

	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]

	plainText, err := aesgcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		return nil, err
	}

	return plainText, nil
}


func EncodePassword(password string, salt string) string {
	encodedPassword := pbkdf2.Key([]byte(password), []byte(salt), 10000, 48, sha256.New) 
	return strings.TrimRight(base64.StdEncoding.EncodeToString(encodedPassword), "=")
}


func EncryptSecret(secret string, key string) (string, error) {
	encryptedSecret, err := AESGCMEncrypt(MD5Encode([]byte(key)), []byte(secret)) 

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encryptedSecret), nil
}


func DecryptSecret(encyptedSecret string, key string) (string, error) {
	encyptedData, err := base64.StdEncoding.DecodeString(encyptedSecret)

	if err != nil {
		return "", err
	}

	secret, err := AESGCMDecrypt(MD5Encode([]byte(key)), []byte(encyptedData))

	if err != nil {
		return "", err
	}

	return string(secret), nil
}
