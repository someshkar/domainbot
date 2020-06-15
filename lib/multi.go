package lib

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/likexian/whois-go"
	whoisparser "github.com/likexian/whois-parser-go"
)

// AllDomainRes returns the Domainbot
// response for multiple popular domains
func AllDomainRes(s string) string {
	domains := CheckDomains(s)

	if len(domains) == 1 {
		res := domains[0] + " may be available!"
		log.Printf("'all %s' returned '%s'", s, res)
		return res
	}

	if len(domains) > 0 {
		res := fmt.Sprintf("%s and %s may be available!", strings.Join(domains[:len(domains)-1], ", "), domains[len(domains)-1])
		log.Printf("'all %s' returned '%s'", s, res)
		return res
	}

	res := "None of the common TLDs are available for this."
	log.Printf("'all %s' returned '%s'", s, res)
	return res
}

// CheckDomains checks which popular domains
// are avaiable for a particular string
func CheckDomains(s string) (available []string) {
	tlds := []string{"com", "org", "net", "co", "io", "dev", "xyz", "tech"}

	var wg sync.WaitGroup
	wg.Add(len(tlds))

	var availableDomains []string

	for _, tld := range tlds {
		go func(tld string) {
			domainAvailable(s+"."+tld, &availableDomains)
			wg.Done()
		}(tld)
	}

	wg.Wait()

	return availableDomains
}

// domainAvailable checks if a domain is available
// and appends it to a slice if it is
func domainAvailable(domain string, availableDomains *[]string) {
	raw, err := whois.Whois(domain)
	if err != nil {
		log.Println(err)
	}

	_, err = whoisparser.Parse(raw)
	if err != nil {
		if err == whoisparser.ErrDomainNotFound {
			*availableDomains = append(*availableDomains, domain)
		}
	}
}
