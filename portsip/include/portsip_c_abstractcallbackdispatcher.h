#ifndef PORTSIP_C_ABSTRACTCALLBACKDISPATCHER_H
#define PORTSIP_C_ABSTRACTCALLBACKDISPATCHER_H

#include "portsip_c_types.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef int (*fnOnMessageCallback)(void* params);


CSDK_EXPORT void* portsip_c_create_abstractcallbackdispatcher();

CSDK_EXPORT void portsip_c_set_onmessage_callback(void* dispatcher, fnOnMessageCallback cbhandle);

CSDK_EXPORT void portsip_c_destory_abstractcallbackdispatcher(void* dispatcher);


#ifdef __cplusplus
}
#endif

#endif
