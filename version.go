package main

import (
	"fmt"
)

var (
	Version            = "dev"
	Commit             = "none"
	RepoUrl            = "unknown"
	BuildDate          = "unknown"
	BuiltBy            = "unknown"
	BuiltWithGoVersion = "unknown"
)

func buildVersion() string {
	var result = fmt.Sprintf("Version: \t%s", Version)
	result = fmt.Sprintf("%s\nBuild Time: \t%s", result, BuildDate)
	result = fmt.Sprintf("%s\nGo Version: \t%s", result, BuiltWithGoVersion)
	result = fmt.Sprintf("%s\nRepo URL: \t%s", result, RepoUrl)
	result = fmt.Sprintf("%s\nCommit Info: \t%s", result, Commit)
	result = fmt.Sprintf("%s\nBuild By: \t%s", result, BuiltBy)
	return result
}
