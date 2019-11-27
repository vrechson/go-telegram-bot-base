package app

import (
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/whoismath/go-telegram-bot-base/commands"
	"github.com/whoismath/go-telegram-bot-base/config"
	repositories "github.com/whoismath/go-telegram-bot-base/database/repositories"
	storages "github.com/whoismath/go-telegram-bot-base/database/storages"
	botFactory "github.com/whoismath/go-telegram-bot-base/factory/bot"
	dbFactory "github.com/whoismath/go-telegram-bot-base/factory/database"
)

// TelegramBotApp is the main structure containing all data types and the bot itself
type TelegramBotApp struct {
	conf       *config.Config
	bot        *tgbotapi.BotAPI
	storage    *storages.Handler
	repository *repositories.Handler
}

// CreateApp setup every structure from TelegramBotApp
func CreateApp(conf *config.Config) *TelegramBotApp {

	// Init the telegram bot
	bot, err := botFactory.BotFactory(conf)
	if err != nil {
		log.Fatal("Can't create a telegram bot: ", err)
	}

	// Init the database
	db, err := dbFactory.DatabaseFactory(conf)
	if err != nil {
		log.Fatal("Couldn't create database: ", err)
	}

	// Init the database storage
	storage, err := storages.CreateStorages(db)
	if err != nil {
		log.Fatal("Couldn't setup database: ", err)
	}

	// Init the database Repository
	repository, err := repositories.CreateRepository(db)
	if err != nil {
		log.Fatal("Couldn't setup database: ", err)
	}

	return &TelegramBotApp{conf, bot, storage, repository}
}

// Start is the function that start the bot service
func (TelegramBotApp *TelegramBotApp) Start() {
	updates, err := TelegramBotApp.getUpdates()
	if err != nil {
		log.Fatal("[!] Error: Can't get updates")
	}
	fmt.Println("[+] Initializating telegram bot")

	// Start the cronjob
	go TelegramBotApp.startCron()

	// Handle telegram api updates
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			go func() {
				TelegramBotApp.handleCommand(&update)
			}()
		}

	}
}

// This method get updates from telegram bot API
func (TelegramBotApp *TelegramBotApp) getUpdates() (tgbotapi.UpdatesChannel, error) {
	// Select between pooling and webhook (heroku works only with webhook)
	if TelegramBotApp.conf.UseWebhook != true {
		return TelegramBotApp.setupPolling()
	}
	return TelegramBotApp.setupWebhook()
}

// This method setup a pooling
func (TelegramBotApp *TelegramBotApp) setupPolling() (tgbotapi.UpdatesChannel, error) {
	TelegramBotApp.bot.RemoveWebhook()
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 5
	fmt.Println("[+] Pooling method selected")
	return TelegramBotApp.bot.GetUpdatesChan(updateConfig)
}

// This method setup a webhook
func (TelegramBotApp *TelegramBotApp) setupWebhook() (tgbotapi.UpdatesChannel, error) {
	_, err := TelegramBotApp.bot.SetWebhook(tgbotapi.NewWebhook(TelegramBotApp.conf.WebhookURL + "/" + TelegramBotApp.bot.Token))
	if err != nil {
		log.Fatal("[!] Webhook problem: ", err)
		//return nil, err
	}
	updates := TelegramBotApp.bot.ListenForWebhook("/" + TelegramBotApp.bot.Token)
	go http.ListenAndServe(":"+TelegramBotApp.conf.Port, nil)

	fmt.Println("[+] Webhook method selected")

	return updates, nil

}

// This method handle commands sent to the bot
func (TelegramBotApp *TelegramBotApp) handleCommand(update *tgbotapi.Update) {
	command := update.Message.Command()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "some text")

	switch command {

	// Showing help menu
	case commands.Help:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "[!] ")
		TelegramBotApp.bot.Send(msg)

	// Command test
	case commands.Ping:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "pong")
		TelegramBotApp.bot.Send(msg)
	}

}
