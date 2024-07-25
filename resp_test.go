package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestReadArray(t *testing.T) {
	data := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	rd := strings.NewReader(data)
	resp := NewResp(rd)
	val, err := resp.Read()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if val.typ != "array" {
		t.Fatalf("Expected type 'array', got '%v'", val.typ)
	}

	if len(val.array) != 2 {
		t.Fatalf("Expected array length 2, got %d", len(val.array))
	}

	if val.array[0].bulk != "foo" {
		t.Fatalf("Expected 'foo', got '%v'", val.array[0].bulk)
	}

	if val.array[1].bulk != "bar" {
		t.Fatalf("Expected 'bar', got '%v'", val.array[1].bulk)
	}
}

func TestReadBulk(t *testing.T) {
	data := "$6\r\nfoobar\r\n"
	rd := strings.NewReader(data)
	resp := NewResp(rd)
	val, err := resp.Read()

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if val.typ != "bulk" {
		t.Fatalf("Expected type 'bulk', got '%v'", val.typ)
	}

	if val.bulk != "foobar" {
		t.Fatalf("Expected 'foobar', got '%v'", val.bulk)
	}
}

func TestMarshalArray(t *testing.T) {
	val := Value{
		typ: "array",
		array: []Value{
			{typ: "bulk", bulk: "foo"},
			{typ: "bulk", bulk: "bar"},
		},
	}

	expected := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	result := string(val.Marshal())

	if result != expected {
		t.Fatalf("Expected '%v', got '%v'", expected, result)
	}
}

func TestMarshalBulk(t *testing.T) {
	val := Value{
		typ:  "bulk",
		bulk: "foobar",
	}

	expected := "$6\r\nfoobar\r\n"
	result := string(val.Marshal())

	if result != expected {
		t.Fatalf("Expected '%v', got '%v'", expected, result)
	}
}

func TestWriter(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	val := Value{
		typ: "array",
		array: []Value{
			{typ: "bulk", bulk: "foo"},
			{typ: "bulk", bulk: "bar"},
		},
	}

	err := writer.Write(val)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
	result := buf.String()

	if result != expected {
		t.Fatalf("Expected '%v', got '%v'", expected, result)
	}
}
