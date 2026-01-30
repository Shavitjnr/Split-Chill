package core

import (
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)

const webContextRequestIdFieldKey = "REQUEST_ID"
const webContextTextualTokenFieldKey = "TOKEN_STRING"
const webContextTokenClaimsFieldKey = "TOKEN_CLAIMS"
const webContextTokenContextFieldKey = "TOKEN_CONTEXT"
const webContextResponseErrorFieldKey = "RESPONSE_ERROR"


const AcceptLanguageHeaderName = "Accept-Language"


const RemoteClientPortHeader = "X-Real-Port"


const ClientTimezoneOffsetHeaderName = "X-Timezone-Offset"


const ClientTimezoneNameHeaderName = "X-Timezone-Name"

const tokenHeaderName = "Authorization"
const tokenHeaderValuePrefix = "bearer "
const tokenQueryStringParam = "token"
const tokenCookieParam = "ebk_auth_token"


type WebContext struct {
	*gin.Context
	
}

func (c *WebContext) ClientPort() uint16 {
	remotePort := c.GetHeader(RemoteClientPortHeader)

	if remotePort != "" {
		remotePortNum, err := strconv.ParseInt(remotePort, 10, 32)

		if err == nil {
			return uint16(remotePortNum)
		}
	}

	if c.Request == nil {
		return 0
	}

	_, remotePort, err := net.SplitHostPort(c.Request.RemoteAddr)

	if err != nil {
		return 0
	}

	remotePortNum, err := strconv.ParseInt(remotePort, 10, 32)

	if err != nil {
		return 0
	}

	return uint16(remotePortNum)
}


func (c *WebContext) SetContextId(requestId string) {
	c.Set(webContextRequestIdFieldKey, requestId)
}


func (c *WebContext) GetContextId() string {
	requestId, exists := c.Get(webContextRequestIdFieldKey)

	if !exists {
		return ""
	}

	return requestId.(string)
}


func (c *WebContext) SetTextualToken(token string) {
	c.Set(webContextTextualTokenFieldKey, token)
}


func (c *WebContext) GetTextualToken() string {
	token, exists := c.Get(webContextTextualTokenFieldKey)

	if !exists {
		return ""
	}

	return token.(string)
}


func (c *WebContext) SetTokenClaims(claims *UserTokenClaims) {
	c.Set(webContextTokenClaimsFieldKey, claims)
}


func (c *WebContext) GetTokenClaims() *UserTokenClaims {
	claims, exists := c.Get(webContextTokenClaimsFieldKey)

	if !exists {
		return nil
	}

	return claims.(*UserTokenClaims)
}


func (c *WebContext) SetTokenContext(context string) {
	c.Set(webContextTokenContextFieldKey, context)
}


func (c *WebContext) GetTokenContext() string {
	context, exists := c.Get(webContextTokenContextFieldKey)

	if !exists {
		return ""
	}

	return context.(string)
}


func (c *WebContext) GetCurrentUid() int64 {
	claims := c.GetTokenClaims()

	if claims == nil {
		return 0
	}

	return claims.Uid
}


func (c *WebContext) GetTokenStringFromHeader() string {
	tokenHeader := c.GetHeader(tokenHeaderName)

	if len(tokenHeader) < 7 || !strings.EqualFold(tokenHeader[:7], tokenHeaderValuePrefix) {
		return ""
	}

	return tokenHeader[7:]
}


func (c *WebContext) GetTokenStringFromQueryString() string {
	return c.Query(tokenQueryStringParam)
}


func (c *WebContext) GetTokenStringFromCookie() string {
	tokenCookie, err := c.Cookie(tokenCookieParam)

	if err != nil {
		return ""
	}

	return tokenCookie
}

func (c *WebContext) SetTokenStringToCookie(token string, tokenExpiredTime int, path string) {
	if token != "" {
		c.SetCookie(tokenCookieParam, token, tokenExpiredTime, path, "", false, true)
	} else {
		c.SetCookie(tokenCookieParam, "", -1, path, "", false, true)
	}
}


func (c *WebContext) GetClientLocale() string {
	value := c.GetHeader(AcceptLanguageHeaderName)

	return value
}

func (c *WebContext) GetClientTimezone() (*time.Location, error) {
	timezoneName := c.getClientTimezoneName()

	if timezoneName != "" {
		location, err := time.LoadLocation(timezoneName)

		if err == nil && location != nil {
			return location, nil
		}
	}

	utcOffset, err := c.getClientTimezoneOffset()

	if err != nil {
		return nil, err
	}

	return time.FixedZone("Client Fixed Timezone", int(utcOffset)*60), nil
}


func (c *WebContext) SetResponseError(error *errs.Error) {
	c.Set(webContextResponseErrorFieldKey, error)
}


func (c *WebContext) GetResponseError() *errs.Error {
	err, exists := c.Get(webContextResponseErrorFieldKey)

	if !exists {
		return nil
	}

	return err.(*errs.Error)
}


func (c *WebContext) getClientTimezoneOffset() (int16, error) {
	value := c.GetHeader(ClientTimezoneOffsetHeaderName)
	offset, err := strconv.Atoi(value)

	if err != nil {
		return 0, err
	}

	return int16(offset), nil
}


func (c *WebContext) getClientTimezoneName() string {
	value := c.GetHeader(ClientTimezoneNameHeaderName)

	return value
}


func WrapWebContext(ginCtx *gin.Context) *WebContext {
	return &WebContext{
		Context: ginCtx,
	}
}
