### TODO
- Better error handling
- Save new events to a database
- Add sig generation

### To run

First
```
cd path/to/nostr_backup
git submodule init
git submodule update
```

As a dev build:
```
go run .
```

As a prod build:
```
go run -tags prod .
```
