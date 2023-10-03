package helper

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// ExecGrepCommand 执行管道命令，如：ps -ef | grep modules | wc -l
func ExecGrepCommand(strCommand string) (*bytes.Buffer, error) {
	fmt.Println("command: ", strCommand)
	commands := strings.Split(strCommand, "|")
	var stdout *bytes.Buffer
	var err error
	for _, command := range commands {
		if command == "" {
			continue
		}
		stdout, err = ExecCommand(command, stdout)
		if err != nil {
			fmt.Println("execute command error: ", err.Error())
			return stdout, err
		}
	}
	return stdout, err
}

// ExecCommand 执行命令，如：ps -ef
func ExecCommand(strCommand string, prevStdout ...*bytes.Buffer) (*bytes.Buffer, error) {
	params := strings.Split(strCommand, " ")
	name := ""
	args := []string{}
	for _, v := range params {
		if v != "" {
			if name == "" {
				name = v
			} else {
				args = append(args, v)
			}
		}
	}
	cmd := exec.Command(name, args...)
	if len(prevStdout) > 0 && prevStdout[0] != nil {
		// 管道依赖输入
		cmd.Stdin = prevStdout[0]
	}

	stdout, cmdErr := SetCommandStd(cmd)
	err := cmd.Run()
	if err != nil {
		err = errors.New(err.Error() + cmdErr.String())
	}
	return stdout, err
}

// SetCommandStd 设置命令输出
func SetCommandStd(cmd *exec.Cmd) (stdout, stderr *bytes.Buffer) {
	stdout = &bytes.Buffer{}
	stderr = &bytes.Buffer{}
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return
}
