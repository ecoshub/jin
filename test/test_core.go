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

func WriteFile(filedir string, buffer []byte) {
	err := ioutil.WriteFile(filedir, buffer, 0666)
	if err != nil {
		fmt.Printf("File Write Error:%v\n", err)
	}
}

func ReadFile(dir string) []byte {
	buff, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Printf("Read File Error:%v\n", err)
		return nil
	}
	return buff
}

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

func GetCurrentDir() string {
	wd, _ := os.Getwd()
	return wd
}

func Sep() string {
	return string(os.PathSeparator)
}
