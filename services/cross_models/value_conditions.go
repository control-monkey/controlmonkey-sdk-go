package cross_models

import (
	"encoding/json"

	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey"
	"github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"
	"github.com/control-monkey/controlmonkey-sdk-go/services/commons"
)

type Condition struct {
	Operator *string   `json:"operator,omitempty"` //commons.VariableConditionOperatorTypes
	Value    *any      `json:"value,omitempty"`
	Values   []*string // logical field to store Value if Value is a Slice

	forceSendFields []string
	nullFields      []string
}

// region Value Conditions

func (o *Condition) UnmarshalJSON(data []byte) error {
	type C Condition
	if err := json.Unmarshal(data, (*C)(o)); err != nil {
		return err
	}

	m := make(map[string]interface{})

	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	if *o.Operator == commons.In {
		values := make([]*string, 0)
		for _, element := range m["value"].([]interface{}) {
			values = append(values, controlmonkey.String(element.(string)))
		}
		o.Values = values
	} else {
		var val = m["value"]
		o.Value = &val
	}

	return nil
}

func (o Condition) MarshalJSON() ([]byte, error) {
	type noMethod Condition
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *Condition) SetOperator(v *string) *Condition {
	if o.Operator = v; o.Operator == nil {
		o.nullFields = append(o.nullFields, "Operator")
	}
	return o
}

func (o *Condition) SetValue(v *any) *Condition {
	if o.Value = v; o.Value == nil {
		o.nullFields = append(o.nullFields, "Value")
	}
	return o
}

//endregion
