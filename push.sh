#!/bin/sh

commit_message=$1
current_branch=$(git rev-parse --abbrev-ref HEAD)

if [ "$commit_message" = "" ] || [ "$commit_message" = " " ]; then
  commit_message="update some scripts"
fi

git add .
git commit -m "$commit_message"
git push origin $current_branch
