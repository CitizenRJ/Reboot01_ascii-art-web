package asciiartfs

import "fmt"

// Print the full outcome
func PrintBanners(banners, arr []string) string {
	art := ""
	num := 0
	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Println()
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				art += arr[int(n)+i]

			}
			art += "\n"

		}
		art += "\n"

	}
	return art
}
