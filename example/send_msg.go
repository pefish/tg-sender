package main

import (
	"log"
	"time"

	i_logger "github.com/pefish/go-interface/i-logger"
	tg_sender "github.com/pefish/tg-sender"
)

func main() {
	err := do()
	if err != nil {
		log.Fatal(err)
		return
	}

	time.Sleep(2 * time.Second)
}

func do() error {
	sender := tg_sender.NewTgSender(&i_logger.DefaultLogger, "")

	err := sender.SendMsg(&tg_sender.MsgStruct{
		ChatId: "",
		Msg:    "test",
	}, 0)
	if err != nil {
		return err
	}

	return nil
}

// go run ./_example/
