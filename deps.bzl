
#
# Copyright 2024 Google Inc. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
load("@bazel_gazelle//:deps.bzl", "go_repository")

def gnsi_deps():
  if not native.existing_rule("com_github_openconfig_gnmi"):
    go_repository(
      name = "com_github_openconfig_gnmi",
      build_directives = [
        "gazelle:proto_import_prefix github.com/openconfig/gnmi",
      ],
      build_file_generation = "on",
      importpath = "github.com/openconfig/gnmi",
      sum = "h1:H7pLIb/o3xObu3+x0Fv9DCK7TH3FUh7mNwbYe+34hFw=",
      version = "v0.11.0",
    )
  if not native.existing_rule("com_github_kylelemons_godebug"):
    go_repository(
      name = "com_github_kylelemons_godebug",
      importpath = "github.com/kylelemons/godebug",
      sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
      version = "v1.1.0",
    )
  if not native.existing_rule("com_github_openconfig_goyang"):
    go_repository(
      name = "com_github_openconfig_goyang",
      importpath = "github.com/openconfig/goyang",
      sum = "h1:Z95LskKYk6nBYOxHtmJCu3YEKlr3pJLWG1tYAaNh3yU=",
      version = "v0.2.9",
    )
  if not native.existing_rule("com_github_openconfig_ygot"):
    go_repository(
      name = "com_github_openconfig_ygot",
      build_directives = [
        "gazelle:proto_import_prefix github.com/openconfig/ygot",
      ],
      importpath = "github.com/openconfig/ygot",
      sum = "h1:EKaeFhx1WwTZGsYeqipyh1mfF8y+z2StaXZtwVnXklk=",
      version = "v0.13.1",
    )
