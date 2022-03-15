package exec_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

// ---------------------------------------------------------------------
//
// go os/exec 简明教程
// https://colobu.com/2020/12/27/go-with-os-exec/
//
// ---------------------------------------------------------------------

// Cmd命令包含输入和输出字段，你可以设置这些字段，实现定制输入和输出：

// Stdin io.Reader
// Stdout io.Writer
// Stderr io.Writer
// 如果Stdin为空，那么进程会从null device(os.DevNull)中读取。
// 如果Stdin是*os.File对象，那么会从这个文件中读取。
// 如果Stdin是os.Stdin,那么会从标准输入比如命令行中读取数据。

// Stdout和Stderr代表外部程序进程的标准输出和错误输出，如果为空，那么输出到null device中。
// 如果Stdout和Stderr是*os.File对象，那么会往文件中输出数据。
// 如果Stdout和Stderr分别设置为os.Stdout、os.Stderr的话，会输出到命令行中。

// 我们改造上一个例子，显示命令输出结果：
func TestStdoutStderr(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

// 默认情况下进程的工作路径是调用这个进程的文件夹，但是你也可以手工指定，比如我们将工作路径指定为根路径：
func TestCmdDir(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "/"
	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

// Cmd.Path是要执行的程序的路径，如果是相对路径，那么它基于Cmd.Dir计算相对路径。如果程序已经在系统$PATH路径下，那么可以直接写程序名。
func TestCmdPath(t *testing.T) {
	cmd := exec.Command("/usr/local/go/bin/go", "env")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("path: %s", cmd.Path)
	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

// Cmd还有一个字段叫做Env,用来设置进程的环境变量，格式是key=value。
// 如果Env为空，那么新进程将使用调用者进程的环境变量。
// 比如下面的例子，我们设置了myvar变量，你可以注释掉cmd.Enc = ...那一行比较一下结果。
func TestCmdEnv(t *testing.T) {
	cmd := exec.Command("bash", "-c", "echo $myvar")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Env = []string{"myvar=wangxb"}
	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

// os/exec是封装的一个便利库，底层它是使用os.StartProcess实现的，所以你可以得到底层的Process对象和ProcessState对象,分别代表进程和进程的状态。
func TestCmdProcess(t *testing.T) {
	cmd := exec.Command("bash", "-c", "sleep 1;echo $myvar")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("failed to call cmd.Start(): %v", err)
	}

	log.Printf("pid: %d", cmd.Process.Pid)
	cmd.Process.Wait()
	log.Printf("exitcode: %d", cmd.ProcessState.ExitCode())
}

// 有时候，你需要在执行一个外部命令的时候需要先检查它是否存在，你可以使用LookPath方法。
// 如果传入参数包含路径分隔符，那么它会基于Cmd.Dir的相对路径或者绝对路径查找这个程序。如果不包含路径分隔符，那么会从PATH环境变量中查找文件。
func TestExecLookPath(t *testing.T) {
	path, err := exec.LookPath("ls")
	if err != nil {
		log.Printf("'ls' not found")
	} else {
		log.Printf("'ls' is in '%s'\n", path)
	}

}

// Cmd提供了Output()方法，如果命令正确执行，可以得到命令执行结果的bytes:
func TestCmdOutput(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	data, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to call Output(): %v", err)
	}
	log.Printf("output: %s", data)

}

// 如果命令出错，错误信息可以通过Stderr获得：
func TestCmdOutputStdErr(t *testing.T) {
	cmd := exec.Command("ls", "-lahxyz")
	cmd.Stderr = os.Stderr
	data, err := cmd.Output()
	if err != nil {
		t.Fatalf("failed to call Output(): %v", err)
	}
	t.Logf("output: %s", data)
}

// 如果你想不管出错与否都能一个方法获取输出结果的话，你可以调用CombinedOutput()方法，它会返回正常输出或者错误输出，并且第二个返回err可以指示是否执行出错：
func TestCmdCombinedOutput(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to call CombinedOutput(): %v", err)
	}
	log.Printf("output: %s", data)
}

// CombinedOutput方法的实现也很简单，其实就是共享同一个bytes.Buffer实现的。
// 了解了CombinedOutput的实现，我们就可以为Stdout和Stderr分别设置bytes.Buffer,来实现独立的读取。
func TestCmdStdoutStdErr(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}
	log.Printf("out:\n%s\nerr:\n%s", stdout.String(), stderr.String())

}

// 既然我们已经能够使用自己的io.Writer设置Stdout/Stderr,那么我们可以做更丰富的功能。
// 比如利用curl命令下载一个大的文件，我们可以实时的显示已下载的数据的大小。
func TestCmdProcessBar(t *testing.T) {
	cmd := exec.Command("curl", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")

	var stdoutProcessStatus bytes.Buffer
	cmd.Stdout = io.MultiWriter(ioutil.Discard, &stdoutProcessStatus)

	done := make(chan struct{})
	go func() {
		tick := time.NewTicker(time.Second)
		defer tick.Stop()

		for {
			select {
			case <-done:
				return
			case <-tick.C:
				log.Printf("downloaded:%d", stdoutProcessStatus.Len())
			}
		}
	}()

	err := cmd.Run()
	if err != nil {
		t.Fatalf("failed to call Run(): %v", err)
	}
	close(done)
}

// 前面几个例子都是演示处理Output的情况，接下来这个例子演示了如何设置Stdin。
// wc命令读取main.go文件，统计它一共有多少行。
func TestCmdStdin(t *testing.T) {
	stdin, err := os.Open("stdpkg_examples_test.go")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	cmd := exec.Command("wc", "-l")
	cmd.Stdin = stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("failed to call cmd.Run(): %v", err)
	}
}

// 你可以将一个命令的输出作为下一个命令的输入，以此类推，将多个命令串成一个管道。
// os/exec提供了StderrPipe、StdinPipe、StdoutPipe方法，获取管道对象。
// 比如下面的命令，将cat main.go的输出作为wc -l命令的输入:
func TestCmdPipe(t *testing.T) {
	cmdCat := exec.Command("cat", "colobu_examples_test.go")
	catout, err := cmdCat.StdoutPipe()
	if err != nil {
		log.Fatalf("failed to get StdoutPipe of cat: %v", err)
	}

	cmdWC := exec.Command("wc", "-l")
	cmdWC.Stdin = catout
	cmdWC.Stdout = os.Stdout

	err = cmdCat.Start()
	if err != nil {
		log.Fatalf("failed to call cmdCat.Run(): %v", err)
	}

	err = cmdWC.Start()
	if err != nil {
		log.Fatalf("failed to call cmdWC.Start(): %v", err)
	}

	cmdCat.Wait()
	cmdWC.Wait()
}

// 下面是一个更通用的创建Cmd管道的方法：
func TestCmdPipeFn(t *testing.T) {
	cmdCat := exec.Command("cat", "colobu_examples_test.go")
	cmdWC := exec.Command("wc", "-l")
	data, err := pipeCommands(cmdCat, cmdWC)
	if err != nil {
		log.Fatalf("failed to call pipeCommands(): %v", err)
	}
	log.Printf("output: %s", data)
}

func pipeCommands(commands ...*exec.Cmd) ([]byte, error) {
	for i, command := range commands[:len(commands)-1] {
		out, err := command.StdoutPipe()
		if err != nil {
			return nil, err
		}
		command.Start()
		commands[i+1].Stdin = out
	}
	final, err := commands[len(commands)-1].Output()
	if err != nil {
		return nil, err
	}
	return final, nil
}

// 如果你通过bash命令执行，你可以使用bash pipe的功能，写起来更简单。
func TestCmdBashPipe(t *testing.T) {
	cmd := exec.Command("bash", "-c", "cat colobu_examples_test.go| wc -l")
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("failed to call pipeCommands(): %v", err)
	}
	log.Printf("output: %s", data)
}

// 孤儿进程
// 当子进程还没有结束的时候，父进程先结束了，那么此时的子进程就叫做孤儿进程，这个时候子进程的ppid就被设置为了1(被系统1号进程接管)。
func TestCmdOrphanProcess(t *testing.T) {
	cmd := exec.Command("curl", "-o", "go1.15.6.linux-amd64.tar.gz", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
	err := cmd.Start()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}
}

// 程序退出时Kill子进程
// 如果我们想在程序退出的时候Kill掉它启动的子进程，那么一个比较笨的办法就是得到子进程的 Process对象，然后调用它的Kill方法将其杀掉。但比较遗憾的是它不能把孙进程杀掉。
// 对于Linux系统，你可以通过下面的设置将孙进程也杀掉：
func TestCmdKillChildProcess(t *testing.T) {
	cmd := exec.Command("curl", "-o", "go1.15.6.linux-amd64.tar.gz", "https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	err := cmd.Start()
	if err != nil {
		log.Fatalf("failed to call Run(): %v", err)
	}

}

// 将父进程打开的文件传给子进程
// 除了标准输入输出0,1,2三个文件外，你还可以将父进程的文件传给子进程，通过Cmd.ExtraFiles字段就可以。
// 比较常用的一个场景就是graceful restart,新的进程继承了老的进程监听的net.Listener,这样网络连接就不需要关闭重打开了。
func TestCmdParentFileToChild(t *testing.T) {
	file := netListener.File() // this returns a Dup()
	path := "/path/to/executable"
	args := []string{
		"-graceful"}
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = []*os.File{file}
	err := cmd.Start()
	if err != nil {
		log.Fatalf("gracefulRestart: Failed to launch, error: %v", err)
	}

}
