// Copyright © 2015 Christin R. Vozar

package ssh

// Define a connection entry.
const Entry = `
Host {{.Host}}
    HostName {{.IP}}
    User {{.User}}{{if .StrictHostKeyChecking}}
    StrictHostKeyChecking yes{{else}}
    StrictHostKeyChecking no{{end}}
`
