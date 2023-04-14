#!/bin/bash
set -euo pipefail

BUCKET_NAME=$1
ARCHIVE_FILE_NAME=$2
LDS_CLIENT_URL=$3

echo "bucket name is $BUCKET_NAME"
echo "archive file name is $ARCHIVE_FILE_NAME"
gsutil cp gs://$BUCKET_NAME/$ARCHIVE_FILE_NAME .
if [[ "$ARCHIVE_FILE_NAME" == *.tar.gz ]]; then
  name=$(basename $ARCHIVE_FILE_NAME .tar.gz)
  mkdir $name
  tar -zxvpf "$ARCHIVE_FILE_NAME" -C $name 
  folders=($(find $name -type d | awk -F/ '{print $NF}'))
  for folder in "${folders[@]}"
  do
      bash upload.sh $LDS_CLIENT_URL $name/$folder
  done
else
  echo "Unknown file type: $ARCHIVE_FILE_NAME"
  exit 1
fi
