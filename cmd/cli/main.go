package main

import (
	"fmt"
	"log"
	"os"

	color "github.com/waldirborbajr/bp-colors"
	"github.com/waldirborbajr/ngtunnel/internal/util"
	"github.com/waldirborbajr/ngtunnel/internal/version"
)

func init() {
	util.CheckOS()
}

func main() {
	port := ""

	if len(os.Args) > 1 {
		param := os.Args
		port = param[1]
	} else {
		fmt.Printf(color.Red + "NGTunnel " + version.AppVersion() + " - inform port number.\n\n" + color.Reset)
		fmt.Printf("Usage:\n")
		fmt.Printf("   ngtunnel [port_number]\n\n")
		os.Exit(1)
	}

	// Kill previous ngrok execution
	if err := util.KillNgrok("ngrok"); err != nil {
		log.Println("Error killing ngrok process")
	}

	// Verify if has curl installed
	curlPath := util.HasCurl()

	util.HasNoHup()

	util.StartNGRok(port)

	// Execute curl te grab public_url from ngrok
	util.GetngrokURL(curlPath)
}
