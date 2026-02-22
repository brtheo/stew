#!/bin/bash

set -e

VERSION="$1"

if [[ -z "$VERSION" ]]; then
  echo "Need a version"
  exit 1
fi
cat << EOF > npm/package.json
{
  "name": "@brtheo/stew",
  "version": "$VERSION",
  "description": "A dumb process manager",
  "main": "index.js",
  "scripts": {
    "postinstall": "node install.js"
  },
  "bin": {
    "stew": "index.js"
  },
  "repository": {
    "type": "git",
    "url": "git+https://github.com/brtheo/stew.git"
  },
  "author": "Brossier théo",
  "license": "MIT",
  "bugs": {
    "url": "https://github.com/brtheo/stew/issues"
  },
  "homepage": "https://github.com/brtheo/stew"
}
EOF

git add npm/package.json
git commit -m "Publish version ${VERSION}"
git tag "v${VERSION}" -f
git push
git push --tags
pushd npm
npm publish --access public .
popd
