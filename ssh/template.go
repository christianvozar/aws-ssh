// Copyright © 2015-2020 Christian R. Vozar <christian@rogueethic.com>

package ssh

// Entry is a SSH connection entry.
const Entry = `
Host {{.Host}}
    HostName {{.IP}}
    User {{.User}}{{if .StrictHostKeyChecking}}
    StrictHostKeyChecking yes{{else}}
    StrictHostKeyChecking no{{end}}
`
