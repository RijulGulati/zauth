package third_party

import (
	"fmt"

	"github.com/grijul/zauth/third_party/andotp"
)

var SupportedExportTypes = []string{"andotp"}

type ExportFile interface {
	// Export exports zauth entries to file. Exported file may be encrypted with password p. If p is empty, unencrypted file is exported.
	// Returns export file path any errors encountered
	Export(p string) (*string, error)
}

// NewExportFile returns ExportFile interface implemented by import type t
func NewExportFile(t *string) (ExportFile, error) {
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
