pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scmGit(branches: [[name: '*/master']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/ritesh-15/mapup-assessment-golang']])
            }
        }

        stage('Run') {
            steps {
                script {
                    docker compose up --build -d
                    echo "App is now running"
                }
            }
        }
    }
}