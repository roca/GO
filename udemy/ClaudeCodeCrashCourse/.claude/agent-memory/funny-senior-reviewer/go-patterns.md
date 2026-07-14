# Go Patterns & Anti-Patterns Observed

## Review 1: fibonacci/main.go (2026-03-12)
- **Good:** Iterative Fibonacci (O(n) time, O(1) space) instead of naive recursion
- **Good:** Idiomatic parallel assignment `a, b = b, a+b`
- **Issue:** No input validation for negative numbers
- **Issue:** No error return (Go convention for invalid inputs)
- **Issue:** Integer overflow not addressed (int64 overflows at F(93))
- **Issue:** Magic number hardcoded (20) without const/var
- **Issue:** No test file present
- **Score:** 7/10
