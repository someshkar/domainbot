import discord
import os
import random
import whois
import json
from dotenv import load_dotenv

from lib import domain_status

load_dotenv()
TOKEN = os.getenv('DISCORD_TOKEN')
GUILD = os.getenv('DISCORD_GUILD')

client = discord.Client()


@client.event
async def on_ready():
    for guild in client.guilds:
        if guild.name == GUILD:
            break

    print(
        f'{client.user} is connected to the following guild(s):\n'
        f'{guild.name}(id: {guild.id})'
    )


@client.event
async def on_message(message):
    if message.author == client.user:
        return

    if message.content.startswith('domain'):
        domain = message.content.split(' ')[1]

        status = domain_status(domain)
        await message.channel.send(status)

client.run(TOKEN)
