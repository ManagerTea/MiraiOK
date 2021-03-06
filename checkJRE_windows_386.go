package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func checkJava() {
	//defer global.Done()
	//检测本地java
	if checkJavaBin() {
		return
	}
	unpackRAR(downFile("MiraiOK/jre-" + RTStr + ".rar"))
	if checkJavaBin() {
		return
	}
	ERROR("无法获取JRE，即将退出...")
	os.Exit(0)
}

func checkJavaBin() bool {
	var stdo bytes.Buffer
	cmd := exec.Command(javaPath, "-version")
	cmd.Stdout = &stdo
	cmd.Stderr = &stdo
	err := cmd.Run()
	if err != nil {
		return false
	}
	for str, err := stdo.ReadString('\n'); err == nil; {
		INFO("JRE:", strings.TrimRight(str, "\r\n"))
		str, err = stdo.ReadString('\n')
	}
	return true
}
