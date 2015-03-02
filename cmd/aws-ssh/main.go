// Copyright © 2015 Christin R. Vozar

package main

import (
	// Standard Library
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	// Internal Packages
	"github.com/christianvozar/aws-ssh/ssh"

	// External Packages
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/ec2"
)

// globalConfig contains global service configuration values.
type globalConfig struct {
	accessKey string
	secretKey string
	region    string
	tag       string
	user      string
}

// Config is the global application configuration state.
var Config globalConfig

func instanceEntryAlias(i ec2.Instance) string {
	for _, t := range i.Tags {
		if t.Key == Config.tag {
			return t.Value
		}
	}
	// No entry for tag is found, return instance ID
	return i.InstanceId
}

func determineAmazonCredentials(ak, sk string) {
	// Set credentials and region if present in ENV.
	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		switch kv[0] {
		case "AWS_ACCESS_KEY":
			Config.accessKey = kv[1]
		case "AWS_SECRET_KEY":
			Config.secretKey = kv[1]
		case "AWS_DEFAULT_REGION":
			Config.region = kv[1]
		}
	}

	if ak != "" {
		Config.accessKey = ak
	}

	if sk != "" {
		Config.secretKey = sk
	}

	if Config.accessKey == "" {
		fmt.Println("Amazon Web Service access key undefined.")
		os.Exit(1)
	}

	if Config.secretKey == "" {
		fmt.Println("Amazon Web Service secret key undefined.")
		os.Exit(1)
	}
}

func main() {
	var entries []ssh.ConnectionEntry

	version := flag.Bool("version", false, "Display version")
	private := flag.Bool("private", false, "Use private IP addresses")
	accessKey := flag.String("access-key", "", "Amazon Web Services API access key")
	secretKey := flag.String("secret-key", "", "Amazon Web Services API secret key")
	tag := flag.String("tag", "Name", "EC2 tag to parse for name")
	user := flag.String("user", "", "Default username if undetermined")
	strict := flag.Bool("strict", true, "SSH strict host key checking")
	flag.Parse()

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	determineAmazonCredentials(*accessKey, *secretKey)
	Config.tag = *tag

	// Connection entry template
	t := template.Must(template.New("entry").Parse(ssh.Entry))

	auth := aws.Auth{AccessKey: Config.accessKey, SecretKey: Config.secretKey}

	e := ec2.New(auth, aws.USEast)
	instances, _ := e.DescribeInstances(nil, nil)
	for _, i := range instances.Reservations {
		for _, j := range i.Instances {
			if j.State.Name == "running" {
				if *private {
					entries = append(entries, ssh.ConnectionEntry{
						Host: instanceEntryAlias(j),
						IP:   j.PrivateIPAddress,
						User: *user,
						StrictHostKeyChecking: *strict,
					})
				} else {
					entries = append(entries, ssh.ConnectionEntry{
						Host: instanceEntryAlias(j),
						IP:   j.IPAddress,
						User: *user,
						StrictHostKeyChecking: *strict,
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
