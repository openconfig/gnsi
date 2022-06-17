# How-To Guide to gNSI.ssh API

## Valid authentication methods

* password

* public key

* certificate

## Password-based

> **_NOTE:_**  The method is strongly discouraged.

The password cannot be set using gNSI.ssh API. The below-provided information is
just for completeness of this guide.

There are two methods to configure a password:

1) Directly on the device.

To change password execute the following command after logging-in to the device
using `ssh` or directly using a RS232 (or similar method):

```
$ echo "TeStP_w0rD" | passwd ${system_role} --stdin
```

2) Using gNSI.console API

> **_NOTE_** DOES NOT EXIST YET!

## Public key-based

In the case of public key based authentication users are authenticated by:
* username
* SSH public key

Provided `username` is checked against the list of known user names that are
stored in `${system_role_home}/.ssh/authorized_users` file.

Provided credentials are checked with the known to the target device public
keys that are stored in `${system_role_home}/.ssh/authorized_keys`

### How it works?

1) Create target's keys (public and private)

```
$ ssh-keygen -f target
```

2) Install required files on the target.
* target.pub (target's public key)
* target (target's private key)

`sshd_config`
```
PasswordAuthentication no
AuthenticationMethods  publickey
PubkeyAuthentication   yes

HostKey /etc/ssh/target
```

3) Create client's keys (public anf private)

```
$ ssh-keygen
```
4) Install client's public key on the target

```
$ ssh-copy-id ${system_role}@target
```

## Certificate-based

### How it works?

1) Create CA keys (public and private)

```
$ ssh-keygen -f CA
```

2) Create target's keys (public and private)

```
$ ssh-keygen -f target
```

3) Sign target's public key with CA's private key

```
$ ssh-keygen -h -s CA -n device.corp.company.com -I ID -V +52w target.pub
```

4) Install required files on the target.
* CA.pub (CA's public key)
* target-cert.pub (target's certificate)
* target.pub (target's public key)
* target (target's private key)

`sshd_config`
```
TrustedUserCAKeys /etc/ssh/CA.pub

HostCertificate /etc/ssh/target-cert.pub
HostKey /etc/ssh/target
```

5) Create client's keys (public anf private)

```
$ ssh-keygen -f client
```

6) Sign client's public key with CA's private key

```
$ ssh-keygen -s CA -I ID -n USERNAME -V +52w client.pub
```

7) Install required files on the client.
* CA.pub (CA's public key)
* client-cert.pub (client's certificate)
* client.pub (client's public key)
* client (client's private key)

`known_hosts`
```
@cert-authority *.company.com <content of CA.pub>
```

## Compromised keys/certificates revocation

OpenSSH server allows for revoking keys/certificates.

To enable this feature first the following lines has to be added to `/etc/ssh/ssd_config` file:
```
# Revoked user keys.
# This can include serials/key IDs for certificates.
RevokedKeys /etc/ssh/sshd_revoked_keys

```

There are two options for the `/etc/ssh/sshd_revoked_keys` file:
* simple text file
* OpenSSH Key Revocation List (KRL)

KRL is a custom binary format to manage revoked keys that is much smaller than the simple text file.

### How it works?

1) Create empty Key Revocation List (KRL) file.

```
$ ssh-keygen -k -f /etc/ssh/sshd_revoked_keys
```

2) Create a key to be revoked
```
$ ssh-keygen -f client
```

3) Revoke client's key.

```
$ ssh-keygen -k -u -f /etc/ssh/sshd_revoked_keys client.pub
```

4) Verify that key was revoked.

```
$ ssh-keygen -Q -f /etc/ssh/sshd_revoked_keys client.pub
```

The result should look like below:
```
client.pub (client.pub): REVOKED
```

The key will be revoked from now on and every attempt to use it will be logged in `auth.log` log file.