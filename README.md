# gNSI - gRPC Network Security Interface

[![License: Apache](https://img.shields.io/badge/license-Apache%202-blue)](https://opensource.org/licenses/Apache-2.0)
[![GitHub Super-Linter](https://github.com/openconfig/gnsi/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/marketplace/actions/super-linter)

A repository which contains security infrastructure services
necessary for safe operations of an OpenConfig platform. These
services include:

  1. Authorization protocol buffer
  2. Accounting protocol buffer
  3. Certificate management
  4. Console access management
  5. Secure Shell (ssh) certificate/key management
  6. Associated YANG models for telemetry collection of gNSI systems.

## Releases

   -  A gNSI server is expected to support single version of gNSI
      (e.g. if v2 releases, then the server will only support v1 until
      it has support for v2, at which point it will drop support for v1).
   -  A gNSI server is expected to be pinned to a release tag on the 
      gNSI repository.

