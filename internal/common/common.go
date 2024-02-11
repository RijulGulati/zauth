package common

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/rijulgulati/zauth/internal/zauth"
	"golang.org/x/term"
)

type ZAuthCommon struct{}

type PasswordReader interface {
	ReadPassword() (string, error)
}

type UserInputReader interface {
	UserInput() (string, error)
}

type ZAuthCommonComp interface {
	PasswordReader
	UserInputReader
}

// GetFileName generates and returns zauth filename for type t.
// Format: zauth-{t}-{Current UnixTime}.json?{.enc}
func GetFileName(t string, enc bool) string {
	name := "zauth-"
	if t != "" {
		name += t
	}

	name += fmt.Sprintf("-%v", time.Now().Unix())
	name += ".json"

	if enc {
		name += ".enc"
	}

	return name
}

// WriteFile writes d bytes to file f.
// If file's directory does it exist, it is created.
// If file exists, it is overwritten.
// Returns write location and any errors occured
func WriteFile(d []byte, f string) (string, error) {
	err := os.MkdirAll(zauth.ZAuthJsonDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	fl := filepath.Join(zauth.ZAuthJsonDir, f)
	return fl, os.WriteFile(fl, d, 0644)
}

// WriteZAuthJson writes ZAuth array objects as json to zauth.json file.
// File is overwritten with new content if ow is true. Else entries in z are appended to existing file.
func WriteZAuthJson(z []zauth.ZAuth, ow bool) error {
	err := os.MkdirAll(zauth.ZAuthJsonDir, os.ModePerm)
	if err != nil {
		return err
	}

	if ow {
		b, err := json.Marshal(z)
		if err != nil {
			return err
		}

		return os.WriteFile(zauth.ZAuthJson, b, 0644)
	} else {

		zj, err := ReadZAuthJson()
		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		}

		zj = append(zj, z...)
		b, err := json.Marshal(zj)
		if err != nil {
			return err
		}

		return os.WriteFile(zauth.ZAuthJson, b, 0644)
	}
}

// ReadZAuthJson reads zauth.json file and returns it's equivalent ZAuth array object
func ReadZAuthJson() ([]zauth.ZAuth, error) {
	b, err := os.ReadFile(zauth.ZAuthJson)
	if err != nil {
		return nil, err
	}

	z := make([]zauth.ZAuth, 0)
	err = json.Unmarshal(b, &z)
	if err != nil {
		return nil, err
	}

	return z, nil
}

func (zc *ZAuthCommon) ReadPassword() (string, error) {
	pass, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		msg := fmt.Sprintf("error reading password: %v", err)
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
		return "", fmt.Errorf(msg)
	}
	return string(pass), nil
}

func (zc *ZAuthCommon) UserInput() (string, error) {
	rd := bufio.NewReader(os.Stdin)
	return rd.ReadString('\n')
}
