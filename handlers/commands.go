package handlers

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"HasiiChatBot/database"
	tele "gopkg.in/telebot.v3"
)

var startVideos = []string{
	"https://telegra.ph/file/9b7e1b820c72a14d90be7.mp4",
	"https://telegra.ph/file/a4d90b0cb759b67d68644.mp4",
	"https://telegra.ph/file/72f349b1386d6d9374a38.mp4",
	"https://telegra.ph/file/2b75449612172a96d4599.mp4",
	"https://telegra.ph/file/b3ac2d77205d5ded860de.mp4",
	"https://telegra.ph/file/58ae4ac86ef70dc8c8f6a.mp4",
	"https://telegra.ph/file/c6c1ac9aee4192a8a3747.mp4",
	"https://telegra.ph/file/55c840c8eba0555318f0d.mp4",
	"https://telegra.ph/file/e97715885d0a0cfbddaaa.mp4",
	"https://telegra.ph/file/943bb99829ec526c3f99a.mp4",
}

func getMainMenu(b *tele.Bot) (string, *tele.ReplyMarkup) {
	caption := "๏ <b>ʜᴇʏ, ɪ ᴀᴍ ʜᴀsɪɪ 🏓</b>\n" +
		"๏ <b>ʜɪᴛ ʜᴇʟᴘ ʙᴜᴛᴛᴏɴ ꜰᴏʀ ʜᴇʟᴘ</b>"

	menu := &tele.ReplyMarkup{}
	
	btnHelp := menu.Data("« ʜᴇʟᴘ »", "help")
		btnSource := menu.URL("❄️ ꜱᴏᴜʀᴄᴇ ❄️", "https://github.com/lakzexe/HasiiChatBot")
	btnAbout := menu.Data("☁️ ᴀʙᴏᴜᴛ ☁️", "about")
	
	btnOwner := menu.URL("🥀 ᴏᴡɴᴇʀ 🥀", "https://t.me/Hasindu_Lakshan")
	btnSupport := menu.URL("✨ ꜱᴜᴘᴘᴏʀᴛ ✨", "https://t.me/Hasindu_Lakshan")
	
	btnAddMe := menu.URL("✦ ᴀᴅᴅ ᴍᴇ ʙᴀʙʏ ✦", "https://t.me/"+b.Me.Username+"?startgroup=true")

	menu.Inline(
		menu.Row(btnAddMe),
		menu.Row(btnHelp),
		menu.Row(btnSource, btnAbout),
		menu.Row(btnOwner, btnSupport),
	)
	return caption, menu
}

func RegisterCommands(b *tele.Bot, db *database.Database) {
	b.Handle("/start", func(c tele.Context) error {
		rand.Seed(time.Now().UnixNano())
		videoURL := startVideos[rand.Intn(len(startVideos))]

		caption, menu := getMainMenu(b)

		video := &tele.Video{
			File:    tele.FromURL(videoURL),
			Caption: caption,
		}

		return c.Send(video, menu, tele.ModeHTML)
	})

	b.Handle("/ping", func(c tele.Context) error {
		return c.Send("Pong! 🏓")
	})

	// toggle bot
	b.Handle("/chatbot", func(c tele.Context) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Send("Usage: /chatbot [on|off]")
		}

		action := strings.ToLower(args[0])
		if action == "on" {
			// remove from disabled list
			_, err := db.Hasii.DeleteOne(context.Background(), map[string]interface{}{"chat_id": c.Chat().ID})
			if err != nil {
				return err
			}
			return c.Send("Chatbot Enabled!")
		} else if action == "off" {
			// add to disabled list
			// avoid dupes
			count, _ := db.Hasii.CountDocuments(context.Background(), map[string]interface{}{"chat_id": c.Chat().ID})
			if count == 0 {
				_, err := db.Hasii.InsertOne(context.Background(), map[string]interface{}{"chat_id": c.Chat().ID})
				if err != nil {
					return err
				}
			}
			return c.Send("Chatbot Disabled!")
		}
		
		return c.Send("Invalid option. Use 'on' or 'off'.")
	})

	b.Handle(&tele.Btn{Unique: "help"}, func(c tele.Context) error {
		helpText := "<b><u>ʜᴀsɪɪ ᴄʜᴀᴛʙᴏᴛ ᴄᴏᴍᴍᴀɴᴅs</u></b>\n\n" +
			"๏ <b>/start</b> - Sᴛᴀʀᴛ ᴛʜᴇ ʙᴏᴛ ᴀɴᴅ sᴇᴇ ᴛʜᴇ ᴍᴀɪɴ ᴍᴇɴᴜ.\n" +
			"๏ <b>/ping</b> - Cʜᴇᴄᴋ ɪꜰ ᴛʜᴇ ʙᴏᴛ ɪs ᴀʟɪᴠᴇ (Rᴇᴘʟɪᴇs ᴡɪᴛʜ Pᴏɴɢ!).\n" +
			"๏ <b>/chatbot on</b> -Eɴᴀʙʟᴇ ᴛʜᴇ ᴄʜᴀᴛʙᴏᴛ ᴛᴏ ᴀᴜᴛᴏ-ʀᴇᴘʟʏ ɪɴ ᴛʜɪs ɢʀᴏᴜᴘ.\n" +
			"๏ <b>/chatbot off</b> - Dɪsᴀʙʟᴇ ᴛʜᴇ ᴄʜᴀᴛʙᴏᴛ ᴀᴜᴛᴏ-ʀᴇᴘʟʏ ɪɴ ᴛʜɪs ɢʀᴏᴜᴘ.\n\n" +
			"<b>Nᴏᴛᴇ: Tʜᴇ ʙᴏᴛ ᴀᴜᴛᴏᴍᴀᴛɪᴄᴀʟʟʏ ʟᴇᴀʀɴs ꜰʀᴏᴍ ᴜsᴇʀs ᴡʜᴇɴ ᴛʜᴇʏ ʀᴇᴘʟʏ ᴛᴏ ᴇᴀᴄʜ ᴏᴛʜᴇʀ!</b>"

		menu := &tele.ReplyMarkup{}
		btnBack := menu.Data("« ʙᴀᴄᴋ »", "back")
		menu.Inline(menu.Row(btnBack))

		// edit msg
		err := c.EditCaption(helpText, menu, tele.ModeHTML)
		if err != nil {
			return err
		}
		return c.Respond()
	})

	// back to main menu
	b.Handle(&tele.Btn{Unique: "back"}, func(c tele.Context) error {
		caption, menu := getMainMenu(b)

		err := c.EditCaption(caption, menu, tele.ModeHTML)
		if err != nil {
			return err
		}
		return c.Respond()
	})
	
	b.Handle(&tele.Btn{Unique: "about"}, func(c tele.Context) error {
		aboutText := "<b><u>ᴀʙᴏᴜᴛ ʜᴀsɪɪ ᴄʜᴀᴛʙᴏᴛ</u></b>\n\n" +
			"ɴᴏᴛʜɪɴɢ ꜱᴘᴇᴄɪᴀʟ, ʙᴜᴛ ɢᴏᴛ ʙɪɢ ᴘᴏᴛᴇɴᴛɪᴀʟ ᴛᴏ ʙᴇ ᴀ ɢʀᴇᴀᴛ ᴄʜᴀᴛʙᴏᴛ ɪɴ ʏᴏᴜʀ ɢʀᴏᴜᴘ.\n" +
			"ɪ ᴏʙꜱᴇʀᴠᴇ ᴄᴏɴᴠᴇʀꜱᴀᴛɪᴏɴꜱ ᴀɴᴅ ʟᴇᴀʀɴ ʜᴏᴡ ᴛᴏ ʀᴇᴘʟʏ ᴛᴏ ᴛᴇxᴛ ᴀɴᴅ ꜱᴛɪᴄᴋᴇʀꜱ ɴᴀᴛᴜʀᴀʟʟʏ.\n\n" +
			"<b>Version:</b> 2.0.0\n" +
			"<b>Developer:</b> @Hasindu_Lakshan"

		menu := &tele.ReplyMarkup{}
		btnBack := menu.Data("« ʙᴀᴄᴋ »", "back")
		menu.Inline(menu.Row(btnBack))

		// edit msg
		err := c.EditCaption(aboutText, menu, tele.ModeHTML)
		if err != nil {
			return err
		}
		return c.Respond()
	})
}
