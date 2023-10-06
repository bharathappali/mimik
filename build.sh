#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Please install Go and try again."
    exit 1
fi

# Check if mimik executable exists, delete the older executable
if [ -f mimik ]; then
    echo "[INFO] Older \`mimik\` executable detected."
    echo -n "[PROCESS] Deleting older \`mimik\` executable ... "
    rm -f ./mimik
    if [ -f mimik ]; then
      echo "FAILED."
    else
      echo "Done."
    fi
fi

# Build mimik executable
go build -o mimik

echo -n "[PROCESS] Creating \`mimik\` ... "
# Check if the executable was created
if [ -f mimik ]; then
    echo "Done."
else
    echo "FAILED."
    exit 1
fi

