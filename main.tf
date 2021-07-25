terraform {

  backend "remote" {
    organization = "elmariofredo"
    workspaces {
      name = "go-count"
    }
  }

  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.10.1"
    }
  }
}

# Add token to terraform cloud
variable "do_token" {}

# Configure the DigitalOcean Provider
provider "digitalocean" {
  token = var.do_token
}

provider "kubernetes" {
  config_path    = "~/.kube/config"
  config_context = "my-context"
}

provider "helm" {
  # kubernetes {
  #   config_path = "~/.kube/config"
  # }

}

resource "digitalocean_kubernetes_cluster" "test" {
  name = "test"
  region = "ams3"
  version = "1.21.2-do.2"

  node_pool {
    name = "default"
    size = "s-1vcpu-2gb"
    node_count = 1
  }

}

resource "helm_release" "nginx_ingress" {
  name       = "nginx-ingress-controller"

  repository = "https://charts.bitnami.com/bitnami"
  chart      = "nginx-ingress-controller"

  set {
    name  = "service.type"
    value = "ClusterIP"
  }
}
