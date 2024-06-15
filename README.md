# Redis Clone

This project is a simplified clone of Redis, supporting basic commands and using the Redis Serialization Protocol (RESP) for communication and Data persistence using Append Only File (AOF).

\*\*This project is for educational purposes and not meant to be on production.\*\*

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [Supported Commands](#supported-commands)
- [Project Structure](#project-structure)

## Overview

This Redis clone supports the following commands:

- `SET`
- `GET`
- `HSET`
- `HGET`
- `HGETALL`

## Installation

To install and run this project:

1. Clone the repository:

   ```sh
   git clone https://github.com/MikhailWahib/redis-clone.git
   cd redis-clone
   ```

2. Build the project:

   ```sh
   make build
   ```

3. Run the server:
   ```sh
   ./bin/run
   ```

## Usage

1. Start your Redis clone server as shown above.
2. Use the [Redis CLI](https://redis.io/docs/latest/develop/connect/cli/) on the default port to interact with your server:
   ```sh
   redis-cli -h 127.0.0.1 -p 6379
   ```

## Supported Commands

- **SET**: Set a key to a string value.
  ```sh
  SET key value
  ```
- **GET**: Get the value of a key.
  ```sh
  GET key
  ```
- **HSET**: Set a field in a hash to a value.
  ```sh
  HSET hash field value
  ```
- **HGET**: Get the value of a field in a hash.
  ```sh
  HGET hash field
  ```
- **HGETALL**: Get all fields and values in a hash.
  ```sh
  HGETALL hash
  ```

## Project Structure

- `main.go`: Entry point of the application.
- `resp.go`: Contains RESP parsing logic.
- `handler.go`: Contains the implementations of supported commands handlers.
- `aof.go`: Contains the logic of AOF.
