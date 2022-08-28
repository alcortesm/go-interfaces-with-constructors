package job2

import "github.com/sirupsen/logrus"

// A Worker knows how to process jobs.
type Worker struct {
	log Log
}

// Log is an interface that describes logrus.Entry.
type Log interface {
	WithField(key string, value interface{}) *logrus.Entry
	Info(args ...any)
}

// Create a Worker using a Log of type L.
func NewWorker(log Log) *Worker {
	return &Worker{
		log: log,
	}
}

// Process forces the worker to work on the given job.
func (w *Worker) Process(jobID string) {
	w.log.WithField("job_id", jobID).Info("starting job")
	// do something
}
