provider "kubernetes" {
  host                   = digitalocean_kubernetes_cluster.test.kube_config.0.host
  cluster_ca_certificate = base64decode(digitalocean_kubernetes_cluster.test.kube_config.0.cluster_ca_certificate)
  token                  = digitalocean_kubernetes_cluster.test.kube_config.0.token
}

provider "helm" {
  kubernetes {
    host                   = digitalocean_kubernetes_cluster.test.kube_config.0.host
    cluster_ca_certificate = base64decode(digitalocean_kubernetes_cluster.test.kube_config.0.cluster_ca_certificate)
    token                  = digitalocean_kubernetes_cluster.test.kube_config.0.token
  }

}

resource "helm_release" "ingress_nginx" {
  name = "ingress-nginx"

  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"

  namespace = "ingress-nginx-system"

  create_namespace = true

  force_update = true

  lint = true

  # https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/Chart.yaml#L5
  version = "3.34.0"

  # https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/values.yaml
  set {
    name  = "controller.replicaCount"
    value = 1
  }
  set {
    name  = "controller.service.type"
    value = "NodePort"
  }
  set {
    name  = "controller.service.nodePorts.http"
    value = "30080" # This is static because I'm unable to receive node port
  }
  set {
    name  = "controller.service.nodePorts.https"
    value = "30443" # This is static because I'm unable to receive node port
  }
}

data "digitalocean_droplet" "k8s_node" {
  id = digitalocean_kubernetes_cluster.test.node_pool.0.nodes.0.droplet_id
}

output "k8s_ingress_endpoint_ip" {
  value = data.digitalocean_droplet.k8s_node.ipv4_address
}

## In case you want to pay extra for LB droplet
# data "kubernetes_service" "ingress_nginx_controller_lb" {
#   metadata {
#     name      = "ingress-nginx-controller"
#     namespace = "ingress-nginx-system"
#   }

#   depends_on = [
#     digitalocean_kubernetes_cluster.test,
#     helm_release.ingress_nginx
#   ]
# }


# output "k8s_ingress_endpoint_ip" {
#   value = data.kubernetes_service.ingress_nginx_controller_lb.status.0.load_balancer.0.ingress.0.ip
# }

resource "digitalocean_domain" "k8s_ingress_domain" {
  name       = "ingress.test-k8s.vejlupek.org"
  ip_address = data.digitalocean_droplet.k8s_node.ipv4_address
  depends_on = [
    data.digitalocean_droplet.k8s_node
  ]
}
