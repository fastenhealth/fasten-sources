# fasten-sources

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

