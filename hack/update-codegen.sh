#!/bin/bash

#
# Generates the typed client for Kubernetes CRDs
# From https://www.openshift.com/blog/kubernetes-deep-dive-code-generation-customresources
#

set -exuo pipefail

# Set the project root directory (modify as necessary)
PROJECT_ROOT=$(pwd)

# Specify your module name (e.g., github.com/yourname/yourproject)
MODULE_NAME="github.com/zhangtbj/KubeOpsAI"

# Define the API group and version, e.g., "yourgroup:v1alpha1"
API_GROUP="core"
API_VERSION="v1alpha1"

echo ""
echo "Using code-generator package version, as instructed in the go.mod file"
echo "The code-generator package is imported via the pkg/kubecodegen dir"
echo "To modify the current version, please modify this in the go.mod"
echo ""

# Run the deepcopy generator
bash ./vendor/k8s.io/code-generator/generate-groups.sh "deepcopy" \
  "${MODULE_NAME}/pkg/generated" \
  "${MODULE_NAME}/pkg/apis" \
  "${API_GROUP}:${API_VERSION}" \
  --go-header-file "${PROJECT_ROOT}/hack/boilerplate.go.txt"