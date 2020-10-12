package main

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func newPls() *pls {
	return &pls{}
}

type pls struct{}

func (pls *pls) handlePls(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		return
	}

	logrus.Debugf("initiating %v", args)

	// decide pls command
	switch args[0] {
	case "spinthebottle":
		if len(args) == 1 {
			pls.handleSpinTheBottle(s, m)
		}
	}
}

func (pls *pls) handleSpinTheBottle(s *discordgo.Session, m *discordgo.MessageCreate) {

	logrus.Debugf("spinning the bottle")

	// Find the channel that the message came from.
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		logrus.Errorf("could not find channel: %v", err)
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		// Could not find guild.
		logrus.Errorf("could not find guild: %v", err)
		return
	}

	// pick a random member that is not the bot
	randomIntegerwithinRange := rand.Intn(len(g.Members)-0) + 0

	var randomMember *discordgo.Member

	for randomMember == nil || randomMember.User.ID == s.State.User.ID {
		randomMember = g.Members[randomIntegerwithinRange]
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("and the chosen one is: %v", randomMember.Mention()))
}
