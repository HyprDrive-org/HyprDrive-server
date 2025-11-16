# HyprDrive-server

HyprDrive is an open-source backend for a privacy-focused storage platform.
This repository currently contains the foundational Go server layout that future
components (database, queue workers, TUI client, and web UI) will build on top of.

## Project layout

```
.
├── cmd/
│   └── server/           # entrypoint for the API server binary
├── internal/
│   ├── config/           # configuration loading and validation helpers
│   ├── httpserver/       # HTTP server setup and routing
│   └── logger/           # thin wrapper around log.Logger for future swaps
├── go.mod
├── LICENSE
└── README.md
```

## Getting started

1. **Run the server**

   ```bash
   go run ./cmd/server
   ```

   The server exposes a `/healthz` endpoint that returns a simple JSON payload.

2. **Configuration**

   Environment variables can override defaults:

   | Variable               | Default      | Description                       |
   | ---------------------- | ------------ | --------------------------------- |
   | `HYPRDRIVE_HTTP_PORT`  | `8080`       | Port to bind the HTTP server      |
   | `HYPRDRIVE_ENV`        | `development`| Application environment label     |

This structure is intentionally minimal but ready to grow with storage drivers,
background workers, and the planned Bubble Tea TUI client.
