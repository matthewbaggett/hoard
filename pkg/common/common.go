package common

import _ "embed"

//go:embed VERSION
var version string

//go:embed BUILD_TIME
var buildTime string

//go:embed GIT_COMMIT_MESSAGE
var gitCommitMessage string

func GetVersion() string {
	return version
}

func GetBuildTime() string {
	return buildTime
}

func GetGitCommitMessage() string {
	return gitCommitMessage
}
