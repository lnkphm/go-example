pipeline {
    agent {
        docker {
            image 'golang:1.21-alpine'
        }
    }
    environment {
        GOPATH = "${env.WORKSPACE}/go"
        GOCACHE = "/tmp"
        IMAGE_REGISTRY = 'registry-1.docker.io'
        IMAGE_REPO = 'lnkphm/go-example'
        IMAGE_TAG = 'latest'
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
                sh 'go clean -testcache'
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
    post {
        success {
            echo 'Pipeline successful'
        }
        failure {
            echo 'Pipeline failed'
        }
        always {
            echo 'Pipeline completed'
        }
    }
}
