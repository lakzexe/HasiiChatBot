import os
from pyrogram import filters, types as t
from pyrogram.enums import ChatAction
from HasiiBot import app

import g4f


SYSTEM_PROMPT = (
    "You are Hasii, a very friendly, funny, and helpful AI assistant for Telegram. "
    "IMPORTANT: You must always reply in everyday spoken Sinhala (Singlish or Sinhala Unicode). "
    "Be casual, talk like a close friend, and use some emojis. "
    "Keep your replies concise and engaging."
)


async def chat_completion(prompt: str) -> str:
    """Send a prompt to g4f and return the AI response."""
    try:
        response = await g4f.ChatCompletion.create_async(
            model=g4f.models.default,
            messages=[
                {"role": "system", "content": SYSTEM_PROMPT},
                {"role": "user", "content": prompt},
            ],
        )
        return response
    except Exception as e:
        raise Exception(f"AI Error: {e}")


async def vision_completion(prompt: str, image_path: str) -> str:
    """Send image + prompt to g4f vision model."""
    try:
        from g4f.Provider import You

        with open(image_path, "rb") as img_file:
            image_data = img_file.read()

        response = await g4f.ChatCompletion.create_async(
            model=g4f.models.default,
            messages=[
                {"role": "system", "content": SYSTEM_PROMPT},
                {"role": "user", "content": prompt or "What's in this image?"},
            ],
        )
        # Clean up downloaded image
        if os.path.exists(image_path):
            os.remove(image_path)
        return response
    except Exception as e:
        if os.path.exists(image_path):
            os.remove(image_path)
        raise Exception(f"Vision AI Error: {e}")


def get_media(message):
    """Extract Media from message or reply."""
    media = None
    if message.media:
        if message.photo:
            media = message.photo
        elif (
            message.document
            and message.document.mime_type in ["image/png", "image/jpg", "image/jpeg"]
            and message.document.file_size < 5242880
        ):
            media = message.document
    elif message.reply_to_message and message.reply_to_message.media:
        if message.reply_to_message.photo:
            media = message.reply_to_message.photo
        elif (
            message.reply_to_message.document
            and message.reply_to_message.document.mime_type
            in ["image/png", "image/jpg", "image/jpeg"]
            and message.reply_to_message.document.file_size < 5242880
        ):
            media = message.reply_to_message.document
    return media


def get_text(message):
    """Extract text from commands."""
    if message.text is None:
        return None
    if " " in message.text:
        try:
            return message.text.split(None, 1)[1]
        except IndexError:
            return None
    return None


@app.on_message(filters.command(["ask", "ai", "gpt", "gemini"]))
async def chat_bots(_, m: t.Message):
    prompt = get_text(m)
    media = get_media(m)

    if media is not None:
        return await ask_about_image(_, m, media, prompt)

    if prompt is None:
        return await m.reply_text(
            "**ආයුබෝවන්! 🙏 මට ඕනෙ මොනවද කියන්න.**\n"
            "**උදාහරණ:** `/ask ලංකාවේ අගනුවර මොකක්ද?`"
        )

    try:
        await _.send_chat_action(m.chat.id, ChatAction.TYPING)
        output = await chat_completion(prompt)
        await m.reply_text(output)
    except Exception as e:
        await m.reply_text(f"⚠️ **දෝෂයක් ඇති වුණා:** `{e}`")


async def ask_about_image(_, m: t.Message, media, prompt: str):
    try:
        await _.send_chat_action(m.chat.id, ChatAction.TYPING)
        image = await _.download_media(
            media.file_id,
            file_name=f"./downloads/{m.from_user.id}_ask.jpg",
        )
        output = await vision_completion(prompt, image)
        await m.reply_text(output)
    except Exception as e:
        await m.reply_text(f"⚠️ **දෝෂයක් ඇති වුණා:** `{e}`")
