package utility

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func KillNgrok() {
	argstr := `killall ngrok`
	_, err := exec.Command("bash", "-c", argstr).Output()
	if err != nil {
		fmt.Println("Error killing NGRok")
	}
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
	curlPath, err := exec.LookPath("curl")
	if err != nil {
		fmt.Println("curl not found. Please install it first and run it again.")
		os.Exit(1)
	}
	return curlPath
}

func GetngrokURL(curlPath string) {
	out, err := exec.Command(curlPath, "-s", "http://127.0.0.1:4040/api/tunnels").Output()
	if err != nil {
		fmt.Println("Error executing curl. Please verify if ngrok it is up and running.")
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
