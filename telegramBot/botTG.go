package telegramBot

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"os"
	"strconv"
	"time"
)

func InitBot() (telebot.ChatID, *telebot.Bot, error) {

	// Получаем переменные
	telegramToken := os.Getenv("TELEGRAM_TOKEN")
	telegramChatID := os.Getenv("TELEGRAM_CHAT_ID")

	// Проверка переменных
	if telegramToken == "" || telegramChatID == "" {
		return telebot.ChatID(0), nil, fmt.Errorf("TELEGRAM_TOKEN или TELEGRAM_CHAT_ID не установлены")
	}

	//var chatID telebot.ChatID

	// Конвертация ChatID в int64
	id, err := strconv.ParseInt(telegramChatID, 10, 64)
	if err != nil {
		return telebot.ChatID(0), nil, fmt.Errorf("некорректный ChatID: %v", err)
	}
	var chatID = telebot.ChatID(id)

	// Создаем бота
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  telegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return telebot.ChatID(0), nil, fmt.Errorf("ошибка создания бота: %v", err)
	}

	return chatID, bot, nil
}
