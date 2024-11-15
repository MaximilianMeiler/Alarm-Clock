package webhook

import (
	"log"

	"github.com/gtuk/discordwebhook"
)

type DiscordWebHook struct {
	username string
	content  string

	title       string
	description string
	color       string
	imgURL      string
}

func NewDiscordWebHook(
	username string,
	content string,
	title string,
	description string,
	color string,
	imgURL string,
) *DiscordWebHook {

	return &DiscordWebHook{
		username:    username,
		content:     content,
		title:       title,
		description: description,
		color:       color,
		imgURL:      imgURL,
	}

}

func (d *DiscordWebHook) SendMessage(webhookURL string) error {
	image := discordwebhook.Image{
		Url: &d.imgURL,
	}

	embeds := discordwebhook.Embed{
		Title:       &d.title,
		Description: &d.description,
		Color:       &d.color,
		Image:       &image,
	}

	message := discordwebhook.Message{
		Username: &d.username,
		Content:  &d.content,
		Embeds:   &[]discordwebhook.Embed{embeds},
	}

	err := discordwebhook.SendMessage(webhookURL, message)

	return err

}

func SendMessageExample(discord_webhook string) {
	var username = "Wake Up"
	var content = "Simple Message"

	var title = "Testing Embed"
	var description = "Embed Description."
	var color = "383483" // Must be a numerical value not nex
	var url2 = "https://media2.giphy.com/media/v1.Y2lkPTc5MGI3NjExbXluOXlraTJzNzU1eHhqdnNta2xld3l2anA0bG0zemM1NTBiajIxYyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/qpCvOBBmBkble/giphy.webp"

	image := discordwebhook.Image{
		Url: &url2,
	}

	embeds := discordwebhook.Embed{
		Title:       &title,
		Description: &description,
		Color:       &color,
		Image:       &image,
	}

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
		Embeds:   &[]discordwebhook.Embed{embeds},
	}

	err := discordwebhook.SendMessage(discord_webhook, message)
	if err != nil {
		log.Fatal(err)
	}
}
