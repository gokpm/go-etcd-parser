package parser

import (
	"strings"

	"go.etcd.io/etcd/api/v3/mvccpb"
)

// Parse converts etcd-style keys into a nested map structure
func Parse(kvPairs []*mvccpb.KeyValue) map[string]interface{} {
	result := make(map[string]interface{})
	for _, kv := range kvPairs {
		parseNested(result, string(kv.Key), string(kv.Value))
	}
	return result
}

// parseNested sets a value in a nested map structure based on a slash-separated key
func parseNested(m map[string]interface{}, key string, value string) {
	parts := strings.Split(key, "/")
	current := m
	// Navigate/create nested maps for all parts except the last one
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		// If the key doesn't exist or isn't a map, create a new map
		if _, exists := current[part]; !exists {
			current[part] = make(map[string]interface{})
		} else if _, isMap := current[part].(map[string]interface{}); !isMap {
			// If it exists but isn't a map, replace it with a map
			current[part] = make(map[string]interface{})
		}
		// Move to the next level
		current = current[part].(map[string]interface{})
	}
	// Set the final value
	if len(parts) > 0 {
		current[parts[len(parts)-1]] = value
	}
}
