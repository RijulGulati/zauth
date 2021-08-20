package test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var TestDir = getTestDir()
var TestAndotpAccountsJson = filepath.Join(TestDir, "andotp_accounts.json")
var TestAndotpAccountsJsonEnc = filepath.Join(TestDir, "andotp_accounts.json.enc")
var TestZAuthJsonDir = filepath.Join(os.TempDir(), "zauth")
var TestZAuthJson = filepath.Join(TestZAuthJsonDir, "zauth.json")
var AndOtpAccountsEncPassword = "testpass"

func getTestDir() string {
	_, fn, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprint(os.Stderr, "Unable to get test directory. No caller info")
		os.Exit(1)
	}

	return filepath.Dir(fn)
}

func RemoveTestFiles() {
	os.RemoveAll(TestZAuthJsonDir)
}
