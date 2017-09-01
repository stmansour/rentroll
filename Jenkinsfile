pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh '''export PATH=$PATH:$GOROOT/bin:$GOPATH/bin:$NVM/versions/node/v8.1.4/bin
make all'''
      }
    }
  }
  environment {
    GOROOT = '/use/local/go'
    GOPATH = '/home/sudip/go-work'
    NVM = '/home/sudip/.nvm'
  }
}