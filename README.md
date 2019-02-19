# Packet Metabot

Packet Metabot! The robot for your Packet metadata!

Available as an executable binary or container, Metabot _knows_ your metadata and will retrieve whatever you want. Metabot's commands are expanding regularly.

Why not just use `curl | jq`? Sure, you could, but Metabot will extract all of that for you quickly and easily, including expanding and formatting, without any of the long url or, worse, complex `jq` formatting..

## Installation

Either download the release from this site, or get the latest docker image `docker pull packethost/metabot`.

## Running
Whether you run as a straight binary or container, the basics are the same: run it and ask it what you want:

```
$ metabot <query> <qualifier>
$ # or
$ docker run --rm packethost/metabot <query> <qualifier>
```

## Queries and Qualifiers
The qualifiers are unique to each query type. Some accept no qualifiers. If a query accepts multiple qualifiers, the order is unimportant; it just keeps applying them. Some qualifiers filter "up", i.e. change the results. See each type.

### Queries
Here are queries metabot understands today:

|query|output|example|
|---|---|---|
|id|ID of this server|b642678f-1d6e-45a2-aed1-bd0a63135fe5|
|hostname|hostname of this server|metabot-test|
|facility|facility ID in which this server exists|ewr1|
|ip|space-separated list of all IP addresses for this host in cidr format|10.0.10.10/31 205.206.207.122/31 2604:1380:1:9200::1/127|

### Qualifiers

For each query type, here are the acceptable qualifiers. If the qualifier is unknown, Metabot will return an error.

#### ip
Metabot query `ip` has multiple qualifiers.

* `address`: return only the address without the `/x` cidr part, e.g. `10.0.10.10 205.206.207.122 2604:1380:1:9200::1`.
* `4`: return only IPv4 addresses.
* `6`: return only IPv6 addresses.
* `private`: return only private IP addresses.
* `parent`: return the parent IP address range instead of the actual IP address on the host.
* `netmask`: return the address in `ip/netmask` format instead of `ip/cidr` format.

## Custom URL
By default, Metabot retrieves data from the Packet URL https://metadata.packet.com/metadata , but you can override it with the `--url <url>` option.

