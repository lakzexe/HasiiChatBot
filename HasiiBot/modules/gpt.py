from HasiiBot import app
from pyrogram.enums import ChatAction, ParseMode
from pyrogram import filters
import g4f


SYSTEM_PROMPT = (
    "You are Hasii, a very friendly, funny, and helpful AI assistant for Telegram. "
    "IMPORTANT: You must always reply in everyday spoken Sinhala (Singlish or Sinhala Unicode). "
    "Be casual, talk like a close friend, and use some emojis. "
    "Keep your replies concise and engaging."
)


@app.on_message(filters.command(["chat"], prefixes=[".", "!", "/"]))
async def chat_gpt(bot, message):
    try:
        await bot.send_chat_action(message.chat.id, ChatAction.TYPING)

        name = message.from_user.first_name if message.from_user else "User"

        if len(message.command) < 2:
            await message.reply_text(
                f"**ආයුබෝවන් {name}! 👋**\n"
                f"**මට ඕනෙ මොනවද කියන්න.**\n"
                f"**උදාහරණ:** `/chat ලංකාවේ අගනුවර මොකක්ද?`"
            )
        else:
            query = message.text.split(" ", 1)[1]
            response = await g4f.ChatCompletion.create_async(
                model=g4f.models.default,
                messages=[
                    {"role": "system", "content": SYSTEM_PROMPT},
                    {"role": "user", "content": query},
                ],
            )
            await message.reply_text(f"{response}")
    except Exception as e:
        await message.reply_text(f"⚠️ **දෝෂයක්:** `{e}`")
