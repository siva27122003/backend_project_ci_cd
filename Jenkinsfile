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
                sh 'go build -o main main.go'
            }
        }

        stage('Zip Build') {
            steps {
                sh '''
                    apt-get update
                    apt-get install zip 
                    mkdir -p artifact_output
                    cd Jenkins_pipeline/src
                    zip ../../artifact_output/bidding_app.zip main
                '''
            }
        }
    }
}
