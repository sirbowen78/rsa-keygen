package cmdline

import (
	"log"
	"os"
	"sort"

	"github.com/sirbowen78/rsa-keygen/rsa"
	"github.com/urfave/cli"
)

// Command initializes the command to generate rsa key pair.
func Command() {
	app := &cli.App{
		Name:  "RSA key generation tool.",
		Flags: flg,
		Action: func(c *cli.Context) error {
			err := rsa.ExportRSAKeysToFile(privKeyFileName, pubkeyFileName, keySize)
			if err != nil {
				return err
			}
			return nil
		},
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	appErr := app.Run(os.Args)
	if appErr != nil {
		log.Fatal(appErr)
	}
}
