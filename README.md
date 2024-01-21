# A-Maze-ing
A personal full stack project to go into a deep dive on all aspects of infrastructure automation, CI/CD, DS & Algo, and frontend + backend web development.

Ultimately, the goal of this project is less about the exact content it serves, which is a fun program that achieves the task of generating and solving perfect mazes. 
 
Instead, by going through each step from ideation to production. I hope to become a more well-rounded developer capable of writing well-tested, structured, and performant code that can meet production quality.

*Self note: progress is incremental, have patience and drink some tea 🍵

> Status: WIP

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