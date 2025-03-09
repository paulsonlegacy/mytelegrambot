package main

import (
	"log"
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Get bot token from environment variable
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		// Get bot token from .env file
		envToken, err := fetchEnv("TELEGRAM_BOT_TOKEN") // Use separate variable
		if err == nil {
			botToken = envToken
		} else {
			log.Fatal("TELEGRAM_BOT_TOKEN not set or empty")
		}
	}

	// Initialize bot
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Bot authorized on account: %s", bot.Self.UserName)
	}

	// Setup update listener
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // Ignore non-message updates
			msgText := update.Message.Text

			// Define a default response
			response := "Hello! I am your personal bot. Ask me anything about my creator"

			// Handle special commands
			if msgText == "/start" {
				response = "Welcome! I am your personal bot. Type anything and I'll reply!"
			} else if msgText == "/about" {
				response = "My creator is a passionate Go developer learning bot development!"
			} else if msgText == "/token" {
				response = botToken
			}

			// Send message back to user
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			bot.Send(msg)
		}
	}

}