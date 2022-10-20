package procrun

import (
	"testing"
)

func TestStartProcess(t *testing.T) {

	err := StartProcess("ngrok", "http", "9091")

	if err != nil {
		t.Errorf("Failed! %s - expected nil, receieved %s",
			t.Name(),
			err.Error(),
		)
	} else {
		t.Logf("Success !")
	}

}

func TestKillProcess(t *testing.T) {

	err := KillProcess()

	if err != nil {
		t.Errorf("Failed! %s - expected nil, receieved %s",
			t.Name(),
			err.Error())
	} else {
		t.Logf("Success !")
	}
}
