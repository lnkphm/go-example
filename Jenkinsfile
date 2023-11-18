pipeline {
    agent {
        docker {
            image 'golang:1.21-alpine'
            args '-u root'
        }
    }
    environment {
        GOPATH = "${env.WORKSPACE}/go"
    }
    stages {
        stage('Build') {
            steps {
                sh 'mkdir -p ${GOPATH}'
                sh 'go mod download'
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
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
