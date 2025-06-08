# Google Keep Clone - Initialization Guide

This guide provides step-by-step instructions to initialize and set up the Google Keep clone project with React frontend and Go backend.

## Quick Start

### Prerequisites
- Node.js 18+
- Go 1.21+
- PostgreSQL
- Redis
- Docker (optional)

### 1. Initialize Project Structure

```bash
# Create the project structure
mkdir -p frontend backend api-tests docker docs
git init
echo "# Google Keep Clone" > README.md

# Create initial progress tracking
cat > docs/PROGRESS.md << EOF
# Development Progress

## Phase 1: Project Setup ⏳
- [ ] Repository structure
- [ ] Frontend setup with Vite + React
- [ ] Backend setup with Go + Fiber
- [ ] Database setup (PostgreSQL)
- [ ] Basic Docker configuration
- [ ] Initial git commits

## Phase 2: Authentication 📋
- [ ] OAuth with Google
- [ ] Email/Password authentication
- [ ] JWT implementation
- [ ] Protected routes

## Phase 3: Core Features 📋
- [ ] Note CRUD operations
- [ ] Real-time synchronization
- [ ] Note categories and labels
- [ ] Search functionality

## Phase 4: Advanced Features 📋
- [ ] File attachments
- [ ] Note sharing
- [ ] Archive functionality
- [ ] Rich text editing

## Phase 5: Production Ready 📋
- [ ] Testing suite
- [ ] Performance optimization
- [ ] Security hardening
- [ ] Deployment setup
EOF

git add . && git commit -m "feat: initial project structure and progress tracking"
```

### 2. Frontend Setup

```bash
cd frontend
npm create vite@latest . -- --template react-ts
npm install

# Install dependencies
npm install \
  @tanstack/react-query \
  @hookform/resolvers \
  react-hook-form \
  react-router-dom \
  zustand \
  zod \
  framer-motion \
  socket.io-client \
  @radix-ui/react-toast \
  @radix-ui/react-dialog \
  @radix-ui/react-dropdown-menu \
  class-variance-authority \
  clsx \
  tailwind-merge \
  lucide-react

# Install dev dependencies
npm install -D \
  @types/node \
  tailwindcss \
  postcss \
  autoprefixer \
  eslint-config-prettier \
  prettier

# Setup Tailwind CSS
npx tailwindcss init -p
```

### 3. Backend Setup

```bash
cd ../backend
go mod init google-keep-clone

# Install dependencies
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/contrib/jwt
go get github.com/golang-jwt/jwt/v5
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/go-redis/redis/v8
go get github.com/google/uuid
go get github.com/joho/godotenv
go get golang.org/x/crypto/bcrypt
go get golang.org/x/oauth2/google
go get github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/fiber-swagger
go get github.com/stretchr/testify
```

### 4. Environment Configuration

Create `.env` file in project root:

```env
# Database
DATABASE_URL=postgres://postgres:postgres@localhost:5432/google_keep_clone?sslmode=disable

# Redis
REDIS_URL=redis://localhost:6379

# JWT
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production

# Google OAuth
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret

# API URLs
VITE_API_URL=http://localhost:8080
VITE_WS_URL=ws://localhost:8080

# Environment
ENVIRONMENT=development
PORT=8080
```

### 5. Database Setup

```bash
# Start PostgreSQL and Redis with Docker
docker run --name postgres -e POSTGRES_DB=google_keep_clone -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres:15-alpine

docker run --name redis -p 6379:6379 -d redis:7-alpine
```

### 6. Development Workflow

Create `Makefile`:

```makefile
.PHONY: dev build test clean docker-up docker-down

# Development
dev:
	docker-compose -f docker/docker-compose.yml up postgres redis -d
	cd backend && air &
	cd frontend && npm run dev

# Build
build:
	cd frontend && npm run build
	cd backend && go build -o bin/server cmd/server/main.go

# Testing
test:
	cd backend && go test ./...
	cd frontend && npm run test

# Install dependencies
install:
	cd backend && go mod tidy
	cd frontend && npm install

# Clean
clean:
	cd backend && go clean
	cd frontend && rm -rf dist node_modules
```

## Development Phases

### Phase 1: Basic Setup
1. Complete project initialization
2. Set up basic frontend and backend structure
3. Configure development environment
4. Test basic connectivity

### Phase 2: Authentication
1. Implement JWT authentication service
2. Create login/register endpoints
3. Add Google OAuth integration
4. Build authentication UI components

### Phase 3: Core Features
1. Design database models for notes
2. Implement CRUD operations
3. Create note management UI
4. Add basic note editing

### Phase 4: Real-time Features
1. Set up WebSocket server
2. Implement real-time synchronization
3. Add live updates to frontend
4. Test multi-device sync

### Phase 5: Production Setup
1. Create Docker configuration
2. Set up CI/CD pipeline
3. Add monitoring and logging
4. Deploy to production

## Technology Stack

### Frontend
- React 18 with TypeScript
- Vite for build tooling
- TailwindCSS + Shadcn/ui
- React Query for server state
- Zustand for client state
- Socket.io for real-time updates

### Backend
- Go with Fiber framework
- PostgreSQL with GORM
- Redis for caching
- JWT authentication
- WebSocket support

### DevOps
- Docker for containerization
- Make for build automation
- Air for Go hot reloading
- Swagger for API docs

## Next Steps

1. **Start Development**: Run `make dev` to start the development environment
2. **Follow Phases**: Implement features following the documented phases
3. **Test Regularly**: Use the provided test files in `api-tests/`
4. **Track Progress**: Update `docs/PROGRESS.md` as you complete features
5. **Commit Often**: Use meaningful commit messages for each feature

## Support

- Check `claude_instructions.md` for detailed implementation guidance
- Use API test files in `api-tests/` directory
- Refer to `docs/PROGRESS.md` for current status
- Follow the structured development phases for best results

Ready to build your Google Keep clone! 🚀