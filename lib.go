package fiftyonedegrees

/*
#cgo CFLAGS: -I . -Wimplicit-function-declaration
#cgo LDFLAGS: -lm
#include "51Degrees.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type FiftyoneDegrees struct {
	dataSet *C.fiftyoneDegreesDataSet
}

func NewFiftyoneDegrees(fileName, properties string) (*FiftyoneDegrees, error) {
	item := &FiftyoneDegrees{dataSet: new(C.fiftyoneDegreesDataSet)}
	status := C.fiftyoneDegreesInitWithPropertyString(C.CString(fileName), item.dataSet, C.CString(properties))
	if status != 0 {
		return nil, errors.New(fmt.Sprintln("InitWithPropertyString Error,Status:", status))
	}
	return item, nil
}
func (this *FiftyoneDegrees) Close() {
	C.fiftyoneDegreesDestroy(this.dataSet)
}

func (this *FiftyoneDegrees) Parse(userAgent string) string {
	ws := C.fiftyoneDegreesCreateWorkset(this.dataSet)
	defer C.fiftyoneDegreesFreeWorkset(ws)
	C.fiftyoneDegreesMatch(ws, C.CString(userAgent))
	resultLength := 50000
	buff := make([]byte, resultLength)
	length := int32(C.fiftyoneDegreesProcessDeviceJSON(ws, (*C.char)(unsafe.Pointer(&buff[0]))))
	result := buff[:length]
	return string(result)
}
