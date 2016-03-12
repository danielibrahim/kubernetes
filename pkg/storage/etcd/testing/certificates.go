/*
Copyright 2014 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testing

// You can use cfssl tool to generate certificates, please refer
// https://github.com/coreos/etcd/tree/master/hack/tls-setup for more details.
const CAFileContent = `
-----BEGIN CERTIFICATE-----
MIIDADCCAoWgAwIBAgIUPyoDaJMWija/6scZvsZIcPjyjqswCgYIKoZIzj0EAwMw
gawxCzAJBgNVBAYTAlVTMSowKAYDVQQKEyFIb25lc3QgQWNobWVkJ3MgVXNlZCBD
ZXJ0aWZpY2F0ZXMxKTAnBgNVBAsTIEhhc3RpbHktR2VuZXJhdGVkIFZhbHVlcyBE
aXZpc29uMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRMwEQYDVQQIEwpDYWxpZm9y
bmlhMRkwFwYDVQQDExBBdXRvZ2VuZXJhdGVkIENBMB4XDTE2MDIyNTEwMzYwMFoX
DTIxMDIyMzEwMzYwMFowgawxCzAJBgNVBAYTAlVTMSowKAYDVQQKEyFIb25lc3Qg
QWNobWVkJ3MgVXNlZCBDZXJ0aWZpY2F0ZXMxKTAnBgNVBAsTIEhhc3RpbHktR2Vu
ZXJhdGVkIFZhbHVlcyBEaXZpc29uMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRMw
EQYDVQQIEwpDYWxpZm9ybmlhMRkwFwYDVQQDExBBdXRvZ2VuZXJhdGVkIENBMHYw
EAYHKoZIzj0CAQYFK4EEACIDYgAEUjwvCgZPvo11v8qf+rBQRdYAt6IdMZJd3t1g
DjSySvBc3b+1WmCe911D/Zdwm/s83M4EQm8qUMeMvt60IqKIBZ6BDBbZdqQRycxJ
DuQwgyyYHQl5G52EDSJx//U1OkrOo2YwZDAOBgNVHQ8BAf8EBAMCAQYwEgYDVR0T
AQH/BAgwBgEB/wIBAjAdBgNVHQ4EFgQUp0oCeNg5O4HvABVa7/iJuLDkjAwwHwYD
VR0jBBgwFoAUp0oCeNg5O4HvABVa7/iJuLDkjAwwCgYIKoZIzj0EAwMDaQAwZgIx
AMuY6J2q53uFus7mZTEfWERXoUrTSvj2DEV+6MrmGD8VW2YaTwIGM0qzKlamb1QJ
rQIxAKtbXrfYzAjKBnrhdLD0kgf06pTQkIqBHj4zLen2K4NnVJWCSsKMua8FG+zP
jqvi0Q==
-----END CERTIFICATE-----
`
const CertFileContent = `
-----BEGIN CERTIFICATE-----
MIIC0zCCAlmgAwIBAgIUHXuZ3c9/pnUh2AOW3iiypCQpYEMwCgYIKoZIzj0EAwMw
gawxCzAJBgNVBAYTAlVTMSowKAYDVQQKEyFIb25lc3QgQWNobWVkJ3MgVXNlZCBD
ZXJ0aWZpY2F0ZXMxKTAnBgNVBAsTIEhhc3RpbHktR2VuZXJhdGVkIFZhbHVlcyBE
aXZpc29uMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRMwEQYDVQQIEwpDYWxpZm9y
bmlhMRkwFwYDVQQDExBBdXRvZ2VuZXJhdGVkIENBMB4XDTE2MDIyNTEwMzYwMFoX
DTE3MDIyNDEwMzYwMFowVTEWMBQGA1UEChMNYXV0b2dlbmVyYXRlZDEVMBMGA1UE
CxMMZXRjZCBjbHVzdGVyMRUwEwYDVQQHEwx0aGUgaW50ZXJuZXQxDTALBgNVBAMT
BGV0Y2QwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAAQ9HJgNWxMIrnns2+Sb8FUj9RBA
Fk/qP9cExp+FmbnjnOUy2poK5pGDdr88TMUAXyzV7J/rbTo6pDmWLWMEcbIgqfWY
W6BRmAaPWuQNLsP/L2k2N3NHvHZfCZK+efuDCGGjgZEwgY4wDgYDVR0PAQH/BAQD
AgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMBAf8EAjAA
MB0GA1UdDgQWBBQCFVaETH9zYZaw5OTHh8uuq4K84DAfBgNVHSMEGDAWgBSnSgJ4
2Dk7ge8AFVrv+Im4sOSMDDAPBgNVHREECDAGhwR/AAABMAoGCCqGSM49BAMDA2gA
MGUCMQC9+42ftFzkhujW6lTKZhf/rW3IyNmm5jXTS+RCPPB7jMqkH4Llq2PWgBR8
YQ8keX4CMGb6h8CLsgmFeVjpRjBTBuFSPSf3DQPPG2BkxPhtFek31g9lpSOjkioZ
fKv0Kz+tUw==
-----END CERTIFICATE-----
`
const KeyFileContent = `
-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDBkmx3mD+yd/qh6WYBTUAFbHZLHKrBv6o4H2AnSfx2HiMQoPm+elwhR
xhWa/tV+8zCgBwYFK4EEACKhZANiAAQ9HJgNWxMIrnns2+Sb8FUj9RBAFk/qP9cE
xp+FmbnjnOUy2poK5pGDdr88TMUAXyzV7J/rbTo6pDmWLWMEcbIgqfWYW6BRmAaP
WuQNLsP/L2k2N3NHvHZfCZK+efuDCGE=
-----END EC PRIVATE KEY-----
`