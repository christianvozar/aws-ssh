// Copyright © 2015-2020 Christian R. Vozar <christian@rogueethic.com>

package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Config contains global service configuration values.
type Config struct {
	// accessKey is the Amazon Web Services API access key
	// Default: nil
	// Required: true
	accessKey string
	// secretKey is the Amazon Web Services API secret key.
	// Default: nil
	// Required: true
	secretKey string
	region    string
	// tag is the EC2 tag to parse for name.
	// Default: "Name"
	tag string
	// user is the default username if undetermined.
	user string
	// private specifies if private IP addresses should be used.
	// Default: false
	private bool
	// strict determines if SSH strict host key checking is enabled.
	// Default: false
	strict bool
}

var (
	config Config
)

// initConfig initializes the global configuration populating settings from
// defaults, then overriding from environment variables.
func initConfig() error {
	if config.strict == false {
		config.strict = true
	}

	processEnvs()

	if config.accessKey == "" {
		return errors.New("Amazon Web Service access key undefined.")
	}

	if config.secretKey == "" {
		return errors.New("Amazon Web Service secret key undefined.")
	}

	if config.tag == "" {
		config.tag = "Name"
	}

	return nil
}

// processEnvs iterates over environment variables, overriding default
// configuration settings.
func processEnvs() {
	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		switch kv[0] {
		case "AWS_ACCESS_KEY":
			config.accessKey = kv[1]
		case "AWS_SECRET_KEY":
			config.secretKey = kv[1]
		case "AWS_DEFAULT_REGION":
			config.region = kv[1]
		case "STRICT":
			config.strict, _ = strconv.ParseBool(kv[1])
		case "PRIVATE":
			config.private, _ = strconv.ParseBool(kv[1])
		case "USER":
			config.user = kv[1]
		case "TAG":
			config.tag = kv[1]
		}
	}
}
