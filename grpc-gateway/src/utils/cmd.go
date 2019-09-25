package utils

import (
	"os"
	"os/exec"
	"strings"
)

func Cmdutil() error {
	logFile, _ := os.OpenFile("", os.O_CREATE | os.O_RDWR, 0777)
	args := strings.Fields("")
	cmd := exec.Command("", args...)
	//cmd.Stdin = logFile
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	err := cmd.Start()
	err = cmd.Wait()
	return err
}
