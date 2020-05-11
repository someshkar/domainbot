import discord
import os
import random
import whois
import json
from dotenv import load_dotenv

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
        f'{client.user} is connected to the following guild:\n'
        f'{guild.name}(id: {guild.id})'
    )


@client.event
async def on_message(message):
    if message.author == client.user:
        return

    if 'domain' in message.content:
        domain = message.content.split(' ')[1]
        domain_whois = whois.query(domain)
        if domain_whois is None:
            await message.channel.send(f'{domain} may be available!')
        else:
            await message.channel.send(f'{domain} is registered at {domain_whois.registrar}')

client.run(TOKEN)
