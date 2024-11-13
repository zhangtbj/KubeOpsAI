//go:build tools
// +build tools

// Require for getting the right code-generator package in the vendor directory
// so that we can generate typed clientsets for CustomResource APIGroups. This is
// called by the hack/update-codegen.sh script

package codegen

import _ "k8s.io/code-generator" // required by hack/update-codegen.sh script
