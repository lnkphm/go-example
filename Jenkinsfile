pipeline {
    agent {
        docker { image 'golang:1.21-alpine'}
    }
    stages {
        stage('Build') {
            steps {
                echo 'Hello world'
                sh 'ls -l'    
                sh 'pwd'
            }
        }
    }
}
