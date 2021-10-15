# OpenTelemetry instrumentations for Go

| Instrumentation Package    | Metrics            | Traces             |
| -------------------------- | ------------------ | ------------------ |
| [otelsql](/otelsql/)       | :heavy_check_mark: | :heavy_check_mark: |
| [otelgorm](/otelgorm/)     | :heavy_check_mark: | :heavy_check_mark: |
| [otellogrus](/otellogrus/) |                    | :heavy_check_mark: |
| [otelzap](/otelzap/)       |                    | :heavy_check_mark: |

## Contributing

If you want to contribute an instrumentation, please make sure to include tests and a runnable
example. Use Docker if you must, but try to avoid it, for example, use SQLite instead of MySQL when
possible. Use [otellogrus](/otellogrus/) as a template.
