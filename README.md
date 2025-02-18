# A-Maze-ing 
 
## Objective
The primary purpose of this application is to generate an interactive 3D maze game for players to navigate and explore!

Through this explorative process of developing an application from start to finish, I hope to develop new skills in IaC automation, CI/CD pipelines, software architecture, 3D rendering on the web and fullstack development.

> Note: From August, 2024 onward, service has been moved to vercel. You can check out this project at https://a-maze-ing-ivory.vercel.app 

[![alt text](/assets/maze_gui.png)](https://a-maze-ing-ivory.vercel.app)

## Quick Start

### Local Dev

- To setup local env, run the script `./set-local-env.sh` or manually copy `.env.local` to `.env` in `frontend/` and `backend/`.

- To run the whole application, do `docker compose up --build -d` ( optionally, build and run each container separately)

#### Run services separately

- To run each microservice seperately, first run docker-compose as above, then kill the frontend and backend containers `docker kill <CONTAINER ID>`
- (**frontend**) `cd frontend && yarn start` -> `localhost:5173/`
- (**api**) `cd backend` then you can either run `go run .` or use the `air` plugin for golang (_preferred_ with live-reload) -> `localhost:8080/`
- (**postgres_db**) locally run image using docker-compose

### API Access

The following endpoints are currently currently well defined:

#### Unauth

1. POST `/maze/generate` generates a new maze and solution
2. POST `/maze/solve` solves an arbitrary perfect maze
3. GET `/maze?limit` returns a list of maze records ordered by `solve_time`
4. GET `/maze/:user` returns a list of maze records belonging to a user
5. GET `/api-health` returns `status = 200` if api is healthy

#### Auth

1. PUT `/maze/:user` updates/create user's maze records
2. DELETE `/maze/:user` deletes a user's maze records

#### Example

```
// generate
{
    width: number, 3 <= w <= 35
    height: number, 3 <= h <= 35
    generator: "prim" | "kruskal"
    solver: "bfs" | "bbfs" | "dfs"
}

// solve
{
    width: number, 3 <= w <= 35
    height: number, 3 <= h <= 35
    generator: "prim" | "kruskal"
    solver: "bfs" | "bbfs" | "dfs"
    maze: {
        start: number,
        end: number,
        grid: [] // width x width matrix, try using output from generate
    }
    solution: [] // user's attempt at solving, program can validate
}
```

e.g generate a 20x20 maze with prim and solve using bfs
![Alt text](/assets/maze_cli.png)

## Tools / Technologies

> ✅ -> Currently included, ❌ -> Not yet / partially

- Frontend
  - Svelte + Vite &nbsp;&nbsp; ✅
  - TailwindCSS &nbsp;&nbsp;✅
  - Skeleton UI &nbsp;&nbsp;✅
  - Vercel for hosting &nbsp;&nbsp;✅
  - Vitest tests &nbsp;&nbsp;✅
  - Threlte for rendering maze / user interaction &nbsp;&nbsp;✅
  - Auth0 (login / sign-up / logout) &nbsp;&nbsp;✅
- Backend
  - Gin-gonic API &nbsp;&nbsp;✅
    - Request validation &nbsp;&nbsp;✅
    - Handlers to perform CRUD to db and generate/solve mazes &nbsp;&nbsp;✅
    - Authentication (validate client's JWT access token from Auth0) &nbsp;&nbsp;✅
    - Middlewares:
      - CORS &nbsp;&nbsp;✅
      - Security (prevent XSS, CSRF)
      - JWT validation &nbsp;&nbsp;✅
      - Logger with rotation &nbsp;&nbsp;✅
      - Global error handling with standardized error output &nbsp;&nbsp;✅
  - Services written in Go &nbsp;&nbsp;✅
    - Maze generation (i.e. Kruskal or Prim to build MST)
    - Maze solving (i.e. Bidirectional Search with concurrency using goroutines, BFS, DFS)
  - Testing: httptest and testify (partial ❌)
  - Database
    - Postgres &nbsp;&nbsp;✅
- DevOps & Infrastructure &nbsp;&nbsp; ✅

  - GitHub actions + Status Checks on PR
  - GitHub main branch protection
  - CircleCi
  - Dockerfiles with multi-stage builds to minimize image size (local dev only currently, but considered pushing/pull to registry as well)
  - Docker-compose (local dev)
    - synchronize local development
    - service ordering with health checks
  - Terraform to provision AWS resources (Not in use anymore)
    - Declarative soltuion with HCL
    - Provisioned VPC, Subnets, IG, EC2, SG, ASG, RDS, ASG, IAM
  - Packer to bake AWS AMI for Jenkins master and workers

## Key Objectives:

1. Set up and experiment with CI/CD (i.e. Github Actions, CircleCI, Jenkins)
2. Generate and optimize Dockerfiles for each container (Multi-stage builds + caching)
3. Work with Docker containers and orchestrate locally with Docker Compose
4. Demonstrate understanding of key Golang concepts by generating and solving mazes (i.e. api setup, concurrency with goroutines + wait groups + channels, etc. )
5. Achieve seamless integration between frontend UI views and backend services.
6. Cloud deployment with AWS/Heroku/Vercel
7. Leverage IaC with Terraform to declaratively provision cloud resources
8. Use Packer to bake custom AMIs

## Limitations:

1. Generate and solve "perfect" mazes (MSTs with only one path from entry to exit)
2. Simple frontend - a working solution is sufficient
3. Limited use of concurrency (can go into a lot more depth but sufficient for the demonstrative purposes of this application)

### Credits:

(Rook model used in the Svelte application)
This work is based on "Low Poly Chess - Rook" (https://sketchfab.com/3d-models/low-poly-chess-rook-cbd416e785f64648bff3675fd45b3594) by marcelo.medeirossilva (https://sketchfab.com/marcelo.medeirossilva) licensed under CC-BY-NC-SA-4.0 (http://creativecommons.org/licenses/by-nc-sa/4.0/)
