FROM alpine:3.2
ADD microsample-srv /microsample-srv
ENTRYPOINT [ "/microsample-srv" ]
