pipeline {
    agent { label 'linux-go-python' }

    options {
        timestamps()
        timeout(time: 10, unit: 'MINUTES')
        disableConcurrentBuilds()
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Go Quality Checks') {
            steps {
                dir('app') {
                    sh '''
                        test -z "$(gofmt -l .)"
                        go vet ./...
                        go test -race -v ./...
                    '''
                }
            }
        }

        stage('Build') {
            steps {
                dir('app') {
                    sh 'go build -o ../releaseflow .'
                }
            }
        }

        stage('Integration Test') {
            steps {
                sh '''
                    APP_VERSION="${GIT_COMMIT:-dev}" \
                        ./releaseflow > app.log 2>&1 &
                    APP_PID=$!
                    trap 'kill "$APP_PID" 2>/dev/null || true' EXIT

                    python3 automation/healthcheck.py \
                        --url http://127.0.0.1:8080/health
                '''
            }
        }
    }

    post {
        always {
            archiveArtifacts(
                artifacts: 'app.log',
                allowEmptyArchive: true
            )
            cleanWs()
        }
    }
}
