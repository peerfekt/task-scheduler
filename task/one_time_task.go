package task

import "time"

type OneTimeTask struct {
	date     time.Time
	callback func()
}

func NewOneTimeTask() *OneTimeTask {
	return &OneTimeTask{date: time.Time{}, callback: DEFAULT_FUNCTION}
}

func (o *OneTimeTask) SetDate(date time.Time) {
	o.date = date
}

func (o *OneTimeTask) SetCallback(callback func()) {
	o.callback = callback
}

func (o *OneTimeTask) GetDate() time.Time {
	return o.date
}

func (o *OneTimeTask) Run() {
	o.callback()
}

//Return bool wether it was executed or not
func (o *OneTimeTask) CheckTimeAndRun() bool {
	if o.date.Before(time.Now()) {
		o.date = time.Time{}
		o.Run()
		return true
	}
	return false
}
