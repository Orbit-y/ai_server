package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "goexport.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func (h *SDKHandle) GetSipMessageHeaderValue(sipMessage string, headerName string, headerValue *string) int {
	if h.libsdkptr != nil {
		csipMessage := C.CString(sipMessage)
		defer C.free(unsafe.Pointer(csipMessage))

		cheaderName := C.CString(headerName)
		defer C.free(unsafe.Pointer(cheaderName))

		buf := make([]byte, 1024)
		cbuf := (*C.char)(unsafe.Pointer(&buf[0]))

		cret := C.portsip_c_getsipmessageheadervalue(h.libsdkptr, csipMessage, cheaderName, cbuf, C.int(len(buf)))
		if cret >= 0 {
			*headerValue = C.GoString(cbuf)
		}
		return int(cret)
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) AddSipMessageHeader(sessionId int64, methodName string, msgType int, headerName string, headerValue string) int64 {
	if h.libsdkptr != nil {
		cmethodName := C.CString(methodName)
		defer C.free(unsafe.Pointer(cmethodName))
		cheaderName := C.CString(headerName)
		defer C.free(unsafe.Pointer(cheaderName))
		cheaderValue := C.CString(headerValue)
		defer C.free(unsafe.Pointer(cheaderValue))

		return int64(C.portsip_c_addsipmessageheader(h.libsdkptr, C.long(sessionId), cmethodName, C.int(msgType), cheaderName, cheaderValue))
	}
	return int64(ECoreNotInitialized)
}

func (h *SDKHandle) RemoveAddSipMessageHeader(addedSipMessageId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_removeaddedsipmessageheader(h.libsdkptr, C.long(addedSipMessageId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) ClearAddSipMessageHeader() {
	if h.libsdkptr != nil {
		C.portsip_c_clearaddedsipmessageheaders(h.libsdkptr)
	}
}

func (h *SDKHandle) ModifySipMessageHeader(sessionId int64, methodName string, msgType int, headerName string, headerValue string) int64 {
	if h.libsdkptr != nil {
		cmethodName := C.CString(methodName)
		defer C.free(unsafe.Pointer(cmethodName))
		cheaderName := C.CString(headerName)
		defer C.free(unsafe.Pointer(cheaderName))
		cheaderValue := C.CString(headerValue)
		defer C.free(unsafe.Pointer(cheaderValue))

		return int64(C.portsip_c_modifysipmessageheader(h.libsdkptr, C.long(sessionId), cmethodName, C.int(msgType), cheaderName, cheaderValue))
	}
	return int64(ECoreNotInitialized)
}

func (h *SDKHandle) RemoveModifiedSipMessageHeader(modifiedSipMessageId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_removemodifiedsipmessageheader(h.libsdkptr, C.long(modifiedSipMessageId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) ClearModifiedSipMessageHeader() {
	if h.libsdkptr != nil {
		C.portsip_c_clearmodifiedsipmessageheaders(h.libsdkptr)
	}
}

func (h *SDKHandle) AddSupportedMimeType(methodName string, mimeType string, subMimeType string) int64 {
	if h.libsdkptr != nil {
		cmethodName := C.CString(methodName)
		defer C.free(unsafe.Pointer(cmethodName))
		cmimeType := C.CString(mimeType)
		defer C.free(unsafe.Pointer(cmimeType))
		csubMimeType := C.CString(subMimeType)
		defer C.free(unsafe.Pointer(csubMimeType))

		return int64(C.portsip_c_addsupportedmimetype(h.libsdkptr, cmethodName, cmimeType, csubMimeType))
	}
	return int64(ECoreNotInitialized)
}

func (h *SDKHandle) SendOptions(userId int, to string, sdp string) int {
	if h.libsdkptr != nil {
		cto := C.CString(to)
		defer C.free(unsafe.Pointer(cto))
		csdp := C.CString(sdp)
		defer C.free(unsafe.Pointer(csdp))

		return int(C.portsip_c_sendoptions(h.libsdkptr, C.int(userId), cto, csdp))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SendInfo(sessionId int64, mimeType string, subMimeType string, infoContents string) int {
	if h.libsdkptr != nil {
		cmimeType := C.CString(mimeType)
		defer C.free(unsafe.Pointer(cmimeType))
		csubMimeType := C.CString(subMimeType)
		defer C.free(unsafe.Pointer(csubMimeType))
		cinfoContents := C.CString(infoContents)
		defer C.free(unsafe.Pointer(cinfoContents))

		return int(C.portsip_c_sendinfo(h.libsdkptr, C.long(sessionId), cmimeType, csubMimeType, cinfoContents))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SendSubscription(userId int, to string, eventName string) int64 {
	if h.libsdkptr != nil {
		cto := C.CString(to)
		defer C.free(unsafe.Pointer(cto))
		ceventName := C.CString(eventName)
		defer C.free(unsafe.Pointer(ceventName))

		return int64(C.portsip_c_sendsubscription(h.libsdkptr, C.int(userId), cto, ceventName))
	}
	return int64(ECoreNotInitialized)
}

func (h *SDKHandle) TerminateSubscription(userId int, subscriptionId int64) {
	if h.libsdkptr != nil {
		C.portsip_c_terminatesubscription(h.libsdkptr, C.int(userId), C.long(subscriptionId))
	}
}

func (h *SDKHandle) SendMessage(userId int, sessionId int64, mimeType string, subMimeType string, message string, messageLength int) int64 {
	if h.libsdkptr != nil {
		cmimeType := C.CString(mimeType)
		defer C.free(unsafe.Pointer(cmimeType))
		csubMimeType := C.CString(subMimeType)
		defer C.free(unsafe.Pointer(csubMimeType))
		cmessage := (*C.uchar)(unsafe.Pointer(C.CString(message)))
		defer C.free(unsafe.Pointer(cmessage))

		return int64(C.portsip_c_sendmessage(h.libsdkptr, C.int(userId), C.long(sessionId), cmimeType, csubMimeType, cmessage, C.int(messageLength)))
	}
	return int64(ECoreNotInitialized)
}

func (h *SDKHandle) SendOutOfDialogMessage(userId int,
	to string,
	mimeType string,
	subMimeType string,
	isSMS bool,
	message string,
	messageLength int,
	displayName string) int64 {
	if h.libsdkptr != nil {
		cto := C.CString(to)
		defer C.free(unsafe.Pointer(cto))
		cmimeType := C.CString(mimeType)
		defer C.free(unsafe.Pointer(cmimeType))
		csubMimeType := C.CString(subMimeType)
		defer C.free(unsafe.Pointer(csubMimeType))
		cmessage := (*C.uchar)(unsafe.Pointer(C.CString(message)))
		defer C.free(unsafe.Pointer(cmessage))
		cdisplayName := C.CString(displayName)
		defer C.free(unsafe.Pointer(cdisplayName))

		return int64(C.portsip_c_sendoutofdialogmessage(h.libsdkptr,
			C.int(userId),
			cto,
			cmimeType,
			csubMimeType,
			C.bool(isSMS),
			cmessage,
			C.int(messageLength),
			cdisplayName))
	}
	return int64(ECoreNotInitialized)
}

func (h *SDKHandle) PresenceSubscribe(userId int, contact string, subject string) int {
	if h.libsdkptr != nil {
		ccontact := C.CString(contact)
		defer C.free(unsafe.Pointer(ccontact))
		csubject := C.CString(subject)
		defer C.free(unsafe.Pointer(csubject))

		return int(C.portsip_c_presencesubscribe(h.libsdkptr, C.int(userId), ccontact, csubject))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) PresenceAcceptSubscribe(userId int, subscribeId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_presenceacceptsubscribe(h.libsdkptr, C.int(userId), C.long(subscribeId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) PresenceRejectSubscribe(userId int, subscribeId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_presencerejectsubscribe(h.libsdkptr, C.int(userId), C.long(subscribeId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetPresenceStatus(userId int, subscribeId int64, statusText string) int {
	if h.libsdkptr != nil {
		cstatusText := C.CString(statusText)
		defer C.free(unsafe.Pointer(cstatusText))
		return int(C.portsip_c_setpresencestatus(h.libsdkptr, C.int(userId), C.long(subscribeId), cstatusText))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetDefaultPublicationTime(userId int, secs uint) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setdefaultpublicationtime(h.libsdkptr, C.int(userId), C.uint(secs)))
	}
	return ECoreNotInitialized
}
