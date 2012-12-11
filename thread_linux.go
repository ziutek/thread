// Package thread is designed for manage OS thread parameters. Usually you
// need to call runtime.LockOSThread before use it.
package thread

import (
	"fmt"
	"github.com/ziutek/sched"
	"syscall"
)

type Thread struct {
	tid int
}

func Current() Thread {
	tid, _, e := syscall.RawSyscall(syscall.SYS_GETTID, 0, 0, 0)
	if e != 0 {
		panic(e)
	}
	return Thread{int(tid)}
}

func (t Thread) String() string {
	return fmt.Sprint("TID=", t.tid)
}

func (t Thread) SetSchedPolicy(policy sched.Policy, param *sched.Param) error {
	return sched.SetPolicy(t.tid, policy, param)
}

func (t Thread) SchedPolicy() (sched.Policy, error) {
	return sched.GetPolicy(t.tid)
}

func (t Thread) SetSchedParam(param *sched.Param) error {
	return sched.SetParam(t.tid, param)
}

func (t Thread) SchedParam(param *sched.Param) error {
	return sched.GetParam(t.tid, param)
}
