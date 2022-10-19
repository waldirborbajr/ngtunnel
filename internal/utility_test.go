package utility

import (
	"log"
	"os"
	"testing"
)

func TestKillNgRok(t *testing.T) {

	// cmd := exec.Command("nohup", "ngrok", "http", "9090")
	// err := cmd.Start()
	//
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	//
	// fmt.Println(cmd.Process.Pid)
	//
	// time.Sleep(5000 * time.Millisecond)

	proc := os.Process{}
	proc.Pid = 4616

	if err := proc.Kill(); err != nil {
		log.Fatal(err.Error())
	}

	// if err := cmd.Process.Kill(); err != nil {
	// 	log.Fatal(err.Error())
	// }
	// time.Sleep(5000 * time.Millisecond)

	// StartNGRok("1010")

	// err := KillNgrok("ngrok")
	// fmt.Println(err.Error())
	// assert.Equal(t, err.Error() != "exit status 1", err.Error())
}

func TestHasNoHup(t *testing.T) {

}

func TestCheckOS(t *testing.T) {

}

func TestHasCurl(t *testing.T) {

}

func TestGetngrokURL(t *testing.T) {

}

func TestStartNGRok(t *testing.T) {}
