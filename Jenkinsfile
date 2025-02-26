pipeline {
  agent {
    node {
      label "linux && java11"
    }
  }
  stages {
    stage('Stage 1'){
      steps {
        echo "Hello World from SCM"
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
      echp "Don't care success or error"
    }
  }
} 
