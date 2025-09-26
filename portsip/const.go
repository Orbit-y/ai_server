package portsip

const (
	LogTypeNone     = 0
	LogTypeFile     = 1
	LogTypeCallBack = 2
	LogTypeCout     = 3
)

const (
	LogLevelNone  = -1
	LogLevelError = 1
	LogLevelWarn  = 2
	LogLevelInfo  = 3
	LogLevelDebug = 4
)

const (
	AUDIOCODEC_NONE = -1
	//Static payload type
	AUDIOCODEC_PCMU = 0  ///< PCMU/G711 U-law 8KHZ 64kbit/s
	AUDIOCODEC_GSM  = 3  ///< GSM 8KHZ 13kbit/s
	AUDIOCODEC_PCMA = 8  ///< PCMA/G711 A-law 8KHZ 64kbit/s
	AUDIOCODEC_G722 = 9  ///< G722 16KHZ 64kbit/s
	AUDIOCODEC_G729 = 18 ///< G729 8KHZ 8kbit/s

	//Dynamic payload type
	AUDIOCODEC_ILBC    = 97  ///< iLBC 8KHZ 30ms-13kbit/s 20 ms-15kbit/s
	AUDIOCODEC_AMR     = 98  ///< Adaptive Multi-Rate (AMR) 8KHZ (4.75,5.15,5.90,6.70,7.40,7.95,10.20,12.20)kbit/s
	AUDIOCODEC_AMRWB   = 99  ///< Adaptive Multi-Rate Wideband (AMR-WB)16KHZ (6.60,8.85,12.65,14.25,15.85,18.25,19.85,23.05,23.85)kbit/s
	AUDIOCODEC_SPEEX   = 100 ///< SPEEX 8KHZ (2-24)kbit/s
	AUDIOCODEC_SPEEXWB = 102 ///< SPEEX 16KHZ (4-42)kbit/s
	AUDIOCODEC_ISACWB  = 103 ///< Internet Speech Audio Codec(iSAC) 16KHZ (32-54)kbit/s
	AUDIOCODEC_ISACSWB = 104 ///< Internet Speech Audio Codec(iSAC) 16KHZ (32-160)kbit/s
	AUDIOCODEC_G7221   = 121 ///< G722.1 16KHZ (16,24,32)kbit/s
	AUDIOCODEC_OPUS    = 105 ///< OPUS 48KHZ (5-510)kbit/s 32kbit/s
	AUDIOCODEC_DTMF    = 101 ///< DTMF RFC 2833
)

const (
	VIDEO_CODEC_NONE      = -1  ///< Do not use video codec.
	VIDEO_CODEC_I420      = 113 ///< I420/YUV420 Raw Video format. It can be used with startRecord only.
	VIDEO_CODEC_H263      = 34  ///< H263 video codec
	VIDEO_CODEC_H263_1998 = 115 ///< H263+/H263 1998 video codec
	VIDEO_CODEC_H264      = 125 ///< H264 video codec
	VIDEO_CODEC_VP8       = 120 ///< VP8 video codec
	VIDEO_CODEC_VP9       = 122 ///< VP9 video codec
)

const (
	DIRECTION_NONE      = 0 ///<	NOT EXIST.
	DIRECTION_SEND_RECV = 1 ///<	both received and sent.
	DIRECTION_SEND      = 2 ///<	Only the sent.
	DIRECTION_RECV      = 3 ///<	Only the received .
	DIRECTION_INACTIVE  = 4 ///<	INACTIVE.
)

const (
	SRTP_POLICY_NONE   = 0 ///< Do not use SRTP. The SDK can receive both encrypted calls (SRTP) and unencrypted calls, but cannot place encrypted outgoing calls.
	SRTP_POLICY_FORCE  = 1 ///< All calls must use SRTP. The SDK allows to receive encrypted calls and place outgoing encrypted calls only.
	SRTP_POLICY_PREFER = 2 ///< Top priority for using SRTP. The SDK allows to receive encrypted and decrypted calls, and place outgoing encrypted calls and unencrypted calls.
)

const (
	SESSION_REFERESH_UAC    = 0 ///< The session refreshment by Caller
	SESSION_REFERESH_UAS    = 1 ///< The session refreshment by Callee
	SESSION_REFERESH_LOCAL  = 2 ///< The session refreshment by Local
	SESSION_REFERESH_REMOTE = 3 ///< The session refreshment by Remote
)

const (
	MEDIA_TYPE_UNKONW = -1
	MEDIA_AUDIO       = 0
	MEDIA_VIDEO       = 1
	MEDIA_SCREEN      = 2
)

const (
	FILEFORMAT_NONE = 0 ///<	Not Recorded.

	//Audio file format
	FILEFORMAT_WAVE = 1 ///<	The record audio file is in WAVE format.
	FILEFORMAT_AMR  = 2 ///<	The record audio file is in AMR format - all voice data are compressed by AMR codec. Mono
	FILEFORMAT_MP3  = 3 ///<	The record audio file is in MP3 format.

	//video file format
	FILEFORMAT_MP = 4
)

const (
	TRANSPORT_NONE = -1 ///< Undefined Transport
	TRANSPORT_UDP  = 0  ///< UDP Transport
	TRANSPORT_TLS  = 1  ///< TLS Transport
	TRANSPORT_TCP  = 2  ///< TCP Transport
)

const (
	PRESENCE_MODE_P2P   = 0 // P2P mode
	PRESENCE_MODE_AGENT = 1 // Presence Agent mode
)

const (
	DTMF_RFC2833 = 0 ///<	Send DTMF tone with RFC 2833. Recommended.
	DTMF_INFO    = 1 ///<	Send DTMF tone with SIP INFO.
)

const (
	SIP_UNKNOWN                          = 0
	SIP_REGISTER_SUCCESS                 = 1  // Register to SIP server succeeded
	SIP_REGISTER_FAILURE                 = 2  // Register to SIP server failed
	SIP_INVITE_INCOMING                  = 3  //	The call is incoming
	SIP_INVITE_TRYING                    = 4  //	The call is trying
	SIP_INVITE_SESSION_PROGRESS          = 5  //	The 183 MESSAGE, early media
	SIP_INVITE_RINGING                   = 6  //	The call is ringing
	SIP_INVITE_ANSWERED                  = 7  //	The callee has answered this call
	SIP_INVITE_FAILURE                   = 8  //	The call is failed
	SIP_INVITE_UPDATED                   = 9  //	The remote party updated this call
	SIP_INVITE_CONNECTED                 = 10 //	The call(dialog) is connected - usually for received or sent the ACK
	SIP_INVITE_BEGINING_FORWARD          = 11 //  When the SDK begin forwarding the call
	SIP_INVITE_CLOSED                    = 12 //	The call is closed
	SIP_DIALOG_STATE_UPDATED             = 13 //  When a subscribed user in a call and the
	SIP_REMOTE_HOLD                      = 14 //	The remote party has hold this call
	SIP_REMOTE_UNHOLD                    = 15 //	The remote party has take off the hold
	SIP_RECEIVED_REFER                   = 16 //	Received a REFER message
	SIP_REFER_ACCEPTED                   = 17 //	The REFER request was accepted
	SIP_REFER_REJECTED                   = 18 //	The REFER request was rejected
	SIP_TRANSFER_TRYING                  = 19 //	The call transfer is trying
	SIP_TRANSFER_RINGING                 = 20 //	The transfer call is ringing
	SIP_ACTV_TRANSFER_SUCCESS            = 21 //	The active transfer call succeeds
	SIP_ACTV_TRANSFER_FAILURE            = 22 //	Active transfer call is failure
	SIP_RECEIVED_SIGNALING               = 23 //	This event will be fired when the SDK received every SIP message
	SIP_SENDING_SIGNALING                = 24 //	This event will be fired when the SDK send every SIP message
	SIP_WAITING_VOICEMESSAGE             = 25 //	If have waiting voice message, this event will be fired(MWI)
	SIP_WAITING_FAXMESSAGE               = 26 //	If have waiting fax message, this event will be fired(MWI)
	SIP_RECV_DTMFTONE                    = 27 //	This event will be fired when received a DTMF tone
	SIP_RECV_MESSAGE                     = 28 //	Received a MESSAGE message in dialog
	SIP_RECV_OUTOFDIALOG_MESSAGE         = 29 //	Received a MESSAGE out of dialog
	SIP_SEND_MESSAGE_SUCCESS             = 30 //	Succeeded to send the message
	SIP_SEND_MESSAGE_FAILURE             = 31 //	Failed to send the message
	SIP_SEND_OUTOFDIALOG_MESSAGE_SUCCESS = 32 //	Succeeded to send the out of dialog message
	SIP_SEND_OUTOFDIALOG_MESSAGE_FAILURE = 33 //	Failed to send the out of dialog message
	SIP_PRESENCE_RECV_SUBSCRIBE          = 34 //	The remote side request subscribe presence state
	SIP_PRESENCE_ONLINE                  = 35 //	The contact is go online
	SIP_PRESENCE_OFFLINE                 = 36 //	The contact is go offline
	SIP_RECV_OPTIONS                     = 37 //	Received an OPTIONS message out of dialog
	SIP_RECV_INFO                        = 38 //	Received a INFO message in dialog
	SIP_RECV_NOTIFY_OF_SUBSCRIPTION      = 39 // Received a NOTIFY of the SUBSCRIPTION
	SIP_SUBSCRIPTION_FAILURE             = 40 // Failed to SUBSCRIBE an event
	SIP_SUBSCRIPTION_TERMINATED          = 41 // The SUBSCRIPTION is terminated
	SIP_PLAY_FILE_FINISHED               = 42 // startPlayingFileToRemote or startPlayingFileLocally finished.
	SIP_CALL_STATISTICS                  = 43 // getStatistics finished.
)
