// Package collector provides utilities for formatting InfluxDB line protocol
// output, intended for use with telegraf exec/execd input plugins.
package collector

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Point represents a single InfluxDB line protocol data point.
type Point struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
	Timestamp   time.Time
}

// NewPoint creates a Point with the given measurement, tags, fields, and timestamp.
func NewPoint(measurement string, tags map[string]string, fields map[string]interface{}, ts time.Time) Point {
	return Point{
		Measurement: measurement,
		Tags:        tags,
		Fields:      fields,
		Timestamp:   ts,
	}
}

// FormatLineProtocol formats a Point as an InfluxDB line protocol string.
func FormatLineProtocol(p Point) string {
	var b strings.Builder

	b.WriteString(escapeMeasurement(p.Measurement))

	// Tags sorted by key for deterministic output
	if len(p.Tags) > 0 {
		keys := make([]string, 0, len(p.Tags))
		for k := range p.Tags {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			b.WriteByte(',')
			b.WriteString(escapeTagKeyValue(k))
			b.WriteByte('=')
			b.WriteString(escapeTagKeyValue(p.Tags[k]))
		}
	}

	b.WriteByte(' ')

	// Fields sorted by key for deterministic output
	if len(p.Fields) > 0 {
		keys := make([]string, 0, len(p.Fields))
		for k := range p.Fields {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			if i > 0 {
				b.WriteByte(',')
			}
			b.WriteString(escapeFieldKey(k))
			b.WriteByte('=')
			b.WriteString(formatFieldValue(p.Fields[k]))
		}
	}

	if !p.Timestamp.IsZero() {
		b.WriteByte(' ')
		b.WriteString(fmt.Sprintf("%d", p.Timestamp.UnixNano()))
	}

	return b.String()
}

func escapeMeasurement(s string) string {
	s = strings.ReplaceAll(s, ",", `\,`)
	s = strings.ReplaceAll(s, " ", `\ `)
	return s
}

func escapeTagKeyValue(s string) string {
	s = strings.ReplaceAll(s, ",", `\,`)
	s = strings.ReplaceAll(s, "=", `\=`)
	s = strings.ReplaceAll(s, " ", `\ `)
	return s
}

func escapeFieldKey(s string) string {
	return escapeTagKeyValue(s)
}

func escapeFieldStringValue(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	return s
}

func formatFieldValue(v interface{}) string {
	switch val := v.(type) {
	case int:
		return fmt.Sprintf("%di", val)
	case int8:
		return fmt.Sprintf("%di", val)
	case int16:
		return fmt.Sprintf("%di", val)
	case int32:
		return fmt.Sprintf("%di", val)
	case int64:
		return fmt.Sprintf("%di", val)
	case uint:
		return fmt.Sprintf("%du", val)
	case uint8:
		return fmt.Sprintf("%du", val)
	case uint16:
		return fmt.Sprintf("%du", val)
	case uint32:
		return fmt.Sprintf("%du", val)
	case uint64:
		return fmt.Sprintf("%du", val)
	case float32:
		return fmt.Sprintf("%g", val)
	case float64:
		return fmt.Sprintf("%g", val)
	case bool:
		if val {
			return "t"
		}
		return "f"
	case string:
		return fmt.Sprintf(`"%s"`, escapeFieldStringValue(val))
	default:
		return fmt.Sprintf(`"%s"`, escapeFieldStringValue(fmt.Sprintf("%v", val)))
	}
}
