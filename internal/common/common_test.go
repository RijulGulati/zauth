package common

import (
	"fmt"
	"testing"
	"time"

	"github.com/grijul/zauth/internal/zauth"
	"github.com/grijul/zauth/test"
)

func init() {
	zauth.ZAuthJson = test.TestZAuthJson
	zauth.ZAuthJsonDir = test.TestZAuthJsonDir
}

func TestGetFileName(t *testing.T) {
	n := GetFileName("test", false)
	r := fmt.Sprintf("zauth-test-%v.json", time.Now().Unix())
	if n != r {
		t.Fatal("unextected output: ", n)
	}

	n = GetFileName("test", true)
	r = fmt.Sprintf("zauth-test-%v.json.enc", time.Now().Unix())
	if n != r {
		t.Fatal("unextected output: ", n)
	}
}

func TestWriteZAuthJson(t *testing.T) {
	test.RemoveTestFiles()
	z := zauth.ZAuth{
		Secret: "test",
		Label:  "test",
	}

	err := WriteZAuthJson([]zauth.ZAuth{z}, false)
	if err != nil {
		t.Fatal(err)
		test.RemoveTestFiles()
	}

	z2 := zauth.ZAuth{
		Secret: "test2",
		Label:  "test2",
	}

	err = WriteZAuthJson([]zauth.ZAuth{z2}, false)
	if err != nil {
		t.Fatal(err)
		test.RemoveTestFiles()
	}
}

func TestReadZAuthJson(t *testing.T) {
	z, err := ReadZAuthJson()
	if err != nil {
		t.Fatal(err)
		test.RemoveTestFiles()
	}
	if len(z) != 2 {
		t.Fatal("expected entries count: 2. received: ", len(z))
		test.RemoveTestFiles()
	}

	test.RemoveTestFiles()
}
