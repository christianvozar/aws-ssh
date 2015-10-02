// Copyright © 2015 Christin R. Vozar

package main

var (
	// GitCommit is the commit during compile; filled in by the compiler.
	GitCommit string
	// GitDescribe is the commit description during compile; filled in by the compiler.
	GitDescribe string
)

const (
	// Version is the Semantic versioning standard for the appliation.
	// http://semver.org/
	Version = "1.0.2"
	// PreRelease of the version for the application.
	PreRelease = "dev"
)
