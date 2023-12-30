# A-Maze-ing-Web-App
A personal project to experiment with using Golang to generate and solve mazes built on a simple Sveltekit Frontend, Gin REST API, and Go backend

## Estimated Timeline:
Time to First Deployment: 2 weeks \
Completion: 1 month

## Key Objectives: 
1. Set up and optimize a CI/CD pipeline with Jenkins and Github Actions (run unit tests, static code analysis, etc)
2. Generate and optimize Dockerfiles for each container
3. Configure & integrate containers following a microservices architecture
4. Demonstrate understanding of key Golang concepts by generating and solving mazes: functional programming, network services, and concurrent program design (via goroutines and multi-cores for parallel computation)
5. Basic frontend setup with Sveltekit to allow user interaction with the maze (mouse drag to solve, view solution, etc)
6. Integrate with Gin REST API framework to connect backend services to frontend views.
7. Cloud Deployment with AWS (EC2 instance or ECS) 
8. Leverage IaC with Terraform to declaratively provision Cloud resources

## Limitations:
1. Generate and solve "perfect" mazes (only one solution: i.e one entry & exit)
2. Very crud frontend (focus of this project is on devops and go backend configurations)
3. Showcase basic idea of concurrent programming in Go using a bidirectional-bfs solution (can go into a lot more depth but sufficient for the demonstrative purposes of this application)