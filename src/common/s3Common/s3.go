package s3Common

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/disintegration/imaging"
	"image/png"
	"math/rand"
	"mime/multipart"
	"time"
)

func ImageUpload(file *multipart.FileHeader, imgType ImgType) (string, error) {
	ctx := context.Background()
	meta, ok := imgMeta[imgType]
	if !ok {
		return "", fmt.Errorf("not available meta info for imgType - %+v", imgType)
	}
	bucket := meta.bucket()
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("fail to open file")
	}
	defer src.Close()
	img, err := imaging.Decode(src)

	if err != nil {
		return "", fmt.Errorf("fail to load image")
	}
	if imgType == ImgTypeProduct {
		if meta.width < 1 || meta.height < 1 {
			if (meta.width < 1 && meta.height < img.Bounds().Size().Y) ||
				(meta.height < 1 && meta.width < img.Bounds().Size().X) {
				img = imaging.Resize(img, meta.width, meta.height, imaging.Lanczos)
			}
		} else {
			img = imaging.Fill(img, meta.width, meta.height, imaging.Center, imaging.Lanczos)
		}
	}

	buf := new(bytes.Buffer)
	if err := imaging.Encode(buf, img, imaging.PNG, imaging.PNGCompressionLevel(png.BestCompression)); err != nil {
		return "", fmt.Errorf("fail to encode png image")
	}

	filename := fmt.Sprintf("%s.png", fileNameGenerateRandom())
	result, err := AwsClientS3Uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(fmt.Sprintf("%s/%s", meta.path, filename)),
		Body:        buf,
		ContentType: aws.String("image/png"),
	})
	fmt.Println(result.Location)
	if err != nil {
		return "", fmt.Errorf("fail to upload image to s3 - bucket:%s / key:%s/%s", bucket, meta.path, filename)
	}
	return filename, nil
}

func ImageGetSignedURL(fileName string, imgType ImgType) (string, error) {
	meta, ok := imgMeta[imgType]
	if !ok {
		return "", fmt.Errorf("not available meta info for imgType - %+v", imgType)
	}
	url := fmt.Sprintf("https://%s/%s/%s", meta.domain(), meta.path, fileName)

	signedURL, err := AwsS3Signer.Sign(url, time.Now().Add(meta.expireTime))
	if err != nil {
		return "", fmt.Errorf("get signed url - origianl url:%s", url)
	}
	return signedURL, nil
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func fileNameGenerateRandom() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
