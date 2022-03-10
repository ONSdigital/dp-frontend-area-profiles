# dp-frontend-area-profiles
MVC app for Geo Journey on ons web


### Getting started

* Run `make debug`

For the frontend build run
* Run `make `
* Run `make `
### Dependencies

* No further dependencies other than those defined in `go.mod`

### Configuration

| Environment variable         | Default   | Description
| ---------------------------- | --------- | -----------
| BIND_ADDR                    | :26600    | The host and port to bind to
| GRACEFUL_SHUTDOWN_TIMEOUT    | 5s        | The graceful shutdown timeout in seconds (`time.Duration` format)
| HEALTHCHECK_INTERVAL         | 30s       | Time between self-healthchecks (`time.Duration` format)
| HEALTHCHECK_CRITICAL_TIMEOUT | 90s       | Time to wait until an unhealthy dependent propagates its state to make this app unhealthy (`time.Duration` format)
| HELLO_WORLD_EMPHASISE        | true      | Example boolean flag to control whether the 'Hello World' greeting should be emphasised with "!"

### Frontend
```
.
├── dist
├── node_modules
├── package-lock.json
├── package.json
public
├── @types
│   └── index.d.ts
├── __test__
│   └── area-landing.spec.ts
├── sass
│   ├── area-landing.scss
│   └── index.scss
└── ts
    ├── area-landing.ts
    └── geography-start.ts
├── tsconfig.json
└── webpack
```

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.

