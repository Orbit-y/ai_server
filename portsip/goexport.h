extern void GofnOnMessageCallback(void* params);
extern int  GofnLogCallback(uintptr_t obj,
	int level,
	char* subsystem,
	char* appName,
	char* file,
	int line,
	char* message,
	char* messageWithHeaders);

extern int  GofnAudioRawCallback(uintptr_t obj,
	long sessionId,
	int type,
	unsigned char* data,
	int dataLength,
	int samplingFreqHz);

extern int GofnVideoRawCallback(uintptr_t obj,
	long sessionId,
	int type,
	int width,
	int height,
	unsigned char* data,
	int dataLength);

extern int GofnRecordAudioRawCallback(uintptr_t obj,
	unsigned char* data,
	int dataLength,
	int samplingFreqHz);

extern  int GofnRTPPacketCallback(uintptr_t obj, long sessionId, int mediaType, int direction, unsigned char* RTPPacket, int packetSize);