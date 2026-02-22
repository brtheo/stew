#!/bin/bash

set -e

VERSION="$1"

if [[ -z "$VERSION" ]]; then
  echo "Need a version"
  exit 1
fi

PLATFORMS=("windows/amd64" "linux/amd64" "darwin/amd64" "linux/arm64" "darwin/arm64")

rm -rf npm/bin/*

for platform in "${PLATFORMS[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name='npm/bin/stew-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo "Building for $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w -X main.VersionNumber=$VERSION" -o $output_name
done
