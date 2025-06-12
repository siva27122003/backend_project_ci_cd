pipeline {
    agent any

    tools {
        go 'Go-1.24.4'
    }

    environment {
        GOPATH = "${env.WORKSPACE}/go"
        APP_ENV = 'docker'
        GOCACHE = "${env.WORKSPACE}/.cache/go-build"
        DOCKER_IMAGE = 'sivasankar123/bidding-app'
        DOCKER_TAG = 'latest'
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/siva27122003/backend_project_ci_cd.git'
            }
        }

        stage('Clean Go Cache') {
            steps {
                sh 'go clean -modcache'
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

        // stage('Lint') {
        //     steps {
        //         sh '''
        //         export GO111MODULE=on
        //         go vet ./...
        //         golint ./... || true
        //         '''
        //     }
        // }

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
        stage('Build docker image & push in docker hub'){
            steps{
                withCredentials([usernamePassword(credentialsId:'docker-hub-creds',usernameVariable:"USERNAME",passwordVariable:"PASSWORD")]){
                sh '''
                echo "$PASSWORD"| docker login -u "$USERNAME" --password-stdin
                docker build -t $DOCKER_IMAGE:$DOCKER_TAG .
                docker push $DOCKER_IMAGE:$DOCKER_TAG
                '''
                }
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
