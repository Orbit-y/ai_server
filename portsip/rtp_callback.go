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
	RTPPacketCallback func(obj cgo.Handle, sessionId int64, mediaType int, direction int, data []byte, dataLength int)
)

var RTPPacketCallbacks = make(map[int64]RTPPacketCallback)
var rtpMutex sync.RWMutex

func AddRTPPacketCallbacks(sessionId int64, callback RTPPacketCallback) {
	rtpMutex.Lock()
	defer rtpMutex.Unlock()
	if callback != nil {
		RTPPacketCallbacks[sessionId] = callback
	}
}

func getRTPPacketCallback(sessionId int64) RTPPacketCallback {
	rtpMutex.RLock()
	defer rtpMutex.RUnlock()
	for key, value := range RTPPacketCallbacks {
		if key == sessionId {
			return value
		}
	}
	return nil
}

func RemoveRTPPacketCallbacks(sessionId int64) {
	rtpMutex.Lock()
	defer rtpMutex.Unlock()
	delete(RTPPacketCallbacks, sessionId)
}

//export GofnRTPPacketCallback
func GofnRTPPacketCallback(obj C.uintptr_t, sessionId C.long, mediaType C.int, direction C.int, RTPPacket *C.uchar, packetSize C.int) C.int {
	cbuf := unsafe.Slice((*byte)(unsafe.Pointer(RTPPacket)), int(packetSize))
	callback := getRTPPacketCallback(int64(sessionId))
	h := cgo.Handle(obj)
	if callback != nil {
		callback(h, int64(sessionId), int(mediaType), int(direction), cbuf, int(packetSize))
	}
	return 0
}
