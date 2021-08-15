package oauthurl

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/grijul/zauth/internal/zauth"
)

func Parse(u string) (*zauth.ZAuth, error) {

	ourl, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	if ourl.Scheme != "otpauth" {
		return nil, fmt.Errorf("invalid url")
	}

	za := &zauth.ZAuth{}
	za.Type = ourl.Host
	za.Label = ourl.Path[1:]

	values := ourl.Query()
	za.Secret = values.Get("secret")
	za.Algorithm = values.Get("algorithm")

	za.Digits, err = strconv.Atoi(values.Get("digits"))
	if err != nil {
		return nil, fmt.Errorf("invalid parameter: digits - " + err.Error())
	}

	za.Issuer = values.Get("issuer")

	if za.Type == "totp" {
		za.Period, err = strconv.ParseInt(values.Get("period"), 10, 0)
		if err != nil {
			return nil, fmt.Errorf("invalid parameter: period - " + err.Error())
		}

	} else {
		za.Counter, err = strconv.ParseInt(values.Get("counter"), 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid parameter: counter - " + err.Error())
		}
	}

	return za, nil
}
