# A-Maze-ing
A personal project to develop new skills in infrastructure automation, CI/CD, DS & Algo, and fullstack web development. 

Ultimately, the goal of this project is less about the exact content it serves, which is a fun program that generates and solves perfect mazes. Instead, by going through each step from ideation to production, I hope to challenge myself to become a better developer.

## Quick Start
### Local Dev
To run the whole application, use the following command
1. `cp ./backend/.env.local ./backend/.env`
2. `cp .env.local .env && docker compose up --build`

or optionally you could build and run each container separately

To access backend API, \
`cd backend && cp .env.local .env && go run .`


The following endpoints are currently currently well defined:
1. POST `localhost:8080/maze/generate` will generate a new maze and solution
2. POST `localhost:8080/maze/solve` will solve arbitrary perfect maze

The request body is as follows: 
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
![Alt text](/assets/maze_example1.png)

For just frontend, \
`cd frontend && yarn dev`

> Frontend is still WIP, will only show default skeleton setup as of now

### Production (Not fully setup)

## Tools / Technologies
- Frontend
    - Sveltekit
    - TailwindCSS
    - Vercel for hosting
    - Jest tests
    - (optional) Three.js for rendering maze / user interaction
- Backend
    - Gin-gonic API
        - HTTP requests validated
        - handlers to perform CRUD to db and generate/solve mazes
        - authication
        - middlewares:
            - custom logger with rotation
            - global error handling with standardized error output
    - Services written in Go
        - Maze generation (i.e. Kruskal to build MST)
        - Maze solving (i.e. Bidirectional search done concurrently with goroutines, channels, and wait groups)
    - Testing with httptest and testify
- Database
    - Postgres
- DevOps & Infrastructure
    - GitHub actions + Status Checks on PR
    - GitHub main branch protection
    - Jenkins Pipeline (CI: build and test, no CD planned currently) 
    - Dockerfiles with multi-stage builds to minimize image size
    - Docker-compose
        - synchronize local development
        - service ordering using healthchecks and condition
    - Terraform to provision AWS resources 
        - declarative structure with HCL
        - VPC, Subnets, IG, EC2, SG, ASG, RDS, ASG, IAM

    - Packer to bake AWS AMI for Jenkins master and workers

## Key Objectives: 
1. Set up and optimize a CI/CD pipeline with Jenkins and Github Actions (run unit tests, static code analysis, etc)
2. Generate and optimize Dockerfiles for each container
3. Configure & integrate containers following a microservices architecture
4. Demonstrate understanding of key Golang concepts by generating and solving mazes: functional programming, network services, and concurrent program design (via goroutines and multi-cores for parallel computation)
5. Basic frontend setup with Sveltekit to allow user interaction with the maze (mouse drag to solve, view solution, etc)
6. Integrate with Gin REST API framework to connect backend services to frontend views.
7. Cloud Deployment with AWS
8. Leverage IaC with Terraform to declaratively provision Cloud resources
9. Use Packer to bake custom AMIs

## Limitations:
1. Generate and solve "perfect" mazes (MSTs with only one path from entry to exit)
2. Crud frontend (focus of this project is on devops and go backend configurations)
3. Limited use of concurrency (can go into a lot more depth but sufficient for the demonstrative purposes of this application)
