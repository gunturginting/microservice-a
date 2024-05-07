pipeline {
    environment {
        APP_NAME = "microservice-a"
        DOCKER_USER = "aey16"
        DOCKER_PASS = 'docker'
        IMAGE_NAME = "${DOCKER_USER}" + "/" + "${APP_NAME}"
        IMAGE_TAG = "${env.GIT_COMMIT.substring(0,7)}"
        REPO_CODE = "https://github.com/gunturginting/microservice-a.git"
        REPO_SECRET = "github-jenkins"
        IMAGE_REPO = "aey16/${APP_NAME}"
        NAMESPACE = "skripsi"
    }
    agent any

    stages {
        stage('SCM') {
            steps {
                git branch: 'main', credentialsId: "${REPO_SECRET}", url: "${REPO_CODE}"
            }
        }
        
        stage('Build Image') {
            steps {
                script {
                    docker.withRegistry('',DOCKER_PASS) {
                        docker_image = docker.build "${IMAGE_NAME}"
                        docker_image.push("${IMAGE_TAG}")
                    }
                }
            }
        }
        
        stage('Deploy to Kubernetes') {
            steps {
                script {
                    withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
                        sh "mkdir -p ~/.kube/"
                        sh "cat ${KUBECONFIG} >> ~/.kube/config"
                        
                        sh """
                        helm upgrade --install ${APP_NAME} ./config/${APP_NAME} \
                        --set-string image.repository=${IMAGE_REPO},image.tag=${IMAGE_TAG} \
                        -f ./config/${APP_NAME}/values.yaml --debug --namespace ${NAMESPACE}
                        """
                    }
                }
            }
        }
        
        stage('Clean') {
            steps {
                script {
                    sh "docker image prune -f"
                    sh "rm ~/.kube/config"
                }
            }
        }
    }
}
