package job3

// note how package user3 doesn't import logrus.

// Log is an interface to describe types that have WithFeild constructors
// and Info methods, for example *logrus.Entry or *zapWrapper.
type Log[L any] interface {
	WithField(key string, value interface{}) L
	Info(args ...any)
}

// A Worker knows how to process jobs.
//
// A Worker has a Log of type L that creates Logs of that same type L.
type Worker[L Log[L]] struct {
	log L
}

// Create a Worker using a Log of type L.
func NewWorker[L Log[L]](log L) *Worker[L] {
	return &Worker[L]{
		log: log,
	}
}

// Process forces the worker to work on the given job.
func (w *Worker[_]) Process(jobID string) {
	w.log.WithField("job_id", jobID).Info("starting job")
	// do something
}
