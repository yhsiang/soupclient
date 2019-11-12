FROM alpine:3.2
ADD go-soupclient /app/go-soupclient
ADD go-soupclient-darwin-amd64 /app/go-soupclient-darwin-amd64
ADD go-soupclient-window-amd64.exe /app/go-soupclient-window-amd64.exe
ENV PATH="/app:${PATH}"