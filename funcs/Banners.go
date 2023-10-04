package funcs

import (
	"strings"
)

func Banners(banner string) string {
	
	var banners = map[string]string{
		"shadow":      "shadow.txt",
		"standard":        "standard.txt",
		"thinkertoy":      "thinkertoy.txt",
		"doom":      "doom.txt",
		
		// Add more banners here as needed
	}
	banner = strings.ToLower(banner)
	if strings.HasSuffix(banner, ".txt"){
		banner = strings.TrimRight(banner, ".txt")
	}
	if banner1,ok := banners[banner] ; ok {
		return banner1
	}
	
	return ("")
}
