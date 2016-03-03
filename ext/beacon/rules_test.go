package beacon

import (
	"testing"

	"github.com/samalba/dockerclient"
)

func TestRuleMatchImage(t *testing.T) {
	image := "redis"

	info := &dockerclient.ContainerInfo{
		Config: &dockerclient.ContainerConfig{
			Image: image,
		},
	}

	b := &Beacon{}

	m := b.isImageMatch(image, info)
	if !m {
		t.Fatalf("expected image match for %s", image)
	}
}

func TestRuleNoMatchImage(t *testing.T) {
	image := "redis"

	info := &dockerclient.ContainerInfo{
		Config: &dockerclient.ContainerConfig{
			Image: "nginx",
		},
	}

	b := &Beacon{}

	m := b.isImageMatch(image, info)
	if m {
		t.Fatalf("expected no match for %s", image)
	}
}

func TestRuleMatchName(t *testing.T) {
	name := "test-container"

	info := &dockerclient.ContainerInfo{
		Name: name,
	}

	b := &Beacon{}

	m := b.isNameMatch(name, info)
	if !m {
		t.Fatalf("expected name match for %s", name)
	}
}

func TestRuleMatchNameWildcard(t *testing.T) {
	name := "test-container"

	info := &dockerclient.ContainerInfo{
		Name: name,
	}

	b := &Beacon{}

	m := b.isNameMatch("test-.*", info)
	if !m {
		t.Fatalf("expected name match for %s", name)
	}
}

func TestRuleNoMatchName(t *testing.T) {
	name := "test-container"

	info := &dockerclient.ContainerInfo{
		Name: name,
	}

	b := &Beacon{}

	m := b.isNameMatch("foo", info)
	if m {
		t.Fatalf("expected no match for %s", name)
	}
}

func TestRuleMatchLabel(t *testing.T) {
	label := "foo=bar"

	info := &dockerclient.ContainerInfo{
		Config: &dockerclient.ContainerConfig{
			Labels: map[string]string{
				"foo": "bar",
			},
		},
	}

	b := &Beacon{}

	m := b.isLabelMatch(label, info)
	if !m {
		t.Fatalf("expected label match for %s", label)
	}
}

func TestRuleNoMatchLabel(t *testing.T) {
	label := "foo=bar"

	info := &dockerclient.ContainerInfo{
		Config: &dockerclient.ContainerConfig{
			Labels: map[string]string{
				"foo": "bar",
			},
		},
	}

	b := &Beacon{}

	m := b.isLabelMatch("baz=foo", info)
	if m {
		t.Fatalf("expected no match for %s", label)
	}
}

func TestRuleMatchLabelWildcard(t *testing.T) {
	label := "foo=.*"

	info := &dockerclient.ContainerInfo{
		Config: &dockerclient.ContainerConfig{
			Labels: map[string]string{
				"foo": "bar",
			},
		},
	}

	b := &Beacon{}

	m := b.isLabelMatch(label, info)
	if !m {
		t.Fatalf("expected label match for %s", label)
	}
}
