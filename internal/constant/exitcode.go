package constant

// ExitCode is a code returned to a parent process by a command
type ExitCode int

const (
	SuccessExitCode ExitCode = iota
	VulnerabilityFoundExitCode

	ErrorExitCode = ExitCode(99)
)
