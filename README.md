# ucp-cli

`ucp-cli` can send SMs using Vodafone SMSCs.

#### setup
- go 1.11
- git

#### installation
```
go get github.com/jaakoob/ucp-cli
```

### config file

config.json:
```
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
