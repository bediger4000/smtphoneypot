# SMTP Honeypot

I want to learn more about [SMTP](https://datatracker.ietf.org/doc/html/rfc2821), and I want to take
revenge on spammers for ruining email.

A time consuming SMTP honeypot.

## Building

```
$ cd $GOPATH/src
$ git clone https://github.com/bediger4000/smtphoneypot.git
$ cd smtphoneypot
$ go build $PWD
```

## Usage

```
Usage of smtphoneypot:
  -a string
        hostname:port form address on which to listen (default "localhost:8976")
  -d    debug output
  -hn string
        host name to present to senders (default "hazard")
  -l string
        log directory (default "./log")
```
