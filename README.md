# Kubi8al Webhook

A webhook receiver that forwards GitHub webhook events to an emitter API.

## Features

- Receives GitHub webhook events
- Extracts important repository information
- Forwards events to a configured emitter API
- Built with Go and Fiber framework
- Containerized deployment with Docker

## Prerequisites

- Go 1.21 or higher
- Docker (for containerized deployment)
- Git

## Environment Variables

The application requires the following environment variables:

- `EMMITER_API_ADDRESS`: URL of the emitter API to forward webhook events
- `WEBHOOK_SECRET`: (Optional) Secret token for webhook authentication
- `PORT`: (Optional, defaults to 8080) Port to listen on

## Building and Running

### Local Development

1. Build the application:
```bash
make build
```

2. Run locally:
```bash
make run
```

### Docker

1. Build Docker image:
```bash
make docker-build
```

2. Run Docker container:
```bash
make docker-run EMMITER_API_ADDRESS=https://your-emitter-api.com WEBHOOK_SECRET=your-secret
```

## Testing

Run the test suite:
```bash
make test
```

## Linting and Formatting

Run linters:
```bash
make lint
```

Format code:
```bash
make fmt
```

## Clean Build Artifacts

Remove build artifacts:
```bash
make clean
```

## Available Make Targets

- `make all`: Build the application (default target)
- `make build`: Build the application binary
- `make run`: Run the application locally
- `make docker-build`: Build Docker image
- `make docker-run`: Run Docker container
- `make test`: Run test suite
- `make lint`: Run linters
- `make fmt`: Format code
- `make clean`: Clean build artifacts
- `make deps`: Install dependencies
- `make update-deps`: Update dependencies
- `make help`: Show available targets

## Deployment

The application is designed to be deployed using Docker. The Docker image is published to GitHub Container Registry (GHCR) with tags matching the git tags.

## Security

- Webhook authentication is supported via the `WEBHOOK_SECRET` environment variable
- All network connections use HTTPS
- Docker image is built with security best practices

## License

MIT License

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Support

For support, please open an issue on the GitHub repository.