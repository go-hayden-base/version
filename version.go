package version

import ver "github.com/hashicorp/go-version"

// IsVersion return v is a version string or not
func IsVersion(v string) bool {
	if _, e := ver.NewVersion(v); e == nil {
		return true
	}
	return false
}

// IsVersionConstraint return c is a version constraint or not
func IsVersionConstraint(c string) bool {
	if _, e := ver.NewConstraint(c); e == nil {
		return true
	}
	return false
}

// CompareVersion compare version a and b.
// If a greater then b, return 1
// If a less then b, return -1
// Else return 0
func CompareVersion(a string, b string) int {
	aVerA, eA := ver.NewVersion(a)
	aVerB, eB := ver.NewVersion(b)
	if eA != nil && eB != nil {
		return 0
	} else if eA != nil {
		return -1
	} else if eB != nil {
		return 1
	}
	return aVerA.Compare(aVerB)
}

// MaxVersion return max version in v array
// If c is not empty and is a valid version constraint,
// the returned max version must match c constraint
func MaxVersion(c string, v ...string) (string, error) {
	if len(v) == 0 {
		return "", nil
	}

	var aConstraint ver.Constraints
	var err error
	if c != "" {
		aConstraint, err = ver.NewConstraint(c)
		if err != nil {
			return "", err
		}
	}

	var compares []string
	if aConstraint != nil {
		compares = make([]string, 0, len(v))
		for _, version := range v {
			aVer, e := ver.NewVersion(version)
			if e != nil {
				continue
			}
			if aConstraint.Check(aVer) {
				compares = append(compares, version)
			}
		}
	} else {
		compares = v
	}

	l := len(compares)
	if l < 1 {
		return "", nil
	}

	max := compares[0]
	for _, version := range compares {
		if aNextVersion, err := ver.NewVersion(version); err != nil {
			continue
		} else if aPreVersion, err := ver.NewVersion(max); err != nil || aNextVersion.GreaterThan(aPreVersion) {
			max = version
			continue
		}
	}
	return max, nil
}

// MatchVersionConstraint return v is match constraint c or not
func MatchVersionConstraint(c string, v string) bool {
	aConstraint, err := ver.NewConstraint(c)
	if err != nil {
		return true
	}
	aVersion, err := ver.NewVersion(v)
	if err != nil {
		return false
	}
	return aConstraint.Check(aVersion)
}
