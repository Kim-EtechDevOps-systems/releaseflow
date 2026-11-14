pipeline {
    agent { label 'linux-go-python' }

    environment {
        CGO_ENABLED = '1'
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
            APP_VERSION="${GIT_COMMIT:-dev}" ./releaseflow > app.log 2>&1 &
            APP_PID=$!

            trap 'kill "$APP_PID" 2>/dev/null || true' EXIT

            if ! python3 automation/healthcheck.py \
                --url http://127.0.0.1:8081/health; then
                echo "Application health check failed"
                cat app.log
                exit 1
            fi
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
