package ssh

// ConnectionEntry represents one host connection entry.
type ConnectionEntry struct {
	IP                    string
	Host                  string
	User                  string
	StrictHostKeyChecking bool
}
