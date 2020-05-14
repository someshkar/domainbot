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


class Domainbot(discord.Client):
    async def on_ready(self):
        print('{} is connected to the following guild(s):\n'.format(self.user))
        for guild in client.guilds:
            print('{0.name} (id: {0.id})'.format(guild))
        print('\n')

    async def on_message(self, message):
        if message.author == self.user:
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


client = Domainbot()
client.run(TOKEN)
