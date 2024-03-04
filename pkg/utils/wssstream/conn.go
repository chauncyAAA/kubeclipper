package wssstream

import (
	"net/http"
	"regexp"
	"strings"
)

var (
	// connectionUpgradeRegex matches any Connection header value that includes upgrade
	connectionUpgradeRegex = regexp.MustCompile(`(^|.*,\s*)upgrade($|\s*,)`)
)

// IsWebSocketRequest returns true if the incoming request contains connection upgrade headers
// for WebSockets.
func IsWebSocketRequest(req *http.Request) bool {
	if !strings.EqualFold(req.Header.Get("Upgrade"), "websocket") {
		return false
	}
	return connectionUpgradeRegex.MatchString(strings.ToLower(req.Header.Get("Connection")))
}
