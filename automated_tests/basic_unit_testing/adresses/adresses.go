package adresses

import (
	"strings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// testing 



// TypeAdresses is a function that receives a string and returns a string
func TypeAdresses(address string) string {
	validTypes := []string{"street", "avenue", "boulevard", "road", "place", "court", "square", "crescent", "trail", "parkway", "terrace", "highway"}
	addressInLowerCase := strings.ToLower(address)
	firstWordAddress := strings.Split(addressInLowerCase, " ")[0]

	addressIsType := false

	for _, validType := range validTypes {
		if firstWordAddress == validType {
			addressIsType = true
		}
	}

	if addressIsType {
		titleCaser := cases.Title(language.English)
		return titleCaser.String(firstWordAddress)
	}

	return "none"
}
