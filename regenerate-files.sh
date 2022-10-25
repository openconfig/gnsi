#!/bin/bash
set -euo pipefail

BASE=$(bazel info bazel-genfiles)
GNSI_NS='github.com/openconfig/gnsi'
bazel build //accounting:all
cp "${BASE}"/accounting/acct_go_proto_/"${GNSI_NS}"/accounting/*.pb.go accounting/
bazel build //authz:all
cp "${BASE}"/authz/authz_go_proto_/"${GNSI_NS}"/authz/*.pb.go authz/
bazel build //certz:all
cp "${BASE}"/certz/certz_go_proto_/"${GNSI_NS}"/certz/*.pb.go certz/
bazel build //credentialz:all
cp "${BASE}"/credentialz/credentialz_go_proto_/"${GNSI_NS}"/credentialz/*.pb.go credentialz/
bazel build //pathz:all
cp "${BASE}"/pathz/pathz_go_proto_/"${GNSI_NS}"/pathz/*.pb.go pathz/
