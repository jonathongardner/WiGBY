FROM appgardner/gocv:v4.6.0 as build
WORKDIR /go/src/dlc
RUN curl -LJO https://gist.githubusercontent.com/jonathongardner/311ff78be0ccec2a32746979b3459c0e/raw/16cc48701146c2aa778275b1b2710bafd1098712/main.go && GO111MODULE=off go build -o /bin/dlc

WORKDIR /go/src/wegyb
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build && dlc wegyb /build

FROM scratch
COPY --from=build /build /
ENV LD_LIBRARY_PATH=/usr/local/lib/

ENTRYPOINT ["/wegyb"]
