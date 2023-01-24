package command

import (
	"github.com/bwmarrin/discordgo"
)

type UserCommand struct {
	Name string
	Desc string
	Disabled bool
	IsDM bool

	Perms int

	SubCommands map[string]SubUserCommand
	Run func(*discordgo.MessageCreate, *Command, []string)
}

type SubUserCommand struct {
	Name string
	Desc string
	IsDM bool

	Perms int

	Run func(*discordgo.MessageCreate, *SubCommand, []string)
}