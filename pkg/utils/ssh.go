package utils

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func ConnectSSH(user, password, host, port string) (*ssh.Client, *ssh.Session, error) {

	log.Printf("Conectando ao SSH %s@%s password: %s", user, host, password)

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), sshConfig)
	if err != nil {
		log.Print(err)
		return nil, nil, fmt.Errorf("falha ao conectar ao SSH: %w", err)
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("falha ao criar sessão SSH: %w", err)
	}

	return client, session, nil
}

func ChangeDirectory(session *ssh.Session, dir string) error {
	cmd := fmt.Sprintf("cd %s", dir)
	return session.Run(cmd)
}

func ExecuteCommand(client *ssh.Client, command string) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("falha ao criar nova sessão: %w", err)
	}
	defer session.Close()

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	err = session.Run(command)
	if err != nil {
		log.Printf("Erro ao executar comando: %s\nSaída: %s\nErro: %s", command, stdout.String(), stderr.String())
		return fmt.Errorf("falha ao executar comando: %w", err)
	}

	log.Printf("Comando executado com sucesso: %s\nSaída: %s", command, stdout.String())
	return nil
}
