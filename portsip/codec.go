package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "goexport.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func (h *SDKHandle) AddAudioCodec(codecType int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_addaudiocodec(h.libsdkptr, C.AUDIOCODEC_TYPE(codecType)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) AddVideoCodec(codecType int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_addvideocodec(h.libsdkptr, C.VIDEOCODEC_TYPE(codecType)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) IsAudioCodecEmpty() bool {
	if h.libsdkptr != nil {
		return bool(C.portsip_c_isaudiocodecempty(h.libsdkptr))
	}
	return false
}

func (h *SDKHandle) IsVideoCodecEmpty() bool {
	if h.libsdkptr != nil {
		return bool(C.portsip_c_isvideocodecempty(h.libsdkptr))
	}
	return false
}

func (h *SDKHandle) SetAudioCodecPayloadtType(codecType int, payloadType int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setaudiocodecpayloadtype(h.libsdkptr, C.AUDIOCODEC_TYPE(codecType), C.int(payloadType)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoodecPayloadtType(codecType int, payloadType int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideocodecpayloadtype(h.libsdkptr, C.VIDEOCODEC_TYPE(codecType), C.int(payloadType)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) ClearAudioCodec() {
	if h.libsdkptr != nil {
		C.portsip_c_clearaudiocodec(h.libsdkptr)
	}
}

func (h *SDKHandle) ClearVideoCodec() {
	if h.libsdkptr != nil {
		C.portsip_c_clearvideocodec(h.libsdkptr)
	}
}

func (h *SDKHandle) SetAudioCodecParameter(codecType int, parameter string) int {
	if h.libsdkptr != nil {
		cparameter := C.CString(parameter)
		defer C.free(unsafe.Pointer(cparameter))

		return int(C.portsip_c_setaudiocodecparameter(h.libsdkptr, C.AUDIOCODEC_TYPE(codecType), cparameter))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideooCodecParameter(codecType int, parameter string) int {
	if h.libsdkptr != nil {
		cparameter := C.CString(parameter)
		defer C.free(unsafe.Pointer(cparameter))

		return int(C.portsip_c_setvideocodecparameter(h.libsdkptr, C.VIDEOCODEC_TYPE(codecType), cparameter))
	}
	return ECoreNotInitialized
}
