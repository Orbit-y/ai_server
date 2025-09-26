#ifndef PORTSIP_C_ICALLBACKPARAMETERS_H
#define PORTSIP_C_ICALLBACKPARAMETERS_H


#include "portsip_c_types.h"

#ifdef __cplusplus
extern "C" {
#endif

CSDK_EXPORT SIP_EVENT portsip_c_params_geteventtype(void* params);

CSDK_EXPORT long portsip_c_params_getsessionid(void* params);

CSDK_EXPORT const char* portsip_c_params_getcallerdisplayname(void* params);
CSDK_EXPORT const char* portsip_c_params_getcalleedisplayname(void* params);

CSDK_EXPORT const char* portsip_c_params_getcaller(void* params);
CSDK_EXPORT const char* portsip_c_params_getcallee(void* params);

CSDK_EXPORT bool portsip_c_params_getexistsearlymedia(void* params);

CSDK_EXPORT int portsip_c_params_getstatuscode(void* params);
CSDK_EXPORT const char* portsip_c_params_getstatustext(void* params);

CSDK_EXPORT long portsip_c_params_getreferid(void* params);
CSDK_EXPORT const char* portsip_c_params_getreferfrom(void* params);

CSDK_EXPORT const char* portsip_c_params_getreferto(void* params);

CSDK_EXPORT const char* portsip_c_params_getforwardto(void* params);

CSDK_EXPORT const unsigned char* portsip_c_params_getmessagedata(void* params);
CSDK_EXPORT int portsip_c_params_getmessagedatalength(void* params);

CSDK_EXPORT const char* portsip_c_params_getsignaling(void* params);

CSDK_EXPORT const char* portsip_c_params_getaudiocodecs(void* params);
CSDK_EXPORT const char* portsip_c_params_getvideocodecs(void* params);
CSDK_EXPORT const char* portsip_c_params_getscreencodecs(void* params);

CSDK_EXPORT bool portsip_c_params_getexistsvideo(void* params);
CSDK_EXPORT bool portsip_c_params_getexistsaudio(void* params);
CSDK_EXPORT bool portsip_c_params_getexistsscreen(void* params);

CSDK_EXPORT const char* portsip_c_params_getwaitingmessageaccount(void* params);

CSDK_EXPORT const char* portsip_c_params_getpresencesubject(void* params);

CSDK_EXPORT int portsip_c_params_geturgentnewwaitingmessagecount(void* params);

CSDK_EXPORT int portsip_c_params_getnewwaitingmessagecount(void* params);

CSDK_EXPORT int portsip_c_params_geturgentoldwaitingmessagecount(void* params);

CSDK_EXPORT int portsip_c_params_getoldwaitingmessagecount(void* params);

CSDK_EXPORT int portsip_c_params_getdtmftone(void* params);

CSDK_EXPORT long portsip_c_params_getsubscribeid(void* params);
CSDK_EXPORT const char* portsip_c_params_getplayedaudiofilename(void* params);

CSDK_EXPORT const char* portsip_c_params_getmimetype(void* params);
CSDK_EXPORT const char* portsip_c_params_getsubmimetype(void* params);
CSDK_EXPORT long portsip_c_params_getmessageid(void* params);

CSDK_EXPORT int portsip_c_params_getuserid(void* params);

CSDK_EXPORT const char* portsip_c_params_getblfmonitoreduri(void* params);
CSDK_EXPORT const char* portsip_c_params_getblfdialogstate(void* params);
CSDK_EXPORT const char* portsip_c_params_getblfdialogdirection(void* params);
CSDK_EXPORT const char* portsip_c_params_getblfdialogid(void* params);

#ifdef __cplusplus
}
#endif
#endif // PORTSIP_C_ICALLBACKPARAMETERS_H
