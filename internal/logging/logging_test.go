package logging

import (
	"net/http"
	"testing"
)

func TestLoggingNewLoggingTransport_IsRoundTripper(t *testing.T) {
	transport := NewLoggingTransport("example", &http.Transport{})
	var _ http.RoundTripper = transport
}
