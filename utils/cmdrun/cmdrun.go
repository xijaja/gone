package cmdrun

import (
	"os/exec"

	"github.com/gookit/color"
)

// RunCmd 执行命令行代码
func RunCmd(args string) {
	_, err := exec.Command("bash", "-c", args).Output()
	color.Info.Println("执行命令:", args)
	if err != nil {
		color.Error.Println("执行命令失败:", err)
	} else {
		color.Success.Println("执行命令成功!")
	}
}
