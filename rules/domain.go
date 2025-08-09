package rules

import (
	"strings"

	"golang.org/x/net/idna"
)

type Domain struct {
	name   string
	values map[string]any
}

func (r *Domain) GetName() string {
	return r.name
}

func (r *Domain) GetValues() map[string]any {
	return r.values
}

func (r *Domain) Validate(value any) bool {
	r.name = "domain"

	if value == nil {
		return false
	}

	switch value.(type) {
	case string:
	default:
		return false
	}

	v := value.(string)

	d, err := idna.ToASCII(v)
	if err != nil {
		return false
	}

	l := len(d)

	if l == 0 || len(d) > 254 {
		return false
	}

	parts := strings.Split(d, ".")

	for _, part := range parts {
		lp := len(part)

		if lp < 2 || lp > 63 {
			return false
		}

		idn := strings.HasPrefix(part, "xn--")
		if !idn && part[2] == '-' && part[3] == '-' {
			return false
		}

		i := 0
		if idn {
			i = 4
		}

		for ; i < lp; i++ {
			if (i == 0 || i == lp-1) && part[i] == '-' {
				return false
			}

			if !(part[i] >= 'a' && part[i] <= 'z' || part[i] >= '0' && part[i] <= '9' || part[i] == '-') {
				return false
			}
		}
	}

	return true
}
