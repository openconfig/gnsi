# gNSI.credentialz

## Bootstrap / Assumptions

The `gNSI.credentialz` API allows changing existing SSH credentials.
Therefore credentials should be set up before credential RPCs are executed.

The following files are expected to be created during the bootstrap process:

* Certificate Authority's public key
  * required for certificate-based client authentication
  * used to check if the client's certificate is valid
* target's certificate
  * required for remote (this) host authentication by the clients
  * presented to the clients who validate it using CA's public key
* target's public key
  * always required
* target's private key
  * always required

## Console access authentication

### Using ``gNSI.credentialz`` API

* Start streaming RPC call to the target device.

```go
stream := RotateAccountCredentials()
```
* Send a password change request message to the target device.

```go
stream.Send(
    RotateAccountCredentialsRequest {
        password: PasswordRequest {
            accounts: Account {
                account: "user",
                password: Password {
                    value: {
                        plaintext: "password",
                    }
                },
                version: "v1.0",
                created_on: 3214451134,
            }
        }
    }
)

resp := stream.Receive()
```

* Check if the new password 'works'

* Finalize the operation

```go
stream.Send(
    RotateAccountCredentialsRequest {
        finalize: FinalizeRequest {}
    }
)
```

## SSH authentication

There are three authentication methods used with SSH:

* password
* public key
* certificate

### Method 1: Password-based

> **_NOTE:_**  The method is strongly discouraged.

Check out the ["Console access authentication"](#console-access-authentication)
section for information how to change account's password.

### Method 2: Public key-based

In the case of public key based authentication users are authenticated by:

* `username`
* SSH public key

#### Update the client's credentials

##### Update the client's authorized key

* Start streaming RPC call to the target device.

```go
stream := RotateAccountCredentials()
```

* Send a authorized keys change request message to the target device.

> **_NOTE:_**  The current list of authorized keys will be **replaced**.

```go
stream.Send(
    RotateAccountCredentialsRequest {
        credential: AuthorizedKeysRequest {
            credentials: AccountCredentials {
                account: "user",
                authorized_keys: AuthorizedKey {
                    authorized_key: "A....=",
                },
                authorized_keys: AuthorizedKey {
                    authorized_key: "A....=",
                },
                version: "v1.0",
                created_on: 3214451134,
            }
        }
    }
)

resp := stream.Receive()
```

* Check if the new SSH keys 'work'

* Finalize the operation

```go
stream.Send(
    RotateAccountCredentialsRequest {
        finalize: FinalizeRequest {}
    }
)
```

#### Update the host's keys with externally created keys

* Start streaming RPC call to the target device.

```go
stream := RotateHostParameters()
```

* Send a server's keys change request message to the target device. The keys must be base64 encoded.

```go
stream.Send(
    RotateHostParametersRequest {
        server_keys: ServerKeysRequest {
            auth_artifacts: []AuthenticationArtifacts{
                private_key: []bytes("...."),
            },
            version: "v1.0",
            created_on: 3214451134,
        }
    }
)

resp := stream.Receive()
```

* Check if the new keys 'work'

* Finalize the operation

```go
stream.Send(
    RotateHostParametersResponse {
        finalize: FinalizeRequest {}
    }
)
```

#### Update the host's keys with generated keys

* Start streaming RPC call to the target device.

```go
stream := RotateHostParameters()
```

* Send a server's keys change request message to the target device. The bytes are expected to be base64 encoded.

```go
stream.Send(
    RotateHostParametersRequest {
        generate_keys: GenerateKeysRequest{
            key_params: KEY_GEN_SSH_KEY_TYPE_RSA_4096,
        }
    }
)
resp, err := stream.Receive()
```

* Check if the new keys 'work'

* Finalize the operation

```go
stream.Send(
    RotateHostParametersResponse {
        finalize: FinalizeRequest {}
    }
)
```

### Method 3: Certificate-based

In this method both ends of the connection present a certificate signed by
the Certificate Authority.
This method is better than the key-based one as both the client and the server
can verify the credentials of the remote side and certificates can expire.

For this method to work the target's server has to have configured:

* Certificate Authority public keys allowed to sign a client's certificate
* A SSH host certificate singed by a Certificate Authority trusted by the client
* server's private key

Similarly, the client has to have the following:

* Certificate Authority public key of the CA that has signed
  the servers's certificate
* A SSH certificate singed by a Certificate Authority trusted by the server
* client's private key

#### Update the CA keys

* Start streaming RPC call to the target device.

```go
stream := RotateHostParameters()
```

* Send a CA key change request message to the target device.

```go
stream.Send(
    RotateHostParametersRequest {
        ssh_ca_public_key: CaPublicKeyRequest {
            ssh_ca_public_keys: "A....=",
            version: "v1.0",
            created_on: 3214451134,
        }
    }
)

resp := stream.Receive()
```

* Check if the new CA key 'works'

* Finalize the operation

```go
stream.Send(
    RotateHostParametersResponse {
        finalize: FinalizeRequest {}
    }
)
```

#### Update the host's certificate

* Start streaming RPC call to the target device.

```go
stream := RotateHostParameters()
```

* Send a server's certificate change request message to the target device. The bytes must be base64 encoded.

```go
stream.Send(
    RotateHostParametersRequest {
        server_keys: ServerKeysRequest {
            auth_artifacts: []AuthenticationArtifacts{
                certificate: []bytes("...."),
            },
            version: "v1.0",
            created_on: 3214451134,
        }
    }
)

resp := stream.Receive()
```

* Check if the new certificate 'works'

* Finalize the operation

```go
stream.Send(
    RotateHostParametersResponse {
        finalize: FinalizeRequest {}
    }
)
```

##### Update the account's authorized `principal` list

* Start streaming RPC call to the target device.

```go
stream := RotateAccountCredentials()
```

* Send a authorized `principal` list change request message to the target device.

> **_NOTE:_**  The current list of authorized `principal`s will be **replaced**.

```go
stream.Send(
    RotateAccountCredentialsRequest {
        user: AuthorizedUsersRequest {
            policies: UserPolicy {
                account: "user",
                authorized_principals: SshAuthorizedPrincipal {
                    authorized_user: "alice",
                },
                authorized_principals: SshAuthorizedPrincipal {
                    authorized_user: "bob",
                },
                version: "v1.0",
                created_on: 3214451134,
            }
        }
    }
)

resp := stream.Receive()
```

* Check if the new list of authorized `principal`s 'works'

* Finalize the operation

```go
stream.Send(
    RotateAccountCredentialsRequest {
        finalize: FinalizeRequest {}
    }
)
```

### Setting Allowed Authentication Types

The default sshd configuration generally allows for password, public key, and
keyboard interactive authentication types. Certificate authentication is implied
by way of setting a TrustedUserCaKeys file. In order to globally disable
specific types, credentialz provides the `AllowedAuthenticationRequest`. Rather
than operating with sshd defaults, this allows the operator to specify which
authentication types are globally permissable.

* Set the list of allowed authentication types.

```go
stream.Send(
    RotateHostParametersRequest {
        authentication_allowed: AllowedAuthenticationRequest {
            authentication_types: AuthenticationType {
                AuthenticationType_PUBKEY.Enum(),
            }
        }
    }
)
```

* Validate that new settings are working as expected.

* Finalize request.

```go
stream.Send(
    RotateHostParametersResponse {
        finalize: FinalizeRequest {}
    }
)
```

### Setting AuthorizedPrincipalsCommand

OpenSSH allows for the use of an tool which can dynamically return the list of
authorized principals for a given system role. This is a global setting and
cannot be set at the same time as the role specific configuration
`authorized_principals` in the `UserPolicy`.

* Set the AuthorizedPrincipalsCommand tool

```go
stream.Send(
    RotateHostParametersRequest {
        authorized_principal_check: AuthorizedPrincipalCheckRequest {
            tool: Tool_TOOL_HIBA_DEFAULT.Enum(),
        }
    }
)
```

* Validate that new settings are working as expected.

* Finalize request.

```go
stream.Send(
    RotateHostParametersResponse {
        finalize: FinalizeRequest {}
    }
)
```

## User Journeys

### Rotate Certificate based on existing key

The most common operation we are expecting to require on devices is the rotation of certificates used for SSH access for devices. This operation expects to reuse the existing host key on the device.

* Get the public key configured on the host.

```go

resp, err := c.GetPublicKeys(&GetPublicKeyRequests{})
```

* Generate certificate based on key.

* Rotate certificate on device.

```go
stream.Send(
    RotateHostParametersRequest {
        server_keys: ServerKeysRequest {
            certificate: "A....=",
            version: "v1.0",
            created_on: 3214451134,
        }
    }
)
```

* Validate that new settings are working as expected.

* Finalize request.

```go
stream.Send(
    RotateHostParametersResponse {
        finalize: FinalizeRequest {}
    }
)
```

### Generate new host key on device and rotate certificate based on the new key

This use case focuses on the rotation of a host key and then generation of the certificate based on the new public key.

* Send request for generation of new private key.

```go
stream.Send(
    RotateHostParametersRequest {
        generate_keys: []GenerateKeysRequest {{
            key_params: KeyGen.KEY_GEN_SSH_KEY_TYPE_EDDSA_ED25519 
        }}
    }
)
```

* Get Response containing public key to generate the certificate.

```go
resp, err := stream.Recv()
data := resp.PublicKeys
```

* The caller will then use this data to generate a certificate.

* Send generated cert to device to rotate.

```go
stream.Send(
    RotateHostParametersRequest {
        server_keys: ServerKeysRequest {
            certificate: "A....=",
            version: "v1.0",
            created_on: 3214451134,
        }
    }
)
```

* Validate the `RotateCredentialsResponse`.

```go
if _, err := stream.Recv(); err != nil {
    ...
}
```

* Validate that new settings are working as expected.

* Finalize request

```go
stream.Send(
    RotateHostParametersResponse {
        finalize: FinalizeRequest {}
    }
)
```

## gNSI.credentialz Telemetry Extension

### `gnsi-credentialz.yang`

An overview of the changes defined in the `gnsi-credentialz.yang` file are
shown below.

```txt
module: gnsi-credentialz

  augment /oc-sys:system:
    +--rw console
       +--rw config
       +--ro state
          +--ro counters
             +--ro access-rejects?       oc-yang:counter64
             +--ro last-access-reject?   oc-types:timeticks64
             +--ro access-accepts?       oc-yang:counter64
             +--ro last-access-accept?   oc-types:timeticks64
  augment /oc-sys:system/oc-sys:ssh-server/oc-sys:state:
    +--ro active-trusted-user-ca-keys-version?      version
    +--ro active-trusted-user-ca-keys-created-on?   created-on
    +--ro active-host-certificate-version?          version
    +--ro active-host-certificate-created-on?       created-on
    +--ro active-host-key-version?                  version
    +--ro active-host-key-version-created-on?       created-on
    +--ro counters
       +--ro access-rejects?       oc-yang:counter64
       +--ro last-access-reject?   oc-types:timeticks64
       +--ro access-accepts?       oc-yang:counter64
       +--ro last-access-accept?   oc-types:timeticks64
  augment /oc-sys:system/oc-sys:aaa/oc-sys:authentication/oc-sys:users/oc-sys:user/oc-sys:state:
    +--ro password-version?                   version
    +--ro password-created-on?                created-on
    +--ro authorized-principals-list-version?      version
    +--ro authorized-principals-list-created-on?   created-on
    +--ro authorized-keys-list-version?       version
    +--ro authorized-keys-list-created-on?    created-on
```

### `openconfig-system` tree

The `openconfig-system` subtree after augments defined in the
`gnsi-credentialz.yang` file is shown below.

For interactive version click [here](gnsi-credentialz.html).

<details>
<summary>
The diagram of the tree.
</summary>

```txt
module: openconfig-system
  +--rw system
     +--rw config
     |  +--rw hostname?       oc-inet:domain-name
     |  +--rw domain-name?    oc-inet:domain-name
     |  +--rw login-banner?   string
     |  +--rw motd-banner?    string
     +--ro state
     |  +--ro hostname?           oc-inet:domain-name
     |  +--ro domain-name?        oc-inet:domain-name
     |  +--ro login-banner?       string
     |  +--ro motd-banner?        string
     |  +--ro current-datetime?   oc-yang:date-and-time
     |  +--ro boot-time?          oc-types:timeticks64
     +--rw clock
     |  +--rw config
     |  |  +--rw timezone-name?   timezone-name-type
     |  +--ro state
     |     +--ro timezone-name?   timezone-name-type
     +--rw dns
     |  +--rw config
     |  |  +--rw search*   oc-inet:domain-name
     |  +--ro state
     |  |  +--ro search*   oc-inet:domain-name
     |  +--rw servers
     |  |  +--rw server* [address]
     |  |     +--rw address    -> ../config/address
     |  |     +--rw config
     |  |     |  +--rw address?   oc-inet:ip-address
     |  |     |  +--rw port?      oc-inet:port-number
     |  |     +--ro state
     |  |        +--ro address?   oc-inet:ip-address
     |  |        +--ro port?      oc-inet:port-number
     |  +--rw host-entries
     |     +--rw host-entry* [hostname]
     |        +--rw hostname    -> ../config/hostname
     |        +--rw config
     |        |  +--rw hostname?       string
     |        |  +--rw alias*          string
     |        |  +--rw ipv4-address*   oc-inet:ipv4-address
     |        |  +--rw ipv6-address*   oc-inet:ipv6-address
     |        +--ro state
     |           +--ro hostname?       string
     |           +--ro alias*          string
     |           +--ro ipv4-address*   oc-inet:ipv4-address
     |           +--ro ipv6-address*   oc-inet:ipv6-address
     +--rw ntp
     |  +--rw config
     |  |  +--rw enabled?              boolean
     |  |  +--rw ntp-source-address?   oc-inet:ip-address
     |  |  +--rw enable-ntp-auth?      boolean
     |  +--ro state
     |  |  +--ro enabled?              boolean
     |  |  +--ro ntp-source-address?   oc-inet:ip-address
     |  |  +--ro enable-ntp-auth?      boolean
     |  |  +--ro auth-mismatch?        oc-yang:counter64
     |  +--rw ntp-keys
     |  |  +--rw ntp-key* [key-id]
     |  |     +--rw key-id    -> ../config/key-id
     |  |     +--rw config
     |  |     |  +--rw key-id?      uint16
     |  |     |  +--rw key-type?    identityref
     |  |     |  +--rw key-value?   string
     |  |     +--ro state
     |  |        +--ro key-id?      uint16
     |  |        +--ro key-type?    identityref
     |  |        +--ro key-value?   string
     |  +--rw servers
     |     +--rw server* [address]
     |        +--rw address    -> ../config/address
     |        +--rw config
     |        |  +--rw address?            oc-inet:host
     |        |  +--rw port?               oc-inet:port-number
     |        |  +--rw version?            uint8
     |        |  +--rw association-type?   enumeration
     |        |  +--rw iburst?             boolean
     |        |  +--rw prefer?             boolean
     |        +--ro state
     |           +--ro address?            oc-inet:host
     |           +--ro port?               oc-inet:port-number
     |           +--ro version?            uint8
     |           +--ro association-type?   enumeration
     |           +--ro iburst?             boolean
     |           +--ro prefer?             boolean
     |           +--ro stratum?            uint8
     |           +--ro root-delay?         uint32
     |           +--ro root-dispersion?    uint64
     |           +--ro offset?             uint64
     |           +--ro poll-interval?      uint32
     +--rw ssh-server
     |  +--rw config
     |  |  +--rw enable?             boolean
     |  |  +--rw protocol-version?   enumeration
     |  |  +--rw timeout?            uint16
     |  |  +--rw rate-limit?         uint16
     |  |  +--rw session-limit?      uint16
     |  +--ro state
     |     +--ro enable?                                              boolean
     |     +--ro protocol-version?                                    enumeration
     |     +--ro timeout?                                             uint16
     |     +--ro rate-limit?                                          uint16
     |     +--ro session-limit?                                       uint16
     |     +--ro gnsi-credz:active-trusted-user-ca-keys-version?      version
     |     +--ro gnsi-credz:active-trusted-user-ca-keys-created-on?   created-on
     |     +--ro gnsi-credz:active-host-certificate-version?          version
     |     +--ro gnsi-credz:active-host-certificate-created-on?       created-on
     |     +--ro gnsi-credz:active-host-key-version?                  version
     |     +--ro gnsi-credz:active-host-key-version-created-on?       created-on
     |     +--ro gnsi-credz:counters
     |        +--ro gnsi-credz:access-rejects?       oc-yang:counter64
     |        +--ro gnsi-credz:last-access-reject?   oc-types:timeticks64
     |        +--ro gnsi-credz:access-accepts?       oc-yang:counter64
     |        +--ro gnsi-credz:last-access-accept?   oc-types:timeticks64
     +--rw telnet-server
     |  +--rw config
     |  |  +--rw enable?          boolean
     |  |  +--rw timeout?         uint16
     |  |  +--rw rate-limit?      uint16
     |  |  +--rw session-limit?   uint16
     |  +--ro state
     |     +--ro enable?          boolean
     |     +--ro timeout?         uint16
     |     +--ro rate-limit?      uint16
     |     +--ro session-limit?   uint16
     +--rw logging
     |  +--rw console
     |  |  +--rw config
     |  |  +--ro state
     |  |  +--rw selectors
     |  |     +--rw selector* [facility severity]
     |  |        +--rw facility    -> ../config/facility
     |  |        +--rw severity    -> ../config/severity
     |  |        +--rw config
     |  |        |  +--rw facility?   identityref
     |  |        |  +--rw severity?   syslog-severity
     |  |        +--ro state
     |  |           +--ro facility?   identityref
     |  |           +--ro severity?   syslog-severity
     |  +--rw remote-servers
     |     +--rw remote-server* [host]
     |        +--rw host         -> ../config/host
     |        +--rw config
     |        |  +--rw host?             oc-inet:host
     |        |  +--rw source-address?   oc-inet:ip-address
     |        |  +--rw remote-port?      oc-inet:port-number
     |        +--ro state
     |        |  +--ro host?             oc-inet:host
     |        |  +--ro source-address?   oc-inet:ip-address
     |        |  +--ro remote-port?      oc-inet:port-number
     |        +--rw selectors
     |           +--rw selector* [facility severity]
     |              +--rw facility    -> ../config/facility
     |              +--rw severity    -> ../config/severity
     |              +--rw config
     |              |  +--rw facility?   identityref
     |              |  +--rw severity?   syslog-severity
     |              +--ro state
     |                 +--ro facility?   identityref
     |                 +--ro severity?   syslog-severity
     +--rw aaa
     |  +--rw config
     |  +--ro state
     |  +--rw authentication
     |  |  +--rw config
     |  |  |  +--rw authentication-method*   union
     |  |  +--ro state
     |  |  |  +--ro authentication-method*   union
     |  |  +--rw admin-user
     |  |  |  +--rw config
     |  |  |  |  +--rw admin-password?          string
     |  |  |  |  +--rw admin-password-hashed?   oc-aaa-types:crypt-password-type
     |  |  |  +--ro state
     |  |  |     +--ro admin-password?          string
     |  |  |     +--ro admin-password-hashed?   oc-aaa-types:crypt-password-type
     |  |  |     +--ro admin-username?          string
     |  |  +--rw users
     |  |     +--rw user* [username]
     |  |        +--rw username    -> ../config/username
     |  |        +--rw config
     |  |        |  +--rw username?   string
     |  |        |  +--rw role?       union
     |  |        +--ro state
     |  |           +--ro username?                                      string
     |  |           +--ro password?                                      string
     |  |           +--ro password-hashed?                               oc-aaa-types:crypt-password-type
     |  |           +--ro role?                                          union
     |  |           +--ro gnsi-credz:password-version?                   version
     |  |           +--ro gnsi-credz:password-created-on?                created-on
     |  |           +--ro gnsi-credz:authorized-principals-list-version?      version
     |  |           +--ro gnsi-credz:authorized-principals-list-created-on?   created-on
     |  |           +--ro gnsi-credz:authorized-keys-list-version?       version
     |  |           +--ro gnsi-credz:authorized-keys-list-created-on?    created-on
     |  +--rw authorization
     |  |  +--rw config
     |  |  |  +--rw authorization-method*   union
     |  |  +--ro state
     |  |  |  +--ro authorization-method*   union
     |  |  +--rw events
     |  |     +--rw event* [event-type]
     |  |        +--rw event-type    -> ../config/event-type
     |  |        +--rw config
     |  |        |  +--rw event-type?   identityref
     |  |        +--ro state
     |  |           +--ro event-type?   identityref
     |  +--rw accounting
     |  |  +--rw config
     |  |  |  +--rw accounting-method*   union
     |  |  +--ro state
     |  |  |  +--ro accounting-method*   union
     |  |  +--rw events
     |  |     +--rw event* [event-type]
     |  |        +--rw event-type    -> ../config/event-type
     |  |        +--rw config
     |  |        |  +--rw event-type?   identityref
     |  |        |  +--rw record?       enumeration
     |  |        +--ro state
     |  |           +--ro event-type?   identityref
     |  |           +--ro record?       enumeration
     |  +--rw server-groups
     |     +--rw server-group* [name]
     |        +--rw name       -> ../config/name
     |        +--rw config
     |        |  +--rw name?   string
     |        |  +--rw type?   identityref
     |        +--ro state
     |        |  +--ro name?   string
     |        |  +--ro type?   identityref
     |        +--rw servers
     |           +--rw server* [address]
     |              +--rw address    -> ../config/address
     |              +--rw config
     |              |  +--rw name?      string
     |              |  +--rw address?   oc-inet:ip-address
     |              |  +--rw timeout?   uint16
     |              +--ro state
     |              |  +--ro name?                  string
     |              |  +--ro address?               oc-inet:ip-address
     |              |  +--ro timeout?               uint16
     |              |  +--ro connection-opens?      oc-yang:counter64
     |              |  +--ro connection-closes?     oc-yang:counter64
     |              |  +--ro connection-aborts?     oc-yang:counter64
     |              |  +--ro connection-failures?   oc-yang:counter64
     |              |  +--ro connection-timeouts?   oc-yang:counter64
     |              |  +--ro messages-sent?         oc-yang:counter64
     |              |  +--ro messages-received?     oc-yang:counter64
     |              |  +--ro errors-received?       oc-yang:counter64
     |              +--rw tacacs
     |              |  +--rw config
     |              |  |  +--rw port?                oc-inet:port-number
     |              |  |  +--rw secret-key?          oc-types:routing-password
     |              |  |  +--rw secret-key-hashed?   oc-aaa-types:crypt-password-type
     |              |  |  +--rw source-address?      oc-inet:ip-address
     |              |  +--ro state
     |              |     +--ro port?                oc-inet:port-number
     |              |     +--ro secret-key?          oc-types:routing-password
     |              |     +--ro secret-key-hashed?   oc-aaa-types:crypt-password-type
     |              |     +--ro source-address?      oc-inet:ip-address
     |              +--rw radius
     |                 +--rw config
     |                 |  +--rw auth-port?             oc-inet:port-number
     |                 |  +--rw acct-port?             oc-inet:port-number
     |                 |  +--rw secret-key?            oc-types:routing-password
     |                 |  +--rw secret-key-hashed?     oc-aaa-types:crypt-password-type
     |                 |  +--rw source-address?        oc-inet:ip-address
     |                 |  +--rw retransmit-attempts?   uint8
     |                 +--ro state
     |                    +--ro auth-port?             oc-inet:port-number
     |                    +--ro acct-port?             oc-inet:port-number
     |                    +--ro secret-key?            oc-types:routing-password
     |                    +--ro secret-key-hashed?     oc-aaa-types:crypt-password-type
     |                    +--ro source-address?        oc-inet:ip-address
     |                    +--ro retransmit-attempts?   uint8
     |                    +--ro counters
     |                       +--ro retried-access-requests?   oc-yang:counter64
     |                       +--ro access-accepts?            oc-yang:counter64
     |                       +--ro access-rejects?            oc-yang:counter64
     |                       +--ro timeout-access-requests?   oc-yang:counter64
     +--rw memory
     |  +--rw config
     |  +--ro state
     |     +--ro physical?   uint64
     |     +--ro reserved?   uint64
     +--ro cpus
     |  +--ro cpu* [index]
     |     +--ro index    -> ../state/index
     |     +--ro state
     |        +--ro index?                union
     |        +--ro total
     |        |  +--ro instant?    oc-types:percentage
     |        |  +--ro avg?        oc-types:percentage
     |        |  +--ro min?        oc-types:percentage
     |        |  +--ro max?        oc-types:percentage
     |        |  +--ro interval?   oc-types:stat-interval
     |        |  +--ro min-time?   oc-types:timeticks64
     |        |  +--ro max-time?   oc-types:timeticks64
     |        +--ro user
     |        |  +--ro instant?    oc-types:percentage
     |        |  +--ro avg?        oc-types:percentage
     |        |  +--ro min?        oc-types:percentage
     |        |  +--ro max?        oc-types:percentage
     |        |  +--ro interval?   oc-types:stat-interval
     |        |  +--ro min-time?   oc-types:timeticks64
     |        |  +--ro max-time?   oc-types:timeticks64
     |        +--ro kernel
     |        |  +--ro instant?    oc-types:percentage
     |        |  +--ro avg?        oc-types:percentage
     |        |  +--ro min?        oc-types:percentage
     |        |  +--ro max?        oc-types:percentage
     |        |  +--ro interval?   oc-types:stat-interval
     |        |  +--ro min-time?   oc-types:timeticks64
     |        |  +--ro max-time?   oc-types:timeticks64
     |        +--ro nice
     |        |  +--ro instant?    oc-types:percentage
     |        |  +--ro avg?        oc-types:percentage
     |        |  +--ro min?        oc-types:percentage
     |        |  +--ro max?        oc-types:percentage
     |        |  +--ro interval?   oc-types:stat-interval
     |        |  +--ro min-time?   oc-types:timeticks64
     |        |  +--ro max-time?   oc-types:timeticks64
     |        +--ro idle
     |        |  +--ro instant?    oc-types:percentage
     |        |  +--ro avg?        oc-types:percentage
     |        |  +--ro min?        oc-types:percentage
     |        |  +--ro max?        oc-types:percentage
     |        |  +--ro interval?   oc-types:stat-interval
     |        |  +--ro min-time?   oc-types:timeticks64
     |        |  +--ro max-time?   oc-types:timeticks64
     |        +--ro wait
     |        |  +--ro instant?    oc-types:percentage
     |        |  +--ro avg?        oc-types:percentage
     |        |  +--ro min?        oc-types:percentage
     |        |  +--ro max?        oc-types:percentage
     |        |  +--ro interval?   oc-types:stat-interval
     |        |  +--ro min-time?   oc-types:timeticks64
     |        |  +--ro max-time?   oc-types:timeticks64
     |        +--ro hardware-interrupt
     |        |  +--ro instant?    oc-types:percentage
     |        |  +--ro avg?        oc-types:percentage
     |        |  +--ro min?        oc-types:percentage
     |        |  +--ro max?        oc-types:percentage
     |        |  +--ro interval?   oc-types:stat-interval
     |        |  +--ro min-time?   oc-types:timeticks64
     |        |  +--ro max-time?   oc-types:timeticks64
     |        +--ro software-interrupt
     |           +--ro instant?    oc-types:percentage
     |           +--ro avg?        oc-types:percentage
     |           +--ro min?        oc-types:percentage
     |           +--ro max?        oc-types:percentage
     |           +--ro interval?   oc-types:stat-interval
     |           +--ro min-time?   oc-types:timeticks64
     |           +--ro max-time?   oc-types:timeticks64
     +--rw processes
     |  +--ro process* [pid]
     |     +--ro pid      -> ../state/pid
     |     +--ro state
     |        +--ro pid?                  uint64
     |        +--ro name?                 string
     |        +--ro args*                 string
     |        +--ro start-time?           oc-types:timeticks64
     |        +--ro cpu-usage-user?       oc-yang:counter64
     |        +--ro cpu-usage-system?     oc-yang:counter64
     |        +--ro cpu-utilization?      oc-types:percentage
     |        +--ro memory-usage?         uint64
     |        +--ro memory-utilization?   oc-types:percentage
     +--ro alarms
     |  +--ro alarm* [id]
     |     +--ro id        -> ../state/id
     |     +--ro config
     |     +--ro state
     |        +--ro id?             string
     |        +--ro resource?       string
     |        +--ro text?           string
     |        +--ro time-created?   oc-types:timeticks64
     |        +--ro severity?       identityref
     |        +--ro type-id?        union
     +--rw messages
     |  +--rw config
     |  |  +--rw severity?   oc-log:syslog-severity
     |  +--ro state
     |  |  +--ro severity?   oc-log:syslog-severity
     |  |  +--ro message
     |  |     +--ro msg?        string
     |  |     +--ro priority?   uint8
     |  |     +--ro app-name?   string
     |  |     +--ro procid?     string
     |  |     +--ro msgid?      string
     |  +--rw debug-entries
     |     +--rw debug-service* [service]
     |        +--rw service    -> ../config/service
     |        +--rw config
     |        |  +--rw service?   identityref
     |        |  +--rw enabled?   boolean
     |        +--ro state
     |           +--ro service?   identityref
     |           +--ro enabled?   boolean
     +--rw license
     |  +--rw licenses
     |     +--rw license* [license-id]
     |        +--rw license-id    -> ../config/license-id
     |        +--rw config
     |        |  +--rw license-id?     string
     |        |  +--rw license-data?   union
     |        |  +--rw active?         boolean
     |        +--ro state
     |           +--ro license-id?        string
     |           +--ro license-data?      union
     |           +--ro active?            boolean
     |           +--ro description?       string
     |           +--ro issue-date?        uint64
     |           +--ro expiration-date?   uint64
     |           +--ro in-use?            boolean
     |           +--ro expired?           boolean
     |           +--ro valid?             boolean
     +--rw oc-sys-grpc:grpc-servers
     |  +--rw oc-sys-grpc:grpc-server* [name]
     |     +--rw oc-sys-grpc:name      -> ../config/name
     |     +--rw oc-sys-grpc:config
     |     |  +--rw oc-sys-grpc:name?                      string
     |     |  +--rw oc-sys-grpc:services*                  identityref
     |     |  +--rw oc-sys-grpc:enable?                    boolean
     |     |  +--rw oc-sys-grpc:port?                      oc-inet:port-number
     |     |  +--rw oc-sys-grpc:transport-security?        boolean
     |     |  +--rw oc-sys-grpc:certificate-id?            string
     |     |  +--rw oc-sys-grpc:metadata-authentication?   boolean
     |     |  +--rw oc-sys-grpc:listen-addresses*          union
     |     |  +--rw oc-sys-grpc:network-instance?          oc-ni:network-instance-ref
     |     +--ro oc-sys-grpc:state
     |        +--ro oc-sys-grpc:name?                      string
     |        +--ro oc-sys-grpc:services*                  identityref
     |        +--ro oc-sys-grpc:enable?                    boolean
     |        +--ro oc-sys-grpc:port?                      oc-inet:port-number
     |        +--ro oc-sys-grpc:transport-security?        boolean
     |        +--ro oc-sys-grpc:certificate-id?            string
     |        +--ro oc-sys-grpc:metadata-authentication?   boolean
     |        +--ro oc-sys-grpc:listen-addresses*          union
     |        +--ro oc-sys-grpc:network-instance?          oc-ni:network-instance-ref
     +--rw gnsi-credz:console
        +--rw gnsi-credz:config
        +--ro gnsi-credz:state
           +--ro gnsi-credz:counters
              +--ro gnsi-credz:access-rejects?       oc-yang:counter64
              +--ro gnsi-credz:last-access-reject?   oc-types:timeticks64
              +--ro gnsi-credz:access-accepts?       oc-yang:counter64
              +--ro gnsi-credz:last-access-accept?   oc-types:timeticks64

```
</details>
