package scheduler

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/peerfekt/task-scheduler/task"
)

type Scheduler struct {
	periodicTasks map[string]*task.PeriodicTask
	oneTimeTasks  map[string]*task.OneTimeTask
	quit          chan bool
}

func NewScheduler() *Scheduler {
	return &Scheduler{periodicTasks: make(map[string]*task.PeriodicTask), oneTimeTasks: make(map[string]*task.OneTimeTask)}
}

//return id of the task. Can be used to delete it from the Scheduler
func (s *Scheduler) AddPeriodicTask(task *task.PeriodicTask) string {
	key := uuid.New().String()
	s.periodicTasks[key] = task
	return key
}

//return id of the task. Can be used to delete it from the Scheduler
func (s *Scheduler) AddOneTimeTask(task *task.OneTimeTask) string {
	key := uuid.New().String()
	s.oneTimeTasks[key] = task
	return key
}

//Deletes a OneTimeTask from the TaskList
func (s *Scheduler) DeleteOneTimeTask(key string) {
	delete(s.oneTimeTasks, key)
}

//Deletes a PeriodicTask from the TaskList
func (s *Scheduler) DeletePeriodicTask(key string) {
	delete(s.periodicTasks, key)
}

func (s *Scheduler) Run() {
	log.Println("Starting the Scheduler")
	for {
		select {
		case <-s.quit:
			log.Println("Scheduler was stopped")
			return
		default:
			s.tick()
		}
		time.Sleep(time.Second)
	}
}

//Tries to run OneTimeTask and deletes it depending on if it was run or not
func (s *Scheduler) runOneTimeTaskAndDelete(key string) {
	run := s.oneTimeTasks[key].CheckTimeAndRun()
	if run {
		s.DeleteOneTimeTask(key)
	}
}

func (s *Scheduler) tick() {
	for _, task := range s.periodicTasks {
		go (*task).RunAndReschedule()
	}
	for key := range s.oneTimeTasks {
		go s.runOneTimeTaskAndDelete(key)
	}
}

//Stops the execution of the scheduler
func (s *Scheduler) Quit() {
	s.quit <- true
}
