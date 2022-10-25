package main

import (
	"log"
	"time"

	"github.com/peerfekt/task-scheduler/scheduler"
	"github.com/peerfekt/task-scheduler/task"
)

func main() {
	runner := scheduler.NewScheduler()
	addExampleTask(runner)
	addExampleOneTimeTask(runner)
	runner.Run()
}

func addExampleTask(s *scheduler.Scheduler) {
	task := task.NewPeriodicTask()
	task.SetCallback(func() {
		log.Println("Iam overwriting the default task function")
	})
	task.SetInterval(time.Second * 10)
	s.AddPeriodicTask(task)
}

func addExampleOneTimeTask(s *scheduler.Scheduler) {
	task := task.NewOneTimeTask()
	task.SetCallback(func() {
		log.Println("Iam overwriting the default task function and iam executed once :)")
	})
	task.SetDate(time.Now().Add(time.Second * 5))

	s.AddOneTimeTask(task)
}
