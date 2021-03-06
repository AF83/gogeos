// Package geos provides support for creating and manipulating spatial data.
// At its core, it relies on the GEOS C library for the implementation of
// spatial operations and geometric algorithms.
package geos

/*
#include <geos_c.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>
#include <string.h>

void gogeos_notice_handler(const char *fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    fprintf(stderr, "NOTICE: ");
    vfprintf(stderr, fmt, ap);
    va_end(ap);
}

#define ERRLEN 256

char gogeos_last_err[ERRLEN];

void gogeos_error_handler(const char *fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    vsnprintf(gogeos_last_err, (size_t) ERRLEN, fmt, ap);
    va_end(ap);
}

char *gogeos_get_last_error(void) {
    return gogeos_last_err;
}

GEOSContextHandle_t gogeos_initGEOS() {
    return initGEOS_r(gogeos_notice_handler, gogeos_error_handler);
}

#cgo LDFLAGS: -lgeos_c
*/
import "C"

import (
	"fmt"
	"sync"
)

var (
	// Required for the thread-safe GEOS C API (the "*_r" functions).
	handle = C.gogeos_initGEOS()
	// Protects the handle from being used concurrently in multiple C threads.
	handlemu sync.Mutex
)

// XXX: store last error message from handler in a global var (chan?)

// Version returns the version of the GEOS C API in use.
func Version() string {
	return C.GoString(cGEOSversion())
}

// Error gets the last error that occured in the GEOS C API as a Go error type.
func Error() error {
	return fmt.Errorf("geos: %s", C.GoString(C.gogeos_get_last_error()))
}
