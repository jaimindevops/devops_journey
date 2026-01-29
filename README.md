# Cloud-Native Go Application on Kubernetes 🚀

This project demonstrates a complete **Cloud-Native** workflow: developing a Go application, containerizing it with Docker, and deploying it to a Kubernetes cluster with High Availability and Load Balancing.

## 🏗 Architecture
* **Application:** Custom Go web server that identifies its own hostname (Pod ID).
* **Containerization:** Docker (Multi-stage build).
* **Orchestration:** Kubernetes (K3s).
* **Networking:** NodePort Service for external access.

## 🛠 Tech Stack
* **Language:** Go (Golang) 1.21
* **Infrastructure:** K3s (Lightweight Kubernetes)
* **OS:** Ubuntu 22.04 LTS (Oracle Cloud Infrastructure)

## 🚀 How to Run

###1. Build the Image

```bash
docker build -t my-go-app:v1 .

### 2. Import to K3s (If using K3s)

Since K3s uses containerd, we import the local image directly to avoid pushing to a registry during development.

Bash
docker save my-go-app:v1 -o my-go-app.tar
sudo k3s ctr images import my-go-app.tar

### 3. Deploy to Cluster

Bash
kubectl apply -f go-deployment.yaml
kubectl apply -f go-nodeport.yaml

### 4. Verify Load Balancing

The application prints the Pod Name. Hitting the endpoint multiple times demonstrates Round-Robin load balancing across replicas.

Bash
curl http://localhost:32667
# Output: "Hello Network Engineer! I am running on Pod: go-app-xyz"
📂 Project Structure
main.go: The Go source code.

Dockerfile: Instructions to package the app.

go-deployment.yaml: K8s Manifest for the Deployment (3 Replicas).

go-nodeport.yaml: K8s Manifest for the Service (Exposing Port 32667).

