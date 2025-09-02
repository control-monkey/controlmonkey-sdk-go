package cross_models

import "github.com/control-monkey/controlmonkey-sdk-go/controlmonkey/util/jsonutil"

type RunTrigger struct {
	Patterns        []*string `json:"patterns,omitempty"`
	ExcludePatterns []*string `json:"excludePatterns,omitempty"`

	forceSendFields []string
	nullFields      []string
}

func (o RunTrigger) MarshalJSON() ([]byte, error) {
	type noMethod RunTrigger
	raw := noMethod(o)
	return jsonutil.MarshalJSON(raw, o.forceSendFields, o.nullFields)
}

func (o *RunTrigger) SetPatterns(v []*string) *RunTrigger {
	if o.Patterns = v; o.Patterns == nil {
		o.nullFields = append(o.nullFields, "Patterns")
	}
	return o
}

func (o *RunTrigger) SetExcludePatterns(v []*string) *RunTrigger {
	if o.ExcludePatterns = v; o.ExcludePatterns == nil {
		o.nullFields = append(o.nullFields, "ExcludePatterns")
	}
	return o
}
