package main

import (
	"goportsipsdk/portsip"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var HANDLE *portsip.SDKHandle

const (
	domain    = "test.com"
	caller    = "100"
	password  = "A1a3s5d7"
	callee    = "101@test.com"
	pbxserver = "192.168.2.125"
	transport = "UDP:8926"
	sipport   = 5060
	pcmpath   = "./park_default"
)

func main() {
	dispatcher, err := portsip.CreateAbstractCallbackDispatcher()
	if err != nil {
		log.Fatalf("Failed to create PortSIP AbstractCallbackDispatcher: %v", err)
	}
	defer dispatcher.DestroyAbstractCallbackDispatcher()
	dispatcher.SetAbstractCallbackDispatcherOnMessageCallback(MyOnMessageCallbackFunc)

	HANDLE, err = portsip.Initialize(dispatcher,
		false,
		"",
		100,
		"PortSIP SDK",
		"0.0.0.0",
		1,
		1,
		"",
		"",
		false,
		nil,
		nil,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to initialize PortSIP SDK: %v", err)
	}
	defer HANDLE.Uninitialize()

	HANDLE.SetLog(
		portsip.LogTypeCallBack,
		portsip.LogLevelError,
		"test",
		"./",
		10,
		0,
	)

	HANDLE.SetLicenseKey("PORTSIP_UC_LICENSE")

	ret := HANDLE.SetRtpPortRange(41000, 43000)
	if ret != 0 {
		log.Fatalf("Failed to SetRtpPortRange: %d", ret)
	}

	HANDLE.SetSrtpPolicy(portsip.SRTP_POLICY_PREFER, true)

	HANDLE.EnableSessionTimer(120, portsip.SESSION_REFERESH_UAC)

	userId := HANDLE.AddUser(
		caller,
		caller,
		caller,
		password,
		transport,
		domain,
		pbxserver,
		sipport,
		"",
		0,
		"",
		0,
	)
	if userId < 0 {
		log.Fatalf("Failed to add user")
	}
	defer HANDLE.RemoveUser(userId)

	HANDLE.AddAudioCodec(portsip.AUDIOCODEC_OPUS)
	HANDLE.AddAudioCodec(portsip.AUDIOCODEC_G729)
	HANDLE.AddAudioCodec(portsip.AUDIOCODEC_PCMA)
	HANDLE.AddAudioCodec(portsip.AUDIOCODEC_PCMU)
	HANDLE.AddAudioCodec(portsip.AUDIOCODEC_G722)
	HANDLE.AddAudioCodec(portsip.AUDIOCODEC_AMRWB)
	HANDLE.AddAudioCodec(portsip.AUDIOCODEC_AMR)
	HANDLE.AddAudioCodec(portsip.AUDIOCODEC_DTMF)

	if result := HANDLE.RegisterServer(userId, 3600, 3); result != 0 {
		log.Printf("Registration failed with code: %d", result)
	}

	defer HANDLE.UnRegisterServer(userId, 0)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c
}
