package logger

import (
	"github.com/fatih/color"
)

// Define color styles globally
var (
	Info    = color.New(color.FgCyan).Add(color.Bold)
	Success = color.New(color.FgGreen)
	Warning = color.New(color.FgYellow).Add(color.Bold)
	Error   = color.New(color.FgRed).Add(color.Bold)
)

// PrintInfo prints informational messages in cyan
func PrintInfo(format string, args ...interface{}) {
	Info.Printf(format, args...)
}

// PrintSuccess prints success messages in green
func PrintSuccess(format string, args ...interface{}) {
	Success.Printf(format, args...)
}

// PrintWarning prints warning messages in yellow
func PrintWarning(format string, args ...interface{}) {
	Warning.Printf(format, args...)
}

// PrintError prints error messages in red
func PrintError(format string, args ...interface{}) {
	Error.Printf(format, args...)
}
