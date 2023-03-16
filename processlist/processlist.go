/*
	Process List in Go
		Written by Lucas Mendes <lucas AT lucasmendes.org>
*/

package main

import "fmt"

type Process struct {
	next    *Process
	command string
	pid     int
}

func (p *Process) SetPID(pid int) {
	fmt.Println("Process.Stop: Stopping", p.command)
}

func (p *Process) Run() {
	fmt.Printf("Process.Run: PID %d\t- Running %s\n", p.pid, p.command)
}

func NewProcess(command string) Process {
	return Process{
		next:    nil,
		command: command,
	}
}

type PIDManager struct {
	nextPid int
}

func (pm *PIDManager) AllocPIDForProcessList(pl *Process) int {
	var isPIDInUse = func(pid int, pl *Process) bool {
		for current := pl; current != nil; current = current.next {
			if current.pid == pid {
				return true
			}
		}
		return false
	}
	var pid int
	for pid = pm.nextPid; isPIDInUse(pid, pl); pid++ {
		fmt.Println("PIDManager.AllocPIDForProcessList: Checking for PID", pid)
	}
	fmt.Println("PIDManager.AllocPIDForProcessList: Using PID", pid)
	return pid
}

func NewPIDManager() PIDManager {
	return PIDManager{
		nextPid: 1,
	}
}

type Scheduler struct {
	processListHead *Process
	processListTail *Process
	processCount    int
	pidManager      PIDManager
}

func (s *Scheduler) RunAll() (preemptedPidCount int) {
	fmt.Println("--- Scheduling ---")
	preemptedPidCount = 0
	for current := s.processListHead; current != nil; current = current.next {
		current.Run()
		preemptedPidCount++
	}
	return
}

func (s *Scheduler) InsertProcess(process Process) (pid int) {
	newPid := s.pidManager.AllocPIDForProcessList(s.processListHead)
	if s.processListHead == nil {
		process.pid = pid
		s.processListHead = &process
		s.processListTail = &process
		s.processCount = 1
		fmt.Printf("Scheduler.InsertProcess: New process %s with PID %d\n",
			process.command, process.pid)
		return newPid
	}
	process.pid = pid
	s.processListTail.next = &process
	s.processListTail = &process
	s.processCount++
	fmt.Printf("Scheduler.InsertProcess: New process %s with PID %d\n",
		process.command, process.pid)
	return newPid
}

func NewScheduler() Scheduler {
	return Scheduler{
		pidManager: NewPIDManager(),
	}
}

func main() {
	s := NewScheduler()
	s.InsertProcess(NewProcess("./one"))
	s.InsertProcess(NewProcess("./two"))
	s.InsertProcess(NewProcess("./three"))
	s.InsertProcess(NewProcess("./four"))
	s.InsertProcess(NewProcess("./five"))
	s.RunAll()
}
