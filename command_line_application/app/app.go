package app

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

// ToGenerate is a function that returns a pointer to a cli.App instance
func ToGenerate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Get IPs And Name Servers"
	
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Get IP Address",
			Flags: flags,
			Action: getIps,
		},
		{
			Name: "ns",
			Usage: "Get Name Servers",
			Flags: flags,
			Action: getNameServers,
		},
	}

	return app
	
}


func getIps(c *cli.Context) {
	host := c.String("host")
    
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to lookup IP")
	}

	for _, ip := range ips {
		//log.Info().Msgf("IP Address: %s", ip.String())
		fmt.Println(ip.String())
	}
}

func getNameServers(c *cli.Context) {
	host := c.String("host")
	
	nameServers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to lookup Name Servers")
	}

	for _, ns := range nameServers {
		//log.Info().Msgf("Name Server: %s", ns.Host)
		fmt.Println(ns.Host)
	}
}