⚠️ This repository will be archived in December 2023 as it is no longer in development. ⚠️

# dp-frontend-renderer

## Deprecation

The renderer is being deprecated, the template and toml (localisation) files are to be migrated to the relevant controller/service.

Refer to the [migration](MIGRATION.md) readme for further information.

### Configuration

| Environment variable          | Default               | Description
| ----------------------------- | --------------------- | -------------------------------
| BIND_ADDR                     | :20010                | The host and port to bind to
| DEBUG                         | false                 | Enable debug mode
| SITE_DOMAIN                   | ons.gov.uk            | Domain for the website
| PATTERN_LIBRARY_ASSETS_PATH   |                       | Pattern library location
| SUPPORTED_LANGUAGES           | [2]string{"en", "cy"} | languages supported
| GRACEFUL_SHUTDOWN_TIMEOUT     | 5s                    | The graceful shutdown timeout
|	HEALTHCHECK_INTERVAL          | 30s                   | Time between calls to healthchecks
|	HEALTHCHECK_CRITICAL_TIMEOUT  | 90s                   | Timeout to consider a failing healthcheck critical

### Adding new page types

1. Create the page model in [dp-frontend-models](https://github.com/ONSdigital/dp-frontend-models)
  - you'll need to update the vendored copy - `govendor update github.com/ONSdigital/dp-frontend-models`
2. Create a new handler - use [handlers/homepage/handler.go](handlers/homepage/handler.go) as an example
  - `render.Handler` will take care of setting global variables, e.g. language
  - the `func()` passed in should be used for any page-specific customisations
3. Create a route in [main.go](main.go)

### NOTE: Code development / testing

Vscode in helpers.go and service.go does not see assets.Asset and assets.AssetNames so debugging won't play ball, and this is expected because the command line make commands utilise go generate to create the required functions in  data.go  ...

So - at the command line, make debug, etc are OK. 

### Licence

Copyright ©‎ 2016, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
