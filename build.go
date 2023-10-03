package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

//go:generate stringer -type ErrCode -linecomment ./global/errcode

// 如果打包go ，请先用这个包
//go install github.com/gogf/gf-cli/v2/gf@master

func main() {
	var platform string
	var module string
	flag.StringVar(&platform, "p", "", "-p 后面的值")
	flag.StringVar(&module, "m", "", "-  后面的值")
	flag.Parse()
	if platform == "" {
		platform = "linux"
	}
	fmt.Printf("编译平台为:%v\n", platform)
	fmt.Println("当前平台为:", runtime.GOOS)
	// darwin window linux
	switch module {
	case "system_api":
		SystemApi(platform)
	}
}

/*
*
编译后端模块
*/
func SystemApi(platform string) {
	buildPlatForm := "linux"
	if platform == "window" {
		buildPlatForm = "windows"
	}
	if platform == "mac" {
		buildPlatForm = "darwin"
	}

	switch runtime.GOOS {
	case "windows":
		Command("rmdir /Q /S release")
		Command("mkdir release")
		Command(" gf build main.go -s linux  -a amd64  -p build -n system_api")
		//Command(" gf build cmd\\cli\\main.go -s linux  -a amd64  -p build -n cli")
		Command("xcopy build\\linux_amd64\\ release\\  /y /e /i /q")

		//Command("xcopy config release\\config  /y /e /i /q")

		// window 下tar有时会报错permission denied，所以放在其他目录
		Command("cd release && tar -zcvf ..\\build\\release.tar.gz .\\*")
		Command("move build\\release.tar.gz .\\system_api_release.tar.gz")
		Command("rmdir /Q /S build")

	case "linux":
		fallthrough
	case "darwin": // mac
		Command("rm -rf release")
		Command("mkdir release")
		setEnv := fmt.Sprintf("CGO_ENABLED=0 GOOS=%v GOARCH=amd64", buildPlatForm)
		Command(" " + setEnv + "  go build -o release/system_api ./main.go")
		//Command(" " + setEnv + "  go build -o release/cli ./cmd/cli/main.go")
		//Command("cp -r config release/config")
		Command("cd release && tar -zcvf release.tar.gz ./*")
		Command("mv release/release.tar.gz ./system_api_release.tar.gz")
	}
}

func Command(sh string) {
	var name string = ""
	var args string = "-c"
	switch runtime.GOOS {
	case "windows":
		name = "cmd"
		args = "/c"
	case "darwin":
		name = "/bin/bash"
	case "linux":
		name = "/bin/bash"
	}
	fmt.Println("执行语句:", sh)
	cmd := exec.Command(name, args, sh)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	e := cmd.Run()
	if e != nil {
		fmt.Println(e)
	}
}
