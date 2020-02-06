package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// WriteFile is NOT FOR PUBLIC USAGE.
// This function is created for test-case creation automation
// Please do not change any thing. And do not use them.
func WriteFile(filedir string, buffer []byte) {
	err := ioutil.WriteFile(filedir, buffer, 0666)
	if err != nil {
		fmt.Printf("File Write Error:%v\n", err)
	}
}

// ReadFile is NOT FOR PUBLIC USAGE.
// This function is created for test-case creation automation
// Please do not change any thing. And do not use them.
func ReadFile(dir string) []byte {
	buff, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Printf("Read File Error:%v\n", err)
		return nil
	}
	return buff
}

// ExecuteNode is NOT FOR PUBLIC USAGE.
// This function is created for test-case creation automation
// Please do not change any thing. And do not use them.
func ExecuteNode(par string) (string, error) {
	cmd := exec.Command("node", "test/test-case-creator.js", par)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// Dir is NOT FOR PUBLIC USAGE.
// This function is created for test-case creation automation
// Please do not change any thing. And do not use them.
func Dir(dir string) []string {
	files := make([]string, 0, 100)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		files = append(files, strings.TrimPrefix(path, dir))
		return nil
	})
	if err != nil {
		fmt.Println("Error walking true path", err)
	}
	return files[1:]
}

// GetCurrentDir is NOT FOR PUBLIC USAGE.
// This function is created for test-case creation automation
// Please do not change any thing. And do not use them.
func GetCurrentDir() string {
	wd, _ := os.Getwd()
	return wd
}

// Sep is NOT FOR PUBLIC USAGE.
// This function is created for test-case creation automation
// Please do not change any thing. And do not use them.
func Sep() string {
	return string(os.PathSeparator)
}
