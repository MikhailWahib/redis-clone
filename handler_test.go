// handler_test.go
package main

import (
	"reflect"
	"testing"
)

func TestPing(t *testing.T) {
	tests := []struct {
		name     string
		args     []Value
		expected Value
	}{
		{
			name:     "No arguments",
			args:     []Value{},
			expected: Value{typ: "string", str: "PONG"},
		},
		{
			name:     "With argument",
			args:     []Value{{typ: "bulk", bulk: "hello"}},
			expected: Value{typ: "string", str: "hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ping(tt.args)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestSetAndGet(t *testing.T) {
	// Clear the SETs map before testing
	SETs = make(map[string]string)

	setArgs := []Value{
		{typ: "bulk", bulk: "mykey"},
		{typ: "bulk", bulk: "myvalue"},
	}
	setResult := set(setArgs)
	expected := Value{typ: "string", str: "OK"}
	if !reflect.DeepEqual(setResult, expected) {
		t.Errorf("SET: Expected %v, got %v", expected, setResult)
	}

	getArgs := []Value{{typ: "bulk", bulk: "mykey"}}
	getResult := get(getArgs)
	expected = Value{typ: "bulk", bulk: "myvalue"}
	if !reflect.DeepEqual(getResult, expected) {
		t.Errorf("GET: Expected %v, got %v", expected, getResult)
	}

	// Test GET with non-existent key
	getArgs = []Value{{typ: "bulk", bulk: "nonexistent"}}
	getResult = get(getArgs)
	expected = Value{typ: "null"}
	if !reflect.DeepEqual(getResult, expected) {
		t.Errorf("GET (non-existent): Expected %v, got %v", expected, getResult)
	}
}

func TestDel(t *testing.T) {
	// Clear and populate the SETs map before testing
	SETs = make(map[string]string)
	SETs["key1"] = "value1"
	SETs["key2"] = "value2"

	args := []Value{
		{typ: "bulk", bulk: "key1"},
		{typ: "bulk", bulk: "key2"},
		{typ: "bulk", bulk: "nonexistent"},
	}
	result := del(args)
	expected := Value{typ: "integer", str: "2"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Verify keys were deleted
	if _, exists := SETs["key1"]; exists {
		t.Errorf("key1 should have been deleted")
	}
	if _, exists := SETs["key2"]; exists {
		t.Errorf("key2 should have been deleted")
	}
}

func TestExists(t *testing.T) {
	// Clear and populate the SETs map before testing
	SETs = make(map[string]string)
	SETs["key1"] = "value1"
	SETs["key2"] = "value2"

	args := []Value{
		{typ: "bulk", bulk: "key1"},
		{typ: "bulk", bulk: "nonexistent"},
		{typ: "bulk", bulk: "key2"},
	}
	result := exists(args)
	expected := Value{typ: "integer", str: "2"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestHSet(t *testing.T) {
	// Clear the HSETs map before testing
	HSETs = make(map[string]map[string]string)

	tests := []struct {
		name     string
		args     []Value
		expected Value
	}{
		{
			name: "Set new hash field",
			args: []Value{
				{typ: "bulk", bulk: "myhash"},
				{typ: "bulk", bulk: "field1"},
				{typ: "bulk", bulk: "value1"},
			},
			expected: Value{typ: "string", str: "OK"},
		},
		{
			name: "Update existing hash field",
			args: []Value{
				{typ: "bulk", bulk: "myhash"},
				{typ: "bulk", bulk: "field1"},
				{typ: "bulk", bulk: "newvalue1"},
			},
			expected: Value{typ: "string", str: "OK"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hset(tt.args)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}

	// Verify the hash was updated correctly
	if HSETs["myhash"]["field1"] != "newvalue1" {
		t.Errorf("Expected myhash[field1] to be 'newvalue1', got '%s'", HSETs["myhash"]["field1"])
	}
}

func TestHGet(t *testing.T) {
	// Set up test data
	HSETs = map[string]map[string]string{
		"myhash": {
			"field1": "value1",
			"field2": "value2",
		},
	}

	tests := []struct {
		name     string
		args     []Value
		expected Value
	}{
		{
			name: "Get existing hash field",
			args: []Value{
				{typ: "bulk", bulk: "myhash"},
				{typ: "bulk", bulk: "field1"},
			},
			expected: Value{typ: "bulk", bulk: "value1"},
		},
		{
			name: "Get non-existent hash field",
			args: []Value{
				{typ: "bulk", bulk: "myhash"},
				{typ: "bulk", bulk: "nonexistent"},
			},
			expected: Value{typ: "null"},
		},
		{
			name: "Get field from non-existent hash",
			args: []Value{
				{typ: "bulk", bulk: "nonexistenthash"},
				{typ: "bulk", bulk: "field1"},
			},
			expected: Value{typ: "null"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hget(tt.args)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestHGetAll(t *testing.T) {
	// Set up test data
	HSETs = map[string]map[string]string{
		"myhash": {
			"field1": "value1",
			"field2": "value2",
		},
		"emptyhash": {},
	}

	tests := []struct {
		name     string
		args     []Value
		expected Value
	}{
		{
			name: "Get all fields from existing hash",
			args: []Value{
				{typ: "bulk", bulk: "myhash"},
			},
			expected: Value{
				typ: "array",
				array: []Value{
					{typ: "bulk", bulk: "field1"},
					{typ: "bulk", bulk: "value1"},
					{typ: "bulk", bulk: "field2"},
					{typ: "bulk", bulk: "value2"},
				},
			},
		},
		{
			name: "Get all fields from empty hash",
			args: []Value{
				{typ: "bulk", bulk: "emptyhash"},
			},
			expected: Value{typ: "array", array: []Value{}},
		},
		{
			name: "Get all fields from non-existent hash",
			args: []Value{
				{typ: "bulk", bulk: "nonexistenthash"},
			},
			expected: Value{typ: "null"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hgetall(tt.args)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
