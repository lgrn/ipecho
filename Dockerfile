
# compile stage

FROM golang:1.19-alpine3.17 as build
RUN addgroup -S ipecho \
&& adduser -S -u 10000 -g ipecho ipecho
WORKDIR /go/ipecho
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ipecho

# create image

FROM scratch
LABEL name="ipecho"
LABEL maintainer="https://github.com/lgrn"
LABEL description="A very basic webserver that listens on port 80 and returns the apparent origin IP."
LABEL url="https://github.com/lgrn/ipecho"
COPY --from=build /go/ipecho/ipecho /ipecho
COPY --from=build /etc/passwd /etc/passwd
USER ipecho
EXPOSE 80
ENTRYPOINT ["./ipecho"]

