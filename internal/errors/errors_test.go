package errors

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestIs(t *testing.T) {
	err1 := New("1", nil)
	erra := New("wrap 2", err1)
	errb := New("wrap 3", erra)

	err3 := New("3", nil)

	poser := &poser{msg: "either 1 or 3", f: func(err error) bool {
		return err == err1 || err == err3
	}}

	testCases := []struct {
		err    error
		target error
		match  bool
	}{
		{err: nil, target: nil, match: true},
		{err: err1, target: nil, match: false},
		{err: err1, target: err1, match: true},
		{err: erra, target: err1, match: true},
		{err: errb, target: err1, match: true},
		{err: err1, target: err3, match: false},
		{err: erra, target: err3, match: false},
		{err: errb, target: err3, match: false},
		{err: poser, target: err1, match: true},
		{err: poser, target: err3, match: true},
		{err: poser, target: erra, match: false},
		{err: poser, target: errb, match: false},
		{err: errorUncomparable{}, target: errorUncomparable{}, match: true},
		{err: errorUncomparable{}, target: &errorUncomparable{}, match: false},
		{err: &errorUncomparable{}, target: errorUncomparable{}, match: true},
		{err: &errorUncomparable{}, target: &errorUncomparable{}, match: false},
		{err: errorUncomparable{}, target: err1, match: false},
		{err: &errorUncomparable{}, target: err1, match: false},
		{err: multiErr{}, target: err1, match: false},
		{err: multiErr{err1, err3}, target: err1, match: true},
		{err: multiErr{err3, err1}, target: err1, match: true},
		{err: multiErr{err1, err3}, target: New("x", nil), match: false},
		{err: multiErr{err3, errb}, target: errb, match: true},
		{err: multiErr{err3, errb}, target: erra, match: true},
		{err: multiErr{err3, errb}, target: err1, match: true},
		{err: multiErr{errb, err3}, target: err1, match: true},
		{err: multiErr{poser}, target: err1, match: true},
		{err: multiErr{poser}, target: err3, match: true},
		{err: multiErr{nil}, target: nil, match: false},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			if got := Is(tc.err, tc.target); got != tc.match {
				t.Errorf("Is(%v, %v) = %v, want %v", tc.err, tc.target, got, tc.match)
			}
		})
	}
}

type poser struct {
	msg string
	f   func(error) bool
}

var poserPathErr = &fs.PathError{Op: "poser"}

func (p *poser) Error() string { return p.msg }

func (p *poser) Is(err error) bool { return p.f(err) }

func (p *poser) As(err any) bool {
	switch x := err.(type) {
	case **poser:
		*x = p
	case *errorT:
		*x = errorT{s: "poser"}
	case **fs.PathError:
		*x = poserPathErr
	default:
		return false
	}
	return true
}

type errorT struct{ s string }

func (e errorT) Error() string { return fmt.Sprintf("errorT(%s)", e.s) }

type multiErr []error

func (m multiErr) Error() string { return "multiError" }

func (m multiErr) Unwrap() []error { return []error(m) }

type errorUncomparable struct {
	f []string
}

func (errorUncomparable) Error() string {
	return "uncomparable error"
}

func (errorUncomparable) Is(target error) bool {
	_, ok := target.(errorUncomparable)
	return ok
}
