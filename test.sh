#!/usr/bin/env bash

go build

cwd=$(pwd)

rm -rf sample-site 2> /dev/null
mkdir -p sample-site/static/images
mkdir -p sample-site/content

testImages=(image1.jpg image2.jpg image3.jpg)

for image in ${testImages[@]}; do
    mkdir -p sample-site/static/images/
    touch sample-site/static/images/$image
done

cd sample-site
pwd
ls ../

../hugo-gallery static/images test_gallery "Test Gallery" domain.com

cd $cwd
