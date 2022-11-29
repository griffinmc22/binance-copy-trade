package general

import (
	"fmt"
	"log"

	wh "github.com/etaaa/go-webhooks"
)

func Sendwebhook(image string, webhook string, itemname string, itemurl string) {
	fmt.Println(image)
	fmt.Println(itemname)
	hook := wh.Webhook{
		Username: "Monitor Name",
		Embeds: []wh.Embed{
			{
				Title:     itemname,
				Url:       itemurl,
				Timestamp: wh.GetTimestamp(),      // RETURNS NEW TIMESTAMP ACCORDING TO DISCORD'S FORMAT
				Color:     wh.GetColor("#00ff00"), // RETURNS COLOR ACCORDING TO DISCORD'S FORMAT
				Footer: wh.EmbedFooter{
					Text: "ChefsOnya#0559",
				},
				Thumbnail: wh.EmbedThumbnail{
					Url: "https:" + image,
				},
				Fields: []wh.EmbedFields{
					{Name: "Restock Alert!!", Value: "In Stock - Click link to visit product"},
				},
			},
		},
	}
	// SEND THE WEBHOOK
	if err := wh.SendWebhook(webhook, hook, true); err != nil {
		log.Fatal(err)
	}
}

func Senderrorwebhook(webhook string) {

	hook := wh.Webhook{
		Username: "Monitor Name",
		Embeds: []wh.Embed{
			{
				Title:     "Error Alert	",
				Timestamp: wh.GetTimestamp(),      // RETURNS NEW TIMESTAMP ACCORDING TO DISCORD'S FORMAT
				Color:     wh.GetColor("#992D22"), // RETURNS COLOR ACCORDING TO DISCORD'S FORMAT
				Footer: wh.EmbedFooter{
					Text: "ChefsOnya#0559",
				},
			},
		},
	}
	// SEND THE WEBHOOK
	if err := wh.SendWebhook(webhook, hook, true); err != nil {
		log.Fatal(err)
	}
}
