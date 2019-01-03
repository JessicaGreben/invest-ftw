FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o invest-ftw

FROM alpine
WORKDIR /app
COPY --from=build-env /src/invest-ftw /app/
COPY --from=build-env /src/static /app/static
ENTRYPOINT ./invest-ftw
