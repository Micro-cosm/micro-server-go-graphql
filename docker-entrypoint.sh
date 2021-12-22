#!/bin/sh

pwd
ls -alR
echo "!!!!! Starting Go server -- /${SERVICE}/docker-entrypoint.sh !!!!!";

/go/${SERVICE}/${EXECUTABLE}