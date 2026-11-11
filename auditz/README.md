# gNSI Runtime Host & Kernel Security Auditing (`gNSI.auditz`) Protocol

## Objective

This protocol buffer definition specifies a vendor-neutral gRPC interface for
streaming continuous operating system, kernel, process, filesystem, privilege,
and control-plane socket security audit telemetry (`gNSI.auditz`) from a network
device (`System`) to a remote collection service (`Collector`).

Whereas `Bootz`, `Enrollz`, and `Attestz` establish hardware root-of-trust and
verify cryptographic measured state at boot and periodic quote intervals, and
`gNSI.acctz` audits management-plane CLI and gRPC commands, `gNSI.auditz` closes
the continuous runtime visibility gap across the underlying OS and kernel
execution layer during device uptime.

## Method of Operation

Runtime security audit records are available at the `gNSI` origin:

```text
gnsi.auditz
```

Events are streamed from a System at the request of a remote Collector via the
unidirectional `EventSubscribe()` RPC. Point-in-time operational status,
resource utilization, and cumulative counters are exported via `gNMI` in
`openconfig-gnsi-auditz.yang`.

Any subsystem or FRU running an independent OS kernel (such as linecards, fabric
cards, or a standby supervisor/RP) must proxy its runtime security events
through the active RP's `gNSI` endpoint and populate `EventResponse.component`
(`gnmi.Path`, e.g., `/components/component[name=linecard0]`).

### Event Categories (Implementation-Agnostic Semantic Contract)

Each `EventResponse` message carries one of five core runtime security event
payloads or a periodic heartbeat. Event categories are specified in terms of
their **semantic security behaviors** rather than prescribing a specific Linux
syscall implementation: if a vendor NOS intercepts or implements operations in
user-space (for example, via `LD_PRELOAD` socket shims or user-space packet
stacks on revenue ports), it must still emit the corresponding event message:

1. **`ProcessEvent`**: Semantic process creation, binary invocation, and
   termination (`ACTION_FORK`, `ACTION_EXEC`, `ACTION_EXIT`) with
   `(pid, fork_time_usec)` identity, parent lineage, UID/GID context, executable
   path, binary `SHA-256` / IMA digest, mount filesystem type (`squashfs`,
   `ext4`, `tmpfs`, `memfd`), command arguments, container namespace, and
   optional correlation with `gNSI.acctz` (`acctz_session_id`).
2. **`KernelEvent`**: Dynamic kernel/OS extension and instrumentation changes—
   including kernel module loads/unloads, eBPF program loads,
   `kprobe`/`uprobe` attachments, packet-filter (Netfilter) hook registrations,
   and security-critical `sysctl` mutations, with signature verification status.
3. **`FilesystemEvent`**: Read-only rootfs remount transitions (`REMOUNT_RW`),
   `dm-verity` / `fs-verity` / IMA appraisal failures, modifications to critical
   system files (`/boot`, `/etc/passwd`, `/etc/shadow`, `/etc/sudoers`,
   `/etc/ssh/sshd_config`), and package manifest drift.
4. **`PrivilegeEvent`**: UID/EUID transitions, capability escalations,
   container/namespace escapes (`unshare` / `setns`), and execution of unlinked
   or anonymous in-memory binaries.
5. **`SocketEvent`**: Semantic socket `BIND`, `LISTEN`, and `CONNECT` operations
   across host, management, and revenue/data-plane VRFs—whether executed via
   standard kernel socket syscalls or intercepted via user-space shims (such as
   `LD_PRELOAD`)—with owning PID and executable lineage.
6. **`AgentHeartbeat`**: Periodic stream heartbeat carrying agent version,
   binary `SHA-256` digest, active probes, dropped-event counters, and CPU/RSS
   utilization.

### Ring-Buffer History & Cursor Resumption

Devices must maintain a bounded in-memory ring buffer of `EventResponse` records
(with zero routine writes to onboard flash/SSD) so that newly connected or
reconnecting Collectors can resume streaming via `EventRequest.timestamp`.
Timestamps assigned to `EventResponse.timestamp` in the ring buffer must be
strictly increasing (if multiple events occur within the same clock tick or
across an NTP clock adjustment, the System increments `nanos` by at least 1
relative to the preceding record). If the requested `timestamp` predates the
oldest record in the ring buffer or if records were shed under load, the first
returned `EventResponse` sets `history_istruncated = true`.

### Control-Plane Safety & Co-Attestation

Implementations of `gNSI.auditz` on network device supervisors must operate
strictly asynchronously (fail-open) without blocking control-plane syscalls or
routing daemons (`BGP`, `ISIS`, `gNMI`, `P4Runtime`). Event severity
classification and correlation are performed by the receiving off-box security
backends; when a backend receives a TCB-impacting `EventResponse` (such as an
unexpected kernel module load, read-only remount, or `dm-verity` failure), it
may immediately trigger an on-demand cryptographic attestation quote (`Attestz`)
on the device.

## Why a Dedicated `gNSI` Service Instead of `gNMI` Telemetry Paths or `gNOI.healthz`?

OpenConfig separates runtime security observability into a **`gNSI` gRPC
streaming service for discrete security events** (`auditz.proto`) paired with
**`gNMI` YANG paths for operational state and cumulative counters**
(`openconfig-gnsi-auditz.yang`), following the exact precedent established by
`gNSI.acctz`:

1. **Event-Log Semantics vs. State-Tree Coalescing (`gNMI`):** `gNMI`
   synchronizes the current state of a hierarchical tree (`ON_CHANGE` / `SAMPLE`)
   and permits coalescing rapid updates to the same path. Ephemeral
   control-plane exploits (such as a `< 5 ms` `/bin/sh -c` execution, a
   transient `remount,rw` followed by `remount,ro`, or a rapid kernel module
   load/unload) are coalesced into a delete/no-op in a state tree, whereas
   `gNSI.auditz` preserves every discrete transition in an append-only event
   stream.
2. **Stateful Cursor Replay Across Disconnects:** When a `gNMI.Subscribe`
   session drops and reconnects, it resyncs current state (`sync_response: true`)
   without replaying historical transitions missed during the outage (even when
   carrying protobuf payloads). `Auditz.EventSubscribe` replays missed events
   from the on-box ring buffer via `EventRequest.timestamp` and flags
   wrap-around via `history_istruncated`.
3. **Protecting the Central NOS State Database:** Modeling a high-rate
   ring buffer as a keyed YANG list in `gNMI` would thrash the central NOS
   state database with continuous list insertions and deletions. A dedicated
   `gNSI` service streams directly from the bounded in-memory ring buffer
   without YANG tree serialization or state database overhead.
4. **Separation from Operational Fault Management (`gNOI.healthz`):**
   `gNOI.healthz` (`Get`, `List`, `Check`, `Artifact`) reports whether hardware
   or software components are `STATUS_HEALTHY` or `STATUS_UNHEALTHY` and
   retrieves post-mortem diagnostic bundles (`show tech`, core dumps). During an
   active control-plane intrusion (e.g., an unauthorized shell or listening
   socket), components remain operationally `STATUS_HEALTHY`. Moreover, keeping
   `Auditz` in `gNSI` allows `gNSI.authz` to restrict sensitive audit records
   (`argv`, UIDs, `acctz_session_id`) to security collectors independently of
   general `gNOI` / `gNMI` operational subscribers.

## OpenConfig Data Models for `gNSI.auditz`

YANG data models for `gNSI.auditz` operational state and counters are defined in
the [OpenConfig public repository](https://github.com/openconfig/public/tree/master/release/models/gnsi)
(`openconfig-gnsi-auditz.yang`). Documentation for OpenConfig models, including
searchable paths and tree representations, is available at
[OpenConfig.net](https://openconfig.net/projects/models/).
