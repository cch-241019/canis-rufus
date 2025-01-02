package main

import (
	"canis-rufus/task/command/shell"
	"fmt"
	"time"
)

/*
* @author: Chen Chiheng
* @date: 2025/1/2 21:22:22
* @description:
**/

func main() {
	cmd := shell.New("ls -al | grep *.go")
	if err := cmd.RunWait(); err != nil {
		panic(err)
	}
	fmt.Printf(cmd.GetStdout().String())

	cmd = shell.New("sleep 30")
	if err := cmd.RunWaitTimeout(2 * time.Second); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf(cmd.GetStdout().String())
}
