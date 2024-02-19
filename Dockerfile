FROM golang:1.18

# 작업 디렉토리 설정
WORKDIR /app

# 의존성 관리 파일 복사
COPY go.mod ./
COPY go.sum ./

# 의존성 설치
#RUN go mod download

# 소스 코드 복사
COPY . .

# 애플리케이션 빌드
RUN go build -o /main .

# 애플리케이션 실행
CMD ["/main"]

