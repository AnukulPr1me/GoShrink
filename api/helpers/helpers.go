package helpers

import (
    "os"
    "strings"
)

func EnforcHTTP(url string) string {
	if url[:4] != "http"{
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {
	if url == os.Getenv("DOMAIN"){
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.split(newURL, "/")[0]

	if newURL == os.Getenv("DOMAIN"){
        return false
    }

	return true
}