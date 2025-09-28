package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gordonklaus/portaudio"
	"goportsipsdk/portsip"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
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

func DownloadTranscribeResult(url, wavPath string) (string, error) {
	dir := filepath.Dir(wavPath)
	outPath := filepath.Join(dir, "output.json")
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	out, err := os.Create(outPath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return outPath, err
}

func main() {
	if err := portaudio.Initialize(); err != nil {
		log.Fatalf("Failed to initialize PortAudio: %v", err)
	}
	log.Printf("PortAudio initialized successfully")
	defer portaudio.Terminate()

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

	a := app.New()
	w := a.NewWindow("录音demo")
	statusLabel := widget.NewLabel("")

	recorder := NewRecorder()
	btn := widget.NewButton("开始录音", nil)
	btn.OnTapped = func() {
		if !recorder.running {
			recorder = NewRecorder()
			btn.SetText("停止录音")
			if err := recorder.Start(); err != nil {
				statusLabel.SetText(fmt.Sprintf("录音失败: %v", err))
			} else {
				statusLabel.SetText("正在录音...")
			}
		} else {
			btn.SetText("开始录音")
			// 录音结束，保存PCM
			os.MkdirAll("output", 0755)
			if err := recorder.StopAndSavePCM("output/output.pcm"); err != nil {
				statusLabel.SetText(fmt.Sprintf("停止录音失败: %v", err))
			} else {
				// PCM转WAV
				if err := PcmToWav("output/output.pcm", "output/output.wav"); err != nil {
					statusLabel.SetText(fmt.Sprintf("PCM转WAV失败: %v", err))
				} else {
					statusLabel.SetText("录音文件已保存为 output/output.pcm 和 output/output.wav")
					// 上传到S3
					bucket := "pbxs3test"
					key := "output.wav"
					statusLabel.SetText("正在上传到S3...")
					if err := UploadToS3("output/output.wav", bucket, key); err != nil {
						statusLabel.SetText(fmt.Sprintf("上传S3失败: %v", err))
						return
					}
					statusLabel.SetText("上传成功，开始转写...")
					jobName := "test-job"
					s3uri := "s3://" + bucket + "/" + key
					uri, err := StartTranscribeJob(jobName, s3uri)
					if err != nil {
						statusLabel.SetText(fmt.Sprintf("转写失败: %v", err.Error()))
						return
					}
					statusLabel.SetText("转写完成，正在下载结果...")
					jsonPath, err := DownloadTranscribeResult(uri, "output/output.wav")
					if err != nil {
						statusLabel.SetText(fmt.Sprintf("下载转写结果失败: %v", err))
						return
					}
					statusLabel.SetText(fmt.Sprintf("转写结果已保存: %s\n格式为JSON", jsonPath))
				}
			}
		}
	}

	w.SetContent(container.NewVBox(btn, statusLabel))
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()

	<-c
}
