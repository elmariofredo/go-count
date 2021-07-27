# Add token to terraform cloud
variable "do_token" {}

# Configure the DigitalOcean Provider
provider "digitalocean" {
  token = var.do_token
}

data "digitalocean_project" "go_count" {
  name = "go-count"
}

resource "digitalocean_project_resources" "go_count" {
  project = data.digitalocean_project.go_count.id
  resources = [
    digitalocean_kubernetes_cluster.test.urn
  ]
}

resource "digitalocean_kubernetes_cluster" "test" {
  name    = "test"
  region  = "ams3"
  version = "1.21.2-do.2"

  node_pool {
    name       = "default"
    size       = "s-1vcpu-2gb"
    node_count = 1
  }

}

output "k8s_host" {
  value = digitalocean_kubernetes_cluster.test.endpoint
}
