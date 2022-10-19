package procrun

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

func getPath() string {
	curDir, _ := os.Getwd()

	return curDir + "/"
}

// StartProcess start a background process
func StartProcess(procName string, protocol string, port string) error {

	cmd := exec.Command("nohup", procName, protocol, port)
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := writePidFile(strconv.Itoa(cmd.Process.Pid)); err != nil {
		return err
	}

	return nil

}

func KillProcess() error {

	pids, err := ioutil.ReadDir(getPath())
	if err != nil {
		return err
	}

	for _, xpid := range pids {
		pidFileName := xpid.Name()

		pkill, _ := strconv.Atoi(filenameWithoutExtension(pidFileName))
		// proc, err := os.FindProcess(pkill)
		// if err != nil {
		// 	removePidFile(pidFileName)
		// 	return err
		// }

		proc := os.Process{}
		proc.Pid = pkill

		if err := proc.Kill(); err != nil {
			removePidFile(pidFileName)
			continue
		}
		removePidFile(pidFileName)

	}
	return nil
}

func removePidFile(pidFile string) error {

	if err := os.Remove(pidFile); err != nil {
		return err
	}
	return nil
}

func filenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func writePidFile(pidFile string) error {
	// // Read in the pid file as a slice of bytes.
	// if piddata, err := ioutil.ReadFile(pidFile); err == nil {
	// 	// Convert the file contents to an integer.
	// 	if pid, err := strconv.Atoi(string(piddata)); err == nil {
	// 		// Look for the pid in the process list.
	// 		if process, err := os.FindProcess(pid); err == nil {
	// 			// Send the process a signal zero kill.
	// 			if err := process.Signal(syscall.Signal(0)); err == nil {
	// 				// We only get an error if the pid isn't running, or it's not ours.
	// 				return fmt.Errorf("pid already running: %d", pid)
	// 			}
	// 		}
	// 	}
	// }
	// If we get here, then the pidfile didn't exist,
	// or the pid in it doesn't belong to the user running this app.
	return ioutil.WriteFile(getPath()+pidFile+".pid", []byte(fmt.Sprintf(pidFile)), 0664)
}
