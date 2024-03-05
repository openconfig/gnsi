# gNSI.certz

## gNSI certz Service Protobuf Definition
**Contributors**: hines@google.com, morrowc@google.com, tmadejski@google.com
**Last Updated**: 2023-05-31

### Background

The certz service definition provides the API to be used for rotating and
testing PKI primitives used on network systems.
The `Rotate()` is bidirectional streaming RPC which permit
mutating Certificates, Root Certificate Bundles, Certificate Revocation
Lists and Authentication Policies. For `Rotate()` stream it is possible to
mutate one or more of the elements, and to send a `Finalize` message once the
in-flight change has been verified to be operational. Failure to send
the `Finalize` message will result in the candidate element being discarded
and the original element being used instead.

### Motivation

Management of the PKI elements for a network system should have
a clear and direct method for installation and update.

#### `Certz.Rotate()`

`Certz.Rotate()` will permit rotation, and
verification of function, of any of the PKI elements.
The normal use-case would be to:

* send an CertificateBundle to a network system as a
  `RotateCertificateRequest`.
* verify that the services which will use the new certificate bundle
  continue to operate normally.
* send a `FinalizeRequest` to finish the rotation process.

#### SSL profiles

SSL profiles logically group a certificate (private and public keys),
Certificate Authority chain of certificates (a.k.a. a CA trust bundle) and
a set of Certificate Revocation Lists into a set that then can be assigned
as a whole to a gRPC server.

There is always at least one profile present on a target - the `system_default_profile` which is vendor provided. This profile cannot be changed. If the use but when the `ssl_profile_id` field in the
`RotateCertificateRequest` message is not set (or set to an empty string) it
also refers this SSL profile. (This statement will be deprecated once all vendors standardize on the key name)

Profiles existing on a target can be discovered using the
`Certz.GetProfileList()` RPC.

A SSL profile can be added using the `Certz.AddProfile()` RPC.

When no longer a profile is needed it can be removed from the target via
`Certz.DeleteProfile()` RPC. Note that the gNxI SSL profile cannot be
removed.

The SSL profile ID of a gRPC server is exposed in the YANG leaf
`ssl-profile-id` which is an augment to the
`/oc-sys:system/oc-sys-grpc:grpc-servers/oc-sys-grpc:grpc-server/oc-sys-grpc:state`
container.

#### Authentication Policy

An authentication policy is a set of rules that defines which CA can be trusted
to sign certificates for which subjects. By rotating authentication policies,
data center admins can ensure all endpoints are updated to validate certificates
presented by their peers during mutual authentication are signed by one of
the authorized CAs in the authentication framework as specified in the policy.
This helps to minimize the impact of a security breach, as it prevents,
for example, an attacker from using a less privileged CA to sign for high value
users/roles.

##### Details

When a client tries to establish a gRPC connection to a gRPC server, the server
must verify that the client is authorized to do so. To do that the client
presents a certificate to the server, which the server verifies to see if it was
issued by a trusted Certificate Authority (CA).

In large scale PKI deployments consisting of multiple signing authorities
assigned to issue certificates of specific users/entities with their own key
hierarchies, use of an authentication policy is one solution to maintain
a single Trust Bundle across all applications. Use of one Trust Bundle
consisting of the root certificates of all signing authorities simplifies
maintenance and avoid endpoint application configuration complexities.

In such deployment, a centrally maintained authentication policy specifies which
signing authorities are permitted to issue certificates for which group of
users. In other words, after validating a connecting peer's certificates against
the Trust Bundle during a TLS handshake, the endpoint will also validate
the peer's and the certificate issuer's identities against the authentication
policy before accepting the connection.

### User Experiences

#### System default SSL profile

The system will always provide a default TLS profile that uses the IDevID cert.
This profile will always be available and cannot be changed. It should use the name
"system_default_profile".

An attempt to change or delete this profile will return an error.

The system will start with this profile and either bootz or enrollz will be responsible for creating an alternate profile during device turnup if those workflows are used.

#### Create a SSL profile

Call `Certz.AddProfile` RPC with the `ssl_profile_id` field specifying the ID
of the new SSL profile.
A new profile can choose to use existing artifacts from other profiles, via sending `Entity` messages with `ExistingEntity` set with the ssl_profile_id set to the source
profile to copy from.

#### Delete a SSL profile

Call `Certz.DeleteProfile` RPC with the `ssl_profile_id` field specifying the
ID of the SSL profile to be deleted.

#### List existing SSL profiles

Call `Certz.GetProfileList` RPC. The response will list all existing
SSL profiles.

#### A Certificate is to be rotated or updated

Create, and test, a new certificate and a private key.

Send that certificate, its private key and all required intermediate certificate
chain to the target network system in the `Certz.UploadRequest`'s
`entity.certificate_chain` field.

Verify that the certificate newly rotated is used by services
which require it.

Send a `Certz.FinalizeRequest` message to the `Certz.Rotate` RPC to close out
the action.

If the stream is disconnected prior to the `Finalize` message being
sent, the proposed configuration is rolled back automatically.

#### A Certificate is rotated, the session breaks before `FinalizeRequest`

Create a new certificate chain and a private key.

Send that certificate, its private key and all required intermediate certificate
chain to the target network system in the `Certz.UploadRequest`'s
`entity.certificate_chain` field.

Verify that the certificate newly deployed is usable by the relevant
services, that the services properly present the certificate upon
new service connections.

The connection to the network system is broken, there is no
`Certz.FinalizeRequest` sent.

The gNSI service rolls back the candidate and re-installs the original
certificate and associated private key.

#### An Authentication Policy is to be rotated or updated

Create a new authentication policy.

Send that authentication policy to the target network system in
the `Certz.UploadRequest`'s `entity.authentication_policy` field.

Verify that the authentication policy newly rotated is used by services
which require it.

Send a `Certz.FinalizeRequest` message to the `Certz.Rotate` RPC to close out
the action.

If the stream is disconnected prior to the `Certz.FinalizeRequest` message being
sent, the proposed authentication policy is rolled back automatically.

#### An Authentication Policy is rotated, the session breaks before `FinalizeRequest`

Create a new authentication policy.

Send that authentication policy to the target network system in
the `Certz.UploadRequest`'s `entity.authentication_policy` field.

Verify that the authentication policy newly deployed is usable by the relevant
services, that the services properly uses the authentication policy upon
new service connections.

The connection to the network system is broken, there is no
`Certz.FinalizeRequest` sent.

The gNSI service rolls back the candidate and re-installs the original
authentication policy.

### Open Questions/Considerations

None to date.

## gNSI.cert Telemetry Extension

### `gnsi-certz.yang`

An overview of the changes defined in the `gnsi-certz.yang` file are shown
below.

```txt
module: gnsi-certz

  augment /oc-sys:system/oc-sys-grpc:grpc-servers/oc-sys-grpc:grpc-server/oc-sys-grpc:state:
    +--ro certificate-version?                             version
    +--ro certificate-created-on?                          created-on
    +--ro ca-trust-bundle-version?                         version
    +--ro ca-trust-bundle-created-on?                      created-on
    +--ro certificate-revocation-list-bundle-version?      version
    +--ro certificate-revocation-list-bundle-created-on?   created-on
    +--ro authentication-policy-version?                   version
    +--ro authentication-policy-created-on?                created-on
    +--ro ssl-profile-id?                                  string
    +--ro counters
       +--ro connection-rejects?       oc-yang:counter64
       +--ro last-connection-reject?   oc-types:timeticks64
       +--ro connection-accepts?       oc-yang:counter64
       +--ro last-connection-accept?   oc-types:timeticks64
```

### `openconfig-system` tree

The  `openconfig-system` subtree after augments defined in the
`gnsi-certz.yang` file is shown below.

For interactive version click [here](gnsi-certz.html).

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
     |     +--ro enable?             boolean
     |     +--ro protocol-version?   enumeration
     |     +--ro timeout?            uint16
     |     +--ro rate-limit?         uint16
     |     +--ro session-limit?      uint16
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
     |  |        |  +--rw username?          string
     |  |        |  +--rw password?          string
     |  |        |  +--rw password-hashed?   oc-aaa-types:crypt-password-type
     |  |        |  +--rw ssh-key?           string
     |  |        |  +--rw role?              union
     |  |        +--ro state
     |  |           +--ro username?          string
     |  |           +--ro password?          string
     |  |           +--ro password-hashed?   oc-aaa-types:crypt-password-type
     |  |           +--ro ssh-key?           string
     |  |           +--ro role?              union
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
        +--rw oc-sys-grpc:grpc-server* [name]
           +--rw oc-sys-grpc:name      -> ../config/name
           +--rw oc-sys-grpc:config
           |  +--rw oc-sys-grpc:name?                      string
           |  +--rw oc-sys-grpc:services*                  identityref
           |  +--rw oc-sys-grpc:enable?                    boolean
           |  +--rw oc-sys-grpc:port?                      oc-inet:port-number
           |  +--rw oc-sys-grpc:transport-security?        boolean
           |  +--rw oc-sys-grpc:certificate-id?            string
           |  +--rw oc-sys-grpc:metadata-authentication?   boolean
           |  +--rw oc-sys-grpc:listen-addresses*          union
           |  +--rw oc-sys-grpc:network-instance?          oc-ni:network-instance-ref
           +--ro oc-sys-grpc:state
              +--ro oc-sys-grpc:name?                                           string
              +--ro oc-sys-grpc:services*                                       identityref
              +--ro oc-sys-grpc:enable?                                         boolean
              +--ro oc-sys-grpc:port?                                           oc-inet:port-number
              +--ro oc-sys-grpc:transport-security?                             boolean
              +--ro oc-sys-grpc:certificate-id?                                 string
              +--ro oc-sys-grpc:metadata-authentication?                        boolean
              +--ro oc-sys-grpc:listen-addresses*                               union
              +--ro oc-sys-grpc:network-instance?                               oc-ni:network-instance-ref
              +--ro gnsi-certz:certificate-version?                             version
              +--ro gnsi-certz:certificate-created-on?                          created-on
              +--ro gnsi-certz:ca-trust-bundle-version?                         version
              +--ro gnsi-certz:ca-trust-bundle-created-on?                      created-on
              +--ro gnsi-certz:certificate-revocation-list-bundle-version?      version
              +--ro gnsi-certz:certificate-revocation-list-bundle-created-on?   created-on
              +--ro gnsi-certz:authentication-policy-version?                   version
              +--ro gnsi-certz:authentication-policy-created-on?                created-on
              +--ro gnsi-certz:ssl-profile-id?                                  string
              +--ro gnsi-certz:counters
                 +--ro gnsi-certz:connection-rejects?       oc-yang:counter64
                 +--ro gnsi-certz:last-connection-reject?   oc-types:timeticks64
                 +--ro gnsi-certz:connection-accepts?       oc-yang:counter64
                 +--ro gnsi-certz:last-connection-accept?   oc-types:timeticks64

```
