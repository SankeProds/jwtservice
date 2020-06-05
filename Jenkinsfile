pipeline {
    agent { label 'master' }
    stages {
        stage('build') {
            steps {
                sh 'docker build -t jwtservice .'
            }
        }
    }
}