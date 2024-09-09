package usecase

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/venture-technology/neji/internal/repository"
)

type NejiUseCase struct {
	nejiRepository repository.INejiRepository
}

func NewNejiUseCase(nejiRepository repository.INejiRepository) *NejiUseCase {
	return &NejiUseCase{
		nejiRepository: nejiRepository,
	}
}

func (uc *NejiUseCase) DeployVenture(s *discordgo.Session, m *discordgo.MessageCreate) {

	id, err := uuid.NewV7()
	if err != nil {
		log.Print(err)
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("# Deploy Backend Venture Inicializado \n - **Started_at: %v** \n - **ID: %s**", time.Now().Format("15:04:05"), id))

	time.Sleep(5 * time.Second)

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("# Deploy Backend Venture Finalizado \n - **Status: Success** \n - **Finished_at: %v**", time.Now().Format("15:04:05")))

}
