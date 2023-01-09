package base

import (
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/gofhir-models/fhir401"
	"time"
)

/*
This function used to extract "metadata" from FHIR resource, and add it to the resource which will be stored in the database.
This additional data includes:
- list of references to other resources (for creating a graph)
- "sort date" - a date when this resource event occurred
- "sort title" - a title for this resource event.
*/
func SourceClientFHIR401ExtractResourceMetadata(resourceRaw interface{}, resource *models.RawResourceFhir) {
	referencedResources := []string{}
	var sortTitle *string
	var sortDate *string

	switch sourceResourceTyped := resourceRaw.(type) {

	case fhir401.AllergyIntolerance:
		if sourceResourceTyped.Code != nil && len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		}

		if len(sourceResourceTyped.Reaction) > 0 && sourceResourceTyped.Reaction[0].Onset != nil {
			sortDate = sourceResourceTyped.Reaction[0].Onset
		} else if sourceResourceTyped.RecordedDate != nil {
			sortDate = sourceResourceTyped.RecordedDate
		}

		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}

		// recorder can contain
		//- Practitioner
		//- PractitionerRole
		//- Patient
		//- RelatedPerson
		if sourceResourceTyped.Recorder != nil && sourceResourceTyped.Recorder.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Recorder.Reference)
		}

		// asserter can contain
		//- Practitioner
		//- PractitionerRole
		//- Patient
		//- RelatedPerson
		if sourceResourceTyped.Asserter != nil && sourceResourceTyped.Asserter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Asserter.Reference)
		}
	case fhir401.Binary:
		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.SecurityContext != nil && sourceResourceTyped.SecurityContext.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.SecurityContext.Reference)
		}
		break
	case fhir401.CarePlan:

		if sourceResourceTyped.Title != nil {
			sortTitle = sourceResourceTyped.Title
		} else if len(sourceResourceTyped.Addresses) > 0 && sourceResourceTyped.Addresses[0].Display != nil {
			sortTitle = sourceResourceTyped.Addresses[0].Display
		}

		if sourceResourceTyped.Period != nil && sourceResourceTyped.Period.Start != nil {
			sortDate = sourceResourceTyped.Period.Start
		} else if sourceResourceTyped.Created != nil {
			sortDate = sourceResourceTyped.Created
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.BasedOn != nil {
			for _, r := range sourceResourceTyped.BasedOn {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Replaces != nil {
			for _, r := range sourceResourceTyped.Replaces {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.PartOf != nil {
			for _, r := range sourceResourceTyped.PartOf {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}

		// encounter can contain
		//- Encounter
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}

		//author can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- CareTeam
		//- RelatedPerson
		if sourceResourceTyped.Author != nil && sourceResourceTyped.Author.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Author.Reference)
		}

		//contributor can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- CareTeam
		//- RelatedPerson
		if sourceResourceTyped.Contributor != nil {
			for _, r := range sourceResourceTyped.Contributor {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}

		}

		//careTeam can contain
		//- CareTeam
		if sourceResourceTyped.CareTeam != nil {
			for _, r := range sourceResourceTyped.CareTeam {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Addresses != nil {
			for _, r := range sourceResourceTyped.Addresses {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.SupportingInfo != nil {
			for _, r := range sourceResourceTyped.SupportingInfo {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Goal != nil {
			for _, r := range sourceResourceTyped.Goal {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Activity != nil {
			for _, r := range sourceResourceTyped.Activity {
				if r.Reference != nil && r.Reference.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference.Reference)
				}
			}
		}
		break
	case fhir401.CareTeam:

		if sourceResourceTyped.Name != nil {
			sortTitle = sourceResourceTyped.Name
		}

		if sourceResourceTyped.Period != nil && sourceResourceTyped.Period.Start != nil {
			sortDate = sourceResourceTyped.Period.Start
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Subject != nil && sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}

		// encounter can contain
		//- Encounter
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
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
		if sourceResourceTyped.Participant != nil {
			for _, r := range sourceResourceTyped.Participant {
				if r.Member != nil && r.Member.Reference != nil {
					referencedResources = append(referencedResources, *r.Member.Reference)
				}
				if r.OnBehalfOf != nil && r.OnBehalfOf.Reference != nil {
					referencedResources = append(referencedResources, *r.OnBehalfOf.Reference)
				}
			}
		}

		if sourceResourceTyped.ReasonReference != nil {
			for _, r := range sourceResourceTyped.ReasonReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//managingOrganization
		//- Organization
		if sourceResourceTyped.ManagingOrganization != nil {
			for _, r := range sourceResourceTyped.ManagingOrganization {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		break
	case fhir401.Condition:

		if sourceResourceTyped.Code != nil && len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		} else if sourceResourceTyped.Code != nil && sourceResourceTyped.Code.Text != nil {
			sortTitle = sourceResourceTyped.Code.Text
		} else if sourceResourceTyped.Code != nil && len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Code != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Code
		}

		if sourceResourceTyped.RecordedDate != nil {
			sortDate = sourceResourceTyped.RecordedDate
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}

		// recorder can contain
		//- Practitioner
		//- PractitionerRole
		//- Patient
		//- RelatedPerson
		if sourceResourceTyped.Recorder != nil && sourceResourceTyped.Recorder.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Recorder.Reference)
		}

		// asserter can contain
		//- Practitioner
		//- PractitionerRole
		//- Patient
		//- RelatedPerson
		if sourceResourceTyped.Asserter != nil && sourceResourceTyped.Asserter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Asserter.Reference)
		}

		if sourceResourceTyped.Evidence != nil {
			for _, r := range sourceResourceTyped.Evidence {
				if r.Detail != nil {
					for _, d := range r.Detail {
						if d.Reference != nil {
							referencedResources = append(referencedResources, *d.Reference)
						}
					}
				}
			}
		}

		break
	case fhir401.Coverage:

		if len(sourceResourceTyped.Identifier) > 0 && sourceResourceTyped.Identifier[0].Value != nil {
			sortTitle = sourceResourceTyped.Identifier[0].Value
		}

		if sourceResourceTyped.Period != nil && sourceResourceTyped.Period.Start != nil {
			sortDate = sourceResourceTyped.Period.Start
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.PolicyHolder != nil && sourceResourceTyped.PolicyHolder.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.PolicyHolder.Reference)
		}
		if sourceResourceTyped.Subscriber != nil && sourceResourceTyped.Subscriber.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subscriber.Reference)
		}
		if sourceResourceTyped.Beneficiary.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Beneficiary.Reference)
		}
		if sourceResourceTyped.Payor != nil {
			for _, r := range sourceResourceTyped.Payor {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Contract != nil {
			for _, r := range sourceResourceTyped.Contract {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		break
	case fhir401.Device:

		if len(sourceResourceTyped.DeviceName) > 0 {
			sortTitle = &sourceResourceTyped.DeviceName[0].Name
		}

		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Definition != nil && sourceResourceTyped.Definition.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Definition.Reference)
		}
		if sourceResourceTyped.Patient != nil && sourceResourceTyped.Patient.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Patient.Reference)
		}
		if sourceResourceTyped.Owner != nil && sourceResourceTyped.Owner.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Owner.Reference)
		}
		if sourceResourceTyped.Location != nil && sourceResourceTyped.Location.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Location.Reference)
		}
		if sourceResourceTyped.Parent != nil && sourceResourceTyped.Parent.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Parent.Reference)
		}
		break
	case fhir401.DiagnosticReport:

		if sourceResourceTyped.Code.Text != nil {
			sortTitle = sourceResourceTyped.Code.Text
		} else if len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		}

		if sourceResourceTyped.Issued != nil {
			sortDate = sourceResourceTyped.Issued
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		//basedOn[x] can contain
		//- CarePlan
		//- ImmunizationRecommendation
		//- MedicationRequest
		//- NutritionOrder
		//- ServiceRequest
		if sourceResourceTyped.BasedOn != nil {
			for _, r := range sourceResourceTyped.BasedOn {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Subject != nil && sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}

		// performer[x] can contain
		//- Practitioner
		//- PractitionerRole
		//- Organization
		//- CareTeam
		if sourceResourceTyped.Performer != nil {
			for _, r := range sourceResourceTyped.Performer {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.ResultsInterpreter != nil {
			for _, r := range sourceResourceTyped.ResultsInterpreter {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Specimen != nil {
			for _, r := range sourceResourceTyped.Specimen {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Result != nil {
			for _, r := range sourceResourceTyped.Result {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.ImagingStudy != nil {
			for _, r := range sourceResourceTyped.ImagingStudy {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Media != nil {
			for _, r := range sourceResourceTyped.Media {
				if r.Link.Reference != nil {
					referencedResources = append(referencedResources, *r.Link.Reference)
				}
			}
		}

		if sourceResourceTyped.PresentedForm != nil {
			for _, r := range sourceResourceTyped.PresentedForm {
				if r.Url != nil && len(*r.Url) > 0 {
					referencedResources = append(referencedResources, *r.Url)
				}
			}
		}

		break
	case fhir401.DocumentReference:

		if len(sourceResourceTyped.Category) > 0 && sourceResourceTyped.Category[0].Text != nil {
			sortTitle = sourceResourceTyped.Category[0].Text
		} else if len(sourceResourceTyped.Category) > 0 && len(sourceResourceTyped.Category[0].Coding) > 0 && sourceResourceTyped.Category[0].Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Category[0].Coding[0].Display
		} else if sourceResourceTyped.Description != nil {
			sortTitle = sourceResourceTyped.Description
		}

		if sourceResourceTyped.Date != nil {
			sortDate = sourceResourceTyped.Date
		} else if sourceResourceTyped.Context != nil && sourceResourceTyped.Context.Period != nil && sourceResourceTyped.Context.Period.Start != nil {
			sortDate = sourceResourceTyped.Context.Period.Start
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Subject != nil && sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}

		//author[x] can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- CareTeam
		//- Device
		if sourceResourceTyped.Author != nil {
			for _, r := range sourceResourceTyped.Author {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//authenticator can contain
		//- Practitioner
		//- Organization
		//- PractitionerRole
		if sourceResourceTyped.Authenticator != nil && sourceResourceTyped.Authenticator.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Authenticator.Reference)
		}

		// custodian can contain
		//- Organization
		if sourceResourceTyped.Custodian != nil && sourceResourceTyped.Custodian.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Custodian.Reference)
		}

		// relatesTo.target
		//- DocumentReference
		if sourceResourceTyped.RelatesTo != nil {
			for _, r := range sourceResourceTyped.RelatesTo {
				if r.Target.Reference != nil {
					referencedResources = append(referencedResources, *r.Target.Reference)
				}
			}
		}

		//content.attachment can contain
		//- Attachment
		if sourceResourceTyped.Content != nil {
			for _, r := range sourceResourceTyped.Content {
				if r.Attachment.Url != nil {
					referencedResources = append(referencedResources, *r.Attachment.Url)
				}
			}
		}
	case fhir401.Encounter:

		if len(sourceResourceTyped.Type) > 0 && sourceResourceTyped.Type[0].Text != nil {
			sortTitle = sourceResourceTyped.Type[0].Text
		} else if len(sourceResourceTyped.Type) > 0 && len(sourceResourceTyped.Type[0].Coding) > 0 && sourceResourceTyped.Type[0].Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Type[0].Coding[0].Display
		}

		if sourceResourceTyped.Period != nil && sourceResourceTyped.Period.Start != nil {
			sortDate = sourceResourceTyped.Period.Start
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Subject != nil && sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.EpisodeOfCare != nil {
			for _, r := range sourceResourceTyped.EpisodeOfCare {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		// basedOn[x] can contain
		//- ServiceRequest
		if sourceResourceTyped.BasedOn != nil {
			for _, r := range sourceResourceTyped.BasedOn {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//participant[x].individual can contain
		//- Practitioner
		//- PractitionerRole
		//- RelatedPerson
		if sourceResourceTyped.Participant != nil {
			for _, r := range sourceResourceTyped.Participant {
				if r.Individual != nil && r.Individual.Reference != nil {
					referencedResources = append(referencedResources, *r.Individual.Reference)
				}
			}
		}

		if sourceResourceTyped.Appointment != nil {
			for _, r := range sourceResourceTyped.Appointment {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//reasonReference[x] can contain
		//- Condition
		//- Procedure
		//- Observation
		//- ImmunizationRecommendation
		if sourceResourceTyped.ReasonReference != nil {
			for _, r := range sourceResourceTyped.ReasonReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Diagnosis != nil {
			for _, r := range sourceResourceTyped.Diagnosis {
				if r.Condition.Reference != nil {
					referencedResources = append(referencedResources, *r.Condition.Reference)
				}
			}
		}

		if sourceResourceTyped.Account != nil {
			for _, r := range sourceResourceTyped.Account {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//hospitalization.origin can contain
		//- Location
		//- Organization
		if sourceResourceTyped.Hospitalization != nil && sourceResourceTyped.Hospitalization.Origin != nil && sourceResourceTyped.Hospitalization.Origin.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Hospitalization.Origin.Reference)
		}

		//hospitalization.destination can contain
		//- Location
		//- Organization
		//referencedResources.push(resourceRaw.hospitalization?.destination?.reference)
		if sourceResourceTyped.Hospitalization != nil && sourceResourceTyped.Hospitalization.Destination != nil && sourceResourceTyped.Hospitalization.Destination.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Hospitalization.Destination.Reference)
		}

		//location[x].location can contain
		//- Location
		if sourceResourceTyped.Location != nil {
			for _, r := range sourceResourceTyped.Location {
				if r.Location.Reference != nil {
					referencedResources = append(referencedResources, *r.Location.Reference)
				}
			}
		}

		//serviceProvider can contain
		//- Organization
		if sourceResourceTyped.ServiceProvider != nil && sourceResourceTyped.ServiceProvider.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.ServiceProvider.Reference)
		}
		if sourceResourceTyped.PartOf != nil && sourceResourceTyped.PartOf.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.PartOf.Reference)
		}

		break
	case fhir401.Goal:

		if len(sourceResourceTyped.Note) > 0 {
			sortTitle = &sourceResourceTyped.Note[0].Text
		} else if sourceResourceTyped.Description.Text != nil {
			sortTitle = sourceResourceTyped.Description.Text
		} else if len(sourceResourceTyped.Description.Coding) > 0 && sourceResourceTyped.Description.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Description.Coding[0].Display
		}

		if sourceResourceTyped.StatusDate != nil {
			sortDate = sourceResourceTyped.StatusDate
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.ExpressedBy != nil && sourceResourceTyped.ExpressedBy.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.ExpressedBy.Reference)
		}
		if sourceResourceTyped.Addresses != nil {
			for _, r := range sourceResourceTyped.Addresses {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.OutcomeReference != nil {
			for _, r := range sourceResourceTyped.OutcomeReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		break
	case fhir401.Immunization:

		if sourceResourceTyped.VaccineCode.Text != nil {
			sortTitle = sourceResourceTyped.VaccineCode.Text
		} else if len(sourceResourceTyped.VaccineCode.Coding) > 0 && sourceResourceTyped.VaccineCode.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.VaccineCode.Coding[0].Display
		}

		if sourceResourceTyped.Recorded != nil {
			sortDate = sourceResourceTyped.Recorded
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Patient.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Patient.Reference)
		}
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}

		// location can contain
		//- Location
		if sourceResourceTyped.Location != nil && sourceResourceTyped.Location.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Location.Reference)
		}

		// manufacturer can contain
		//- Organization
		if sourceResourceTyped.Manufacturer != nil && sourceResourceTyped.Manufacturer.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Manufacturer.Reference)
		}

		//performer[x].actor can contain
		//- Practitioner | PractitionerRole | Organization
		if sourceResourceTyped.Performer != nil {
			for _, r := range sourceResourceTyped.Performer {
				if r.Actor.Reference != nil {
					referencedResources = append(referencedResources, *r.Actor.Reference)
				}
			}
		}

		//reasonReference[x] can contain
		//- Condition | Observation | DiagnosticReport
		if sourceResourceTyped.ReasonReference != nil {
			for _, r := range sourceResourceTyped.ReasonReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Education != nil {
			for _, r := range sourceResourceTyped.Education {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Reaction != nil {
			for _, r := range sourceResourceTyped.Reaction {
				if r.Detail != nil && r.Detail.Reference != nil {
					referencedResources = append(referencedResources, *r.Detail.Reference)
				}
			}
		}

		//protocolApplied[x].authority can contain
		//- Organization
		if sourceResourceTyped.ProtocolApplied != nil {
			for _, r := range sourceResourceTyped.ProtocolApplied {
				if r.Authority != nil && r.Authority.Reference != nil {
					referencedResources = append(referencedResources, *r.Authority.Reference)
				}
			}
		}

		break
	case fhir401.Location:

		if sourceResourceTyped.Name != nil {
			sortTitle = sourceResourceTyped.Name
		}

		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		// managingOrganization can contain
		//- Organization
		if sourceResourceTyped.ManagingOrganization != nil && sourceResourceTyped.ManagingOrganization.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.ManagingOrganization.Reference)
		}

		// partOf can contain
		//- Location
		if sourceResourceTyped.PartOf != nil && sourceResourceTyped.PartOf.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.PartOf.Reference)
		}
		if sourceResourceTyped.Endpoint != nil {
			for _, r := range sourceResourceTyped.Endpoint {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		break
	case fhir401.Medication:

		if sourceResourceTyped.Code != nil && sourceResourceTyped.Code.Text != nil {
			sortTitle = sourceResourceTyped.Code.Text
		} else if sourceResourceTyped.Code != nil && len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		}

		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Manufacturer != nil && sourceResourceTyped.Manufacturer.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Manufacturer.Reference)
		}
		break
	case fhir401.MedicationRequest:

		if len(sourceResourceTyped.Identifier) > 0 && sourceResourceTyped.Identifier[0].Value != nil {
			sortTitle = sourceResourceTyped.Identifier[0].Value
		}

		if sourceResourceTyped.AuthoredOn != nil {
			sortDate = sourceResourceTyped.AuthoredOn
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}
		if sourceResourceTyped.SupportingInformation != nil {
			for _, r := range sourceResourceTyped.SupportingInformation {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
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
		if sourceResourceTyped.Requester != nil && sourceResourceTyped.Requester.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Requester.Reference)
		}

		// performer can contain
		//- Practitioner | PractitionerRole | Organization | Patient | Device | RelatedPerson | CareTeam
		if sourceResourceTyped.Performer != nil && sourceResourceTyped.Performer.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Performer.Reference)
		}

		// recorder can contain
		//- Practitioner | PractitionerRole
		if sourceResourceTyped.Recorder != nil && sourceResourceTyped.Recorder.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Recorder.Reference)
		}

		if sourceResourceTyped.ReasonReference != nil {
			for _, r := range sourceResourceTyped.ReasonReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.BasedOn != nil {
			for _, r := range sourceResourceTyped.BasedOn {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Insurance != nil {
			for _, r := range sourceResourceTyped.Insurance {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		// dispenseRequest.performer can contain
		//- Organization
		if sourceResourceTyped.DispenseRequest != nil && sourceResourceTyped.DispenseRequest.Performer != nil && sourceResourceTyped.DispenseRequest.Performer.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.DispenseRequest.Performer.Reference)
		}

		if sourceResourceTyped.PriorPrescription != nil && sourceResourceTyped.PriorPrescription.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.PriorPrescription.Reference)
		}

		if sourceResourceTyped.DetectedIssue != nil {
			for _, r := range sourceResourceTyped.DetectedIssue {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.EventHistory != nil {
			for _, r := range sourceResourceTyped.EventHistory {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		break
	case fhir401.Observation:

		if sourceResourceTyped.Code.Text != nil {
			sortTitle = sourceResourceTyped.Code.Text
		} else if len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		}

		if sourceResourceTyped.Issued != nil {
			sortDate = sourceResourceTyped.Issued
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		//basedOn[x] can contain
		//- CarePlan | DeviceRequest | ImmunizationRecommendation | MedicationRequest | NutritionOrder | ServiceRequest
		if sourceResourceTyped.BasedOn != nil {
			for _, r := range sourceResourceTyped.BasedOn {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		// partOf[x] can contain
		//- MedicationAdministration | MedicationDispense | MedicationStatement | Procedure | Immunization | ImagingStudy
		if sourceResourceTyped.PartOf != nil {
			for _, r := range sourceResourceTyped.PartOf {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Subject != nil && sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.Focus != nil {
			for _, r := range sourceResourceTyped.Focus {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}

		// performer[x] can contain
		//- Practitioner | PractitionerRole | Organization | CareTeam | Patient | RelatedPerson
		if sourceResourceTyped.Performer != nil {
			for _, r := range sourceResourceTyped.Performer {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Specimen != nil && sourceResourceTyped.Specimen.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Specimen.Reference)
		}
		// device can contain
		//- Device | DeviceMetric
		if sourceResourceTyped.Device != nil && sourceResourceTyped.Device.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Device.Reference)
		}
		if sourceResourceTyped.HasMember != nil {
			for _, r := range sourceResourceTyped.HasMember {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.DerivedFrom != nil {
			for _, r := range sourceResourceTyped.DerivedFrom {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		break
	case fhir401.Organization:

		if sourceResourceTyped.Name != nil {
			sortTitle = sourceResourceTyped.Name
		}

		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.PartOf != nil && sourceResourceTyped.PartOf.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.PartOf.Reference)
		}
		if sourceResourceTyped.Endpoint != nil {
			for _, r := range sourceResourceTyped.Endpoint {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		break
	case fhir401.Patient:

		if sourceResourceTyped.GeneralPractitioner != nil {
			for _, r := range sourceResourceTyped.GeneralPractitioner {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.ManagingOrganization != nil && sourceResourceTyped.ManagingOrganization.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.ManagingOrganization.Reference)
		}
		break
	case fhir401.Practitioner:

		if len(sourceResourceTyped.Name) > 0 && sourceResourceTyped.Name[0].Text != nil {
			sortTitle = sourceResourceTyped.Name[0].Text
		} else if len(sourceResourceTyped.Name) > 0 && len(sourceResourceTyped.Name[0].Given) > 0 && sourceResourceTyped.Name[0].Family != nil {
			name := fmt.Sprintf("%s, %s", *sourceResourceTyped.Name[0].Family, sourceResourceTyped.Name[0].Given[0])
			sortTitle = &name
		}

		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		break
	case fhir401.PractitionerRole:

		if sourceResourceTyped.Practitioner != nil && sourceResourceTyped.Practitioner.Display != nil {
			sortTitle = sourceResourceTyped.Practitioner.Display
		} else if len(sourceResourceTyped.Code) > 0 && sourceResourceTyped.Code[0].Text != nil {
			sortTitle = sourceResourceTyped.Code[0].Text
		}

		if sourceResourceTyped.Period != nil && sourceResourceTyped.Period.Start != nil {
			sortDate = sourceResourceTyped.Period.Start
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		// practitioner can contain
		//- Practitioner
		if sourceResourceTyped.Practitioner != nil && sourceResourceTyped.Practitioner.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Practitioner.Reference)
		}

		//organization can contain
		//- Organization
		if sourceResourceTyped.Organization != nil && sourceResourceTyped.Organization.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Organization.Reference)
		}

		//location can contain
		//- Location
		if sourceResourceTyped.Location != nil {
			for _, r := range sourceResourceTyped.Location {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//TODO: healthcareService
		if sourceResourceTyped.HealthcareService != nil {
			for _, r := range sourceResourceTyped.HealthcareService {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//TODO: endpoint
		break
	case fhir401.Procedure:

		if sourceResourceTyped.Code != nil && sourceResourceTyped.Code.Text != nil {
			sortTitle = sourceResourceTyped.Code.Text
		} else if sourceResourceTyped.Code != nil && len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		}

		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.BasedOn != nil {
			for _, r := range sourceResourceTyped.BasedOn {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.PartOf != nil {
			for _, r := range sourceResourceTyped.PartOf {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}
		if sourceResourceTyped.Recorder != nil && sourceResourceTyped.Recorder.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Recorder.Reference)
		}
		if sourceResourceTyped.Asserter != nil && sourceResourceTyped.Asserter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Asserter.Reference)
		}
		if sourceResourceTyped.Performer != nil {
			for _, r := range sourceResourceTyped.Performer {
				if r.Actor.Reference != nil {
					referencedResources = append(referencedResources, *r.Actor.Reference)
				}
			}
		}
		if sourceResourceTyped.Location != nil && sourceResourceTyped.Location.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Location.Reference)
		}
		if sourceResourceTyped.ReasonReference != nil {
			for _, r := range sourceResourceTyped.ReasonReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Report != nil {
			for _, r := range sourceResourceTyped.Report {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.ComplicationDetail != nil {
			for _, r := range sourceResourceTyped.ComplicationDetail {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		break
	case fhir401.Provenance:
		break
	case fhir401.RelatedPerson:
		if len(sourceResourceTyped.Name) > 0 && sourceResourceTyped.Name[0].Text != nil {
			sortTitle = sourceResourceTyped.Name[0].Text
		} else if len(sourceResourceTyped.Name) > 0 && len(sourceResourceTyped.Name[0].Given) > 0 && sourceResourceTyped.Name[0].Family != nil {
			name := fmt.Sprintf("%s, %s", *sourceResourceTyped.Name[0].Family, sourceResourceTyped.Name[0].Given[0])
			sortTitle = &name
		}

		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}
		break
	case fhir401.ServiceRequest:

		if sourceResourceTyped.Code != nil && sourceResourceTyped.Code.Text != nil {
			sortTitle = sourceResourceTyped.Code.Text
		} else if sourceResourceTyped.Code != nil && len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		} else if len(sourceResourceTyped.OrderDetail) > 0 && sourceResourceTyped.OrderDetail[0].Text != nil {
			sortTitle = sourceResourceTyped.OrderDetail[0].Text
		}

		if sourceResourceTyped.AuthoredOn != nil {
			sortDate = sourceResourceTyped.AuthoredOn
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		// basedOn[x] can contain
		//- CarePlan | ServiceRequest | MedicationRequest
		if sourceResourceTyped.BasedOn != nil {
			for _, r := range sourceResourceTyped.BasedOn {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Replaces != nil {
			for _, r := range sourceResourceTyped.Replaces {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.Encounter != nil && sourceResourceTyped.Encounter.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Encounter.Reference)
		}
		//requester can contain
		//- Practitioner
		//- Organization
		//- Patient
		//- PractitionerRole
		//- RelatedPerson
		//- Device
		if sourceResourceTyped.Requester != nil && sourceResourceTyped.Requester.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Requester.Reference)
		}

		//performer[x] can contain
		//- Practitioner | PractitionerRole | Organization | CareTeam | HealthcareService | Patient | Device | RelatedPerson
		if sourceResourceTyped.Performer != nil {
			for _, r := range sourceResourceTyped.Performer {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//locationReference[x] an contain
		//-Location
		if sourceResourceTyped.LocationReference != nil {
			for _, r := range sourceResourceTyped.LocationReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//reasonReference[x] can contain
		//-Condition
		//-Observation
		//-DiagnosticReport
		//-DocumentReference
		if sourceResourceTyped.ReasonReference != nil {
			for _, r := range sourceResourceTyped.ReasonReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		//insurance[x] can contain
		//- Coverage | ClaimResponse
		if sourceResourceTyped.Insurance != nil {
			for _, r := range sourceResourceTyped.Insurance {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.SupportingInfo != nil {
			for _, r := range sourceResourceTyped.SupportingInfo {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Specimen != nil {
			for _, r := range sourceResourceTyped.Specimen {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.RelevantHistory != nil {
			for _, r := range sourceResourceTyped.RelevantHistory {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		break
	case fhir401.Specimen:

		if sourceResourceTyped.ReceivedTime != nil {
			sortDate = sourceResourceTyped.ReceivedTime
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Subject != nil && sourceResourceTyped.Subject.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Subject.Reference)
		}
		if sourceResourceTyped.Parent != nil {
			for _, r := range sourceResourceTyped.Parent {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Request != nil {
			for _, r := range sourceResourceTyped.Request {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		if sourceResourceTyped.Request != nil {
			for _, r := range sourceResourceTyped.Request {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}
		break
	}

	// remove all null values, remove all duplicates
	cleanResourceRefs := removeDuplicateStr(referencedResources)
	resource.ReferencedResources = cleanResourceRefs

	if sortTitle != nil {
		resource.SortTitle = sortTitle
	}
	if sortDate != nil {
		sortDateTime, err := time.Parse(time.RFC3339, *sortDate)
		if err == nil {
			resource.SortDate = &sortDateTime
		}
	}

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
