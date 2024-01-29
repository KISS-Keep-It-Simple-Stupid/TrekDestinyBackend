package helper

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type JwtClaims struct {
	jwt.StandardClaims
	CustomField string `json:"customField"`
}

func TestDecodeToken(t *testing.T) {
	originalJWTKey := viper.GetString("JWTKEY")
	defer func() {
		viper.Set("JWTKEY", originalJWTKey)
	}()

	viper.Set("JWTKEY", "mockJWTKey")

	mockClaims := &JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
		CustomField: "example",
	}
	mockToken := generateMockToken(mockClaims)

	_, _ = DecodeToken(mockToken)

	invalidToken := "invalidToken"
	_, _ = DecodeToken(invalidToken)

	viper.Set("JWTKEY", "expiredJWTKey")
	expiredToken := generateMockToken(mockClaims)
	_, _ = DecodeToken(expiredToken)
}

func generateMockToken(claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(viper.GetString("JWTKEY"))
	signedToken, _ := token.SignedString(key)
	return signedToken
}

func TestGetImageURL(t *testing.T) {
	originalBucketName := viper.GetString("OBJECT_STORAGE_BUCKET_NAME")
	defer func() {
		viper.Set("OBJECT_STORAGE_BUCKET_NAME", originalBucketName)
	}()

	viper.Set("OBJECT_STORAGE_BUCKET_NAME", "mockBucketName")

	mockS3Client := createMockS3Client()

	objectKey := "mockObjectKey"
	_, _ = GetImageURL(mockS3Client, objectKey)

	viper.Set("OBJECT_STORAGE_BUCKET_NAME", "")
	_, _ = GetImageURL(mockS3Client, objectKey)
}

func createMockS3Client() *s3.S3 {
	return s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})))
}

type MockS3Client struct{}

func (m *MockS3Client) DeleteObject(input *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	return nil, nil
}

func TestDeleteImage(t *testing.T) {
	_ = &MockS3Client{}

	objectKey := "test_object_key"

	viper.Set("OBJECT_STORAGE_BUCKET_NAME", "test_bucket")

	s3Session := s3.New(session.New(), &aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Endpoint:         aws.String("http://localhost:4569"),
	})

	_ = DeleteImage(s3Session, objectKey)
}