package common

import _ "embed"

//go:embed VERSION
var version string

//go:embed BUILD_TIME
var build_time string

func GetVersion() string {
	return version
}

func GetBuildTime() string {
	return build_time
}
