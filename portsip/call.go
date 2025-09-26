package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "goexport.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func (h *SDKHandle) Call(userId int, callee string, sendSdp bool, videoCall bool, headerNames string, headerValues string) int {
	if h.libsdkptr != nil {
		ccallee := C.CString(callee)
		defer C.free(unsafe.Pointer(ccallee))
		cheaderName := C.CString(headerNames)
		defer C.free(unsafe.Pointer(cheaderName))
		cheaderValue := C.CString(headerValues)
		defer C.free(unsafe.Pointer(cheaderValue))

		return int(C.portsip_c_call(h.libsdkptr, C.int(userId), ccallee, C.bool(sendSdp), C.bool(videoCall), cheaderName, cheaderValue))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) RejectCall(sessionId int64, code int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_rejectcall(h.libsdkptr, C.long(sessionId), C.int(code)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) Hangup(sessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_hangup(h.libsdkptr, C.long(sessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) AnswerCall(sessionId int64, videoCall bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_answercall(h.libsdkptr, C.long(sessionId), C.bool(videoCall)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) UpdateCall(sessionId int64, enableAudio bool, enableVideo bool, enableScreen bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_updatecall(h.libsdkptr, C.long(sessionId), C.bool(enableAudio), C.bool(enableVideo), C.bool(enableScreen)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) Hold(sessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_hold(h.libsdkptr, C.long(sessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) UnHold(sessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_unhold(h.libsdkptr, C.long(sessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) Refer(sessionId int64, referTo string) int {
	if h.libsdkptr != nil {
		creferTo := C.CString(referTo)
		defer C.free(unsafe.Pointer(creferTo))

		return int(C.portsip_c_refer(h.libsdkptr, C.long(sessionId), creferTo))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) Refer2(sessionId int64, referTo string, headerNames string, headerValues string) int {
	if h.libsdkptr != nil {
		creferTo := C.CString(referTo)
		defer C.free(unsafe.Pointer(creferTo))
		cheaderName := C.CString(headerNames)
		defer C.free(unsafe.Pointer(cheaderName))
		cheaderValue := C.CString(headerValues)
		defer C.free(unsafe.Pointer(cheaderValue))

		return int(C.portsip_c_refer2(h.libsdkptr, C.long(sessionId), creferTo, cheaderName, cheaderValue))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) AttendedRefer(sessionId int64, replaceSessionId int64, referTo string) int {
	if h.libsdkptr != nil {
		creferTo := C.CString(referTo)
		defer C.free(unsafe.Pointer(creferTo))

		return int(C.portsip_c_attendedrefer(h.libsdkptr, C.long(sessionId), C.long(replaceSessionId), creferTo))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) AccepRefer(userId int, referId int64, referSignaling string) int {
	if h.libsdkptr != nil {
		creferSignaling := C.CString(referSignaling)
		defer C.free(unsafe.Pointer(creferSignaling))

		return int(C.portsip_c_acceptrefer(h.libsdkptr, C.int(userId), C.long(referId), creferSignaling))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) RejectRefer(userId int, referId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_rejectrefer(h.libsdkptr, C.int(userId), C.long(referId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) MuteSession(sessionId int64, muteIncomingAudio bool, muteOutgoingAudio bool, muteIncomingVideo bool, muteOutgoingVideo bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_mutesession(h.libsdkptr,
			C.long(sessionId),
			C.bool(muteIncomingAudio),
			C.bool(muteOutgoingAudio),
			C.bool(muteIncomingVideo),
			C.bool(muteOutgoingVideo)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) HoldSession(sessionId int64, holdIncomingAudio bool, holdOutgoingAudio bool, holdIncomingVideo bool, holdOutgoingVideo bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_holdsession(h.libsdkptr,
			C.long(sessionId),
			C.bool(holdIncomingAudio),
			C.bool(holdOutgoingAudio),
			C.bool(holdIncomingVideo),
			C.bool(holdOutgoingVideo)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) Redirect(sessionId int64, toSessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_redirect(h.libsdkptr, C.long(sessionId), C.long(toSessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) ForwardCall(sessionId int64, forwardTo string) int {
	if h.libsdkptr != nil {
		cforwardTo := C.CString(forwardTo)
		defer C.free(unsafe.Pointer(cforwardTo))

		return int(C.portsip_c_forwardcall(h.libsdkptr, C.long(sessionId), cforwardTo))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SendDtmf(sessionId int64, dtmfMethod int, code int, dtmfDuration int, playDtmfTone bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_senddtmf(h.libsdkptr, C.long(sessionId), C.DTMF_METHOD(dtmfMethod), C.int(code), C.int(dtmfDuration), C.bool(playDtmfTone)))
	}
	return ECoreNotInitialized
}
