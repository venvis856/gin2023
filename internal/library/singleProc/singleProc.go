package singleProc

// 脚本进程单例模式相关方法封装

import (
	"fmt"
	"github.com/gogf/gf/os/grpool"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

// SingleProcess 判断进程的pid是否存在，存在则不再执行
// 如进程退出则删除对应的pid文件
func SingleProcess(sigfile string) bool {
	// 打开一个文件，如果文件不存在 则创建
	f, err := os.OpenFile(sigfile, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return true
	}

	// 下面是文件加锁
	if syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB) != nil {
		return true
	}

	// 下面是写入当前进程的pid到文件
	if ioutil.WriteFile(sigfile, []byte(fmt.Sprintln(os.Getpid())), os.ModePerm) != nil {
		return true
	}

	// 监听退出信号，如退出则删了pid文件
	NotifySignal(func(data ...interface{}) {
		os.Remove(sigfile)
	})
	return false
}

// GetFilePid 获取pid文件里的pid，如果存在对应的进程则返回pid，不存在则删除该文件
func GetFilePid(sigfile string) int {
	if pidByte, err := ioutil.ReadFile(sigfile); err == nil {
		pid, _ := strconv.Atoi(strings.TrimSpace(string(pidByte)))
		runos := runtime.GOOS
		if runos == "darwin" {
			// 判断pid是否存在
			_, err := os.FindProcess(pid)
			if err == nil {
				return pid
			}
		} else if CheckProcessExists(pid) {
			// linux系统
			return pid
		} else {
			os.Remove(sigfile)
		}
	}
	return 0
}

// CheckProcessExists 根据pid判断是否存在该进程（适用于linux）
func CheckProcessExists(pid int) bool {
	if _, err := os.Stat(filepath.Join("/proc", strconv.Itoa(pid))); err == nil {
		return true
	}
	return false
}

// KillPid 结束进程
func KillPid(pid int, signum syscall.Signal) error {
	if pid == 0 {
		return nil
	}
	return syscall.Kill(pid, signum)
}

// NotifySignal 监听信号
func NotifySignal(callback Callback) {
	// 创建chan
	c := make(chan os.Signal)

	// 监听所有信号
	signal.Notify(c)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("out signal: ", s)
				// 进程退出处理...
				callback(s)
				os.Exit(0)
			case syscall.SIGUSR1:
				fmt.Println("usr1 signal: ", s)
			case syscall.SIGUSR2:
				fmt.Println("usr1 signal: ", s)
			default:
				fmt.Println("other signal: ", s)
			}
		}
	}()
}

// Daemon 创建多个协程回调业务方法处理业务逻辑
func Daemon(maxProc int, callback func(d ...interface{}) bool) {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// 监听退出信号，退出协程
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	// 定义最大协程数，如果无数据处理时，自动定义为1
	loopNum := maxProc

	isLoop := make(chan bool, 1)
	go func() {

		// 协程都处理完业务逻辑后，才启动新的协程
		wg := sync.WaitGroup{}

		// 第一次发送数据表示要处理业务
		isLoop <- true

		for {
			select {
			case sig := <-sigs:
				fmt.Println(sig, "send exit command ...")
				// 发现进程退出信号，告诉channel结束协程
				done <- true
			case <-isLoop:

				if loopNum == 1 {
					// 如果只有一个协程时，说明数据少，sleep一下，避免cpu占用
					time.Sleep(time.Millisecond * 10)
				}

				// 判断chan长度小于指定长度时，则创建协程，使得协程数维持在指定数量内执行业务逻辑
				for i := 0; i < loopNum; i++ {
					wg.Add(1)
					num := i
					grpool.Add(func() {
						// 业务逻辑处理方法
						flag := callback(num)
						if !flag {
							// 无数据处理或异常时，协程数变为1，降低cpu占用
							loopNum = 1
						} else {
							loopNum = maxProc
						}
						wg.Done()
					})
				}
				wg.Wait()
				// 当前所有的协程结束后，进入下一次循环处理业务
				isLoop <- true
			}
		}
	}()
	fmt.Println("awaitng signal ...")
	<-done
	fmt.Println("exiting ...")
}

// Callback 定义回调
type Callback func(data ...interface{})
