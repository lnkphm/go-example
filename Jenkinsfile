pipeline {
    agent {
        docker { image 'golang:1.21-alpine'}
    }
    stages {
        stage('Build') {
            // Set GOPATH to current workspace
            script {
                def gopath = pwd()
                env.GOPATH = "${gopath}/go"
                sh "mkdir -p ${gopath}/go"
            }
            steps {
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
