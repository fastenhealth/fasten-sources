// Copyright (C) Fasten Health, Inc. - All Rights Reserved.

package pkg

type PlatformType string

const (
	PlatformTypeManual PlatformType = "manual"
	PlatformTypeFasten PlatformType = "fasten"
	// platform
	PlatformTypeAdvancedmd       PlatformType = "advancedmd"
	PlatformTypeAetna            PlatformType = "aetna"
	PlatformTypeAllscripts       PlatformType = "allscripts"
	PlatformTypeAthena           PlatformType = "athena"
	PlatformTypeBcbsal           PlatformType = "bcbsal"
	PlatformTypeAnthem           PlatformType = "anthem"
	PlatformTypeCareevolution    PlatformType = "careevolution"
	PlatformTypeCerner           PlatformType = "cerner"
	PlatformTypeCigna            PlatformType = "cigna"
	PlatformTypeEclinicalworks   PlatformType = "eclinicalworks"
	PlatformTypeEdifecs          PlatformType = "edifecs"
	PlatformTypeEpic             PlatformType = "epic"
	PlatformTypeEpicLegacy       PlatformType = "epic-legacy"
	PlatformTypeHumana           PlatformType = "humana"
	PlatformTypeKaiser           PlatformType = "kaiser"
	PlatformTypeMedicare         PlatformType = "medicare"
	PlatformTypeMeditech         PlatformType = "meditech"
	PlatformTypeNextgen          PlatformType = "nextgen"
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

		PlatformTypeAdvancedmd,
		PlatformTypeAetna,
		PlatformTypeAllscripts,
		PlatformTypeAthena,
		PlatformTypeBcbsal,
		PlatformTypeAnthem,
		PlatformTypeCareevolution,
		PlatformTypeCerner,
		PlatformTypeCigna,
		PlatformTypeEclinicalworks,
		PlatformTypeEdifecs,
		PlatformTypeEpic,
		PlatformTypeEpicLegacy,
		PlatformTypeHumana,
		PlatformTypeKaiser,
		PlatformTypeMedicare,
		PlatformTypeMeditech,
		PlatformTypeNextgen,
		PlatformTypeUnitedhealthcare,
		PlatformTypeVahealth,

		PlatformTypeHealthit,
		PlatformTypeLogica,
	}
}
