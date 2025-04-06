// Package model holds common model structs and methods.
package model

import "errors"

// Error vars defines different model error constants.
var (
	ErrTitleNotFound      = errors.New("title not found inside http content page")
	ErrNoLinkTargetsFound = errors.New("no link targets found")
)
