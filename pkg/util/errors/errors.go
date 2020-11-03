package errors

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"strings"
)

type Errors []error

func (errs Errors) Error() string {
	s := make([]string, 0, len(errs))

	for _, err := range errs {
		s = append(s, err.Error())
	}

	return strings.Join(s, "\n")
}

func (errs Errors) AsError() error {
	if len(errs) == 0 {
		return nil // https://golang.org/doc/faq#nil_error
	}

	return errs
}
