# Google Keep Clone

A production-ready Google Keep clone built with React frontend and Go backend, featuring real-time synchronization, OAuth authentication, and modern UI/UX.

## Technology Stack

### Frontend
- **React 18** with TypeScript
- **Vite** for build tooling
- **TailwindCSS** + Shadcn/ui for styling
- **React Query** for server state management
- **Zustand** for client state management
- **Socket.io** for real-time updates

### Backend
- **Go** with Fiber framework
- **PostgreSQL** with GORM ORM
- **Redis** for caching and sessions
- **JWT** for authentication
- **WebSocket** support for real-time features

### DevOps
- **Docker** for containerization
- **Task** for build automation
- **Air** for Go hot reloading
- **Swagger** for API documentation

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go 1.21+** - [Download Go](https://golang.org/dl/)
- **Node.js 18+** - [Download Node.js](https://nodejs.org/)
- **Docker** - [Download Docker](https://www.docker.com/get-started)
- **Task** - [Install Task](https://taskfile.dev/installation/)

### Installing Task (Task Runner)

```bash
# macOS (Homebrew)
brew install go-task/tap/go-task

# Linux/Windows (Go)
go install github.com/go-task/task/v3/cmd/task@latest

# Or download binary from https://github.com/go-task/task/releases
```

## Quick Start

```bash
# 1. Clone the repository
git clone <repository-url>
cd google-keep-clone

# 2. Setup the project (installs dependencies and creates .env)
task setup

# 3. Start development environment
task dev

# 4. Open your browser
# Frontend: http://localhost:5173
# Backend API: http://localhost:8080
# Health Check: http://localhost:8080/health
```

## Task Commands Reference

Task is used as the primary build tool. Run `task --list-all` to see all available commands.

### 🚀 Quick Commands

| Command | Description |
|---------|-------------|
| `task` | Show all available tasks |
| `task help` | Show detailed help |
| `task setup` | Install all dependencies and setup project |
| `task dev` | Start full development environment |
| `task build` | Build production artifacts |
| `task test` | Run all tests |
| `task clean` | Clean all build artifacts |

### 🔧 Development Commands

| Command | Description |
|---------|-------------|
| `task dev` | Start database + backend + frontend |
| `task dev:backend` | Start only backend server |
| `task dev:frontend` | Start only frontend server |
| `task dev:stop` | Stop all development servers |
| `task start` | Alias for `task dev` |
| `task stop` | Stop all services |
| `task restart` | Restart development environment |

### 📊 Database Commands

| Command | Description |
|---------|-------------|
| `task db:start` | Start PostgreSQL and Redis with Docker |
| `task db:stop` | Stop database services |
| `task db:reset` | Reset database (stop, clean, start) |
| `task db:clean` | Remove database containers and volumes |
| `task db:wait` | Wait for database to be ready |
| `task db:logs` | Show database logs |

### 🏗️ Build Commands

| Command | Description |
|---------|-------------|
| `task build` | Build both frontend and backend |
| `task build:frontend` | Build frontend for production |
| `task build:backend` | Build backend binary |

### 🧪 Testing Commands

| Command | Description |
|---------|-------------|
| `task test` | Run all tests |
| `task test:backend` | Run backend tests |
| `task test:frontend` | Run frontend tests |
| `task test:watch` | Run tests in watch mode |
| `task test:api` | Test API endpoints (requires running server) |

### 🔍 Code Quality Commands

| Command | Description |
|---------|-------------|
| `task lint` | Run linting for both frontend and backend |
| `task lint:backend` | Run Go linting (golangci-lint or go vet) |
| `task lint:frontend` | Run frontend linting (ESLint) |
| `task format` | Format code for both frontend and backend |
| `task format:backend` | Format Go code with `go fmt` |
| `task format:frontend` | Format frontend code (Prettier) |

### 🐳 Docker Commands

| Command | Description |
|---------|-------------|
| `task docker:build` | Build Docker images for production |
| `task docker:up` | Start production environment with Docker Compose |
| `task docker:down` | Stop Docker Compose environment |
| `task docker:logs` | Show Docker Compose logs |

### 🛠️ Utility Commands

| Command | Description |
|---------|-------------|
| `task setup:backend` | Setup backend dependencies only |
| `task setup:frontend` | Setup frontend dependencies only |
| `task setup:env` | Setup environment configuration |
| `task clean:backend` | Clean backend build artifacts |
| `task clean:frontend` | Clean frontend build artifacts and dependencies |
| `task deps:update` | Update all dependencies |
| `task tools:install` | Install development tools (Air, golangci-lint, swag) |
| `task docs:generate` | Generate API documentation with Swagger |
| `task info` | Show project information and status |
| `task status` | Show status of all services |
| `task logs` | Show application logs |

## Development Workflow

### 1. Initial Setup
```bash
# Install dependencies and setup environment
task setup

# Install development tools (optional but recommended)
task tools:install
```

### 2. Daily Development
```bash
# Start development environment
task dev

# In separate terminals you can:
task test:watch    # Run tests in watch mode
task lint         # Check code quality
task logs         # Monitor logs
```

### 3. Before Committing
```bash
# Run all checks
task lint
task test
task build

# Optional: Update dependencies
task deps:update
```

### 4. Production Build
```bash
# Build production artifacts
task build

# Or build Docker images
task docker:build
task docker:up
```

## Environment Variables

The project uses a `.env` file for configuration. Copy `.env.example` to `.env` and modify as needed:

```env
# Database
DATABASE_URL=postgres://postgres:postgres@localhost:5432/google_keep_clone?sslmode=disable

# Redis
REDIS_URL=redis://localhost:6379

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Google OAuth (optional)
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret

# API URLs
VITE_API_URL=http://localhost:8080
VITE_WS_URL=ws://localhost:8080

# Environment
ENVIRONMENT=development
PORT=8080
```

## API Documentation

### Health Check
```bash
curl http://localhost:8080/health
```

### Authentication Endpoints
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user
- `GET /auth/me` - Get current user (protected)
- `POST /auth/logout` - Logout user

### Notes Endpoints
- `GET /notes` - Get all notes
- `POST /notes` - Create note
- `GET /notes/:id` - Get note by ID
- `PUT /notes/:id` - Update note
- `DELETE /notes/:id` - Delete note
- `PATCH /notes/:id/pin` - Toggle pin status
- `PATCH /notes/:id/archive` - Toggle archive status
- `PATCH /notes/:id/color` - Update note color
- `GET /notes/search?q=query` - Search notes
- `GET /notes/pinned` - Get pinned notes
- `GET /notes/archived` - Get archived notes

### API Testing
Use the REST files in `api-tests/` directory:
- `api-tests/auth.rest` - Authentication endpoints
- `api-tests/notes.rest` - Notes endpoints

## Project Structure

```
google-keep-clone/
├── frontend/                 # React TypeScript frontend
│   ├── src/
│   │   ├── components/      # React components
│   │   ├── services/        # API services
│   │   ├── types/           # TypeScript types
│   │   └── lib/             # Utility functions
│   ├── package.json
│   └── vite.config.ts
├── backend/                  # Go backend
│   ├── cmd/server/          # Main application
│   ├── internal/
│   │   ├── handlers/        # HTTP handlers
│   │   ├── services/        # Business logic
│   │   ├── repositories/    # Data access layer
│   │   ├── models/          # Database models
│   │   ├── middleware/      # HTTP middleware
│   │   └── validators/      # Request validators
│   ├── go.mod
│   └── go.sum
├── api-tests/               # REST API test files
├── docker/                 # Docker configuration
├── docs/                   # Documentation
├── Taskfile.yaml           # Task runner configuration
├── Makefile               # Make commands (legacy)
├── .env.example           # Environment variables template
└── README.md
```

## Features

### ✅ Implemented
- **User Authentication**: Email/password registration and login
- **JWT Security**: Secure token-based authentication
- **Notes Management**: Full CRUD operations
- **Note Organization**: Pin important notes, archive old ones
- **Visual Customization**: Color-coded notes
- **Search Functionality**: Find notes by title and content
- **Responsive Design**: Works on desktop and mobile
- **API Testing**: Comprehensive REST API tests
- **Type Safety**: Full TypeScript implementation
- **Development Tools**: Hot reloading, linting, testing

### 🚧 Planned
- **Real-time Synchronization**: WebSocket-based live updates
- **Google OAuth**: Social login integration
- **Rich Text Editing**: Advanced note formatting
- **File Attachments**: Upload and attach files to notes
- **Note Sharing**: Share notes with other users
- **Labels System**: Organize notes with custom labels
- **Dark Mode**: Theme switching capability

## Troubleshooting

### Common Issues

1. **Database Connection Failed**
   ```bash
   task db:reset  # Reset database
   task db:wait   # Wait for database to be ready
   ```

2. **Port Already in Use**
   ```bash
   task stop      # Stop all services
   task dev       # Restart development
   ```

3. **Dependencies Issues**
   ```bash
   task clean     # Clean all dependencies
   task setup     # Reinstall everything
   ```

4. **Development Tools Missing**
   ```bash
   task tools:install  # Install Air, golangci-lint, swag
   ```

### Getting Help

- Run `task help` for quick help
- Run `task info` to check project status
- Run `task status` to see service status
- Check logs with `task logs`

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run `task lint && task test` to ensure quality
5. Submit a pull request

## License

MIT License - see LICENSE file for details.

---

**Built with ❤️ using Go, React, and modern development practices.**