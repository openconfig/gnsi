# gNSI Accounting Protocol

## Objective

This proto definition serves to describe a method of transfering
accounting records from a System, which may be a network device, to a
remote collection service, primarily over a gRPC transport connection.

## Method of Operation

Accounting Records are available at a gNSI origin:
   gnsi.acctz

Records may be streamed from a system at request of the remote collector,
via the RecordSubscribe() service/rpc.

Configuration of the Accounting service is made through standard
gNxI methods using the defined YANG model.

Records will be streamed to the receiver as individual Record
messages as they are defined in the gnsi.acctz protocol buffer
definition.

Each Record() message contains a timestamp element, this represents the
time at which the accounted event occured, local to the system which sends
the message. This could be different from the time received at the Collector
and the time the Record was emitted from the system.

The stream continues as new records are received by the accounting subsystem,
ending when the gNSI session ends or an error occurs.

Devices should maintain a history of accounting records so that they can be
retrieved periodically by newly and already connected Collectors.  The depth
of this history should be configurable by the administrator.  The default
depth and configurability are subject to implementation support, but should
be documented.

## OpenConfig Extension for the gMNI gRPC-based Accounting telemetry
### gnsi-acctz.yang
An overview of the changes defined in the gnsi-acctz.yang file are shown below.

```txt
module: gnsi-acctz
  augment /oc-sys:system/oc-sys-grpc:grpc-servers/oc-sys-grpc:grpc-server:
    +--ro counters
       +--ro last-cleared-on?   oc-types:timeticks64
       +--ro client-counters
       |  +--ro history_istruncated?   oc-yang:counter64
       |  +--ro IdleTimeouts?          oc-yang:counter64
       |  +--ro RecordRequests?        oc-yang:counter64
       |  +--ro RecordResponses?       oc-yang:counter64
       +--ro source-counters
          +--ro source-records* [service type]
             +--ro service    service-request
             +--ro type       service-type
             +--ro records?   oc-yang:counter64
```
