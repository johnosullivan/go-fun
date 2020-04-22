
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
              sh 'echo $REPOSITORY_NAME'
              sh 'go version'
              sh 'go get'
          }
      }
      stage('Compile') {
          steps {
              sh 'go build'
          }
      }
      stage('Testing') {
          steps {
              echo 'Testing Working ;)'
          }
      }
      stage('Building Image') {
          steps {
              sh 'docker build -t $REPOSITORY_NAME .'
          }
      }
      stage('Publish Image') {
          steps {
              sh 'docker tag $REPOSITORY_NAME:latest $REPOSITORY_URI/$REPOSITORY_NAME:latest'
              sh 'docker push $REPOSITORY_URI/$REPOSITORY_NAME:latest'
          }
      }
      stage('Clean Up') {
          steps {
              sh 'docker rmi $REPOSITORY_URI/$REPOSITORY_NAME'
              sh 'docker rmi $REPOSITORY_NAME:latest'
          }
      }
  }
}
