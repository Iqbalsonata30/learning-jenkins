pipeline {
  agent none
  environment {
    AUTHOR = "Iqbal Sonata"
    AUTHOR_EMAIL = "iqbalsonata2@gmail.com"
    AUTHOR_WEB = "https://iqbalsonata.xyz"
  }
  triggers {
    // cron "*/5 * * * *"
    pollSCM "* * * * *"
  }
  parameters {
    string(name: 'NAME', defaultValue: 'Guest',description: 'What is your name?')
    text(name: 'DESCRIPTION', defaultValue: '',description: 'Tell me about you')
    booleanParam(name: 'DEPLOY', defaultValue: false,description: 'Need to deploy?')
    choice(name: 'SOCIAL_MEDIA', choices: ['Instagram','Twitter','Facebook'],description: 'Which social media?')
    password(name: 'SECRET',defaultValue: "",description: 'Encrypt Key')
  }
  options {
    disableConcurrentBuilds()
    timeout(time: 1,unit: 'MINUTES')
  }
  stages {
    stage('Preparation'){
      agent {
        node {
          label "linux && java11"
        }
      }
      stages {
        stage('Prepare Golang'){
          steps{
            echo "Prepare Go"
            sleep(5)
          }
        }
        stage('Prepare React'){
          steps{
            echo "Prepare React"
            sleep(5)
          }
          }
        }
      }
    }
    
    stage('Parameter'){
      agent {
        node {
          label "linux && java11"
        }
      }
      steps {
        echo "Hello ${params.NAME}"
        echo "You are ${params.DESCRIPTION}"
        echo "Okay to deploy: ${params.DEPLOY}"
        echo "Your Social Media : ${params.SOCIAL_MEDIA}"
        echo "Your Secret Key : ${params.SECRET}"
      }
    }
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
      input {
        message "Can we deploy?"
        ok "Yes,of course"
        submitter "iqbalsonata"
        parameters {
          choice(name: "TARGET_ENV", choices:['QA','DEV','PROD'],description: "Which environment?")
        }
      }
      agent {
        node {
          label "linux && java11"
        }
      }

      steps{
        echo "Deploy to ${TARGET_ENV}"
      }
    }
    stage('Release'){
      when {
        expression {
          return params.DEPLOY
        }
      }
      agent {
        node {
          label "linux && java11"
        }
      }
      steps{
        echo("Release it")
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
