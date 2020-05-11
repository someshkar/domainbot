# Discord Domainbot

A bot that lets you check domain availability while brainstorming names for an organization or event on Discord.

## Adding to your server

Go to [this](https://discord.com/api/oauth2/authorize?client_id=709320758475751495&permissions=18432&scope=bot) link to add the bot to your server. You need to have `manage` permissions on your Discord server to be able to add it.

### Usage

Simply type `domain example.com` in a channel with domainbot in it to check if it's available! If it isn't available, it'll let you know which registrar it was registered at.

## Self-hosting

This bot is currently hosted on AWS, but if you want to host it youself, make sure you've installed the `whois` package on your linux machine with `sudo apt install whois`.

Copy the `.env.example` file to `.env` and populate the environment variables.

After that, simply run the following commands to get it up and running:

```console
pip install -U discord.py whois python-dotenv
python3 bot.py
```

Issues and Pull Requests welcome :)
