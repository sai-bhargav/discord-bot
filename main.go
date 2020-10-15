package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/discord-bot/commands"
)

func main() {
	fmt.Println("started")

	BotToken := os.Getenv("BOT_TOKEN")

	dg, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println("errored")
		fmt.Println(err)
	}

	dg.AddHandler(ready)
	dg.AddHandler(guildCreate)
	dg.AddHandler(messageCreate)
	dg.AddHandler(commands.XkcdComics)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("chillara bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
	fmt.Println("Started")
}

func ready(s *discordgo.Session, event *discordgo.Ready) {

	s.UpdateStatus(0, "starting")
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			_, _ = s.ChannelMessageSend(channel.ID, "Chillara bot under progress")
			return
		}
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	msgContent := m.Content
	res := strings.HasPrefix(msgContent, "+")

	if !res {
		return
	}

	msgContent = msgContent[1:]

	if msgContent == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if msgContent == "help" {
		s.ChannelMessageSend(m.ChannelID, "yem help kavali ra neeku , what are your wants")
	}
}
