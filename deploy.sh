#!/bin/sh
source config.sh
if [ "$hugo_gallery_domain" == "" ]; then
    echo "Missing configuration values"
    exit 1
fi

s3cmd sync static/images/ s3://$hugo_gallery_domain/
