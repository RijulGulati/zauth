package andotp

import (
	"os"
	"testing"

	"github.com/grijul/zauth/internal/zauth"
	"github.com/grijul/zauth/test"
)

func init() {
	zauth.ZAuthJson = test.TestZAuthJson
	zauth.ZAuthJsonDir = test.TestZAuthJsonDir
}

func TestImport(t *testing.T) {
	o := NewAndOtp()
	test.RemoveTestFiles()

	// check for unencrypted file
	_, err := o.Import(test.TestAndotpAccountsJson, "", true)
	if err != nil {
		test.RemoveTestFiles()
		t.Fatal(err)
	}

	// check for encrypted file
	_, err = o.Import(test.TestAndotpAccountsJsonEnc, test.AndOtpAccountsEncPassword, false)
	if err != nil {
		test.RemoveTestFiles()
		t.Fatal(err)
	}

}

func TestExport(t *testing.T) {
	o := NewAndOtp()

	// check for unencrypted file
	f, err := o.Export("")
	if err != nil {
		t.Fatal(err)
		test.RemoveTestFiles()
	}

	os.Remove(*f)

	// check for encrypted file
	f, err = o.Export("testpass")
	if err != nil {
		t.Fatal(err)
		test.RemoveTestFiles()
	}

	test.RemoveTestFiles()
	os.Remove(*f)
}
