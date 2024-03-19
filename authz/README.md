# gNSI.authz
## The idea

Implementation of reliable and fast APIs to control remote network-connected
device like a switch or a router is not easy due to complexity of the
communication over a computer network. Fortunately, by using `Remote Procedure
Call` (`RPC`) technique all (or most) of this complexity can be hidden from
a user, but because those APIs can be used to create havoc in mission-critical
networks, not everybody should be able to perform all RPCs provided by those
management APIs.

`gNSI.authz` defines an API that allows for configuration of the RPC service on
a switch to control which user can and cannot access specific RPCs.

## The gRPC-level Authorization Policy

The policy to be enforced is defined in the form of a JSON string whose
structure depends on the requirements of the RPC server.

In the case of a `gRPC`-based server the JSON string's schema can be found
[here](https://github.com/grpc/proposal/blob/master/A43-grpc-authorization-api.md).
It also can be described using the following PROTOBUF definition.

```protobuf
// Copyright 2021 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

syntax = "proto3";

package grpc.auth.v1;

option go_api_flag = "OPEN_TO_OPAQUE_HYBRID";  // See http://go/go-api-flag.

// Peer specifies attributes of a peer. Fields in the Peer are ANDed together,
// once we support multiple fields in the future.
message Peer {
  // Optional. A list of peer identities to match for authorization. The
  // principals are one of, i.e., it matches if one of the principals matches.
  // The field supports Exact, Prefix, Suffix, and Presence matches.
  // - Exact match: "abc" will match on value "abc".
  // - Prefix match: "abc*" will match on value "abc" and "abcd".
  // - Suffix match: "*abc" will match on value "abc" and "xabc".
  // - Presence match: "*" will match when the value is not empty.
  repeated string principals = 1;
}

// Specification of HTTP header match attributes.
message Header {
  // Required. The name of the HTTP header to match. The following headers are
  // *not* supported: "hop-by-hop" headers (e.g., those listed in "Connection"
  // header), HTTP/2 pseudo headers (":"-prefixed), the "Host" header, and
  // headers prefixed with "grpc-".
  string key = 1;

  // Required. A list of header values to match. The header values are ORed
  // together, i.e., it matches if one of the values matches. This field
  // supports Exact, Prefix, Suffix, and Presence match. Multi-valued headers
  // are considered a single value with commas added between values.
  // - Exact match: "abc" will match on value "abc".
  // - Prefix match: "abc*" will match on value "abc" and "abcd".
  // - Suffix match: "*abc" will match on value "abc" and "xabc".
  // - Presence match: "*" will match when the value is not empty.
  repeated string values = 2;
}

// Request specifies attributes of a request. Fields in the Request are ANDed
// together.
message Request {
  // Optional. A list of paths to match for authorization. This is the fully
  // qualified name in the form of "/package.service/method". The paths are ORed
  // together, i.e., it matches if one of the paths matches. This field supports
  // Exact, Prefix, Suffix, and Presence matches.
  // - Exact match: "abc" will match on value "abc".
  // - Prefix match: "abc*" will match on value "abc" and "abcd".
  // - Suffix match: "*abc" will match on value "abc" and "xabc".
  // - Presence match: "*" will match when the value is not empty.
  repeated string paths = 1;

  // Optional. A list of HTTP header key/value pairs to match against, for
  // potentially advanced use cases. The headers are ANDed together, i.e., it
  // matches only if *all* the headers match.
  repeated Header headers = 3;
}

// Specification of rules.
message Rule {
  // Required. The name of an authorization rule.
  // It is mainly for monitoring and error message generation.
  // This name must be unique within the list of deny (or allow) rules.
  string name = 1;

  // Optional. If not set, no checks will be performed against the source. An
  // empty rule is always matched (i.e., both source and request are empty).
  Peer source = 2;

  // Optional. If not set, no checks will be performed against the request. An
  // empty rule is always matched (i.e., both source and request are empty).
  Request request = 3;
}

// AuthorizationPolicy defines which principals are permitted to access which
// resource. Resources are RPC methods scoped by services.
//
// In the following yaml policy example, a peer identity from ["admin1",
// "admin2", "admin3"] is authorized to access any RPC methods in pkg.service,
// and peer identity "dev" is authorized to access the "foo" and "bar" RPC
// methods.
//
// name: example-policy
// allow_rules:
// - name: admin-access
//   source:
//     principals:
//     - "spiffe://foo.com/sa/admin1"
//     - "spiffe://foo.com/sa/admin2"
//     - "spiffe://foo.com/sa/admin3"
//   request:
//     paths: ["/pkg.service/*"]
// - name: dev-access
//   source:
//     principals: ["spiffe://foo.com/sa/dev"]
//   request:
//     paths: ["/pkg.service/foo", "/pkg.service/bar"]

message AuthorizationPolicy {
  // Required. The name of an authorization policy.
  // It is mainly for monitoring and error message generation.
  string name = 1;

  // Optional. List of deny rules to match. If a request matches any of the deny
  // rules, then it will be denied. If none of the deny rules matches or there
  // are no deny rules, the allow rules will be evaluated.
  repeated Rule deny_rules = 2;

  // Required. List of allow rules to match. The allow rules will only be
  // evaluated after the deny rules. If a request matches any of the allow
  // rules, then it will be allowed. If none of the allow rules match, it
  // will be denied.
  repeated Rule allow_rules = 3;
}
```

## An example

Below is an example of a gRPC-level Authorization Policy that allows two admins,
Alice and Bob, access to all RPCs that are defined by the `gNSI.ssh` interface.
Nobody else will be able to call any of the `gNSI.ssh` RPCs.

```json
{
  "name": "gNSI.ssh policy",
  "allow_rules": [{
    "name": "admin-access",
    "source": {
      "principals": [
        "spiffe://company.com/sa/alice",
        "spiffe://company.com/sa/bob"
      ]
    },
    "request": {
      "paths": [
        "/gnsi.ssh.Ssh/*"
      ]
    }
  }]
}
```

## Managing the gRPC-based Authorization Policy

### Initial (factory reset) state assumption

When a device boots for the first time it should have:

1. The `gNSI.authz` service transitions to up and running.
1. The default gRPC-level Authorization Policy for all active gRPC services.

   The default gRPC-level Authorization Policy must allow access to all RPCs.
1. Once a gNSI policy is set (uploaded and Finalized), the default policy
   disposition becomes deny, as mentioned in the AuthorizationPolicy message
   documentation above.

### Updating the policy

Every policy needs changes from time to time and the `gNSI.authz.Rotate()` RPC
is designed to do this task.

There are 5 steps in the process of updating (rotating) an gRPC-level
Authorization Policy, namely:

1. Starting the `gNSI.authz.Rotate()` streaming RPC.

   As the result a streaming connection is created between the server (the
   switch) and the client (the management application) that is used in the
   following steps.

   > **⚠ Warning**
   > Only one `gNSI.authz.Rotate()` can be in progress.

1. The client uploads new gRPC-level Authorization Policy using
   the `UploadRequest` message.

   For example:
   ```json
   {
     "version": "version-1",
     "created_on": "1632779276520673693",
     "policy": {
       "name": "gNSI.ssh policy",
       "allow_rules": [{
         "name": "admin-access",
         "source": {
           "principals": [
             "spiffe://company.com/sa/alice",
             "spiffe://company.com/sa/bob"
           ]
         },
         "request": {
           "paths": [
             "/gnsi.ssh.Ssh/*"
           ]
         }
       }],
       "deny_rules": [{
         "name": "sales-access",
         "source": {
           "principals": [
             "spiffe://company.com/sa/marge",
             "spiffe://company.com/sa/don"
           ]
         },
         "request": {
           "paths": [
             "/gnsi.ssh.Ssh/MutateAccountCredentials",
             "/gnsi.ssh.Ssh/MutateHostCredentials"
           ]
         }
       }]
     }
   }
   ```

   > **⚠ Warning**
   > There is only one gRPC-level Authorization Policy on the device therefore
   > it is "declarative" for all gRPC servers and services on the device.
   > In other words: all policies must be defined in the policy being rotated as
   > this rotate operation will replace all previously defined/used policies
   > once the `Finalize` message is sent.

   The information passed in both the `version` and the `created_on` fields is
   not used internally by the `gNSI.authz` service and is designed to help keep
   track of what gRPC-level Authorization Policy is active on a particular
   switch.

1. After syntactic validation and activating the new policy, the server sends
   the `UploadResponse` back to the client

1. The client verifies the correctness of the new gRPC-level Authorization
   Policy using separate `gNSI.authz.Probe()` RPC(s)

1. The client sends the `Finalize` message indicating the previous gRPC-level
   Authorization Policy can be deleted.

   > **⚠ Warning**
   > Closing the stream without sending the `Finalize` message will result in
   > abandoning the uploaded policy and rollback to the one that was active
   > before the Rotation RPC started.

### Evaluating the rules

In a simple deployment, the set of rules in the gRPC-level Authorization Policy
most likely will be clear enough for a human to analyze, but in a data-center
environment the list of rules will likely be long and complex, and therefore
difficult to reason about.

To help this process the `gNSI.authz` API includes the `gNSI.authz.Probe()` RPC.

This RPC allows for checking the response of the gRPC-level Authorization Policy
engine to a RPC performed by a specific user based the installed policy.

Because the policy uploaded during the `gNSI.authz.Rotate()` call becomes active
immediately, the `gNSI.authz.Probe()` can be used to check if the uploaded
policy provides the expected response, without attempting the (potentially
destructive) RPC in question, while the `gNSI.authz.Rotate()` is still active
(the stream is still open and the `Finalize` message has not been sent yet).

For example, to check if `alice` can perform the
`gNSI.ssh.MutateAccountCredentials()` RPC the `gNSI.authz.Probe()` should be
called with the following parameters:

```json
{
  "user": "spiffe://company.com/sa/alice",
  "rpc": "gNSI.ssh.MutateAccountCredentials"
}
```

As `alice` is listed in the example policy in the `allow_rules` section the
expected result of the `gNSI.authz.Probe()` RPC is:

```json
{
  "action": "ACTION_PERMIT",
  "version": "<a version string provided during in the UploadRequest>"
}
```

## OpenConfig data models for gNSI Authorization Policy

Yang data models for authz are defined in the [OpenConfig public repository(https://github.com/openconfig/public/tree/master/release/models/gnsi)].  Documentation for OpenConfig including searchable list of paths and tree representations are at [OpenConfig.net](https://openconfig.net/projects/models/)
