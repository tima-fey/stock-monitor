package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartBot(token string)  (bot *tgbotapi.BotAPI, err error){
	bot, err = tgbotapi.NewBotAPI(token)
	//if err != nil {
	//	log.Panic(err)
	//	return
	//}
	return bot, err
	//bot.Debug = true
	//
	//log.Printf("Authorized on account %s", bot.Self.UserName)
	//
	//u := tgbotapi.NewUpdate(0)
	//u.Timeout = 60
	//
	//updates, err := bot.GetUpdatesChan(u)

	//for update := range updates {
	//	if update.Message == nil { // ignore any non-Message Updates
	//		continue
	//	}
	//
	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//	msg.ReplyToMessageID = update.Message.MessageID
	//
	//	bot.Send(msg)
	//}
}
