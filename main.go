package main

import (
	"log"
	"os"

	gconfig "github.com/MrNinso/og-proxy/commands/genarate-config"

	"github.com/urfave/cli/v2"
)

func main() {
	proxyCli := cli.NewApp()

	proxyCli.Commands = []*cli.Command{
		gconfig.Cmd(),
	}

	if err := proxyCli.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
