# Google Keep Clone

A production-ready Google Keep clone built with React frontend and Go backend, featuring real-time synchronization, OAuth authentication, and modern UI/UX.

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

## Quick Start

1. **Install dependencies**: `make install`
2. **Start development**: `make dev`
3. **Run tests**: `make test`
4. **Build for production**: `make build`

## Development

See `init.md` for detailed setup instructions and `claude_instructions.md` for complete development guide.

## Features

- User authentication (email/password + Google OAuth)
- CRUD operations for notes
- Real-time synchronization
- Note colors and labels
- Search functionality
- Archive and pin notes
- Responsive design

## License

MIT License