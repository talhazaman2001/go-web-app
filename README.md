# Go Web App - Enterprise-Grade DevOps Pipeline

![Infrastructure](https://img.shields.io/badge/Infrastructure-AWS-orange)
![Container](https://img.shields.io/badge/Container-Docker-blue)
![Orchestration](https://img.shields.io/badge/Orchestration-Kubernetes-blue)
![IaC](https://img.shields.io/badge/IaC-Terraform-purple)
![CI/CD](https://img.shields.io/badge/CI/CD-GitHub_Actions-green)
![GitOps](https://img.shields.io/badge/GitOps-ArgoCD-red)
![Templating](https://img.shields.io/badge/Templating-Helm-blue)

A production-ready DevOps implementation that transforms a simple Go application into a robust cloud service through five strategic stages of DevOps transformation.

## 🚀 Overview

In today's competitive tech landscape, deployment velocity and infrastructure reliability are no longer optional — they're essential competitive advantages. This project demonstrates a comprehensive, enterprise-grade DevOps implementation that transforms a simple Go application into a production-ready cloud service.

## 📋 Project Implementation Stages

### 1️⃣ Containerisation with Docker
- Efficient Docker image build process
- Multi-stage builds for optimized container size
- Proper base image selection for security and performance

### 2️⃣ Infrastructure as Code with Terraform
- AWS EKS cluster provisioning
- Kubernetes infrastructure setup
- Network configuration and security groups

### 3️⃣ Traffic Management with Nginx Ingress Controller
- External traffic routing to Kubernetes services
- TLS termination and certificate management
- Load balancing configuration

### 4️⃣ Application Deployment with Helm
- Templated Kubernetes manifests
- Environment-specific configurations
- Release management and versioning

### 5️⃣ CI/CD Pipeline with GitHub Actions and ArgoCD
- Automated build and testing
- Container vulnerability scanning
- GitOps-based deployment workflow

## 🏗️ Architecture

The combination of Terraform for IaC, Docker for Containerisation, Kubernetes for container orchestration, Helm for templating and environment management, and GitOps with ArgoCD creates a powerful, automated workflow that ensures consistency and reliability.

This architecture not only streamlines deployments but also embodies DevOps best practices by:
- Maintaining infrastructure as code
- Enabling version control for all components
- Providing clear audit trails
- Ensuring consistency across environments

## 🛠️ Prerequisites

- AWS Account
- Docker
- Kubectl
- Helm
- Terraform
- GitHub account

## 🚦 Getting Started

1. Clone the repository:
   ```
   git clone https://github.com/talhazaman2001/go-web-app.git
   cd go-web-app
   ```

2. Follow the instructions in each directory to implement the respective stage.

## 📚 Detailed Documentation

For a step-by-step guide and detailed explanation of each implementation stage, please refer to the [Medium article](https://talhazaman01.medium.com/5-stage-devops-transformation-blueprint-for-a-go-web-app-e81a8cfe32c3).

## 💡 Key Features

- **Infrastructure as Code**: Every aspect of the infrastructure is defined in code
- **GitOps Workflow**: Changes propagate through git commits
- **Multi-Environment Support**: Production, staging, and development environments
- **Automated Security Scanning**: Container and dependency vulnerability scanning
- **Scalable Architecture**: Built on Kubernetes for horizontal scaling

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📞 Contact

For questions or feedback, please reach out via LinkedIn - link on my GitHub profile.

---

This project serves as a blueprint of foundational DevOps principles which can be easily applied to more complex solutions in the future.
