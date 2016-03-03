package beacon

import (
	"regexp"
	"strings"

	"github.com/samalba/dockerclient"
)

func (b *Beacon) ruleMatch(info *dockerclient.ContainerInfo) bool {
	// iterate through rules and check type / match
	isMatch := false

	for _, rule := range b.cfg.Rules {
		switch rule.Type {
		case "label":
			m := b.isLabelMatch(rule.Regex, info)
			if m {
				isMatch = m
				break
			}
		case "name":
			m := b.isNameMatch(rule.Regex, info)
			if m {
				isMatch = m
				break
			}
		case "image":
			m := b.isImageMatch(rule.Regex, info)
			if m {
				isMatch = m
				break
			}
		default:
			log().Errorf("unknown rule type: %s", rule.Type)
		}
	}

	return isMatch
}

func (b *Beacon) isLabelMatch(rule string, info *dockerclient.ContainerInfo) bool {
	key := rule
	val := ""

	parts := strings.Split(rule, "=")
	if len(parts) > 1 {
		key = parts[0]
		val = parts[1]
	}

	for k, v := range info.Config.Labels {
		if k == key {
			r := regexp.MustCompile(val)
			m := r.MatchString(v)
			return m
		}
	}

	return false
}

func (b *Beacon) isNameMatch(rule string, info *dockerclient.ContainerInfo) bool {
	r := regexp.MustCompile(rule)

	m := r.MatchString(info.Name)

	return m
}

func (b *Beacon) isImageMatch(rule string, info *dockerclient.ContainerInfo) bool {
	image := info.Config.Image

	r := regexp.MustCompile(rule)

	m := r.MatchString(image)

	return m
}
