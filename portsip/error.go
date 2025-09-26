package portsip

const (
	INVALID_SESSION_ID    = -1
	CONFERENCE_SESSION_ID = 0x7FFF //Conference Session Id, only use on enableVideoStreamCallback API
	LOCALVIDEO_SESSION_ID = 0x7FFE //Local Video Session Id, only use on enableVideoStreamCallback API

	ECoreAlreadyInitialized             = -60000
	ECoreNotInitialized                 = -60001
	ECoreSDKObjectNull                  = -60002
	ECoreArgumentNull                   = -60003
	ECoreInitializeWinsockFailure       = -60004
	ECoreUserNameAuthNameEmpty          = -60005
	ECoreInitializeStackFailure         = -60006
	ECorePortOutOfRange                 = -60007
	ECoreAddTcpTransportFailure         = -60008
	ECoreAddTlsTransportFailure         = -60009
	ECoreAddUdpTransportFailure         = -60010
	ECoreMiniAudioPortOutOfRange        = -60011
	ECoreMaxAudioPortOutOfRange         = -60012
	ECoreMiniVideoPortOutOfRange        = -60013
	ECoreMaxVideoPortOutOfRange         = -60014
	ECoreMiniAudioPortNotEvenNumber     = -60015
	ECoreMaxAudioPortNotEvenNumber      = -60016
	ECoreMiniVideoPortNotEvenNumber     = -60017
	ECoreMaxVideoPortNotEvenNumber      = -60018
	ECoreAudioVideoPortOverlapped       = -60019
	ECoreAudioVideoPortRangeTooSmall    = -60020
	ECoreAlreadyRegistered              = -60021
	ECoreSIPServerEmpty                 = -60022
	ECoreExpiresValueTooSmall           = -60023
	ECoreCallIdNotFound                 = -60024
	ECoreNotRegistered                  = -60025
	ECoreCalleeEmpty                    = -60026
	ECoreInvalidUri                     = -60027
	ECoreAudioVideoCodecEmpty           = -60028
	ECoreNoFreeDialogSession            = -60029
	ECoreCreateAudioChannelFailed       = -60030
	ECoreSessionTimerValueTooSmall      = -60040
	ECoreAudioHandleNull                = -60041
	ECoreVideoHandleNull                = -60042
	ECoreCallIsClosed                   = -60043
	ECoreCallAlreadyHold                = -60044
	ECoreCallNotEstablished             = -60045
	ECoreCallNotHold                    = -60050
	ECoreSipMessaegEmpty                = -60051
	ECoreSipHeaderNotExist              = -60052
	ECoreSipHeaderValueEmpty            = -60053
	ECoreSipHeaderBadFormed             = -60054
	ECoreBufferTooSmall                 = -60055
	ECoreSipHeaderValueListEmpty        = -60056
	ECoreSipHeaderParserEmpty           = -60057
	ECoreSipHeaderValueListNull         = -60058
	ECoreSipHeaderNameEmpty             = -60059
	ECoreAudioSampleNotmultiple         = -60060 //	The audio sample should be multiple of 10
	ECoreAudioSampleOutOfRange          = -60061 //	The audio sample ranges from 10 to 60
	ECoreInviteSessionNotFound          = -60062
	ECoreStackException                 = -60063
	ECoreMimeTypeUnknown                = -60064
	ECoreDataSizeTooLarge               = -60065
	ECoreSessionNumsOutOfRange          = -60066
	ECoreNotSupportCallbackMode         = -60067
	ECoreNotFoundSubscribeId            = -60068
	ECoreCodecNotSupport                = -60069
	ECoreCodecParameterNotSupport       = -60070
	ECorePayloadOutofRange              = -60071 //  Dynamic Payload ranges from 96 to 127
	ECorePayloadHasExist                = -60072 //  Duplicate Payload values are not allowed.
	ECoreFixPayloadCantChange           = -60073
	ECoreCodecTypeInvalid               = -60074
	ECoreCodecWasExist                  = -60075
	ECorePayloadTypeInvalid             = -60076
	ECoreArgumentTooLong                = -60077
	ECoreMiniRtpPortMustIsEvenNum       = -60078
	ECoreCallInHold                     = -60079
	ECoreNotIncomingCall                = -60080
	ECoreCreateMediaEngineFailure       = -60081
	ECoreAudioCodecEmptyButAudioEnabled = -60082
	ECoreVideoCodecEmptyButVideoEnabled = -60083
	ECoreNetworkInterfaceUnavailable    = -60084
	ECoreWrongDTMFTone                  = -60085
	ECoreWrongLicenseKey                = -60086
	ECoreTrialVersionLicenseKey         = -60087
	ECoreOutgoingAudioMuted             = -60088
	ECoreOutgoingVideoMuted             = -60089
	ECoreFailedCreateSdp                = -60090
	ECoreTrialVersionExpired            = -60091
	ECoreStackFailure                   = -60092
	ECoreTransportExists                = -60093
	ECoreUnsupportTransport             = -60094
	ECoreAllowOnlyOneUser               = -60095
	ECoreUserNotFound                   = -60096
	ECoreTransportsIncorrect            = -60097
	ECoreCreateTransportFailure         = -60098
	ECoreTransportNotSet                = -60099
	ECoreECreateSignalingFailure        = -60100
	ECoreArgumentIncorrect              = -60101
	ECoreSipMethodNameEmpty             = -60102
	ECoreSipAlreadySubscribed           = -60103
	ECoreStartRecordFailure             = -60104
	ECoreParsedSdpFailure               = -60105
	EOpenPlayFileFailure                = -60106
	EFilePlayerNotExist                 = -60107

	// IVR
	ECoreIVRObjectNull      = -61001
	ECoreIVRIndexOutOfRange = -61002
	ECoreIVRReferFailure    = -61003
	ECoreIVRWaitingTimeOut  = -61004

	// Conference
	EConferenceAlreadyExists          = -62001
	EConferenceNotExist               = -62002
	EConferenceCreateAudioConfFailure = -62003
	EConferenceCreateVideoConfFailure = -62004
	EConferenceUnsupportedLayout      = -62005

	// Audio

	EAudioFileNameEmpty         = -70000
	EAudioChannelNotFound       = -70001
	EAudioPlayFileAlreadyEnable = -70006

	EAudioPlaySteamNotEnabled          = -70008
	EAudioRegisterCallbackFailure      = -70009
	EAudioCreateAudioConferenceFailure = -70010
	EAudioPlayFileModeNotSupport       = -70012
	EAudioPlayFileFormatNotSupport     = -70013
	EAudioPlaySteamAlreadyEnabled      = -70014
	EAudioCodecNotSupport              = -70016
	EAudioPlayFileNotEnabled           = -70017
	EAudioPlayFileGetPositionFailure   = -70018
	EAudioVolumeOutOfRange             = -70020
	EAudioNotSupportDTMF2833           = -70021

	// Video
	EVideoFileNameEmpty              = -80000
	EVideoGetDeviceNameFailure       = -80001
	EVideoGetDeviceIdFailure         = -80002
	EVideoStartCaptureFailure        = -80003
	EVideoChannelNotFound            = -80004
	EVideoStartSendFailure           = -80005
	EVideoGetStatisticsFailure       = -80006
	EVideoStartPlayAviFailure        = -80007
	EVideoSendAviFileFailure         = -80008
	EVideoRecordUnknowCodec          = -80009
	EVideoCantSetDeviceIdDuringCall  = -80010
	EVideoUnsupportCaptureRotate     = -80011
	EVideoUnsupportCaptureResolution = -80012
	ECameraSwitchTooOften            = -80013
	EMTUOutOfRange                   = -80014
	EVideoCodecNotSupport            = -80015
	EVideoSendStreamAlreadyExists    = -80016

	// Device
	EDeviceObjectNull           = -90000
	EDeviceGetDeviceNameFailure = -90001

	//Screen
	EScreenCapturerNotSupported   = -90002
	EScreenSourceIdNotFound       = -90003
	EScreenChannelNotFound        = -90004
	EScreenCapturerNotInitialized = -90005
	EScreenCapturerRuning         = -90006
)
