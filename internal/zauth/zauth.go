package zauth

import (
	"os"
	"path"
)

type ZAuth struct {
	Secret    string                 `json:"secret"`
	Label     string                 `json:"label"`
	Issuer    string                 `json:"issuer"`
	Digits    int                    `json:"digits"`
	Algorithm string                 `json:"algorithm"`
	Counter   int64                  `json:"counter"`
	Period    int64                  `json:"period"`
	Type      string                 `json:"type"`
	Misc      map[string]interface{} `json:"misc"` // to store all other nodes (that are not part of ZAuth at time of import)
}

type ZAuthOtp struct {
	Otp       string
	Remaining int64
}

const DefaultDigits = 6
const DefaultType = "totp"
const DefaultCounter = 0
const DefaultAlgo = "sha1"
const DefaultPeriod = 30

var ZAuthJsonDir string = path.Join(os.Getenv("HOME"), ".zauth")
var ZAuthJson string = path.Join(ZAuthJsonDir, "zauth.json")
