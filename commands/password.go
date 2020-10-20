package commands

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	util "github.com/discord-bot/utils"
)

// PASSWORD_CHARS has all the characters that can be included to generate the password
const PASSWORD_CHARS = "abcdefghijklmnopqrstuvwxyz!\"#$%&'()*+,-./:;<=>?@[]^_`{|}~\\1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// PASSWORD_REGEX is command verification regex
var PASSWORD_REGEX = regexp.MustCompile(`^pwd (?P<length>[0-9]*$)`)

// PwdGenerator will return a random generated password
func PwdGenerator(s *discordgo.Session, m *discordgo.MessageCreate) {

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

	// Defaults to 10 characters password if length is not specified
	if msgContent == "pwd" {
		s.ChannelMessageSend(m.ChannelID, "pwd 10")
	} else {
		s.ChannelMessageSend(m.ChannelID, generator(msgContent))
	}
}

func generator(content string) string {
	regexCheck := PASSWORD_REGEX.MatchString(content)

	if regexCheck == false {
		return "Please check the input"
	}
	passwordLength := extractPasswordLength(content)

	return generatePassword(passwordLength)
}

func generatePassword(pwdLength string) string {

	if len(pwdLength) > 3 {
		pwdLength = pwdLength[0:3]
	}

	length, err := strconv.Atoi(pwdLength)
	if err != nil {
		util.Logger("Errored converting password to integer")
		log.Fatal(err)
	}

	fmt.Println(length)

	res := ""
	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())
		res = fmt.Sprintf("%v%v", res, string(PASSWORD_CHARS[rand.Intn(93)]))
	}

	return res
}

func extractPasswordLength(content string) string {
	match := PASSWORD_REGEX.FindStringSubmatch(content)
	result := make(map[string]string)
	for i, name := range PASSWORD_REGEX.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result["length"]
}
