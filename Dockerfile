FROM golang:alpine3.15
WORKDIR /dist

COPY . /src
RUN cd /src && go build -o /dist/thefeed
RUN rm -rf /src

CMD [ "./thefeed" ]