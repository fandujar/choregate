#!/bin/bash

set -e

CLUSTER_NAME="choregate"
KUBECONFIG_DIR="hack/development"
KUBECONFIG_FILENAME="kind-kubeconfig.yaml"
KUBECONFIG_PATH="$KUBECONFIG_DIR/$KUBECONFIG_FILENAME"

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to get the latest release tag from GitHub
get_latest_release() {
    curl --silent "https://api.github.com/repos/$1/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")'
}

# Determine architecture
ARCH=$(uname -m)
case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64 | arm64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Ensure required commands are available
for cmd in curl grep awk kubectl; do
    if ! command_exists $cmd; then
        echo "Error: $cmd is not installed."
        exit 1
    fi
done

# Install kind if not present
if ! command_exists kind; then
    echo "kind is not installed. Installing kind..."
    KIND_VERSION=$(get_latest_release "kubernetes-sigs/kind")
    curl -Lo ./kind "https://kind.sigs.k8s.io/dl/$KIND_VERSION/kind-linux-$ARCH"
    chmod +x ./kind
    sudo mv ./kind /usr/local/bin/kind
fi

# Create kind cluster if it doesn't exist
if ! kind get clusters | grep -i "$CLUSTER_NAME"; then
    echo "Creating kind cluster..."
    kind create cluster --config hack/development/kind-config.yaml
    kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
    kubectl create namespace choregate
else
    echo "Kind cluster already exists."
fi

# Export kubeconfig
kind get kubeconfig --name "$CLUSTER_NAME" > "$KUBECONFIG_PATH"

# Check if kubeconfig file was created
if [[ ! -f "$KUBECONFIG_PATH" ]]; then
    echo "Error: Kubeconfig file not found at $KUBECONFIG_PATH"
    exit 1
fi

# Create a temporary file for modifications
KUBECONFIG_TMP_PATH=$(mktemp)

# Replace 127.0.0.1 with host.docker.internal
echo "Running awk command to replace 127.0.0.1 with host.docker.internal"
awk '{gsub(/127.0.0.1/, "host.docker.internal"); print}' "$KUBECONFIG_PATH" > "$KUBECONFIG_TMP_PATH" && mv "$KUBECONFIG_TMP_PATH" "$KUBECONFIG_PATH"

# Remove certificate-authority-data line
echo "Running awk command to remove certificate-authority-data line"
awk '!/certificate-authority-data/' "$KUBECONFIG_PATH" > "$KUBECONFIG_TMP_PATH" && mv "$KUBECONFIG_TMP_PATH" "$KUBECONFIG_PATH"

# Add insecure-skip-tls-verify: true
echo "Running awk command to add insecure-skip-tls-verify: true"
awk '/server: https:\/\/host.docker.internal:[0-9]+/ {print; print "    insecure-skip-tls-verify: true"; next}1' "$KUBECONFIG_PATH" > "$KUBECONFIG_TMP_PATH" && mv "$KUBECONFIG_TMP_PATH" "$KUBECONFIG_PATH"

echo "Script completed successfully."
