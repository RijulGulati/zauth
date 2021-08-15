package otp

import (
	"strings"
	"time"

	"github.com/grijul/otpgen"
	"github.com/grijul/zauth/internal/zauth"
)

// GenerateOTP generates TOTP/HOTP code and returns pointer to ZAuthOtp object and any errors encountered.
func GenerateOTP(z *zauth.ZAuth) (*zauth.ZAuthOtp, error) {
	t := strings.ToUpper(z.Type)
	if t == "TOTP" {
		return generateTOTP(z)

	} else {
		return generateHOTP(z)
	}
}

// generateTOTP generates TOTP code and returns pointer to ZAuthOtp object and any errors encountered.
func generateTOTP(z *zauth.ZAuth) (*zauth.ZAuthOtp, error) {
	tm := time.Now().Unix()
	t := &otpgen.TOTP{
		Secret:    z.Secret,
		Digits:    z.Digits,
		Algorithm: z.Algorithm,
		Period:    z.Period,
		UnixTime:  tm,
	}

	o, err := t.Generate()
	return &zauth.ZAuthOtp{
		Otp:       o,
		Remaining: z.Period - (tm % z.Period),
	}, err
}

// generateHOTP generates HOTP code and returns pointer to ZAuthOtp object and any errors encountered.
func generateHOTP(z *zauth.ZAuth) (*zauth.ZAuthOtp, error) {
	h := otpgen.HOTP{
		Secret:  z.Secret,
		Digits:  z.Digits,
		Counter: z.Counter,
	}

	o, err := h.Generate()
	return &zauth.ZAuthOtp{
		Otp:       o,
		Remaining: 0,
	}, err
}
