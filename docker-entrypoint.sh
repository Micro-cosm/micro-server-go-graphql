#!/bin/sh

cd /go/graphql;
ls -al;
file graphql
echo "!!!!! Starting Go server(/${SERVICE}/docker-entrypoint.sh) !!!!! ${EXECUTABLE}";
/go/${SERVICE}/${EXECUTABLE}