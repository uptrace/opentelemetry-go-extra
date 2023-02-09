## [0.1.21](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.20...v0.1.21) (2023-02-09)



## [0.1.20](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.19...v0.1.20) (2023-02-01)



## [0.1.19](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.18...v0.1.19) (2023-01-31)



## [0.1.18](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.17...v0.1.18) (2023-01-11)


### Bug Fixes

* **otelgorm:** restore original context after "after" callback ([7cd8508](https://github.com/uptrace/opentelemetry-go-extra/commit/7cd850876ad1dfaf4ee7efd5b3d226d75060d7c9))
* otelsql attr passthrough when reporting db stats ([6172588](https://github.com/uptrace/opentelemetry-go-extra/commit/617258842465e7f6bf085f5f8c53bd0cd864f474))


### Features

* **otelgorm:** ignore dry run spans by default ([323b5a9](https://github.com/uptrace/opentelemetry-go-extra/commit/323b5a99e338125fc77216a274bfd09cb932fb3a))
* **otelgorm:** optionally ignore dryrun callbacks ([5239b0d](https://github.com/uptrace/opentelemetry-go-extra/commit/5239b0d90875322fa57a07338a8ac47e74be50e2))



## [0.1.17](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.16...v0.1.17) (2022-10-07)



## [0.1.16](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.15...v0.1.16) (2022-09-05)


### Bug Fixes

* change SugaredLoggerWithCtx.Desuger to return LoggerWithCtx ([12f061e](https://github.com/uptrace/opentelemetry-go-extra/commit/12f061e380d1e2a383207614b139b03e51e946bf))
* copy skipCaller in With. Fixes [#62](https://github.com/uptrace/opentelemetry-go-extra/issues/62) ([4b58eac](https://github.com/uptrace/opentelemetry-go-extra/commit/4b58eac802c1cd9f564db06a7f832f7e417455c7))
* **otelzap:** do not panic on invalid input ([#69](https://github.com/uptrace/opentelemetry-go-extra/issues/69)) ([a41206e](https://github.com/uptrace/opentelemetry-go-extra/commit/a41206ed2c49fc7c817f470f61c5082b124cf267))


### Features

* **otelzap:** add extra fields option ([aa0e1a0](https://github.com/uptrace/opentelemetry-go-extra/commit/aa0e1a011e1dd8340143e95d1ab2e99300e7dc0b))
* **otelzap:** make the caller depth configurable ([38a6a03](https://github.com/uptrace/opentelemetry-go-extra/commit/38a6a03114bcaf15e53c5965eaf03b9be8ff0631))



## [0.1.15](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.14...v0.1.15) (2022-07-25)


### Bug Fixes

* **otelzap:** add missing DebugwContext function ([ea2b33a](https://github.com/uptrace/opentelemetry-go-extra/commit/ea2b33aa1e368a61b65fa31f4cf31ac794cd06ba))
* **otelzap:** fix WithTraceIDField for Debugw-like API ([c937b65](https://github.com/uptrace/opentelemetry-go-extra/commit/c937b653c20110be254e2232a4fc22d3de700c94))



## [0.1.14](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.13...v0.1.14) (2022-05-28)


### Bug Fixes

* import error ([8838b79](https://github.com/uptrace/opentelemetry-go-extra/commit/8838b79f5a8096bc7713925498531c26d6757fab))
* set span kind to client for otelgorm ([0250bbd](https://github.com/uptrace/opentelemetry-go-extra/commit/0250bbd1c85b89b8b1ade546f05d1db2a693ab34))



## [0.1.13](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.12...v0.1.13) (2022-05-08)


### Features

* upgrade to opentelemetry-go v1.7.0 ([bfa8f4e](https://github.com/uptrace/opentelemetry-go-extra/commit/bfa8f4ea4d83ca45da4429d17069e2ca531140f5))



## [0.1.12](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.11...v0.1.12) (2022-04-12)


### Bug Fixes

* **otelgorm:** ignore gorm.ErrRecordNotFound and other errors from otelsql/otel.go ([#48](https://github.com/uptrace/opentelemetry-go-extra/issues/48)) ([1c5d1f7](https://github.com/uptrace/opentelemetry-go-extra/commit/1c5d1f712afede4daf0dbdecbc48b6c245fbda39))


### Features

* **otelgorm:** add query formatter option ([e2d9787](https://github.com/uptrace/opentelemetry-go-extra/commit/e2d97873222a38c867c9d1342a3e60df588c98aa))
* **otelsql:** add an option: WithQueryFormatter ([b7a9f06](https://github.com/uptrace/opentelemetry-go-extra/commit/b7a9f0695fd18ec2e81eb0668bd694a5647a1dd9))



## [0.1.11](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.10...v0.1.11) (2022-03-29)



## [0.1.10](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.9...v0.1.10) (2022-03-16)


### Bug Fixes

* **otelzap:** print correct caller in InfoContext ([bcc0fa9](https://github.com/uptrace/opentelemetry-go-extra/commit/bcc0fa9898947d973b1a8a82218cdce9c751eb01))


### Features

* **otelgorm:** added an option to not report DB stats metrics ([9e52a0f](https://github.com/uptrace/opentelemetry-go-extra/commit/9e52a0f9b54cad994b24f221494445685415be40))



## [0.1.9](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.8...v0.1.9) (2022-02-27)

- **otelgorm**: added an option to otelgorm to exclude query variables

## [0.1.8](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.7...v0.1.8) (2022-01-28)

### Bug Fixes

- **otelsql:** implement driver.NamedValueChecker to support pgx
  ([c47a5de](https://github.com/uptrace/opentelemetry-go-extra/commit/c47a5de9a98df03d1c5575f5fe9a60c1eedac25a))
- **otelzap:** withoptions(zap.fields(...)) will keep fields with logger
  ([5e91392](https://github.com/uptrace/opentelemetry-go-extra/commit/5e91392104ad59e612bac1da80f5fe65debd5a3f))

### Features

- **otelzap:** add LoggerWithCtx.Sugar
  ([4792401](https://github.com/uptrace/opentelemetry-go-extra/commit/479240184b44f36f1623c6a3c5426e5ff0468c25))

## [0.1.7](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.6...v0.1.7) (2021-12-13)

### Bug Fixes

- **otelzap:** skip caller frame in ctx-aware API
  ([495c2e5](https://github.com/uptrace/opentelemetry-go-extra/commit/495c2e50d14e8a046b0e18624d9609b10885baf5))

### Features

- add otelzap.Ctx shortcut
  ([2d3c044](https://github.com/uptrace/opentelemetry-go-extra/commit/2d3c044adc7b624b596aaa4cdc3a566505fa4b91))

## [0.1.6](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.5...v0.1.6) (2021-11-25)

### Bug Fixes

- ctx fields not propegating to logger
  ([1520b8c](https://github.com/uptrace/opentelemetry-go-extra/commit/1520b8c4ab1a79539c91b274b25394d7b4cebb0c))

## [0.1.5](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.4...v0.1.5) (2021-11-17)

- Update to OpenTelemetry v1.2.0

## [0.1.4](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.3...v0.1.4) (2021-11-05)

- Update OpenTelemetry

## [0.1.3](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.2...v0.1.3) (2021-10-31)

### Features

- otelgraphql-go instrumentation ([#9](https://github.com/uptrace/opentelemetry-go-extra/issues/9))
  ([5cf626d](https://github.com/uptrace/opentelemetry-go-extra/commit/5cf626db67dd1e6f5c90b786259ea0a9091d08d3))

## [0.1.2](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.1...v0.1.2) (2021-10-26)

### Bug Fixes

- update Go module import for otelsqlx package
  ([2d517f7](https://github.com/uptrace/opentelemetry-go-extra/commit/2d517f7c01dcd5a6166e2ef4049ec983ec512c75))

## [0.1.1](https://github.com/uptrace/opentelemetry-go-extra/compare/v0.1.0...v0.1.1) (2021-10-18)

### Features

- add sqlx instrumentation
  ([bf92fbe](https://github.com/uptrace/opentelemetry-go-extra/commit/bf92fbe5873a96dd86ec5cc682758c1cc9303aba))
- **otelzap:** add missing globals
  ([a511657](https://github.com/uptrace/opentelemetry-go-extra/commit/a5116579029afd7b7f9d42125ce0abc12b93264d))

# 0.1.0 (2021-10-15)

### Features

- add otelgorm instrumentation
  ([d7a4276](https://github.com/uptrace/opentelemetry-go-extra/commit/d7a4276dd7de25cb1256828bd1c142ea61f3f1e1))
- add otellogrus instrumentation
  ([c86e9dd](https://github.com/uptrace/opentelemetry-go-extra/commit/c86e9dd73da4df87013d4241c0682c058ce89b4f))
- add otelsql instrumentation
  ([68269f9](https://github.com/uptrace/opentelemetry-go-extra/commit/68269f9c88cbdde75175526974eee10f1f03aa7b))
- add otelzap instrumentation
  ([f0691fa](https://github.com/uptrace/opentelemetry-go-extra/commit/f0691fa8573cb44691ddddfa00e32141bfa15095))
