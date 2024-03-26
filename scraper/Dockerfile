FROM ubuntu:22.04

RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates

copy ./scraper /bin/scraper

ENTRYPOINT ["/bin/scraper"]
