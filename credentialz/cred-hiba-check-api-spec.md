# gNSI Credentialz w/ hiba-chk API Specification

This is an API which creates a contract for how gNSI Credentialz and hiba-chk will be expected to behave together. The purpose of this API is to make it clear how these two different components will map messages passed over Credentialz to hiba-chk. A vendor implementing gNSI Credentialz will also need to implement this API for one or more versions to ensure support for hiba-chk in gNSI credentialz.

If changes are made to either gNSI Credentialz or hiba-chk which change how arguments are mapped this API will need to be incremented to handle this scenario. A vendor will then need to produce a new image which fulfills the new API version.

## API v1.0
- Applicable to gNSI credentialz with proto version **0.2.0**
- Applicable to hiba-chk utility with commit `cae606dafc5692240ac49441fd57653d7aade99b`

Messages are mapped as follows:

- Credentialz message `ServerKeysRequest::AuthenticationArtifacts::certificate` is mapped to hiba-chk as: `-i {ServerKeysRequest::AuthenticationArtifacts::certificate}`. If multiple certificates are provided then the last certificate provided will be used for this mapping.

The full set of mappings and command-line arguments is
`hiba-chk -i {ServerKeysRequest::AuthenticationArtifacts::certificate} -r %u %k`
