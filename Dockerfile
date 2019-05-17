FROM golang:1.7.3-alpine
LABEL author Tim Urista

COPY cloud-native-go /go/cloud-native-go

RUN chmod +x /go/cloud-native-go

ENV PORT 8080
EXPOSE 8080

# ENTRYPOINT cloud-native-go
# ENTRYPOINT ["./app/cloud-native-go"]



CMD ["./go/cloud-native-go"]