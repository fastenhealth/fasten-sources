package catalog

import (
	"github.com/fastenhealth/fasten-sources/pkg"
)

type CatalogQueryOptions struct {
	Id                        string
	LighthouseEnvType         pkg.FastenLighthouseEnvType
	IncludeSuspendedEndpoints bool

	CachedEndpointsLookup *map[string]PatientAccessEndpoint
	CachedPortalsLookup   *map[string]PatientAccessPortal
}
