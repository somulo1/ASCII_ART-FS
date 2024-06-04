#!/bin/bash

# git config --global credential.helper cache

echo $(git add .)

commit=$1

echo $(git commit -m "$commit")

read=$(git rev-parse --abbrev-ref HEAD)

echo $(git push github $branch)

echo $(git push gitea $branch)
