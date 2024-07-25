package main

import (
	"os"
	"reflect"
	"testing"
)

func TestAof(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "aof_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Create a new AOF instance
	aof, err := NewAof(tmpfile.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer aof.Close()

	// Test writing to AOF
	testValue := Value{
		typ: "array",
		array: []Value{
			{typ: "bulk", bulk: "SET"},
			{typ: "bulk", bulk: "mykey"},
			{typ: "bulk", bulk: "myvalue"},
		},
	}

	err = aof.Write(testValue)
	if err != nil {
		t.Fatalf("Error writing to AOF: %v", err)
	}

	// Test reading from AOF
	var readValue Value
	err = aof.Read(func(value Value) {
		readValue = value
	})
	if err != nil {
		t.Fatalf("Error reading from AOF: %v", err)
	}

	// Compare written and read values
	if !reflect.DeepEqual(testValue, readValue) {
		t.Errorf("Read value does not match written value. Expected %v, got %v", testValue, readValue)
	}
}
