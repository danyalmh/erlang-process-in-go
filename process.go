package process

import (
	"fmt"
	"reflect"
)

var maxRecipient int = 100
var maxElemtPid int8 = 100

// Dynamic is type
type Dynamic interface{}

// Pid is process id
type Pid struct {
	p1 int8
	p2 int8
	p3 int8
}

type process struct {
	pid       Pid
	recipient chan Dynamic
	function  func(Dynamic)
}

var currentPid Pid
var keeper []*process

func newPid() Pid {

	if currentPid.p3 < maxElemtPid {

		var pid = currentPid
		currentPid.p3++
		return pid

	} else if currentPid.p3 == maxElemtPid {

		if currentPid.p2 < maxElemtPid {

			var pid = currentPid
			currentPid.p2++
			currentPid.p3 = 0
			return pid
		}

	} else if currentPid.p3 == maxElemtPid {

		if currentPid.p2 == maxElemtPid {

			if currentPid.p1 < maxElemtPid {

				var pid = currentPid
				currentPid.p1++
				currentPid.p2 = 0
				currentPid.p3 = 0
				return pid
			}
		}

	}

	return Pid{-1, -1, -1}

}

func findProcess(Pid Pid) *process {

	for _, pp := range keeper {

		if pp.pid.p1 == Pid.p1 &&
			pp.pid.p2 == Pid.p2 &&
			pp.pid.p3 == Pid.p3 {
			return pp
		}
	}
	return nil
}

func indexProcess(Pid Pid) int {

	for index, pp := range keeper {

		if pp.pid.p1 == Pid.p1 &&
			pp.pid.p2 == Pid.p2 &&
			pp.pid.p3 == Pid.p3 {

			return index
		}
	}
	return -1
}

func startProcess(p *process, eternal bool) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println("RECOVER", a)
		}
	}()

	if eternal {

		keeper = append(keeper, p)
		for {
			select {
			case msg := <-p.recipient:

				switch msg.(type) {
				case string:
					if msg == "crash" {
						var i = indexProcess(p.pid)
						keeper = append(keeper[0:i], keeper[i:]...)
						break

					} else {
						p.function(msg)
					}
				default:
					p.function(msg)

				}
			}
		}

	} else {

		select {
		case msg := <-p.recipient:
			p.function(msg)
		}
		var i = indexProcess(p.pid)
		keeper = append(keeper[0:i], keeper[i:]...)
	}

}

// ToSliceString is cast (dynamic to []string)
func ToSliceString(lst Dynamic) []string {

	var newLst []string

	var object = reflect.ValueOf(lst)
	var iter = object.Index(0)
	var typ = iter.Type()

	for i := 0; i < iter.NumField(); i++ {

		newLst = append(newLst, typ.Field(i).Name)
	}

	return newLst
}

// Bang send message to recipient
func Bang(Pid Pid, message Dynamic) {

	var process = findProcess(Pid)
	process.recipient <- message
}

// Crash yes crahs! insted of panic use Crash for finish process
func Crash(p *process) {

	p.recipient <- "crash"
}

// Spawn new process
func Spawn(eternal bool, Fun func(Dynamic)) Pid {

	var Process = process{

		pid:       newPid(),
		recipient: make(chan Dynamic, maxRecipient),
		function:  Fun,
	}
	keeper = append(keeper, &Process)

	go startProcess(&Process, eternal)
	return Process.pid
}

// PidPrint show pid
func (p *Pid) PidPrint() {
	fmt.Printf("<%v.%v.%v>\n", p.p1, p.p2, p.p3)
}
