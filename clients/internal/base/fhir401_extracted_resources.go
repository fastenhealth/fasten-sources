package base

import (
	"github.com/fastenhealth/gofhir-models/fhir401"
)

func (c *SourceClientFHIR401) ExtractResourceReference(resourceRaw interface{}) []string {
	resourceRefs := []string{}

	switch sourceResourceType := resourceRaw.(type) {

	case fhir401.AllergyIntolerance:
		// recorder can contain
		//- Practitioner
		//- PractitionerRole
		//- Patient
		//- RelatedPerson
		if sourceResourceType.Recorder != nil && sourceResourceType.Recorder.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Recorder.Reference)
		}

		// asserter can contain
		//- Practitioner
		//- PractitionerRole
		//- Patient
		//- RelatedPerson
		if sourceResourceType.Asserter != nil && sourceResourceType.Asserter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Asserter.Reference)
		}

	case fhir401.CarePlan:
		// encounter can contain
		//- Encounter
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
		}

		//author can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- CareTeam
		//- RelatedPerson
		if sourceResourceType.Author != nil && sourceResourceType.Author.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Author.Reference)
		}

		//contributor can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- CareTeam
		//- RelatedPerson
		if sourceResourceType.Contributor != nil {
			for _, r := range sourceResourceType.Contributor {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}

		}

		//careTeam can contain
		//- CareTeam
		if sourceResourceType.CareTeam != nil {
			for _, r := range sourceResourceType.CareTeam {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		break
	case fhir401.CareTeam:
		// encounter can contain
		//- Encounter
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
		}

		//participant[x].member can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- CareTeam
		//- RelatedPerson
		//participant[x].onBehalfOf can contain
		//- Organization
		if sourceResourceType.Participant != nil {
			for _, r := range sourceResourceType.Participant {
				if r.Member != nil && r.Member.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Member.Reference)
				}
				if r.OnBehalfOf != nil && r.OnBehalfOf.Reference != nil {
					resourceRefs = append(resourceRefs, *r.OnBehalfOf.Reference)
				}
			}
		}

		//managingOrganization
		//- Organization
		if sourceResourceType.ManagingOrganization != nil {
			for _, r := range sourceResourceType.ManagingOrganization {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		break
	case fhir401.Condition:
		// recorder can contain
		//- Practitioner
		//- PractitionerRole
		//- Patient
		//- RelatedPerson
		if sourceResourceType.Recorder != nil && sourceResourceType.Recorder.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Recorder.Reference)
		}

		// asserter can contain
		//- Practitioner
		//- PractitionerRole
		//- Patient
		//- RelatedPerson
		if sourceResourceType.Asserter != nil && sourceResourceType.Asserter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Asserter.Reference)
		}

		break
	case fhir401.DiagnosticReport:
		//basedOn[x] can contain
		//- CarePlan
		//- ImmunizationRecommendation
		//- MedicationRequest
		//- NutritionOrder
		//- ServiceRequest
		if sourceResourceType.BasedOn != nil {
			for _, r := range sourceResourceType.BasedOn {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		// performer[x] can contain
		//- Practitioner
		//- PractitionerRole
		//- Organization
		//- CareTeam
		if sourceResourceType.Performer != nil {
			for _, r := range sourceResourceType.Performer {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.ResultsInterpreter != nil {
			for _, r := range sourceResourceType.ResultsInterpreter {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.Specimen != nil {
			for _, r := range sourceResourceType.Specimen {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.Result != nil {
			for _, r := range sourceResourceType.Result {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.ImagingStudy != nil {
			for _, r := range sourceResourceType.ImagingStudy {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		break
	case fhir401.DocumentReference:
		//author[x] can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- CareTeam
		//- Device
		if sourceResourceType.Author != nil {
			for _, r := range sourceResourceType.Author {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//authenticator can contain
		//- Practitioner
		//- Organization
		//- PractitionerRole
		if sourceResourceType.Authenticator != nil && sourceResourceType.Authenticator.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Authenticator.Reference)
		}

		// custodian can contain
		//- Organization
		if sourceResourceType.Custodian != nil && sourceResourceType.Custodian.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Custodian.Reference)
		}

		// relatesTo.target
		//- DocumentReference
		if sourceResourceType.RelatesTo != nil {
			for _, r := range sourceResourceType.RelatesTo {
				if r.Target.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Target.Reference)
				}
			}
		}

		//content.attachment can contain
		//- Attachment

	case fhir401.Encounter:
		// basedOn[x] can contain
		//- ServiceRequest
		if sourceResourceType.BasedOn != nil {
			for _, r := range sourceResourceType.BasedOn {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//participant[x].individual can contain
		//- Practitioner
		//- PractitionerRole
		//- RelatedPerson
		if sourceResourceType.Participant != nil {
			for _, r := range sourceResourceType.Participant {
				if r.Individual != nil && r.Individual.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Individual.Reference)
				}
			}
		}

		//reasonReference[x] can contain
		//- Condition
		//- Procedure
		//- Observation
		//- ImmunizationRecommendation
		if sourceResourceType.ReasonReference != nil {
			for _, r := range sourceResourceType.ReasonReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//hospitalization.origin can contain
		//- Location
		//- Organization
		if sourceResourceType.Hospitalization != nil && sourceResourceType.Hospitalization.Origin != nil && sourceResourceType.Hospitalization.Origin.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Hospitalization.Origin.Reference)
		}

		//hospitalization.destination can contain
		//- Location
		//- Organization
		//resourceRefs.push(resourceRaw.hospitalization?.destination?.reference)
		if sourceResourceType.Hospitalization != nil && sourceResourceType.Hospitalization.Destination != nil && sourceResourceType.Hospitalization.Destination.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Hospitalization.Destination.Reference)
		}

		//location[x].location can contain
		//- Location
		if sourceResourceType.Location != nil {
			for _, r := range sourceResourceType.Location {
				if r.Location.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Location.Reference)
				}
			}
		}

		//serviceProvider can contain
		//- Organization
		if sourceResourceType.ServiceProvider != nil && sourceResourceType.ServiceProvider.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.ServiceProvider.Reference)
		}

		if sourceResourceType.Diagnosis != nil {
			for _, r := range sourceResourceType.Diagnosis {
				if r.Condition.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Condition.Reference)
				}
			}
		}

		break
	case fhir401.Immunization:
		// location can contain
		//- Location
		if sourceResourceType.Location != nil && sourceResourceType.Location.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Location.Reference)
		}

		// manufacturer can contain
		//- Organization
		if sourceResourceType.Manufacturer != nil && sourceResourceType.Manufacturer.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Manufacturer.Reference)
		}

		//performer[x].actor can contain
		//- Practitioner | PractitionerRole | Organization
		if sourceResourceType.Performer != nil {
			for _, r := range sourceResourceType.Performer {
				if r.Actor.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Actor.Reference)
				}
			}
		}

		//reasonReference[x] can contain
		//- Condition | Observation | DiagnosticReport
		if sourceResourceType.ReasonReference != nil {
			for _, r := range sourceResourceType.ReasonReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//protocolApplied[x].authority can contain
		//- Organization
		if sourceResourceType.ProtocolApplied != nil {
			for _, r := range sourceResourceType.ProtocolApplied {
				if r.Authority != nil && r.Authority.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Authority.Reference)
				}
			}
		}

		break
	case fhir401.Location:
		// managingOrganization can contain
		//- Organization
		if sourceResourceType.ManagingOrganization != nil && sourceResourceType.ManagingOrganization.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.ManagingOrganization.Reference)
		}

		// partOf can contain
		//- Location
		if sourceResourceType.PartOf != nil && sourceResourceType.PartOf.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.PartOf.Reference)
		}

		break
	case fhir401.MedicationRequest:

		// requester can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- RelatedPerson
		//- Device
		if sourceResourceType.Requester != nil && sourceResourceType.Requester.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Requester.Reference)
		}

		// performer can contain
		//- Practitioner | PractitionerRole | Organization | Patient | Device | RelatedPerson | CareTeam
		if sourceResourceType.Performer != nil && sourceResourceType.Performer.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Performer.Reference)
		}

		// recorder can contain
		//- Practitioner | PractitionerRole
		if sourceResourceType.Recorder != nil && sourceResourceType.Recorder.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Recorder.Reference)
		}

		//TODO: reasonReference
		if sourceResourceType.ReasonReference != nil {
			for _, r := range sourceResourceType.ReasonReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//TODO: basedOn
		if sourceResourceType.BasedOn != nil {
			for _, r := range sourceResourceType.BasedOn {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//TODO: insurance
		if sourceResourceType.Insurance != nil {
			for _, r := range sourceResourceType.Insurance {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		// dispenseRequest.performer can contain
		//- Organization
		if sourceResourceType.DispenseRequest != nil && sourceResourceType.DispenseRequest.Performer != nil && sourceResourceType.DispenseRequest.Performer.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.DispenseRequest.Performer.Reference)
		}
		break
	case fhir401.Observation:
		//basedOn[x] can contain
		//- CarePlan | DeviceRequest | ImmunizationRecommendation | MedicationRequest | NutritionOrder | ServiceRequest
		if sourceResourceType.BasedOn != nil {
			for _, r := range sourceResourceType.BasedOn {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		// partOf[x] can contain
		//- MedicationAdministration | MedicationDispense | MedicationStatement | Procedure | Immunization | ImagingStudy
		if sourceResourceType.PartOf != nil {
			for _, r := range sourceResourceType.PartOf {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		// performer[x] can contain
		//- Practitioner | PractitionerRole | Organization | CareTeam | Patient | RelatedPerson
		if sourceResourceType.Performer != nil {
			for _, r := range sourceResourceType.Performer {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		// device can contain
		//- Device | DeviceMetric
		if sourceResourceType.Device != nil && sourceResourceType.Device.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Device.Reference)
		}

		break
	case fhir401.PractitionerRole:
		// practitioner can contain
		//- Practitioner
		if sourceResourceType.Practitioner != nil && sourceResourceType.Practitioner.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Practitioner.Reference)
		}

		//organization can contain
		//- Organization
		if sourceResourceType.Organization != nil && sourceResourceType.Organization.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Organization.Reference)
		}

		//location can contain
		//- Location
		if sourceResourceType.Location != nil {
			for _, r := range sourceResourceType.Location {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//TODO: healthcareService
		if sourceResourceType.HealthcareService != nil {
			for _, r := range sourceResourceType.HealthcareService {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//TODO: endpoint
		break
	case fhir401.ServiceRequest:
		// basedOn[x] can contain
		//- CarePlan | ServiceRequest | MedicationRequest
		if sourceResourceType.BasedOn != nil {
			for _, r := range sourceResourceType.BasedOn {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//requester can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- RelatedPerson
		//- Device
		if sourceResourceType.Requester != nil && sourceResourceType.Requester.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Requester.Reference)
		}

		//performer[x] can contain
		//- Practitioner | PractitionerRole | Organization | CareTeam | HealthcareService | Patient | Device | RelatedPerson
		if sourceResourceType.Performer != nil {
			for _, r := range sourceResourceType.Performer {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//locationReference[x] an contain
		//-Location
		if sourceResourceType.LocationReference != nil {
			for _, r := range sourceResourceType.LocationReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//reasonReference[x] can contain
		//-Condition
		//-Observation
		//-DiagnosticReport
		//-DocumentReference
		if sourceResourceType.ReasonReference != nil {
			for _, r := range sourceResourceType.ReasonReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		//insurance[x] can contain
		//- Coverage | ClaimResponse
		if sourceResourceType.Insurance != nil {
			for _, r := range sourceResourceType.Insurance {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		break
	}

	// remove all null values, remove all duplicates
	cleanResourceRefs := removeDuplicateStr(resourceRefs)
	return cleanResourceRefs
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
