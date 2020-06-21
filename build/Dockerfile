FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV INVENTORY_BINARY=/usr/local/bin/inventory \
    USER_UID=1001 \
    USER_NAME=inventory

# install inventory ms binary
#COPY ../inventory ${INVENTORY_BINARY}

COPY bin /usr/local/bin
#RUN  /usr/local/bin/inventory

#ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
