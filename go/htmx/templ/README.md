### Dependencies

- make
- go
- npm
- go binaries templ and air to be in $PATH (Found in $HOME/go/bin on my linux machine)

### To install

`make install`

### To build

`cp .env-example .env`

Update .env as necessary

`make build`

### To configure keycloak realm etc.

Use this guide `https://www.keycloak.org/getting-started/getting-started-docker`

@MarkFix Save keycloak config to json and config it on starting the docker container

### For watch / hot reloading

`make watch`

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