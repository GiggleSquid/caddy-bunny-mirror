package bunny

import (
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	bunny "github.com/GiggleSquid/caddy-bunny-dns-mirror"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *bunny.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.bunny",
		New: func() caddy.Module { return &Provider{new(bunny.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.AccessKey = repl.ReplaceAll(p.Provider.AccessKey, "")
	p.Provider.Zone = repl.ReplaceAll(p.Provider.Zone, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	bunny {
//	    access_key <access_token>
//	    zone <zone>
//	}
//
// Expansion of placeholders in the API token is left to the JSON config caddy.Provisioner (above).
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "access_key":
				if d.NextArg() {
					p.Provider.AccessKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "zone":
				if d.NextArg() {
					p.Provider.Zone = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("Unrecognised subdirective '%s'", d.Val())
			}
		}
	}

	fmt.Println("Initialising with zone ", p.Provider.Zone)

	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
