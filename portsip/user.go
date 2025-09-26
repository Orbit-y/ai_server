package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "goexport.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func (h *SDKHandle) AddUser(userName, displayName, authName, password, useTransportName, sipDomain, sipServerAddr string,
	sipServerPort int, stunServerAddr string, stunServerPort int, outboundServerAddr string, outboundServerPort int) int {
	if h.libsdkptr != nil {
		cUserName := C.CString(userName)
		defer C.free(unsafe.Pointer(cUserName))
		cDisplay := C.CString(displayName)
		defer C.free(unsafe.Pointer(cDisplay))
		cAuth := C.CString(authName)
		defer C.free(unsafe.Pointer(cAuth))
		cPass := C.CString(password)
		defer C.free(unsafe.Pointer(cPass))
		cTransport := C.CString(useTransportName)
		defer C.free(unsafe.Pointer(cTransport))
		cDomain := C.CString(sipDomain)
		defer C.free(unsafe.Pointer(cDomain))
		cSipServer := C.CString(sipServerAddr)
		defer C.free(unsafe.Pointer(cSipServer))
		cStun := C.CString(stunServerAddr)
		defer C.free(unsafe.Pointer(cStun))
		cOutbound := C.CString(outboundServerAddr)
		defer C.free(unsafe.Pointer(cOutbound))

		return int(C.portsip_c_adduser(
			h.libsdkptr,
			cUserName,
			cDisplay,
			cAuth,
			cPass,
			cTransport,
			cDomain,
			cSipServer,
			C.int(sipServerPort),
			cStun,
			C.int(stunServerPort),
			cOutbound,
			C.int(outboundServerPort),
		))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) UpdateUser(userId int, userName, displayName, authName, password, useTransportName, sipDomain, sipServerAddr string,
	sipServerPort int, stunServerAddr string, stunServerPort int, outboundServerAddr string, outboundServerPort int) int {
	if h.libsdkptr != nil {
		cUserName := C.CString(userName)
		defer C.free(unsafe.Pointer(cUserName))
		cDisplay := C.CString(displayName)
		defer C.free(unsafe.Pointer(cDisplay))
		cAuth := C.CString(authName)
		defer C.free(unsafe.Pointer(cAuth))
		cPass := C.CString(password)
		defer C.free(unsafe.Pointer(cPass))
		cTransport := C.CString(useTransportName)
		defer C.free(unsafe.Pointer(cTransport))
		cDomain := C.CString(sipDomain)
		defer C.free(unsafe.Pointer(cDomain))
		cSipServer := C.CString(sipServerAddr)
		defer C.free(unsafe.Pointer(cSipServer))
		cStun := C.CString(stunServerAddr)
		defer C.free(unsafe.Pointer(cStun))
		cOutbound := C.CString(outboundServerAddr)
		defer C.free(unsafe.Pointer(cOutbound))

		return int(C.portsip_c_updateuser(
			h.libsdkptr,
			C.int(userId),
			cUserName,
			cDisplay,
			cAuth,
			cPass,
			cTransport,
			cDomain,
			cSipServer,
			C.int(sipServerPort),
			cStun,
			C.int(stunServerPort),
			cOutbound,
			C.int(outboundServerPort),
		))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) RemoveUser(userId int) {
	if h.libsdkptr != nil {
		C.portsip_c_removeuser(
			h.libsdkptr,
			C.int(userId),
		)
	}
}

func (h *SDKHandle) RegisterServer(userId, regExpires, retryTimes int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_registerserver(h.libsdkptr, C.int(userId), C.int(regExpires), C.int(retryTimes)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) UnRegisterServer(userId, waitMS int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_unregisterserver(h.libsdkptr, C.int(userId), C.int(waitMS)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetDisplayName(userId int, displayName string) int {
	if h.libsdkptr != nil {
		cDisplay := C.CString(displayName)
		defer C.free(unsafe.Pointer(cDisplay))

		return int(C.portsip_c_setdisplayname(h.libsdkptr, C.int(userId), cDisplay))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetInstanceId(uuid string) int {
	if h.libsdkptr != nil {
		cuuid := C.CString(uuid)
		defer C.free(unsafe.Pointer(cuuid))

		return int(C.portsip_c_setinstanceid(h.libsdkptr, cuuid))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetUserAgent(sipAgentString string) int {
	if h.libsdkptr != nil {
		csipAgentString := C.CString(sipAgentString)
		defer C.free(unsafe.Pointer(csipAgentString))

		return int(C.portsip_c_setuseragent(h.libsdkptr, csipAgentString))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetReliableProvisional(mode int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setreliableprovisional(h.libsdkptr, C.int(mode)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) Enable3gpptags(userId int, enable bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enable3gpptags(h.libsdkptr, C.int(userId), C.bool(enable)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableCallbackSignaling(enableSending bool, enableReceived bool) {
	C.portsip_c_enablecallbacksignaling(h.libsdkptr, C.bool(enableSending), C.bool(enableReceived))
}

func (h *SDKHandle) EnableCallForward(userId int, forBusyOnly bool, forwardTo string) int {
	if h.libsdkptr != nil {
		cforwardTo := C.CString(forwardTo)
		defer C.free(unsafe.Pointer(cforwardTo))

		return int(C.portsip_c_enablecallforward(h.libsdkptr, C.int(userId), C.bool(forBusyOnly), cforwardTo))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) DisableCallForward(userId int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_disablecallforward(h.libsdkptr, C.int(userId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableSessionTimer(timerSeconds int, refreshMode int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enablesessiontimer(h.libsdkptr, C.int(timerSeconds), C.SESSION_REFRESH_MODE(refreshMode)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) DisableSessionTimer() int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_disablesessiontimer(h.libsdkptr))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetDonotDisturb(userId int, state bool) {
	if h.libsdkptr != nil {
		C.portsip_c_setdonotdisturb(h.libsdkptr, C.int(userId), C.bool(state))
	}
}

func (h *SDKHandle) SetRtpKeepAlive(state bool, keepAlivePayloadType int, deltaTransmitTimeMS int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setrtpkeepalive(h.libsdkptr, C.bool(state), C.int(keepAlivePayloadType), C.int(deltaTransmitTimeMS)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetKeepAliveTime(keepAliveTime int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setkeepalivetime(h.libsdkptr, C.int(keepAliveTime)))
	}
	return ECoreNotInitialized
}
