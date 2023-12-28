package helper

import (
	"time"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func DecodeToken(token string) (*models.JwtClaims, error) {
	claims := &models.JwtClaims{}
	key := viper.Get("JWTKEY").(string)
	jwttoken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if !jwttoken.Valid {
		return nil, err
	}
	return claims, nil
}

func GetImageURL(s *s3.S3, object_key string) (string, error) {
	bucketName := viper.Get("OBJECT_STORAGE_BUCKET_NAME").(string)
	req, _ := s.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(object_key),
	})
	url, err := req.Presign(1 * time.Hour)
	if err != nil {
		return "", err
	}
	return url, nil
}

func DeleteImage(s *s3.S3, object_key string) error {
	bucketName := viper.Get("OBJECT_STORAGE_BUCKET_NAME").(string)
	_, err := s.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(object_key),
	})
	return err
}
