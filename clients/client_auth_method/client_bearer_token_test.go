package client_auth_method_test

import (
	"context"
	"github.com/fastenhealth/fasten-sources/clients/client_auth_method"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/fasten-sources/clients/testutils"
	definitionsModels "github.com/fastenhealth/fasten-sources/definitions/models"
	"github.com/fastenhealth/fasten-sources/pkg/models/catalog"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
	"testing"
	"time"
)

func TestClientBearerTokenIntrospectToken_Epic(t *testing.T) {
	//setup
	endpointDef := definitionsModels.LighthouseSourceDefinition{
		PatientAccessEndpoint: &catalog.PatientAccessEndpoint{
			IntrospectionEndpoint: "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/introspect",
		},
	}

	testHttpClient := testutils.OAuthVcrSetup(t, false)

	tokenData := &oauth2.Token{
		TokenType:    "Bearer",
		AccessToken:  "PLACEHOLDER_ACCESS_TOKEN",
		RefreshToken: "PLACEHOLDER_REFRESH_TOKEN",
		Expiry:       time.Now().Add(5 * time.Minute),
	}

	//test
	response, err := client_auth_method.ClientBearerTokenAuthIntrospectToken(
		context.TODO(),
		logrus.New(),
		endpointDef,
		tokenData,
		models.TokenIntrospectTokenTypeRefresh,
		"PLACEHOLDER_REFRESH_TOKEN",
		testHttpClient,
	)
	require.NoError(t, err)

	//assert
	assert.NotNil(t, response)
	assert.True(t, response.Active)
	assert.Equal(t, "patient/AllergyIntolerance.read patient/Binary.read patient/CarePlan.read patient/CareTeam.read patient/Condition.read patient/Device.read patient/DiagnosticReport.read patient/DocumentReference.read patient/Encounter.read patient/Goal.read patient/Immunization.read patient/Location.read patient/Medication.read patient/MedicationRequest.read patient/Observation.read patient/Organization.read patient/Patient.read patient/Practitioner.read patient/PractitionerRole.read patient/Procedure.read patient/Provenance.read patient/RelatedPerson.read launch/patient offline_access", response.Scope)
	assert.Equal(t, "1e3ce324-5c1b-45a9-8ad1-7a0c1f51df43", response.ClientID)
}

func TestClientBearerTokenUserInfoGetPatientId_Epic(t *testing.T) {
	t.Skipf("Skipping test since Epic does not support userinfo")

	userInfoEndpoint := "https://fhir.epic.com/interconnect-fhir-oauth/oauth2/userinfo"

	testHttpClient := testutils.OAuthVcrSetup(t, false)

	tokenData := &oauth2.Token{
		TokenType:    "Bearer",
		AccessToken:  "PLACEHOLDER_ACCESS_TOKEN",
		RefreshToken: "PLACEHOLDER_REFRESH_TOKEN",
		Expiry:       time.Now().Add(5 * time.Minute),
	}

	patientID, err := client_auth_method.ClientBearerTokenAuthUserInfoGetPatientId(context.TODO(), logrus.New(), tokenData, userInfoEndpoint, testHttpClient)
	require.NoError(t, err)
	assert.Equal(t, "test-patient-id", patientID)
}
