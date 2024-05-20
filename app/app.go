package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {

	app := cli.NewApp()
	app.Name = "Simple CLI"
	app.Usage = "Search IPs and Server Names On The Internet."

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search Public Ips",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchIP,
		},
		{
			Name:  "server",
			Usage: "Search ServerName",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchServerName,
		},
	}
	return app
}

func searchIP(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServerName(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, server := range servers {
		fmt.Println(server)
	}
}
