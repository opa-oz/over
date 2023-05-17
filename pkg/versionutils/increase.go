package versionutils

import (
	"fmt"

	"github.com/hashicorp/go-version"
)

func Increase(v *version.Version, patch, minor, major bool) (*version.Version, error) {
	var majorNum int
	var minorNum int
	var patchNum int

	segments := v.Segments()
	if len(segments) > 0 {
		majorNum = segments[0]
	}
	if len(segments) > 1 {
		minorNum = segments[1]
	}
	if len(segments) > 2 {
		patchNum = segments[2]
	}

	if patch {
		patchNum += 1
	} else if minor {
		patchNum = 0
		minorNum += 1
	} else if major {
		patchNum = 0
		minorNum = 0
		majorNum += 1
	}

	return version.NewVersion(fmt.Sprintf("%d.%d.%d", majorNum, minorNum, patchNum))
}
