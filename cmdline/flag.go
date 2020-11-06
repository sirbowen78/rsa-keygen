package cmdline

import "github.com/urfave/cli"

var (
	pubkeyFileName, privKeyFileName string
	keySize                         int
)

var flg = []cli.Flag{
	&cli.StringFlag{
		Name:        "pub",
		Value:       "pubkey.pem",
		Usage:       "Public Key Filename",
		Destination: &pubkeyFileName,
	},
	&cli.StringFlag{
		Name:        "pte",
		Value:       "privkey.pem",
		Usage:       "Private Key Filename",
		Destination: &privKeyFileName,
	},
	&cli.IntFlag{
		Name:        "b",
		Value:       2048,
		Usage:       "RSA key length in bits",
		Destination: &keySize,
	},
}
