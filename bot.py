import discord
import os
import random
import whois
import json
import validators
from dotenv import load_dotenv

from lib import domain_status, featured_domains

load_dotenv()
TOKEN = os.getenv('DISCORD_TOKEN')

client = discord.Client()


@client.event
async def on_ready():
    print('{} is connected to the following guild(s):\n'.format(client.user))
    for guild in client.guilds:
        print('{}(id: {})'.format(guild.name, guild.id))
    print('\n')


@client.event
async def on_message(message):
    if message.author == client.user:
        return

    if message.content.startswith('domain'):
        domain = message.content.split(' ')[1]

        if ' all ' in message.content:
            domain = message.content.split(' ')[2]
            status = featured_domains(domain)

            print('"{}" returned "{}"'.format(domain, status))
            await message.channel.send(status)
            return

        status = domain_status(domain)

        print('"{}" returned "{}"'.format(domain, status))
        await message.channel.send(status)
        return

client.run(TOKEN)
