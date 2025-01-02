package fes3tools

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/goccy/go-yaml"
)

type ComInfo struct {
	Url     string `yaml:"url"`
	User    string `yaml:"user"`
	Pwd     string `yaml:"pwd"`
	Email   string `yaml:"email"`
	Name    string `yaml:"name"`
	Pgpkey  string `yaml:"pgpkey"`
	Signkey string `yaml:"signkey"`
}

func GetComInfo(bucket string, partner string) ComInfo {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(partner),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer output.Body.Close()

	// Read the file content into memory
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, output.Body)
	if err != nil {
		fmt.Println("Error reading object body:", err)
	}

	// unmarshal yaml and return struct
	var c ComInfo
	if err := yaml.Unmarshal(buf.Bytes(), &c); err != nil {
		fmt.Println(err)
	}

	return c
}

func GetFile(bucket string, fname string) string {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fname),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer output.Body.Close()

	// Read the file content into memory
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, output.Body)
	if err != nil {
		fmt.Println("Error reading object body:", err)
	}

	return buf.String()
}

func PutFile(bucket string, fname string, data string) string {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	// This uploads the contents of the buffer to S3
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fname),
		Body:   bytes.NewReader([]byte(data)),
	})
	if err != nil {
		return err.Error()
	}

	return "Suceess"
}

func CheckFile(bucket string, fname string) bool {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	_, err = client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fname),
	})
	if err != nil {
		return false
	}

	return true
}
