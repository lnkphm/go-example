pipeline {
    agent {
        docker {
            image 'golang:1.21-alpine'
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
        GOPATH = "${env.WORKSPACE}/go"
        GOCACHE = "/tmp"
        IMAGE_REGISTRY = 'registry-1.docker.io'
        IMAGE_REPO = 'lnkphm/go-example'
        IMAGE_TAG = 'latest'
    }
    stages {
        stage('Cleanup') {
            steps {
                deleteDir()
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
