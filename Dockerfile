FROM gcr.io/distroless/base

COPY testgo ./
EXPOSE 8888

CMD ["/TestGit"]