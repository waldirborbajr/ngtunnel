package util

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5/plumbing/color"
)

func KillNgrok(processName string) error {
	argstr := `killall ` + processName
	if _, err := exec.Command("bash", "-c", argstr).Output(); err != nil {
		return err
	}
	return nil
}

func HasNoHup() {
	_, err := os.Stat("nohup")

	if !os.IsNotExist(err) {
		os.Remove("nohup")
	}
}

func CheckOS() {
	if strings.EqualFold(strings.ToLower(runtime.GOOS), strings.ToLower("windows")) {
		fmt.Println("This program is not applicable for Windows machine.")

		os.Exit(1)
	}
}

func HasCurl() string {
	var curlPath string
	var err error

	if curlPath, err = exec.LookPath("curl"); err != nil {
		fmt.Printf(color.Red + "curl not found. Please install it first and run it again.\n" + color.Reset)
		os.Exit(1)
	}
	return curlPath
}

func GetngrokURL(curlPath string) {
	out, err := exec.Command(curlPath, "-s", "http://127.0.0.1:4040/api/tunnels").Output()
	if err != nil {
		fmt.Println(color.Red + "Error executing curl. Please verify if ngrok it is up and running.\n" + color.Reset)
		os.Exit(1)
	}
	output := out[:]
	fmt.Println(processRegexp(string(output)))
}

func processRegexp(output string) string {
	str := ""

	re := regexp.MustCompile(`"public_url":"https://([^"]+)"`)
	reurl := regexp.MustCompile(`"https://([^"]+)"`)

	if len(re.FindStringIndex(output)) > 0 {
		str = re.FindString(output)

		if len(reurl.FindStringIndex(str)) > 0 {
			addr := reurl.FindString(str)
			addr = strings.ReplaceAll(addr, "\"", "")
			fmt.Println(addr)
			os.Exit(0)
		}
	}

	return ""
}

func StartNGRok(port string) {
	const two = 2

	devnull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0o755)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("nohup", "ngrok", "http", port)
	cmd.Stdout = devnull

	if err := cmd.Start(); err != nil {
		fmt.Println("Error ....")
	}

	time.Sleep(two * time.Second)
}
