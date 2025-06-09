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
                sh 'mkdir -p bin'
                sh 'go build -o bin/server'
                sh 'ls -l bin'
            }
        }
    }

    post {
        success {
            archiveArtifacts artifacts: 'bin/server', fingerprint: true
        }
        always {
            cleanWs()
        }
    }
}
