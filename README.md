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

### 1. Build the Image
```bash
docker build -t my-go-app:v1 .
