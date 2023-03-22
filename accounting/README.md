# gNMI Accounting Protocol

## Objective

This proto definition serves to describe a method of transfering
accounting records from a network device (or other system) to a
streaming telemetry collection system, primarily over a gNMI
transport connection.

## Method of Operation

Accounting Records are available at a gNMI origin:
   gnmi.accounting

Records maybe streamed from a system either at request of the remote
caller, via the ClientStream() rpc, or at the request of the system
to a remote endpoint, via the ServerStream() rpc.

Configuration of the Accounting service is made through standard
gNMI methods using the defined YANG model.

Records will be streamed to the receiver as individual Record
messages as they are defined in the gnsi.accounting protocol buffer
definition.

Each Record() message contains a timestamp element, this represents the
time at which the accounted event occured, local to the device which sends
the message. This could be different from the time received at the distant
end, and the time the Record was emitted from the system.
