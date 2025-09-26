package portsip

/*
#include "include/portsip_c_serverbase.h"
#include "include/portsip_c_abstractcallbackdispatcher.h"
#include "goexport.h"

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type (
	OnMessageCallbackFunc func(pointer unsafe.Pointer)
)

var OnMessageCallback OnMessageCallbackFunc

//export GofnOnMessageCallback
func GofnOnMessageCallback(params unsafe.Pointer) {
	if OnMessageCallback != nil {
		OnMessageCallback(params)
	}
}

type DispatcherHandle struct {
	dispatcher unsafe.Pointer
	callback   OnMessageCallbackFunc
}

func CreateAbstractCallbackDispatcher() (*DispatcherHandle, error) {
	dispatcher := C.portsip_c_create_abstractcallbackdispatcher()
	if dispatcher == nil {
		return nil, fmt.Errorf("Failed to create PortSIP AbstractcallbackDispatcher(")
	}
	C.portsip_c_set_onmessage_callback(dispatcher, (C.fnOnMessageCallback)(C.GofnOnMessageCallback))
	return &DispatcherHandle{dispatcher: dispatcher}, nil
}

func (d *DispatcherHandle) SetAbstractCallbackDispatcherOnMessageCallback(callback OnMessageCallbackFunc) error {
	OnMessageCallback = callback
	return nil
}

func (d *DispatcherHandle) DestroyAbstractCallbackDispatcher() error {
	if d.dispatcher != nil {
		C.portsip_c_destory_abstractcallbackdispatcher(d.dispatcher)
		d.dispatcher = nil
	}
	return nil
}
