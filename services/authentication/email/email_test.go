package email

import (
	"testing"

	"github.com/spf13/viper"
)

func TestSend(t *testing.T) {
	viper.Set("SMTPHOST", "your_test_key")
	viper.Set("SMTPPORT", "your_test_key")
	mockEmail := Email{}
	mockEmail.Send()
}