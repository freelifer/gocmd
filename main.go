package main

import (
	"bytes"
	"fmt"
	"github.com/freelifer/gocmd/publish"
	"os"
	"os/exec"
	"strings"
)

func main() {

	publish.Publish()
	// 配置 gobuild.toml
	// xx
	cmd := exec.Command("go", "build", "main.go")

	var stderr bytes.Buffer
	var out bytes.Buffer

	cmd.Stderr = &stderr // 标准错误
	cmd.Stdout = &out

	ChangeCmdEnvironment(cmd)
	err := cmd.Run()
	if err != nil {
		fmt.Printf(string(stderr.Bytes()))
		os.Exit(2)
	}
	fmt.Printf(string(out.Bytes()))
	fmt.Println("Done!")
	os.Exit(0)
}

func beforeBuild() {

}

func afterBuild() {

}

func ChangeCmdEnvironment(cmd *exec.Cmd) error {
	env := os.Environ()
	cmdEnv := []string{}

	for _, e := range env {
		i := strings.Index(e, "=")
		if i > 0 && (e[:i] == "GOOS") {
			// cmdEnv = append(cmdEnv, "windows")
		} else {
			cmdEnv = append(cmdEnv, e)
		}
	}
	// cmdEnv = append(cmdEnv, "GOOS=windows")
	cmd.Env = cmdEnv

	return nil
}
