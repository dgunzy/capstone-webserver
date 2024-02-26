#!/bin/bash

NEW_TAG=$(date +%Y-%m-%d-%H%M%S)
echo "Creating new tag: $NEW_TAG"

git fetch --prune --tags 2>/dev/null

if git rev-parse "$NEW_TAG" >/dev/null 2>&1; then
    echo "Tag $NEW_TAG already exists. Exiting without creating a new tag."
    exit 1
else
    git tag $NEW_TAG
    git push --tags
    echo "Tagged with $NEW_TAG"
fi

echo ::set-output name=git-tag::$NEW_TAG

exit 0
