package test

import (
	"os/exec"
	"io/ioutil"
	"fmt"
	"bytes"
)

func WriteFile(filedir string, buffer []byte) {
	err := ioutil.WriteFile(filedir, buffer, 0666)	
	if err != nil {
		fmt.Printf("File Write Error:%v\n", err)
	}
}

func ReadFile(dir string) []byte{
	buff, err := ioutil.ReadFile(dir)
	if err != nil {
		fmt.Printf("Read File Error:%v\n", err)
		return nil
	}
	return buff
}

func ExecuteNode(par string) (string, error){
	cmd := exec.Command("node", "test/test-case-creator.js", par)	
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}else{
		return out.String(), nil
	}
}