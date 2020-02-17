// Copyright © 2015-2020 Christian R. Vozar <christian@rogueethic.com>

package main

import (
	// Standard Library
	"fmt"
	"log"
	"os"
	"text/template"

	// Internal Packages
	"github.com/christianvozar/aws-ssh/ssh"

	// External Packages
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/ec2"
)

func instanceEntryAlias(i ec2.Instance) string {
	for _, t := range i.Tags {
		if t.Key == config.tag {
			return t.Value
		}
	}
	// No entry for tag is found, return instance ID
	return i.InstanceId
}

func main() {
	var entries []ssh.ConnectionEntry

	if err := initConfig(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Connection entry template
	t := template.Must(template.New("entry").Parse(ssh.Entry))

	auth := aws.Auth{AccessKey: config.accessKey, SecretKey: config.secretKey}

	e := ec2.New(auth, aws.USEast)
	instances, _ := e.DescribeInstances(nil, nil)
	for _, i := range instances.Reservations {
		for _, j := range i.Instances {
			if j.State.Name == "running" {
				if config.private {
					entries = append(entries, ssh.ConnectionEntry{
						Host: instanceEntryAlias(j),
						IP:   j.PrivateIPAddress,
						User: config.user,
						StrictHostKeyChecking: config.strict,
					})
				} else {
					entries = append(entries, ssh.ConnectionEntry{
						Host: instanceEntryAlias(j),
						IP:   j.IPAddress,
						User: config.user,
						StrictHostKeyChecking: config.strict,
					})
				}
			}
		}
	}

	// Render entry for each instance
	for _, ent := range entries {
		err := t.Execute(os.Stdout, ent)
		if err != nil {
			log.Println("executing connection entry template:", err)
		}
	}
}
