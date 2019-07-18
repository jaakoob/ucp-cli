# ucp-cli

`ucp-cli` is a pure [Go](https://golang.org) implementation of the [UCP](https://wiki.wireshark.org/UCP) protocol primarily used to connect to short message service centres (SMSCs),  in order to send and receive short messages (SMS). This fork is adapted to work with the Vodafone SMSCs.

#### setup
- go 1.11
- git

#### installation
```
go get github.com/jaakoob/ucp-cli
```

### config file

```
config.json:
{
  "Address": "1.1.1.1",
  "Port": 5001,
  "Username": "foobar",
  "Password": "p4ssw0rd"
}
```

#### usage
```
./ucp-cli -config="config.json" -from=00491721234567 -to=00491721234568 -message="This is a test message."
```
