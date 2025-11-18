package helper

import (
	"strings"
	"testing"
)

func TestGetDefaultCapJsTmpl(t *testing.T) {
	tmpl := GetDefaultCapJsTmpl()

	// Verify it returns a non-empty string
	if tmpl == "" {
		t.Error("GetDefaultCapJsTmpl returned empty string")
	}

	// Verify it contains expected HTML elements
	expectedElements := []string{
		"<html>",
		"</html>",
		"<head>",
		"</head>",
		"<body>",
		"</body>",
		"<cap-widget",
		"</cap-widget>",
		"{{ .FrontendJS }}",
		"{{ .ChallengeURL }}",
	}

	for _, elem := range expectedElements {
		if !strings.Contains(tmpl, elem) {
			t.Errorf("Template missing expected element: %s", elem)
		}
	}

	// Verify it's valid HTML structure (basic check)
	if !strings.HasPrefix(tmpl, "<html>") {
		t.Error("Template should start with <html>")
	}
	if !strings.HasSuffix(strings.TrimSpace(tmpl), "</html>") {
		t.Error("Template should end with </html>")
	}
}
