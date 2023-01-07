/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package version

import (
	"fmt"

	goVersion "github.com/hashicorp/go-version"
)

// ProtocolVersion marks the protobuf interface version
// // Update this if breaking changes are made to the protobuf spec (for example column types)
var ProtocolVersion int64 = 20220201

// Version is the main version number that is being run at the moment.
var version = "5.1.0"

// A pre-release marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
var prerelease = "rc.0"

// semVer is an instance of version.Version. This has the secondary
// benefit of verifying during tests and init time that our version is a
// proper semantic version, which should always be the case.
var semVer *goVersion.Version

func init() {
	versionString := version
	if prerelease != "" {
		versionString = fmt.Sprintf("%s-%s", version, prerelease)
	}
	semVer = goVersion.Must(goVersion.NewVersion(versionString))
}

// String returns the complete version string, including prerelease
func String() string {
	return semVer.String()
}
