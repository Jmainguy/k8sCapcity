#!/bin/bash
# Set REPO To update
export METAREPO="slmingol/opsMarina"
export DIRNAME=$(echo $METAREPO | awk -F'/' '{print $2}')
# Set this repo name, to gather changelog from here, to add to meta repo.
export THISREPO="jmainguy/k8sCapcity"

git clone "https://$GITHUB_TOKEN@github.com/$METAREPO.git"
cd $DIRNAME

# Only Curl once
CURL=$(curl --silent "https://$GITHUB_TOKEN@api.github.com/repos/$METAREPO/releases/latest")
# Set release tag
LASTRELEASE=$(echo $CURL | jq -r .tag_name)
echo "LASTRELEASE: $LASTRELEASE"
TODAY=$(date +"%Y-%m-%d")
echo "TODAY: $TODAY"
LASTRELEASEDAY=$(echo $LASTRELEASE | awk -F"-" '{print $1"-"$2"-"$3}')
if [[ $TODAY != $LASTRELEASEDAY ]]; then
    export TRAVIS_TAG=$(echo "$TODAY-1")
else
    N=$(echo $LASTRELEASE | awk -F'-' '{print $NF}')
    ((N++))
    export TRAVIS_TAG=$(echo "$TODAY-$N")
fi
echo "TRAVIS_TAG: $TRAVIS_TAG"

# Tag and Push
CURL=$(curl --silent "https://$GITHUB_TOKEN@api.github.com/repos/$THISREPO/releases/latest")
COMMENT=$(echo $CURL | jq -r .body)
THISRELEASE=$(echo $CURL | jq -r .tag_name)
git tag -a $TRAVIS_TAG -m "Release generated by $THISREPO - Release - $THISRELEASE.

$COMMENT"
git push origin $TRAVIS_TAG
