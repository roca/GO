package web

import (
	"io"
	"mime"
	"net/http"
	"regexp"
)

const (
	injectionScanBudget = 64 * 1024 // 64KB

	maxJSONBodyBytes = 1 << 20 // 1MB

	maxMultipartBytes = 10 << 20 // 10MB
)

var injectionPatterns = []*regexp.Regexp{
	// "ignore (all|any|the) (previous|prior|above|earlier) (instructions|prompts|messages)"
	regexp.MustCompile(`(?i)\bignore\s+(?:all\s+|any\s+|the\s+|your\s+)?(?:previous|prior|above|earlier|prior\s+system)\s+(?:instructions?|prompts?|messages?|directives?|rules?)`),
	// "disregard ..."
	regexp.MustCompile(`(?i)\bdisregard\s+(?:all\s+|any\s+|the\s+|your\s+|previous\s+|prior\s+|above\s+|system\s+)+(?:instructions?|prompts?|messages?|directives?|rules?)`),
	// "forget everything ..."
	regexp.MustCompile(`(?i)\bforget\s+(?:all\s+|everything\s+|prior\s+|previous\s+|your\s+)+(?:instructions?|prompts?|messages?|context|rules?|above)`),
	// Role-takeover: "you are now X", "act as Y", "pretend to be Z"
	regexp.MustCompile(`(?i)\b(?:you\s+are\s+now|act\s+as|pretend\s+to\s+be|roleplay\s+as)\s+(?:dan|jailbreak|developer\s+mode|unrestricted|an?\s+ai\s+with\s+no\s+(?:restrictions|filters|rules))`),
	// "Do Anything Now" / "DAN mode" / "developer mode enabled"
	regexp.MustCompile(`(?i)\b(?:do\s+anything\s+now|dan\s+mode(?:\s+enabled)?|developer\s+mode\s+(?:enabled|activated|on))\b`),
	// Reveal-the-prompt attacks
	regexp.MustCompile(`(?i)\breveal\s+(?:your\s+|the\s+)?(?:system\s+)?(?:prompt|instructions|directives)`),
	regexp.MustCompile(`(?i)\b(?:what|show|tell\s+me)\s+(?:are|were|is|me)?\s*(?:your|the)?\s*(?:system\s+)?(?:prompt|instructions|directives)\b`),
	regexp.MustCompile(`(?i)\b(?:print|output|repeat|echo|show)\s+(?:everything\s+|all\s+(?:of\s+)?|me\s+|verbatim\s+)?(?:above|prior|previous|your\s+(?:system\s+)?(?:prompt|instructions))`),
	// Safety-bypass language
	regexp.MustCompile(`(?i)\b(?:override|bypass|circumvent|disable|turn\s+off)\s+(?:your\s+|the\s+|all\s+)?(?:safety|guardrails|restrictions|filters|content\s+policy|safeguards)`),
	// ChatML / role-tag smuggling. Local templates that honor these
	// tokens will treat the embedded role as authoritative.
	regexp.MustCompile(`(?i)<\s*\|?\s*im_start\s*\|?\s*>`),
	regexp.MustCompile(`(?i)<\s*\|?\s*(?:system|assistant|developer)\s*\|?\s*>`),
	regexp.MustCompile(`(?i)\[\s*(?:INST|SYSTEM|/SYSTEM)\s*\]`),
}

func scanForInjection(s string) string {
	if len(s) > injectionScanBudget {
		s = s[:injectionScanBudget]
	}

	for _, p := range injectionPatterns {
		if p.MatchString(s) {
			return p.String()
		}
	}

	return ""
}

func injectionDefense(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mt, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		switch mt {
		case "application/json":
		// inspectJSON
		case "multipart/form-data":
		// inspectMultipart
		default:
			next.ServeHTTP(w, r)
		}
	})
}

// TODO: must implement the following functions to inspect the request body for injection patterns
func inspectJSON(w http.ResponseWriter, r *http.Request, next http.Handler) {
}
