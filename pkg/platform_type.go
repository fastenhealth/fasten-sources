// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package pkg

type PlatformType string

const (
	PlatformTypeManual PlatformType = "manual"
	PlatformTypeFasten PlatformType = "fasten"

	// TEFCA direct & TEFCA facilitated FHIR platforms
	PlatformTypeTEFCA     PlatformType = "tefca"      //tefca is special because it requires IAL2 verification
	PlatformTypeTEFCAEpic PlatformType = "tefca-epic" //tefca-epic is a facilitated fhir platform

	// platform
	PlatformTypeAdvancedmd     PlatformType = "advancedmd"
	PlatformTypeAetna          PlatformType = "aetna"
	PlatformTypeAllscripts     PlatformType = "allscripts"
	PlatformTypeAnthem         PlatformType = "anthem"
	PlatformTypeAthena         PlatformType = "athena"
	PlatformTypeBcbsal         PlatformType = "bcbsal"
	PlatformTypeCareevolution  PlatformType = "careevolution"
	PlatformTypeCerner         PlatformType = "cerner"
	PlatformTypeCHBase         PlatformType = "chbase"
	PlatformTypeCigna          PlatformType = "cigna"
	PlatformTypeEclinicalworks PlatformType = "eclinicalworks"
	PlatformTypeEdifecs        PlatformType = "edifecs"
	PlatformTypeEpic           PlatformType = "epic"

	// deprecated - use PlatformTypeEpic instead
	PlatformTypeEpicLegacy          PlatformType = "epic-legacy"
	PlatformTypeFlatiron            PlatformType = "flatiron"
	PlatformTypeDrChrono            PlatformType = "drchrono"
	PlatformTypeDynamicHealthIT     PlatformType = "dynamichealthit"
	PlatformTypeHumana              PlatformType = "humana"
	PlatformTypeKaiser              PlatformType = "kaiser"
	PlatformTypeMaximeyes           PlatformType = "maximeyes"
	PlatformTypeMedhost             PlatformType = "medhost"
	PlatformTypeMedicare            PlatformType = "medicare"
	PlatformTypeMeditech            PlatformType = "meditech"
	PlatformTypeMeldrx              PlatformType = "meldrx"
	PlatformTypeNHS                 PlatformType = "nhs"
	PlatformTypeNetsmart            PlatformType = "netsmart"
	PlatformTypeNextgen             PlatformType = "nextgen"
	PlatformTypeOneMedical          PlatformType = "onemedical"
	PlatformTypePracticeFusion      PlatformType = "practicefusion"
	PlatformTypeQualifactsCareLogic PlatformType = "qualifacts-carelogic"
	PlatformTypeQualifactsCredible  PlatformType = "qualifacts-credible"
	PlatformTypeQualifactsInSync    PlatformType = "qualifacts-insync"
	PlatformTypeQuestDiagnostics    PlatformType = "questdiagnostics"

	PlatformTypeUnitedhealthcare PlatformType = "unitedhealthcare"
	PlatformTypeVahealth         PlatformType = "vahealth"

	// sandbox only
	PlatformTypeHealthit PlatformType = "healthit"
	PlatformTypeLogica   PlatformType = "logica"
)

func GetPlatformTypes() []PlatformType {
	return []PlatformType{
		PlatformTypeManual,
		PlatformTypeFasten,
		PlatformTypeTEFCA,

		PlatformTypeAdvancedmd,
		PlatformTypeAetna,
		PlatformTypeAllscripts,
		PlatformTypeAnthem,
		PlatformTypeAthena,
		PlatformTypeBcbsal,
		PlatformTypeCareevolution,
		PlatformTypeCerner,
		PlatformTypeCHBase,
		PlatformTypeCigna,
		PlatformTypeEclinicalworks,
		PlatformTypeEdifecs,
		PlatformTypeEpic,
		PlatformTypeEpicLegacy,
		PlatformTypeFlatiron,
		PlatformTypeDrChrono,
		PlatformTypeDynamicHealthIT,
		PlatformTypeHumana,
		PlatformTypeKaiser,
		PlatformTypeMaximeyes,
		PlatformTypeMedhost,
		PlatformTypeMedicare,
		PlatformTypeMeditech,
		PlatformTypeMeldrx,
		PlatformTypeNHS,
		PlatformTypeNetsmart,
		PlatformTypeNextgen,
		PlatformTypeOneMedical,
		PlatformTypePracticeFusion,
		PlatformTypeQualifactsCareLogic,
		PlatformTypeQualifactsCredible,
		PlatformTypeQualifactsInSync,
		PlatformTypeQuestDiagnostics,
		PlatformTypeUnitedhealthcare,
		PlatformTypeVahealth,

		PlatformTypeHealthit,
		PlatformTypeLogica,
	}
}
