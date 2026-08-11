package inspect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildItemPath(t *testing.T) {

	testCases := []struct {
		name     string
		segments []string
		expected string
	}{
		{
			name:     "/",
			segments: []string{"/"},
			expected: "/",
		},
		{
			name:     "//",
			segments: []string{"//"},
			expected: "/",
		},
		{
			name:     "empty",
			segments: []string{""},
			expected: "/",
		},
		{
			name:     "no input",
			segments: []string{},
			expected: "",
		},
		{
			name:     "//ip",
			segments: []string{"//ip"},
			expected: "/ip",
		},
		{
			name:     "/,/ip",
			segments: []string{"/", "/ip"},
			expected: "/ip",
		},
		{
			name:     "/ip/",
			segments: []string{"//ip/"},
			expected: "/ip",
		},
		{
			name:     "/,/,/",
			segments: []string{"/", "/", "/"},
			expected: "/",
		},
		{
			name:     "/,/,/ip/",
			segments: []string{"/", "/", "/ip/"},
			expected: "/ip",
		},
		{
			name:     "/,//ip/,dns/static/",
			segments: []string{"/", "//ip/", "dns/static"},
			expected: "/ip/dns/static",
		},
		{
			name:     "ip,dns,static",
			segments: []string{"ip", "dns", "static"},
			expected: "/ip/dns/static",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := buildItemPath(tc.segments...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
