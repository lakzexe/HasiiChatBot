# HasiiChatBot

<p align="center">
   <img src="https://files.catbox.moe/bti5oz.png" alt="HasiiChatBot Banner" width="420" />
</p>

> AI-first Telegram assistant with MongoDB memory, Pyrogram power, and VPS-only deployment.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Bot Info](#bot-info)
- [Requirements](#requirements)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Running the Bot](#running-the-bot)
- [VPS Deployment Checklist](#vps-deployment-checklist)
- [Support](#support)

## Overview
HasiiChatBot is an intelligent conversational Telegram bot that runs best on a self-managed VPS. It uses Pyrogram + MongoDB to remember chats, respects Telegram FloodWait limits, and provides a smooth experience in both groups and private chats.

## Features
- AI-powered replies with GPT-backed fallback pipeline
- Chatbot toggle per group with inline admin controls
- MongoDB memory for contextual responses (text + stickers)
- FloodWait-safe send helper to keep the bot online
- Sticker-to-text learning so users can teach new reactions
- `/chatbot` admin command plus `/start` and `/help` flows
- Ready-to-run VPS scripts (`start`) for hands-free bootstrapping

## Bot Info
| Item | Details |
| --- | --- |
| Owner | [@Hasindu_Lakshan](https://t.me/Hasindu_Lakshan) |
| Tech Stack | Python 3 • Pyrogram • MongoDB |
| Hosting Target | Linux-based VPS (Ubuntu 20.04+ recommended) |
| Database | Mongo Atlas / self-hosted MongoDB |
| Source | Public repo (this project) |

## Requirements
- Python 3.10+
- MongoDB URI with write access
- Telegram API credentials (`API_ID`, `API_HASH`)
- Bot token from [@BotFather](https://t.me/BotFather)
- VPS or local machine running Linux with `bash`

## Quick Start
1. Clone the repository:
   ```bash
   git clone https://github.com/lakzexe/HasiiChatBot.git
   cd HasiiChatBot
   ```
2. Install dependencies:
   ```bash
   pip3 install -r requirements.txt
   ```
3. Copy `sample.env` to `.env` and fill in your secrets.
4. Launch the bot:
   ```bash
   bash start
   ```

## Configuration
Provide these environment variables (or update `config.py`).

| Variable | Description |
| --- | --- |
| `API_ID` | Telegram API ID from my.telegram.org |
| `API_HASH` | Telegram API hash paired with the ID |
| `BOT_TOKEN` | Bot token issued by BotFather |
| `MONGO_URL` | MongoDB connection string (SRV or standard) |
| `OWNER_ID` | Numeric Telegram user ID of the bot owner |
| `OWNER_USERNAME` | Owner username (used in helper texts) |
| `GPT_API` | OpenAI/GPT API key for AI replies |
| `SUPPORT_GROUP` | Optional support group username/link |
| `UPDATE_CHANNEL` | Optional channel username/link |

## Running the Bot
```bash
source .env      # or export the vars manually
pip3 install -r requirements.txt
bash start       # runs the launch script included in /start
```

The `start` script handles screen/tmux friendly launches on a VPS. Customize it if you want systemd services or Docker alternatives.

## VPS Deployment Checklist
1. Update and secure the VPS (`apt update && apt upgrade` + firewall).
2. Install Python build tools: `sudo apt install python3-pip python3-venv -y`.
3. Clone the repo to `/opt/hasiiChatbot` (or your preferred path).
4. Create a virtualenv (`python3 -m venv venv && source venv/bin/activate`).
5. Install requirements and configure `.env`.
6. Run `bash start` inside a process manager (tmux, screen, or systemd).
7. Monitor logs for FloodWait messages or Mongo connectivity issues.

## Support
- Owner: [@Hasindu_Lakshan](https://t.me/Hasindu_Lakshan)
- Source Code: this GitHub repository
- Support Group / Updates: *Coming soon*
