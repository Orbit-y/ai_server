package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "goexport.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func (h *SDKHandle) SetAudioSample(ptime int, maxPtime int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setaudiosamples(h.libsdkptr, C.int(ptime), C.int(maxPtime)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetAudioDeviceId(inputDeviceId int, outputDeviceId int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setaudiodeviceid(h.libsdkptr, C.int(inputDeviceId), C.int(outputDeviceId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoDeviceId(deviceId int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideodeviceid(h.libsdkptr, C.int(deviceId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoResolution(width int, height int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideoresolution(h.libsdkptr, C.int(width), C.int(height)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetAudioBitrate(sessionId int64, codecType int, bitrateKbps int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setaudiobitrate(h.libsdkptr, C.long(sessionId), C.AUDIOCODEC_TYPE(codecType), C.int(bitrateKbps)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoBitrate(sessionId int64, bitrateKbps int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideobitrate(h.libsdkptr, C.long(sessionId), C.int(bitrateKbps)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoFramerate(sessionId int64, frameRate int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideoframerate(h.libsdkptr, C.long(sessionId), C.int(frameRate)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SendVideo(sessionId int64, sendState bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_sendvideo(h.libsdkptr, C.long(sessionId), C.bool(sendState)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoOrientation(rotation int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideoorientation(h.libsdkptr, C.int(rotation)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) MuteMicrophone(mute bool) {
	if h.libsdkptr != nil {
		C.portsip_c_mutemicrophone(h.libsdkptr, C.bool(mute))
	}
}

func (h *SDKHandle) MuteSpeaker(mute bool) {
	if h.libsdkptr != nil {
		C.portsip_c_mutespeaker(h.libsdkptr, C.bool(mute))
	}
}

func (h *SDKHandle) SetChannelOutputVolumeScaling(sessionId int64, scaling int) {
	if h.libsdkptr != nil {
		C.portsip_c_setchanneloutputvolumescaling(h.libsdkptr, C.long(sessionId), C.int(scaling))
	}
}

func (h *SDKHandle) SetRemoteVideoWindow(sessionId int64, remoteVideoWindow unsafe.Pointer) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setremotevideowindow(h.libsdkptr, C.long(sessionId), remoteVideoWindow))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) DisplayLocalVideo(state bool, localVideoWindow unsafe.Pointer) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_displaylocalvideo(h.libsdkptr, C.bool(state), localVideoWindow))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoNackstatus(state bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideonackstatus(h.libsdkptr, C.bool(state)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableSendPcmStreamToRemote(sessionId int64, state bool, streamSamplesPerSec int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enablesendpcmstreamtoremote(h.libsdkptr, C.long(sessionId), C.bool(state), C.int(streamSamplesPerSec)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SendPcmStreamToRemote(sessionId int64, data string, dataLength int) int {
	if h.libsdkptr != nil {
		cdata := C.CString(data)
		defer C.free(unsafe.Pointer(cdata))
		return int(C.portsip_c_sendpcmstreamtoremote(h.libsdkptr, C.long(sessionId), (*C.uchar)(unsafe.Pointer(cdata)), C.int(dataLength)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableSendVideoStreamToRemote(sessionId int64, state bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enablesendvideostreamtoremote(h.libsdkptr, C.long(sessionId), C.bool(state)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SendVideoStreamToRemote(sessionId int64, data string, dataLength int, width int, height int) int {
	if h.libsdkptr != nil {
		cdata := C.CString(data)
		defer C.free(unsafe.Pointer(cdata))
		return int(C.portsip_c_sendvideostreamtoremote(h.libsdkptr,
			C.long(sessionId),
			(*C.uchar)(unsafe.Pointer(cdata)),
			C.int(dataLength),
			C.int(width),
			C.int(height)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableAudioStreamCallback(sessionId int64, enable bool, callbackMode int, obj uintptr) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enableaudiostreamcallback(h.libsdkptr,
			C.long(sessionId),
			C.bool(enable),
			C.DIRECTION_MODE(callbackMode),
			C.uintptr_t(obj),
			C.fnAudioRawCallback(C.GofnAudioRawCallback)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableVideoStreamCallback(sessionId int64, callbackMode int, obj uintptr) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enablevideostreamcallback(h.libsdkptr,
			C.long(sessionId),
			C.DIRECTION_MODE(callbackMode),
			C.uintptr_t(obj),
			C.fnVideoRawCallback(C.GofnVideoRawCallback)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SendRtpPacketToRemote(sessionId int64, mediaType int, data string, dataLength int) int {
	if h.libsdkptr != nil {
		cdata := C.CString(data)
		defer C.free(unsafe.Pointer(cdata))
		return int(C.portsip_c_sendrtppackettoremote(h.libsdkptr, C.long(sessionId), C.int(mediaType), (*C.uchar)(unsafe.Pointer(cdata)), C.int(dataLength)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableRtpCallback(sessionId int64, mediaType int, callbackMode int, obj uintptr) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enablertpcallback(h.libsdkptr,
			C.long(sessionId),
			C.int(mediaType),
			C.DIRECTION_MODE(callbackMode),
			C.uintptr_t(obj),
			C.fnRTPPacketCallback(C.GofnRTPPacketCallback)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) RequestKeyFrame(sessionId int64, mediaType int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_requestkeyframe(h.libsdkptr, C.long(sessionId), C.int(mediaType)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) StartRecord(sessionId int64,
	recordFilePath string,
	recordFileName string,
	appendTimeStamp bool,
	channels int,
	recordFileFormat int,
	audioRecordMode int,
	videoRecordMode int) int {
	if h.libsdkptr != nil {
		crecordFilePath := C.CString(recordFilePath)
		defer C.free(unsafe.Pointer(crecordFilePath))
		crecordFileName := C.CString(recordFileName)
		defer C.free(unsafe.Pointer(crecordFileName))

		return int(C.portsip_c_startrecord(h.libsdkptr,
			C.long(sessionId),
			crecordFilePath,
			crecordFileName,
			C.bool(appendTimeStamp),
			C.int(channels),
			C.FILE_FORMAT(recordFileFormat),
			C.RECORD_MODE(audioRecordMode),
			C.RECORD_MODE(videoRecordMode)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) StopRecord(sessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_stoprecord(h.libsdkptr, C.long(sessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) StartPlayingFileToRemote(sessionId int64, fileUrl string, loop bool, playAudio int) int {
	if h.libsdkptr != nil {
		cfileUrl := C.CString(fileUrl)
		defer C.free(unsafe.Pointer(cfileUrl))

		return int(C.portsip_c_startplayingfiletoremote(h.libsdkptr, C.long(sessionId), cfileUrl, C.bool(loop), C.int(playAudio)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) StopPlayingFileToRemote(sessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_stopplayingfiletoremote(h.libsdkptr, C.long(sessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetAudioRtcpBandwidth(sessionId int64, BitsRR int, BitsRS int, KBitsAS int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setaudiortcpbandwidth(h.libsdkptr, C.long(sessionId), C.int(BitsRR), C.int(BitsRS), C.int(KBitsAS)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoRtcpBandwidth(sessionId int64, BitsRR int, BitsRS int, KBitsAS int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideortcpbandwidth(h.libsdkptr, C.long(sessionId), C.int(BitsRR), C.int(BitsRS), C.int(KBitsAS)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) GetStatistics(sessionId int64) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_getstatistics(h.libsdkptr, C.long(sessionId)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableVAD(state bool) {
	if h.libsdkptr != nil {
		C.portsip_c_enablevad(h.libsdkptr, C.bool(state))
	}
}

func (h *SDKHandle) EnableAEC(state bool) {
	if h.libsdkptr != nil {
		C.portsip_c_enableaec(h.libsdkptr, C.bool(state))
	}
}

func (h *SDKHandle) EnableCNG(state bool) {
	if h.libsdkptr != nil {
		C.portsip_c_enablecng(h.libsdkptr, C.bool(state))
	}
}

func (h *SDKHandle) EnableAGC(state bool) {
	if h.libsdkptr != nil {
		C.portsip_c_enableagc(h.libsdkptr, C.bool(state))
	}
}

func (h *SDKHandle) EnableANS(state bool) {
	if h.libsdkptr != nil {
		C.portsip_c_enableans(h.libsdkptr, C.bool(state))
	}
}

func (h *SDKHandle) EnableAudioQos(state bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enableaudioqos(h.libsdkptr, C.bool(state)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) EnableVideoQos(state bool) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_enablevideoqos(h.libsdkptr, C.bool(state)))
	}
	return ECoreNotInitialized
}

func (h *SDKHandle) SetVideoMTU(mtu int) int {
	if h.libsdkptr != nil {
		return int(C.portsip_c_setvideomtu(h.libsdkptr, C.int(mtu)))
	}
	return ECoreNotInitialized
}
