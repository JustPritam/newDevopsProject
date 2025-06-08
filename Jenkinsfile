pipeline {
    agent any

    environment {
        IMAGE_NAME = 'todo-app'
        PORT = '8081'
    }

    stages {
        stage('Clone') {
            steps {
                git 'git@github.com:JustPritam/newDevopsproject.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t pritamcalsoft/new_proj:1 .'
            }
        }

        stage('Run App') {
            steps {
                sh '''
                docker rm -f pritamcalsoft/new_proj:1 || true
                docker run -d -p 8081:8081 --name pritamcalsoft/new_proj:1 pritamcalsoft/new_proj:1 
                '''
            }
        }
    }
}
