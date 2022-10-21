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
		t.Logf("%s - Success ! ", t.Name())
	}

}

func TestKillProcess(t *testing.T) {

	err := KillProcess()

	if err != nil {
		t.Errorf("Failed! %s - expected nil, receieved %s",
			t.Name(),
			err.Error())
	} else {
		t.Logf("%s - Success !", t.Name())
	}
}

func TestGetPath(t *testing.T) {
	result := getPath()

	if result != "" {
		t.Logf("%s - Success !", t.Name())
	} else {
		t.Errorf("Failed! %s - expected nil, receieved %s",
			t.Name(),
			result)
	}

}

func TestFilenameWithoutExtension(t *testing.T) {

	input := "abc.pid"
	result := "abc"

	noExtension := filenameWithoutExtension(input)

	if noExtension != result {
		t.Errorf("Failed! %s - expected %s, receieved %s",
			t.Name(),
			input,
			noExtension)
	} else {
		t.Logf("%s - Success !", t.Name())
	}

}
