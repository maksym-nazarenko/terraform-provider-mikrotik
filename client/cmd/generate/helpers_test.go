package main

import (
	"testing"

	"github.com/ddelnano/terraform-provider-mikrotik/client/pkg/inspect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindSubNode(t *testing.T) {
	testCases := []struct {
		name      string
		basePath  string
		rootNode  *inspect.Node
		expected  *inspect.Node
		expectErr bool
	}{
		{
			name:     "path is root node",
			basePath: "/ip/dns/static",
			rootNode: &inspect.Node{
				Self: "/ip/dns/static",
			},
			expected: &inspect.Node{
				Self: "/ip/dns/static",
			},
		},
		{
			name:     "path leads to a child",
			basePath: "/ip/dns/static",
			rootNode: &inspect.Node{
				Self: "/ip",
				ChildrenMap: map[string]*inspect.Node{
					"dns": {
						Self: "/ip/dns",
						ChildrenMap: map[string]*inspect.Node{
							"cache": {
								Self: "/ip/dns/cache",
							},
							"static": {
								Self: "/ip/dns/static",
							},
						},
					},
				},
			},
			expected: &inspect.Node{
				Self: "/ip/dns/static",
			},
		},
		{
			name:     "path has different path prefix from root node",
			basePath: "/ip/dns/static",
			rootNode: &inspect.Node{
				Self: "/bridge/port",
			},
			expectErr: true,
		},
		{
			name:     "path leads to a missing child",
			basePath: "/ip/dns/static",
			rootNode: &inspect.Node{
				Self: "/ip",
				ChildrenMap: map[string]*inspect.Node{
					"dns": {
						Self: "/ip/dns",
						ChildrenMap: map[string]*inspect.Node{
							"cache": {
								Self: "/ip/dns/cache",
							},
							"adlist": {
								Self: "/ip/dns/adlist",
							},
						},
					},
				},
			},
			expectErr: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := findSubNode(tc.rootNode, tc.basePath)
			require.True(t, (err == nil) == (tc.expectErr == false), "expected error: %v, got: %v", tc.expectErr, err)
			if tc.expectErr {
				return
			}

			assert.Equal(t, tc.expected, actual)
		})
	}
}
