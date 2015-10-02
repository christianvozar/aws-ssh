// Copyright © 2015 Christin R. Vozar

package ssh

// ConnectionEntry represents one host connection entry.
type ConnectionEntry struct {
	IP                    string
	Host                  string
	User                  string
	StrictHostKeyChecking bool
}
