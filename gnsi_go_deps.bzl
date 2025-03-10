# Copyright 2025 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
"""Go dependencies to build gnmi."""

load("@bazel_gazelle//:deps.bzl", "go_repository")

def gnsi_go_deps():
    """Declare the third-party Go dependencies necessary to build gnmi"""
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:G2HAfAmvm/GcKan2oOQpBXOd2tT2G57ZnZGWa1PxPBQ=",
        version = "v4.1.1",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:UL815xU9SqsFlibzuggzjXhog7bL6oX9BbNZnL2UFvs=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:QVw89YDxXxEe+l8gU8ETbOasdwEV+avkR75ZzsVV9WI=",
        version = "v0.0.0-20240905190251-b4127c9b8d78",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:vPfJZCkob6yTMEgS+0TwfTUfbHjfy/6vOJ8hUWX/uXE=",
        version = "v0.13.1",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:tntQDh69XqOCOZsDz0lVJQez/2L6Uu2PdjCQwWCJ3bM=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        sum = "h1:6pFjapn8bFcIbiKo3XT4j/BhANplGihG6tvd+8rYgrY=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:1+mZ9upx1Dh6FmUTFR1naJ77miKiXgALjWOZ3NVFPmY=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_detectors_gcp",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp",
        sum = "h1:cZpsGsWTIFKymTA0je7IIvi1O7Es7apb9CF3EQlOcfE=",
        version = "v1.24.2",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        importpath = "github.com/mitchellh/go-wordwrap",
        sum = "h1:TLuKupo69TCn6TQSyGxwI1EblZZEsQ0vMlAFQflz0v0=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_openconfig_gnmi",
        importpath = "github.com/openconfig/gnmi",
        sum = "h1:tv9HygDMXnoGyWuLmNCodMV2+PK6+uT/ndAxDVzsUUQ=",
        version = "v0.0.0-20220617175856-41246b1b3507",
    )
    go_repository(
        name = "com_github_openconfig_goyang",
        importpath = "github.com/openconfig/goyang",
        sum = "h1:5MyIz4bN4vpH6aHDN339bkWXAjTkhg1ZKMhR4aIi5Rk=",
        version = "v0.0.0-20200115183954-d0a48929f0ea",
    )
    go_repository(
        name = "com_github_openconfig_grpctunnel",
        importpath = "github.com/openconfig/grpctunnel",
        sum = "h1:tSMKTQlWcHhdxQhn6P9myhLcoI+SzxF9e6hHWstCagU=",
        version = "v0.0.0-20220524190229-125331eabdde",
    )
    go_repository(
        name = "com_github_openconfig_ygot",
        importpath = "github.com/openconfig/ygot",
        sum = "h1:kJJFPBrczC6TDnz/HMlFTJEdW2CuyUftV13XveIukg0=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_planetscale_vtprotobuf",
        importpath = "github.com/planetscale/vtprotobuf",
        sum = "h1:GFCKgmp0tecUJ0sJuv4pzYCqS9+RGSn52M3FUwPs+uo=",
        version = "v0.6.1-0.20240319094008-0393e58bdf10",
    )
    go_repository(
        name = "com_github_protocolbuffers_txtpbfmt",
        importpath = "github.com/protocolbuffers/txtpbfmt",
        sum = "h1:AKJY61V2SQtJ2a2PdeswKk0NM1qF77X+julRNYRxPOk=",
        version = "v0.0.0-20220608084003-fc78c767cd6a",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:UxK4uu/Tn+I3p2dYWTfiX4wva7aYlKixAHn3fyqngqo=",
        version = "v0.5.2",
    )
    go_repository(
        name = "dev_cel_expr",
        importpath = "cel.dev/expr",
        sum = "h1:RwRhoH17VhAu9U5CMvMhH1PDVgf0tuz9FT+24AfMLfU=",
        version = "v0.16.2",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_detectors_gcp",
        importpath = "go.opentelemetry.io/contrib/detectors/gcp",
        sum = "h1:G1JQOreVrfhRkner+l4mrGxmfqYCAuy76asTDAo0xsA=",
        version = "v1.31.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:NsJcKPIW0D0H3NgzPDHmo0WW6SptzPdqg/L1zsIm2hY=",
        version = "v1.31.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:FSErL0ATQAmYHUIzSezZibnyVlft1ybhy4ozRPcF2fE=",
        version = "v1.31.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:xLY3abVHYZ5HSfOg3l2E5LUj2Cwva5Y7yGxnSW9H5Gk=",
        version = "v1.31.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_metric",
        importpath = "go.opentelemetry.io/otel/sdk/metric",
        sum = "h1:i9hxxLJF/9kkvfHppyLL55aW7iIJz4JjxTeYusH7zMc=",
        version = "v1.31.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:ffjsj1aRouKewfr85U2aGagJ46+MvodynlQ1HYdmJys=",
        version = "v1.31.0",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:qnrhhl4uoNFepTqE28u11llFcDH07Z6r/cQxpGR97A4=",
        version = "v0.0.0-20210811021853-ddbe55d93216",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:fVoAXEKA4+yufmbdVYv+SE73+cPZbbbe8paLsHfkK+U=",
        version = "v0.0.0-20241015192408-796eee8c2d53",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:3UsHvIr4Wc2aW4brOaSCmcxh9ksica6fHEr8P1XhkYw=",
        version = "v0.0.0-20250106144421-5f5ef82da422",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:U3S9QEtbXC0bYNvRtcoklF3xGtLViumSYxWykJS+7AU=",
        version = "v1.69.2",
    )
    go_repository(
        name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
        importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
        sum = "h1:F29+wU6Ee6qgu9TddPgooOdaqsxTMunOoj8KA5yuS5A=",
        version = "v1.5.1",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        build_file_proto_mode = "disable_global",  # Manually added to fix build. See https://github.com/golang/protobuf/issues/1611
        importpath = "google.golang.org/protobuf",
        sum = "h1:R8FeyR1/eLmkutZOM5CWghmo5itiG9z0ktFlTVLuTmU=",
        version = "v1.36.2",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:euUpcYgM8WcP71gNpTqQCn6rC2t6ULUPiOzfWaXVVfc=",
        version = "v0.32.0",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:zY54UmvipHiNd+pm+m0x9KhZ9hl1/7QNMyxXbc6ICqA=",
        version = "v0.17.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:Mb7Mrk043xzHgnRM88suvJFwzVrRfHEHJEl5/71CKw0=",
        version = "v0.34.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:PbgcYx2W7i4LvjJWEbf0ngHV6qJYr86PkAV3bXdLEbs=",
        version = "v0.23.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:3NQrjDixjgGwUOCaF8w2+VYHv0Ve/vGYSbdkTa98gmQ=",
        version = "v0.10.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:TPYlXGxvx1MGTn2GiZDhnjPA9wZzZeGKHHmKhHYvgaU=",
        version = "v0.29.0",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:/Ts8HFuMR2E6IP/jlo7QVLZHggjKQbhu/7H0LJFr3Gg=",
        version = "v0.28.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:zyQAAkrwaneQ066sspRyJaG9VNi/YJ1NfzcGB3hZ/qo=",
        version = "v0.21.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:vU5i/LfpvrRCpgM/VPfJLg5KjxD3E+hfT1SH+d9zLwg=",
        version = "v0.21.1-0.20240508182429-e35e4ccd0d2d",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
        version = "v0.0.0-20200804184101-5ec99f83aff1",
    )
