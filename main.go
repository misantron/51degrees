package main

/*
#cgo CFLAGS: -I . -Wimplicit-function-declaration
#include "51Degrees.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

func main() {
	properties := "DeviceType, IsMobile, IsSmartPhone, IsTablet, IsTv, HardwareName, HardwareVendor, HardwareModel, BrowserName, BrowserVersion, PlatformName, PlatformVersion, ScreenPixelsWidth, ScreenPixelsHeight"
	testUA := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.124 Safari/537.36"
	//item, err := NewFiftyoneDegrees("51Degrees-EnterpriseV3_1.dat", properties)
	item, err := NewFiftyoneDegrees("51Degrees-PremiumV3_2.dat", properties)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(item.Parse(testUA))
	item.Close()
}

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
