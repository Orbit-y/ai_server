package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
)

func UploadToS3(localFile, bucket, key string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := s3.NewFromConfig(cfg)
	file, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	return err
}

func StartTranscribeJob(jobName, s3uri string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}
	client := transcribe.NewFromConfig(cfg)
	_, err = client.StartTranscriptionJob(context.TODO(), &transcribe.StartTranscriptionJobInput{
		TranscriptionJobName: &jobName,
		LanguageCode:         "zh-CN",
		MediaFormat:          "wav",
		Media: &types.Media{
			MediaFileUri: &s3uri,
		},
	})
	if err != nil {
		return "", err
	}
	for {
		out, err := client.GetTranscriptionJob(context.TODO(), &transcribe.GetTranscriptionJobInput{
			TranscriptionJobName: &jobName,
		})
		if err != nil {
			return "", err
		}
		status := out.TranscriptionJob.TranscriptionJobStatus
		if status == "COMPLETED" {
			return *out.TranscriptionJob.Transcript.TranscriptFileUri, nil
		}
		if status == "FAILED" {
			return "", fmt.Errorf("transcribe failed")
		}
		time.Sleep(5 * time.Second)
	}
}
