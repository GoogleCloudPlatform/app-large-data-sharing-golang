#!/bin/bash

if [[ -z $3 ]]; then
  echo "Usage:"
  echo "download_dummy_images.sh {category name} {number of categories} {number of images per category}"
  exit 1
fi

catename=$1
catenum=$2
imgnum=$3

for i in $(seq 1 $catenum); do
  dirname=${catename}${i}
  mkdir ${dirname}
  for j in $(seq 1 $imgnum); do
    name="${dirname}_${j}.jpg"
    curl "https://picsum.photos/1920/1080.jpg" -L --output ${dirname}/${name}
  done
done