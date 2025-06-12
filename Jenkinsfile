pipeline {
    agent any

    tools {
        go 'Go-1.24.4'
    }

    environment {
        GOPATH = "${env.WORKSPACE}/go"
        APP_ENV = 'docker'
        GOCACHE = "${env.WORKSPACE}/.cache/go-build"
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/siva27122003/backend_project_ci_cd.git'
            }
        }

        stage('Install Dependencies') {
            steps {
                sh '''
                export GO111MODULE=on
                go mod tidy
                '''
            }
        }

        stage('Lint') {
            steps {
                sh '''
                export GO111MODULE=on
                go vet ./...
                golint ./... || true
                '''
            }
        }

        stage('Test') {
            steps {
                sh '''
                export GO111MODULE=on
                go test ./Handler -v -cover
                '''
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o bin/server'
            }
        }

        stage('Zip Build') {
            steps {
                sh 'zip -r build_bidding_app.zip bin/'
            }
        }

        stage('Archive Artifact') {
            steps {
                archiveArtifacts artifacts: 'build_bidding_app.zip', fingerprint: true
            }
        }
    }
}
