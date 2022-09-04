# gNMI Accounting Protocol

## Objective

This proto definition serves to describe a method of transfering
accounting records from a network device (or other system) to a
streaming telemetry collection system, primarily over a gNMI
transport connection.

## Method of Operation

Accounting Records are available at a gNMI origin:

```
   gnmi.accounting
```

Records will be streamed to the receiver as individual Record
messages as they are defined in the gnsi.accounting protocol buffer
definition.

Each Record() message contains a timestamp element, this represents the
timeat which the accounted event occured, local to the device which sends
the message.
