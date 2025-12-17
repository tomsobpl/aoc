package main

// These values are intended to be overridden at link time via -ldflags
// Example: -X 'main.Version=1.2.3' -X 'main.Commit=abc123' -X 'main.Date=2025-12-17T00:00:00Z'
var (
	Version = "dev"
	Commit  = ""
	Date    = ""
)
