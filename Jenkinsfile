pipeline {
    agent any
    tools {
        go 'go-1.14.2'
    }
    environment {
        GO111MODULE = 'on'
        registry = "docker_hub_account/repository_name"
        registryCredential = 'dockerhub'
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
        stage('Tests') {
            steps {
                echo 'Testing was good!'
            }
        }
        stage('Building Docker Image') {
            steps {
                sh 'docker build -t gofun .'
            }
        }
    }
}
