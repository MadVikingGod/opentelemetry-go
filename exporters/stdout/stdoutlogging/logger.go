package stdoutlogging

import (
	"context"
	"fmt"
	"io"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/logbridge"
)

type Exporter struct {
	writer io.Writer
	noTime bool
}

func NewExporter(writer io.Writer) *Exporter {
	return &Exporter{
		writer: writer,
	}
}

func NewExporterNoTime(writer io.Writer) *Exporter {
	return &Exporter{
		writer: writer,
		noTime: true,
	}
}

func (e *Exporter) Export(_ context.Context, records []*logbridge.Record) {
	for _, r := range records {
		if r == nil {
			continue
		}
		e.flattenRecord(r)
		e.flattenAttributes(r.Attributes)
		e.writer.Write([]byte("\n"))
	}
}

func (e *Exporter) flattenRecord(r *logbridge.Record) {
	if !e.noTime {
		fmt.Fprintf(e.writer, `%s %s "`, r.Timestamp, r.ObservedTimestamp)
	}
	fmt.Fprintf(e.writer, `%s body="%s"`, r.SeverityText, r.Body)
}

func (e *Exporter) flattenAttributes(attrs []attribute.KeyValue) {
	for _, a := range attrs {
		fmt.Fprintf(e.writer, ` %s="%s"`, a.Key, a.Value.Emit())
	}
}
