package slog

import (
	"context"
	"fmt"
	slog "log/slog"
	"slices"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/internal/global"
	"go.opentelemetry.io/otel/logbridge"
)

type handler struct {
	logger logbridge.Logger
	attrs  []attribute.KeyValue
}

func NewHandler(opts ...Options) slog.Handler {
	h := &handler{}
	for _, opt := range opts {
		opt.apply(h)
	}
	return h
}

type Options interface {
	apply(*handler)
}
type optionFunc func(*handler)

func (o optionFunc) apply(h *handler) { o(h) }

func WithLogger(l logbridge.Logger) Options {
	return optionFunc(func(h *handler) {
		h.logger = l
	})
}
func WithAttrs(attrs []attribute.KeyValue) Options {
	return optionFunc(func(h *handler) {
		h.attrs = attrs
	})
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
// It is called early, before any arguments are processed,
// to save effort if the log event should be discarded.
// If called from a Logger method, the first argument is the context
// passed to that method, or context.Background() if nil was passed
// or the method does not take a context.
// The context is passed so Enabled can use its values
// to make a decision.
func (h *handler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

// Handle handles the Record.
// It will only be called when Enabled returns true.
// The Context argument is as for Enabled.
// It is present solely to provide Handlers access to the context's values.
// Canceling the context should not affect record processing.
// (Among other things, log messages may be necessary to debug a
// cancellation-related problem.)
//
// Handle methods that produce output should observe the following rules:
//   - If r.Time is the zero time, ignore the time.
//   - If r.PC is zero, ignore it.
//   - Attr's values should be resolved.
//   - If an Attr's key and value are both the zero value, ignore the Attr.
//     This can be tested with attr.Equal(Attr{}).
//   - If a group's key is empty, inline the group's Attrs.
//   - If a group has no Attrs (even if it has a non-empty key),
//     ignore it.
func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	lvl := convertLevel(r.Level)

	attrs := slices.Clone(h.attrs)
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, convertAttr(a))
		return true
	})

	h.logger.Emit(ctx, &logbridge.Record{
		Timestamp:         r.Time,
		ObservedTimestamp: time.Now(),
		Severity:          lvl,
		SeverityText:      lvl.String(),
		Body:              r.Message,
		Attributes:        attrs,
	})

	return nil
}

func convertLevel(l slog.Level) logbridge.Severity {
	return logbridge.Severity(l + 9)
}

func convertAttr(attr slog.Attr) attribute.KeyValue {
	val := convertValue(attr.Value, false)
	return attribute.KeyValue{Key: attribute.Key(attr.Key), Value: val}
}

func convertValue(v slog.Value, recurse bool) attribute.Value {
	switch v.Kind() {
	case slog.KindAny:
		return attribute.StringValue(fmt.Sprintf("%+v", v.Any()))
	case slog.KindBool:
		return attribute.BoolValue(v.Bool())
	case slog.KindDuration:
		return attribute.Int64Value(v.Duration().Nanoseconds())
	case slog.KindFloat64:
		return attribute.Float64Value(v.Float64())
	case slog.KindInt64:
		return attribute.Int64Value(v.Int64())
	case slog.KindString:
		return attribute.StringValue(v.String())
	case slog.KindTime:
		return attribute.Int64Value(v.Time().UnixNano())
	case slog.KindUint64:
		return attribute.Int64Value(int64(v.Uint64()))
	case slog.KindLogValuer:
		if !recurse {
			return convertValue(v.LogValuer().LogValue(), true)
		}
	}
	global.Warn("unknown attribute kind", "kind", v.Kind())
	return attribute.Value{}
}

// WithAttrs returns a new Handler whose attributes consist of
// both the receiver's attributes and the arguments.
// The Handler owns the slice: it may retain, modify or discard it.
func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	next := &handler{
		logger: h.logger,
		attrs:  slices.Clone(h.attrs),
	}
	next.attrs = slices.Grow(next.attrs, len(attrs))
	for _, a := range attrs {
		next.attrs = append(next.attrs, convertAttr(a))
	}
	return next
}

// WithGroup returns a new Handler with the given group appended to
// the receiver's existing groups.
// The keys of all subsequent attributes, whether added by With or in a
// Record, should be qualified by the sequence of group names.
//
// How this qualification happens is up to the Handler, so long as
// this Handler's attribute keys differ from those of another Handler
// with a different sequence of group names.
//
// A Handler should treat WithGroup as starting a Group of Attrs that ends
// at the end of the log event. That is,
//
//	logger.WithGroup("s").LogAttrs(level, msg, slog.Int("a", 1), slog.Int("b", 2))
//
// should behave like
//
//	logger.LogAttrs(level, msg, slog.Group("s", slog.Int("a", 1), slog.Int("b", 2)))
//
// If the name is empty, WithGroup returns the receiver.
func (f *handler) WithGroup(name string) slog.Handler {
	panic("not implemented") // TODO: Implement
}
