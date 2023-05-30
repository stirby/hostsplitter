# HostSplitter
HostSplitter is an HTTP reverse proxy and load balancer that distributes requests to an arbitrary amount of sites based on the Host header.


## Motivation
I commonly run into an issue developing small golang websites: I want to use the same IP address for many sites that aren't large enough to justify their own VPS.

## Site files
HostSplitter will look for site files by default in "/etc/hostsplitter/". HostSplitter will only read files with the .json extension.

A each site file should look like
```json
{
	"hostnames": [
		"skirby.dev",
		"www.skirby.dev"
	],
	"backends": [
		"127.0.0.1:9000"
	],
	"secret": "puppies1234"
}
```

The "secret" field is passed along with every request to that site in the ``X-Hostsplitter-Secret`` header. This is intended to be checked before trusting the passed along IP.

## Real IP 
The original requester's IP is located in the ``X-Forwarded-For`` header.

## Reloading
HostSplitter provides 0 downtime reload functionality via SIGUSR1. E.g
```bash
pkill -10 hostsplitter

## Roadmap
- SSL
