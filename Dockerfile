FROM 381491935073.dkr.ecr.us-east-1.amazonaws.com/wsc:golang
 
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-docker

EXPOSE 8000

CMD [ "/go-docker" ]