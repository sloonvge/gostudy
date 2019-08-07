package goos

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func TestOpenFile(t *testing.T) {
	t.Run("exist", func(t *testing.T) {
		file, err := os.Open("ex.txt")
		if err != nil {
			log.Fatal(err)
		}
		_ = file
	})
	t.Run("not exist", func(t *testing.T) {
		file, err := os.Open("notex.txt")
		if err != nil {
			log.Fatal(err)
		}
		_ = file
	})
}


func TestNewFile(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		file := os.NewFile(uintptr(syscall.Stderr), "stderr.log")
		if file != nil {

			_, err := file.WriteString("Write content to stdout.\n")
			if err != nil {
				log.Fatal(err)
			}
		}
		err :=  file.Close()
		if err != nil {
			log.Fatal(err)
		}
	})
}

func TestStartProcess(t *testing.T) {
	name := "dir"
	args := []string{"."}
	attr := &os.ProcAttr{}
	proc, err := os.StartProcess(name, args, attr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = proc.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

// https://www.cnblogs.com/benlightning/articles/4441192.html
// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter10/10.1.html
func TestExec(t *testing.T) {
	t.Run("command", func(t *testing.T) {
		dateCmd := exec.Command("go", "version")
		dateOut, err := dateCmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(dateOut))
	})
	t.Run("pipe", func(t *testing.T) {
		cmd := exec.Command("go", "version")
		in, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		out, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}
		cmd.Start()
		in.Write([]byte("hello go\ngoodbay go"))
		in.Close()
		bytes, err := ioutil.ReadAll(out)
		cmd.Wait()
		log.Println()
		log.Println(string(bytes))
	})
	t.Run("look", func(t *testing.T) {
		binary, err := exec.LookPath("go")
		if err != nil {
			log.Fatal(err)
		}
		args := []string{"version"}
		env := os.Environ()
		err = syscall.Exec(binary, args, env)
		if err != nil {
			log.Fatal(err)
		}
	})
}
