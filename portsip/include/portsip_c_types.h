#ifndef PORTSIP_C_TYPES_H
#define PORTSIP_C_TYPES_H


#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#if defined(CSDK_SHARED)
#if defined(_MSC_VER)
#ifdef CSDK_EXPORTS
#define CSDK_EXPORT __declspec(dllexport)
#else
#define CSDK_EXPORT __declspec(dllimport)
#endif
#else // defined(_MSC_VER)
#ifdef CSDK_EXPORTS
#define CSDK_EXPORT __attribute__((visibility("default")))
#else
#define CSDK_EXPORT
#endif
#endif
#else // defined(CSDK_SHARED)
#define CSDK_EXPORT
#endif


#ifndef ET_BUFFER_SIZE
#define ET_BUFFER_SIZE 512
#endif

#ifndef SIP_MESSAGE_MAXSIZE
#define SIP_MESSAGE_MAXSIZE 1024 * 64
#endif

typedef enum
{
    SIP_UNKNOWN = 0,
    SIP_REGISTER_SUCCESS = 1, // Register to SIP server succeeded
    SIP_REGISTER_FAILURE,     // Register to SIP server failed

    SIP_INVITE_INCOMING,         //	The call is incoming
    SIP_INVITE_TRYING,           //	The call is trying
    SIP_INVITE_SESSION_PROGRESS, //	The 183 MESSAGE, early media
    SIP_INVITE_RINGING,          //	The call is ringing
    SIP_INVITE_ANSWERED,         //	The callee has answered this call
    SIP_INVITE_FAILURE,          //	The call is failed
    SIP_INVITE_UPDATED,          //	The remote party updated this call
    SIP_INVITE_CONNECTED,        //	The call(dialog) is connected - usually for received or sent the ACK
    SIP_INVITE_BEGINING_FORWARD, //  When the SDK begin forwarding the call
    SIP_INVITE_CLOSED,           //	The call is closed

    SIP_DIALOG_STATE_UPDATED, //  When a subscribed user in a call and the
                              //  call state is changed, this event will be triggered

    SIP_REMOTE_HOLD,   //	The remote party has hold this call
    SIP_REMOTE_UNHOLD, //	The remote party has take off the hold

    SIP_RECEIVED_REFER,        //	Received a REFER message
    SIP_REFER_ACCEPTED,        //	The REFER request was accepted
    SIP_REFER_REJECTED,        //	The REFER request was rejected
    SIP_TRANSFER_TRYING,       //	The call transfer is trying
    SIP_TRANSFER_RINGING,      //	The transfer call is ringing
    SIP_ACTV_TRANSFER_SUCCESS, //	The active transfer call succeeds
    SIP_ACTV_TRANSFER_FAILURE, //	Active transfer call is failure

    SIP_RECEIVED_SIGNALING, //	This event will be fired when the SDK received every SIP message
    SIP_SENDING_SIGNALING,  //	This event will be fired when the SDK send every SIP message

    SIP_WAITING_VOICEMESSAGE, //	If have waiting voice message, this event will be fired(MWI)
    SIP_WAITING_FAXMESSAGE,   //	If have waiting fax message, this event will be fired(MWI)

    SIP_RECV_DTMFTONE, //	This event will be fired when received a DTMF tone

    SIP_RECV_MESSAGE,             //	Received a MESSAGE message in dialog
    SIP_RECV_OUTOFDIALOG_MESSAGE, //	Received a MESSAGE out of dialog

    SIP_SEND_MESSAGE_SUCCESS, //	Succeeded to send the message
    SIP_SEND_MESSAGE_FAILURE, //	Failed to send the message

    SIP_SEND_OUTOFDIALOG_MESSAGE_SUCCESS, //	Succeeded to send the out of dialog message
    SIP_SEND_OUTOFDIALOG_MESSAGE_FAILURE, //	Failed to send the out of dialog message

    SIP_PRESENCE_RECV_SUBSCRIBE, //	The remote side request subscribe presence state
    SIP_PRESENCE_ONLINE,         //	The contact is go online
    SIP_PRESENCE_OFFLINE,        //	The contact is go offline

    SIP_RECV_OPTIONS, //	Received an OPTIONS message out of dialog
    SIP_RECV_INFO,    //	Received a INFO message in dialog

    SIP_RECV_NOTIFY_OF_SUBSCRIPTION, // Received a NOTIFY of the SUBSCRIPTION
    SIP_SUBSCRIPTION_FAILURE,        // Failed to SUBSCRIBE an event
    SIP_SUBSCRIPTION_TERMINATED,     // The SUBSCRIPTION is terminated

    SIP_PLAY_FILE_FINISHED, // startPlayingFileToRemote or startPlayingFileLocally finished.
    SIP_CALL_STATISTICS,    // getStatistics finished.

} SIP_EVENT;

/// Media type
typedef enum
{
	MEDIA_TYPE_UNKONW = -1,
	MEDIA_AUDIO = 0,
	MEDIA_VIDEO = 1,
	MEDIA_SCREEN = 2
}CHANNEL_MEDIA_TYPE;

/// Audio codec type
typedef enum
{
	AUDIOCODEC_NONE = -1,
	//Static payload type
	AUDIOCODEC_PCMU = 0,	///< PCMU/G711 U-law 8KHZ 64kbit/s
	AUDIOCODEC_GSM = 3,	///< GSM 8KHZ 13kbit/s
	AUDIOCODEC_PCMA = 8,	///< PCMA/G711 A-law 8KHZ 64kbit/s
	AUDIOCODEC_G722 = 9,	///< G722 16KHZ 64kbit/s
	AUDIOCODEC_G729 = 18,	///< G729 8KHZ 8kbit/s

	//Dynamic payload type
	AUDIOCODEC_ILBC = 97,	///< iLBC 8KHZ 30ms-13kbit/s 20 ms-15kbit/s
	AUDIOCODEC_AMR = 98,	///< Adaptive Multi-Rate (AMR) 8KHZ (4.75,5.15,5.90,6.70,7.40,7.95,10.20,12.20)kbit/s
	AUDIOCODEC_AMRWB = 99,	///< Adaptive Multi-Rate Wideband (AMR-WB)16KHZ (6.60,8.85,12.65,14.25,15.85,18.25,19.85,23.05,23.85)kbit/s
	AUDIOCODEC_SPEEX = 100,	///< SPEEX 8KHZ (2-24)kbit/s
	AUDIOCODEC_SPEEXWB = 102,	///< SPEEX 16KHZ (4-42)kbit/s
	AUDIOCODEC_ISACWB = 103,	///< Internet Speech Audio Codec(iSAC) 16KHZ (32-54)kbit/s
	AUDIOCODEC_ISACSWB = 104,	///< Internet Speech Audio Codec(iSAC) 16KHZ (32-160)kbit/s
	AUDIOCODEC_G7221 = 121,	///< G722.1 16KHZ (16,24,32)kbit/s
	AUDIOCODEC_OPUS = 105,	///< OPUS 48KHZ (5-510)kbit/s 32kbit/s
	AUDIOCODEC_DTMF = 101	///< DTMF RFC 2833
}AUDIOCODEC_TYPE;

/// Video codec type
typedef enum
{
	VIDEO_CODEC_NONE = -1,	///< Do not use video codec.
	VIDEO_CODEC_I420 = 113,	///< I420/YUV420 Raw Video format. It can be used with startRecord only.
	VIDEO_CODEC_H263 = 34,	///< H263 video codec
	VIDEO_CODEC_H263_1998 = 115,	///< H263+/H263 1998 video codec
	VIDEO_CODEC_H264 = 125,	///< H264 video codec
	VIDEO_CODEC_VP8 = 120,	///< VP8 video codec
	VIDEO_CODEC_VP9 = 122	///< VP9 video codec
}VIDEOCODEC_TYPE;


/// The record file format
typedef enum
{
	FILEFORMAT_NONE = 0,	///<	Not Recorded.

	//Audio file format
	FILEFORMAT_WAVE = 1,	///<	The record audio file is in WAVE format.
	FILEFORMAT_AMR,			///<	The record audio file is in AMR format - all voice data are compressed by AMR codec. Mono
	FILEFORMAT_MP3,			///<	The record audio file is in MP3 format.

	//video file format
	FILEFORMAT_MP4,			///<	The record video file is in MP4(AAC and H264) format.
}FILE_FORMAT;

typedef enum
{
	DIRECTION_NONE = 0,		///<	NOT EXIST.
	DIRECTION_SEND_RECV = 1,///<	both received and sent.
	DIRECTION_SEND,			///<	Only the sent.
	DIRECTION_RECV,			///<	Only the received .
	DIRECTION_INACTIVE,	///<	INACTIVE.
}DIRECTION_MODE;

///The audio/Video record mode
typedef enum
{
	RECORD_NONE = 0,		///<	Not Recorded.
	RECORD_RECV = 1,		///<	Only record the received data.
	RECORD_SEND,			///<	Only record the sent data.
	RECORD_BOTH				///<	Record both received and sent data.
}RECORD_MODE;

/// Log Type
typedef enum
{
	LOG_TYPE_NONE = 0,
	LOG_TYPE_FILE = 1,
	LOG_TYPE_CALLBACK = 2,
	LOG_TYPE_COUT = 3,		///< DEBUG To cout (xcode output window, Visual studio Debug Window)
}PORTSIP_LOG_TYPE;

/// Log level
typedef enum
{
	PORTSIP_LOG_NONE = -1,
	PORTSIP_LOG_ERROR = 1,
	PORTSIP_LOG_WARNING = 2,
	PORTSIP_LOG_INFO = 3,
	PORTSIP_LOG_DEBUG = 4
}PORTSIP_LOG_LEVEL;

/// SRTP Policy
typedef enum
{
	SRTP_POLICY_NONE = 0,	///< Do not use SRTP. The SDK can receive both encrypted calls (SRTP) and unencrypted calls, but cannot place encrypted outgoing calls.
	SRTP_POLICY_FORCE,		///< All calls must use SRTP. The SDK allows to receive encrypted calls and place outgoing encrypted calls only.
	SRTP_POLICY_PREFER		///< Top priority for using SRTP. The SDK allows to receive encrypted and decrypted calls, and place outgoing encrypted calls and unencrypted calls.
}SRTP_POLICY;

/// Transport for SIP signaling.
typedef enum
{
	TRANSPORT_NONE = -1,	///< Undefined Transport
	TRANSPORT_UDP = 0,	///< UDP Transport
	TRANSPORT_TLS,		///< TLS Transport
	TRANSPORT_TCP,		///< TCP Transport
}TRANSPORT_TYPE;

///The session refreshment mode
typedef enum
{
	SESSION_REFERESH_UAC = 0,	///< The session refreshment by Caller
	SESSION_REFERESH_UAS,		///< The session refreshment by Callee
	SESSION_REFERESH_LOCAL,		///< The session refreshment by Local
	SESSION_REFERESH_REMOTE		///< The session refreshment by Remote
}SESSION_REFRESH_MODE;

/// type of Presence mode
typedef enum
{
	PRESENCE_MODE_P2P = 0,			// P2P mode
	PRESENCE_MODE_AGENT,            // Presence Agent mode
}PRESENCE_MODES;

///send DTMF tone with two methods
typedef enum
{
	DTMF_RFC2833 = 0,	///<	Send DTMF tone with RFC 2833. Recommended.
	DTMF_INFO = 1	///<	Send DTMF tone with SIP INFO.
}DTMF_METHOD;

typedef int (*fnLogCallback)(uintptr_t obj,
	int level,
	const char* subsystem,
	const char* appName,
	const char* file,
	int line,
	const char* message,
	const char* messageWithHeaders);

// Audio and video callback function prototype, for Visual C++ only
typedef int (*fnAudioRawCallback)(uintptr_t obj,
	long sessionId,
	DIRECTION_MODE type,
	unsigned char* data,
	int dataLength,
	int samplingFreqHz);

typedef int (*fnVideoRawCallback)(uintptr_t obj,
	long sessionId,
	DIRECTION_MODE type,
	int width,
	int height,
	unsigned char* data,
	int dataLength);

typedef int (*fnRecordAudioRawCallback)(uintptr_t obj,
	unsigned char* data,
	int dataLength,
	int samplingFreqHz);

///  Callback functions for received and sent RTP packets. Visual C++ only
///  mediaType 0 - audio, 1 - video, 2 - screen
typedef int (*fnRTPPacketCallback)(uintptr_t obj, long sessionId, int mediaType, DIRECTION_MODE direction, unsigned char* RTPPacket, int packetSize);

#endif
