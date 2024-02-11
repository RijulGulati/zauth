package otp

import (
	"testing"

	"github.com/rijulgulati/zauth/internal/oauthurl"
)

func TestGenerateOTP(t *testing.T) {
	totpUrl := "otpauth://totp/SomeOrg:test@example.com?secret=NBSWY3DPO5XXE3DEBI======&issuer=SomeOrg&algorithm=SHA1&digits=6&period=30"
	z, err := oauthurl.Parse(totpUrl)
	if err != nil {
		t.Fatal(err)
	}

	_, err = GenerateOTP(z)
	if err != nil {
		t.Fatal(err)
	}

	hotpUrl := "otpauth://hotp/SomeOrg:test@example.com?secret=NBSWY3DPO5XXE3DEBI======&issuer=SomeOrg&algorithm=SHA1&digits=6&counter=5"
	z, err = oauthurl.Parse(hotpUrl)
	if err != nil {
		t.Fatal(err)
	}

	_, err = GenerateOTP(z)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateTOTP(t *testing.T) {
	totpUrl := "otpauth://totp/SomeOrg:test@example.com?secret=NBSWY3DPO5XXE3DEBI======&issuer=SomeOrg&algorithm=SHA1&digits=6&period=30"
	z, err := oauthurl.Parse(totpUrl)
	if err != nil {
		t.Fatal(err)
	}

	zo, err := generateTOTP(z)
	if err != nil {
		t.Fatal(err)
	}

	if zo.Otp == "" {
		t.Fatal("received empty otp.")
	}

}

func TestGenerateHOTP(t *testing.T) {
	hotpUrl := "otpauth://hotp/SomeOrg:test@example.com?secret=NBSWY3DPO5XXE3DEBI======&issuer=SomeOrg&algorithm=SHA1&digits=6&counter=5"
	z, err := oauthurl.Parse(hotpUrl)
	if err != nil {
		t.Fatal(err)
	}

	zo, err := generateHOTP(z)
	if err != nil {
		t.Fatal(err)
	}

	if zo.Otp == "" {
		t.Fatal("received empty otp.")
	}

}
