package portsip

/*
#include "include/portsip_c_types.h"
#include "goexport.h"

*/
import "C"
import (
	"log"
)

type (
	LogCallbackFunc func(logLevel int, message string)
)

//export GofnLogCallback
func GofnLogCallback(obj C.uintptr_t, level C.int, subsystem *C.char, appName *C.char, file *C.char, line C.int, message *C.char, messageWithHeaders *C.char) C.int {
	log.Printf("loglevel : %d , msg : s%  ", int(level), C.GoString(message))
	return 0
}

//export GofnRecordAudioRawCallback
func GofnRecordAudioRawCallback(obj C.uintptr_t, data *C.uchar, dataLength C.int, samplingFreqHz C.int) C.int {

	return 0
}
