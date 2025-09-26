package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "include/portsip_c_icallbackparameters.h"
#include "goexport.h"

*/
import "C"
import (
	"unsafe"
)

func DestroyParam(params unsafe.Pointer) {
	if params != nil {
		C.portsip_c_delcallbackparameters(params)
	}
}

func ParamGetEventtype(params unsafe.Pointer) int {
	if params == nil {
		return SIP_UNKNOWN
	}
	sip_event := C.portsip_c_params_geteventtype(params)
	return int(sip_event)
}

func ParamGetSessionid(params unsafe.Pointer) int64 {
	if params == nil {
		return -1
	}
	sessonid := C.portsip_c_params_getsessionid(params)
	return int64(sessonid)
}

func ParamGetUserid(params unsafe.Pointer) int {
	if params == nil {
		return -1
	}
	userid := C.portsip_c_params_getuserid(params)
	return int(userid)
}

func ParamGetCallerDisplayName(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	name := C.portsip_c_params_getcallerdisplayname(params)
	return C.GoString(name)
}

func ParamGetCalleeDisplayName(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	name := C.portsip_c_params_getcalleedisplayname(params)
	return C.GoString(name)
}

func ParamGetCaller(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	caller := C.portsip_c_params_getcaller(params)
	return C.GoString(caller)
}

func ParamGetCallee(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	caller := C.portsip_c_params_getcallee(params)
	return C.GoString(caller)
}

func ParamGetExistsEarlyMedia(params unsafe.Pointer) bool {
	if params == nil {
		return false
	}
	exist := C.portsip_c_params_getexistsearlymedia(params)
	return bool(exist)
}

func ParamGetStatusCode(params unsafe.Pointer) int {
	if params == nil {
		return -1
	}
	code := C.portsip_c_params_getstatuscode(params)
	return int(code)
}

func ParamGetStatusText(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getstatustext(params)
	return C.GoString(ret)
}

func ParamGetReferId(params unsafe.Pointer) int64 {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_getreferid(params)
	return int64(ret)
}

func ParamGetReferFrom(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getreferfrom(params)
	return C.GoString(ret)
}

func ParamGetRefertTo(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getreferto(params)
	return C.GoString(ret)
}

func ParamGetForwardTo(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getforwardto(params)
	return C.GoString(ret)
}

func ParamGetMessageData(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getmessagedata(params)
	return C.GoString((*C.char)(unsafe.Pointer(ret)))
}

func ParamGetMessageDataLength(params unsafe.Pointer) int {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_getmessagedatalength(params)
	return int(ret)
}

func ParamGetSignaling(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getsignaling(params)
	return C.GoString(ret)
}

func ParamGetAudioCodecs(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getaudiocodecs(params)
	return C.GoString(ret)
}

func ParamGetVideoCodecs(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getvideocodecs(params)
	return C.GoString(ret)
}

func ParamGetScreenCodecs(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getscreencodecs(params)
	return C.GoString(ret)
}

func ParamGetExistsVideo(params unsafe.Pointer) bool {
	if params == nil {
		return false
	}
	ret := C.portsip_c_params_getexistsvideo(params)
	return bool(ret)
}

func ParamGetExistsAudio(params unsafe.Pointer) bool {
	if params == nil {
		return false
	}
	ret := C.portsip_c_params_getexistsaudio(params)
	return bool(ret)
}

func ParamGetExistsScreen(params unsafe.Pointer) bool {
	if params == nil {
		return false
	}
	ret := C.portsip_c_params_getexistsscreen(params)
	return bool(ret)
}

func ParamGetWaitingMessageAccount(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getwaitingmessageaccount(params)
	return C.GoString(ret)
}

func ParamGetPresenceSubject(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getpresencesubject(params)
	return C.GoString(ret)
}

func ParamGetUrgentNewWaitingMessageCount(params unsafe.Pointer) int {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_geturgentnewwaitingmessagecount(params)
	return int(ret)
}

func ParamGetNewWaitingMessageCount(params unsafe.Pointer) int {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_getnewwaitingmessagecount(params)
	return int(ret)
}

func ParamGetUrgentOldWaitingMessageCount(params unsafe.Pointer) int {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_geturgentoldwaitingmessagecount(params)
	return int(ret)
}

func ParamGetOldWaitingMessageCount(params unsafe.Pointer) int {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_getoldwaitingmessagecount(params)
	return int(ret)
}

func ParamGetDtmfTone(params unsafe.Pointer) int {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_getdtmftone(params)
	return int(ret)
}

func ParamGetSubscribeId(params unsafe.Pointer) int64 {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_getsubscribeid(params)
	return int64(ret)
}

func ParamGetPlayedAudioFileName(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getplayedaudiofilename(params)
	return C.GoString(ret)
}

func ParamGetMimeType(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getmimetype(params)
	return C.GoString(ret)
}

func ParamGetSubMimeType(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getsubmimetype(params)
	return C.GoString(ret)
}

func ParamGetMessageId(params unsafe.Pointer) int64 {
	if params == nil {
		return -1
	}
	ret := C.portsip_c_params_getmessageid(params)
	return int64(ret)
}

func ParamGetBlfMonitoredUri(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getblfmonitoreduri(params)
	return C.GoString(ret)
}

func ParamGetBlfDialogState(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getblfdialogstate(params)
	return C.GoString(ret)
}

func ParamGetBlfDialogDirection(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getblfdialogdirection(params)
	return C.GoString(ret)
}

func ParamGetBlfDialogId(params unsafe.Pointer) string {
	if params == nil {
		return ""
	}
	ret := C.portsip_c_params_getblfdialogid(params)
	return C.GoString(ret)
}
