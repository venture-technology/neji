package bot

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/venture-technology/neji/config"
	"github.com/venture-technology/neji/internal/usecase"
)

type DiscordBot struct {
	Token string
	uc    *usecase.NejiUseCase
}

func NewDiscordBot(Token string, uc *usecase.NejiUseCase) *DiscordBot {
	return &DiscordBot{
		Token: Token,
		uc:    uc,
	}
}

func (bot *DiscordBot) Setup() {

	conf := config.Get()

	dg, err := discordgo.New("Bot " + conf.Discord.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(bot.MessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. ")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()

}

func (bot *DiscordBot) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "deploy venture qa" {
		bot.uc.DeployVenture(s, m)
	}

}
