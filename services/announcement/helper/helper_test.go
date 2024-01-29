package helper

import (
	"testing"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func TestDecodeToken(t *testing.T) {
	// Set up the key for testing (replace with your actual key)
	viper.Set("JWTKEY", "your_test_key")
	defer viper.Reset()

	// Create a sample JWT token
	claims := &models.JwtClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(viper.GetString("JWTKEY")))
	if err != nil {
		t.Fatalf("Failed to sign JWT token: %v", err)
	}
	DecodeToken(signedToken)
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
