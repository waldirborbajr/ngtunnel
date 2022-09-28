package main

import (
	"fmt"
	"os"

	"localhost/ngtunnel/utility"

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
		fmt.Printf(color.Green + "NGTunnel v1.1.6 - inform port number.\n\n" + color.Reset)
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
