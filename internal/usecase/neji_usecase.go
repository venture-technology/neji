package usecase

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("# Deploy Backend Venture Inicializado \n - **Started_at: %v** \n - **ID: %s**", time.Now().Add(-5*time.Hour).Format("15:04:05"), id))

	if err := os.Chdir("/root/venture"); err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
		return
	}

	if err := runCommand("git", "pull"); err != nil {
		fmt.Println("Erro ao executar git pull:", err)
		return
	}

	if err := runCommand("go", "test", "./..."); err != nil {
		fmt.Println("Falha nos testes do Go:", err)
		return
	}

	if err := runCommand("docker", "compose", "up", "--build", "-d"); err != nil {
		fmt.Println("Erro ao executar docker compose:", err)
		return
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("# Deploy Backend Venture Finalizado \n - **Status: Success** \n - **Finished_at: %v**", time.Now().Add(-5*time.Hour).Format("15:04:05")))

}

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	logFile, err := os.OpenFile("deploy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo de log: %w", err)
	}
	defer logFile.Close()

	cmd.Stdout = logFile
	cmd.Stderr = logFile

	fmt.Printf("Executando: %s %s\n", name, args)
	return cmd.Run()
}
