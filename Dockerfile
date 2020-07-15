FROM golang:1.13

MAINTAINER Alex Bondar <abondar1992@gmail.com>

WORKDIR /app
RUN mkdir json
COPY . /app

RUN go build -o app github.com/abondar24/RecipeCountAnalyzer

ENTRYPOINT ["./app"]
