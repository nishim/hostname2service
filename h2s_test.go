package h2s

import "testing"

func TestUndefined(t *testing.T) {
	e := "undefined"
	a := Detect("qwerty")
	if e != a {
		t.Errorf("expected: %v, actual: %v", e, a)
	}
}

func TestDefined(t *testing.T) {
	cases := []struct {
		service string
		name    string
	}{
		{service: "AWS EC2", name: "ec2-0-0-0-0.ap-northeast-1.compute.amazonaws.com"},
		{service: "AWS CloudFront", name: "server-0-0-0-0.nrt00.r.cloudfront.net"},
	}

	for _, c := range cases {
		s := Detect(c.name)
		if s != c.service {
			t.Errorf("expected: %v, actual: %v", c.service, s)
		}
	}
}
