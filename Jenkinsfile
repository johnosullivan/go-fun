pipeline {
  agent any
  tools {
      go 'go'
  }
  environment {
      GO111MODULE = 'on'
  }
  stages {
      /*
        Initialize all the pipline parameters and thresholds.
      */
      stage("Initialize") {
        steps {
          initialize()
        }
      }
      stage('Compile') {
          steps {
              sh 'go build'
          }
      }
      stage('Testing') {
          steps {
              echo 'go test -v'
          }
      }
      stage('Building Image') {
          steps {
              sh 'docker build -t $REPOSITORY_NAME .'
          }
      }
      /*stage('Publish Image') {
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
      }*/
  }
}

def initialize() {

}
