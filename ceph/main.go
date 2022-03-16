package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var CredentialsProvider = aws.CredentialsProviderFunc(
	func(context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     "L5S0PMK8RKUWT3118YRJ",
			SecretAccessKey: "IZOEAt9oSzx65S4Ougkyrx1x1jcFeXegjHm2v5MC",
		}, nil
	})

var EndpointResolver = aws.EndpointResolverWithOptionsFunc(
	func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               "http://39.104.181.232:51011",
			HostnameImmutable: true,
		}, nil
	})

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(CredentialsProvider),
		config.WithEndpointResolverWithOptions(EndpointResolver),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	//PutObject(client)
	PutString(client)
	//GetFile(client)
	GetObject(client)
	ListBuckets(client)
	ListObjects(client)
	ListObjectsV2(client)
}

type S3ListBucketsAPI interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}

func GetAllBuckets(c context.Context, api S3ListBucketsAPI, input *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	return api.ListBuckets(c, input)
}

func ListBuckets(client *s3.Client) {
	input := &s3.ListBucketsInput{}

	result, err := GetAllBuckets(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving buckets:")
		fmt.Println(err)
		return
	}

	fmt.Println("Buckets:")

	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name + ": " + bucket.CreationDate.Format("2006-01-02 15:04:05 Monday"))
	}
}

func ListObjectsV2(client *s3.Client) {
	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("test"),
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}

type S3ListObjectsAPI interface {
	ListObjectsV2(ctx context.Context,
		params *s3.ListObjectsV2Input,
		optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)
}

func GetObjects(c context.Context, api S3ListObjectsAPI, input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	return api.ListObjectsV2(c, input)
}

func ListObjects(client *s3.Client) {
	bucket := "test"
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	}

	resp, err := GetObjects(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got error retrieving list of objects:")
		fmt.Println(err)
		return
	}

	fmt.Println("Objects in " + bucket + ":")

	for _, item := range resp.Contents {
		fmt.Println("Name:          ", *item.Key)
		fmt.Println("Last modified: ", *item.LastModified)
		fmt.Println("Size:          ", item.Size)
		fmt.Println("Storage class: ", item.StorageClass)
		fmt.Println("")
	}

	fmt.Println("Found", len(resp.Contents), "items in bucket", bucket)
	fmt.Println("")
}

type S3PutObjectAPI interface {
	PutObject(ctx context.Context,
		params *s3.PutObjectInput,
		optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func PutFile(c context.Context, api S3PutObjectAPI, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return api.PutObject(c, input)
}

func PutObject(client *s3.Client) {
	bucket := "test"
	filename := "/Users/wyb/temp/user2addresses.xlsx"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to open file " + filename)
		return
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	}

	_, err = PutFile(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got error uploading file:")
		fmt.Println(err)
		return
	}
}

func PutString(client *s3.Client) {
	bucket := "test"
	key := "abc"
	buf := bytes.NewBufferString("123abc! 你好，中国🇨🇳🌎🌍🌏¼½¾³²∞√㏒！")

	input := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(key),
		Body:          buf,
		ContentLength: int64(buf.Len()),
	}

	_, err := client.PutObject(context.TODO(), input, s3.WithAPIOptions(v4.SwapComputePayloadSHA256ForUnsignedPayloadMiddleware))
	if err != nil {
		fmt.Println("Got error uploading bytes:")
		fmt.Println(err)
		return
	}
}

func GetObject(client *s3.Client) {
	bucket := "test"
	key := "abc"

	request := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	response, err := client.GetObject(context.TODO(), request)
	if err != nil {
		panic(err)
	}

	fmt.Println("GetObject:", response)

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("GetObject:", string(result))

	filename := "/Users/wyb/temp/abc.txt"
	if err != ioutil.WriteFile(filename, result, os.FileMode(0644)) {
		panic(err)
	}
}

func GetFile(client *s3.Client) {
	bucket := "test"
	key := "/Users/wyb/temp/user2addresses.xlsx"
	filename := "/Users/wyb/temp/testfile.xlsx"

	request := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	response, err := client.GetObject(context.TODO(), request)
	if err != nil {
		panic(err)
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if err != ioutil.WriteFile(filename, result, os.ModePerm) {
		panic(err)
	}
}
