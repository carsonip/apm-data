package input

// ConsumeLogsResult contains the number of rejected log records and error message for partial success response.
type ConsumeLogsResult struct {
	ErrorMessage       string
	RejectedLogRecords int64
}
