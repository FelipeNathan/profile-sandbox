pipeline {
    agent any

    stages {
        stage('Running compose') { 
            steps {
                sh '''
                    docker stop web
                    docker rm web
                    docker compose up -d --build web
                '''
            }
            
        }
    }
}

