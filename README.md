# xui-exporter

Prometheus exporter for an x-ui panel. It authenticates with the panel API,
collects inbound and client traffic statistics, and exposes them at `/metrics`.

## Requirements

- Go 1.25 or newer for local builds
- Access to an x-ui panel
- Docker and Docker Compose for containerized deployment

## Configuration

The exporter is configured with environment variables.

| Variable | Required | Default | Description |
| --- | --- | --- | --- |
| `PANEL_URL` | Yes | - | Base URL of the x-ui panel, without a trailing slash |
| `USERNAME` | Yes | - | x-ui panel username |
| `PASSWORD` | Yes | - | x-ui panel password |
| `LISTEN_ADDR` | No | `0.0.0.0:9100` | Address and port used by the metrics HTTP server |
| `INSECURE_TLS` | Yes | - | Set to `true` to skip TLS certificate verification, otherwise `false` |

Skipping TLS verification reduces connection security. Use
`INSECURE_TLS=false` when the panel has a valid certificate.

## Run Locally

```sh
export PANEL_URL="https://xui.example.com"
export USERNAME="admin"
export PASSWORD="change-me"
export INSECURE_TLS="false"

go run ./cmd
```

Metrics are then available at:

```text
http://localhost:9100/metrics
```

To build a binary:

```sh
go build -o xui-exporter ./cmd
./xui-exporter
```

## Run with Docker

Build the image:

```sh
docker build -t xui-exporter .
```

Create a `.env` file:

```dotenv
PANEL_URL=https://xui.example.com
USERNAME=admin
PASSWORD=change-me
INSECURE_TLS=false
```

Start the exporter:

```sh
docker compose up -d
```

The provided Compose configuration publishes the exporter at
`http://localhost:9111/metrics`.

## Prometheus Configuration

For a local exporter:

```yaml
scrape_configs:
  - job_name: xui-exporter
    static_configs:
      - targets:
          - localhost:9100
```

For Prometheus running in the same Compose network, use the service name and
container port:

```yaml
scrape_configs:
  - job_name: xui-exporter
    static_configs:
      - targets:
          - xui-exporter:9100
```

## Metrics

| Metric | Type | Description |
| --- | --- | --- |
| `xui_scrape_success` | Gauge | Whether the most recent x-ui API scrape succeeded |
| `xui_inbound_up_bytes` | Gauge | Uploaded traffic for an inbound |
| `xui_inbound_down_bytes` | Gauge | Downloaded traffic for an inbound |
| `xui_inbound_total_bytes` | Gauge | Traffic limit for an inbound; zero usually means unlimited |
| `xui_inbound_enabled` | Gauge | Whether an inbound is enabled |
| `xui_client_up_bytes` | Gauge | Uploaded traffic for a client |
| `xui_client_down_bytes` | Gauge | Downloaded traffic for a client |
| `xui_client_total_bytes` | Gauge | Traffic limit for a client; zero usually means unlimited |
| `xui_client_enabled` | Gauge | Whether a client is enabled |
| `xui_client_expiry_timestamp_ms` | Gauge | Client expiry time as a Unix timestamp in milliseconds |

Inbound metrics include the `id`, `remark`, `protocol`, and `port` labels.
Client metrics include the `inbound_id`, `inbound_remark`, and `email` labels.
Enabled and scrape-success metrics use `1` for true and `0` for false.

## Verify the Exporter

```sh
curl http://localhost:9100/metrics
```

When using the provided Compose configuration:

```sh
curl http://localhost:9111/metrics
```

The exporter logs in during startup. If authentication fails, the process exits.
If a later metrics scrape fails, `xui_scrape_success` is set to `0` and the
error is written to the application log.
