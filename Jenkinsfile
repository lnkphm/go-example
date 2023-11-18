pipeline {
    agent none
    options {
        buildDiscarder(
            logRotator(
                artifactDaysToKeepStr: "50",
                artifactNumToKeepStr: "5",
                daysToKeepStr: "100",
                numToKeepStr: "10"
            )
        )
    }
    stages {
        stage('Build') {
            agent {
                docker {
                    image 'golang:1.21-alpine'
                    args '-u root'
                    reuseNode true
                }
            }
            steps {
                sh 'go mod download'
                sh 'go build -o build/go-example'
            }
        }
        stage('Test') {
            agent {
                docker {
                    image 'golang:1.21-alpine'
                    args '-u root'
                    reuseNode true
                }
            }
            steps {
                sh 'go clean -testcache'
                sh 'go test ./... -v -short'
            }
        }
        stage('Lint') {
            agent {
                docker {
                    image 'golang:1.21-alpine'
                    args '-u root'
                    reuseNode true
                }
            }
            steps {
                sh 'wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2'
                sh '${GOPATH}/bin/golangci-lint --version'
                sh 'golangci-lint run'
            }
        }
        stage('Publish') {
            agent any
            when {
                branch 'main'
            }
            environment {
                IMAGE_REGISTRY = "https://index.docker.io/v1/"
                IMAGE_REPO = "lnkphm/go-example"
                IMAGE_TAG = "latest"
            }
            steps {
                script {
                    image = docker.build("${IMAGE_REPO}:${IMAGE_TAG}", ".")
                    docker.withRegistry("${IMAGE_REGISTRY}", "dockerhub") {
                        image.push()
                    }
                }
            }
        }
        stage('Deploy') {
            agent any
            when {
                branch 'main'
            }
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
            agent any
            cleanWs(
                cleanWhenNotBuilt: false,
                deleteDirs: true,
                disableDeferredWipeout: true,
                notFailBuild: true,
                patterns: [
                    [pattern: '.gitignore', type: 'INCLUDE'],
                    [pattern: '.propsfile', type: 'EXCLUDE']
                ]
            )
        }
    }
}
