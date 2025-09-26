package portsip

/*
#include "include/portsip_c_types.h"
#include "goexport.h"

*/
import "C"
import (
	"runtime/cgo"
	"sync"
	"unsafe"
)

type (
	AudioRawCallbackFunc func(obj cgo.Handle, sessionId int64, dtype int, data []byte, dataLength int, samplingFreqHz int)
)

var AudioRawCallbacks = make(map[int64]AudioRawCallbackFunc)
var audioMutex sync.RWMutex

func AddAudioRawCallbacks(sessionId int64, callback AudioRawCallbackFunc) {
	audioMutex.Lock()
	defer audioMutex.Unlock()
	if callback != nil {
		AudioRawCallbacks[sessionId] = callback
	}
}

func getAudioRawCallback(sessionId int64) AudioRawCallbackFunc {
	audioMutex.RLock()
	defer audioMutex.RUnlock()
	for key, value := range AudioRawCallbacks {
		if key == sessionId {
			return value
		}
	}
	return nil
}

func RemoveAudioRawCallbacks(sessionId int64) {
	audioMutex.Lock()
	defer audioMutex.Unlock()
	delete(AudioRawCallbacks, sessionId)
}

//export GofnAudioRawCallback
func GofnAudioRawCallback(obj C.uintptr_t, sessionId C.long, dtype C.int, data *C.uchar, dataLength C.int, samplingFreqHz C.int) C.int {
	cbuf := unsafe.Slice((*byte)(unsafe.Pointer(data)), int(dataLength))
	callback := getAudioRawCallback(int64(sessionId))
	h := cgo.Handle(obj)
	if callback != nil {
		callback(h, int64(sessionId), int(dtype), cbuf, int(dataLength), int(samplingFreqHz))
	}
	return 0
}
