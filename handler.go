package main

import "sync"

var Handlers = map[string]func([]Value) Value{
	"PING": ping,
	"SET":  set,
	"GET":  get,
	// "HSET":    hset,
	// "HGET":    hget,
	// "HGETALL": hgetall,
}

func ping(args []Value) Value {
	if len(args) == 0 {
		return Value{typ: "string", str: "PONG"}
	}

	return Value{typ: "string", str: args[0].bulk}
}

var SETs = map[string]string{}
var SETsMu = sync.RWMutex{}

func set(args []Value) Value {
	if len(args) != 2 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}
	SETsMu.Lock()
	defer SETsMu.Unlock()

	key := args[0].bulk
	value := args[1].bulk

	SETs[key] = value

	return Value{typ: "string", str: "OK"}
}

func get(args []Value) Value {
	if len(args) != 1 {
		return Value{typ: "error", str: "ERR wrong number of arguments for 'set' command"}
	}
	SETsMu.RLock()
	defer SETsMu.RUnlock()

	key := args[0].bulk

	value, ok := SETs[key]
	if !ok {
		return Value{typ: "null"}
	}

	return Value{typ: "bulk", bulk: value}
}