pipeline {
    agent any

    tools {
      go 'Go-1.24.4'
    }

    environment {
        GOPATH = "${env.WORKSPACE}/go"
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/siva27122003/backend_project_ci_cd.git'
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o bin/server ./cmd/server'
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
