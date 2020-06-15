package handlers

import (
	"strings"

	d "github.com/bwmarrin/discordgo"
	"github.com/someshkar/domainbot/lib"
)

// MainHandler handles single domain lookups
func MainHandler(s *d.Session, m *d.MessageCreate) {
	// Ignore messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, "domain ") {
		mSlice := strings.Fields(m.Content)

		if mSlice[1] == "all" {
			// Handle multiple domains
			s.ChannelMessageSend(m.ChannelID, lib.AllDomainRes(mSlice[2]))
			return
		}

		// Check if it's a valid domain
		if !lib.IsDomain(mSlice[1]) {
			s.ChannelMessageSend(m.ChannelID, "Please enter a valid domain!")
			return
		}

		// Handle single domain
		s.ChannelMessageSend(m.ChannelID, lib.SingleDomainRes(mSlice[1]))
	}
}
