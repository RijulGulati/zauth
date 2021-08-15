package third_party

import (
	"fmt"

	"github.com/grijul/zauth/internal/zauth"
	"github.com/grijul/zauth/third_party/andotp"
)

var SupportedImportTypes = []string{"andotp"}

type ImportFile interface {
	// Import imports encrypted/decrypted file f and returns ZAuth object and any errors encountered.
	// Password pwd is used for decrypting file if file is encrypted.
	// If ow is true, existing zauth.json file is overwritten. Else entries are appended.
	Import(f string, p string, ow bool) ([]zauth.ZAuth, error)
}

// NewImportFile returns ImportFile interface implemented by import type t
func NewImportFile(t *string) (ImportFile, error) {
	switch *t {
	case "andotp":
		{
			return andotp.NewAndOtp(), nil
		}

	default:
		{
			return nil, fmt.Errorf("%s", "file type not supported")
		}
	}
}
