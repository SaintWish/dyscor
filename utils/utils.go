package utils

import (
	"time"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

var (
	UserIDRegex = regexp.MustCompile("<@!?([0-9]+)>")
	RoleIDRegex = regexp.MustCompile("<@&?([0-9]+)>")
	ChannelIDRegex = regexp.MustCompile("<#([0-9]+)>")
)

func ComesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}

func GetGuildID(chanID string, s *discordgo.Session) (string, error) {
	channel, err := ChannelDetails(chanID, s)
	if err != nil {
		return nil, err
	}

	return channel.GuildID, nil
}

func SendUserDM(s *discordgo.Session, userID string, msg string) error {
	ch, err := s.UserChannelCreate(userID)
	if err != nil {
		return err
	}

	s.ChannelMessageSend(ch.ID, msg)
	return nil
}

func ChannelDetails(channelID string, s *discordgo.Session) (*discordgo.Channel, error) {
	channelDetails, err := s.State.Channel(channelID)
	if err == discordgo.ErrStateNotFound {
		channelDetails, err = s.Channel(channelID)
	}

	return channelDetails, err
}

func PermissionDetails(authorID, channelID string, s *discordgo.Session) (int, error) {
	perms, err = s.State.UserChannelPermissions(authorID, channelID)
	if err == discordgo.ErrStateNotFound {
		perms, err = s.UserChannelPermissions(authorID, channelID)
	}
	
	return perms, err
}

func GuildDetails(guildID string, s *discordgo.Session) (*discordgo.Guild, error) {
	guildDetails, err = s.State.Guild(guildID)
	if err == discordgo.ErrStateNotFound {
		guildDetails, err = s.Guild(guildID)
	}
	return guildDetails, err
}

func MemberDetails(guildID, memberID string, s *discordgo.Session) (*discordgo.Member, error) {
	member, err = s.State.Member(guildID, memberID)
	if err == discordgo.ErrStateNotFound {
		member, err = s.GuildMember(guildID, memberID)
	}
}

func DeleteMessage(m *discordgo.Message, s *discordgo.Session) {
	if m != nil {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
	}
}

func GetMessages(amount int, channelID string, s *discordgo.Session) (list []*discordgo.Message, err error) {
	list, err = s.ChannelMessages(channelID, amount, "", "", "")
}

func GetMessageAge(msg *discordgo.Message, s *discordgo.Session, m *discordgo.MessageCreate) (time.Duration, error) {
	then, err := msg.Timestamp.Parse()
	if err != nil {
		return time.Duration(0), err
	}
	return time.Since(then), nil
}

func ChangeStatus(s *discordgo.Session, msg string) {
	s.UpdateStatusComplex(discordgo.UpdateStatusData{Status: msg})
}