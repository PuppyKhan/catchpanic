package catchpanic

import (
	"testing"
)

const badDay = "very bad day"

type OrdinaryError struct{}

func (f OrdinaryError) Error() string {
	return "ordinary day"
}

func TestCatchPanic(t *testing.T) {
	tableTests := []struct {
		// givenPanictype int
		givenTestFunc func(t *testing.T) error

		wantErr error
	}{
		{
			// USER,
			func(t *testing.T) (err error) {
				defer ToError(&err)
				err = OrdinaryError{}
				panic(badDay)
				// return
			},

			SystemError{badDay},
		},
		{
			// POINTER,
			func(t *testing.T) (err error) {
				defer ToError(&err)
				err = OrdinaryError{}
				ptr := []string{}
				_ = ptr[42] // cause a panic
				t.Error("Should never get here")
				return
			},

			SystemError{"runtime error: index out of range"}, // r.(runtime.Error)
		},
		// { // doesn't look possible https://github.com/golang/go/issues/14162 "...reasonable to abort the program" -davecheney #facepalm
		// 	// MEMORY,
		// 	func(t *testing.T) (err error) {
		// 		defer ToError(&err)
		// 		err = OrdinaryError{}
		//		// stack
		// 		// var x [math.MaxInt32 / 32]int // "runtime: goroutine stack exceeds 1000000000-byte limit"
		// 		// var x [math.MaxInt32]int // "stack frame too large (>2GB)"
		//		// heap
		// 		// x := []int{}
		// 		// x = make([]int, math.MaxInt32, math.MaxInt32) // slow, but does not panic
		// 		type messyStruct struct {
		// 			Initial   byte
		// 			ID        int64
		// 			Name      string
		// 			SomeMap   map[string]string
		// 			SliceDice []string
		// 			Ptr       *int
		// 		}
		// 		x := []messyStruct{}
		// 		x = make([]messyStruct, math.MaxInt32, math.MaxInt32) // "signal: killed"
		// 		// http://stackoverflow.com/questions/43100127/golang-error-while-make-test-signal-killed
		// 		t.Error("Should never get here", x)
		// 		return
		// 	},

		// 	SystemError{"type?"}, // r.(runtime.Error)
		// },
	}

	for i, v := range tableTests {
		gotErr := v.givenTestFunc(t)

		if gotErr == nil {
			t.Errorf("%d: Failed to catch user panic, returned nothing", i)
		} else if _, ok := gotErr.(OrdinaryError); ok {
			t.Errorf("%d: Failed to catch user panic, returned func result", i)
		} else if gotErr.Error() != v.wantErr.Error() {
			t.Errorf("%d: Failed to catch user panic, returned %s", i, gotErr.Error())
		}
	}
}
