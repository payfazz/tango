package util

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// RunScript create sh file, run it, and delete sh file
func RunScript(fileName string, content string, args ...string) error {
	// Create sh file
	initFile, err := os.Create(fileName)
	if nil != err {
		return errors.New(fmt.Sprintf("failed to create %s: %s", fileName, err))
	}

	// Write script to sh file
	_, err = initFile.Write([]byte(content))
	if nil != err {
		return errors.New(fmt.Sprintf("failed to write %s: %s", fileName, err))
	}

	// Close file after writing
	err = initFile.Close()
	if nil != err {
		return errors.New(fmt.Sprintf("failed to close file %s: %s", fileName, err))
	}

	// Set chmod to 755
	err = os.Chmod(fileName, 0755)
	if nil != err {
		return errors.New(fmt.Sprintf("failed to chmod %s: %s", fileName, err))
	}

	cmd := exec.Command(fmt.Sprintf("./%s", fileName), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	// Remove sh file after running
	rErr := os.Remove(fileName)
	if nil != rErr {
		return errors.New(fmt.Sprintf("failed to remove %s: %s", fileName, err))
	}

	if nil != err {
		return errors.New(fmt.Sprintf("failed to run %s: %s", fileName, err))
	}

	return nil
}
