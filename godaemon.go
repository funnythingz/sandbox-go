package main

import (
	"github.com/VividCortex/godaemon"
	"log"
	"os"
	"syscall"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

var (
	err error
	f   *os.File
)

func main() {

	f, err = os.OpenFile("./gin-bin", os.O_RDONLY, 0755)
	failOnError(err)

	err = syscall.Flock(int(f.Fd()), syscall.LOCK_EX)
	failOnError(err)

	godaemon.MakeDaemon(&godaemon.DaemonAttr{
		ProgramName: "plog",
		Files:       []**os.File{&f},
	})

	f.Close()
}
