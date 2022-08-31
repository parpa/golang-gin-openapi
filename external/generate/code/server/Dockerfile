FROM openapitools/openapi-generator-cli

COPY generate.sh /usr/local/bin/generate-server
RUN chmod a+x /usr/local/bin/generate-server

ENTRYPOINT [ "generate-server" ]