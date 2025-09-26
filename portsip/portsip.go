package portsip

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./lib -lportsip_c_serverbase

#include "include/portsip_c_serverbase.h"
#include "include/portsip_c_abstractcallbackdispatcher.h"
#include "include/portsip_c_icallbackparameters.h"
#include "goexport.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type SDKHandle struct {
	libsdkptr unsafe.Pointer
}

func Initialize(dispatcher *DispatcherHandle, singleAccountMode bool, transports string, maxCallSessions int,
	sipAgentString string, localIp string, audioDeviceLayer int, videoDeviceLayer int, TLSCertificatesRootPath string,
	TLSCipherList string, verifyTLSCertificate bool, encoderFactory unsafe.Pointer, encoderFactoryCallback unsafe.Pointer,
	decoderFactoryCallback unsafe.Pointer) (*SDKHandle, error) {

	cTransports := C.CString(transports)
	defer C.free(unsafe.Pointer(cTransports))
	cSipAgent := C.CString(sipAgentString)
	defer C.free(unsafe.Pointer(cSipAgent))
	cLocalIp := C.CString(localIp)
	defer C.free(unsafe.Pointer(cLocalIp))
	cTLSPath := C.CString(TLSCertificatesRootPath)
	defer C.free(unsafe.Pointer(cTLSPath))
	cCipherList := C.CString(TLSCipherList)
	defer C.free(unsafe.Pointer(cCipherList))

	var errorCode C.int
	libsdkptr := C.portsip_c_initialize(
		dispatcher.dispatcher,
		C.bool(singleAccountMode),
		cTransports,
		C.int(maxCallSessions),
		cSipAgent,
		cLocalIp,
		C.int(audioDeviceLayer),
		C.int(videoDeviceLayer),
		cTLSPath,
		cCipherList,
		C.bool(verifyTLSCertificate),
		encoderFactory,
		encoderFactoryCallback,
		decoderFactoryCallback,
		&errorCode,
	)

	if libsdkptr == nil {
		return nil, fmt.Errorf("Failed to initialize PortSIP SDK error code %d", errorCode)
	}

	return &SDKHandle{libsdkptr: libsdkptr}, nil
}

func (h *SDKHandle) Uninitialize() {
	if h.libsdkptr != nil {
		C.portsip_c_uninitialize(h.libsdkptr)
		h.libsdkptr = nil
	}
}

func (h *SDKHandle) DiscardReceivedPacket(state bool) {
	if h.libsdkptr != nil {
		C.portsip_c_discardreceivedpacket(h.libsdkptr, C.bool(state))
	}
}

func (h *SDKHandle) GetnNicNums() int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_getnicnums(h.libsdkptr))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) GetLocalIpAddress(index int, ip *string) int {
	if h.libsdkptr != nil {
		buf := make([]byte, 1024)
		cbuf := (*C.char)(unsafe.Pointer(&buf[0]))

		cret := C.portsip_c_getlocalipaddress(h.libsdkptr, C.int(index), cbuf, C.int(len(buf)))
		if cret >= 0 {
			*ip = C.GoString(cbuf)
		}
		return int(cret)
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) UpdateLocalIp(ip string) int {
	if h.libsdkptr != nil {
		cip := C.CString(ip)
		defer C.free(unsafe.Pointer(cip))

		return int(C.portsip_c_updatelocalip(h.libsdkptr, cip))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetLog(logType int, logLevel int, appName string, logFilePath string, maxFileSizeMB int, logObj uintptr) {
	if h.libsdkptr != nil {
		cAppName := C.CString(appName)
		defer C.free(unsafe.Pointer(cAppName))
		cLogPath := C.CString(logFilePath)
		defer C.free(unsafe.Pointer(cLogPath))

		C.portsip_c_setlog(
			h.libsdkptr,
			C.PORTSIP_LOG_TYPE(logType),
			C.PORTSIP_LOG_LEVEL(logLevel),
			cAppName,
			cLogPath,
			C.size_t(maxFileSizeMB),
			C.uintptr_t(logObj),
			C.fnLogCallback(C.GofnLogCallback),
		)
	}
}

func (h *SDKHandle) SetLicenseKey(key string) int {
	if h.libsdkptr != nil {
		cKey := C.CString(key)
		defer C.free(unsafe.Pointer(cKey))
		return int(C.portsip_c_setlicensekey(h.libsdkptr, cKey))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetRtpPortRange(startRTPPort int, endPort int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setrtpportrange(
			h.libsdkptr,
			C.int(startRTPPort),
			C.int(endPort),
		))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetSrtpPolicy(policy int, allowSrtpOverUnsecureTransport bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setsrtppolicy(
			h.libsdkptr,
			C.SRTP_POLICY(policy),
			C.bool(allowSrtpOverUnsecureTransport),
		))
	}
	return ECoreNotInitialized
}
