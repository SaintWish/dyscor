package command

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

type CommandSystem struct {
	s *discordgo.Session

	slashCommands map[string]SlashCommand
	slashCmdLock sync.RWMutex

	userCommands map[string]UserCommand
	userCmdLock sync.RWMutex
}

func NewCommandSystem(session *discordgo) *CommandSystem {
	return &CommandApp {
		s: session
	}
}