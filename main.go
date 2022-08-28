package main

import (
	"github.com/alcortesm/go-interfaces-with-constructors/job1"
	"github.com/alcortesm/go-interfaces-with-constructors/job2"
	"github.com/alcortesm/go-interfaces-with-constructors/job3"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

func main() {
	do1()
	do2()
	do3()
}

// do1 uses package job1, which expect you to pass a logrus.Entry
func do1() {
	l := logrus.New()

	w := job1.NewWorker(l)
	w.Process("some_job_id")
}

// do2 uses package job2, which expect you to pass a type whose
// WithField method returns a logrus.Entry, which is better than
// job1 but still too coupled with logrus.
func do2() {
	l := logrus.New()

	w := job2.NewWorker(l)
	w.Process("some_job_id")
}

// do3 uses package job3, which expect you to pass a generic L logger
// that satisfies user3.Log[L], this is, its WithField method will
// return a type L. Package job3 is decoupled from the logrus library.
func do3() {
	// log using logrus
	{
		l := logrus.New()
		entry := logrus.NewEntry(l)

		w := job3.NewWorker(entry)
		w.Process("some_job_id")
	}

	// log using zap, via a zap-to-user3.Log wrapper
	{
		sugar := zap.NewExample().Sugar()
		defer sugar.Sync()
		l := &zapWrapper{inner: sugar}

		w := job3.NewWorker(l)
		w.Process("some_job_id")
	}
}

// zapWrapper wraps a zap logger to turn it into a type that satisfies user3.Log.
type zapWrapper struct {
	inner *zap.SugaredLogger
}

func (l *zapWrapper) WithField(key string, value any) *zapWrapper {
	return &zapWrapper{
		inner: l.inner.With(key, value),
	}
}

func (l *zapWrapper) Info(args ...any) {
	l.inner.Info(args...)
}
