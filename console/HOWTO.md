# How-To Guide to gNSI.console API

## Console access

There are two methods to configure a password:

1) Directly on the device.

To change password execute the following command after logging-in to the device
using `ssh` or directly using a console (for example a RS232-based one or
simmilar method):

```
$ echo "TeStP_w0rD" | passwd ${account} --stdin
```

2) Using `gNSI.console` API

* Start streaming RPC call to the target device.

```
stream := MutateAccountPassword()
```
* Send a password change request message to the target device.

```
stream.Send(
    MutateAccountPasswordRequest {
        set_password: SetPasswordRequest {
            pairs: AccountPassword {
                account: "user",
                password: "password",
            }
        }
    }
)

resp := stream.Receive()
```

* Check if the new password 'works'

```
```

* Finalize the operation

```
stream.Send(
    MutateAccountPasswordRequest {
        finalize: FinalizeRequest {}
    }
)
```

# Apropos

## SSH with Password-based authentication

> **_NOTE:_**  The SSH password authentication method is strongly discouraged.

The password cannot be set using `gNSI.ssh` API. It can be changed using
`gNSI.console` API.

