package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "goexport.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func (h *SDKHandle) CreateSipSession(userId int, sessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_createsipsession(h.libsdkptr, C.int(userId), C.long(sessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) DestorySipSession(sessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_destorysipsession(h.libsdkptr, C.long(sessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) CreateSdpOffer(sessionId int64,
	audioDirection int,
	videoDirection int,
	screenDirection int,
	newSdp *string) int {
	if h.libsdkptr != nil {
		buf := make([]byte, 1024)
		cbuf := (*C.char)(unsafe.Pointer(&buf[0]))

		cret := C.portsip_c_createsdpoffer(h.libsdkptr,
			C.long(sessionId),
			C.DIRECTION_MODE(audioDirection),
			C.DIRECTION_MODE(videoDirection),
			C.DIRECTION_MODE(screenDirection),
			cbuf,
			C.int(len(buf)))
		if cret >= 0 {
			*newSdp = C.GoString(cbuf)
		}
		return int(cret)
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) CreateSdpAnswer(sessionId int64,
	audioDirection int,
	videoDirection int,
	screenDirection int,
	newSdp *string) int {
	if h.libsdkptr != nil {
		buf := make([]byte, 1024)
		cbuf := (*C.char)(unsafe.Pointer(&buf[0]))

		cret := C.portsip_c_createsdpanswer(h.libsdkptr,
			C.long(sessionId),
			C.DIRECTION_MODE(audioDirection),
			C.DIRECTION_MODE(videoDirection),
			C.DIRECTION_MODE(screenDirection),
			cbuf,
			C.int(len(buf)))
		if cret >= 0 {
			*newSdp = C.GoString(cbuf)
		}
		return int(cret)
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) ReceivedRemoteSdp(sessionId int64,
	remoteSdp string,
	sdpSize int,
	audioDirection *int,
	videoDirection *int,
	screenDirection *int) int {
	if h.libsdkptr != nil {
		var audio, video, screen C.int
		cremoteSdp := C.CString(remoteSdp)
		defer C.free(unsafe.Pointer(cremoteSdp))

		cret := C.portsip_c_receivedremotesdp(h.libsdkptr,
			C.long(sessionId),
			cremoteSdp,
			C.int(sdpSize),
			(*C.DIRECTION_MODE)(unsafe.Pointer(&audio)),
			(*C.DIRECTION_MODE)(unsafe.Pointer(&video)),
			(*C.DIRECTION_MODE)(unsafe.Pointer(&screen)))
		if cret >= 0 {
			*audioDirection = int(audio)
			*videoDirection = int(video)
			*screenDirection = int(screen)
		}
		return int(cret)
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) RegisterDtmfDetection(observer unsafe.Pointer) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_registerdtmfdetection(h.libsdkptr, observer))
	}
	return ECoreNotInitialized
}
