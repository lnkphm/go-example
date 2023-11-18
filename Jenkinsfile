pipeline {
    agent {
        docker {
            image 'golang:1.21-alpine'
            args '-u root'
        }
    }
    options {
        buildDiscarder(
            logRotator(
                numToKeepStr: '5',
                artifactNumToKeepStr: '5'
            )
        )
    }
    environment {
        IMAGE_REGISTRY = "registry-1.docker.io"
        IMAGE_REPO = "lnkphm/go-example"
        IMAGE_TAG = "latest"
    }
    stages {
        stage('Build') {
            steps {
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
        stage('Lint') {
            steps {
                sh 'wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2'
                sh '${GOPATH}/bin/golangci-lint --version'
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
            cleanWs()
        }
    }
}
