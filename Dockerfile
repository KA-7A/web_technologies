FROM golang:latest

WORKDIR /server

COPY *go* ./
COPY ./src ./src


RUN go mod download
RUN go build -o main.out main.go
EXPOSE 8080
RUN echo Well, lets start this shit
CMD ["./main.out"]