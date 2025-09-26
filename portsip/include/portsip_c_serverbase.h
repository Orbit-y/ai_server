#ifndef LIB_PORTSIP_C_SERVERBASE_H
#define LIB_PORTSIP_C_SERVERBASE_H

#include "portsip_c_types.h"

#ifdef __cplusplus
extern "C" {
#endif

CSDK_EXPORT void portsip_c_delcallbackparameters(void* parameters);
CSDK_EXPORT const char* portsip_c_getversion();

CSDK_EXPORT void* portsip_c_initialize(void* callbackDispatcher,
                         bool singleAccountMode,
                         const char* transports,
                         int maxCallSessions,
                         const char* sipAgentString,
                         const char* localIp,
                         int audioDeviceLayer,
                         int videoDeviceLayer,
                         const char* TLSCertificatesRootPath,
                         const char* TLSCipherList,
                         bool verifyTLSCertificate,
                         void* encoderFactory,
                         void* encoderFactoryCallback,
                         void* decoderFactoryCallback,
                         int* errorCode);

CSDK_EXPORT void portsip_c_uninitialize(void* libSDK);

CSDK_EXPORT void portsip_c_discardreceivedpacket(void* libSDK, bool state);

CSDK_EXPORT void portsip_c_setlog(void* libSDK,
                               PORTSIP_LOG_TYPE logType,
                               PORTSIP_LOG_LEVEL logLevel,
                               const char* APPName,
                               const char* logFilePath,
                               size_t maxFileSizeMB,
                               uintptr_t logObj,
                               fnLogCallback logCallback);

CSDK_EXPORT int portsip_c_setlicensekey(void* libSDK, const char* key);
CSDK_EXPORT int portsip_c_getnicnums(void* libSDK);
CSDK_EXPORT int portsip_c_getlocalipaddress(void* libSDK, int index, char* ip, int ipSize);
CSDK_EXPORT int portsip_c_updatelocalip(void* libSDK, const char* ip);
CSDK_EXPORT int portsip_c_adduser(void* libSDK,
                               const char* userName,
                               const char* displayName,
                               const char* authName,
                               const char* password,
                               const char* useTransportName,
                               const char* sipDomain,
                               const char* sipServerAddr,
                               int sipServerPort,
                               const char* stunServerAddr,
                               int stunServerPort,
                               const char* outboundServerAddr,
                               int outboundServerPort);

CSDK_EXPORT int portsip_c_updateuser(void* libSDK,
                                  int userId,
                                  const char* userName,
                                  const char* displayName,
                                  const char* authName,
                                  const char* password,
                                  const char* useTransportName,
                                  const char* sipDomain,
                                  const char* sipServerAddr,
                                  int sipServerPort,
                                  const char* stunServerAddr,
                                  int stunServerPort,
                                  const char* outboundServerAddr,
                                  int outboundServerPort);

CSDK_EXPORT void portsip_c_removeuser(void* libSDK, int userId);
CSDK_EXPORT int portsip_c_setdisplayname(void* libSDK, int userId, const char* displayName);

CSDK_EXPORT int portsip_c_setinstanceid(void* libSDK, const char* uuid);
CSDK_EXPORT int portsip_c_setuseragent(void* libSDK, const char* sipAgentString);

CSDK_EXPORT int portsip_c_registerserver(void* libSDK, int userId, int regExpires, int retryTimes);
CSDK_EXPORT int portsip_c_unregisterserver(void* libSDK, int userId, int waitMS);

CSDK_EXPORT int portsip_c_setreliableprovisional(void* libSDK, int mode);
CSDK_EXPORT int portsip_c_enable3gpptags(void* libSDK, int userId, bool enable);
CSDK_EXPORT void portsip_c_enablecallbacksignaling(void* libSDK, bool enableSending, bool enableReceived);

CSDK_EXPORT int portsip_c_addaudiocodec(void* libSDK, AUDIOCODEC_TYPE codecType);
CSDK_EXPORT int portsip_c_addvideocodec(void* libSDK, VIDEOCODEC_TYPE codecType);
CSDK_EXPORT bool portsip_c_isaudiocodecempty(void* libSDK);
CSDK_EXPORT bool portsip_c_isvideocodecempty(void* libSDK);
CSDK_EXPORT int portsip_c_setaudiocodecpayloadtype(void* libSDK, AUDIOCODEC_TYPE codecType, int payloadType);
CSDK_EXPORT int portsip_c_setvideocodecpayloadtype(void* libSDK, VIDEOCODEC_TYPE codecType, int payloadType);
CSDK_EXPORT void portsip_c_clearaudiocodec(void* libSDK);
CSDK_EXPORT void portsip_c_clearvideocodec(void* libSDK);
CSDK_EXPORT int portsip_c_setaudiocodecparameter(void* libSDK, AUDIOCODEC_TYPE codecType, const char* parameter);
CSDK_EXPORT int portsip_c_setvideocodecparameter(void* libSDK, VIDEOCODEC_TYPE codecType, const char* parameter);

CSDK_EXPORT int portsip_c_setsrtppolicy(void* libSDK, SRTP_POLICY policy, bool allowSrtpOverUnsecureTransport);
CSDK_EXPORT int portsip_c_setrtpportrange(void* libSDK, int minimumRtpPort, int maximumRtpPort);

CSDK_EXPORT int portsip_c_enablecallforward(void* libSDK, int userId, bool forBusyOnly, const char* forwardTo);
CSDK_EXPORT int portsip_c_disablecallforward(void* libSDK, int userId);
CSDK_EXPORT int portsip_c_enablesessiontimer(void* libSDK, int timerSeconds, SESSION_REFRESH_MODE refreshMode);
CSDK_EXPORT int portsip_c_disablesessiontimer(void* libSDK);
CSDK_EXPORT void portsip_c_setdonotdisturb(void* libSDK, int userId, bool state);
CSDK_EXPORT int portsip_c_setrtpkeepalive(void* libSDK, bool state, int keepAlivePayloadType, int deltaTransmitTimeMS);
CSDK_EXPORT int portsip_c_setkeepalivetime(void* libSDK, int keepAliveTime);

CSDK_EXPORT int portsip_c_getsipmessageheadervalue(void* libSDK, const char* sipMessage, const char* headerName, char* headerValue, int headerValueLength);
CSDK_EXPORT long portsip_c_addsipmessageheader(void* libSDK, long sessionId, const char* methodName, int msgType, const char* headerName, const char* headerValue);
CSDK_EXPORT int portsip_c_removeaddedsipmessageheader(void* libSDK, long addedSipMessageId);
CSDK_EXPORT void portsip_c_clearaddedsipmessageheaders(void* libSDK);
CSDK_EXPORT long portsip_c_modifysipmessageheader(void* libSDK, long sessionId, const char* methodName, int msgType, const char* headerName, const char* headerValue);
CSDK_EXPORT int portsip_c_removemodifiedsipmessageheader(void* libSDK, long modifiedSipMessageId);
CSDK_EXPORT void portsip_c_clearmodifiedsipmessageheaders(void* libSDK);
CSDK_EXPORT int portsip_c_addsupportedmimetype(void* libSDK, const char* methodName, const char* mimeType, const char* subMimeType);

CSDK_EXPORT int portsip_c_setaudiosamples(void* libSDK, int ptime, int maxPtime);
CSDK_EXPORT int portsip_c_setaudiodeviceid(void* libSDK, int inputDeviceId, int outputDeviceId);
CSDK_EXPORT int portsip_c_setvideodeviceid(void* libSDK, int deviceId);
CSDK_EXPORT int portsip_c_setvideoresolution(void* libSDK, int width, int height);
CSDK_EXPORT int portsip_c_setaudiobitrate(void* libSDK, long sessionId, AUDIOCODEC_TYPE codecType, int bitrateKbps);
CSDK_EXPORT int portsip_c_setvideobitrate(void* libSDK, long sessionId, int bitrateKbps);
CSDK_EXPORT int portsip_c_setvideoframerate(void* libSDK, long sessionId, int frameRate);
CSDK_EXPORT int portsip_c_sendvideo(void* libSDK, long sessionId, bool sendState);

CSDK_EXPORT int portsip_c_setvideoorientation(void* libSDK, int rotation);
CSDK_EXPORT void portsip_c_mutemicrophone(void* libSDK, bool mute);
CSDK_EXPORT void portsip_c_mutespeaker(void* libSDK, bool mute);
CSDK_EXPORT void portsip_c_setchanneloutputvolumescaling(void* libSDK, long sessionId, int scaling);
CSDK_EXPORT int portsip_c_setremotevideowindow(void* libSDK, long sessionId, void* remoteVideoWindow);
CSDK_EXPORT int portsip_c_displaylocalvideo(void* libSDK, bool state, void* localVideoWindow);
CSDK_EXPORT int portsip_c_setvideonackstatus(void* libSDK, bool state);

CSDK_EXPORT long portsip_c_call(void* libSDK, int userId, const char* callee, bool sendSdp, bool videoCall, const char* headerNames, const char* headerValues);
CSDK_EXPORT int portsip_c_rejectcall(void* libSDK, long sessionId, int code);
CSDK_EXPORT int portsip_c_hangup(void* libSDK, long sessionId);
CSDK_EXPORT int portsip_c_answercall(void* libSDK, long sessionId, bool videoCall);
CSDK_EXPORT int portsip_c_updatecall(void* libSDK, long sessionId, bool enableAudio, bool enableVideo, bool enableScreen);
CSDK_EXPORT int portsip_c_hold(void* libSDK, long sessionId);
CSDK_EXPORT int portsip_c_unhold(void* libSDK, long sessionId);
CSDK_EXPORT int portsip_c_refer(void* libSDK, long sessionId, const char* referTo);
CSDK_EXPORT int portsip_c_refer2(void* libSDK, long sessionId, const char* referTo, const char* headerName, const char* headerValue);
CSDK_EXPORT int portsip_c_attendedrefer(void* libSDK, long sessionId, long replaceSessionId, const char* referTo);
CSDK_EXPORT long portsip_c_acceptrefer(void* libSDK, int userId, long referId, const char* referSignaling);
CSDK_EXPORT int portsip_c_rejectrefer(void* libSDK, int userId, long referId);
CSDK_EXPORT int portsip_c_mutesession(void* libSDK,
                                   long sessionId,
                                   bool muteIncomingAudio,
                                   bool muteOutgoingAudio,
                                   bool muteIncomingVideo,
                                   bool muteOutgoingVideo);
CSDK_EXPORT int portsip_c_holdsession(void* libSDK,
                                   long sessionId,
                                   bool holdIncomingAudio,
                                   bool holdOutgoingAudio,
                                   bool holdIncomingVideo,
                                   bool holdOutgoingVideo);

CSDK_EXPORT int portsip_c_redirect(void* libSDK, long sessionId, long toSessionId);
CSDK_EXPORT int portsip_c_forwardcall(void* libSDK, long sessionId, const char* forwardTo);
CSDK_EXPORT int portsip_c_senddtmf(void* libSDK, long sessionId, DTMF_METHOD dtmfMethod, int code, int dtmfDuration, bool playDtmfTone);

CSDK_EXPORT int portsip_c_enablesendpcmstreamtoremote(void* libSDK, long sessionId, bool state, int streamSamplesPerSec);
CSDK_EXPORT int portsip_c_sendpcmstreamtoremote(void* libSDK, long sessionId, const unsigned char* data, int dataLength);

CSDK_EXPORT int portsip_c_enablesendvideostreamtoremote(void* libSDK, long sessionId, bool state);
CSDK_EXPORT int portsip_c_sendvideostreamtoremote(void* libSDK, long sessionId, const unsigned char* data, int dataLength, int width, int height);

CSDK_EXPORT int portsip_c_enableaudiostreamcallback(void* libSDK, long sessionId, bool enable, DIRECTION_MODE callbackMode, uintptr_t obj, fnAudioRawCallback callbackFunc);
CSDK_EXPORT int portsip_c_enablevideostreamcallback(void* libSDK, long sessionId, DIRECTION_MODE callbackMode, uintptr_t obj, fnVideoRawCallback callbackFunc);
CSDK_EXPORT int portsip_c_sendrtppackettoremote(void* libSDK, long sessionId, int mediaType, const unsigned char* data, int dataLength);

CSDK_EXPORT int portsip_c_enablertpcallback(void* libSDK, long sessionId, int mediaType, DIRECTION_MODE mode, uintptr_t obj, fnRTPPacketCallback callback);

CSDK_EXPORT int portsip_c_requestkeyframe(void* libSDK, long sessionId, int mediaType);

CSDK_EXPORT int portsip_c_startrecord(void* libSDK,
                                   long sessionId,
                                   const char* recordFilePath,
                                   const char* recordFileName,
                                   bool appendTimeStamp,
                                   int channels,
                                   FILE_FORMAT recordFileFormat,
                                   RECORD_MODE audioRecordMode,
                                   RECORD_MODE videoRecordMode);

CSDK_EXPORT int portsip_c_stoprecord(void* libSDK, long sessionId);

CSDK_EXPORT int portsip_c_startplayingfiletoremote(void* libSDK,
                                                long sessionId,
                                                const char* fileUrl,
                                                bool loop,
                                                int playAudio);

CSDK_EXPORT int portsip_c_stopplayingfiletoremote(void* libSDK, long sessionId);

CSDK_EXPORT int portsip_c_setaudiortcpbandwidth(void* libSDK, long sessionId, int BitsRR, int BitsRS, int KBitsAS);
CSDK_EXPORT int portsip_c_setvideortcpbandwidth(void* libSDK, long sessionId, int BitsRR, int BitsRS, int KBitsAS);

CSDK_EXPORT int portsip_c_getstatistics(void* libSDK, long sessionId);

CSDK_EXPORT void portsip_c_enablevad(void* libSDK, bool state);
CSDK_EXPORT void portsip_c_enableaec(void* libSDK, bool state);
CSDK_EXPORT void portsip_c_enablecng(void* libSDK, bool state);
CSDK_EXPORT void portsip_c_enableagc(void* libSDK, bool state);
CSDK_EXPORT void portsip_c_enableans(void* libSDK, bool state);

CSDK_EXPORT int portsip_c_enableaudioqos(void* libSDK, bool state);
CSDK_EXPORT int portsip_c_enablevideoqos(void* libSDK, bool state);
CSDK_EXPORT int portsip_c_setvideomtu(void* libSDK, int mtu);

CSDK_EXPORT int portsip_c_sendoptions(void* libSDK, int userId, const char* to, const char* sdp);
CSDK_EXPORT int portsip_c_sendinfo(void* libSDK, long sessionId, const char* mimeType, const char* subMimeType, const char* infoContents);
CSDK_EXPORT long portsip_c_sendsubscription(void* libSDK, int userId, const char* to, const char* eventName);
CSDK_EXPORT void portsip_c_terminatesubscription(void* libSDK, int userId, long subscriptionId);

CSDK_EXPORT long portsip_c_sendmessage(void* libSDK, int userId,
                                    long sessionId,
                                    const char* mimeType,
                                    const char* subMimeType,
                                    const unsigned char* message,
                                    int messageLength);

CSDK_EXPORT long portsip_c_sendoutofdialogmessage(void* libSDK,
                                               int userId,
                                               const char* to,
                                               const char* mimeType,
                                               const char* subMimeType,
                                               bool isSMS,
                                               const unsigned char* message,
                                               int messageLength,
                                               const char* displayName);

CSDK_EXPORT int portsip_c_presencesubscribe(void* libSDK, int userId, const char* contact, const char* subject);
CSDK_EXPORT int portsip_c_presenceacceptsubscribe(void* libSDK, int userId, long subscribeId);
CSDK_EXPORT int portsip_c_presencerejectsubscribe(void* libSDK, int userId, long subscribeId);
CSDK_EXPORT int portsip_c_setpresencestatus(void* libSDK, int userId, long subscribeId, const char* statusText);
CSDK_EXPORT int portsip_c_setdefaultpublicationtime(void* libSDK, int userId, unsigned int secs);

// use for webrtc engine
CSDK_EXPORT long portsip_c_createsipsession(void* libSDK, int userId, long sessionId);
CSDK_EXPORT int portsip_c_destorysipsession(void* libSDK, long sessionId);
CSDK_EXPORT int portsip_c_createsdpoffer(void* libSDK, long sessionId,
                                      DIRECTION_MODE audioDirection,
                                      DIRECTION_MODE videoDirection,
                                      DIRECTION_MODE screenDirection,
                                      char* newSdp, int sdpSize);
CSDK_EXPORT int portsip_c_createsdpanswer(void* libSDK, long sessionId,
                                       DIRECTION_MODE audioDirection,
                                       DIRECTION_MODE videoDirection,
                                       DIRECTION_MODE screenDirection,
                                       char* newSdp, int sdpSize);
CSDK_EXPORT int portsip_c_receivedremotesdp(void* libSDK, long sessionId, const char* remoteSdp, int sdpSize,
                                         DIRECTION_MODE* audioDirection,
                                         DIRECTION_MODE* videoDirection,
                                         DIRECTION_MODE* screenDirection);

CSDK_EXPORT int portsip_c_registerdtmfdetection(void* libSDK, void* observer);
//////////////////////////////////////////////////////////////////////////
//
// Device Manage functions.
//
//////////////////////////////////////////////////////////////////////////

CSDK_EXPORT int portsip_c_getnumofrecordingdevices(void* libSDK);

CSDK_EXPORT int portsip_c_getnumofplayoutdevices(void* libSDK);

CSDK_EXPORT int portsip_c_getrecordingdevicename(void* libSDK,
                                              int index,
                                              char* nameUTF8,
                                              int nameUTF8Length);

CSDK_EXPORT int portsip_c_getplayoutdevicename(void* libSDK,
                                            int index,
                                            char* nameUTF8,
                                            int nameUTF8Length);

CSDK_EXPORT int portsip_c_setspeakervolume(void* libSDK, unsigned int volume);
CSDK_EXPORT int portsip_c_getspeakervolume(void* libSDK);

CSDK_EXPORT int portsip_c_setmicvolume(void* libSDK, unsigned int volume);
CSDK_EXPORT int portsip_c_getmicvolume(void* libSDK);
CSDK_EXPORT void portsip_c_audioplayloopbacktest(void* libSDK, bool enable);

// Video device
CSDK_EXPORT int portsip_c_getnumofvideocapturedevices(void* libSDK);
CSDK_EXPORT int portsip_c_getvideocapturedevicename(void* libSDK,
                                                 unsigned int index,
                                                 char* uniqueIdUTF8,
                                                 const unsigned int uniqueIdUTF8Length,
                                                 char* deviceNameUTF8,
                                                 const unsigned int deviceNameUTF8Length);

CSDK_EXPORT int portsip_c_showvideocapturesettingsdialogbox(void* libSDK,
                                                         const char* uniqueIdUTF8,
                                                         const unsigned int uniqueIdUTF8Length,
                                                         const char* dialogTitle,
                                                         void* parentWindow,
                                                         const unsigned int x,
                                                         const unsigned int y);

#ifdef __cplusplus
}
#endif

#endif // LIB_PORTSIP_C_SERVERBASE_H
