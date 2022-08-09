FROM golang:1.18-alpine AS builder
ENV GOPATH /go
WORKDIR /go/src/keyboardify-server
COPY . .

RUN --mount=type=secret,id=APP \
   --mount=type=secret,id=PORT \
   --mount=type=secret,id=URL \
   --mount=type=secret,id=CLIENT \
   --mount=type=secret,id=STRIPE_SECRET_KEY \
   --mount=type=secret,id=FIREBASE_PRIVATE_KEY_JSON \
   --mount=type=secret,id=FIREBASE_TYPE \
   --mount=type=secret,id=FIREBASE_PROJECT_ID \
   --mount=type=secret,id=FIREBASE_PRIVATE_KEY_ID \
   --mount=type=secret,id=FIREBASE_PRIVATE_KEY \
   --mount=type=secret,id=FIREBASE_CLIENT_EMAIL \
   --mount=type=secret,id=FIREBASE_CLIENT_ID \
   --mount=type=secret,id=FIREBASE_AUTH_URI \
   --mount=type=secret,id=FIREBASE_TOKEN_URI \
   --mount=type=secret,id=FIREBASE_AUTH_PROVIDER_X509_CERT_URL \
   --mount=type=secret,id=FIREBASE_CLIENT_X509_CERT_URL \
   echo -e "APP=\"$(cat /run/secrets/APP)\"\nPORT=\"$(cat /run/secrets/PORT)\"\nURL=\"$(cat /run/secrets/URL)\"\nCLIENT=\"$(cat /run/secrets/CLIENT)\"\nSTRIPE_SECRET_KEY=\"$(cat /run/secrets/STRIPE_SECRET_KEY)\"\nFIREBASE_PRIVATE_KEY_JSON=\"$(cat /run/secrets/FIREBASE_PRIVATE_KEY_JSON)\"\nFIREBASE_TYPE=\"$(cat /run/secrets/FIREBASE_TYPE)\"\nFIREBASE_PROJECT_ID=\"$(cat /run/secrets/FIREBASE_PROJECT_ID)\"\nFIREBASE_PRIVATE_KEY_ID=\"$(cat /run/secrets/FIREBASE_PRIVATE_KEY_ID)\"\nFIREBASE_PRIVATE_KEY=\"$(cat /run/secrets/FIREBASE_PRIVATE_KEY)\"\nFIREBASE_CLIENT_EMAIL=\"$(cat /run/secrets/FIREBASE_CLIENT_EMAIL)\"\nFIREBASE_CLIENT_ID=\"$(cat /run/secrets/FIREBASE_CLIENT_ID)\"\nFIREBASE_AUTH_URI=\"$(cat /run/secrets/FIREBASE_AUTH_URI)\"\nFIREBASE_TOKEN_URI=\"$(cat /run/secrets/FIREBASE_TOKEN_URI)\"\nFIREBASE_AUTH_PROVIDER_X509_CERT_URL=\"$(cat /run/secrets/FIREBASE_AUTH_PROVIDER_X509_CERT_URL)\"\nFIREBASE_CLIENT_X509_CERT_URL=\"$(cat /run/secrets/FIREBASE_CLIENT_X509_CERT_URL)\"" > .env

RUN pwd
RUN ls -a

RUN apk update -q
RUN apk add --no-cache git
RUN apk add --no-cache gcc musl-dev
RUN apk add --no-cache ca-certificates

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build .

RUN ls -a
RUN pwd


FROM alpine
WORKDIR /keyboardify-server
COPY --from=builder /go/src/keyboardify-server/keyboardify-server .
COPY --from=builder /go/src/keyboardify-server/.env .
EXPOSE 3000

RUN ls -a
RUN pwd

CMD [ "./keyboardify-server" ]