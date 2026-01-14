# Compare Anything

A React frontend with Go backend application for comparing data in tables. This POC demonstrates the functionality using Nikon Z mount lens data.

## Features

- View all available Nikon Z mount lenses
- Select multiple lenses to compare
- Side-by-side comparison table with all specifications
- Modern, responsive UI

## Prerequisites

- Go 1.21 or higher
- Node.js 16+ and npm

## Setup

### Backend (Go)

1. Navigate to the backend directory:
```bash
cd backend
```

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

The backend will start on port 8080 (or the port specified in the PORT environment variable).

### Frontend (React)

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm start
```

The frontend will start on port 3000 and automatically open in your browser.

## Usage

1. Start both the backend and frontend servers
2. Open your browser to http://localhost:3000
3. Click on lens cards to select them for comparison
4. View the comparison table that appears below
5. Click "Clear Selection" to start over

## Project Structure

```
.
├── backend/
│   ├── main.go          # Go backend server
│   ├── go.mod           # Go module file
│   └── go.sum           # Go dependencies
├── frontend/
│   ├── public/
│   │   └── index.html   # HTML template
│   ├── src/
│   │   ├── App.js       # Main React component
│   │   ├── App.css      # Styles
│   │   ├── index.js     # React entry point
│   │   └── index.css    # Global styles
│   └── package.json     # Node dependencies
└── README.md
```

## API Endpoints

- `GET /api/lenses` - Get all lenses
- `GET /api/lenses/{id}` - Get a specific lens by ID


