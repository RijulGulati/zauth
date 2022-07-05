package args

import (
	"fmt"
	"os"
	"testing"

	"github.com/grijul/zauth/internal/zauth"
	"github.com/grijul/zauth/test"
)

type zauthCommonTest struct{}

var zc *zauthCommonTest

var isErroredPassword bool
var isEmptyPassword bool

func init() {
	zauth.ZAuthJson = test.TestZAuthJson
	zauth.ZAuthJsonDir = test.TestZAuthJsonDir
}

func TestParseNoSubcommandArgs(t *testing.T) {
	defer test.RemoveTestFiles()
	test.RemoveTestFiles()

	// error: no file found
	os.Args = []string{"zauth"}
	err := ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when no zauth json is found")
	}

	// import sample file to create zauth json file. The file will then be exported
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson)}
	ParseArgs(zc)

	// no error
	os.Args = []string{"zauth"}
	err = ParseArgs(zc)
	if err != nil {
		t.Fatal(err)
	}

	// invalid subcommand
	// error: no file found
	os.Args = []string{"zauth", "testxyz"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when invalid subcommand is given")
	}

}

func TestParseImportArgs(t *testing.T) {
	defer test.RemoveTestFiles()

	// valid arguments:overwrite=true
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson), "-overwrite"}
	err := ParseArgs(zc)
	if err != nil {
		t.Fatal(err)
	}

	//valid arguments:overwrite=false
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson)}
	err = ParseArgs(zc)
	if err != nil {
		test.RemoveTestFiles()
		t.Fatal(err)
	}

	//invalid arguments:empty type
	os.Args = []string{"zauth", "import", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson), "-overwrite=true"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when import type is empty")
	}

	//invalid arguments:empty file
	os.Args = []string{"zauth", "import", "-type=andotp"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when import file is empty")
	}

	//invalid arguments:empty file
	os.Args = []string{"zauth", "import", "-type=andotp"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when import file is empty")
	}

	//invalid arguments:invalid file
	os.Args = []string{"zauth", "import", "-type=andotp", "-file=test"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when import file is invalid")
	}

	// invalid type
	zauth.ZAuthJson = test.TestZAuthJson
	os.Args = []string{"zauth", "import", "-type=invalid", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson)}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when type is invalid")
	}

	// encrypt
	// invalid arguments: wrong file
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson), "-decrypt"}
	err = ParseArgs(zc)
	if err != nil {
		t.Fatal(err)
	}

	// invalid arguments: correct file
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJsonEnc), "-decrypt"}
	err = ParseArgs(zc)
	if err != nil {
		t.Fatal(err)
	}

	resetPasswordFlags()

	// errored password
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson), "-decrypt"}
	isErroredPassword = true
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when read password is errored")
	}

	resetPasswordFlags()

	// empty password
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson), "-decrypt"}
	isEmptyPassword = true
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when password is empty")
	}

	resetPasswordFlags()

}

func TestParseExportArgs(t *testing.T) {
	defer test.RemoveTestFiles()
	test.RemoveTestFiles()

	// zauth file does not exist
	os.Args = []string{"zauth", "export", "-type=andotp"}
	err := ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when zauth file does not exist")
	}

	// import sample file to create zauth json file. The file will then be exported
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson)}
	ParseArgs(zc)

	// valid arguments
	os.Args = []string{"zauth", "export", "-type=andotp"}
	err = ParseArgs(zc)
	if err != nil {
		t.Fatal(err)
	}

	// invalid type
	os.Args = []string{"zauth", "export", "-type=invalid"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when type is invalid")
	}

	// empty type
	os.Args = []string{"zauth", "export"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when type is empty")
	}

	// encrypt
	// valid arguments
	os.Args = []string{"zauth", "export", "-type=andotp", "-encrypt"}
	err = ParseArgs(zc)
	if err != nil {
		t.Fatal(err)
	}

	resetPasswordFlags()

	// errored password
	os.Args = []string{"zauth", "export", "-type=andotp", "-encrypt"}
	isErroredPassword = true
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when read password is errored")
	}

	resetPasswordFlags()

	// empty password
	os.Args = []string{"zauth", "export", "-type=andotp", "-encrypt"}
	isEmptyPassword = true
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when password is empty")
	}

	resetPasswordFlags()
}

func TestParseEntryArgs(t *testing.T) {
	defer test.RemoveTestFiles()
	test.RemoveTestFiles()

	//entry list

	// no zauth.json present
	os.Args = []string{"zauth", "entry", "-list"}
	err := ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail when zauth file does not exist")
	}

	// import sample file to create zauth json file. The file will then be exported
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson)}
	ParseArgs(zc)

	// no zauth.json present
	os.Args = []string{"zauth", "entry", "-list"}
	err = ParseArgs(zc)
	if err != nil {
		t.Fatal(err)
	}

	//entry edit
	os.Args = []string{"zauth", "entry", "-edit"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail since functionality is not yet implemented")
	}

	//entry delete
	os.Args = []string{"zauth", "entry", "-delete"}
	err = ParseArgs(zc)
	if err == nil {
		t.Fatal("expected test to fail since functionality is not yet implemented")
	}
}
func TestParseOtpArgs(t *testing.T) {
	defer test.RemoveTestFiles()

	// import sample file to create zauth json file. The file will then be exported
	os.Args = []string{"zauth", "import", "-type=andotp", fmt.Sprintf("-file=%s", test.TestAndotpAccountsJson)}
	ParseArgs(zc)
	//otp get
	zauth.ZAuthJson = test.TestZAuthJson
	os.Args = []string{"zauth", "otp", "-get", "Some Org"}
	err := ParseArgs(zc)
	if err != nil {
		t.Fatal(err)
	}

}

func (*zauthCommonTest) ReadPassword() (string, error) {
	if isEmptyPassword {
		return "", nil
	}

	if isErroredPassword {
		return "", fmt.Errorf("password error")
	}

	return test.AndOtpAccountsEncPassword, nil
}

func resetPasswordFlags() {
	isEmptyPassword = false
	isErroredPassword = false
}

func (*zauthCommonTest) UserInput() (string, error) {
	return "test", nil
}
