package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	if len(os.Args) < 2 {
		println("go run goadvent2015 <count>")
		return 1
	}
	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println(err)
		println("go run goadvent2015 <count>")
		return 1
	}

	if count < 10000 {
		println("count should be greater than 10000")
		return 1
	}

	println("Go version: " + runtime.Version())
	RunStoreRawCPointer(count)
	RunStoreUnsafePointer(count)
	RunStoreUintptr(count)
	return 0
}

func timeGC() {
	start := time.Now()
	runtime.GC()
	fmt.Printf("\tGC time: %d μs\n", time.Since(start).Nanoseconds())
}

func RunStoreRawCPointer(count int) {
	fmt.Println("Starting run for *C.char GC times...")
	list := make([]*C.char, count)
	for i := range list {
		list[i] = C.CString(strconv.Itoa(i))
	}
	defer func() {
		for _, p := range list {
			C.free(unsafe.Pointer(p))
		}
	}()

	// run GC once to "normalize"
	runtime.GC()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected, and to prevent them from being optimized away
		fmt.Fprintf(ioutil.Discard, "value @ index %d: %v\n", i*(count/6), list[i*(count/6)])

		timeGC()
	}
}

func RunStoreUnsafePointer(count int) {
	fmt.Println("Starting run for unsafe.Pointer GC times...")
	list := make([]unsafe.Pointer, count)
	for i := range list {
		list[i] = unsafe.Pointer(C.CString(strconv.Itoa(i)))
	}
	defer func() {
		for _, p := range list {
			C.free(p)
		}
	}()

	// run GC once to "normalize"
	runtime.GC()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected, and to prevent them from being optimized away
		fmt.Fprintf(ioutil.Discard, "value @ index %d: %v\n", i*(count/6), list[i*(count/6)])

		timeGC()
	}
}

func RunStoreUintptr(count int) {
	fmt.Println("Starting run for uintptr GC times...")
	list := make([]uintptr, count)
	for i := range list {
		list[i] = uintptr(unsafe.Pointer(C.CString(strconv.Itoa(i))))
	}
	defer func() {
		for _, p := range list {
			C.free(unsafe.Pointer(p))
		}
	}()

	// run GC once to "normalize"
	runtime.GC()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected, and to prevent them from being optimized away
		fmt.Fprintf(ioutil.Discard, "value @ index %d: %v\n", i*(count/6), list[i*(count/6)])

		timeGC()
	}
}
