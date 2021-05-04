FROM scratch

ENTRYPOINT ["/bin/go-count"]

COPY go-count /bin/go-count
