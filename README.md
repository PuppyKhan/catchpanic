# catchpanic
Recovering panics into errors

Experimenting with system memory panics

## Usage

Import package:
```go
import "github.com/PuppyKhan/catchpanic"
```

Add catchall with `defer`:
```go
func foo() (err error) {
	defer catchpanic.ToError(&err)

	// panic prone code
}
```

## Issues

There are actually 3 distinct types of panics in Go

* Recoverable user thrown panics
* Recoverable minor system panics
* Unrecoverable system panics

### Out Of Memory Panics

The benevolent overlords of Go deemed it impossible to catch all system panics, especially when it comes to running out of memory.

> Fundamentally, the role of a garbage collector is to present to the programmer the illusion of an infinite free store.
> 
> If that illusion is broken, is seems reasonable to abort the program.
 -[davecheney](https://github.com/davecheney) #facepalm

https://github.com/golang/go/issues/14162 


## Author

Luigi Kapaj <puppy at viahistoria.com>
