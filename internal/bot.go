package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func Run() {
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	dg.Identify.Intents = discordgo.IntentsMessageContent
	dg.AddHandler(msgHandler)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Bot online now")
	out := make(chan os.Signal, 1)
	signal.Notify(out, os.Interrupt, syscall.SIGTERM)
	<-out

	dg.Close()
}

func msgHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
		return
	}
	// channel, _ := s.UserChannelCreate(m.Author.ID)
	// s.ChannelMessageSend(channel.ID, "Pong!")
}
