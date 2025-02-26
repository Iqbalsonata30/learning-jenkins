pipeline {
  agent {
    node {
      label "linux && java11"
    }
  }
  stages {
    stage('Build'){
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
