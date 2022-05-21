package testutil

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Diff compares two items and returns a human-readable diff string. If the
// items are equal, the string is empty.
func Diff[T any](got, want T, opts ...cmp.Option) string {
	if len(opts) == 0 {
		opts = cmp.Options{
			cmp.Exporter(func(reflect.Type) bool { return true }),
			cmpopts.EquateEmpty(),
		}
	}
	diff := cmp.Diff(got, want, opts...)
	if diff != "" {
		return "\n-got +want\n" + diff
	}
	return ""
}

// Callers prints the stack trace of everything up til the line where Callers()
// was invoked.
func Callers() string {
	var pc [50]uintptr
	n := runtime.Callers(2, pc[:]) // skip runtime.Callers + Callers
	callsites := make([]string, 0, n)
	frames := runtime.CallersFrames(pc[:n])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callsites = append(callsites, frame.File+":"+strconv.Itoa(frame.Line))
	}
	callsites = callsites[:len(callsites)-1] // skip testing.tRunner
	if len(callsites) == 1 {
		return ""
	}
	var b strings.Builder
	b.WriteString("\n[")
	for i := len(callsites) - 1; i >= 0; i-- {
		if i < len(callsites)-1 {
			b.WriteString(" -> ")
		}
		b.WriteString(filepath.Base(callsites[i]))
	}
	b.WriteString("]")
	return b.String()
}
