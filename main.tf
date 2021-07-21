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

resource "digitalocean_droplet" "web" {
  name = "test"
  size = "s-1vcpu-2gb"
  region = "fra1"
  image = "ubuntu-18-04-x64"
}
resource "digitalocean_droplet" "ams3" {
  name = "test"
  size = "s-1vcpu-2gb"
  region = "ams3"
  image = "ubuntu-18-04-x64"
}
resource "digitalocean_droplet" "lon1" {
  name = "test"
  size = "s-1vcpu-2gb"
  region = "lon1"
  image = "ubuntu-18-04-x64"
}
