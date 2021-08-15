package test

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

var TestDir = getTestDir()
var TestAndotpAccountsJson = path.Join(TestDir, "andotp_accounts.json")
var TestAndotpAccountsJsonEnc = path.Join(TestDir, "andotp_accounts.json.enc")
var TestZAuthJsonDir = path.Join(os.TempDir(), "zauth")
var TestZAuthJson = path.Join(TestZAuthJsonDir, "zauth.json")
var AndOtpAccountsEncPassword = "testpass"

func getTestDir() string {
	_, fn, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprint(os.Stderr, "Unable to get test directory. No caller info")
		os.Exit(1)
	}

	return path.Dir(fn)
}

func RemoveTestFiles() {
	os.RemoveAll(TestZAuthJsonDir)
}
