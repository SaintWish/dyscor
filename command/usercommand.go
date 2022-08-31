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
	Run func(*discordgo.Session, *discordgo.MessageCreate, *Command, []string)
}

type SubUserCommand struct {
	Name string
	Desc string
	IsDM string

	Perms int

	Run func(*discordgo.Session, *discordgo.MessageCreate, *SubCommand, []string)
}