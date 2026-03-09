package collector

import (
	"testing"
	"time"
)

func TestFormatLineProtocol_Basic(t *testing.T) {
	ts := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	p := NewPoint(
		"cpu",
		map[string]string{"host": "server01", "region": "us-west"},
		map[string]interface{}{"value": 0.64},
		ts,
	)
	got := FormatLineProtocol(p)
	want := `cpu,host=server01,region=us-west value=0.64 1705320000000000000`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatLineProtocol_IntegerFields(t *testing.T) {
	ts := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	p := NewPoint(
		"test",
		map[string]string{},
		map[string]interface{}{"count": 42, "small": int8(1)},
		ts,
	)
	got := FormatLineProtocol(p)
	want := `test count=42i,small=1i 1705320000000000000`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatLineProtocol_BooleanFields(t *testing.T) {
	ts := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	p := NewPoint(
		"test",
		map[string]string{},
		map[string]interface{}{"active": true, "paused": false},
		ts,
	)
	got := FormatLineProtocol(p)
	want := `test active=t,paused=f 1705320000000000000`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatLineProtocol_StringFields(t *testing.T) {
	ts := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	p := NewPoint(
		"test",
		map[string]string{},
		map[string]interface{}{"message": `hello "world"`},
		ts,
	)
	got := FormatLineProtocol(p)
	want := `test message="hello \"world\"" 1705320000000000000`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatLineProtocol_EscapedTagsAndMeasurement(t *testing.T) {
	ts := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	p := NewPoint(
		"my measurement",
		map[string]string{"tag key": "tag,value"},
		map[string]interface{}{"val": 1},
		ts,
	)
	got := FormatLineProtocol(p)
	want := `my\ measurement,tag\ key=tag\,value val=1i 1705320000000000000`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatLineProtocol_NoTimestamp(t *testing.T) {
	p := NewPoint(
		"test",
		map[string]string{},
		map[string]interface{}{"value": 1.0},
		time.Time{},
	)
	got := FormatLineProtocol(p)
	want := `test value=1`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}

func TestFormatLineProtocol_NoTags(t *testing.T) {
	ts := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	p := NewPoint(
		"test",
		map[string]string{},
		map[string]interface{}{"value": 3.14},
		ts,
	)
	got := FormatLineProtocol(p)
	want := `test value=3.14 1705320000000000000`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}
