FROM golang:1.11.13

RUN go get github.com/labstack/echo     

WORKDIR /go/src/github.com/SimonRininsland/gofino

# Add Local folder to WORKDIR once in build
ADD . .

# get depencies, build and get go watcher
RUN go get ./... && \ 
    go build -v && \
    go get github.com/loov/watchrun

# Monitoring Current Directory And Build Automatically
CMD ["sh","-c","watchrun -monitor *.go go build -v && go run main.go"]

# executing fino.go
#CMD ["go", "run", "main.go"]
