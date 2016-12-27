// Copyright © 2015-2017 Christian R. Vozar <christian@rogueethic.com>
// Fabriqué en Nouvelle Orléans ⚜

package ssh

// ConnectionEntry represents one host connection entry.
type ConnectionEntry struct {
	IP                    string
	Host                  string
	User                  string
	StrictHostKeyChecking bool
}
