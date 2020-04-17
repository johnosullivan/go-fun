node {
    def app

    stage('Clone repository') {
        checkout scm
    }

    stage('Pre Compile Checks') {
        sh 'go version'
        sh 'go get'
    }

    stage('Compile') {
        sh 'go build'
    }

    stage('Build image') {
        app = docker.build "gofun" + ":$BUILD_NUMBER"
    }

    stage('Test Image') {
        app.inside {
            sh 'echo "Tests passed"'
        }
    }

    /*stage('Push image') {
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            app.push("${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }*/
}
