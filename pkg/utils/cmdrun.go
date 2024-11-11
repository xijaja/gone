package utils

import (
	"github.com/gookit/color"
	"os/exec"
)

// RunCmd 在终端执行命令
func RunCmd(args string) {
	_, err := exec.Command("bash", "-c", args).Output()
	color.Info.Println("执行命令:", args)
	if err != nil {
		color.Error.Println("执行命令失败:", err)
	} else {
		color.Success.Println("执行命令成功!")
	}
}
