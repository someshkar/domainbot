import json
import whois
import validators

with open('./tlds.json') as f:
    tlds = json.loads(f.read())


def domain_status(domain, expiry=True):
    domain = domain.lower()

    if not validators.domain(domain):
        return "Please enter a valid domain!"

    tld = domain.split('.', 1)[1]

    if tld not in tlds['supported']:
        return "This TLD isn't currently supported!"

    domain_whois = whois.query(domain)

    if domain_whois is None:
        return '{} may be available!'.format(domain)
    else:
        if not expiry:
            args = [domain, domain_whois.registrar]
            return '{} is registered at {}'.format(*args)
        args = [domain, domain_whois.registrar,
                domain_whois.expiration_date.strftime("%B %d, %Y")]
        return '{} is registered at {} and will expire on {}.'.format(*args)


def domain_registered(domain):
    domain = domain.lower()
    domain_whois = whois.query(domain)
    if domain_whois is None:
        return False
    else:
        return True


def featured_domains(s):
    s = s.lower()

    available = []
    message = ' may be available!'

    for tld in tlds['featured']:
        domain = s + '.' + tld
        if not domain_registered(domain):
            available.append(domain)

    if len(available) > 1:
        args = [', '.join(available[:-1]), available[-1]]
        return '{} and {} may be available!'.format(*args)
    elif len(available) == 1:
        return '{} may be available!'.format(available[0])
    else:
        return 'None of the common TLDs are available for this.'
