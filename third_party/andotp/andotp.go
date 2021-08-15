package andotp

import (
	"encoding/json"
	"os"

	ga "github.com/grijul/go-andotp/andotp"
	"github.com/grijul/zauth/internal/common"
	"github.com/grijul/zauth/internal/zauth"
)

type andotpNode struct {
	Secret         string        `json:"secret"`
	Issuer         string        `json:"issuer"`
	Label          string        `json:"label"`
	Digits         int           `json:"digits"`
	Type           string        `json:"type"`
	Algorithm      string        `json:"alogrithm"`
	Thumbnail      string        `json:"thumbnail"`
	Last_used      int64         `json:"last_used"`
	Used_frequency int           `json:"used_frequency"`
	Period         int           `json:"period"`
	Tags           []interface{} `json:"tags"`
}

type AndOtpImportExport struct{}

func NewAndOtp() AndOtpImportExport {
	return AndOtpImportExport{}
}

// Import imports andOTP encrypted/decrypted file f and returns ZAuth object and any errors encountered.
// Password pwd is used for decrypting file.
// If ow is true, existing zauth.json file is overwritten. Else entries are appended.
func (a AndOtpImportExport) Import(f string, pwd string, ow bool) ([]zauth.ZAuth, error) {

	an := make([]andotpNode, 0)
	zl := make([]zauth.ZAuth, 0)
	z := zauth.ZAuth{}

	fc, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fc, &an)
	if err != nil {
		if pwd != "" {
			fc, err = ga.Decrypt(fc, pwd)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(fc, &an)
			if err != nil {
				return nil, err
			}

		} else {
			return nil, err
		}
	}

	for _, node := range an {
		z.Algorithm = node.Algorithm
		z.Digits = node.Digits
		z.Issuer = node.Issuer
		z.Label = node.Label
		z.Period = int64(node.Period)
		z.Secret = node.Secret
		z.Type = node.Type
		if z.Misc == nil {
			z.Misc = make(map[string]interface{})
		}
		z.Misc["thumbnail"] = node.Thumbnail
		z.Misc["tags"] = node.Tags
		z.Misc["last_used"] = node.Last_used
		z.Misc["used_frequency"] = node.Used_frequency

		zl = append(zl, z)
		z = zauth.ZAuth{}
	}

	err = common.WriteZAuthJson(zl, ow)
	if err != nil {
		return nil, err
	}

	return zl, nil
}

// Export exports zauth's entries to andOTP-compatible json file.
// Output file is encrypted with password pwd if provided.
func (a AndOtpImportExport) Export(pwd string) (*string, error) {
	var d []byte
	enc := false
	al := make([]andotpNode, 0)
	an := andotpNode{}

	zj, err := common.ReadZAuthJson()
	if err != nil {
		return nil, err
	}

	for _, z := range zj {
		an.Algorithm = z.Algorithm
		an.Digits = z.Digits
		an.Issuer = z.Issuer
		an.Label = z.Label
		an.Period = int(z.Period)
		an.Secret = z.Secret
		an.Type = z.Type
		if z.Misc == nil {
			an.Thumbnail = "Default"
		} else {
			if z.Misc["thumbnail"] != nil {
				an.Thumbnail = z.Misc["thumbnail"].(string)
			} else {
				an.Thumbnail = "Default"
			}
			if z.Misc["tags"] != nil {
				an.Tags = z.Misc["tags"].([]interface{})
			}
			if z.Misc["last_used"] != nil {
				an.Last_used = int64(z.Misc["last_used"].(float64))
			}
			if z.Misc["used_frequency"] != nil {
				an.Used_frequency = int(z.Misc["used_frequency"].(float64))
			}
		}

		al = append(al, an)
		an = andotpNode{}
	}

	d, err = json.Marshal(al)
	if err != nil {
		return nil, err
	}

	if pwd != "" {
		enc = true
		d, err = ga.Encrypt(d, pwd)
		if err != nil {
			return nil, err
		}
	}

	fl, err := common.WriteFile(d, common.GetFileName("andotp", enc))
	if err != nil {
		return nil, err
	}

	return &fl, nil
}
