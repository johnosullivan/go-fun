pipeline {
    agent any
    tools {
        go 'go-1.14.2'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Pre Compile Checks') {
            steps {
                sh 'go version'
                sh 'ls'
                sh 'go mod init'
                sh 'go get'
            }
        }
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
    }
}
