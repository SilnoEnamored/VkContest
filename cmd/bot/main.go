package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN_API")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message Updates
			continue
		}

		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, я бот! Нажми на кнопку, чтобы увидеть еще кнопки")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Кнопка 1"),
					tgbotapi.NewKeyboardButton("Кнопка 2"),
					tgbotapi.NewKeyboardButton("Кнопка 3"),
					tgbotapi.NewKeyboardButton("Кнопка 4"),
				),
			)
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		case "Кнопка 1":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы выбрали Кнопку 1")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Кнопка 1.1"),
					tgbotapi.NewKeyboardButton("Кнопка 1.2"),
				),
			)
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		case "Кнопка 2":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы выбрали Кнопку 2")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Кнопка 2.1"),
					tgbotapi.NewKeyboardButton("Кнопка 2.2"),
				),
			)
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		case "Кнопка 3":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы выбрали Кнопку 3")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Кнопка 3.1"),
					tgbotapi.NewKeyboardButton("Кнопка 3.2"),
				),
			)
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		case "Кнопка 4":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы выбрали Кнопку 4")
			msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Кнопка 4.1"),
					tgbotapi.NewKeyboardButton("Кнопка 4.2"),
				),
			)
			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
