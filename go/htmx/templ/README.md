### Dependencies

- go
- go binaries (templ, air, etc.) to be in $PATH after install (Found in $HOME/go/bin on my linux machine)
- make
- docker

### To install

`make install`

This will:
- Install all go binaries used
- Install all go dependencies
- Create ./database/database.db and run all migrations

### To build

You must install first. Then...

`cp .env-example .env`

Update .env as necessary

`make build`

Or to make a prod build

`make buildProd`

### To run

`make kc`

And then in separate terminal

`make run`

### For watch / hot reloading

`make watch`

That will not give you browser hot reloading with the templ templates.
For that, you should run

`make watchProxy`

and use http://127.0.0.1:7331

### To configure keycloak realm etc.

Keycloak should use the backup templ-realm.json found in ./keycloak/, but if you need to recreate it:

- Use this guide for setting up a new realm/client/users/etc. `https://www.keycloak.org/getting-started/getting-started-docker`
- Once setup to your satisfaction, then `docker exec -it keycloak /bin/bash` to get a prompt inside the keycloak container and then `./opt/keycloak/bin/kc.sh export --file /opt/keycloak/data/import/templ-realm.json --realm templ-realm`
- The new backup file will be in ./keycloak/. It will overwrite the existing ./keycloak/templ-realm.json file

The backup templ-realm.json contains an admin user/password admin/admin and a test user/password test/test

### To generate self-signed certs for TLS

You can generate a self-signed cert for your own machine using `mkcert`

```
mkcert -install
mkcert localhost
```

then add to .env `TLS_KEY="localhost-key.pem"` and `TLS_CRT="localhost.pem"`

### For VSCode

Install `Go` plugin

Install `templ-vscode` plugin

Install `HTMX Attributes` for htmx autocompletion