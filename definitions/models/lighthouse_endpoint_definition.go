package models

import (
	"github.com/fastenhealth/fasten-sources/pkg"
	"github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"strings"
)

// LighthouseSource
// https://vteapif1.aetna.com/fhirdemo/v1/fhirserver_auth/oauth2/authorize?
// response_type=code&
// client_id=5c47935b-29a7-4346-a01b-649a11d94dc5&
// scope=launch%2Fpatient%20patient%2FObservation.rs%20patient%2FPatient.rs%20offline_access&
// redirect_uri=https://sandbox.fastenhealth.com/api/callback/aetna&
// aud=https://vteapif1.aetna.com/fhirdemo&
// code_challenge=XXXXXX&
// code_challenge_method=XXXXX
//
// Similar in functionality to https://build.fhir.org/ig/HL7/smart-app-launch/conformance.html#example-request
// /apis/fhir/.well-known/smart-configuration
type LighthouseSourceDefinition struct {
	BrandId  string `json:"brand_id,omitempty" yaml:"-" validate:"omitempty,uuid"`
	PortalId string `json:"portal_id,omitempty" yaml:"-" validate:"omitempty,uuid"`

	*catalog.PatientAccessEndpoint `json:",inline" yaml:"-" validate:"required"`

	// Smart-On-FHIR configuration
	// https://build.fhir.org/ig/HL7/smart-app-launch/conformance.html#example-request

	Scopes                 []string `json:"scopes_supported" yaml:"scopes_supported" validate:"required"`
	GrantTypesSupported    []string `json:"grant_types_supported" yaml:"grant_types_supported" validate:"required"`
	ResponseType           []string `json:"response_types_supported" yaml:"response_types_supported" validate:"required"`
	ResponseModesSupported []string `json:"response_modes_supported" yaml:"response_modes_supported" validate:"required"`
	// If populated: PKCE is supported (can be used with Confidential true or false)
	CodeChallengeMethodsSupported []string `json:"code_challenge_methods_supported" yaml:"code_challenge_methods_supported" validate:"required"`

	//if enabled, requires client_secret to authenticate with provider (PKCE)
	Confidential bool `json:"confidential" yaml:"confidential"`
	//if enabled, will dynamically register client with provider (https://oauth.net/2/dynamic-client-registration/)
	DynamicClientRegistrationMode string `json:"dynamic_client_registration_mode" yaml:"dynamic_client_registration_mode"`
	//if true, requires CORS proxy/relay, as provider does not return proper response to CORS preflight
	CORSRelayRequired bool `json:"cors_relay_required" yaml:"cors_relay_required"`
	//the secret key prefix to use, if empty (default) will use the sourceType value
	SecretKeyPrefix string `json:"-" yaml:"secret_key_prefix"`

	//optional - required for some providers
	Issuer        string            `json:"issuer" yaml:"-" validate:"omitempty,http_url"`
	Audience      string            `json:"aud" yaml:"-"  validate:"omitempty,http_url"`
	Documentation string            `json:"-" yaml:"documentation"  validate:"omitempty"`
	ClientHeaders map[string]string `json:"-" yaml:"client_headers"  validate:"omitempty"`

	//Client configuration
	MissingOpPatientEverything bool `json:"-" yaml:"missing_op_patient_everything"  validate:"omitempty"`
	//can only be set if MissingOpPatientEverything is true
	CustomOpPatientEverything string   `json:"-" yaml:"custom_op_patient_everything"  validate:"omitempty"`
	ClientSupportedResources  []string `json:"-" yaml:"client_supported_resources"  validate:"omitempty"`

	//set by the Populate() function
	PlatformType pkg.PlatformType `json:"platform_type" yaml:"platform_type" validate:"required"`
	ClientId     string           `json:"client_id" yaml:"-" validate:"required"`
	//set by the Populate() function - lighthouse url the provider will redirect to (registered with App)
	RedirectUri string `json:"redirect_uri" yaml:"-" validate:"required,http_url"`
}

func (def *LighthouseSourceDefinition) Populate(
	endpoint *catalog.PatientAccessEndpoint,
	env pkg.FastenLighthouseEnvType,
	clientIdLookup map[pkg.PlatformType]string,
) {

	//must be done first
	def.PatientAccessEndpoint = endpoint
	def.PlatformType = endpoint.GetPlatformType() //this handles platform type overrides

	//Hide sources for platform types which are still under development
	//if platformType == pkg.PlatformTypeAllscripts {
	//	sourceDef.Hidden = true
	//}

	if !(def.PlatformType == pkg.PlatformTypeCigna ||
		def.PlatformType == pkg.PlatformTypeNextgen ||
		def.PlatformType == pkg.PlatformTypeVahealth) {
		//most providers use the same url for API endpoint and Audience. These are the exceptions
		def.Audience = def.Url
	}

	if def.PlatformType == pkg.PlatformTypeCareevolution ||
		def.PlatformType == pkg.PlatformTypeAnthem ||
		def.PlatformType == pkg.PlatformTypeEclinicalworks ||
		def.PlatformType == pkg.PlatformTypeMedhost {
		//remove trailing slash for audience for CareEvolution & Anthem
		def.Audience = strings.TrimSuffix(def.Audience, "/")
	}

	if def.PlatformType == pkg.PlatformTypeCerner {
		def.IntrospectionEndpoint = "https://authorization.cerner.com/tokeninfo"
	}

	if def.PlatformType == pkg.PlatformTypeKaiser && env == pkg.FastenLighthouseEnvSandbox {
		def.Scopes = append(def.Scopes, "sandbox")
	}

	//Common defaults. All customizations should be above this line
	def.Issuer = def.Url
	// retrieve client-id, if available
	if clientId, clientIdOk := clientIdLookup[def.PlatformType]; clientIdOk {
		def.ClientId = clientId
	}
	def.RedirectUri = pkg.GetCallbackEndpoint(string(def.PlatformType))

}
