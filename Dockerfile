FROM golang:1.23 

WORKDIR /app 

COPY go.mod . 

COPY go.sum . 

RUN go mod tidy 

COPY . . 

COPY .env . 

RUN go build -o ewallet-notification 


RUN chmod +x ewallet-notification 

EXPOSE 7003 

CMD [ "./ewallet-notification" ]