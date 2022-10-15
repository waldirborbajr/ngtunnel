package main

import (
	"fmt"
	"os"

	utility "localhost/ngtunnel/internal"

	color "github.com/waldirborbajr/bp-colors"
)

const Version = "0.6.1"

func init() {
	utility.CheckOS()
}

func main() {
	port := ""

	if len(os.Args) > 1 {
		param := os.Args
		port = param[1]
	} else {
		fmt.Printf(color.Green + "NGTunnel " + Version + " - inform port number.\n\n" + color.Reset)
		fmt.Printf("Usage:\n")
		fmt.Printf("   ngtunnel [port_number]\n\n")
		os.Exit(0)
	}

	// Kill previous ngrok execution
	utility.KillNgrok()

	// Verify if has curl installed
	curlPath := utility.HasCurl()

	utility.HasNoHup()

	utility.StartNGRok(port)

	// Execut curl te grab public_url from ngrok
	utility.GetngrokURL(curlPath)
}
