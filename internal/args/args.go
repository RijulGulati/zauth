package args

import (
	"encoding/base32"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/grijul/zauth/internal/common"
	"github.com/grijul/zauth/internal/otp"
	"github.com/grijul/zauth/internal/zauth"
	"github.com/grijul/zauth/third_party"
	"github.com/rodaine/table"
)

type ZAuthArgsEntry struct{}

type SecretReader interface {
	ReadSecret(common.ZAuthCommonComp) (string, error)
}

type IssuerReader interface {
	ReadIssuer(common.ZAuthCommonComp) (string, error)
}

type IdentifierReader interface {
	ReadIdentifier(common.ZAuthCommonComp) (string, error)
}

type TypeReader interface {
	ReadType(common.ZAuthCommonComp) (string, error)
}

type DigitsReader interface {
	ReadDigits(common.ZAuthCommonComp) (int, error)
}

type AlgorithmReader interface {
	ReadAlgorithm(common.ZAuthCommonComp) (string, error)
}

type CounterReader interface {
	ReadCounter(common.ZAuthCommonComp) (int64, error)
}

type PeriodReader interface {
	ReadPeriod(common.ZAuthCommonComp) (int64, error)
}

type ZAuthArgsEntryInput interface {
	SecretReader
	PeriodReader
	CounterReader
	AlgorithmReader
	DigitsReader
	TypeReader
	IdentifierReader
	IssuerReader
}

const usage = `COMMANDS:
  entry			zauth entry operations (add/edit/delete/list) (see zauth entry --help)
  import		import file(s) to zauth (see zauth import --help)
  export		export zauth entries to file (see zauth export --help)`

func ParseArgs(zc common.ZAuthCommonComp) error {
	var msg string
	ze := &ZAuthArgsEntry{}

	// import cmd
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	importType := importCmd.String("type", "", "Import type (required)")
	importFile := importCmd.String("file", "", "Import file (required)")
	importFileOverwrite := importCmd.Bool("overwrite", false, "Overwrite existing entries with new entries (optional)")
	importFileDecrypt := importCmd.Bool("decrypt", false, "Decrypt import file with password (Enter password on prompt) (optional)")

	// export cmd
	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	exportType := exportCmd.String("type", "", "Export type (required)")
	exportFileEncrypt := exportCmd.Bool("encrypt", false, "Encrypt output file with password (Enter password on prompt) (optional)")

	// entry cmd
	entryCmd := flag.NewFlagSet("entry", flag.ExitOnError)
	entryNew := entryCmd.Bool("new", false, "Create new entry")
	entryEdit := entryCmd.Bool("edit", false, "Edit existing entry")
	entryDelete := entryCmd.Bool("delete", false, "Delete existing entry")
	entryList := entryCmd.Bool("list", false, "List all entries")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <command>\n\n%v\n", os.Args[0], usage)
	}

	flag.Parse()

	if len(os.Args) == 1 {
		return printZAuthOtpTable()
	} else {
		switch os.Args[1] {
		case "import":
			{
				var pwd string
				importCmd.Usage = func() {
					printUsage("import", importCmd)
					fmt.Fprintf(flag.CommandLine.Output(), "\nSupported import types: %v\n\n", strings.Join(third_party.SupportedImportTypes, ", "))
				}

				importCmd.Parse(os.Args[2:])

				if *importType == "" {
					msg = "import type cannot be empty"
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					fmt.Fprint(flag.CommandLine.Output(), "see -h for help\n")
					return fmt.Errorf(msg)
				}

				if *importFile == "" {
					msg = "import file cannot be empty"
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					fmt.Fprint(flag.CommandLine.Output(), "see -h for help\n")
					return fmt.Errorf(msg)
				}

				if *importFileDecrypt {
					fmt.Print("Password: ")
					p, err := zc.ReadPassword()
					if err != nil {
						msg = fmt.Sprintf("An error occured while capturing password: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}
					if p == "" {
						msg = "password cannot be empty"
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}

					pwd = p
				}

				tpi, err := third_party.NewImportFile(importType)
				if err != nil {
					msg = fmt.Sprintf("An error occured while importing file: %v", err)
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					return fmt.Errorf(msg)
				}
				_, err = tpi.Import(*importFile, pwd, *importFileOverwrite)
				if err != nil {
					msg = fmt.Sprintf("An error occured while importing file: %v", err)
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					return fmt.Errorf(msg)
				}

				fmt.Println("\n\n1 file imported successfully")
				return nil
			}
		case "export":
			{
				var pwd string
				exportCmd.Usage = func() {
					printUsage("export", exportCmd)
					fmt.Fprintf(flag.CommandLine.Output(), "\nSupported export types: %v\n\n", strings.Join(third_party.SupportedImportTypes, ", "))
				}

				exportCmd.Parse(os.Args[2:])

				if *exportType == "" {
					msg = "export type cannot be empty"
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					fmt.Fprint(flag.CommandLine.Output(), "see -h for help\n")
					return fmt.Errorf(msg)
				}

				if *exportFileEncrypt {
					fmt.Print("Password: ")
					p, err := zc.ReadPassword()
					if err != nil {
						msg = fmt.Sprintf("An error occured while capturing password: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}
					if p == "" {
						msg = "password cannot be empty"
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}
					pwd = p
				}

				tpe, err := third_party.NewExportFile(exportType)
				if err != nil {
					msg = fmt.Sprintf("An error occured while exporting file: %v", err)
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					return fmt.Errorf(msg)
				}

				fl, err := tpe.Export(pwd)
				if err != nil {
					msg = fmt.Sprintf("An error occured while exporting file: %v", err)
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					return fmt.Errorf(msg)
				}

				fmt.Println("\n\n1 file exported successfully")
				fmt.Println("export location: ", *fl)
				return nil
			}
		case "entry":
			{
				entryCmd.Usage = func() {
					printUsage("entry", entryCmd)
				}

				entryCmd.Parse(os.Args[2:])

				if *entryNew {
					z := &zauth.ZAuth{}
					fmt.Println("zauth new entry")
					fmt.Printf("-----------------------\n\n")

					sec, err := ze.ReadSecret(zc)
					if err != nil {
						msg = fmt.Sprintf("An error occured while reading secret: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}
					z.Secret = sec

					iss, err := ze.ReadIssuer(zc)
					if err != nil {
						msg = fmt.Sprintf("An error occured while reading issuer: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}
					z.Issuer = iss

					id, err := ze.ReadIdentifier(zc)
					if err != nil {
						msg = fmt.Sprintf("An error occured while reading identifier: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}
					z.Label = fmt.Sprintf("%s:%s", z.Issuer, id)

					tp, err := ze.ReadType(zc)
					if err != nil {
						msg = fmt.Sprintf("n error occured while reading type: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}
					z.Type = tp

					dt, err := ze.ReadDigits(zc)
					if err != nil {
						msg = fmt.Sprintf("An error occured while reading digits: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}
					z.Digits = dt

					if z.Type == "totp" {

						algo, err := ze.ReadAlgorithm(zc)
						if err != nil {
							msg = fmt.Sprintf("An error occured while reading algorithm: %v", err)
							fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
							return fmt.Errorf(msg)
						}
						z.Algorithm = algo

						pd, err := ze.ReadPeriod(zc)
						if err != nil {
							msg = fmt.Sprintf("An error occured while reading period: %v", err)
							fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
							return fmt.Errorf(msg)
						}
						z.Period = pd

					} else {
						z.Algorithm = "sha1"

						ctr, err := ze.ReadCounter(zc)
						if err != nil {
							msg = fmt.Sprintf("An error occured while reading counter: %v", err)
							fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
							return fmt.Errorf(msg)
						}
						z.Counter = ctr
					}

					err = common.WriteZAuthJson([]zauth.ZAuth{*z}, false)
					if err != nil {
						msg = fmt.Sprintf("An error occured while creating entry: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}

					fmt.Println("\n1 entry created successfully!")
					return nil

				} else if *entryList {
					lst, err := common.ReadZAuthJson()
					if err != nil {
						msg = fmt.Sprintf("An error occured while listing entries: %v", err)
						fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
						return fmt.Errorf(msg)
					}

					fmt.Println("zauth entries")
					fmt.Printf("-----------------------\n\n")
					for i, l := range lst {
						out := fmt.Sprintf("[%d] (%s) %s (%s)", i+1, l.Issuer, l.Label, strings.ToUpper(l.Type))
						fmt.Println(out)
					}
					return nil
				} else if *entryEdit {
					msg = "functionality is yet to be implemented"
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					return fmt.Errorf(msg)
				} else if *entryDelete {
					msg = "functionality is yet to be implemented"
					fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
					return fmt.Errorf(msg)
				}

				return nil
			}

		default:
			{
				msg = "invalid input.\nPlease see -h for available commands"
				fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
				return fmt.Errorf(msg)
			}

		}
	}
}

func (za *ZAuthArgsEntry) ReadSecret(zc common.ZAuthCommonComp) (string, error) {
	for {
		fmt.Print("Secret (required): ")
		sec, err := zc.UserInput()
		if err != nil {
			return "", err
		}

		sec = strings.TrimSpace(sec)
		if sec == "" {
			fmt.Fprint(flag.CommandLine.Output(), "secret cannot be empty\n")
		} else {
			_, err := base32.StdEncoding.DecodeString(sec)
			if err != nil {
				fmt.Fprint(flag.CommandLine.Output(), "invalid base32 secret. please try again.\n")
			} else {
				return sec, nil
			}
		}
	}
}

func (za *ZAuthArgsEntry) ReadIssuer(zc common.ZAuthCommonComp) (string, error) {
	for {
		fmt.Print("Issuer (eg: GitHub/Google..) (required): ")
		iss, err := zc.UserInput()
		if err != nil {
			return "", err
		}

		iss = strings.TrimSpace(iss)
		if iss == "" {
			fmt.Fprint(flag.CommandLine.Output(), "issuer cannot be empty\n")
		} else {
			return iss, nil
		}
	}
}

func (za *ZAuthArgsEntry) ReadIdentifier(zc common.ZAuthCommonComp) (string, error) {
	for {
		fmt.Print("Identifier (eg: Username/email) (required): ")
		acc, err := zc.UserInput()
		if err != nil {
			return "", err
		}
		acc = strings.TrimSpace(acc)
		if acc == "" {
			fmt.Fprint(flag.CommandLine.Output(), "identifier cannot be empty\n")
		} else {
			return acc, nil
		}
	}
}

func (za *ZAuthArgsEntry) ReadType(zc common.ZAuthCommonComp) (string, error) {
	for {
		fmt.Print("Type (TOTP/HOTP) (default: TOTP): ")
		typ, err := zc.UserInput()
		if err != nil {
			return "", err
		}

		typ = strings.TrimSpace(typ)
		if typ == "" {
			return zauth.DefaultType, nil
		}

		t := strings.ToLower(typ)
		if t == "totp" || t == "hotp" {
			return t, nil
		} else {
			fmt.Fprint(flag.CommandLine.Output(), "bad input\n")
		}
	}
}

func (za *ZAuthArgsEntry) ReadDigits(zc common.ZAuthCommonComp) (int, error) {
	for {
		fmt.Print("Digits (default: 6): ")
		sdig, err := zc.UserInput()
		if err != nil {
			return 0, err
		}

		sdig = strings.TrimSpace(sdig)
		if sdig == "" {
			return zauth.DefaultDigits, nil
		}

		dig, err := strconv.Atoi(sdig)
		if err != nil {
			fmt.Fprint(flag.CommandLine.Output(), "bad input\n")
		} else {
			return dig, nil
		}
	}
}

func (za *ZAuthArgsEntry) ReadAlgorithm(zc common.ZAuthCommonComp) (string, error) {
	for {
		fmt.Print("Algorithm (sha1/sha256/sha512) (default: sha1): ")
		algo, err := zc.UserInput()
		if err != nil {
			return "", err
		}

		algo = strings.TrimSpace(algo)
		if algo == "" {
			return zauth.DefaultAlgo, nil
		}

		algo = strings.ToLower(algo)
		if algo == "sha1" || algo == "sha256" || algo == "sha512" {
			return algo, nil
		} else {
			fmt.Fprint(flag.CommandLine.Output(), "bad input\n")
		}
	}
}

func (za *ZAuthArgsEntry) ReadPeriod(zc common.ZAuthCommonComp) (int64, error) {

	for {
		fmt.Print("Period (default: 30): ")
		sper, err := zc.UserInput()
		if err != nil {
			return 0, err
		}

		sper = strings.TrimSpace(sper)
		if sper == "" {
			return zauth.DefaultPeriod, nil
		}
		per, err := strconv.ParseInt(sper, 10, 64)
		if err != nil {
			fmt.Fprint(flag.CommandLine.Output(), "bad input\n")
		} else {
			return per, nil
		}
	}
}

func (za *ZAuthArgsEntry) ReadCounter(zc common.ZAuthCommonComp) (int64, error) {
	for {
		fmt.Print("Counter (default: 0): ")
		sctr, err := zc.UserInput()
		if err != nil {
			return 0, err
		}

		sctr = strings.TrimSpace(sctr)
		if sctr == "" {
			return zauth.DefaultCounter, nil
		}

		ctr, err := strconv.ParseInt(sctr, 10, 64)
		if err != nil {
			fmt.Fprint(flag.CommandLine.Output(), "bad input\n")
		} else {
			return ctr, nil
		}
	}
}

func printUsage(t string, f *flag.FlagSet) {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s %s [OPTIONS]\n\nOPTIONS:\n", os.Args[0], t)
	f.PrintDefaults()
}

func printZAuthOtpTable() error {
	tbl := table.New("ISSUER", "IDENTIFIER", "TYPE", "OTP", "REMAINING")

	tbl.WithPadding(5)

	zl, err := common.ReadZAuthJson()
	if err != nil {
		msg := fmt.Sprintf("An error occured while reading entries: %v", err)
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", msg)
		return fmt.Errorf(msg)
	}

	for _, z := range zl {
		otp, err := otp.GenerateOTP(&z)
		if err != nil {
			fmt.Fprintf(flag.CommandLine.Output(), "An error occured while generating OTP for %s: %v\n", z.Label, err)
		}

		var lbl string

		if strings.Contains(z.Label, ":") {
			lbl = strings.Split(z.Label, ":")[1]
		} else {
			lbl = z.Label
		}

		tbl.AddRow(z.Issuer, lbl, strings.ToUpper(z.Type), otp.Otp, otp.Remaining)
	}
	tbl.Print()

	return nil
}
