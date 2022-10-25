package task

import (
	"time"
)

type PeriodicTask struct {
	callback            func()
	interval            time.Duration
	last_execution_time time.Time
}

func NewPeriodicTask() *PeriodicTask {
	return &PeriodicTask{interval: DEFAULT_INTERVAL, callback: DEFAULT_FUNCTION, last_execution_time: time.Time{}}
}

func (p *PeriodicTask) SetInterval(interval time.Duration) {
	p.interval = interval
}

func (p *PeriodicTask) SetCallback(callback func()) {
	p.callback = callback
}

func (p *PeriodicTask) GetInterval() time.Duration {
	return p.interval
}

func (p *PeriodicTask) Run() {
	p.callback()
}

func (p *PeriodicTask) RunAndReschedule() {
	if p.last_execution_time.Add(p.interval).Before(time.Now()) {
		p.last_execution_time = time.Now()
		p.Run()
	}
}
