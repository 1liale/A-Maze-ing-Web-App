pipeline {
    agent any 
    
    stages{
        stage("Clone Code"){
            steps {
                echo "Cloning from GitHub..."
                git url:"https://github.com/1liale/A-Maze-ing-Generator-Solver.git", branch: "release"
            }
        }
        stage("Backend Build"){
            steps {
                echo "Building backend image..."
                sh "pwd"
                dir('backend') {
                    sh "pwd"
                }
                sh "pwd"
            }
        }
    }
}