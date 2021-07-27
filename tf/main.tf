terraform {

  # For debbuging purposes it's simpler to use local state file
  # backend "remote" {
  #   organization = "elmariofredo"
  #   workspaces {
  #     name = "go-count"
  #   }
  # }

  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "2.10.1"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "2.2.0"
    }
  }
}
