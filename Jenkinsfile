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
                artifactDaysToKeepStr: "50",
                artifactNumToKeepStr: "5",
                daysToKeepStr: "100",
                numToKeepStr: "10"
            )
        )
    }
    environment {
        IMAGE_REGISTRY = "https://index.docker.io/v1/"
        IMAGE_REPO = "lnkphm/go-example"
        IMAGE_TAG = "latest"
    }
    stages {
        stage('Build') {
            steps {
                sh 'go mod download'
                sh 'go build -o build/go-example'
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
                sh 'golangci-lint run'
            }
        }
        stage('Publish') {
            agent {
                any
            }
            when {
                branch 'main'
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
