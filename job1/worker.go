package job1

import (
	"github.com/sirupsen/logrus"
)

// A Worker knows how to process jobs.
type Worker struct {
	log logrus.FieldLogger
}

// Create a Worker using a Log of type L.
func NewWorker(log logrus.FieldLogger) *Worker {
	return &Worker{
		log: log,
	}
}

// Process forces the worker to work on the given job.
func (w *Worker) Process(jobID string) {
	w.log.WithField("job_id", jobID).Info("starting job")
	// do something
}
