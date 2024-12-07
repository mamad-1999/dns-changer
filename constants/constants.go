package constants

// Config URL for downloading the default configuration.
const ConfigURL = "https://raw.githubusercontent.com/mamad-1999/dns-changer/refs/heads/master/config.json"

// Config directory and file names.
const (
	ConfigDir  = ".config/dns-changer"
	ConfigFile = "config.json"
	BackupDir  = ".config/dns-changer/backups" // New backup directory
	BackupFile = "resolv.conf.backup"          // Backup filename
)

// Error messages.
const (
	ErrFindingHomeDir     = "Error finding home directory: %s"
	ErrCreatingConfigDir  = "Error creating config directory: %s"
	ErrHandlingConfigFile = "Error handling config.json: %s"
	ErrParsingDnsConfig   = "Error parsing DNS config: %s"
	ErrReadingInput       = "Error reading input. Please try again."
	ErrInvalidInput       = "Invalid input. Please enter a valid number."
	ErrInvalidChoice      = "Invalid choice. Please select a number between 1 and %d, or 0 to exit."
	ErrWritingToResolv    = "Error writing to /etc/resolv.conf: %s"
	ErrDnsPingFailed      = "DNS ping failed: %s"
)

// Success messages.
const (
	SuccessDnsChanged = "Successfully changed DNS to %s"
	SuccessExit       = "Exiting the program."
)

// Prompt messages.
const (
	SelectDnsPrompt = "Select a DNS server by number (or 0 to exit): "
)
