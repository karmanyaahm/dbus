# NoProvider2Push D-Bus

**Warning**: Only use this for development. It's not production-ready.

Config:
This typically goes into `~/.config/unifiedpush/distributors/np2p.conf`. The following are some possible values you can fill in (don't put in multiple of the same key into one config file).
```ini
proxyurl = direct
# for testing or if your computer is publicly exposed to a static IP for some reason (still not recommended because np2p doesn't support tls(https))
# or
proxyurl = https://mynp2p.proxy.tld

port = 30043
# defaults to this so no need to fill in unless you want to change it

IP = 192.168.0.99
# ipv4
IP = 2001:0DB8::123
# your ipv6 address

# depends on your proxy setup
# this defaults to a Yggdrasil IP address if you're running that in the background
IP = 201:be::0123
```

Run this with:
```sh
git clone https://github.com/NoProvider2Push/dbus.git
cd dbus
go run .
```


Roadmap: 
- alpha: currently
- beta: once builds are set up
- stable: v1.0 should be released once dbus UP platform is proven stable - don't know timeline

## Library

The distributor package can be used as a module in your own distributor if you wish. Other parts like config and storage can also be copied with the appropriate license. NP2P is usually the 'example distributor' in UnifiedPush due to its simplicity.

```sh
go get -u unifiedpush.org/go/np2p_dbus
```
