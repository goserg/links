FROM golang:1.16
RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN make build
EXPOSE 8080
CMD ["bin/links"]
