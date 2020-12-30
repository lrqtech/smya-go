package exec

import (
	"fmt"
	"os/exec"
	"runtime"
	"smya/wol"
)

// 执行命令
func Shell(tp int, shell string, name string) {
	if tp == 5 {
		if name == "wol" || name == "WOL" {
			wol.WolWake(shell, "", "", "")
		} else {
			switch runtime.GOOS {
			case "linux":
				switch runtime.GOARCH {
				case "386", "amd64":
					fmt.Printf("Run: %s", name)
					cmd := exec.Command("/bin/bash", "-c", shell)
					err := cmd.Run()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("Command has been executed")
				case "mips", "mipsle", "mips64", "mips64le":
					fmt.Printf("Run: %s", name)
					cmd := exec.Command("/bin/ash", "-c", shell)
					err := cmd.Run()
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println("Command has been executed")
				}
			case "windows":
				fmt.Printf("Run: %s", name)
				cmd := exec.Command("cmd.exe", "/c", "start "+shell)
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Command has been executed")
			case "darwin":
				fmt.Printf("Run: %s", name)
				cmd := exec.Command("/usr/bin/", "-c", shell)
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Command has been executed")
			}
		}
	} else {
		fmt.Println("Command dose not support")
		fmt.Println("Please use GUI program")
		fmt.Println("More information: https://smya.cn/download")
	}
}
