# Project Structure

```text
.
├── README.md                 # Entry point for project, containing basic information and setup instructions.
├── LICENSE                   # The license file for project (MIT).
├── go.mod                    # Go module definition file.
├── go.sum                    # Go module dependencies lock file.
├── Makefile                  # Build and development automation commands.
├── compose.yaml              # Docker Compose configuration for containerized service.
├── docs/                    # Contains all project documentation.
│   ├── api-endpoints.md     # API endpoint documentation.
│   ├── project-overview.md   # High-level overview of project goals and scope.
│   ├── project-structure.md  # This file. Explains organization of codebase.
│   └── star-wars-unlimited.md # Domain knowledge and context about Star Wars Unlimited trading card game.
├── service/                 # Backend service implementation.
│   ├── main.go             # Pocketbase application entry point with custom routes.
│   └── Dockerfile          # Multi-stage Dockerfile for building the service.
└── tests/                  # Behavioral tests for the application.
    └── helloworld_test.go  # Test for the /hello endpoint.
```
