# OpenTelemetry instrumentations for Go

| Instrumentation Package   | Metrics            | Traces             |
| ------------------------- | ------------------ | ------------------ |
| [database/sql](/otelsql/) | :heavy_check_mark: | :heavy_check_mark: |
| [GORM](/otelgorm/)        | :heavy_check_mark: | :heavy_check_mark: |
| [logrus](/otellogrus/)    |                    | :heavy_check_mark: |
| [Zap](/otelzap/)          |                    | :heavy_check_mark: |

## Contributing

To simiplify maintenance, we use a single version and a shared [changelog](CHANGELOG.md) for all
instrumentations. The changelog is auto-generated from
[conventional commits](https://www.conventionalcommits.org/en/v1.0.0/).

If you want to contribute an instrumentation, please make sure to include tests and a runnable
example. Use Docker if you must, but try to avoid it, for example, you can use SQLite instead of
MySQL to test database/sql instrumentation. Use [otellogrus](/otellogrus/) instrumentation as a
template.
