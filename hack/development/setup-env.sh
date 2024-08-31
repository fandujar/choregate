#!/bin/bash

set -e

CLUSTER_NAME="choregate"
KUBECONFIG_PATH="/tmp/kind-kubeconfig.yaml"

command_exists() {
    command -v "$1" >/dev/null 2>&1
}

get_latest_release() {
    curl --silent "https://api.github.com/repos/$1/releases/latest" | grep -Po '"tag_name": "\K.*?(?=")'
}

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

if ! command_exists kind; then
    echo "kind is not installed. Installing kind..."
    KIND_VERSION=$(get_latest_release "kubernetes-sigs/kind")
    curl -Lo ./kind "https://kind.sigs.k8s.io/dl/$KIND_VERSION/kind-linux-$ARCH"
    chmod +x ./kind
    sudo mv ./kind /usr/local/bin/kind
fi

if ! kind get clusters | grep -i "$CLUSTER_NAME"; then
    echo "Creating kind cluster..."
    kind create cluster --config hack/development/kind-config.yaml
    kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
    kubectl create namespace choregate
else
    echo "Kind cluster already exists."
fi

kind get kubeconfig --name "$CLUSTER_NAME" > "$KUBECONFIG_PATH"

if [[ ! -f "$KUBECONFIG_PATH" ]]; then
    echo "Error: Kubeconfig file not found at $KUBECONFIG_PATH"
    exit 1
fi

echo "Running sed command to replace 127.0.0.1 with host.docker.internal"
sed 's/127.0.0.1/host.docker.internal/g' "$KUBECONFIG_PATH" > /tmp/kind-kubeconfig.tmp && mv /tmp/kind-kubeconfig.tmp "$KUBECONFIG_PATH"

echo "Running sed command to remove certificate-authority-data line"
sed '/certificate-authority-data/d' "$KUBECONFIG_PATH" > /tmp/kind-kubeconfig.tmp && mv /tmp/kind-kubeconfig.tmp "$KUBECONFIG_PATH"

echo "Running awk command to add insecure-skip-tls-verify: true"
awk '/server: https:\/\/host.docker.internal:[0-9]+/ {print; print "    insecure-skip-tls-verify: true"; next}1' "$KUBECONFIG_PATH" > /tmp/kind-kubeconfig.tmp && mv /tmp/kind-kubeconfig.tmp "$KUBECONFIG_PATH"
