package services

import "slices"

type FlagRouter struct {
	r Registry
}

func NewFlagRouter(r Registry) *FlagRouter {
	return &FlagRouter{
		r,
	}
}

func (m *FlagRouter) IsFeatureEnabled(flag []string) bool {
	return slices.ContainsFunc(m.r.Config().Applictaion.Features, func(feature string) bool {
		var feat bool
		for _, f := range flag {
			if f == feature {
				feat = true
			} else {
				feat = false
			}
		}

		return feat
	})
}

func (m *FlagRouter) SetFlag(flag string, value bool) {
	if value {
		m.r.Config().Applictaion.Features = append(m.r.Config().Applictaion.Features, flag)
	} else {
		m.r.Config().Applictaion.Features = remove(m.r.Config().Applictaion.Features, "two")
	}
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
