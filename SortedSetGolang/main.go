package main

import (
	"fmt"
	"time"
)

type Task struct {
}

func (t *Task) GetID() string {
	return "id-1"
}

type Worker struct {
}

func (w *Worker) ExecuteTask(t *Task) {
	fmt.Println(fmt.Sprintf("Executing task: %v in worker", t.GetID()))

}

//------------------------------------------------------------------------------------------------------------------------

type TaskExecutor struct {
	w     chan *Worker
	tasks chan *Task // stores all the tasks to be executed at time in the key of the map
}

func (t *TaskExecutor) Execute() {
	for {
		task := <-t.tasks
		worker := <-t.w
		go func() {
			worker.ExecuteTask(task) // once a task and worker is found, execute the task using that worker
			defer func() {
				t.w <- worker
			}()
		}()
	}
}

func (t *TaskExecutor) AddTasksToScheduler(tasks []*Task) {
	for _, task := range tasks {
		t.tasks <- task
	}
}

func (t *TaskExecutor) AddWorker(worker *Worker) {
	t.w <- worker
}

//---------------------------------------------------------------------------------------------------------------------

type TaskScheduler struct {
	tasks map[time.Time][]*Task
}

func (s *TaskScheduler) getUndoneTasksBeforeTime(t time.Time) []*Task {
	return s.tasks[t]
}

func (s *TaskScheduler) Runnable(e *TaskExecutor) {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				// check for the tasks in this time
				tasks := s.getUndoneTasksBeforeTime(t)
				e.AddTasksToScheduler(tasks)
				// TODO: Calculate next executiontime of task and add it to scheduler
			}
		}
	}()
	time.Sleep(100 * time.Hour)
	done <- true
}

//---------------------------------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------------------------------

func main() {

}
