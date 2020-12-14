package core

import (
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
)

// 执行命令
func ExecShell(tp int, shell string) {
	if tp == 5 {
		switch runtime.GOOS {
		case "linux":
			cmd := exec.Command("/bin/bash", "-c", shell)
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Command has been executed")
		case "windows":
			cmd := exec.Command("cmd.exe", "/c", "start "+shell)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Command has been executed")
		case "darwin":
			cmd := exec.Command("/usr/bin/", "-c", shell)
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Command has been executed")
		}
	} else {
		fmt.Println("Command dose not support")
		fmt.Println("Please use GUI program")
		fmt.Println("More information: https://smya.cn/download")
	}
}
