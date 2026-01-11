package tui

import (
	"strings"
)

func (m *model) getRows() int {
	return (len(m.Relays) + 1) / 2
}

func getStringToSort(desc string) string {
	// gets whats inside ()
	start := strings.LastIndex(desc, "(")
	end := strings.LastIndex(desc, ")")

	if start != -1 && end != -1 && end > start {
		return strings.ToLower(desc[start+1 : end])
	}
	return strings.ToLower(desc)
}

func (m *model) getSelectedIps() []string {
	var ips []string
	for index := range m.RelaysChecked {
		pop := m.Relays[index]

		for _, relay := range pop.Relays {
			if relay.Ipv4 != "" {
				ips = append(ips, relay.Ipv4)
			}
		}
	}
	return ips
}

func (m *model) getUnSelectedIps() []string {
	var ips []string
	for i, pop := range m.Relays {
		_, checked := m.RelaysChecked[i]
		if !checked {
			for _, relay := range pop.Relays {
				if relay.Ipv4 != "" {
					ips = append(ips, relay.Ipv4)
				}
			}
		}
	}
	return ips
}
