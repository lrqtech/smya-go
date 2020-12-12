package core

import (
	"fmt"
	"os/exec"
)

// 执行命令
func ExecShell(tp int, shell string) {
	switch tp {
	case 1:

	case 2:

	case 3:

	case 4:

	case 5:
		err := exec.Command(shell)
		if err != nil {
			fmt.Println(err)
		}
	}
}
