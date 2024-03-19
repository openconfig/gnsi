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

## OpenConfig Data models for gNSI credentialz

Yang data models for certz are defined in the [OpenConfig public repository(https://github.com/openconfig/public/tree/master/release/models/gnsi)].  Documentation for OpenConfig including searchable list of paths and tree representations are at [OpenConfig.net](https://openconfig.net/projects/models/)
