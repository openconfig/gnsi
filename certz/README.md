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
as a whole to a gRPC service.

There is always at least one profile present on a target - the `system_default_profile`
which is vendor provided.
This profile cannot be changed or deleted.
See the the [System default SSL profile](#system-default-ssl-profile) section below.

Profiles existing on a target can be discovered using the
`Certz.GetProfileList()` RPC.

A SSL profile can be added using the `Certz.AddProfile()` RPC.

When a profile is no longer needed it can be removed from the target via
`Certz.DeleteProfile()` RPC. Note that the system_default_profile SSL
profile cannot be removed.

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
A new profile can choose to use existing artifacts from other profiles, via sending `Entity` messages with `ExistingEntity` set with the `ssl_profile_id` set to the source
profile to copy from, and the `entity_type` field set to the type of entity to be copied.

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

#### MBM-Boot workflow

See the [TCG Reference Integrity Manifest (RIM) Information
Model](https://trustedcomputinggroup.org/resource/tcg-reference-integrity-manifest-rim-information-model/)
for more details on the following workflow.

Call `Certz.GetIntegrityManifest`. The `Certz.GetIntegrityManifestResponse`'s
`manifest` field will contain the reference integrity manifest. Determine the
PCRs to be included and all allowable digest values.

Send a `Certz.GenerateCSRRequest` to the `Certz.Rotate` endpoint, containing a
`Certz.ReferenceIntegritySpec`. Using the returned `Certz.GenerateCSRResponse`
and the `MBMData` within, do the following:

* Verify the `ek_leaf_cert` using the `ek_cert_chain` and your trust anchor.
* Optional: Verify that the AK matches your expectations, using the
  `ak_creation_data` struct.
* Validate the `ak_signature` over the `ak_attestation` struct which was
  certified by the EK, and validate its contents. This verifies the AK.
* Validate the `signature` over `quoted` by the AK. Then validate that the PCRs
  match one of the allowed ones.
* Validate the `csr_signature` over the `certificate_signing_request` by the AK.
  This verifies the CSR.

Get a new certificate issued by a trusted CA using the CSR. Then `Certz.Rotate`
as normal.

### Open Questions/Considerations

None to date.

## OpenConfig Data models for gNSI certz

Yang data models for certz are defined in the [OpenConfig public repository(https://github.com/openconfig/public/tree/master/release/models/gnsi)].  Documentation for OpenConfig including searchable list of paths and tree representations are at [OpenConfig.net](https://openconfig.net/projects/models/)
