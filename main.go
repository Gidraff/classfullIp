package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ip",
				Value: "127.0.0.1",
				Usage: "IP address to be parsed",
			},
		},

		Action: func(ctx *cli.Context) error {
			ip := "127.0.0."
			if ctx.NArg() > 0 {
				ip = ctx.Args().Get(0)
			}
			result := net.ParseIP(ip)
			if result == nil {
				log.Print(errors.New("invalid ip address"))
			}
			octetList := strings.Split(result.String(), ".")
			ipInt, err := strconv.ParseInt(octetList[0], 10, 32)
			if err != nil {
				log.Fatal(err.Error())
				os.Exit(1)
			}

			if 0 <= ipInt && ipInt <= 127 {
				fmt.Println(result.String() + "/8")
			}
			if ipInt >= 128 && ipInt <= 191 {
				fmt.Println(result.String() + "/16")
			}
			if ipInt >= 192 && ipInt <= 223 {
				fmt.Println(result.String() + "/24")
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
