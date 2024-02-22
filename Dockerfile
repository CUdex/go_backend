FROM golang:1.22.0

# 작업 디렉토리 설정
WORKDIR /go/src

# 의존성 관리 파일 복사
COPY go.mod ./
COPY go.sum ./
# 소스 코드 복사
COPY . .

# 의존성 설치
# RUN go mod init tag-controller
# RUN go mod tidy
RUN go mod download

# 애플리케이션 빌드
RUN go build -o /backController .

# 애플리케이션 실행
CMD ["/backController"]