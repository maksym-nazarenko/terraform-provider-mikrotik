package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSmoke_InterfaceWiFiConfiguration(t *testing.T) {
	c := NewClient(GetConfigFromEnv())

	expectedResource := &InterfaceWiFiConfiguration{
		Comment:  "A comment",
		Disabled: true,
		Name:     "sample",
		Ssid:     "sample-ssid",
	}

	createdResource, err := c.AddInterfaceWiFiConfiguration(expectedResource)
	require.NoError(t, err)

	defer func() {
		err := c.Delete(createdResource)
		if !IsNotFoundError(err) {
			assert.NoError(t, err)
		}
	}()
	assert.NotEmpty(t, createdResource.Id)

	foundResource, err := c.Find(expectedResource)
	require.NoError(t, err)
	assert.Equal(t, createdResource, foundResource)

	// cleanup
	err = c.Delete(foundResource)
	assert.NoError(t, err)

	_, err = c.Find(expectedResource)
	assert.True(t, IsNotFoundError(err), "expected not found error, got: %v", err)
}
