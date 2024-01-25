package tg_sender

import (
	go_test_ "github.com/pefish/go-test"
	"testing"
)

func TestTelegramSender_SendMsg(t *testing.T) {
	successCount := 0

	go_test_.Equal(t, 0, successCount)
}
