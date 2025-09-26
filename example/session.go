package main

import "C"
import (
	"goportsipsdk/portsip"
	"log"
	"os"
	"runtime/cgo"
)

type Session struct {
	localData    []byte
	localOffset  int64
	localLength  int64
	remoteData   []byte
	remoteLength int64
	sessionId    int64
}

func CreateSession(sessionId int64) (*Session, error) {
	session := new(Session)
	session.sessionId = sessionId
	return session, nil
}

func (S *Session) SetLocalPlayPcmData(localpcm string) error {
	data, err := os.ReadFile(localpcm)
	if err != nil {
		return err
	}
	S.localData = data

	fileInfo, err := os.Stat(localpcm)
	if err != nil {
		return err
	}
	S.localLength = fileInfo.Size()
	log.Printf("load local pcm data length %d", S.localLength)
	return nil
}

func (S *Session) RegisterAudioStreamCallback() {
	portsip.AddAudioRawCallbacks(S.sessionId, audioRawCallbackFunc)
	handle := cgo.NewHandle(S)
	ret := HANDLE.EnableAudioStreamCallback(S.sessionId, true, portsip.DIRECTION_SEND_RECV, uintptr(handle))
	if ret != 0 {
		log.Printf("Failed to enable local  audio stream call back for session %d , ret %d", S.sessionId, ret)
	}
}

func (S *Session) RemoveAudioStreamCallback() {
	portsip.RemoveAudioRawCallbacks(S.sessionId)
	ret := HANDLE.EnableAudioStreamCallback(S.sessionId, false, portsip.DIRECTION_SEND_RECV, 0)
	if ret != 0 {
		log.Printf("Failed to stop local  audio stream call back for session %d , ret %d", S.sessionId, ret)
	}
}

func (S *Session) putData(data []byte, dataLength int) {
	S.remoteData = append(S.remoteData, data[:dataLength]...)
	S.remoteLength = S.remoteLength + int64(dataLength)
	log.Printf("session %d recived data length %d", S.sessionId, S.remoteLength)
}

func (S *Session) readData(data []byte, dataLength int) {
	if S.localOffset+int64(dataLength) > S.localLength {
		S.localOffset = 0
		return
	}
	log.Printf("session %d send data offset %d length %d", S.sessionId, S.localOffset, dataLength)
	copy(data, S.localData[S.localOffset:S.localOffset+int64(dataLength)])
	S.localOffset = S.localOffset + int64(dataLength)
}

func audioRawCallbackFunc(obj cgo.Handle, sessionId int64, dtype int, data []byte, dataLength int, samplingFreqHz int) {
	if obj == 0 {
		return
	}
	session := obj.Value().(*Session)
	switch dtype {
	case portsip.DIRECTION_SEND:
		session.readData(data, dataLength)
	case portsip.DIRECTION_RECV:
		session.putData(data, dataLength)
	}
}
