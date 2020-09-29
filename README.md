# Joule LND Proxy

Local proxy for a remote LND node.

This app allows you to proxy requests to a remote node through localhost.

This solves the invalid SSL issue when connecting the [Joule lightning](https://lightningjoule.com/) extension to a remote node.

The browsers complain about the self-signed LND certificate. This proxy allows you to connect joule to localhost through http and 
requests get proxied to the remote node using the correct certificate.

## Screenshots

#### Main Screen
![](https://github.com/bumi/lnd-proxy/raw/master/screenshot-main.png)

#### Runnig Proxy
![](https://github.com/bumi/lnd-proxy/raw/master/screenshot-running.png)


For more information see: https://github.com/joule-labs/joule-extension/issues/106


## Download

See [releases](https://github.com/bumi/lnd-proxy/releases)
