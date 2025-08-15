package otelsql

import (
	"database/sql/driver"

	"go.opentelemetry.io/otel/trace"
)

type countableRows struct {
	rows driver.Rows
	span trace.Span

	count int
}

func newCountableRows(
	rows driver.Rows,
	span trace.Span,
) *countableRows {
	return &countableRows{rows: rows, span: span}
}

func (c *countableRows) Columns() []string {
	return c.rows.Columns()
}

func (c *countableRows) Close() error {
	defer c.span.End()

	err := c.rows.Close()
	c.span.AddEvent("Close")
	if err != nil {
		c.span.RecordError(err)
	}

	c.span.SetAttributes(dbRowsScanned.Int(c.count))

	return err
}

func (c *countableRows) Next(dest []driver.Value) error {
	err := c.rows.Next(dest)
	if err == nil {
		c.count++
	}
	return err
}
