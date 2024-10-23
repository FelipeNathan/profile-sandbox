pipeline {
    agent any

    stages {
        stage('Running compose') { 
            steps {
                sh '''
                    docker compose up -d
                '''
            }
            
        }
    }
}

