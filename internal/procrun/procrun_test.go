package procrun

import (
	"testing"
)

func TestStartProcess(t *testing.T) {

	err := StartProcess("ngrok", "http", "9090")

	if err != nil {
		t.Error("Expected nil, receieved ", err.Error())
	}

}

func TestKillProcess(t *testing.T) {
	err := KillProcess()

	if err != nil {
		t.Error("Expected nil, receieved ", err.Error())
	}
}
