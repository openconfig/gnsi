# gNSI cert Service Protobuf Definition
**Contributors**: hines@google.com, morrowc@google.com, tmadejski@google.com
**Last Updated**: 2022-07-12

## Background

The cert service definition provides the API to be used for installing,
rotating, deleting, and testing PKI primatives used on network systems.
The Rotate and Install are bidirectional streaming rpcs which permit
mutating Certificates, Root Certificate Bundles, Certificate Revocation
Lists.  For either an Install or Rotate stream it is possible to mutate
one or more of the elements, and to send a Finalize message once the
in-flight change has been verified to be operational. Failure to send
the Finalize message will result in the candidate element being discarded
and the original element being used instead.

## Motivation

Management of the PKI elements for a network system should have
a clear and direct method for installation and update.

### CertificateManagement.Install

CertificateManagemetn.Install will permit installation, and
verification of function, of any of the PKI elements. The normal
use-case would be to:

* send an Certificate and Key to a network system as an
InstallCertificateRequest.
* verify that the services which will use the certificate continues
to operate normally.
* send a FinalizeRequest to finish the installation process.

### CertificateManagement.Rotate

CertificateManagement.Rotate will permit rotation, and
verification of function, of any of the PKI elements.
The normal use-case would be to:

* send an CertificateBundle to a network system as a
RotateCertificateRequest.
* verify that the services which will use the new certificate bundle
continue to operate normally.
* send a FinalizeRequest to finish the rotateion process.

## User Experiences

### A Certificate is to be installed

Create a new Certificate and Key.

Send that certificate to the target network system with a
cert.InstallCertificateRequest to the cert.Install rpc. The
InstallCertificateRequest's install_request will be a
cert.Certificate.

Verify that the certificate newly deployed is usable by the relevant
services, that the services properly present the certificate upon
new service connections.

Send a cert.Finalize message to the cert.Install rpc to close
out the action.

If the stream is disconnected prior to the Finalize message being
sent, the proposed configuration is rolled back automatically.

### A CertificateRevocationList is to be installed

Create a new CertificateRevocationList (CRL).

Send that certificate to the target network system with a
cert.InstallCertificateRequest to the cert.Install rpc. The
InstallCertificateRequest's install_request will be a
cert.CertificateRevocationListBundle.

Optional, verify that the CRL newly deployed is usable by the relevant
services.

Send a cert.Finalize message to the cert.Install rpc to close
out the action.

If the stream is disconnected prior to the Finalize message being
sent, the proposed configuration is rolled back automatically.

### A CertificateBundle is to be rotated or updated

Create, and test, a new CertificateBundle.

Send that policy to the target network system with a
cert.RotateCertificateRequest to cert.Rotate rpc. The
RotateCertificateRequest's rotate_request will be a cert.CertificateBundle.

Verify that the CertificateBundle newly rotated is used by services
which require it.

Send a Finalize message to the cert.Rotate rpc to close out the action.

If the stream is disconnected prior to the Finalize message being
sent, the proposed configuration is rolled back automatically.

### A Certificate is rotated, the session breaks before Finalize

Create a new Certificate and Key.

Send that certificate to the target network system with a
cert.InstallCertificateRequest to the cert.Install rpc. The
InstallCertificateRequest's install_request will be a
cert.Certificate.

Verify that the certificate newly deployed is usable by the relevant
services, that the services properly present the certificate upon
new service connections.

The connection to the network system is broken, there is no Finalize sent.

The gNSI service rolls back the candidate and re-installs the original
certificate and key.


## Open Questions/Considerations

None to date.
