package main

import (
	"encoding/json"
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
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

type TranscribeResult struct {
	Results struct {
		Transcripts []struct {
			Transcript string `json:"transcript"`
		} `json:"transcripts"`
	} `json:"results"`
}

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

type ChatMessage struct {
	Question  string
	Answer    string
	PCMPath   string
	WAVPath   string
	JSONPath  string
	MP3Path   string
	Timestamp string
}

var chatHistory []ChatMessage

func refreshChatUI(chatBox *fyne.Container) {
	chatBox.Objects = nil
	for _, msg := range chatHistory {
		qLabel := widget.NewLabelWithStyle("Q: "+msg.Question, fyne.TextAlignLeading, fyne.TextStyle{})
		qLabel.Wrapping = fyne.TextWrapWord
		aLabel := widget.NewLabelWithStyle("A: "+msg.Answer, fyne.TextAlignLeading, fyne.TextStyle{})
		aLabel.Wrapping = fyne.TextWrapWord
		vbox := container.NewVBox(aLabel)
		scroll := container.NewVScroll(vbox)
		scroll.SetMinSize(fyne.NewSize(780, 100)) // 设置最小高度，可根据需要调整
		playBtn := widget.NewButton("播放语音", func(mp3 string) func() {
			return func() {
				exec.Command("cmd", "/C", "start", mp3).Run()
			}
		}(msg.MP3Path))
		chatBox.Add(container.NewVBox(
			qLabel,
			aLabel,
			playBtn,
			widget.NewLabel("时间: "+msg.Timestamp),
			widget.NewSeparator(),
		))
	}
	chatBox.Refresh()
}

func DownloadTranscribeResult(url, jsonPath string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	out, err := os.Create(jsonPath)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return jsonPath, err
}

func main() {
	if err := portaudio.Initialize(); err != nil {
		log.Fatalf("Failed to initialize PortAudio: %v", err)
	}
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
	HANDLE.SetRtpPortRange(41000, 43000)
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

	var chatBox *fyne.Container
	a := app.New()
	w := a.NewWindow("录音demo")
	statusLabel := widget.NewLabel("")
	isRecording := false
	var stopChan chan struct{}
	var recorder *Recorder

	oneClickBtn := widget.NewButton("开始对话", nil)
	oneClickBtn.OnTapped = func() {
		if !isRecording {
			isRecording = true
			oneClickBtn.SetText("停止")
			stopChan = make(chan struct{})
			go func() {
				os.RemoveAll("output")
				os.MkdirAll("output", 0755)

				timestamp := time.Now().Format("20060102_150405")
				pcmPath := fmt.Sprintf("output/output_%s.pcm", timestamp)
				wavPath := fmt.Sprintf("output/output_%s.wav", timestamp)
				jsonPath := fmt.Sprintf("output/output_%s.json", timestamp)
				mp3Path := fmt.Sprintf("output/answer_%s.mp3", timestamp)

				recorder = NewRecorder()
				statusLabel.SetText("请提问...")
				if err := recorder.Start(); err != nil {
					statusLabel.SetText(fmt.Sprintf("录音失败: %v", err))
					isRecording = false
					oneClickBtn.SetText("开始对话")
					return
				}

				select {
				case <-time.After(60 * time.Second):
				case <-stopChan:
				}
				recorder.StopAndSavePCM(pcmPath)
				PcmToWav(pcmPath, wavPath)

				statusLabel.SetText("上传到S3...")
				bucket := "awstestpbx"
				key := filepath.Base(wavPath)
				if err := UploadToS3(wavPath, bucket, key); err != nil {
					statusLabel.SetText(fmt.Sprintf("上传S3失败: %v", err))
					isRecording = false
					oneClickBtn.SetText("开始对话")
					return
				}

				statusLabel.SetText("转写中...")
				jobName := fmt.Sprintf("job-%s", timestamp)
				s3uri := "s3://" + bucket + "/" + key
				uri, err := StartTranscribeJob(jobName, s3uri)
				if err != nil {
					statusLabel.SetText(fmt.Sprintf("转写失败: %v", err))
					isRecording = false
					oneClickBtn.SetText("开始对话")
					return
				}
				_, err = DownloadTranscribeResult(uri, jsonPath)
				if err != nil {
					statusLabel.SetText(fmt.Sprintf("下载转写结果失败: %v", err))
					isRecording = false
					oneClickBtn.SetText("开始对话")
					return
				}

				jsonData, err := os.ReadFile(jsonPath)
				var transcript string
				if err == nil {
					var result TranscribeResult
					if err := json.Unmarshal(jsonData, &result); err == nil && len(result.Results.Transcripts) > 0 {
						transcript = result.Results.Transcripts[0].Transcript
					}
				}
				if transcript == "" {
					statusLabel.SetText("未找到转写内容")
					isRecording = false
					oneClickBtn.SetText("开始对话")
					return
				}

				messages := []map[string]string{
					{"role": "system", "content": "请用自然对话语气回答，不要用md格式，要像一个客服一样。"},
				}
				for _, msg := range chatHistory {
					messages = append(messages, map[string]string{"role": "user", "content": msg.Question})
					messages = append(messages, map[string]string{"role": "assistant", "content": msg.Answer})
				}
				messages = append(messages, map[string]string{"role": "user", "content": transcript})

				statusLabel.SetText("分析中...")
				answer, err := AskDeepSeek(messages)
				if err != nil {
					statusLabel.SetText(fmt.Sprintf("分析失败: %v", err))
					isRecording = false
					oneClickBtn.SetText("开始对话")
					return
				}

				statusLabel.SetText("语音合成中...")
				if err := SynthesizeSpeech(answer, mp3Path); err != nil {
					statusLabel.SetText(fmt.Sprintf("语音合成失败: %v", err))
					isRecording = false
					oneClickBtn.SetText("开始对话")
					return
				}

				chatHistory = append(chatHistory, ChatMessage{
					Question:  transcript,
					Answer:    answer,
					PCMPath:   pcmPath,
					WAVPath:   wavPath,
					JSONPath:  jsonPath,
					MP3Path:   mp3Path,
					Timestamp: timestamp,
				})
				refreshChatUI(chatBox)
				statusLabel.SetText("完成！")
				isRecording = false
				oneClickBtn.SetText("开始对话")
			}()
		} else {
			close(stopChan)
			statusLabel.SetText("录音已停止，处理中...")
		}
	}

	chatBox = container.NewVBox()
	scroll := container.NewVScroll(chatBox)
	scroll.SetMinSize(fyne.NewSize(780, 500))       // 设置较高的高度
	scroll.Direction = container.ScrollVerticalOnly // 只允许上下滚动

	w.SetContent(container.NewVBox(
		scroll, // 防止横向滚动条
		statusLabel,
		oneClickBtn,
	))

	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
	<-c
}
