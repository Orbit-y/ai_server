package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "goexport.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func (h *SDKHandle) GetNumOfRecordingDevices() int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_getnumofrecordingdevices(h.libsdkptr))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) GetNumOfPlayoutDevices() int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_getnumofplayoutdevices(h.libsdkptr))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) GetRecordingDeviceName(index int, nameUTF8 *string) int {
	if h.libsdkptr != nil {
		buf := make([]byte, 1024)
		cbuf := (*C.char)(unsafe.Pointer(&buf[0]))

		cret := C.portsip_c_getrecordingdevicename(h.libsdkptr, C.int(index), cbuf, C.int(len(buf)))
		if cret >= 0 {
			*nameUTF8 = C.GoString(cbuf)
		}
		return int(cret)
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) GetPlayoutDeviceName(index int, nameUTF8 *string) int {
	if h.libsdkptr != nil {
		buf := make([]byte, 1024)
		cbuf := (*C.char)(unsafe.Pointer(&buf[0]))

		cret := C.portsip_c_getplayoutdevicename(h.libsdkptr, C.int(index), cbuf, C.int(len(buf)))
		if cret >= 0 {
			*nameUTF8 = C.GoString(cbuf)
		}
		return int(cret)
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetSpeakerVolume(volume uint) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setspeakervolume(h.libsdkptr, C.uint(volume)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) GetSpeakerVolume() int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_getspeakervolume(h.libsdkptr))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetMicVolume(volume uint) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setmicvolume(h.libsdkptr, C.uint(volume)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) GetMicVolume() int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_getmicvolume(h.libsdkptr))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) AudioPlayLoopbackTest(enable bool) {
	if h.libsdkptr != nil {
		C.portsip_c_audioplayloopbacktest(h.libsdkptr, C.bool(enable))
	}
}

func (h *SDKHandle) GetNumOfVideoCaptureDevices() int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_getnumofvideocapturedevices(h.libsdkptr))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) GetVideoCaptureDeviceName(index uint, uniqueIdUTF8 *string, deviceNameUTF8 *string) int {
	if h.libsdkptr != nil {
		uniquebuf := make([]byte, 1024)
		cuniquebuf := (*C.char)(unsafe.Pointer(&uniquebuf[0]))
		devicebuf := make([]byte, 1024)
		cdevicebuf := (*C.char)(unsafe.Pointer(&devicebuf[0]))

		cret := C.portsip_c_getvideocapturedevicename(h.libsdkptr, C.uint(index), cuniquebuf, C.uint(len(uniquebuf)), cdevicebuf, C.uint(len(devicebuf)))
		if cret >= 0 {
			*uniqueIdUTF8 = C.GoString(cuniquebuf)
			*deviceNameUTF8 = C.GoString(cdevicebuf)
		}
		return int(cret)
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) ShowVideoCaptureSettingsDialogbox(uniqueIdUTF8 string,
	uniqueIdUTF8Length uint,
	dialogTitle string,
	parentWindow unsafe.Pointer,
	x uint,
	y uint) int {
	if h.libsdkptr != nil {
		cuniqueIdUTF8 := C.CString(uniqueIdUTF8)
		defer C.free(unsafe.Pointer(cuniqueIdUTF8))
		cdialogTitle := C.CString(dialogTitle)
		defer C.free(unsafe.Pointer(cdialogTitle))

		return int(C.portsip_c_showvideocapturesettingsdialogbox(h.libsdkptr,
			cuniqueIdUTF8,
			C.uint(uniqueIdUTF8Length),
			cdialogTitle,
			parentWindow,
			C.uint(x),
			C.uint(y)))
	}
	return ECoreNotInitialized
}
