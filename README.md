# go-utilities

Shared Go packages for personal projects.

## collector

Package `collector` provides utilities for formatting [InfluxDB line protocol](https://docs.influxdata.com/influxdb/v2/reference/syntax/line-protocol/) and writing it to stdout. It is designed for use with Telegraf's [`inputs.exec`](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/exec) and [`inputs.execd`](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/execd) plugins.

### Usage

```go
import "github.com/iwvelando/go-utilities/collector"

points := []collector.Point{
    collector.NewPoint(
        "measurement_name",
        map[string]string{"host": "server01"},
        map[string]interface{}{"value": 42, "active": true},
        time.Now(),
    ),
}
collector.WritePoints(points)
```

### API

- `NewPoint(measurement, tags, fields, timestamp)` - creates a `Point`
- `FormatLineProtocol(point)` - formats a single `Point` as a line protocol string
- `WritePoints(points)` - writes all points to stdout with sync

### Consumers

- [forecast-solar-collector](https://github.com/iwvelando/forecast-solar-collector) (exec)
- [daylight-timeseries](https://github.com/iwvelando/daylight-timeseries) (exec)
- [ecobee_influx_connector](https://github.com/iwvelando/ecobee_influx_connector) (execd)
- [sleepnumber-stats-collector](https://github.com/iwvelando/sleepnumber-stats-collector) (execd)
- [airgradient-influxdb](https://github.com/iwvelando/airgradient-influxdb) (execd)