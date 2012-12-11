package main

import (
	"fmt"
	"github.com/ziutek/sched"
	"github.com/ziutek/thread"
	"os"
	"runtime"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	runtime.LockOSThread()

	t := thread.Current()
	fmt.Println("Current thread:", t)

	var p sched.Param

	policy, err := t.SchedPolicy()
	checkErr(err)
	err = t.SchedParam(&p)
	checkErr(err)

	fmt.Printf("Current policy/priority: %s/%d\n", policy, p.Priority)

	p.Priority = sched.FIFO.MaxPriority()
	err = t.SetSchedPolicy(sched.FIFO, &p)
	checkErr(err)

	policy, err = t.SchedPolicy()
	checkErr(err)
	err = t.SchedParam(&p)
	checkErr(err)

	fmt.Printf("New policy/priority: %s/%d\n", policy, p.Priority)
}
