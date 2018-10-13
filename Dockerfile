# build stage
FROM golang:alpine AS build-stage
COPY . /src
RUN cd /src && go build -o mvno 

# final stage
FROM alpine
COPY --from=build-stage /src/mvno /usr/local/bin/
ENTRYPOINT ["mvno"]
