pipeline {
  agent none
  environment {
    AUTHOR = "Iqbal Sonata"
    AUTHOR_EMAIL = "iqbalsonata2@gmail.com"
    AUTHOR_WEB = "https://iqbalsonata.xyz"
  }
  stages {
    stage('Prepare'){
      agent {
        node {
          label "linux && java11"
        }
      }

      environment {
        APP = credentials("iqbal_rahasia")
      }
      steps {
        echo "Author : ${AUTHOR}"
        echo "Author Email : ${AUTHOR_EMAIL}"
        echo "Author Web : ${AUTHOR_WEB}"
        echo "Start Job : ${env.JOB_NAME}"
        echo "Start build : ${env.BUILD_NUMBER}"
        echo "Branch Name  : ${env.BRANCH_NAME}"
        echo "-------------------------------------"
        echo "Username : ${APP_USR}"
        sh('echo "Password : $APP_PSW" > secretfile.txt')

      }
    }
    stage('Build'){
      agent {
        node {
          label "linux && java11"
        }
      }

      steps {
        script {
          for (int i = 0; i < 10;i++){
            echo "Script ${i}"
          }
        }
        echo 'Start build'
        sh 'go build'
        echo 'Finish Build'
      }
    }
    stage('Test'){
      agent {
        node {
          label "linux && java11"
        }
      }

      steps {
        script {
          def data = [
              "firstName" : "Iqbal",
              "lastName" : "Sonata",
          ]
          writeJSON(file: "data.json",json : data)
        }
        echo 'Start Test'
        sh 'go test'
        echo 'Finish Test'
      }
    }
    stage('Deploy'){
      agent {
        node {
          label "linux && java11"
        }
      }

      steps{
        echo 'Starting deploy to public...'
      }
    }
  }

  post {
    always {
      echo "I'll always say Hello again!"
    }
    success {
      echo "Yay,success"
    }
    failure {
      echo "Oh no, it's failure"
    }
    cleanup {
      echo "Don't care success or error"
    }
  }
} 
