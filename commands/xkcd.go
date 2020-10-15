package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	util "github.com/discord-bot/utils"

	"github.com/bwmarrin/discordgo"
)

type xkcdResponse struct {
	Month string `json:'month'`
	Num   int    `json:'num'`
	Img   string `json:'img'`
}

func XkcdComics(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	msgContent := m.Content
	res := strings.HasPrefix(msgContent, "+")

	if !res {
		return
	}

	msgContent = msgContent[1:]

	if msgContent == "xkcd" {
		s.ChannelMessageSend(m.ChannelID, randomComic())
	}

}

func randomComic() string {

	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(2372)

	fetchUrl := fmt.Sprintf("https://xkcd.com/%v/info.0.json", res)

	r := util.GetRequest(fetchUrl)
	defer r.Body.Close()

	x := new(xkcdResponse)
	json.NewDecoder(r.Body).Decode(x)

	return x.Img
}
