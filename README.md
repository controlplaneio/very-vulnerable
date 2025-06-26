# very-vulnerable

Note: Please do not merge Github dependabot fixes to this repo. It is suppose to be vulnerable!

TODO: Publish to dockerhub https://hub.docker.com/u/controlplane

## About

This repo is filled with bad security practices designed to help test security scanners and CI/CD pipelines. 

- Bad Go code
    - Utilizing vulnerable packages
    - following vulnerable code paths
- Packed into a container
    - With a very old and vulnerable base container
    - With hard coded secret files

## Vulns

- GO-2023-1737 (https://pkg.go.dev/vuln/GO-2023-1737) and CVE-2023-29401 = [Gin](https://github.com/gin-gonic/gin) `FileAttachment() is vulnerable to Reflect File Download`
- Use of md5 hash
- Leaks SSH Private key and Github Personal Access Token
- Base Image contains a number of vulns
    - https://hub.docker.com/layers/library/debian/trixie-20240904-slim/images/sha256-d80214935426b02c162971a3213e1a14312f01a632c68f0ef09f27dcd6f33267

## [osv-scanner](https://github.com/google/osv-scanner) results:

### GO

```
╭─────────────────────────────────────┬──────┬───────────┬────────────────────────────┬───────────────────────────────────┬────────╮
│ OSV URL                             │ CVSS │ ECOSYSTEM │ PACKAGE                    │ VERSION                           │ SOURCE │
├─────────────────────────────────────┼──────┼───────────┼────────────────────────────┼───────────────────────────────────┼────────┤
│ https://osv.dev/GO-2023-1737        │ 4.3  │ Go        │ github.com/gin-gonic/gin   │ 1.8.2                             │ go.mod │
│ https://osv.dev/GHSA-2c4m-59x9-fr2g │      │           │                            │                                   │        │
│ https://osv.dev/GHSA-3vp4-m3rf-835h │ 5.6  │ Go        │ github.com/gin-gonic/gin   │ 1.8.2                             │ go.mod │
│ https://osv.dev/GO-2023-1571        │ 7.5  │ Go        │ golang.org/x/net           │ 0.4.0                             │ go.mod │
│ https://osv.dev/GHSA-vvpx-j8f3-3w6h │      │           │                            │                                   │        │
│ https://osv.dev/GO-2023-2102        │ 7.5  │ Go        │ golang.org/x/net           │ 0.4.0                             │ go.mod │
│ https://osv.dev/GHSA-4374-p667-p6c8 │      │           │                            │                                   │        │
│ https://osv.dev/GO-2024-2687        │ 5.3  │ Go        │ golang.org/x/net           │ 0.4.0                             │ go.mod │
│ https://osv.dev/GHSA-4v7x-pqxf-cx7m │      │           │                            │                                   │        │
│ https://osv.dev/GHSA-qppj-fm5r-hxr3 │ 6.9  │ Go        │ golang.org/x/net           │ 0.4.0                             │ go.mod │
│ https://osv.dev/GO-2025-3749        │      │ Go        │ stdlib                     │ 1.24.2                            │ go.mod │
│ https://osv.dev/GO-2025-3750        │      │ Go        │ stdlib                     │ 1.24.2                            │ go.mod │
├─────────────────────────────────────┼──────┼───────────┼────────────────────────────┼───────────────────────────────────┼────────┤
│ Uncalled vulnerabilities            │      │           │                            │                                   │        │
├─────────────────────────────────────┼──────┼───────────┼────────────────────────────┼───────────────────────────────────┼────────┤
│ https://osv.dev/GO-2021-0356        │ 7.5  │ Go        │ golang.org/x/crypto        │ 0.0.0-20211215153901-e495a2d5b3d3 │ go.mod │
│ https://osv.dev/GHSA-8c26-wmh5-6g9v │      │           │                            │                                   │        │
│ https://osv.dev/GO-2023-2402        │ 5.9  │ Go        │ golang.org/x/crypto        │ 0.0.0-20211215153901-e495a2d5b3d3 │ go.mod │
│ https://osv.dev/GHSA-45x7-px36-x8w8 │      │           │                            │                                   │        │
│ https://osv.dev/GO-2024-2961        │      │ Go        │ golang.org/x/crypto        │ 0.0.0-20211215153901-e495a2d5b3d3 │ go.mod │
│ https://osv.dev/GO-2024-3321        │ 9.1  │ Go        │ golang.org/x/crypto        │ 0.0.0-20211215153901-e495a2d5b3d3 │ go.mod │
│ https://osv.dev/GHSA-v778-237x-gjrc │      │           │                            │                                   │        │
│ https://osv.dev/GO-2025-3487        │ 7.5  │ Go        │ golang.org/x/crypto        │ 0.0.0-20211215153901-e495a2d5b3d3 │ go.mod │
│ https://osv.dev/GHSA-hcg3-q754-cr77 │      │           │                            │                                   │        │
│ https://osv.dev/GO-2023-1988        │ 6.1  │ Go        │ golang.org/x/net           │ 0.4.0                             │ go.mod │
│ https://osv.dev/GHSA-2wrh-6pvc-2jm9 │      │           │                            │                                   │        │
│ https://osv.dev/GO-2024-3333        │      │ Go        │ golang.org/x/net           │ 0.4.0                             │ go.mod │
│ https://osv.dev/GO-2025-3503        │ 4.4  │ Go        │ golang.org/x/net           │ 0.4.0                             │ go.mod │
│ https://osv.dev/GHSA-qxp5-gwg8-xv66 │      │           │                            │                                   │        │
│ https://osv.dev/GO-2025-3595        │ 5.3  │ Go        │ golang.org/x/net           │ 0.4.0                             │ go.mod │
│ https://osv.dev/GHSA-vvgc-356p-c3xw │      │           │                            │                                   │        │
│ https://osv.dev/GO-2024-2611        │ 7.5  │ Go        │ google.golang.org/protobuf │ 1.28.1                            │ go.mod │
│ https://osv.dev/GHSA-8r3f-844c-mc37 │      │           │                            │                                   │        │
│ https://osv.dev/GO-2025-3751        │      │ Go        │ stdlib                     │ 1.24.2                            │ go.mod │
╰─────────────────────────────────────┴──────┴───────────┴────────────────────────────┴───────────────────────────────────┴────────╯
```

### Docker

```
Container Scanning Result (Debian GNU/Linux trixie/sid):
Total 22 packages affected by 53 known vulnerabilities (3 Critical, 10 High, 15 Medium, 1 Low, 24 Unknown) from 2 ecosystems.
16 vulnerabilities can be fixed.


Go
╭────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│ Source:artifact:very-vulnerable                                                                                                │
├────────────────────────────┬───────────────────────────────────┬───────────────┬────────────┬──────────────────┬───────────────┤
│ PACKAGE                    │ INSTALLED VERSION                 │ FIX AVAILABLE │ VULN COUNT │ INTRODUCED LAYER │ IN BASE IMAGE │
├────────────────────────────┼───────────────────────────────────┼───────────────┼────────────┼──────────────────┼───────────────┤
│ github.com/gin-gonic/gin   │ 1.8.2                             │ Fix Available │          2 │ # 5 Layer        │ --            │
│ golang.org/x/crypto        │ 0.0.0-20211215153901-e495a2d5b3d3 │ Fix Available │          5 │ # 5 Layer        │ --            │
│ golang.org/x/net           │ 0.4.0                             │ Fix Available │          8 │ # 5 Layer        │ --            │
│ google.golang.org/protobuf │ 1.28.1                            │ Fix Available │          1 │ # 5 Layer        │ --            │
╰────────────────────────────┴───────────────────────────────────┴───────────────┴────────────┴──────────────────┴───────────────╯
Debian
╭─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│ Source:os:var/lib/dpkg/status                                                                                                           │
├────────────────┬──────────────────────────┬──────────────────┬────────────┬─────────────────────────┬──────────────────┬────────────────┤
│ SOURCE PACKAGE │ INSTALLED VERSION        │ FIX AVAILABLE    │ VULN COUNT │ BINARY PACKAGES (COUNT) │ INTRODUCED LAYER │ IN BASE IMAGE  │
├────────────────┼──────────────────────────┼──────────────────┼────────────┼─────────────────────────┼──────────────────┼────────────────┤
│ bash           │ 5.2.32-1                 │ No fix available │          1 │ bash                    │ # 0 Layer        │ library/debian │
│ coreutils      │ 9.4-3.1                  │ No fix available │          2 │ coreutils               │ # 0 Layer        │ library/debian │
│ db5.3          │ 5.3.28+dfsg2-7           │ No fix available │          1 │ libdb5.3t64             │ # 0 Layer        │ library/debian │
│ glibc          │ 2.40-2                   │ No fix available │          6 │ libc-bin, libc6         │ # 0 Layer        │ library/debian │
│ gnupg2         │ 2.2.43-8                 │ No fix available │          1 │ gpgv                    │ # 0 Layer        │ library/debian │
│ gnutls28       │ 3.8.6-2                  │ No fix available │          1 │ libgnutls30t64          │ # 0 Layer        │ library/debian │
│ libcap2        │ 1:2.66-5                 │ No fix available │          1 │ libcap2                 │ # 0 Layer        │ library/debian │
│ libgcrypt20    │ 1.11.0-6                 │ No fix available │          1 │ libgcrypt20             │ # 0 Layer        │ library/debian │
│ libtasn1-6     │ 4.19.0-3+b2              │ No fix available │          1 │ libtasn1-6              │ # 0 Layer        │ library/debian │
│ libzstd        │ 1.5.6+dfsg-1             │ No fix available │          1 │ libzstd1                │ # 0 Layer        │ library/debian │
│ ncurses        │ 6.5-2                    │ No fix available │          2 │ libtinfo6... (3)        │ # 0 Layer        │ library/debian │
│ openssl        │ 3.3.1-7                  │ No fix available │          6 │ libssl3t64... (2)       │ # 0 Layer        │ library/debian │
│ pam            │ 1.5.3-7                  │ No fix available │          4 │ libpam-modules... (4)   │ # 0 Layer        │ library/debian │
│ perl           │ 5.38.2-5                 │ No fix available │          3 │ perl-base               │ # 0 Layer        │ library/debian │
│ systemd        │ 256.5-1                  │ No fix available │          1 │ libsystemd0... (2)      │ # 0 Layer        │ library/debian │
│ tzdata         │ 2024a-4                  │ No fix available │          3 │ tzdata                  │ # 0 Layer        │ library/debian │
│ xz-utils       │ 5.6.2-2                  │ No fix available │          1 │ liblzma5                │ # 0 Layer        │ library/debian │
│ zlib           │ 1:1.3.dfsg+really1.3.1-1 │ No fix available │          1 │ zlib1g                  │ # 0 Layer        │ library/debian │
╰────────────────┴──────────────────────────┴──────────────────┴────────────┴─────────────────────────┴──────────────────┴────────────────╯

Filtered Vulnerabilities:
╭─────────────┬───────────┬───────────────────────────┬─────────────────────┬────────────────╮
│ PACKAGE     │ ECOSYSTEM │ INSTALLED VERSION         │ FILTERED VULN COUNT │ FILTER REASONS │
├─────────────┼───────────┼───────────────────────────┼─────────────────────┼────────────────┤
│ apt         │ Debian    │ 2.9.8                     │                   1 │ Unimportant    │
│ coreutils   │ Debian    │ 9.4-3.1                   │                   2 │ Unimportant    │
│ glibc       │ Debian    │ 2.40-2                    │                   2 │ Unimportant    │
│ gnupg2      │ Debian    │ 2.2.43-8                  │                   1 │ Unimportant    │
│ gnutls28    │ Debian    │ 3.8.6-2                   │                   1 │ Unimportant    │
│ libgcrypt20 │ Debian    │ 1.11.0-6                  │                   2 │ Unimportant    │
│ pcre2       │ Debian    │ 10.42-4+b1                │                   1 │ Unimportant    │
│ perl        │ Debian    │ 5.38.2-5                  │                   2 │ Unimportant    │
│ shadow      │ Debian    │ 1:4.16.0-4                │                   2 │ Unimportant    │
│ systemd     │ Debian    │ 256.5-1                   │                   2 │ Unimportant    │
│ tar         │ Debian    │ 1.35+dfsg-3               │                   1 │ Unimportant    │
│ util-linux  │ Debian    │ 1:2.40.2-7                │                   1 │ Unimportant    │
│ util-linux  │ Debian    │ 1:4.16.0-2+really2.40.2-7 │                   1 │ Unimportant    │
│ util-linux  │ Debian    │ 2.40.2-7                  │                   1 │ Unimportant    │
╰─────────────┴───────────┴───────────────────────────┴─────────────────────┴────────────────╯
```
