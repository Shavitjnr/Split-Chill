package core

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type IPPattern struct {
	Pattern string
	regex   *regexp.Regexp
}


func (p *IPPattern) Match(ip string) bool {
	if p.regex == nil {
		return false
	}

	return p.regex.MatchString(ip)
}


func (p *IPPattern) GobEncode() ([]byte, error) {
	return []byte(p.Pattern), nil
}


func (p *IPPattern) GobDecode(data []byte) error {
	pattern := string(data)

	if pattern == "" {
		p.Pattern = ""
		p.regex = nil
		return nil
	}

	newPattern, err := ParseIPPattern(pattern)

	if err != nil {
		return err
	}

	p.Pattern = newPattern.Pattern
	p.regex = newPattern.regex
	return nil
}


func ParseIPPattern(ipPattern string) (*IPPattern, error) {
	if ipPattern == "" {
		return nil, nil
	}

	hasDot := false
	hasSemicolon := false

	for i := 0; i < len(ipPattern); i++ {
		ch := rune(ipPattern[i])

		if ch == '.' { 
			if hasSemicolon {
				return nil, errs.ErrInvalidIpAddressPattern
			}
			hasDot = true
		} else if ch == ':' { 
			if hasDot {
				return nil, errs.ErrInvalidIpAddressPattern
			}
			hasSemicolon = true
		}
	}

	if hasDot {
		return ParseIPv4Pattern(ipPattern)
	} else if hasSemicolon {
		return ParseIPv6Pattern(ipPattern)
	} else {
		return nil, errs.ErrInvalidIpAddressPattern
	}
}


func ParseIPv4Pattern(ipPattern string) (*IPPattern, error) {
	items := strings.Split(ipPattern, ".")

	if len(items) != 4 {
		return nil, errs.ErrInvalidIpAddressPattern
	}

	regexBuilder := strings.Builder{}
	regexBuilder.WriteRune('^')

	for i := 0; i < len(items); i++ {
		item := strings.TrimSpace(items[i])

		if item == "*" {
			regexBuilder.WriteString("[0-9]{1,3}")
		} else if item == "" {
			return nil, errs.ErrInvalidIpAddressPattern
		} else {
			num, err := strconv.Atoi(item)

			if err != nil || num < 0 || num > 255 {
				return nil, errs.ErrInvalidIpAddressPattern
			}

			regexBuilder.WriteString(item)
		}

		if i < len(items)-1 {
			regexBuilder.WriteRune('\\')
			regexBuilder.WriteRune('.')
		}
	}

	regexBuilder.WriteRune('$')
	regex, err := regexp.Compile(regexBuilder.String())

	if err != nil {
		return nil, errs.ErrInvalidIpAddressPattern
	}

	return &IPPattern{
		Pattern: ipPattern,
		regex:   regex,
	}, nil
}


func ParseIPv6Pattern(ipPattern string) (*IPPattern, error) {
	items := strings.Split(ipPattern, ":")

	if len(items) < 2 || len(items) > 8 {
		return nil, errs.ErrInvalidIpAddressPattern
	}

	regexBuilder := strings.Builder{}
	regexBuilder.WriteRune('^')

	for i := 0; i < len(items); i++ {
		item := strings.TrimSpace(items[i])

		if item == "*" {
			regexBuilder.WriteString("[0-9a-fA-F]{1,4}")
		} else if i < len(items)-1 && item == "" {
			
		} else {
			num, err := strconv.ParseInt(item, 16, 32)

			if err != nil || num < 0 || num > 0xFFFF {
				return nil, errs.ErrInvalidIpAddressPattern
			}

			regexBuilder.WriteString(item)
		}

		if i < len(items)-1 {
			regexBuilder.WriteRune(':')
		}
	}

	regexBuilder.WriteRune('$')
	regex, err := regexp.Compile(regexBuilder.String())

	if err != nil {
		return nil, errs.ErrInvalidIpAddressPattern
	}

	return &IPPattern{
		Pattern: ipPattern,
		regex:   regex,
	}, nil
}
