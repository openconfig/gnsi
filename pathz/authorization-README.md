# gNMI Authorization Protocol

## Objective

This proto definition and the reference code(to be delivered separately) serve
to describe an authorization framework for controlling which gNMI paths of a
network device users can access. The authorization policy is initially intended
to be deployed to a device, with the ability to define:

* Policy rules - each rule defines a single authorization policy.
* Groups of users - as a method to logically group users in the administrative
  domain, for instance: operators or administrators.
* Users - individuals referenced in rules or group definitions.

Authentication information is not included in this Authorization configuration.

Policy rules are matched based on the best match for the authorization request,
not the first match against a policy rule. Best match enables a configuration
which permits a user or group access to particular gNMI paths while denying
subordinate portions of the permitted paths, or the converse, without regard
to ordering of the rules in the configuration.

## Best Match

Authorization is performed for a singular user, gNMI path access, and access
methodology (READ/WRITE). The result of an Authorization evaluation is an
Action (PERMIT/DENY), policy version, and rule identifier.

Among all matching policies, the best, or most specific match,
is determined from the following rules in order:

1. A longer matching path is preferred over a shorter one.
1. Definite keys over wildcards keys. A rule with more definite keys is
   preferred over one with fewer.
1. User over group. A rule that matches with the user is preferred over one
   with matches with a group a user belongs to.
1. Deny over permit. If all above are equal, prefer the rule with DENY action.

Match rules permit a match against:

* User or Group (not both)
* an gNMI path
* an access mode (READ / WRITE)

An implicit deny is assumed, if there is no matching rule in the policy.

As a request is evaluated against the configured policy, a READ (gNMI `Get` or
`Subscribe`) request for the configuration tree may traverse all of the tree
and subtrees. The client request must have an explicit permit for the path or
a parent path of the request for the request to be permitted. For portions of
the tree for which the user has no access no data will be returned. A WRITE
request which attempts to write to a denied gNMI path or element will return
a "Permission Denied" error to the caller.

[gNMI paths](https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-specification.md#222-paths)
are hierarchical, and rooted at a defined "origin". gNMI may contain paths
such as:

```proto
    /a/b/c/d
```

Paths may also have attributes associated with the path elements such as:

```proto
    /a/b[key=foo]/c/d
```

Attributes may be wildcarded in the policy, such as:

```proto
    /a/b[key=*]/c/d
```

Wildcards are only valid in policy rules when applied to attributes.
Permitted use:

```proto
    /a/b[key=*]/c/d
```

Not permitted use:

```proto
    /a/b[key=foo]/*/d
```

An example of two rules with similar paths that differ only with respect
to attribute wildcarding:

```proto
   /a/b[key=FOO]/c/d
   /a/b[key=*]/c/d
```

The first rule specifies a single key, that rule is more specific than
the second rule, which specifies 'any value is accepted'.

### Conflict Resolution

The policy must be evaluated in the order described by the "Best Match" section.
In case of conflicting paths or group membership, preferring DENY over ALLOW ensures
only a single action is applicable.

### Examples

Probe Path: `/network-instances/network-instance[name=DEFAULT]/protocols/protocol[identifier=BGP]`  
Group Memberships: `admin: [stevie]; engineers: [stevie]`

Example 1

Installed Rules

```sh
/network-instances/network-instance[name=DEFAULT]/protocols/protocol[identifier=BGP] -> group admin, action PERMIT
/network-instances/network-instance[name=*]/protocols/protocol[identifier=BGP] -> user stevie, action DENY
```

Result: PERMIT, both paths are the same length, the first policy has more definite keys,
so it is a better match, even though it applies to a group while the second applies to the user.

Example 2

Installed Rules

```sh
/network-instances/network-instance[name=DEFAULT]/protocols/protocol[identifier=BGP] -> user stevie, action PERMIT
/network-instances/network-instance[name=DEFAULT]/protocols/protocol[identifier=BGP] -> group admin, action DENY
```

Result: PERMIT, both paths are the same length, the first policy applies to the user,
so it is prefer over the rule that applies to group.

Example 3

Installed Rules

```sh
/network-instances/network-instance[name=DEFAULT]/protocols/protocol[identifier=*] -> user stevie, action PERMIT
/network-instances/network-instance[name=*]/protocols/protocol[identifier=BGP] -> user stevie, action DENY
```

Result: DENY, both paths are the same length, have the same amount of definite keys,
apply to the user (not group), so prefer DENY over PERMIT.

Example 4

Installed Rules

```sh
/network-instances/network-instance[name=DEFAULT]/protocols/protocol[identifier=BGP] -> group admin, action PERMIT
/network-instances/network-instance[name=DEFAULT]/protocols/protocol[identifier=BGP] -> group engineers, action DENY
```

Result: DENY, both paths are the same length, have the same amount of definite keys,
apply to groups that user is a member of, so prefer DENY over PERMIT.

Example 5

Installed Rules

```sh
/interfaces/interface[name=*] -> group core-controllers, action PERMIT
/interfaces/interface[name=*] -> group core-eng, action PERMIT
/interfaces/interface[name=et-1/0/1]/state/counters -> user customer-controller1, action PERMIT
/interfaces/interface[name=et-1/0/2]/state/counters -> user customer-controller2, action PERMIT
/interfaces/interface[name=et-1/0/1] -> core-controller, action DENY
/interfaces/interface[name=et-1/0/2] -> core-controller, action DENY

```

Result:

This will assume that core-controller1 is member of core-controllers
This will assume that eng1 is member of core-eng

* `gnmi.Subscribe(10.0.0.10:515253, eng1, /interfaces/interface/state/counters, ONCE)`

Subscribe will be accepted and all subtrees will be returned as eng1 is member core-eng group

* `gnmi.Subscribe(10.0.0.10:515253, customer-controller1, /interfaces/interface/state/counters, ONCE)`

Subscribe will be rejected as user does not have access at that container

* `gnmi.Subscribe(10.0.0.10:515253, customer-controller1, /interfaces/interface[name=et-1/0/1]/state/counters, ONCE)`

Subscribe will be accepted and all subtrees will be returned.

* `gnmi.Subscribe(10.0.0.10:515253, core-contollers, /interfaces/interface/state/counters, ONCE)`

Subscribe will be accepted, only interfaces not matching the deny rule will be returned.

* `gnmi.Subscribe(10.0.0.10:515253, core-controllers, /interfaces/interface[name=et-1/0/1]/state/counters, ONCE)`

Subscribe will be rejected due to explicit DENY rule for path.

## Bootstrap / Install Options

System bootstrap, or install, operations may include an authorization policy
delivered during bootstrap operations. It is suggested that the bootstrap
process include the complete authorization policy so all production tools
and services have immediate authorized access to finish installation and
move devices into production in a timely manner.

Using [Bootz](https://github.com/openconfig/bootz) or the Secure Zero Touch
Provisioning (sZTP - RFC8572)process for bootstrap/installation is a recommended
method for accomplishing this delivery, and the delivery of all other bootstrap
artifacts in a secure manner.

## An Example Authorization Protobuf

```proto
version: "UUID-1234-123123-123123"
created_on: 1234567890
# Define 2 groups.
group {
  name: "family-group"
  members { name: "stevie" }
  members { name: "brian" }
}
group {
  name: "test-group"
  members { name: "crusty" }
  members { name: "the-clown" }
}
# Action stevie to access /this/is/a/message_path in a READ manner.
policy {
  id: "one"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: READ
  user { name: "stevie" }
}
# Action members of family-group to access /this/is/a/different/message_path in
# READ mode.
policy {
  id: "two-read"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem { name: "different" }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: READ
  group { name: "family-group" }
}
# Action members of family-group to access /this/is/a/different/message_path in
# WRIITE mode.
policy {
  id: "two-write"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem { name: "different" }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: WRITE
  group { name: "family-group" }
# Demonstrate READ access to a key with an attribute defined.
policy {
  id: "key"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem {
      name: "keyed"
      key{
        key: "Ethernet"
        value: "1/2/3"
      }
    }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: READ
  group { name: "test-group" }
}
# Demonstrate READ access to a key with a wildcard attribute.
policy {
  id: "wyld"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem {
      name: "keyed"
      key{
        key: "Ethernet"
        value: "*"
      }
    }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: READ
  group { name: "family-group" }
}
# Demonstrate a key with a wildcard attribute and a user specific match.
# The previous policy matches all family-group users and permits a command
# path, the policy rule below specifically denies brian access to these paths.
policy {
  id: "wyld-stallions"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem {
      name: "keyed"
      key{
        key: "Ethernet"
        value: "*"
      }
    }
    elem { name: "message_path" }
  }
  action: DENY
  mode: READ
  user { name: "brian" }
}
# Add a final rule which is an implicit deny rule.
```

The example first policy rule:

```proto
# Action stevie to access /this/is/a/message_path for READ.
policy {
  id: "one"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: READ
  user { name: "stevie" }
}
```

permits the singular user "stevie" to access the path:

```shell
    /this/is/a/message_path
```

Additionally, "stevie" is permitted access to all paths below the defined path,
in a READ only mode, such as:

```shell
    /this/is/a/message_path/the
    /this/is/a/message_path/the/one
    /this/is/a/message_path/the/one/that
    /this/is/a/message_path/the/one/that/knocks
```

The second policy rule:

```proto
# Action members of family-group to run /this/is/a/different/message_path
policy {
  id: "two-read"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem { name: "different" }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: READ
  group { name: "family-group" }
}
```

example policy permits members or the family-group access to a single path, for
reading:

```shell
    /this/is/a/different/message_path
```

and all path elements beyond "message_path":

```shell
    /this/is/a/different/message_path/foo
    /this/is/a/different/message_path/bar
    /this/is/a/different/message_path/foo/baz/bing/boop
```

The third policy rule:

```proto
# Action members of family-group to run /this/is/a/different/message_path
policy {
  id: "two-write"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem { name: "different" }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: WRITE
  group { name: "family-group" }
}
```

example policy permits members or the family-group access to a single path, for
writing:

```shell
    /this/is/a/different/message_path
```

and all path elements beyond "message_path":

```shell
    /this/is/a/different/message_path/foo
    /this/is/a/different/message_path/bar
    /this/is/a/different/message_path/foo/baz/bing/boop
```

The fourth policy rule:

```proto
# Demonstrate a key with an attribute defined.
policy {
  id: "key"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem {
      name: "keyed"
      key{
        key: "name"
        value: "Ethernet1/2/3"
      }
    }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: READ
  group { name: "test-group" }
}
```

Permits access by the "test-group" users to the keyed path, in a
read only manner:

```shell
    /this/is/a/keyed[name=Ethernet1/2/3]/message_path
```

and all path elements beyond "message_path". The final policy rule:

```proto
# Demonstrate a key with a wildcard attribute.
policy {
  id: "wyld"
  path {
    origin: "foo"
    elem { name: "this" }
    elem { name: "is" }
    elem { name: "a" }
    elem {
      name: "keyed"
      key{
        key: "name"
        value: "*"
      }
    }
    elem { name: "message_path" }
  }
  action: PERMIT
  mode: READ
  group { name: "family-group" }
}
```

permits access by the "family-group" users to the keyed path, with no
restrictions on the key values, but still as read-only:

```shell
    /this/is/a/keyed[name=Ethernet1/2/3]/message_path
    /this/is/a/keyed[name=POS3]/message_path
    /this/is/a/keyed[name=Serial4/1]/message_path
    /this/is/a/keyed[name=HSSI2]/message_path
```

Additionally, the path elements beyond "message_path" are available for access
to this group as well.

The wildcard character "*" (asterisk) may only be used as a value in keyed
elements, if the keys are missing in a keyed path a wildcard is assumed. The
wildcard is only used to mask out all possible values but not portions of
values, for instance:

```shell
    /this/is/a/keyed[name=*]/things - permitted usage of wildcard
    /this/is/a/keyed[name=Ethernet1/*/3]/things - NOT permitted usage of wildcard
```

The end of every policy includes an implicit deny policy rule.  This rule will
cause all matches to be counted.
