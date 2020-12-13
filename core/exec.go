package core

import (
	"fmt"
	"os/exec"
	"runtime"
)

// 执行命令
func ExecShell(tp int, shell string) {
	if tp == 5 {
		switch runtime.GOOS {
		case "linux":

		case "windows":
			cmd := exec.Command("cmd.exe", "/c", "start "+shell)
			// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}
		case "darwin":

		}
	}
	fmt.Println("Command dose not support")
	fmt.Println("Please use GUI program")
	fmt.Println("More information: https://smya.cn/download")
}
