module github.com/ONSdigital/dp-frontend-renderer

go 1.17

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/ONSdigital/dp-frontend-models v1.12.2
	github.com/ONSdigital/dp-healthcheck v1.1.0
	github.com/ONSdigital/go-ns v0.0.0-20210831102424-ebdecc20fe9e
	github.com/ONSdigital/log.go/v2 v2.0.9
	github.com/c2h5oh/datasize v0.0.0-20200825124411-48ed595a09d2
	github.com/gorilla/mux v1.8.0
	github.com/gosimple/slug v1.9.0
	github.com/jteeuwen/go-bindata v3.0.7+incompatible // indirect
	github.com/justinas/alice v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nicksnyder/go-i18n/v2 v2.1.2
	github.com/smartystreets/goconvey v1.6.4
	github.com/unrolled/render v1.0.3
	golang.org/x/text v0.3.7
	gopkg.in/russross/blackfriday.v2 v2.1.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/fatih/color v1.12.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20210202160940-bed99a852dfe // indirect
	github.com/hokaccha/go-prettyjson v0.0.0-20210113012101-fb4e108d2519 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/smartystreets/assertions v1.2.0 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)

// This dependency fails if used directly. We need to do this replace. More info: https://github.com/russross/blackfriday/issues/500
replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
