terraform {
  required_providers {
    bindrole = {
      source = "techjavelin/bindrole"
    }
  }
}

provider "bindrole" {}

resource "bindrole_host" "netsrv-01" {
  domain   = "my.domain.tld"
  hostname = "netsrv-01"
  ip       = "10.0.0.1"
  aliases = [
    "ns-01",
    "dhcp-01"
  ]
}