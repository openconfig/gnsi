# gNSI.authz
## The idea

Implementation of reliable and fast APIs to control remote network-connected
device like a switch or a router is not easy due to complexity of the
communication over a computer network. Fortunately by using `Remote Procedure
Call` (`RPC`) technique all (or most) of this complexity can be hidden from
a user but because those APIs can be used to create havoc in mission-critical
networks not everybody should be able to perform all RPC provided by those
management APIs.

`gNSI.authz` defines an API that allows for configuration of the RPC service on
a switch to control which user can and cannot access specific RPCs.

## The gRPC-level Authorization Policy

The policy the RPC is to enforce is defined in a form of a JSON string whose
structure depends on the requirements of the RPC server.

In the case of a `gRPC`-based server the JSON string's schema can be found
[here](https://github.com/grpc/proposal/pull/246).
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
  // rules, then it will allowed. If none of the allow rules matches, it will be
  // denied.
  repeated Rule allow_rules = 3;
}
```

## An example

Below is an example of a gRPC-level Authorization Policy that allows two admins:
Alice and Bob access to all RPCs that are defined by the `gNSI.ssh` interface.
Nobody else will be able to call any of the `gNSI.ssh` RPCs.

```json
{
  "name": "gNSI.ssh policy",
  "allow_rules": [{
    "name": "admin-access",
    "principals": [
      "spiffe://company.com/sa/alice",
      "spiffe://company.com/sa/bob"
      ],
    "request": {
      "paths": [
        "/gnsi.ssh.Ssh/*"
      ]
    }
  }]
}
```

## Managing the gRPC-based Authorization Policy

### Initial (fresh out of the box) state assumption

When a device boots for the first time it should have:

1. The `gNSI.authz` service up and running
1. A default gRPC-level Authorization Policy for every active gRPC service.

   The initial default gRPC-level Authorization Policy can either allow access
   to all RPCs or deny access to all RPCs except for the `gNSI.authz` family of
   RPCs.

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
     "policy": "
        {
          "name": "gNSI.ssh policy",
          "allow_rules": [{
            "name": "admin-access",
            "principals": [
              "spiffe://company.com/sa/alice",
              "spiffe://company.com/sa/bob"
              ],
            "request": {
              "paths": [
                "/gnsi.ssh.Ssh/*"
              ]
            }
          }]
        }"
   }
   ```

   The information passed in both the `version` and the `created_on` fields is
   not used internally by the `gNSI.authz` service and is designed to help keep
   track of what gRPC-level Authorization Policy is active on a particular
   switch.

1. After pre-validating and activating the new policy, the server sends the
   `UploadResponse` is sent back to the client

1. The client verifies the correctness of the new gRPC-level Authorization
   Policy using separate `gNSI.authz.Probe()` RPC(s)

1. The client sends the `Finalize` message indicating the previous gRPC-level
   Authorization Policy can be deleted.

   > **⚠ Warning**
   > Closing the stream without sending the `Finalize` message will result in
   > abandoning the uploaded policy and rollback of the one that was active
   > before the RPC started.

### Evaluating the rules

In a simple deployment the set of rules in the gRPC-level Authorization Policy
most likely will be clear enough for a human to analyze but in a data-center
environment most likely the list of rules will be long and complex and therefore
hard to reason about.

To help this process the `gNSI.authz` API includes the `gNSI.authz.Probe()` RPC.

This RPC allows for checking the response of the gRPC-level Authorization Policy
engine to a RPC performed by a specific user based the installed policy.

Because the policy uploaded during the `gNSI.authz.Rotate()` call becomes active
immediately, the `gNSI.authz.Probe()` can be used to check if the uploaded
policy provides the expected response without attempting performing the
(potentially destructive) RPC in question while the `gNSI.authz.Rotate()` is
still active (the stream is opened and the `Finalize` message has not been sent
yet.

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

## OpenConfig Extension for the gMNI gRPC-based Authorization Policy telemetry

### `gnsi-authz.yang`

An overview of the changes defined in the `gnmi-authz.yang` file are shown
below.

```txt
module: gnsi-authz

  augment /oc-sys:system/oc-sys:aaa/oc-sys:authorization/oc-sys:state:
    +--ro grpc-authz-policy-version?      version
    +--ro grpc-authz-policy-created-on?   created-on
```

### `openconfig-system` tree
The  `openconfig-system` subtree after augments defined in the `gnsi-authz.yang`
file is shown below.

<details>
<summary>
The diagram of the tree.
</summary>

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
     |  |  |  +--ro authorization-method*                      union
     |  |  |  +--ro gnsi-authz:grpc-authz-policy-version?      version
     |  |  |  +--ro gnsi-authz:grpc-authz-policy-created-on?   created-on
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
              +--ro oc-sys-grpc:name?                      string
              +--ro oc-sys-grpc:services*                  identityref
              +--ro oc-sys-grpc:enable?                    boolean
              +--ro oc-sys-grpc:port?                      oc-inet:port-number
              +--ro oc-sys-grpc:transport-security?        boolean
              +--ro oc-sys-grpc:certificate-id?            string
              +--ro oc-sys-grpc:metadata-authentication?   boolean
              +--ro oc-sys-grpc:listen-addresses*          union
              +--ro oc-sys-grpc:network-instance?          oc-ni:network-instance-ref

```
</details>

For interactive version click [here](gnsi-authz.html).
