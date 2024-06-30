### Dependencies

- make
- go
- npm
- installed go binaries to be in $PATH (for templ, air)

### To install

`make install`

### To build

`cp .env-example .env`
`make build`

### For watch / hot reloading

`make watch`

### For VSCode

Install `HTMX Attributes` for htmx autocompletion

Install `Tailwind CSS IntelliSense` for tailwindcss completion
AND
Add this to your settings.json
```
"tailwindCSS.includeLanguages": {
    "templ": "html",
},
```