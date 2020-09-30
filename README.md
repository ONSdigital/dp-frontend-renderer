dp-frontend-renderer
====================

### Configuration

| Environment variable          | Default               | Description
| ----------------------------- | --------------------- | -------------------------------
| BIND_ADDR                     | :20000                | The host and port to bind to
| DEBUG                         | false                 | Enable debug mode
| SITE_DOMAIN                   | ons.gov.uk            | Domain for the website
| PATTERN_LIBRARY_ASSETS_PATH   |                       | Pattern library location
| SUPPORTED_LANGUAGES           | [2]string{"en", "cy"} | languages supported
| GRACEFUL_SHUTDOWN_TIMEOUT     | 5s                    | The graceful shutdown timeout
|	HEALTHCHECK_INTERVAL          | 30s                   | Time between calls to healthchecks
|	HEALTHCHECK_CRITICAL_TIMEOUT  | 90s                   | Timeout to consider a failing healthcheck critical

### Adding new page types

1. Create the page model in [https://github.com/ONSdigital/dp-frontend-models](dp-frontend-models)
  - you'll need to update the vendored copy - `govendor update github.com/ONSdigital/dp-frontend-models`
2. Create a new handler - use [handlers/homepage/handler.go](handlers/homepage/handler.go) as an example
  - `render.Handler` will take care of setting global variables, e.g. language
  - the `func()` passed in should be used for any page-specific customisations
3. Create a route in [main.go](main.go)


### Licence

Copyright ©‎ 2016, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
