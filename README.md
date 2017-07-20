# Test cipher in browser

## Prepare

Install Gopher.js https://github.com/gopherjs/gopherjs .

Install Go packages.

```
go get -u gopkg.in/macaron.v1
go get -u github.com/skycoin/skycoin/src/cipher
go get -u github.com/go-macaron/binding
go get -u github.com/go-macaron/gzip
```

## Build

Build Cipher

```
cd ./gopherjs
gopherjs build main.go
```

## Run

```
go run main.go
```

Open url in browser http://localhost:4000