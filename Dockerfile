FROM golang:alpine AS Buildstage

WORKDIR /app

COPY . ./

RUN go mod download

EXPOSE 8090

RUN go build -o /mastenk_pplgawad

FROM golang:alpine AS Main

WORKDIR /

COPY --from=Buildstage /mastenk_pplgawad /

COPY local.env ./

EXPOSE 8090

ENTRYPOINT [ "/mastenk_pplgawad" ]