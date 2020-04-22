pipeline {
    agent any
    tools {
        go 'go'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Pre Compile Checks') {
            steps {
                sh 'go version'
                sh 'go get'
            }
        }
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
        stage('Building Image') {
              steps {
                sh 'docker build -t go-fun .'
              }
        }
        stage('Publish Image') {
              steps {
                sh 'docker tag go-fun:latest 396527728813.dkr.ecr.us-west-1.amazonaws.com/go-fun:latest'
                sh 'docker push 396527728813.dkr.ecr.us-west-1.amazonaws.com/go-fun:latest'
              }
        }
    }
}
