import json
import whois
import validators

with open('./supported_tlds.json') as f:
    supported_tlds = json.loads(f.read())


def domain_status(domain, expiry=True):
    if not validators.domain(domain):
        return "Please enter a valid domain!"

    tld = domain.split('.', 1)[1]

    if tld not in supported_tlds:
        return "This TLD isn't currently supported!"

    domain_whois = whois.query(domain)

    if domain_whois is None:
        return '{} may be available!'.format(domain)
    else:
        if expiry:
            args = [domain, domain_whois.registrar,
                    domain_whois.expiration_date.strftime("%B %d, %Y")]
            return '{} is registered at {} and will expire on {}.'.format(*args)
        args = [domain, domain_whois.registrar]
        return '{} is registered at {}'.format(*args)
