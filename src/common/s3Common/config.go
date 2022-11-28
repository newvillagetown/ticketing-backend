package s3Common

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/cloudfront/sign"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"main/common/envCommon"
	"time"
)

var AwsClientS3 *s3.Client
var AwsClientS3Uploader *manager.Uploader
var AwsClientS3Downloader *manager.Downloader
var AwsS3Signer *sign.URLSigner

type ImgType uint8

const (
	ImgTypeProduct = ImgType(0)
)

type imgMetaStruct struct {
	bucket     func() string
	domain     func() string
	path       string
	width      int
	height     int
	expireTime time.Duration
}

var imgMeta = map[ImgType]imgMetaStruct{
	ImgTypeProduct: {
		bucket: func() string {
			return fmt.Sprintf("%s-%s-s3", envCommon.Env.Env, envCommon.Env.Project)
		},
		domain:     func() string { return fmt.Sprintf("%s-%s-s3.breathings.net", envCommon.Env.Env, envCommon.Env.Project) },
		path:       "image",
		width:      512,
		height:     512,
		expireTime: 24 * time.Hour,
	},
}
