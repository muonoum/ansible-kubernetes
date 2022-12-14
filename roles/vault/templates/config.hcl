storage "raft" {
  path = "/var/lib/vault"
  node_id = "{{ node_name }}"
}

listener "tcp" {
  address = "{{ node_address }}:8200"
  cluster_address = "{{ node_address }}:8201"
  tls_cert_file = "/etc/vault/server.chain"
  tls_key_file = "/etc/vault/server.key"
  tls_min_version = "tls13"
  tls_client_ca_file = "/etc/vault/root.crt"
}

cluster_name = "{{ cluster_name }}"
api_addr = "https://vault:8200"
cluster_addr = "https://vault:8201"
ui = true
