import json
import whois

with open('./supported_tlds.json') as f:
    supported_tlds = json.loads(f.read())


def domain_status(domain):
    tld = domain.split('.')[1]

    if tld not in supported_tlds:
        return "This TLD isn't currently supported!"

    domain_whois = whois.query(domain)

    if domain_whois is None:
        return '{} may be available!'.format(domain)
    else:
        args = [domain, domain_whois.registrar]
        return '{} is registered at {}'.format(*args)
