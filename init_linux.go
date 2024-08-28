//go:build linux
// +build linux

package tibrv

/*
	Choose between statical or dynamical linking to Tibco RendezVous libraries
	Put at FIRST the chosen option
*/

/*
#cgo LDFLAGS: -Wl,-Bstatic -ltibrvcm64 -ltibrvcmq64 -ltibrvft64 -ltibrv64 -Wl,-Bdynamic
#include <tibrv/tibrv.h>
*/
import "C"

// Open init Tibco RendezVous Environment
func Open() error {
	if code := C.tibrv_Open(); code != C.TIBRV_OK {
		return NewRvError(code)
	}
	return nil
}

// Close destroy Tibco RendezVous Environment
func Close() error {
	if code := C.tibrv_Close(); code != C.TIBRV_OK {
		return NewRvError(code)
	}
	return nil
}
