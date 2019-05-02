package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
)

var RunningCommand *exec.Cmd

type FakeIn struct {
}

func (f *FakeIn) Read(b []byte) (n int, err error) {
	n, e := os.Stdin.Read(b)
	return n, e
}

func main() {

	if len(os.Args) > 1 {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		checked := false

		go func() {
			for range c {
				if checked {
					if RunningCommand != nil {
						_ = RunningCommand.Process.Kill()
					}
					os.Exit(0)

				} else {
					checked = true
					//this currently breaks interactive commands
					fmt.Println("^C again to quit.")
				}

			}
		}()

		err := Run()
		checkError(err)
	} else {
		fmt.Println("Usage: wrapper command args")
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Run() error {

	RunningCommand := exec.Command("htop")
	RunningCommand.Stderr = os.Stderr
	RunningCommand.Stdout = os.Stdout
	RunningCommand.Stdin = &FakeIn{}
	return RunningCommand.Run()
}
