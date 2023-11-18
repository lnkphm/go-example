pipeline {
    agent {
        docker { image 'golang:1.21-alpine'}
    }
    stages {
        stage('Build') {
            steps {
                sh 'go clean -cache'
                sh 'go mod download'
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                sh 'go clean -cache'
                sh 'go test ./... -v -short'
            }
        }
        stage('Publish') {
            steps {
                echo "Publish image here"
            }
        }
        stage('Deploy') {
            steps {
                echo "Trigger deployment here"
            }
        }
    }
}
