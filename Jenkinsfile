pipeline {
  agent {
    node {
      label "linux && java11"
    }
  }
  stages {
    stage('Build'){
      steps {
        echo 'Build'
      }
    }
    stage('Test'){
      steps {
        echo 'Test'
        sh 'error'
      }
    }
    stage('Deploy'){
      steps{
        echo 'Deploy'
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
