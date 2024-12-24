package datatypes

import (
	"github.com/go-playground/validator/v10"
)

type Telecom struct {
	System string `json:"system,omitempty" yaml:"system,omitempty" validate:"required,oneof=phone fax email pager url sms other"`
	Value  string `json:"value,omitempty" yaml:"value,omitempty" validate:"required,min=1"`
}

func (t *Telecom) Validate() error {
	valid := validator.New()
	err := valid.Struct(t)
	if err != nil {
		return err
	}

	return nil
}
