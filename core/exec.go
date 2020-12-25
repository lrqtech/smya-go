package core

import (
	"bufio"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"io/ioutil"
	"os/exec"
	"runtime"
)

// 执行命令
func ExecShell(client mqtt.Client, tp int, shell string, name string) {
	if tp == 5 {
		if name == "wol" || name == "WOL" {
			WolWake(shell, "", "", "")
		} else {
			switch runtime.GOOS {
			case "linux":
				switch runtime.GOARCH {
				case "386", "amd64":
					cmd := exec.Command("/bin/bash", "-c", shell)
					stdout, err := cmd.StdoutPipe()
					if err != nil {
						fmt.Println(err)
					}
					_ = cmd.Start()
					t, err := ioutil.ReadAll(bufio.NewReader(stdout))
					if err != nil {
						fmt.Println(err)
					}
					s := "Command has been executed"
					text := fmt.Sprintf("Run: %s \n %s \n %s \n", name, t, s)
					Publish(client, MsgFmtJson(text))
				case "mips", "mipsle":
					cmd := exec.Command("/bin/ash", "-c", shell)
					stdout, err := cmd.StdoutPipe()
					if err != nil {
						fmt.Println(err)
					}
					_ = cmd.Start()
					t, err := ioutil.ReadAll(bufio.NewReader(stdout))
					if err != nil {
						fmt.Println(err)
					}
					s := "Command has been executed"
					text := fmt.Sprintf("Run: %s \n %s \n %s \n", name, t, s)
					Publish(client, MsgFmtJson(text))
				}
			case "windows":
				cmd := exec.Command("cmd.exe", "/c", "start "+shell)
				stdout, err := cmd.StdoutPipe()
				if err != nil {
					fmt.Println(err)
				}
				_ = cmd.Start()
				t, err := ioutil.ReadAll(bufio.NewReader(stdout))
				if err != nil {
					fmt.Println(err)
				}
				s := "Command has been executed"
				text := fmt.Sprintf("Run: %s \n %s \n %s \n", name, t, s)
				Publish(client, MsgFmtJson(text))
			case "darwin":
				cmd := exec.Command("/usr/bin/", "-c", shell)
				stdout, err := cmd.StdoutPipe()
				if err != nil {
					fmt.Println(err)
				}
				_ = cmd.Start()
				t, err := ioutil.ReadAll(bufio.NewReader(stdout))
				if err != nil {
					fmt.Println(err)
				}
				s := "Command has been executed"
				text := fmt.Sprintf("Run: %s \n %s \n %s \n", name, t, s)
				Publish(client, MsgFmtJson(text))
			}
		}
	} else {
		fmt.Println("Command dose not support")
		fmt.Println("Please use GUI program")
		fmt.Println("More information: https://smya.cn/download")
	}
}
