package main

import (
	"fmt"
	"testing"
)

func BenchmarkProcessCreation(b *testing.B) {
	scheduler := NewScheduler()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		scheduler.InsertProcess(NewProcess(fmt.Sprintf("%d", i)))
	}
}

func TestSchedule(t *testing.T) {
	scheduler := NewScheduler()
	previousPid, pid := 0, 0
	pid = scheduler.InsertProcess(NewProcess("./one"))
	if pid == previousPid {
		t.Fatalf("New PID is the same as previous PID")
	}
	pid = scheduler.InsertProcess(NewProcess("./two"))
	if pid == previousPid {
		t.Fatalf("New PID is the same as previous PID")
	}
	pid = scheduler.InsertProcess(NewProcess("./three"))
	if pid == previousPid {
		t.Fatalf("New PID is the same as previous PID")
	}
	pid = scheduler.InsertProcess(NewProcess("./four"))
	if pid == previousPid {
		t.Fatalf("New PID is the same as previous PID")
	}
	pid = scheduler.InsertProcess(NewProcess("./five"))
	if pid == previousPid {
		t.Fatalf("New PID is the same as previous PID")
	}
	preemptedPidCount := scheduler.RunAll()
	if preemptedPidCount != 5 {
		t.Fatalf("Expected scheduler.RunAll() to return 5 PIDs. Got %d instead", preemptedPidCount)
	}
}
