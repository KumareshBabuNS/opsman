package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	const VERSION string = "0.0.2"
	app := &cli.App{
		Name: "opsman-cli",
		HelpName: "opsman-cli",
		Version: VERSION,
	}
	app.Commands = []*cli.Command{
		{
			Name:   "token",
			Usage:  "retieve token",
			Action: RetrieveToken,
			Flags:  []cli.Flag{OpsManagerURLFlag, Username, Password, SkipSSL},
		},
		{
			Name:   "download",
			Usage:  "download product",
			Action: DownloadProduct,
			Flags:  []cli.Flag{ProductURL, PivnetToken, SaveProductTo},
		},
		{
			Name:   "upload",
			Usage:  "upload product",
			Action: UploadProduct,
			Flags:  []cli.Flag{OpsManagerURLFlag, Username, Password, UploadProductFrom, SkipSSL},
		},
	}
	app.Run(os.Args)
}
