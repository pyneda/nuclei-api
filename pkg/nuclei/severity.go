package nuclei

import (
	"errors"
	"github.com/projectdiscovery/nuclei/v2/pkg/model/types/severity"
	"strings"
)

var severityMappings = map[severity.Severity]string{
	severity.Info:     "info",
	severity.Low:      "low",
	severity.Medium:   "medium",
	severity.High:     "high",
	severity.Critical: "critical",
	severity.Unknown:  "unknown",
}

func normalizeValue(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func stringToSeverity(severityString string) (severity.Severity, error) {
	normalizedValue := normalizeValue(severityString)
	for key, currentValue := range severityMappings {
		if normalizedValue == currentValue {
			return key, nil
		}
	}
	return -1, errors.New("Invalid severity: " + severityString)
}
