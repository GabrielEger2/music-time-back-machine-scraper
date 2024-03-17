package main

import "fmt"

func main() {
	music := scrapeMusicData()

	// Print data
	for _, m := range music {
		fmt.Println(m.Rank, m.Title, m.Artist)
	}
}
