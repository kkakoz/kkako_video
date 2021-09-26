FROM alpine:latest

WORKDIR /app

ADD user_repo_linux /app/

EXPOSE 9001

ENTRYPOINT ["./user_repo_linux -conf=true"]