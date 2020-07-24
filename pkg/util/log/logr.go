package log

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/sirupsen/logrus"
)

type logrWrapper struct {
	entry *logrus.Entry
	level int
}

func (lw *logrWrapper) Enabled() bool {
	return lw.level <= int(logrus.GetLevel())
}

func (lw *logrWrapper) Error(err error, msg string, keysAndValues ...interface{}) {
	lw.withKeysAndValues(keysAndValues).Error(msg, " ", err)
}

func (lw *logrWrapper) withKeysAndValues(keysAndValues []interface{}) *logrus.Entry {
	if len(keysAndValues) == 0 {
		return lw.entry
	}
	key := ""
	fields := logrus.Fields{}
	for _, item := range keysAndValues {
		if key == "" {
			key = fmt.Sprint(item)
		} else {
			fields[key] = fmt.Sprint(item)
			key = ""
		}
	}
	if key != "" {
		// key with no value
		fields[key] = ""
	}

	return lw.entry.WithFields(fields)
}

func (lw *logrWrapper) Info(msg string, keysAndValues ...interface{}) {
	if !lw.Enabled() {
		return
	}
	lw.withKeysAndValues(keysAndValues).Info(msg)
}

func (lw *logrWrapper) V(level int) logr.InfoLogger {
	return &logrWrapper{
		entry: lw.entry,
		level: level,
	}
}

func (lw *logrWrapper) WithValues(keysAndValues ...interface{}) logr.Logger {
	return &logrWrapper{
		entry: lw.withKeysAndValues(keysAndValues),
		level: lw.level,
	}
}

func (lw *logrWrapper) WithName(name string) logr.Logger {
	return &logrWrapper{
		entry: lw.withKeysAndValues([]interface{}{name, ""}),
		level: lw.level,
	}
}

func LogrWrapper(logger *logrus.Entry) logr.Logger {
	return &logrWrapper{
		entry: logger,
		level: int(logrus.GetLevel()),
	}
}
