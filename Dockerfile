FROM gcr.io/distroless/base

COPY TestGit ./
EXPOSE 8888

CMD ["/TestGit"]