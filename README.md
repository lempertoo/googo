# Port Scanner

A lightweight TCP port scanner written in Go.

## Overview

Port Scanner checks whether TCP ports are open on a target host.

The project demonstrates:

- Goroutines
- Channels
- TCP networking
- Modular architecture

## Features

- Concurrent scanning
- Configurable worker count
- Sorted output
- Scan summary

## Structure

```
.
├── main.go
├── scanner.go
├── worker.go
├── formatter.go
├── config.go
└── ports.go
```

## Run

```bash
go run .
```

## Future Improvements

- UDP scan
- Banner grabbing
- CLI arguments
- JSON export
- Progress bar

