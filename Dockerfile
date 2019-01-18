FROM golang:latest 
RUN mkdir /quizzup 
ADD . /quizzup/ 
WORKDIR /quizzup 
RUN go build -o main . 
CMD ["/quizzup/main"]