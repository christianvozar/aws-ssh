// Copyright © 2015-2020 Christian R. Vozar <christian@rogueethic.com>

package ssh

// ConnectionEntry represents one host connection entry.
type ConnectionEntry struct {
	IP                    string
	Host                  string
	User                  string
	StrictHostKeyChecking bool
}
