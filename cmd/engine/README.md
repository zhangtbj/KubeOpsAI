# Engine Controller

The `Engine` controller is a Kubernetes controller for managing `Question` custom resources. It listens for new `Question` resources, processes them, and updates their status with an answer, phase, message, and timestamp. This controller is built using [Kubebuilder](https://book.kubebuilder.io/), making it easy to manage custom Kubernetes resources.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Directory Structure](#directory-structure)
- [Running the Controller Locally](#running-the-controller-locally)
- [Deploying to a Cluster](#deploying-to-a-cluster)
- [Testing the Controller](#testing-the-controller)
- [Contributing](#contributing)

## Prerequisites
To develop this controller, ensure you have the following tools installed:
- [Go](https://golang.org/doc/install) 1.19 or higher
- [Kubebuilder](https://book.kubebuilder.io/quick-start.html#installation) (for scaffolding and code generation)
- [Docker](https://docs.docker.com/get-docker/) (for containerizing the controller)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/) (to interact with Kubernetes)
- Access to a Kubernetes cluster (local or remote)

## Getting Started
1. Clone the repository:
   ```bash
    git clone https://github.com/zhangtbj/KubeOpsAI.git
    cd KubeOpsAI
   ```

2. Install dependencies:
   ```bash
    go mod tidy
   ```

3. Generate necessary code, including deepcopy functions and CRD manifests:
   ```bash
    make generate
    make manifests
   ```

## Directory Structure

Here’s an overview of the main directories:
- `api/v1alpha1`: Contains the `Question` CRD Go definitions.
- `controllers`: Contains the main logic for the `Engine` controller.
- `config`: Kubernetes manifests for deploying the CRD and controller. 
  - `config/crd`: Stores the CRD YAML files.
  - `config/default`: Default configuration for the controller.
  - `config/manager`: Deployment manifest for the controller manager.
  - `config/rbac`: RBAC settings for accessing `Question` resources.

## Running the Controller Locally

To test the controller locally without deploying it to a cluster, you can use kubebuilder’s `make run`:

1. Ensure your Kubernetes cluster is running and your `kubectl` context is set correctly.
2. Apply the CRD to your cluster:
   ```bash
    make install
   ```
3. Run the controller locally:
   ```bash
    make run
   ```
This will start the controller and watch for `Question` resources in your Kubernetes cluster.

## Deploying to a Cluster

To deploy the controller to a Kubernetes cluster:

1. Build the Docker image:
   ```bash
    docker build -t yourusername/engine-controller:latest .
   ```
2. Push the Docker image to a container registry:
   ```bash
    docker push yourusername/engine-controller:latest
   ```
3. Update the image in the deployment YAML: Modify `config/manager/manager.yaml` to use your image:
   ```yaml
    spec:
    containers:
    - name: manager
      image: yourusername/engine-controller:latest
   ```
4. Deploy the controller:
   ```bash
    make deploy
   ```
This command will apply the necessary RBAC, CRD, and controller manager manifests to your cluster.

## Testing the Controller

To test the controller, you can create an Engine custom resource and check if the controller updates its status correctly.

1. Create an Engine resource YAML, e.g., sample_engine.yaml:
   ```yaml
    apiVersion: core.kubeopsai.ai/v1alpha1
    kind: Engine
    metadata:
    name: sample-engine
    spec:
    question: "What is the status of my Kubernetes cluster?"
   ```

2. Apply the Engine resource:
   ```bash
    kubectl apply -f sample_engine.yaml
   ```
3. Check the status of the Engine resource:
   ```bash
    kubectl get engine sample-engine -o yaml
   ```
You should see the status fields (answer, phase, message, lastUpdated) populated by the controller.

## Contributing

Contributions are welcome! To contribute:

- Fork the repository.
- Create a new branch for your feature or bug fix.
- Make your changes and add tests if applicable.
- Open a pull request with a description of your changes.

License
This project is licensed under the Apache-2.0 License.
