package version

import (
	"testing"
)

func TestIsVersion(t *testing.T) {
	a := IsVersion("2.1.3")
	b := IsVersion("2.3-beta")
	c := IsVersion(">= 2.1.3")
	d := IsVersion("~> 2.1.3")
	if !a || !b || c || d {
		t.Error("TestIsVersion Fail:", a, b, c, d)
	}
}

func TestIsVersionConstraint(t *testing.T) {
	a := IsVersionConstraint("2.1.3")
	b := IsVersionConstraint("2.3-beta")
	c := IsVersionConstraint(">= 2.1.3")
	d := IsVersionConstraint("~> 2.1.3")
	f := IsVersionConstraint("abc")
	e := IsVersionConstraint("")
	if !a || !b || !c || !d || f || e {
		t.Error("TestIsVersionConstraint Fial:", a, b, c, d, f, e)
	}
}

func TestCompareVersion(t *testing.T) {
	a := CompareVersion("2.1.0", "2.1.1")
	b := CompareVersion("2.1-beta", "2.1.0")
	if a > -1 || b > -1 {
		t.Error(a, b)
	}
}

func TestMaxVersion(t *testing.T) {
	v := []string{"2.1.1", "2.1.2", "2.1-beta", "2.2.0"}
	a, errA := MaxVersion("", v...)
	b, errB := MaxVersion("~> 2.1.0", v...)
	c, errC := MaxVersion("< 2.1.0", v...)
	if errA != nil {
		t.Error(errA.Error())
	} else if errB != nil {
		t.Error(errB.Error())
	} else if errC != nil {
		t.Error(errC.Error())
	} else if a != "2.2.0" || b != "2.1.2" || c != "2.1-beta" {
		t.Error(a, b, c)
	}
}

func TestMatchVersionConstraint(t *testing.T) {
	a := MatchVersionConstraint("", "2.1.1")
	b := MatchVersionConstraint("~> 2.1.0", "2.1.2")
	c := MatchVersionConstraint("~> 2.1.0", "2.2.2")
	d := MatchVersionConstraint("> 2.1.0", "2.2.2")
	e := MatchVersionConstraint("< 2.1.0", "2.2.2")
	if !a || !b || c || !d || e {
		t.Error(a, b, c, d, e)
	}
}
