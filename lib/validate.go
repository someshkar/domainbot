package lib

import "regexp"

// IsDomain checks if a string is a valid domain
func IsDomain(toTest string) bool {
	// rxPat := regexp.MustCompile(`^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$`)
	rxPat := regexp.MustCompile(`^[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.[a-zA-Z]{2,}$`)

	// TODO: Load tlds.json and check is tld is supported
	return rxPat.MatchString(toTest)
}