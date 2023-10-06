package asciiartfs

import "strings"

// Print the full outcome
func PrintBanners(banners, arr []string) string {
	art := ""
	for _, word := range banners {
		if word == "\n" {
			art += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, j := range word {
				if j >= 32 {
					artSplit := strings.Split(arr[j-32], "\n")
					if i < len(artSplit) {
						for _, ascii := range artSplit[i] {
							art += string(ascii)
						}
					} else {
						continue
					}
				}
			}
			art += "\n"
		}
	}
	return art
}
