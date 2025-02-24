FROM golang:1.22 as base

WORKDIR /app

COPY go.mod .

RUN go mod download 

COPY . . 

RUN go build -o main. . 

# final stage -- distroless image 

FROM gcr.io/distroless/base 

COPY --from=base /app/main. .

COPY --from=base /app/static ./static 

EXPOSE 8000

CMD [ "./main." ]

