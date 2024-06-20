package main

import (
	"time"

	go_logger "github.com/pefish/go-logger"
	tg_sender "github.com/pefish/tg-sender"
)

func main() {
	err := do()
	if err != nil {
		go_logger.Logger.Error(err)
		return
	}

	time.Sleep(2 * time.Second)
}

func do() error {
	sender := tg_sender.NewTgSender("")

	err := sender.SendMsg(&tg_sender.MsgStruct{
		ChatId: "",
		Msg:    "test",
	}, 0)
	if err != nil {
		return err
	}

	return nil
}
