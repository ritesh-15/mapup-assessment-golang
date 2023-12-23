pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                script {
                    checkout scmGit(branches: [[name: '*/master']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/ritesh-15/mapup-assessment-golang']])
                }
            }
        }

        stage('Run') {
            steps {
                script {
                    sh "sudo docker compose down"
                    sh "sudo docker compose up --build -d"
                }
            }
        }
    }
}