package web3go

import (
	"testing"
	"time"

	"github.com/mcuadros/go-defaults"
	"gotest.tools/assert"
)

func TestConfigurationDefault(t *testing.T) {
	c := ClientOption{}
	defaults.SetDefaults(&c)
	assert.Equal(t, c.RetryCount, 3)
	assert.Equal(t, c.RetryInterval, 1*time.Second)
	assert.Equal(t, c.RequestTimeout, 3*time.Second)

	c = ClientOption{RetryCount: 10}
	defaults.SetDefaults(&c)
	assert.Equal(t, c.RetryCount, 10)
	assert.Equal(t, c.RetryInterval, 1*time.Second)
	assert.Equal(t, c.RequestTimeout, 3*time.Second)

}
