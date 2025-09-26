package main

import (
	"goportsipsdk/portsip"
	"log"
	"unsafe"
)

var SESSION *Session

func MyOnMessageCallbackFunc(params unsafe.Pointer) {

	defer portsip.DestroyParam(params)
	sip_event := portsip.ParamGetEventtype(params)
	userid := portsip.ParamGetUserid(params)
	log.Printf("recived sip messge type : %d  userid : %d", sip_event, userid)
	switch sip_event {
	case portsip.SIP_REGISTER_SUCCESS:
		log.Printf("userid : %d register success", userid)
		onRegisterSuccess(params)
	case portsip.SIP_REGISTER_FAILURE:
		log.Printf("userid : %d register failure, error_code : %d error_msg : %s ",
			userid,
			portsip.ParamGetStatusCode(params),
			portsip.ParamGetStatusText(params))
	case portsip.SIP_INVITE_FAILURE:
		log.Printf("userid : %d invite failure, error_code : %d error_msg : %s ",
			userid,
			portsip.ParamGetStatusCode(params),
			portsip.ParamGetStatusText(params))
	case portsip.SIP_INVITE_INCOMING:
		onIncoming(params)
	case portsip.SIP_INVITE_ANSWERED:
		onAnswered(params)
	case portsip.SIP_INVITE_CONNECTED:
		onSIPCallConnected(params)
	case portsip.SIP_INVITE_CLOSED:
		onSIPCallClosed(params)
	}

	return
}

func onRegisterSuccess(params unsafe.Pointer) {
	userid := portsip.ParamGetUserid(params)
	sessionid := HANDLE.Call(userid, callee, true, false, "Remote-Party-ID", caller+"@"+domain)
	if sessionid < 0 {
		log.Printf("Failed to call %s , ret %d", callee, sessionid)
		return
	}
}

func onAnswered(params unsafe.Pointer) {
	sessionid := portsip.ParamGetSessionid(params)
	session, err := CreateSession(sessionid)
	if err != nil {
		log.Print(err)
		HANDLE.Hangup(sessionid)
		return
	}
	SESSION = session
}

func onIncoming(params unsafe.Pointer) {
	sessionid := portsip.ParamGetSessionid(params)
	session, err := CreateSession(sessionid)
	if err != nil {
		log.Print(err)
		HANDLE.RejectCall(sessionid, 486)
		return
	}

	ret := HANDLE.AnswerCall(sessionid, false)
	if ret != 0 {
		HANDLE.RejectCall(sessionid, 486)
		caller := portsip.ParamGetCaller(params)
		log.Printf("Failed to answer call from %s", caller)
		return
	}
	SESSION = session
}

func onSIPCallConnected(params unsafe.Pointer) {
	sessionid := portsip.ParamGetSessionid(params)
	err := SESSION.SetLocalPlayPcmData(pcmpath)
	if err != nil {
		log.Print(err)
		HANDLE.Hangup(sessionid)
		return
	}
	SESSION.RegisterAudioStreamCallback()
}

func onSIPCallClosed(params unsafe.Pointer) {
	SESSION.RemoveAudioStreamCallback()
	return
}
