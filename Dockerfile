FROM golang:1.18-alpine

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
   echo -e "APP=\"$(cat /run/secrets/APP)\"\n" \
   "PORT=\"$(cat /run/secrets/PORT)\"\n" \
   "URL=\"$(cat /run/secrets/URL)\"\n" \
   "CLIENT=\"$(cat /run/secrets/CLIENT)\"\n" \
   "STRIPE_SECRET_KEY=\"$(cat /run/secrets/STRIPE_SECRET_KEY)\"\n" \
   "FIREBASE_PRIVATE_KEY_JSON=\"$(cat /run/secrets/FIREBASE_PRIVATE_KEY_JSON)\"\n" \
   "FIREBASE_TYPE=\"$(cat /run/secrets/FIREBASE_TYPE)\"\n" \
   "FIREBASE_PROJECT_ID=\"$(cat /run/secrets/FIREBASE_PROJECT_ID)\"\n" \
   "FIREBASE_PRIVATE_KEY_ID=\"$(cat /run/secrets/FIREBASE_PRIVATE_KEY_ID)\"\n" \
   "FIREBASE_PRIVATE_KEY=\"$(cat /run/secrets/FIREBASE_PRIVATE_KEY)\"\n" \
   "FIREBASE_CLIENT_EMAIL=\"$(cat /run/secrets/FIREBASE_CLIENT_EMAIL)\"\n" \
   "FIREBASE_CLIENT_ID=\"$(cat /run/secrets/FIREBASE_CLIENT_ID)\"\n" \
   "FIREBASE_AUTH_URI=\"$(cat /run/secrets/FIREBASE_AUTH_URI)\"\n" \
   "FIREBASE_TOKEN_URI=\"$(cat /run/secrets/FIREBASE_TOKEN_URI)\"\n" \
   "FIREBASE_AUTH_PROVIDER_X509_CERT_URL=\"$(cat /run/secrets/FIREBASE_AUTH_PROVIDER_X509_CERT_URL)\"\n" \
   "FIREBASE_CLIENT_X509_CERT_URL=\"$(cat /run/secrets/FIREBASE_CLIENT_X509_CERT_URL)\"" > .env

COPY /.env /go/src/github.com/foxsaysderp/keyboardify-server/.env
WORKDIR /go/src/github.com/foxsaysderp/keyboardify-server
COPY . .

RUN apk add --no-cache git
RUN apk add --no-cache gcc musl-dev
RUN apk --no-cache add ca-certificates
EXPOSE 3000
ENTRYPOINT [ "go", "run", "main.go"]