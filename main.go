package main

import (
	"Golang/process"
	"fmt"
	"time"
)

func main() {

	// define input
	var a1 = []string{"erlang", "golang", "I", "Love"}

	// spawn new process (first arg is process run forEver or finish after done)
	var pid = process.Spawn(false, gprint)

	// show <pid> is a id of each process like erlang ;-)
	pid.PidPrint()

	// Bang is a function for send a message to process
	// erlang syntax ->  pid ! msg
	// golang syntax ->  Bang(pid,msg)
	process.Bang(pid, process.Dynamic(a1))

	time.Sleep(time.Minute * 2)
}

func gprint(dy process.Dynamic) {

	switch dy.(type) {
	case []string:
		fmt.Printf("%v \n", dy)
	}

}
