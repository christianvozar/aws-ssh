// Copyright © 2015-2017 Christian R. Vozar <christian@rogueethic.com>
// Fabriqué en Nouvelle Orléans ⚜

package ssh

// Define a connection entry.
const Entry = `
Host {{.Host}}
    HostName {{.IP}}
    User {{.User}}{{if .StrictHostKeyChecking}}
    StrictHostKeyChecking yes{{else}}
    StrictHostKeyChecking no{{end}}
`
