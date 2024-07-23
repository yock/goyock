---
title: No-CSP
publishDate: "2020-12-06"
path: no-csp
---

In November 2020 it [came to light][tech-republic] that Apple's latest release of macOS, Big Sur, included a feature which sent identifying information to Apple each time you open an application on your computer. Online Certificate Status Protocol (OCSP) has been billed as a security feature that Apple claims to use to identify and protect users from malware. This not withstanding it could easily be used to gather analytics on its users which might be valuable to both benign and malicious organizations. The debate on whether or not this is reasonable for Apple to do will continue well into the future but I wanted to be able to block this on my home network with a high level of confidence. This is how I did it.

First I must give a nod to my friend [Adam Simpson][adam-simpson] ([@a_simpson][adam-simpson-twitter]). It was him who first let me know about OCSP and it was with his guidance that I landed on a plan of action. That plan was to run a DNS server on my local network on a separate non-Apple device. That way I could resolve the OCSP subdomain (ocsp.apple.com) to a non-routable address. No amount of trickery on the Apple devices can work around this short of hard-coding the IP address in the OS.

There seemed to be one obvious choice for hardware. The Raspbery Pi 4 is overpowered for the specific task of resolving domain names for a handful of consumer devices at home but it's ultimately affordable and popular enough to ensure software support for the foreseeable future. [Unbound][unbound] serves as the DNS software running on the Pi. Configuring Unbound isn't complicated but there were a few bumps I had to iron out.

## Configuration

```
# Unbound configuration
server:
	verbosity: 0
	interface: 0.0.0.0

	access-control: 127.0.0.0/8 allow
	access-control: 192.168.1.0/24 allow
	do-ip4: yes
	do-tcp: yes
	do-udp: yes

	root-hints: "/var/lib/unbound/root.hints"

	# Block Apple OCSP
	local-zone: "ocsp.apple.com" redirect
	local-data: "ocsp.apple.com A 0.0.0.0"

	# Netgear Orbi administration
	local-data: "orbilogin.com A 192.168.1.1"
```

The first thing to call attention to here is the `root-hints` directive. In this application Unbound needs to know how to find authoritative DNS servers for any hosts not configured locally and this file contains the location of those root dns servers. The content for this file comes care of the [IANA][root-hints-iana].

The next two lines are where the money is. They ensure that `ocsp.apple.com` resolves to the non-routable address `0.0.0.0`. I'll be the first to admit that I'm not a DNS expert and the specifics underlying DNS and this terminology largely elude me. That said my understanding is that the first line ensures that the local Unbound instance is authoritative for the configured name. That means Unbound will never forward requests to resolve this domain to another DNS, which is what we want. The second line creates the `A` record to resolve the name to the IP address we want.

The last line won't be relevant to most people but it adds convenience for my network. My Netgear Orbi system resolves this domain to the address of the local router. Now that I'm bypassing the DNS features of my router I added this line so I could keep using the domain instead of the IP address of the router.

## The Future

DNS is a popular choice for whole network ad-blocking and there are sources on the web for scripts that will build Unbound configuration files that define these same kinds of rules for common ad hosts. One drawback to this approach is the difficulty in pausing ad-blocking or selectively allowing certain hosts. Not limited to DNS tricks the Raspberry Pi makes an excellent platform for software-defined radio. If you're into aviation they're a common platform for ADSB receivers and some flight tracking networks offer perks for feeding them data.

[tech-republic]: https://www.techrepublic.com/article/security-experts-level-criticism-at-apple-after-big-sur-launch-issues/
[adam-simpson]: https://adamsimpson.net/
[adam-simpson-twitter]: https://twitter.com/a_simpson
[unbound]: https://nlnetlabs.nl/projects/unbound/about/
[root-hints-iana]: https://www.iana.org/domains/root/files
