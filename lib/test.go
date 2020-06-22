package lib

import (
	"fmt"
	"strings"
	"time"

	d "github.com/bwmarrin/discordgo"
)

func TestRes(s string, m *d.MessageCreate) *d.MessageEmbed {
	domains := []string{"someshkar.com", "someshkar.net"}

	embed := &d.MessageEmbed{
		Author: &d.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: m.Author.AvatarURL("32"),
		},
		Title:       "Lookup for ***google***",
		Description: "Looks like a few common TLDs are available for ***google***",
		Color:       0x2977f5,
		Fields: []*d.MessageEmbedField{
			&d.MessageEmbedField{
				Name:  "Available Domains",
				Value: fmt.Sprintf("%s\n\n", strings.Join(domains, "\n")),
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return embed
}

// // TestRes is a test handler
// func TestRes(s string, m *d.MessageCreate) *d.MessageEmbed {
// 	embed := &d.MessageEmbed{
// 		Author: &d.MessageEmbedAuthor{
// 			Name:    m.Author.Username,
// 			IconURL: m.Author.AvatarURL("32"),
// 		},
// 		Title:       "Lookup for google.com",
// 		URL: "https://www.google.com",
// 		Description: "Seems like google.com isn't available :(",
// 		Color:       0x2977f5,
// 		Fields: []*d.MessageEmbedField{
// 			&d.MessageEmbedField{
// 				Name:   "Registrar",
// 				Value:  "Namecheap Inc.",
// 				Inline: true,
// 			},
// 			&d.MessageEmbedField{
// 				Name:   "Expiry",
// 				Value:  "January 26, 2029",
// 				Inline: true,
// 			},
// 		},
// 		Timestamp: time.Now().Format(time.RFC3339),
// 	}

// 	return embed
// }
