# gNSI Authorization Protobuf Definition

**Contributors**: <hines@google.com>, <morrowc@google.com>, <tmadejski@google.com>
**Last Updated**: 2022-07-12

## Background

The authorization proto definition provides a clear method to define
and implement policy authorizing user or group access to the management
systems and interfaces of a network deployment. The ability to permit
or deny access to these systems and interfaces is intended to help
operate a network in a more safe and secure manner.

## Motivation

Authorization services for network systems can, at times, be operated
as remote services. One example of this is a TACACS+ service, there are
inherent challenges with operating a network and relying upon a remote
service for critical functions performed on network systems. A solution
to provide on-device Authorization with the fidelity of remote services
should be provided.

## Architecture

There is no requirement for the Authorization policy to be evaluated:
on the local network system, in a microservice operated on the network
system, or in every gNMI service individually on the network system.
It is expected that gNMI services enabled on a network system respect
the AuthorizationPolicy installed, however.

The pathz.Rotate rpc is a bi-directional streaming RPC, it's
possible to send more than one policy change through an
open stream checkpointing the policy with pathz.Finalize messages or
replacing the candidate AuthorizationPolicy to verify functionality
changes prior to the Finalize message. If the stream is closed prior
to the Finalize message being received at the server, the candidate
AuthorizationPolicy is discarded and the existing policy again becomes
active.

### Pathz.Rotate

Pathz.Rotate will permit rotation, and verification of function,
of an AuthorizationPolicy. The normal use-case would be to:

* send an AuthorizationPolicy to a network system as a
RotatePathzRequest
* verify access/authorization has changed to the desired state
through existing gNMI methods, or with pathz.Probe requests.
* send a FinalizeRequest to finish the installation process.

### Pathz.Probe

Pathz.Probe provides a method to test the AuthorizationPolicy
with a path.ProbeRequest which includes a user and gNMI path. This
enables network operations and management systems to verify that
the access expected is either permitted or denied in accordance
with the expected to be deployed AuthroizationPolicy.

## User Experiences

### gNMI interaction with Subscribe, Set or Get

When a client makes a request to gNMI with a system with a pathz policy
installed, the pathz policy is evaluated against the paths requested. The
if there is no explicit permit to the path requested the client the RPC must
be rejected. This keeps the gNMI server from being potentially DOS'ed by
clients requesting top level paths which then have to be recursed for all
possible accepts which might be a lower levels of the tree.

### An AuthorizationPolicy is to be installed

Create, and test, a new AuthorizationPolicy.

Send that policy to the target network system with a
pathz.InstallAuthzRequest to the pathz.Install rpc. The
InstallAuthzRequest's install_request will be a pathz.UploadRequest.

Verify that the policy newly deployed performs according to the documented
intent, by sending pathz.Probe requests to the network system.

Send a pathz.Finalize message to the pathz.Install() rpc to close
out the action.

If the stream is disconnected prior to the Finalize message being
sent, the proposed configuration is rolled back automatically.

### An AuthorizationPolicy is to be rotated or updated

Create, and test, a new AuthorizationPolicy.

Send that policy to the target network system with a
pathz.RotateAuthzRequest to pathz.Rotate rpc. The
RotateAuthzRequest's rotate_request will be a pathz.UploadRequest.

Verify that the policy newly rotated performs according to the documented
intent, by sending pathz.Probe requests to the network system.

Send a Finalize message to the pathz.Rotate rpc to close
out the action.

If the stream is disconnected prior to the Finalize message being
sent, the proposed configuration is rolled back automatically.

### AuthorizationPolicy rotation interaction with gNMI Subscribe and gSII Modify RPCs

When the pathz policy is rotated, ongoing gNMI Subscribe RPCs will be disconnected.
Any ongoing gSII Modify RPCs will also be disconnected.

This forces gNMI/gSII clients to reconnect and ensures that all gNMI subscriptions
or gSII Modify streams are proceeding using the most recent pathz policy.

## Open Questions/Considerations

None to date.
