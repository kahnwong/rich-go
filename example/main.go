package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kahnwong/rich-go/client"
)

func main() {
	err := client.Login(os.Getenv("DISCORD_APP_ID"))
	if err != nil {
		panic(err)
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		ActivityType: client.ActivityTypes.Playing,
		State:        "Heyy!!!",
		Details:      "I'm running on rich-go :)",
		LargeImage:   "largeimageid",
		LargeText:    "This is the large image :D",
		SmallImage:   "smallimageid",
		SmallText:    "And this is the small image",
		Party: &client.Party{
			ID:         "-1",
			Players:    15,
			MaxPlayers: 24,
		},
		Timestamps: &client.Timestamps{
			Start: &now,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "GitHub",
				Url:   "https://github.com/hugolgst/rich-go",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	// Discord will only show the presence if the app is running
	// Sleep for a few seconds to see the update
	fmt.Println("Sleeping...")
	time.Sleep(time.Second * 10)
}
