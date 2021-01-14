package interfaces



type StatusResponse struct {
	Active bool
	FailureRate float64
}

type Message struct {
	InwardFailure string `json:"todayFailureRateInward"`
	OutwardFailure string `json:"todayFailureRateOutward"`
}

type Result struct {
	Loaded bool `json:"loaded"`
	Msg Message `json:"msg"`
}

type ErrorResponse struct{
	Message string
}

type LogEntry struct {
	Method string
	URL string
	RequestHeaders string
}