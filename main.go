package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/discord-bot/commands"
	util "github.com/discord-bot/utils"
	"github.com/gorilla/mux"
)

// Ping test to prevent Heroku dyno from idling
type PingResponse struct {
	Message string
}

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
	dg.AddHandler(commands.PwdGenerator)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	fmt.Println("chillara bot is now running.  Press CTRL-C to exit.")
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/ping", pingTest)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

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

	util.Logger(msgContent + "-" + m.Author.Username)
	msgContent = msgContent[1:]

	if msgContent == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if msgContent == "help" {
		s.ChannelMessageSend(m.ChannelID, "yem help kavali ra neeku , what are your wants")
	}
}

func pingTest(w http.ResponseWriter, r *http.Request) {

	util.Logger("Ping successful")

	data := PingResponse{}
	data.Message = "Awake and Alive!!!!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
