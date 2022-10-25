# Task Scheduler in Golang

Basic implementation of a Task Scheduler in Golang.
This can be used to schedule function executions periodically or as a onetime task.

To-Do`s:
- [ ] unified interface for tasks - ability to add own task implementations
- [ ] scheduler returns waitgroup to indicate that execution is stopped due to no scheduled tasks
- [ ] unit tests
- [ ] optimization for multi-threading

[Example of use](example/example.go)