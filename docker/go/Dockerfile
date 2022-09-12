FROM cosmtrek/air

ENV TZ="Chile/Continental"
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
ENTRYPOINT ["air"]
