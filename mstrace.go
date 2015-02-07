package mseed

//#cgo CFLAGS: -I/usr/local/include
//#cgo LDFLAGS: /usr/local/lib/libmseed.a
//#include <libmseed.h>
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

type MSTrace _Ctype_MSTrace

func NewMSTrace() *MSTrace {
	return (*MSTrace)(C.mst_init(nil))
}

func FreeMSTrace(t *MSTrace) {
	C.mst_free((**_Ctype_struct_MSTrace_s)((unsafe.Pointer)(&t)))
}

func (t *MSTrace) print() {

	sec := int64(t.starttime) / 1000000
	nsec := 1000 * (int64(t.starttime) % 1000000)

	starttime := time.Unix(sec, nsec)

	fmt.Printf("print trace [%s] %d %s\n", C.GoStringN(&t.station[0], 11), int(t.samplecnt), starttime.UTC().Format(time.RFC3339Nano))
}
