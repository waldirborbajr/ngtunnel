package main

import (
	"fmt"
	"log"
	"os"

	utility "localhost/ngtunnel/internal"
	cfg "localhost/ngtunnel/pkg/config"

	color "github.com/waldirborbajr/bp-colors"
)

func init() {
	utility.CheckOS()
}

func main() {
	port := ""

	if len(os.Args) > 1 {
		param := os.Args
		port = param[1]
	} else {
		fmt.Printf(color.Green + "NGTunnel " + cfg.Verzion + " - inform port number.\n\n" + color.Reset)
		fmt.Printf("Usage:\n")
		fmt.Printf("   ngtunnel [port_number]\n\n")
		os.Exit(0)
	}

	// Kill previous ngrok execution
	if err := utility.KillNgrok("ngrok"); err != nil {
		log.Println("Error kill ngrok process")
	}

	// Verify if has curl installed
	curlPath := utility.HasCurl()

	utility.HasNoHup()

	utility.StartNGRok(port)

	// Execute curl te grab public_url from ngrok
	utility.GetngrokURL(curlPath)
}
