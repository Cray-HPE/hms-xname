//
//  MIT License
//
//  (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP
//
//  Permission is hereby granted, free of charge, to any person obtaining a
//  copy of this software and associated documentation files (the "Software"),
//  to deal in the Software without restriction, including without limitation
//  the rights to use, copy, modify, merge, publish, distribute, sublicense,
//  and/or sell copies of the Software, and to permit persons to whom the
//  Software is furnished to do so, subject to the following conditions:
//
//  The above copyright notice and this permission notice shall be included
//  in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
//  THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
//  OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
//  ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//  OTHER DEALINGS IN THE SOFTWARE.
//
package xname

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/Cray-HPE/hms-xname/xnametypes"
)

func TestExample(t *testing.T) {
	n := Node{
		Cabinet:       1000,
		Chassis:       1,
		ComputeModule: 7,
		NodeBMC:       1,
		Node:          0,
	}

	nodeXname := n.String()
	expectedNodeXname := "x1000c1s7b1n0"
	if expectedNodeXname != nodeXname {
		t.Fatalf("Unexpected node xname (%s), expected (%s)", nodeXname, expectedNodeXname)
	}

	nodeBMCXname := n.Parent().String()
	expectedNodeBMCXname := "x1000c1s7b1"
	if expectedNodeBMCXname != nodeBMCXname {
		t.Fatalf("Unexpected node bmc xname (%s), expected (%s)", nodeBMCXname, expectedNodeXname)
	}

	computeModuleXname := n.Parent().Parent().String()
	expectedComputeModuleXname := "x1000c1s7"
	if expectedComputeModuleXname != computeModuleXname {
		t.Fatalf("Unexpected compute module xname (%s), expected (%s)", computeModuleXname, expectedComputeModuleXname)
	}

	chassisXname := n.Parent().Parent().Parent().String()
	expectedChassisXname := "x1000c1"
	if expectedChassisXname != chassisXname {
		t.Fatalf("Unexpected chassis xname (%s), expected (%s)", chassisXname, expectedChassisXname)
	}

	cabinetXname := n.Parent().Parent().Parent().Parent().String()
	expectedCabinetXname := "x1000"
	if expectedCabinetXname != cabinetXname {
		t.Fatalf("Unexpected cabinet xname (%s), expected (%s)", cabinetXname, expectedCabinetXname)
	}

	systemXname := n.Parent().Parent().Parent().Parent().Parent().String()
	expectedSystemXname := "s0"
	if expectedSystemXname != systemXname {
		t.Fatalf("Unexpected system xname (%s), expected (%s)", systemXname, expectedSystemXname)
	}

	// n = Cabinet{Cabinet: 1000}.Chassis(1).NodeBMC(7, 1).Node(0)
	// t.Log("Node:", n)
	expectedNodeXname = "x1000c1s7b1n0"
	n = Cabinet{Cabinet: 1000}.Chassis(1).ComputeModule(7).NodeBMC(1).Node(0)
	t.Log("Node:", n)
	if expectedNodeXname != n.String() {
		t.Fatalf("Unexpected node xname (%s), expected (%s)", n.String(), expectedNodeXname)
	}

	n = System{}.
		Cabinet(1000).
		Chassis(1).
		ComputeModule(7).
		NodeBMC(1).
		Node(0)
	t.Log("Node:", n)
	if expectedNodeXname != n.String() {
		t.Fatalf("Unexpected node xname (%s), expected (%s)", n.String(), expectedNodeXname)
	}

	// n = System{}.
	// 	Cabinet(1000).
	// 	Chassis(1).
	// 	NodeBMC(7, 1).
	// 	Node(0)
	// t.Log("Node:", n)

	n = Node{
		Cabinet:       1000,
		Chassis:       1,
		ComputeModule: 7,
		NodeBMC:       1,
		Node:          0,
	}

	hmsType, err := GetHMSType(n)
	if err != nil {
		t.Log("GetHMSType error:", err)
		t.FailNow()
		return
	}
	t.Log("HMS Type:", hmsType)

	formatStr, numArgs, err := xnametypes.GetHMSTypeFormatString(hmsType)
	if err != nil {
		t.Log("GetHMSTypeFormatString error:", err)
		t.FailNow()
		return
	}
	t.Log("Format String args:", numArgs)
	t.Log("Format String:", formatStr)

	cduSwitch := System{}.CDU(0).CDUMgmtSwitch(1)
	t.Log("CDU Switch:", cduSwitch)

	ms := MgmtSwitch{
		Cabinet:    1,  // X: 0-999
		Chassis:    0,  // C: 0-7
		MgmtSwitch: 32, // W: 1-48
	}
	t.Log("MgmtSwitch:", ms)

}

func TestRegex(t *testing.T) {
	xname := "x1c2s3b4n5"
	hmsType := xnametypes.GetHMSType(xname)
	t.Log("HMS Type:", hmsType)

	re, err := xnametypes.GetHMSTypeRegex(hmsType)
	if err != nil {
		t.Fatal("GetHMSTypeRegex error", err)
		return
	}

	_, argCount, err := xnametypes.GetHMSTypeFormatString(hmsType)
	if err != nil {
		t.Fatal("GetHMSTypeFormatString error", err)
		return
	}
	t.Log("Format String Args:", argCount)

	matchesRaw := re.FindStringSubmatch(xname)
	t.Log("Matches Raw", matchesRaw)

	if (argCount + 1) != len(matchesRaw) {
		t.Fatal("Unexpected number of matches found:", len(matchesRaw), "expected:", argCount)
		return
	}

	matches := []int{}
	for _, matchRaw := range matchesRaw[1:] {
		// If we have gotten to this point these matches should be integers
		match, err := strconv.Atoi(matchRaw)
		if err != nil {
			t.Fatal("unable to convert match to integer:", matchRaw, "error:", err)
			return
		}

		matches = append(matches, match)
	}

	t.Log("Matches", matches)

	node := Node{
		Cabinet:       matches[0],
		Chassis:       matches[1],
		ComputeModule: matches[2],
		NodeBMC:       matches[3],
		Node:          matches[4],
	}

	t.Log("Node", node)

}

func TestToFromXnames(t *testing.T) {
	// Note, not all of the xnames in the following tests are valid. Each ordinal is incremented by 1 to verify that each ordinal is being properly
	// handled and not getting lost or switched around.
	tests := []struct {
		xname             string
		hmsType           xnametypes.HMSType
		expectedComponent interface{}
	}{
		{
			"s0",
			xnametypes.System,
			System{},
		}, {
			"d0",
			xnametypes.CDU,
			CDU{
				CDU: 0,
			},
		}, {
			"d0w1", xnametypes.CDUMgmtSwitch,
			CDUMgmtSwitch{
				CDU:           0,
				CDUMgmtSwitch: 1,
			},
		}, {
			"x1",
			xnametypes.Cabinet,
			Cabinet{
				Cabinet: 1,
			},
		}, {
			"x1c2",
			xnametypes.Chassis,
			Chassis{
				Cabinet: 1,
				Chassis: 2,
			},
		}, {
			"x1c2b0",
			xnametypes.ChassisBMC,
			ChassisBMC{
				Cabinet:    1,
				Chassis:    2,
				ChassisBMC: 0,
			},
			// }, { // TODO This causes a panic
			// 	"x1c2b3",
			// 	xnametypes.ChassisBMC,
			// 	ChassisBMC{
			// 		Cabinet: 1,
			// 		Chassis: 2,
			// 		BMC:     3,
			// 	},
		}, {
			"x1c2h3",
			xnametypes.MgmtHLSwitchEnclosure,
			MgmtHLSwitchEnclosure{
				Cabinet:               1,
				Chassis:               2,
				MgmtHLSwitchEnclosure: 3,
			},
		}, {
			"x1c2h3s4",
			xnametypes.MgmtHLSwitch,
			MgmtHLSwitch{
				Cabinet:               1,
				Chassis:               2,
				MgmtHLSwitchEnclosure: 3,
				MgmtHLSwitch:          4,
			},
		}, {
			"x1c2w3",
			xnametypes.MgmtSwitch,
			MgmtSwitch{
				Cabinet:    1,
				Chassis:    2,
				MgmtSwitch: 3,
			},
		}, {
			"x1c2w3j4",
			xnametypes.MgmtSwitchConnector,
			MgmtSwitchConnector{
				Cabinet:             1,
				Chassis:             2,
				MgmtSwitch:          3,
				MgmtSwitchConnector: 4,
			},
		}, {
			"x1c2r3",
			xnametypes.RouterModule,
			RouterModule{
				Cabinet:      1,
				Chassis:      2,
				RouterModule: 3,
			},
		}, {
			"x1c2r3b4",
			xnametypes.RouterBMC,
			RouterBMC{
				Cabinet:      1,
				Chassis:      2,
				RouterModule: 3,
				RouterBMC:    4,
			},
		}, {
			"x1c2s3",
			xnametypes.ComputeModule,
			ComputeModule{
				Cabinet:       1,
				Chassis:       2,
				ComputeModule: 3,
			},
		}, {
			"x1c2s3b4",
			xnametypes.NodeBMC,
			NodeBMC{
				Cabinet:       1,
				Chassis:       2,
				ComputeModule: 3,
				NodeBMC:       4,
			},
		}, {
			"x1c2s3b4n5",
			xnametypes.Node,
			Node{
				Cabinet:       1,
				Chassis:       2,
				ComputeModule: 3,
				NodeBMC:       4,
				Node:          5,
			},
		},
	}

	for _, test := range tests {
		xname := test.xname
		expectedHMSType := test.hmsType

		// Just a sanity check to verify that out test data is good
		if hmsType := xnametypes.GetHMSType(xname); hmsType != expectedHMSType {
			t.Errorf("unexpected HMS Type (%s) for xname (%s) in test data, expected (%s)", hmsType, xname, expectedHMSType)
		}

		// Verify FromString returns the HMS Type
		componentRaw, hmsType := FromString(xname)
		if expectedHMSType != hmsType {
			t.Error("Unexpected HMS Type:", hmsType, "expected:", expectedHMSType)
		}

		// Verify FromString returns the correct xname struct values
		if componentRaw != test.expectedComponent {
			t.Errorf("Unexpected xname struct (%v), expected (%v)", componentRaw, test.expectedComponent)
		}

		// Verify that GetHMSType works
		objXnameType, err := GetHMSType(componentRaw)
		if err != nil {
			t.Error("GetHMSType error:", err)
		}
		if expectedHMSType != objXnameType {
			t.Error("Unexpected HMS Type for xname struct:", objXnameType, "expected:", expectedHMSType)
		}

		// Verify the xname string built from the xname struct matches what was given to FromString
		generatedXname := componentRaw.(fmt.Stringer).String()
		if xname != generatedXname {
			t.Error("Unexpected generated xname:", generatedXname, "expected:", xname)
		}

		// Verify the HMS Type of the xname built FromString has the expected HMS Type
		generatedXnameType := xnametypes.GetHMSType(generatedXname)
		if expectedHMSType != generatedXnameType {
			t.Errorf("Unexpected generated xname %s (%s), expected (%s) %s", generatedXnameType, generatedXname, expectedHMSType, xname)
		}
	}
}

//
//
// Tests to verify that Parent/Children functions behave as expected
//
//

func TestSystemChildren(t *testing.T) {
	system := System{}

	// Create a child CDU
	cdu := system.CDU(1)
	expectedCDU := CDU{
		CDU: 1,
	}
	if !reflect.DeepEqual(expectedCDU, cdu) {
		t.Errorf("TestSystemChildren FAIL: Expected cdu=%v but instead got cdu=%v", expectedCDU, cdu)
	}

	// Create a child cabinet
	cabinet := system.Cabinet(1)
	expectedCabinet := Cabinet{
		Cabinet: 1,
	}
	if !reflect.DeepEqual(expectedCabinet, cabinet) {
		t.Errorf("TestSystemChildren FAIL: Expected cabinet=%v but instead got cabinet=%v", expectedCabinet, cabinet)
	}
}

func TestSystemParent(t *testing.T) {
	// A System doesn't have a parent
}

func TestCDUChildren(t *testing.T) {
	cdu := CDU{
		CDU: 1,
	}

	// Create a child CDUMgmtSwitch
	cduMgmtSwitch := cdu.CDUMgmtSwitch(2)
	expectedCDUMgmtSwitch := CDUMgmtSwitch{
		CDU:           1,
		CDUMgmtSwitch: 2,
	}
	if !reflect.DeepEqual(expectedCDUMgmtSwitch, cduMgmtSwitch) {
		t.Errorf("TestCDUChildren FAIL: Expected cduMgmtSwitch=%v but instead got cduMgmtSwitch=%v", expectedCDUMgmtSwitch, cduMgmtSwitch)
	}
}

func TestCDUParent(t *testing.T) {
	cdu := CDU{
		CDU: 1,
	}

	parent := cdu.Parent()
	expectedParent := System{}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestCDUParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestCabinetChildren(t *testing.T) {
	cabinet := Cabinet{
		Cabinet: 1,
	}

	// Create a child CabinetPDUController
	cabinetPDUController := cabinet.CabinetPDUController(2)
	expectedCabinetPDUController := CabinetPDUController{
		Cabinet:              1,
		CabinetPDUController: 2,
	}
	if !reflect.DeepEqual(expectedCabinetPDUController, cabinetPDUController) {
		t.Errorf("TestCabinetChildren FAIL: Expected cabinetPDUController=%v but instead got cabinetPDUController=%v", expectedCabinetPDUController, cabinetPDUController)
	}

	// Create a child Chassis
	chassis := cabinet.Chassis(2)
	expectedChassis := Chassis{
		Cabinet: 1,
		Chassis: 2,
	}
	if !reflect.DeepEqual(expectedChassis, chassis) {
		t.Errorf("TestCabinetChildren FAIL: Expected chassis=%v but instead got chassis=%v", expectedChassis, chassis)
	}
}

func TestCabinetParent(t *testing.T) {
	cabinet := Cabinet{
		Cabinet: 1,
	}

	parent := cabinet.Parent()
	expectedParent := System{}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestCabinetParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestCabinetPDUControllerChildren(t *testing.T) {
	// TODO no children structures have bene defined yet, but child xname formats have been defined
}

func TestCabinetPDUControllerParent(t *testing.T) {
	cabinetPDUController := CabinetPDUController{
		Cabinet:              1,
		CabinetPDUController: 2,
	}

	parent := cabinetPDUController.Parent()
	expectedParent := Cabinet{
		Cabinet: 1,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestCabinetPDUControllerParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestChassisChildren(t *testing.T) {
	chassis := Chassis{
		Cabinet: 1,
		Chassis: 2,
	}

	// Create a child ComputeModule
	computeModule := chassis.ComputeModule(3)
	expectedComputeModule := ComputeModule{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
	}
	if !reflect.DeepEqual(expectedComputeModule, computeModule) {
		t.Errorf("TestChassisChildren FAIL: Expected computeModule=%v but instead got computeModule=%v", expectedComputeModule, computeModule)
	}

	// Create a child MgmtSwitch
	mgmtSwitch := chassis.MgmtSwitch(3)
	expectedMgmtSwitch := MgmtSwitch{
		Cabinet:    1,
		Chassis:    2,
		MgmtSwitch: 3,
	}
	if !reflect.DeepEqual(expectedMgmtSwitch, mgmtSwitch) {
		t.Errorf("TestChassisChildren FAIL: Expected mgmtSwitch=%v but instead got mgmtSwitch=%v", expectedMgmtSwitch, mgmtSwitch)
	}

	// Create a child MgmtHLSwitchEnclosure
	mgmtHLSwitchEnclosure := chassis.MgmtHLSwitchEnclosure(3)
	expectedMgmtHLSwitchEnclosure := MgmtHLSwitchEnclosure{
		Cabinet:               1,
		Chassis:               2,
		MgmtHLSwitchEnclosure: 3,
	}
	if !reflect.DeepEqual(expectedMgmtHLSwitchEnclosure, mgmtHLSwitchEnclosure) {
		t.Errorf("TestChassisChildren FAIL: Expected mgmtHLSwitchEnclosure=%v but instead got mgmtHLSwitchEnclosure=%v", expectedMgmtHLSwitchEnclosure, mgmtHLSwitchEnclosure)
	}

	// Create a child RouterModule
	routerModule := chassis.RouterModule(3)
	expetedRouterModule := RouterModule{
		Cabinet:      1,
		Chassis:      2,
		RouterModule: 3,
	}
	if !reflect.DeepEqual(expetedRouterModule, routerModule) {
		t.Errorf("TestChassisChildren FAIL: Expected routerModule=%v but instead got routerModule=%v", expetedRouterModule, routerModule)
	}
}

func TestChassisParent(t *testing.T) {
	chassis := Chassis{
		Cabinet: 1,
		Chassis: 2,
	}

	parent := chassis.Parent()
	expectedParent := Cabinet{
		Cabinet: 1,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestChassisParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestChassisBMCChildren(t *testing.T) {
	// TODO no children structures have bene defined yet, but child xname formats have been defined
}

func TestChassisBMCParent(t *testing.T) {
	chassisBMC := ChassisBMC{
		Cabinet:    1,
		Chassis:    2,
		ChassisBMC: 0,
	}

	parent := chassisBMC.Parent()
	expectedParent := Chassis{
		Cabinet: 1,
		Chassis: 2,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestChassisBMCParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestMgmtSwitchChildren(t *testing.T) {
	mgmtSwitch := MgmtSwitch{
		Cabinet:    1,
		Chassis:    2,
		MgmtSwitch: 3,
	}

	// Create a child MgmtSwitchConnector
	mgmtSwitchConnector := mgmtSwitch.MgmtSwitchConnector(4)
	expectedMgmtSwitchConnector := MgmtSwitchConnector{
		Cabinet:             1,
		Chassis:             2,
		MgmtSwitch:          3,
		MgmtSwitchConnector: 4,
	}
	if !reflect.DeepEqual(expectedMgmtSwitchConnector, mgmtSwitchConnector) {
		t.Errorf("TestMgmtSwitchChildren FAIL: Expected mgmtSwitchConnector=%v but instead got mgmtSwitchConnector=%v", expectedMgmtSwitchConnector, mgmtSwitchConnector)
	}
}

func TestMgmtSwitchParent(t *testing.T) {
	mgmtSwitch := MgmtSwitch{
		Cabinet:    1,
		Chassis:    2,
		MgmtSwitch: 3,
	}

	parent := mgmtSwitch.Parent()
	expectedParent := Chassis{
		Cabinet: 1,
		Chassis: 2,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestMgmtSwitchParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestMgmtSwitchConnectorChildren(t *testing.T) {
	// There are no childlen for a MgmtSwitchConnector
}

func TestMgmtSwitchConnectorParent(t *testing.T) {
	mgmtSwitchConnector := MgmtSwitchConnector{
		Cabinet:             1,
		Chassis:             2,
		MgmtSwitch:          3,
		MgmtSwitchConnector: 4,
	}

	parent := mgmtSwitchConnector.Parent()
	expectedParent := MgmtSwitch{
		Cabinet:    1,
		Chassis:    2,
		MgmtSwitch: 3,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestMgmtSwitchConnectorParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestMgmtHLSwitchEnclosureChildren(t *testing.T) {
	mgmtHLSwitchEnclosure := MgmtHLSwitchEnclosure{
		Cabinet:               1,
		Chassis:               2,
		MgmtHLSwitchEnclosure: 3,
	}

	// Create a child MgmtHLSwitch
	mgmtHLSwitch := mgmtHLSwitchEnclosure.MgmtHLSwitch(4)
	expectedMgmtHLSwitch := MgmtHLSwitch{
		Cabinet:               1,
		Chassis:               2,
		MgmtHLSwitchEnclosure: 3,
		MgmtHLSwitch:          4,
	}
	if !reflect.DeepEqual(expectedMgmtHLSwitch, mgmtHLSwitch) {
		t.Errorf("TestMgmtHLSwitchEnclosureChildren FAIL: Expected mgmtHLSwitch=%v but instead got mgmtHLSwitch=%v", expectedMgmtHLSwitch, mgmtHLSwitch)
	}
}

func TestMgmtHLSwitchEnclosureParent(t *testing.T) {
	mgmtHLSwitchEnclosure := MgmtHLSwitchEnclosure{
		Cabinet:               1,
		Chassis:               2,
		MgmtHLSwitchEnclosure: 3,
	}

	parent := mgmtHLSwitchEnclosure.Parent()
	expectedParent := Chassis{
		Cabinet: 1,
		Chassis: 2,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestMgmtHLSwitchEnclosureParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestMgmtHLSwitchChildren(t *testing.T) {
	// TODO no children structures have bene defined yet, and currently no child xname formats have been defined
}

func TestMgmtHLSwitchParent(t *testing.T) {
	mgmtHLSwitch := MgmtHLSwitch{
		Cabinet:               1,
		Chassis:               2,
		MgmtHLSwitchEnclosure: 3,
		MgmtHLSwitch:          4,
	}

	parent := mgmtHLSwitch.Parent()
	expectedParent := MgmtHLSwitchEnclosure{
		Cabinet:               1,
		Chassis:               2,
		MgmtHLSwitchEnclosure: 3,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestMgmtHLSwitchParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestRouterModuleChildren(t *testing.T) {
	routerModule := RouterModule{
		Cabinet:      1,
		Chassis:      2,
		RouterModule: 3,
	}

	// Create a child RouterBMC
	routerBMC := routerModule.RouterBMC(4)
	expectedRouterBMC := RouterBMC{
		Cabinet:      1,
		Chassis:      2,
		RouterModule: 3,
		RouterBMC:    4,
	}
	if !reflect.DeepEqual(expectedRouterBMC, routerBMC) {
		t.Errorf("TestRouterModuleChildren FAIL: Expected routerBMC=%v but instead got routerBMC=%v", expectedRouterBMC, routerBMC)
	}
}

func TestRouterModuleParent(t *testing.T) {
	routerModule := RouterModule{
		Cabinet:      1,
		Chassis:      2,
		RouterModule: 3,
	}

	parent := routerModule.Parent()
	expectedParent := Chassis{
		Cabinet: 1,
		Chassis: 2,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestRouterModuleParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestRouterBMCChildren(t *testing.T) {
	// TODO no children structures have bene defined yet, but child xname formats have been defined
}

func TestRouterBMCParent(t *testing.T) {
	routerModule := RouterBMC{
		Cabinet:      1,
		Chassis:      2,
		RouterModule: 3,
		RouterBMC:    4,
	}

	parent := routerModule.Parent()
	expectedParent := RouterModule{
		Cabinet:      1,
		Chassis:      2,
		RouterModule: 3,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestRouterBMCParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestComputeModuleChildren(t *testing.T) {
	computeModule := ComputeModule{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
	}

	// Create a child NodeBMC
	nodeBMC := computeModule.NodeBMC(4)
	expectedNodeBMC := NodeBMC{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
		NodeBMC:       4,
	}

	if !reflect.DeepEqual(expectedNodeBMC, nodeBMC) {
		t.Errorf("TestComputeModuleChildren FAIL: Expected nodeBMC=%v but instead got nodeBMC=%v", expectedNodeBMC, nodeBMC)
	}
}

func TestComputeModuleParent(t *testing.T) {
	computeModule := ComputeModule{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
	}

	parent := computeModule.Parent()
	expectedParent := Chassis{
		Cabinet: 1,
		Chassis: 2,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestComputeModuleParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestNodeBMCChildren(t *testing.T) {
	nodeBMC := NodeBMC{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
		NodeBMC:       4,
	}

	// Create a child Node
	node := nodeBMC.Node(0)
	expectedNode := Node{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
		NodeBMC:       4,
		Node:          0,
	}
	if !reflect.DeepEqual(expectedNode, node) {
		t.Errorf("TestNodeBMCChildren FAIL: Expected node=%v but instead got node=%v", expectedNode, node)
	}
}

func TestNodeBMCParent(t *testing.T) {
	nodeBMC := NodeBMC{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
		NodeBMC:       4,
	}

	parent := nodeBMC.Parent()
	expectedParent := ComputeModule{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestNodeBMCParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

func TestNodeChildren(t *testing.T) {
	// TODO no children structures have bene defined yet, but child xname formats have been defined
}

func TestNodeParent(t *testing.T) {
	node := Node{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
		NodeBMC:       4,
		Node:          0,
	}

	parent := node.Parent()
	expectedParent := NodeBMC{
		Cabinet:       1,
		Chassis:       2,
		ComputeModule: 3,
		NodeBMC:       4,
	}
	if !reflect.DeepEqual(expectedParent, parent) {
		t.Errorf("TestNodeParent FAIL: Expected parent=%v but instead got parent=%v", expectedParent, parent)
	}
}

// TestValidateValidXnames unit test function for Validate
func TestValidateValidXnames(t *testing.T) {
	invalidXnames := []interface{}{
		CDU{
			CDU: -1,
		},
		CDUMgmtSwitch{
			CDU:           0,
			CDUMgmtSwitch: -1,
		},
		Cabinet{
			Cabinet: -1,
		},
		Chassis{
			Cabinet: 1,
			Chassis: -1,
		},
		ChassisBMC{
			Cabinet:    1,
			Chassis:    1,
			ChassisBMC: 1, // This value needs to be 0
		},
		ChassisBMC{
			Cabinet:    -1,
			Chassis:    -1,
			ChassisBMC: -1,
		},
	}

	for _, xname := range invalidXnames {
		if err := xname.(Validator).Validate(); err == nil {
			t.Errorf("No validation error returned for xname (%v)", xname)
		}
	}

}

// TestValidateInvalidXnames unit test function for Validate
func TestValidateInvalidXnames(t *testing.T) {

	validXnames := []string{
		"d0",
		"d0w0",
		"d0w1",
		"d0w30",
		"x0",
		"x0b0",
		"x0c0",
		"x0c0b0",
		"x0c0f0",
		"x0c0r0",
		"x0c0r0t0f0",
		"x0c0r0t0f1",
		"x0c0r1a0",
		"x0c0r1a0l64",
		"x0c0s0",
		"x0c0s4j1",
		"x0c0s64",
		"x0c0t0",
		"x0d0",
		"x0e0",
		"x0m0",
		"x0m0i0",
		"x0m0p0",
		"x0m0p0j1",
		"x1",
		"x10",
		"x100",
		"x1000",
		"x1234c0r1t0f0",
		"x16c2s3b0n1",
		"x16c2s3b0n1g1",
		"x16c2s3b0n1g1k1",
		"x16c2s3b0n1p1",
		"x16c3r2f0",
		"x16c3r4j7",
		"x16c3r4j7p0",
		"x16c3s0b0n1d3",
		"x16c3s0b0n1h1",
		"x16c3s0b0n1i1",
		"x16c3s0b1n0a1",
		"x16c3s1b1",
		"x16c3s4b1f0",
		"x1c0",
		"x1c0h1s2",
		"x1c0r2b0",
		"x1c0r2e0",
		"x1c0r47j31",
		"x1c0w1",
		"x1c2r3b0i0",
		"x1m0p0j1",
		"x1m0p1j10",
		"x1m1p0j64",
		"x2",
		"x3",
		"x3c0s16e0",
	}

	for _, xname := range validXnames {
		x, xType := FromString(xname)
		if xType == xnametypes.HMSTypeInvalid {
			t.Errorf("FromString failed on %s", xname)
			continue
		}

		if err := x.(Validator).Validate(); err != nil {
			t.Errorf("Unexpected validation error (%v) with xname (%s)", err, xname)
		}

	}

}
