package lib

import (
	"fmt"
	"log"
	"time"

	d "github.com/bwmarrin/discordgo"
	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

// SingleDomainRes generates the Domainbot
// response for a single domain
func SingleDomainRes(domain string, m *d.MessageCreate) string {
	isRegistered, registrar, expiryStr := checkDomain(domain)

	if isRegistered {
		t, _ := time.Parse(time.RFC3339, expiryStr)
		expiry := t.Format("January 2, 2006")

		// res := domain + " is registered at " + registrar + " and will expire on " + expiry + "."
		res := fmt.Sprintf("%s %s is registered at %s and will expire on %s.",
			m.Author.Mention(), domain, registrar, expiry)
		log.Printf("'%s' returned '%s'", domain, res)
		return res
	}

	// res := domain + " may be available!"
	res := fmt.Sprintf("%s %s may be available!", m.Author.Mention(), domain)
	log.Printf("'%s' returned '%s'", domain, res)
	return res
}

// checkDomain checks if a domain is available and returns relevant info if it is
func checkDomain(domain string) (taken bool, registrar string, expiryDate string) {
	raw, err := whois.Whois(domain)
	if err != nil {
		log.Println(err)
	}

	result, err := whoisparser.Parse(raw)
	if err != nil {
		if err == whoisparser.ErrDomainNotFound {
			return false, "", ""
		}
		log.Println(err)
	}

	return true, result.Registrar.Name, result.Domain.ExpirationDate
}
