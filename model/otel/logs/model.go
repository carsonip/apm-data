package logs

import (
	"fmt"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	// timestampFormat formats timestamps according to Elasticsearch's
	// strict_date_optional_time date format, which includes a fractional
	// seconds component.
	timestampFormat = "2006-01-02T15:04:05.000Z07:00"
)

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format(timestampFormat))
	return []byte(stamp), nil
}

type LogRecord struct {
	TimeUnixNano           JSONTime       `json:"@timestamp"`
	ObservedTimeUnixNano   JSONTime       `json:"observed_timestamp"`
	SeverityNumber         int64          `json:"severity_number"`
	SeverityText           string         `json:"severity_text"`
	BodyText               string         `json:"body_text"`
	BodyStructured         map[string]any `json:"body_structured"`
	Attributes             map[string]any `json:"attributes"`
	DroppedAttributesCount int64          `json:"dropped_attributes_count"`
	Flags                  uint8          `json:"trace_flags"`
	TraceId                string         `json:"trace_id"`
	SpanId                 string         `json:"span_id"`
	Scope                  Scope          `json:"scope"`
	Resource               Resource       `json:"resource"`

	DataStream DataStream `json:"data_stream"`
}

type Scope struct {
	ScopeName              string         `json:"name"`
	ScopeVersion           string         `json:"version"`
	ScopeAttributes        map[string]any `json:"attributes"`
	DroppedAttributesCount int64          `json:"dropped_attributes_count"`
	SchemaUrl              string         `json:"schema_url"`
}

type Resource struct {
	ResourceAttributes             map[string]any `json:"attributes"`
	ResourceDroppedAttributesCount int64          `json:"dropped_attributes_count"`
	SchemaUrl                      string         `json:"schema_url"`
}

type DataStream struct {
	Type      string `json:"type,omitempty"`
	Dataset   string `json:"dataset,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

func toLogRecord(resource pcommon.Resource, scope pcommon.InstrumentationScope, log plog.LogRecord) LogRecord {
	r := Resource{
		ResourceAttributes:             resource.Attributes().AsRaw(),
		ResourceDroppedAttributesCount: int64(resource.DroppedAttributesCount()),
		SchemaUrl:                      "",
	}
	s := Scope{
		ScopeName:              scope.Name(),
		ScopeVersion:           scope.Version(),
		ScopeAttributes:        resource.Attributes().AsRaw(),
		DroppedAttributesCount: int64(scope.DroppedAttributesCount()),
		SchemaUrl:              "",
	}
	var bodyStructured map[string]any
	if log.Body().Type() == pcommon.ValueTypeMap {
		bodyStructured = log.Body().Map().AsRaw()
	}
	l := LogRecord{
		TimeUnixNano:           JSONTime{log.Timestamp().AsTime()},
		ObservedTimeUnixNano:   JSONTime{log.ObservedTimestamp().AsTime()},
		SeverityNumber:         int64(log.SeverityNumber()),
		SeverityText:           log.SeverityText(),
		BodyText:               log.Body().AsString(),
		BodyStructured:         bodyStructured,
		Attributes:             log.Attributes().AsRaw(),
		DroppedAttributesCount: int64(log.DroppedAttributesCount()),
		Flags:                  uint8(log.Flags()),
		TraceId:                log.TraceID().String(),
		SpanId:                 log.SpanID().String(),
		Scope:                  s,
		Resource:               r,
	}
	return l
}

func Flatten(logs plog.Logs) (res []LogRecord) {
	res = make([]LogRecord, 0, logs.LogRecordCount())

	resourceLogs := logs.ResourceLogs()
	for i := 0; i < resourceLogs.Len(); i++ {
		resourceLog := resourceLogs.At(i)
		resource := resourceLog.Resource()
		scopeLogs := resourceLog.ScopeLogs()
		for j := 0; j < scopeLogs.Len(); j++ {
			scopeLog := scopeLogs.At(j)
			scope := scopeLog.Scope()
			logRecords := scopeLog.LogRecords()
			for k := 0; k < logRecords.Len(); k++ {
				res = append(res, toLogRecord(resource, scope, logRecords.At(k)))
			}
		}
	}
	return res
}
