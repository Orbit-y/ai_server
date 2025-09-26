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
	VideoRawCallbackFunc func(obj cgo.Handle, sessionId int64, dtype int, width int, height int, data []byte, dataLength int)
)

var VideoRawCallbacks = make(map[int64]VideoRawCallbackFunc)
var videoMutex sync.RWMutex

func getVideoRawCallback(sessionId int64) VideoRawCallbackFunc {
	videoMutex.RLock()
	defer videoMutex.RUnlock()
	for key, value := range VideoRawCallbacks {
		if key == sessionId {
			return value
		}
	}
	return nil
}

func AddVideoRawCallbacks(sessionId int64, callback VideoRawCallbackFunc) {
	videoMutex.Lock()
	defer videoMutex.Unlock()
	if callback != nil {
		VideoRawCallbacks[sessionId] = callback
	}
}

func RemoveVideoRawCallbacks(sessionId int64) {
	videoMutex.Lock()
	defer videoMutex.Unlock()
	delete(VideoRawCallbacks, sessionId)
}

//export GofnVideoRawCallback
func GofnVideoRawCallback(obj C.uintptr_t, sessionId C.long, dtype C.int, width C.int, height C.int, data *C.uchar, dataLength C.int) C.int {
	cbuf := unsafe.Slice((*byte)(unsafe.Pointer(data)), int(dataLength))
	callback := getVideoRawCallback(int64(sessionId))
	h := cgo.Handle(obj)
	if callback != nil {
		callback(h, int64(sessionId), int(dtype), int(width), int(height), cbuf, int(dataLength))
	}
	return 0
}
