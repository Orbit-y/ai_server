package main

import (
	"encoding/binary"
	"github.com/gordonklaus/portaudio"
	"log"
	"os"
	"time"
)

const (
	sampleRate = 8000
	channels   = 1
)

type Recorder struct {
	stream   *portaudio.Stream
	buf      []int16
	frames   []int16
	running  bool
	stopChan chan struct{}
}

func NewRecorder() *Recorder {
	r := &Recorder{
		buf:      make([]int16, 1024),
		frames:   make([]int16, 0, sampleRate*60),
		stopChan: make(chan struct{}),
	}
	return r
}

func (r *Recorder) Start() error {
	var err error
	r.stream, err = portaudio.OpenDefaultStream(channels, 0, float64(sampleRate), len(r.buf), &r.buf)
	if err != nil {
		return err
	}
	if err := r.stream.Start(); err != nil {
		return err
	}

	r.running = true
	go r.loop()
	return nil
}

func (r *Recorder) loop() {
	for r.running {
		select {
		case <-r.stopChan:
			return
		default:
			if err := r.stream.Read(); err != nil {
				log.Printf("录音读取失败:%v", err)
				return
			}
			r.frames = append(r.frames, r.buf...)
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func (r *Recorder) StopAndSavePCM(filename string) error {
	r.running = false
	close(r.stopChan)
	if r.stream != nil {
		r.stream.Stop()
		r.stream.Close()
	}
	return savePCM(filename, r.frames)
}

func savePCM(filename string, data []int16) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, v := range data {
		f.Write([]byte{byte(v), byte(v >> 8)})
	}
	return nil
}

func PcmToWav(pcmPath, wavPath string) error {
	pcmData, err := os.ReadFile(pcmPath)
	if err != nil {
		return err
	}
	wavFile, err := os.Create(wavPath)
	if err != nil {
		return err
	}
	defer wavFile.Close()
	var (
		audioFormat   uint16 = 1
		numChannels   uint16 = 1
		sampleRate    uint32 = 8000
		bitsPerSample uint16 = 16
		byteRate             = sampleRate * uint32(numChannels) * uint32(bitsPerSample/8)
		blockAlign    uint16 = numChannels * bitsPerSample / 8
		dataLen              = uint32(len(pcmData))
	)
	wavFile.Write([]byte("RIFF"))
	binary.Write(wavFile, binary.LittleEndian, uint32(36+dataLen))
	wavFile.Write([]byte("WAVEfmt "))
	binary.Write(wavFile, binary.LittleEndian, uint32(16))
	binary.Write(wavFile, binary.LittleEndian, audioFormat)
	binary.Write(wavFile, binary.LittleEndian, numChannels)
	binary.Write(wavFile, binary.LittleEndian, sampleRate)
	binary.Write(wavFile, binary.LittleEndian, byteRate)
	binary.Write(wavFile, binary.LittleEndian, blockAlign)
	binary.Write(wavFile, binary.LittleEndian, bitsPerSample)
	wavFile.Write([]byte("data"))
	binary.Write(wavFile, binary.LittleEndian, dataLen)
	wavFile.Write(pcmData)
	return nil
}
