package oauthurl

import (
	"fmt"
	"testing"
)

func TestParseUrl(t *testing.T) {
	fmt.Println("Testing ParseUrl()")
	totpUrl := "otpauth://totp/SomeOrg:test@example.com?secret=NBSWY3DPO5XXE3DEBIFA====&issuer=SomeOrg&algorithm=SHA1&digits=6&period=30"
	z, err := Parse(totpUrl)
	if err != nil {
		t.Fatal(err)
	}

	if z.Issuer != "SomeOrg" {
		t.Fatal("incorrect issuer: ", z.Issuer)
	}

	if z.Type != "totp" {
		t.Fatal("incorrect type: ", z.Type)
	}

	if z.Label != "SomeOrg:test@example.com" {
		t.Fatal("incorrect label: ", z.Label)
	}

	if z.Secret != "NBSWY3DPO5XXE3DEBIFA====" {
		t.Fatal("incorrect secret: ", z.Secret)
	}

	if z.Algorithm != "SHA1" {
		t.Fatal("incorrect algorithm: ", z.Algorithm)
	}

	if z.Digits != 6 {
		t.Fatal("incorrect digits: ", z.Digits)
	}

	if z.Period != 30 {
		t.Fatal("incorrect period: ", z.Period)
	}

	hotpUrl := "otpauth://hotp/SomeOrg:test@example.com?secret=NBSWY3DPO5XXE3DEBIFA====&issuer=SomeOrg&algorithm=SHA1&digits=6&counter=20"

	z, err = Parse(hotpUrl)
	if err != nil {
		t.Fatal(err)
	}

	if z.Type != "hotp" {
		t.Fatal("incorrect type: ", z.Type)
	}

	if z.Counter != 20 {
		t.Fatal("incorrect counter: ", z.Counter)
	}
}
