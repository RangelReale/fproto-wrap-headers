package ptypes_headers

import (
	"github.com/RangelReale/fproto-wrap-headers/gowrap/gwproto"
)

func Import(source *gwproto.Headers) map[string][]string {
	if source == nil || source.Headers == nil {
		return nil
	}

	ret := make(map[string][]string)
	for hn, hv := range source.Headers {
		ret[hn] = hv.Values
	}
	return ret
}

func Export(source map[string][]string) *gwproto.Headers {
	if source == nil || len(source) == 0 {
		return nil
	}

	ret := &gwproto.Headers{
		Headers: make(map[string]*gwproto.Headers_Values),
	}
	for hn, hv := range source {
		ret.Headers[hn] = &gwproto.Headers_Values{
			Values: hv,
		}
	}

	return ret
}
