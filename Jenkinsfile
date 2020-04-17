pipeline {
    agent {
      docker {
        image 'golang'
      }
    }
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
        stage('Building Image') {
          steps {
            script {
              docker.build registry + ":$BUILD_NUMBER"
            }
          }
        }
    }
}
