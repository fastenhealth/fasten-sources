package base

import (
	"fmt"
	"github.com/fastenhealth/fasten-sources/clients/models"
	"github.com/fastenhealth/gofhir-models/fhir401"
	"strings"
	"time"
)

/*
This function used to extract "metadata" from FHIR resource, and add it to the resource which will be stored in the database.
This additional data includes:
- list of references to other resources (for creating a graph)
- "sort date" - a date when this resource event occurred
- "sort title" - a title for this resource event.
*/
func SourceClientFHIR401ExtractResourceMetadata(resourceRaw interface{}, resource *models.RawResourceFhir, internalFragmentReferenceLookup map[string]string) {
	referencedResources := []string{}
	var sortTitle *string
	var sortDate *string

	switch sourceResourceTyped := resourceRaw.(type) {

	case fhir401.AllergyIntolerance:
		if sourceResourceTyped.Code != nil && len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		}

		if sourceResourceTyped.OnsetPeriod != nil && sourceResourceTyped.OnsetPeriod.Start != nil {
			sortDate = sourceResourceTyped.OnsetPeriod.Start
		} else if sourceResourceTyped.OnsetDateTime != nil {
			sortDate = sourceResourceTyped.OnsetDateTime
		} else if len(sourceResourceTyped.Reaction) > 0 && sourceResourceTyped.Reaction[0].Onset != nil {
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
				if r.OutcomeReference != nil {
					for _, or := range r.OutcomeReference {
						if or.Reference != nil {
							referencedResources = append(referencedResources, *or.Reference)
						}
					}
				}
				if r.Reference != nil && r.Reference.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference.Reference)
				}

				if r.Detail != nil {
					if r.Detail.ReasonReference != nil {
						for _, dr := range r.Detail.ReasonReference {
							if dr.Reference != nil {
								referencedResources = append(referencedResources, *dr.Reference)
							}
						}
					}
					if r.Detail.Goal != nil {
						for _, dg := range r.Detail.Goal {
							if dg.Reference != nil {
								referencedResources = append(referencedResources, *dg.Reference)
							}
						}
					}
					if r.Detail.Location != nil && r.Detail.Location.Reference != nil {
						referencedResources = append(referencedResources, *r.Detail.Location.Reference)
					}
					if r.Detail.Performer != nil {
						for _, dp := range r.Detail.Performer {
							if dp.Reference != nil {
								referencedResources = append(referencedResources, *dp.Reference)
							}
						}
					}
					if r.Detail.ProductReference != nil && r.Detail.ProductReference.Reference != nil {
						referencedResources = append(referencedResources, *r.Detail.ProductReference.Reference)
					}
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

		if sourceResourceTyped.OnsetPeriod != nil && sourceResourceTyped.OnsetPeriod.Start != nil {
			sortDate = sourceResourceTyped.OnsetPeriod.Start
		} else if sourceResourceTyped.OnsetDateTime != nil {
			sortDate = sourceResourceTyped.OnsetDateTime
		} else if sourceResourceTyped.RecordedDate != nil {
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

		if sourceResourceTyped.Stage != nil {
			for _, s := range sourceResourceTyped.Stage {
				if s.Assessment != nil {
					for _, a := range s.Assessment {
						if a.Reference != nil {
							referencedResources = append(referencedResources, *a.Reference)
						}
					}
				}
			}
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

		if len(sourceResourceTyped.Class) > 0 && sourceResourceTyped.Class[0].Name != nil {
			sortTitle = sourceResourceTyped.Class[0].Name
		} else if len(sourceResourceTyped.Identifier) > 0 && sourceResourceTyped.Identifier[0].Value != nil {
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

		if sourceResourceTyped.EffectiveDateTime != nil {
			sortDate = sourceResourceTyped.EffectiveDateTime
		} else if sourceResourceTyped.EffectivePeriod != nil && sourceResourceTyped.EffectivePeriod.Start != nil {
			sortDate = sourceResourceTyped.EffectivePeriod.Start
		} else if sourceResourceTyped.Issued != nil {
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

		if sourceResourceTyped.Context != nil {
			if sourceResourceTyped.Context.Encounter != nil {
				for _, er := range sourceResourceTyped.Context.Encounter {
					if er.Reference != nil {
						referencedResources = append(referencedResources, *er.Reference)
					}
				}
			}
			if sourceResourceTyped.Context.SourcePatientInfo != nil {
				if sourceResourceTyped.Context.SourcePatientInfo.Reference != nil {
					referencedResources = append(referencedResources, *sourceResourceTyped.Context.SourcePatientInfo.Reference)
				}
			}
			if sourceResourceTyped.Context.Related != nil {
				for _, rr := range sourceResourceTyped.Context.Related {
					if rr.Reference != nil {
						referencedResources = append(referencedResources, *rr.Reference)
					}
				}
			}
		}
		break
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
	case fhir401.ExplanationOfBenefit:

		if sourceResourceTyped.Diagnosis != nil && len(sourceResourceTyped.Diagnosis) > 0 {
			if sourceResourceTyped.Diagnosis[0].DiagnosisCodeableConcept.Text != nil {
				sortTitle = sourceResourceTyped.Diagnosis[0].DiagnosisCodeableConcept.Text
			} else if sourceResourceTyped.Diagnosis[0].DiagnosisReference.Display != nil {
				sortTitle = sourceResourceTyped.Diagnosis[0].DiagnosisReference.Display
			}
		} else if sourceResourceTyped.Type.Text != nil {
			sortTitle = sourceResourceTyped.Type.Text
		} else if len(sourceResourceTyped.Type.Coding) > 0 && sourceResourceTyped.Type.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Type.Coding[0].Display
		}

		if sourceResourceTyped.BillablePeriod != nil && sourceResourceTyped.BillablePeriod.Start != nil {
			sortDate = sourceResourceTyped.BillablePeriod.Start
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Patient.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Patient.Reference)
		}

		if sourceResourceTyped.Enterer != nil && sourceResourceTyped.Enterer.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Enterer.Reference)
		}

		if sourceResourceTyped.Insurer.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Insurer.Reference)
		}

		if sourceResourceTyped.Provider.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Provider.Reference)
		}

		if sourceResourceTyped.Related != nil {
			for _, r := range sourceResourceTyped.Related {
				if r.Claim != nil && r.Claim.Reference != nil {
					referencedResources = append(referencedResources, *r.Claim.Reference)
				}
			}
		}

		if sourceResourceTyped.Prescription != nil && sourceResourceTyped.Prescription.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Prescription.Reference)
		}

		if sourceResourceTyped.OriginalPrescription != nil && sourceResourceTyped.OriginalPrescription.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.OriginalPrescription.Reference)
		}

		if sourceResourceTyped.Payee != nil && sourceResourceTyped.Payee.Party != nil && sourceResourceTyped.Payee.Party.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Payee.Party.Reference)
		}

		if sourceResourceTyped.Referral != nil && sourceResourceTyped.Referral.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Referral.Reference)
		}
		if sourceResourceTyped.Facility != nil && sourceResourceTyped.Facility.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Facility.Reference)
		}
		if sourceResourceTyped.Claim != nil && sourceResourceTyped.Claim.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Claim.Reference)
		}
		if sourceResourceTyped.ClaimResponse != nil && sourceResourceTyped.ClaimResponse.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.ClaimResponse.Reference)
		}

		if sourceResourceTyped.CareTeam != nil {
			for _, r := range sourceResourceTyped.CareTeam {
				if r.Provider.Reference != nil {
					referencedResources = append(referencedResources, *r.Provider.Reference)
				}
			}
		}

		if sourceResourceTyped.SupportingInfo != nil {
			for _, r := range sourceResourceTyped.SupportingInfo {
				if r.ValueReference != nil && r.ValueReference.Reference != nil {
					referencedResources = append(referencedResources, *r.ValueReference.Reference)
				}

				if r.ValueAttachment != nil && r.ValueAttachment.Url != nil {
					referencedResources = append(referencedResources, *r.ValueAttachment.Url)
				}
			}
		}

		if sourceResourceTyped.Diagnosis != nil {
			for _, r := range sourceResourceTyped.Diagnosis {
				if r.DiagnosisReference.Reference != nil {
					referencedResources = append(referencedResources, *r.DiagnosisReference.Reference)
				}
			}
		}

		if sourceResourceTyped.Procedure != nil {
			for _, r := range sourceResourceTyped.Procedure {
				if r.ProcedureReference.Reference != nil {
					referencedResources = append(referencedResources, *r.ProcedureReference.Reference)
				}
				if r.Udi != nil {
					for _, u := range r.Udi {
						if u.Reference != nil {
							referencedResources = append(referencedResources, *u.Reference)
						}
					}
				}
			}
		}

		if sourceResourceTyped.Insurance != nil {
			for _, r := range sourceResourceTyped.Insurance {
				if r.Coverage.Reference != nil {
					referencedResources = append(referencedResources, *r.Coverage.Reference)
				}
			}
		}
		if sourceResourceTyped.Accident != nil && sourceResourceTyped.Accident.LocationReference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Accident.LocationReference.Reference)
		}

		if sourceResourceTyped.Item != nil {
			for _, r := range sourceResourceTyped.Item {

				if r.LocationReference != nil {
					referencedResources = append(referencedResources, *r.LocationReference.Reference)
				}

				if r.Udi != nil {
					for _, u := range r.Udi {
						if u.Reference != nil {
							referencedResources = append(referencedResources, *u.Reference)
						}
					}
				}

				if r.Encounter != nil {
					for _, u := range r.Encounter {
						if u.Reference != nil {
							referencedResources = append(referencedResources, *u.Reference)
						}
					}
				}
				if r.Detail != nil {
					for _, u := range r.Detail {
						if u.Udi != nil {
							for _, u := range u.Udi {
								if u.Reference != nil {
									referencedResources = append(referencedResources, *u.Reference)
								}
							}
						}
						//TODO: subdetail
					}
				}
			}
		}
		if sourceResourceTyped.AddItem != nil {
			for _, r := range sourceResourceTyped.AddItem {
				if r.Provider != nil {
					for _, p := range r.Provider {
						if p.Reference != nil {
							referencedResources = append(referencedResources, *p.Reference)
						}
					}
				}
			}
		}

		if sourceResourceTyped.Form != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Form.Url)
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

		if sourceResourceTyped.StartDate != nil {
			sortDate = sourceResourceTyped.StartDate
		} else if sourceResourceTyped.StatusDate != nil {
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
		} else if len(sourceResourceTyped.OccurrenceDateTime) > 0 {
			sortDate = &sourceResourceTyped.OccurrenceDateTime
		} else if len(sourceResourceTyped.OccurrenceString) > 0 {
			sortDate = &sourceResourceTyped.OccurrenceString
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
	case fhir401.Media:

		if sourceResourceTyped.Type != nil && sourceResourceTyped.Type.Text != nil {
			sortTitle = sourceResourceTyped.Type.Text
		} else if sourceResourceTyped.Type != nil && len(sourceResourceTyped.Type.Coding) > 0 && sourceResourceTyped.Type.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Type.Coding[0].Display
		}

		if sourceResourceTyped.CreatedDateTime != nil {
			sortDate = sourceResourceTyped.CreatedDateTime
		} else if sourceResourceTyped.CreatedPeriod != nil && sourceResourceTyped.CreatedPeriod.Start != nil {
			sortDate = sourceResourceTyped.CreatedPeriod.Start
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.BasedOn != nil && len(sourceResourceTyped.BasedOn) > 0 && sourceResourceTyped.BasedOn[0].Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.BasedOn[0].Reference)
		}
		if sourceResourceTyped.PartOf != nil && len(sourceResourceTyped.PartOf) > 0 {
			for _, r := range sourceResourceTyped.PartOf {
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

		if sourceResourceTyped.Operator != nil && sourceResourceTyped.Operator.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Operator.Reference)
		}

		if sourceResourceTyped.Device != nil && sourceResourceTyped.Device.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Device.Reference)
		}

		if sourceResourceTyped.Content.Url != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Content.Url)
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

		if sourceResourceTyped.ReportedReference != nil && sourceResourceTyped.ReportedReference.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.ReportedReference.Reference)
		}

		if sourceResourceTyped.MedicationReference.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.MedicationReference.Reference)
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

		if sourceResourceTyped.EffectivePeriod != nil && sourceResourceTyped.EffectivePeriod.Start != nil {
			sortDate = sourceResourceTyped.EffectivePeriod.Start
		} else if sourceResourceTyped.EffectiveDateTime != nil {
			sortDate = sourceResourceTyped.EffectiveDateTime
		} else if sourceResourceTyped.Issued != nil {
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

		if len(sourceResourceTyped.Name) > 0 && sourceResourceTyped.Name[0].Text != nil {
			sortTitle = sourceResourceTyped.Name[0].Text
		} else if len(sourceResourceTyped.Name) > 0 && len(sourceResourceTyped.Name[0].Given) > 0 && sourceResourceTyped.Name[0].Family != nil {
			name := fmt.Sprintf("%s, %s", *sourceResourceTyped.Name[0].Family, sourceResourceTyped.Name[0].Given[0])
			sortTitle = &name
		}

		if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Photo != nil {
			for _, r := range sourceResourceTyped.Photo {
				if r.Url != nil {
					referencedResources = append(referencedResources, *r.Url)
				}
			}
		}

		if sourceResourceTyped.Contact != nil {
			for _, r := range sourceResourceTyped.Contact {
				if r.Organization != nil && r.Organization.Reference != nil {
					referencedResources = append(referencedResources, *r.Organization.Reference)
				}
			}
		}

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

		if sourceResourceTyped.Link != nil {
			for _, r := range sourceResourceTyped.Link {
				if r.Other.Reference != nil {
					referencedResources = append(referencedResources, *r.Other.Reference)
				}
			}
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

		if sourceResourceTyped.Photo != nil {
			for _, r := range sourceResourceTyped.Photo {
				if r.Url != nil {
					referencedResources = append(referencedResources, *r.Url)
				}
			}
		}

		if sourceResourceTyped.Qualification != nil {
			for _, r := range sourceResourceTyped.Qualification {
				if r.Issuer != nil && r.Issuer.Reference != nil {
					referencedResources = append(referencedResources, *r.Issuer.Reference)
				}
			}
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

		if sourceResourceTyped.HealthcareService != nil {
			for _, r := range sourceResourceTyped.HealthcareService {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		break
	case fhir401.Procedure:

		if sourceResourceTyped.Code != nil && sourceResourceTyped.Code.Text != nil {
			sortTitle = sourceResourceTyped.Code.Text
		} else if sourceResourceTyped.Code != nil && len(sourceResourceTyped.Code.Coding) > 0 && sourceResourceTyped.Code.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Code.Coding[0].Display
		}

		if sourceResourceTyped.PerformedDateTime != nil {
			sortDate = sourceResourceTyped.PerformedDateTime
		} else if sourceResourceTyped.PerformedPeriod != nil && sourceResourceTyped.PerformedPeriod.Start != nil {
			sortDate = sourceResourceTyped.PerformedPeriod.Start
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

		if sourceResourceTyped.FocalDevice != nil {
			for _, r := range sourceResourceTyped.FocalDevice {
				if r.Manipulated.Reference != nil {
					referencedResources = append(referencedResources, *r.Manipulated.Reference)
				}
			}
		}

		if sourceResourceTyped.UsedReference != nil {
			for _, r := range sourceResourceTyped.UsedReference {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		break
	case fhir401.Provenance:

		if sourceResourceTyped.Activity != nil && sourceResourceTyped.Activity.Text != nil {
			sortTitle = sourceResourceTyped.Activity.Text
		} else if sourceResourceTyped.Activity != nil && len(sourceResourceTyped.Activity.Coding) > 0 && sourceResourceTyped.Activity.Coding[0].Display != nil {
			sortTitle = sourceResourceTyped.Activity.Coding[0].Display
		}

		if sourceResourceTyped.OccurredPeriod != nil && sourceResourceTyped.OccurredPeriod.Start != nil {
			sortDate = sourceResourceTyped.OccurredPeriod.Start
		} else if sourceResourceTyped.OccurredDateTime != nil {
			sortDate = sourceResourceTyped.OccurredDateTime
		} else if sourceResourceTyped.Meta != nil && sourceResourceTyped.Meta.LastUpdated != nil {
			sortDate = sourceResourceTyped.Meta.LastUpdated
		}

		if sourceResourceTyped.Target != nil {
			for _, r := range sourceResourceTyped.Target {
				if r.Reference != nil {
					referencedResources = append(referencedResources, *r.Reference)
				}
			}
		}

		if sourceResourceTyped.Location != nil && sourceResourceTyped.Location.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Location.Reference)
		}

		if sourceResourceTyped.Agent != nil {
			for _, r := range sourceResourceTyped.Agent {
				if r.Who.Reference != nil {
					referencedResources = append(referencedResources, *r.Who.Reference)
				}
				if r.OnBehalfOf != nil && r.OnBehalfOf.Reference != nil {
					referencedResources = append(referencedResources, *r.OnBehalfOf.Reference)
				}
			}
		}

		if sourceResourceTyped.Entity != nil {
			for _, e := range sourceResourceTyped.Entity {
				if e.What.Reference != nil {
					referencedResources = append(referencedResources, *e.What.Reference)
				}
			}
		}

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

		if sourceResourceTyped.Patient.Reference != nil {
			referencedResources = append(referencedResources, *sourceResourceTyped.Patient.Reference)
		}

		if sourceResourceTyped.Photo != nil {
			for _, r := range sourceResourceTyped.Photo {
				if r.Url != nil {
					referencedResources = append(referencedResources, *r.Url)
				}
			}
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

		if sourceResourceTyped.OccurrencePeriod != nil && sourceResourceTyped.OccurrencePeriod.Start != nil {
			sortDate = sourceResourceTyped.OccurrencePeriod.Start
		} else if sourceResourceTyped.OccurrenceDateTime != nil {
			sortDate = sourceResourceTyped.OccurrenceDateTime
		} else if sourceResourceTyped.AuthoredOn != nil {
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

	//before storing resource references, we need to determine if any of them are internal bundle references, and replace them if so.
	for ndx, _ := range cleanResourceRefs {
		internalRef := cleanResourceRefs[ndx]
		cleanResourceRefs[ndx] = normalizeReferenceId(internalRef, internalFragmentReferenceLookup)
	}
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

//normalizeReference takes a reference string and returns a normalized reference string
//FHIR references come in multiple forms:
//
// [id] the logical [id] of a resource using a local reference (i.e. a relative reference).
// [type]/[id] the logical [id] of a resource of a specified type using a local reference (i.e. a relative reference), for when the reference can point to different types of resources (e.g. Observation.subject).
// [type]/[id]/_history/[version] TU the logical [id] of a resource of a specified type using a local reference (i.e. a relative reference), for when the reference can point to different types of resources and a specific version is requested. Note that server implementations MAY return an error when using this syntax if resource versions are not supported. For more information, see References and Versions.
// [url] where the [url] is an absolute URL - a reference to a resource by its absolute location, or by its canonical URL
// [url]|[version] TU where the search element is a canonical reference, the [url] is an absolute URL, and a specific version or partial version is desired. For more information, see References and Versions.
//
// see https://build.fhir.org/search.html#reference
// see https://build.fhir.org/references.html
func normalizeReferenceId(originalReference string, internalFragmentReferenceLookup map[string]string) string {
	if strings.HasPrefix(originalReference, "urn:uuid:") {
		if relativeReference, relativeReferenceOk := internalFragmentReferenceLookup[originalReference]; relativeReferenceOk {
			//replace internal reference with relative reference
			return relativeReference
		}
	}

	// handle absolute urls
	if strings.HasPrefix(originalReference, "http://") || strings.HasPrefix(originalReference, "https://") {
		//do nothing, absolute urls must be handled as-is
		return originalReference

	} else {
		if strings.Contains(originalReference, "|") {
			//split on | and drop second part
			originalReference = strings.SplitN(originalReference, "|", 2)[0]
		}

		if strings.Contains(originalReference, "/_history/") {
			//split on _history and drop second part
			return strings.SplitN(originalReference, "/_history/", 2)[0]
		}
	}

	//fallback, do nothing
	return originalReference
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
