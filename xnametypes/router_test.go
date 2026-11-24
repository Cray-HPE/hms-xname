// SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
// SPDX-License-Identifier: MIT

package xnametypes

import "testing"

func TestRouterRecognition(t *testing.T) {
	for _, tt := range []struct {
		name  string
		input string
		want  string
	}{
		{"NVLink OS", "x4900c0r1b0n0", "x4900c0r1b0n0"},
		{"minimum ordinals", "x0c0r0b0n0", "x0c0r0b0n0"},
		{"router bounds", "x9999c7r123b45n67", "x9999c7r123b45n67"},
		{"normalization", " X04900C0R01B00N00 ", "x4900c0r1b0n0"},
		{"missing OS index", "x4900c0r1b0n", ""},
		{"missing BMC", "x4900c0r1n0", ""},
		{"negative OS index", "x4900c0r1b0n-1", ""},
		{"chassis out of range", "x4900c8r1b0n0", ""},
		{"cabinet out of range", "x10000c0r1b0n0", ""},
		{"extra suffix", "x4900c0r1b0n0garbage", ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			normalized := VerifyNormalizeCompID(tt.input)
			if normalized != tt.want {
				t.Fatalf("normalized=%q, want %q", normalized, tt.want)
			}
			if tt.want == "" {
				return
			}
			if got := GetHMSTypeString(normalized); got != "Router" {
				t.Errorf("type=%q, want Router", got)
			}
			if !IsHMSCompIDValid(normalized) {
				t.Error("normalized Router is invalid")
			}
		})
	}
}

func TestRouterTypeMetadata(t *testing.T) {
	router := Router
	if got := ToHMSType("rOuTeR"); got != router {
		t.Errorf("type=%q, want Router", got)
	}
	if got := VerifyNormalizeType("router"); got != "Router" {
		t.Errorf("normalized type=%q", got)
	}
	format, argc, err := GetHMSTypeFormatString(router)
	if err != nil || format != "x%dc%dr%db%dn%d" || argc != 5 {
		t.Fatalf("format=%q argc=%d error=%v", format, argc, err)
	}
	entry := GetHMSCompRecognitionTable()["router"]
	if entry.ParentType != RouterBMC {
		t.Errorf("parent type=%q, want RouterBMC", entry.ParentType)
	}
	if got := GetHMSCompParent("x4900c0r1b0n0"); got != "x4900c0r1b0" {
		t.Errorf("parent=%q", got)
	}
	if IsHMSTypeController(router) || IsHMSTypeContainer(router) {
		t.Error("Router is an OS component, not a controller or chassis container")
	}
	for id, want := range map[string]HMSType{"x4900c0r1b0": RouterBMC, "x4900c0r1": RouterModule, "x4900c0s1b0n0": Node} {
		if got := GetHMSType(id); got != want {
			t.Errorf("%s type=%q, want %q", id, got, want)
		}
	}
}

func TestRouterInTypeAndIDLists(t *testing.T) {
	found := false
	for _, name := range GetHMSTypeList() {
		if name == "Router" {
			found = true
		}
	}
	if !found {
		t.Error("Router missing from type list")
	}
	valid, invalid := ValidateCompIDs([]string{"x4900c0r1b0n0", "x4900c0r1b0n"}, false)
	if len(valid) != 1 || valid[0] != "x4900c0r1b0n0" || len(invalid) != 1 || invalid[0] != "x4900c0r1b0n" {
		t.Errorf("valid=%v invalid=%v", valid, invalid)
	}
}
