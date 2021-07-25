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
