### Dependencies

- go
- go binaries templ and air to be in $PATH (Found in $HOME/go/bin on my linux machine)
- make
- npm
- docker

### To install

`make install`

### To build

`cp .env-example .env`

Update .env as necessary

`make build`

### To run

`docker-compose up keycloak`
and then
`make run`

### For watch / hot reloading

`make watch`

### To configure keycloak realm etc.

Keycloak should use the backup templ.json found in ./keycloak/, but if you need to recreate it:

- Use this guide for setting up a new realm/client/users/etc. `https://www.keycloak.org/getting-started/getting-started-docker`
- Once setup to your satisfaction, then `docker exec -it keycloak /bin/bash` and then inside the container `./opt/keycloak/bin/kc.sh export --file /opt/keycloak/data/import/templ-realm.json --realm templ-realm`
- The new backup file should be in the ./keycloak directory

### For VSCode

Install `Go` plugin

Install `templ-vscode` plugin

Install `HTMX Attributes` for htmx autocompletion

Install `Tailwind CSS IntelliSense` for tailwindcss completion
AND
Add this to your settings.json
```
"tailwindCSS.includeLanguages": {
    "templ": "html",
},
```