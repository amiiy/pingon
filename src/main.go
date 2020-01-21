package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sparrc/go-ping"

	"github.com/urfave/cli"
)

func main() {

	app := &cli.App{
		Name:  "pingon",
		Usage: "pingon [host, host, ...]",
		Action: func(c *cli.Context) error {
			fmt.Println("hello from pingon")
			hostCount := len(c.Args().Slice())
			fmt.Printf("number of hosts: %d", hostCount)
			RunPing(c.Args().Slice())
			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}

func RunPing(hosts []string) {

	fmt.Println()
	for i := 0; i < len(hosts); i++ {
		fmt.Printf("running ping on: %v \n", hosts[i])
		RunPingOnHost(hosts[i])
	}
}

func RunPingOnHost(host string) {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		log.Fatal(err)
	}

	pinger.Count = 10
	pinger.Run()
	stats := pinger.Statistics()

	fmt.Printf("got: %v ---------- on host: %v", stats.AvgRtt, stats.Addr)
	fmt.Println()

}
