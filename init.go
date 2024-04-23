package tibrv

/*
	Choose between statical or dynamical linking to Tibco RendezVous libraries
	Put at FIRST the chosen option
*/

/*
//DYNAMIC RV LINKAGE
#cgo LDFLAGS: -ltibrvcm -ltibrvcmq -ltibrvft -ltibrv
//STATIC RV LINKAGE
#cgo LDFLAGS: -Wl,-Bstatic -ltibrvcm -ltibrvcmq -ltibrvft -ltibrv -Wl,-Bdynamic
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
