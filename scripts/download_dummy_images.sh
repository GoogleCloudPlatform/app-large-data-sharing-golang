#!/bin/bash

# Copyright 2023 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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