package procrun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartProcess(t *testing.T) {

	err := StartProcess("ngrok", "http", "9091")

	assert.Nil(t, err)

}

func TestKillProcess(t *testing.T) {

	err := KillProcess()

	assert.Nil(t, err)

}

func TestFilenameWithoutExtension(t *testing.T) {

	inputEq := "abc.pid"
	inputNEq := "abc.pid"
	result := "abc"

	noExtension := filenameWithoutExtension(inputEq)

	assert.Equal(t, noExtension, result)
	assert.NotEqual(t, noExtension, inputNEq)

}
