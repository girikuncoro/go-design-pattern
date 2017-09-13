package main

import "fmt"

type botPrototype interface {
	clone() botPrototype
	SendMessage() string
}

type BotCache struct {
	bot botPrototype
}

func (self *BotCache) Store(bot botPrototype) {
	self.bot = bot
}

func (self *BotCache) Load() botPrototype {
	bot := self.bot
	return bot.clone()
}

type Bot struct {
	message string
}

func (self *Bot) Init() {
	// init bot takes resources and time
}

func (self *Bot) SendMessage() string {
	return self.message
}

func (self *Bot) clone() botPrototype {
	return &Bot{self.message}
}

func main() {
	bot := Bot{"Hello"}
	bot.Init()

	botCache := BotCache{}
	botCache.Store(&bot)

	cloned := botCache.Load()

	fmt.Printf("%p %s\n", &bot, bot.SendMessage())
	fmt.Printf("%p %s\n", &cloned, cloned.SendMessage())
}
