// SPDX-FileCopyrightText: © 2026 OpenCHAMI a Series of LF Projects, LLC
// SPDX-License-Identifier: MIT

package xnames

import (
	"testing"

	"github.com/Cray-HPE/hms-xname/xnametypes"
)

func TestRouterRoundTrip(t *testing.T) {
	parent := RouterBMC{Cabinet: 4900, Chassis: 0, RouterModule: 1, RouterBMC: 0}
	router := parent.Router(0)
	const id = "x4900c0r1b0n0"
	if got := router.String(); got != id {
		t.Fatalf("xname=%q, want %q", got, id)
	}
	if err := router.Validate(); err != nil {
		t.Fatal(err)
	}
	if got := router.Type(); got != xnametypes.Router {
		t.Errorf("type=%q, want Router", got)
	}
	if got := router.Parent(); got != parent {
		t.Errorf("parent=%v, want %v", got, parent)
	}
	if got := router.ParentInterface(); got != parent {
		t.Errorf("parent interface=%v, want %v", got, parent)
	}
	if got := FromString(id); got != router {
		t.Errorf("parsed=%v, want %v", got, router)
	}
	if got := FromStringToStruct[Router](id); got == nil || *got != router {
		t.Errorf("typed parsed=%v, want %v", got, router)
	}
	for _, tt := range []struct {
		name  string
		value interface{}
	}{
		{"value", router}, {"pointer", &router},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := GetHMSType(tt.value); err != nil || got != xnametypes.Router {
				t.Errorf("type=%q error=%v, want Router", got, err)
			}
		})
	}
}

func TestRouterRejectsInvalidOrdinals(t *testing.T) {
	for _, tt := range []struct {
		name   string
		router Router
	}{
		{"negative OS", Router{Cabinet: 4900, Chassis: 0, RouterModule: 1, RouterBMC: 0, Router: -1}},
		{"invalid chassis", Router{Cabinet: 4900, Chassis: 8, RouterModule: 1, RouterBMC: 0, Router: 0}},
		{"invalid cabinet", Router{Cabinet: 10000, Chassis: 0, RouterModule: 1, RouterBMC: 0, Router: 0}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.router.Validate(); err == nil {
				t.Error("expected validation error")
			}
			if got := FromString(tt.router.String()); got != nil {
				t.Errorf("parsed invalid Router: %v", got)
			}
		})
	}
}
