package catchpanic

import (
	"fmt"
)

type MemoryError struct {
	Original string
}

func (f MemoryError) Error() string {
	return fmt.Sprintf("Internal Memory Error: %s", f.Original)
}

type SystemError struct {
	Original string
}

func (f SystemError) Error() string {
	return fmt.Sprintf("Internal System Error: %s", f.Original)
}

// ToError recovers from a panic and returns it as an error instead
func ToError(parentErr *error) {
	r := recover()
	if r == nil {
		return
	}

	// // Is this necessary, or will it include the mem issue we want to handle?
	// if _, ok := r.(runtime.Error); ok {
	// 	// fmt.Println("Runtime error...")
	// 	panic(r)
	// }

	retErr := SystemError{""}
	err, ok := r.(error)
	if !ok {
		err, ok := r.(string)
		if !ok {
			retErr.Original = fmt.Sprintf("%v", r)
		} else {
			retErr.Original = err
		}
	} else {
		retErr.Original = err.Error()
	}
	*parentErr = retErr
	return
}
