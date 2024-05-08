# Bunny.net module for Caddy

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Bunny.net accounts.

## Examples

```
tls {
  acme_dns bunny {
    access_key {env.BUNNY_API_KEY}
    zone {env.BUNNY_ZONE}
  }

  propagation_timeout -1
}
```

```
{env.APP_DOMAIN} {
  tls {
    dns bunny {
      access_key {env.BUNNY_API_KEY}
      zone {env.BUNNY_ZONE}
    }

    propagation_timeout -1
  }
}
```

Replace `{env.BUNNY_API_KEY}` or `{env.BUNNY_ZONE}` with actual values if preferred. Not providing a `zone` will make Bunny.net look for a zone named `_acme-challenge.{env.APP_DOMAIN}`. Similarly, `propagation_timeout -1` is required.

Remember to update your zone when declaring `dns_challenge_override_domain` because of a `CNAME` record.

## Authenticating

To authenticate you need to provide a Bunny.net [API Key](https://dash.bunny.net/account/settings).
