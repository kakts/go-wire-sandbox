# go-wire-sandbox
A sandbox project for [google/wire](https://github.com/google/wire)
/


# How to develop
## update wire injector
- update wire.go
- `make wire` or `make generate` to re-generate wire_gen.go
- check wire_gen.go

## Build this app
```bash
make build
```

`main` file to be created.

execute main
```bash
./main

```

## Test app
```bash
make test
```

If you want to test with verbose mode
```bash
make testv
```
