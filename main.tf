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
