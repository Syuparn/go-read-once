package goreadonce

import (
	"errors"
	"sync/atomic"
)

// Password represents user's password.
type Password struct {
	value *atomic.Pointer[string]
}

// NewPassword generates an user's password.
func NewPassword(s string) *Password {
	// TODO: validation check for password requirements
	v := &atomic.Pointer[string]{}
	v.Store(&s)
	return &Password{value: v}
}

// Value returns the password string. The password string can be obtained only once.
func (p *Password) Value() (string, error) {
	s := p.value.Swap(nil)
	// NOTE: nil comparison checks both types
	if s == (*string)(nil) {
		return "", errors.New("sensitive value has already been consumed")
	}
	return *s, nil
}

// String implements fmt.Stringer.
func (p *Password) String() string {
	// hide credentials
	return "{credential}"
}

// GoString implements fmt.GoStringer.
func (p *Password) GoString() string {
	// hide credentials (called by %#v)
	return "{credential}"
}
