package schema

import (
	"encoding/json"
	"testing"

	"github.com/shiningrush/droplet"
	"github.com/stretchr/testify/assert"
)

func TestPlugin(t *testing.T) {
	// init
	handler := &Handler{}
	assert.NotNil(t, handler)

	// plugin list(old api, return name only)
	listInput := &ListInput{}
	ctx := droplet.NewContext()
	ctx.SetInput(listInput)
	list, err := handler.Plugins(ctx)
	assert.Nil(t, err)
	assert.Contains(t, list.([]string), "limit-count")

	// plugin list(return all fields of plugin)
	listInput = &ListInput{
		All: true,
	}
	ctx = droplet.NewContext()
	ctx.SetInput(listInput)
	list, err = handler.Plugins(ctx)
	assert.Nil(t, err)
	plugins := list.([]map[string]interface{})
	var authPlugins []string
	var basicAuthConsumerSchema string
	for _, plugin := range plugins {
		if plugin["type"] == "auth" {
			authPlugins = append(authPlugins, plugin["name"].(string))
		}
		if plugin["name"] == "basic-auth" {
			consumerSchemaByte, err := json.Marshal(plugin["consumer_schema"])
			basicAuthConsumerSchema = string(consumerSchemaByte)
			assert.Nil(t, err)
		}
	}
	// plugin type
	assert.ElementsMatch(t, []string{"basic-auth", "jwt-auth", "hmac-auth", "key-auth", "wolf-rbac", "ldap-auth"}, authPlugins)
	// consumer schema
	assert.Equal(t, `{"properties":{"password":{"type":"string"},"username":{"type":"string"}},"required":["password","username"],"title":"work with consumer object","type":"object"}`, basicAuthConsumerSchema)
}
