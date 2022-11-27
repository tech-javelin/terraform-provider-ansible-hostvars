resource "bindrole_zone" "mydomain_com" {
    name = "mydomain.com"
    hostmaster_email = "admin"
    networks = [ "192.168.0" ]
    mail_servers = [
        {
            name = "mail1.mydomain.com"
            preference = 10
        },
        {
            name = "mail2.mydomain.com"
            preference = 20
        }
    ]
    name_servers = [ "ns1", "ns2" ]
    primaries = [ "192.168.0.11" ]
}