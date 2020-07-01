package util

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/payfazz/go-apt/pkg/fazzcommon/formatter"

	"github.com/pkg/errors"
)

// SnakeToPascalCase convert str from snake_case to PascalCase
func SnakeToPascalCase(str string) string {
	var link = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
	return link.ReplaceAllStringFunc(str, func(s string) string {
		return strings.ToUpper(strings.Replace(s, "_", "", -1))
	})
}

// SnakeToCamelCase convert str from snake_case to camelCase
func SnakeToCamelCase(str string) string {
	result := SnakeToPascalCase(str)
	return formatter.ToLowerFirst(result)
}

// PascalToCamelCase convert str from PascalCase to camelCase
func PascalToCamelCase(str string) string {
	return formatter.ToLowerFirst(str)
}

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

func PullRepoFolder() {
	homeDir, _ := os.UserHomeDir()
	tangoDir := homeDir + "/.tango"
	if _, err := os.Stat(homeDir + "/.tango"); os.IsNotExist(err) {
		fmt.Println("Get tango repository for project template..")
		cmd := exec.Command("git", "clone", "git@github.com:payfazz/tango.git", tangoDir)
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
