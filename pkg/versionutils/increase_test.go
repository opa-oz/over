package versionutils_test

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/opa-oz/over/pkg/versionutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	In                  string
	Exp                 string
	Patch, Minor, Major bool
}

func TestIncrease(t *testing.T) {
	testTable := []TestCase{
		{
			In:    "1.0.0",
			Exp:   "1.0.1",
			Patch: true,
		},
		{
			In:    "1.0.0",
			Exp:   "1.1.0",
			Minor: true,
		},
		{
			In:    "1.0.0",
			Exp:   "2.0.0",
			Major: true,
		},
		{
			In:    "1.0.1",
			Exp:   "1.0.2",
			Patch: true,
		},
		{
			In:    "1.0.1",
			Exp:   "1.1.0",
			Minor: true,
		},
		{
			In:    "1.1.1",
			Exp:   "2.0.0",
			Major: true,
		},
	}

	for index, tcase := range testTable {
		t.Run(fmt.Sprintf("#%d - %s -> %s", index, tcase.In, tcase.Exp), func(t *testing.T) {
			parsed, err := version.NewSemver(tcase.In)
			assert.NoError(t, err)
			actual, err := versionutils.Increase(parsed, tcase.Patch, tcase.Minor, tcase.Major)

			assert.NoError(t, err)
			assert.Equal(t, tcase.Exp, actual.String())
		})
	}
}
