package goworker

type Worker interface {
        Work() error
}

type workerFunc func() Worker

var workers map[string]workerFunc

// Registers a goworker Worker.
// `class` refers to the Ruby name of the class which enqueues the job.
func Register(class string, workerFunc workerFunc) {
        workers[class] = workerFunc
}

func init() {
        workers = make(map[string]workerFunc)
}
