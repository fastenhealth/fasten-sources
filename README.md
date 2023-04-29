<p align="center">
  <a href="https://github.com/fastenhealth/fasten-onprem">
  <img width="400" alt="fasten_view" src="frontend/src/assets/banner/banner.png">
  </a>
</p>

# fasten-sources

[![CI](https://github.com/fastenhealth/fasten-sources/actions/workflows/ci.yaml/badge.svg)](https://github.com/fastenhealth/fasten-sources/actions/workflows/ci.yaml)
[![codecov](https://codecov.io/gh/fastenhealth/fasten-sources/branch/main/graph/badge.svg?token=0FD2L52DTK)](https://codecov.io/gh/fastenhealth/fasten-sources)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/fastenhealth/fasten-sources?style=flat-square)](https://github.com/fastenhealth/fasten-sources/releases/latest)
[![Discord Join](https://img.shields.io/badge/discord-join-blueviolet?style=flat-square&logo=discord)](https://discord.gg/Bykz6BAN8p)
[![Request Providers](https://img.shields.io/static/v1?label=request+providers&message=form&color=orange&style=flat-square)](https://forms.gle/4oU8372y4KyM8DbdA)
[![Join Mailing List](https://img.shields.io/static/v1?label=join&message=mailing+list&color=blue&style=flat-square)](https://forms.gle/SNsYX9BNMXB6TuTw6)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/analogj?style=flat-square)](https://github.com/sponsors/AnalogJ/)


The Fasten Sources is a library that defines medical provider metadata (`definitions` - OpenID Metadata documents)
and http clients (OAuth2/Smart-on-FHIR clients) which can be used to retrieve data from various Medical
Providers (`clients`).


> See [SOURCE_LIST.md](./SOURCE_LIST.md) for a full list of sources available in this repo. 


# Types

There are multiple protocols used by the Medical Provider industry to transfer patient data, the following mechanisms are the
ones that Fasten supports



| Definition Folder                                        | Description                                                                                                                                     |
|----------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| [clients/factory](./clients/factory)                     | Automatically created client initializer                                                                                                        |
| [clients/internal/base](./clients/internal/base])        | Manually created base OAuth clients for various FHIR versions (R3/R4). These are the base clients that all platforms inherit from.              |
| [clients/internal/platform](./clients/internal/platform) | Automatically created OAuth clients for various EMR platforms. **Manually created test files**                                                  |
| [clients/internal/sandbox](./clients/internal/sandbox)   | Automatically created OAuth clients for accessing test FHIR servers full of synthetic data. **Manually created test files**                     |
| [clients/internal/source](./clients/internal/source)     | Automatically created OAuth clients for accessing production data from various healthcare institutions. Usually inherit from Platform clients   |
| [definitions](./definitions/)                            | Automatically created definition files. These files are generated from files created by `fasten-sources-gen` and are used by Fasten Lighthouse. |

