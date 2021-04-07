package config

import (
	"fmt"

	"github.com/MrNinso/og-proxy/types"
	"github.com/urfave/cli/v2"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "genarate-config",
		Aliases: []string{"gc"},
		Action:  CmdAction,
	}
}

func CmdAction(c *cli.Context) error {
	s, err := types.GetJson().MarshalIndent(buildExampleConfig(), "", "    ")

	fmt.Println(string(s))

	return err
}

func buildExampleConfig() types.Config {
	return types.Config{
		Servers: []interface{}{
			types.ServerHTTPS{
				Type: "HTTPS",
				Hosts: []types.HostHTTP{
					types.HostHTTP{
						Hostname: "http://host1.example.com",
						Path:     "",
						Port:     3000,
						Weight:   1,
					},

					types.HostHTTP{
						Hostname: "http://host2.example.com",
						Path:     "/api/v1",
						Port:     3002,
						Weight:   2,
					},
				},
				Certificate: "/path/to/ctr/file",
				Key:         "/path/to/key/file",
				ServerName:  "app.example.com",
				ServerPort:  443,
			},
			types.ServerTCP{
				Type: "TCP",
				Port: 600,
				Hosts: []types.HostTCPUDP{
					types.HostTCPUDP{
						Hostname: "server1.exemple.com",
						Port:  601,
						Weight: 1,
					},
				},
			},
			types.ServerUDP{
				Type: "UPD",
				Port: 603,
				Hosts: []types.HostTCPUDP{
					types.HostTCPUDP{
						Hostname: "server1.exemple.com",
						Port:  603,
						Weight: 1,
					},
				},
			},
		},
	}
}
