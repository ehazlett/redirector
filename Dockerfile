FROM scratch
COPY redirector /bin/redirector
COPY ./certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/bin/redirector"]
EXPOSE 8080
CMD ["-h"]
