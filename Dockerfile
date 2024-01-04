#Build based on Golang minimal image
#Set initial image as build base for next stage
FROM golang:alpine AS Buildstage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8090

RUN go build -o /mastenk_pplgawad

#Build operational image from previous stage, with minimal alpine image
#Alpine edge needed for linux library
FROM alpine:edge AS Main

WORKDIR /app

COPY --from=Buildstage ./mastenk_pplgawad /mastenk_pplgawad

COPY local.env .

EXPOSE 8090

ENTRYPOINT [ "/mastenk_pplgawad" ]