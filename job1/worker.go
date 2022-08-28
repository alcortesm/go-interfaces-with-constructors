package job1

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// Repeat returns a new string consisting of count copies of the string s.
//
// It returns s if count is negative.
func Repeat(s string, count int, log logrus.FieldLogger) string {
	if count < 0 {
		log.WithField("count", count).Info("protecting strings.Repeat from a negative count")
		return s
	}

	return strings.Repeat(s, count)
}

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
