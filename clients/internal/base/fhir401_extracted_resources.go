package base

import (
	"github.com/fastenhealth/gofhir-models/fhir401"
)

func (c *SourceClientFHIR401) ExtractResourceReference(resourceRaw interface{}) []string {
	resourceRefs := []string{}

	switch sourceResourceType := resourceRaw.(type) {

	case fhir401.AllergyIntolerance:

		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
		}

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
	case fhir401.Binary:
		if sourceResourceType.SecurityContext != nil && sourceResourceType.SecurityContext.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.SecurityContext.Reference)
		}
		break
	case fhir401.CarePlan:

		if sourceResourceType.BasedOn != nil {
			for _, r := range sourceResourceType.BasedOn {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.Replaces != nil {
			for _, r := range sourceResourceType.Replaces {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.PartOf != nil {
			for _, r := range sourceResourceType.PartOf {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}

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

		if sourceResourceType.Addresses != nil {
			for _, r := range sourceResourceType.Addresses {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.SupportingInfo != nil {
			for _, r := range sourceResourceType.SupportingInfo {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.Goal != nil {
			for _, r := range sourceResourceType.Goal {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.Activity != nil {
			for _, r := range sourceResourceType.Activity {
				if r.Reference != nil && r.Reference.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference.Reference)
				}
			}
		}
		break
	case fhir401.CareTeam:

		if sourceResourceType.Subject != nil && sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}

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

		if sourceResourceType.ReasonReference != nil {
			for _, r := range sourceResourceType.ReasonReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
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
		if sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
		}

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

		if sourceResourceType.Evidence != nil {
			for _, r := range sourceResourceType.Evidence {
				if r.Detail != nil {
					for _, d := range r.Detail {
						if d.Reference != nil {
							resourceRefs = append(resourceRefs, *d.Reference)
						}
					}
				}
			}
		}

		break
	case fhir401.Coverage:
		if sourceResourceType.PolicyHolder != nil && sourceResourceType.PolicyHolder.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.PolicyHolder.Reference)
		}
		if sourceResourceType.Subscriber != nil && sourceResourceType.Subscriber.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subscriber.Reference)
		}
		if sourceResourceType.Beneficiary.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Beneficiary.Reference)
		}
		if sourceResourceType.Payor != nil {
			for _, r := range sourceResourceType.Payor {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.Contract != nil {
			for _, r := range sourceResourceType.Contract {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		break
	case fhir401.Device:
		if sourceResourceType.Definition != nil && sourceResourceType.Definition.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Definition.Reference)
		}
		if sourceResourceType.Patient != nil && sourceResourceType.Patient.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Patient.Reference)
		}
		if sourceResourceType.Owner != nil && sourceResourceType.Owner.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Owner.Reference)
		}
		if sourceResourceType.Location != nil && sourceResourceType.Location.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Location.Reference)
		}
		if sourceResourceType.Parent != nil && sourceResourceType.Parent.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Parent.Reference)
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

		if sourceResourceType.Subject != nil && sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
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

		if sourceResourceType.Media != nil {
			for _, r := range sourceResourceType.Media {
				if r.Link.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Link.Reference)
				}
			}
		}

		if sourceResourceType.PresentedForm != nil {
			for _, r := range sourceResourceType.PresentedForm {
				if r.Url != nil && len(*r.Url) > 0 {
					resourceRefs = append(resourceRefs, *r.Url)
				}
			}
		}

		break
	case fhir401.DocumentReference:
		if sourceResourceType.Subject != nil && sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}

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
		if sourceResourceType.Content != nil {
			for _, r := range sourceResourceType.Content {
				if r.Attachment.Url != nil {
					resourceRefs = append(resourceRefs, *r.Attachment.Url)
				}
			}
		}
	case fhir401.Encounter:

		if sourceResourceType.Subject != nil && sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.EpisodeOfCare != nil {
			for _, r := range sourceResourceType.EpisodeOfCare {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

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

		if sourceResourceType.Appointment != nil {
			for _, r := range sourceResourceType.Appointment {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
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

		if sourceResourceType.Diagnosis != nil {
			for _, r := range sourceResourceType.Diagnosis {
				if r.Condition.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Condition.Reference)
				}
			}
		}

		if sourceResourceType.Account != nil {
			for _, r := range sourceResourceType.Account {
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
		if sourceResourceType.PartOf != nil && sourceResourceType.PartOf.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.PartOf.Reference)
		}

		break
	case fhir401.Goal:
		if sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.ExpressedBy != nil && sourceResourceType.ExpressedBy.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.ExpressedBy.Reference)
		}
		if sourceResourceType.Addresses != nil {
			for _, r := range sourceResourceType.Addresses {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.OutcomeReference != nil {
			for _, r := range sourceResourceType.OutcomeReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		break
	case fhir401.Immunization:
		if sourceResourceType.Patient.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Patient.Reference)
		}
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
		}

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
		if sourceResourceType.Education != nil {
			for _, r := range sourceResourceType.Education {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		if sourceResourceType.Reaction != nil {
			for _, r := range sourceResourceType.Reaction {
				if r.Detail != nil && r.Detail.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Detail.Reference)
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
		if sourceResourceType.Endpoint != nil {
			for _, r := range sourceResourceType.Endpoint {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

		break
	case fhir401.Medication:
		if sourceResourceType.Manufacturer != nil && sourceResourceType.Manufacturer.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Manufacturer.Reference)
		}
		break
	case fhir401.MedicationRequest:
		if sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
		}
		if sourceResourceType.SupportingInformation != nil {
			for _, r := range sourceResourceType.SupportingInformation {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}

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

		if sourceResourceType.ReasonReference != nil {
			for _, r := range sourceResourceType.ReasonReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.BasedOn != nil {
			for _, r := range sourceResourceType.BasedOn {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
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

		if sourceResourceType.PriorPrescription != nil && sourceResourceType.PriorPrescription.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.PriorPrescription.Reference)
		}

		if sourceResourceType.DetectedIssue != nil {
			for _, r := range sourceResourceType.DetectedIssue {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.EventHistory != nil {
			for _, r := range sourceResourceType.EventHistory {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
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

		if sourceResourceType.Subject != nil && sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.Focus != nil {
			for _, r := range sourceResourceType.Focus {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
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

		if sourceResourceType.Specimen != nil && sourceResourceType.Specimen.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Specimen.Reference)
		}
		// device can contain
		//- Device | DeviceMetric
		if sourceResourceType.Device != nil && sourceResourceType.Device.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Device.Reference)
		}
		if sourceResourceType.HasMember != nil {
			for _, r := range sourceResourceType.HasMember {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.DerivedFrom != nil {
			for _, r := range sourceResourceType.DerivedFrom {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		break
	case fhir401.Organization:
		if sourceResourceType.PartOf != nil && sourceResourceType.PartOf.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.PartOf.Reference)
		}
		if sourceResourceType.Endpoint != nil {
			for _, r := range sourceResourceType.Endpoint {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		break
	case fhir401.Patient:

		if sourceResourceType.GeneralPractitioner != nil {
			for _, r := range sourceResourceType.GeneralPractitioner {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.ManagingOrganization != nil && sourceResourceType.ManagingOrganization.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.ManagingOrganization.Reference)
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
	case fhir401.Procedure:
		if sourceResourceType.BasedOn != nil {
			for _, r := range sourceResourceType.BasedOn {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.PartOf != nil {
			for _, r := range sourceResourceType.PartOf {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
		}
		if sourceResourceType.Recorder != nil && sourceResourceType.Recorder.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Recorder.Reference)
		}
		if sourceResourceType.Asserter != nil && sourceResourceType.Asserter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Asserter.Reference)
		}
		if sourceResourceType.Performer != nil {
			for _, r := range sourceResourceType.Performer {
				if r.Actor.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Actor.Reference)
				}
			}
		}
		if sourceResourceType.Location != nil && sourceResourceType.Location.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Location.Reference)
		}
		if sourceResourceType.ReasonReference != nil {
			for _, r := range sourceResourceType.ReasonReference {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.Report != nil {
			for _, r := range sourceResourceType.Report {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.ComplicationDetail != nil {
			for _, r := range sourceResourceType.ComplicationDetail {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
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
		if sourceResourceType.Replaces != nil {
			for _, r := range sourceResourceType.Replaces {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.Encounter != nil && sourceResourceType.Encounter.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Encounter.Reference)
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

		if sourceResourceType.SupportingInfo != nil {
			for _, r := range sourceResourceType.SupportingInfo {
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

		if sourceResourceType.RelevantHistory != nil {
			for _, r := range sourceResourceType.RelevantHistory {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		break
	case fhir401.Specimen:

		if sourceResourceType.Subject != nil && sourceResourceType.Subject.Reference != nil {
			resourceRefs = append(resourceRefs, *sourceResourceType.Subject.Reference)
		}
		if sourceResourceType.Parent != nil {
			for _, r := range sourceResourceType.Parent {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.Request != nil {
			for _, r := range sourceResourceType.Request {
				if r.Reference != nil {
					resourceRefs = append(resourceRefs, *r.Reference)
				}
			}
		}
		if sourceResourceType.Request != nil {
			for _, r := range sourceResourceType.Request {
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
