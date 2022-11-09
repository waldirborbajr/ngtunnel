package procrun

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/waldirborbajr/ngtunnel/internal/util"
)

// StartProcess start a background process
func StartProcess(procName string, protocol string, port string) error {

	cmd := exec.Command("nohup", procName, protocol, port)

	if err := cmd.Start(); err != nil {
		return err
	}

	time.Sleep(3 * time.Second)

	if err := writePidFile(strconv.Itoa(cmd.Process.Pid)); err != nil {
		return err
	}

	return nil
}

func KillProcess() error {

	pids, err := ioutil.ReadDir(util.GetPath())
	if err != nil {
		return err
	}

	for _, xpid := range pids {
		pidFileName := xpid.Name()

		extension := strings.ToLower(filepath.Ext(pidFileName))

		if extension == ".pid" {
			pkill, _ := strconv.Atoi(filenameWithoutExtension(pidFileName))

			proc := os.Process{}
			proc.Pid = pkill

			if err := proc.Kill(); err != nil {
				removePidFile(pidFileName)
				continue
			}
			removePidFile(pidFileName)
		}
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
	return ioutil.WriteFile(util.GetPath()+pidFile+".pid", []byte(fmt.Sprintf(pidFile)), 0664)
}
