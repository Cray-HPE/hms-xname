// MIT License
//
// (C) Copyright [2021] Hewlett Packard Enterprise Development LP
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
// OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package xname

import (
	"fmt"

	"github.com/Cray-HPE/hms-xname/xnametypes"
)

// s0
type System struct{}

func (s System) String() string {
	return "s0"
}

func (s System) Validate() error {
	return nil
}

func (s System) CDU(coolingGroup int) CDU {
	return CDU{
		CoolingGroup: coolingGroup,
	}
}

func (s System) Cabinet(cabinet int) Cabinet {
	return Cabinet{
		Cabinet: cabinet,
	}
}

// dD
type CDU struct {
	CoolingGroup int // D: 0-999
}

func (c CDU) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.CDU)
	return fmt.Sprintf(formatStr, c.CoolingGroup)
}

func (c CDU) Parent() System {
	return System{}
}

func (c CDU) Validate() error {
	xname := c.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid CDU xname: %s", xname)
	}

	return nil
}

func (c CDU) CDUMgmtSwitch(slot int) CDUMgmtSwitch {
	return CDUMgmtSwitch{
		CoolingGroup: c.CoolingGroup,
		Slot:         slot,
	}
}

// dDwW
type CDUMgmtSwitch struct {
	CoolingGroup int // D: 0-999
	Slot         int // W: 0-31
}

func (cms CDUMgmtSwitch) String() string { 
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.CDUMgmtSwitch)
	return fmt.Sprintf(formatStr, cms.CoolingGroup, cms.Slot)
}

func (cms CDUMgmtSwitch) Parent() CDU {
	return CDU{
		CoolingGroup: cms.CoolingGroup,
	}
}

func (cms CDUMgmtSwitch) Validate() error {
	xname := cms.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid CDUMgmtSwitch xname: %s", xname)
	}

	return nil
}

// xX
type Cabinet struct {
	Cabinet int // X: 0-999
}

func (c Cabinet) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.Cabinet)
	return fmt.Sprintf(formatStr, c.Cabinet)
}

func (c Cabinet) Parent() System {
	return System{}
}

func (c Cabinet) Validate() error {
	xname := c.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid Cabinet xname: %s", xname)
	}

	return nil
}

func (c Cabinet) Chassis(chassis int) Chassis {
	return Chassis{
		Cabinet: c.Cabinet,
		Chassis: chassis,
	}
}

func (c Cabinet) CabinetPDUController(pduController int) CabinetPDUController {
	return CabinetPDUController{
		Cabinet:       c.Cabinet,
		PDUController: pduController,
	}
}

// xXmM
type CabinetPDUController struct {
	Cabinet       int // X: 0-999
	PDUController int // M: 0-3
}

func (p CabinetPDUController) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.CabinetPDUController)
	return fmt.Sprintf(formatStr, p.Cabinet, p.Cabinet)
}

func (p CabinetPDUController) Parent() Cabinet {
	return Cabinet{
		Cabinet: p.Cabinet,
	}
}

func (p CabinetPDUController) Validate() error {
	xname := p.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid CabinetPDUController xname: %s", xname)
	}

	return nil
}

// xXcC
type Chassis struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
}

func (c Chassis) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.Chassis)
	return fmt.Sprintf(formatStr, c.Cabinet, c.Chassis)
}

func (c Chassis) Parent() Cabinet {
	return Cabinet{
		Cabinet: c.Cabinet,
	}
}

func (c Chassis) Validate() error {
	xname := c.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid Chassis xname: %s", xname)
	}

	return nil
}

func (c Chassis) ChassisBMC(bmc int) ChassisBMC {
	return ChassisBMC{
		Cabinet: c.Cabinet,
		Chassis: c.Chassis,
		BMC: bmc,
	}
}

func (c Chassis) MgmtHLSwitchEnclosure(slot int) MgmtHLSwitchEnclosure {
	return MgmtHLSwitchEnclosure{
		Cabinet: c.Cabinet,
		Chassis: c.Chassis,
		Slot:    slot,
	}
}

func (c Chassis) MgmtSwitch(slot int) MgmtSwitch {
	return MgmtSwitch{
		Cabinet: c.Cabinet,
		Chassis: c.Chassis,
		Slot:    slot,
	}
}

// This is a convience function, as we normally do not work with MgmtHLSwitchEnclosures directly
func (c Chassis) MgmtHLSwitch(slot, space int) MgmtHLSwitch {
	return c.MgmtHLSwitchEnclosure(slot).MgmtHLSwitch(space)
}

func (c Chassis) RouterModule(slot int) RouterModule {
	return RouterModule{
		Cabinet: c.Cabinet,
		Chassis: c.Chassis,
		Slot:    slot,
	}
}

// This is a convince function, as we normally do not work with RouterModules directly.
func (c Chassis) RouterBMC(slot, bmc int) RouterBMC {
	return c.RouterModule(slot).RouterBMC(bmc)
}

func (c Chassis) ComputeModule(slot int) ComputeModule {
	return ComputeModule{
		Cabinet: c.Cabinet,
		Chassis: c.Chassis,
		Slot:    slot,
	}
}

func (c Chassis) NodeBMC(slot, bmc int) NodeBMC {
	return c.ComputeModule(slot).NodeBMC(bmc)
}

// xXcCbB
// Mountain and Hill have only b0
// River does not have ChassisBMCs
type ChassisBMC struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	BMC int // B: 0
}

func (c ChassisBMC) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.ChassisBMC)
	return fmt.Sprintf(formatStr, c.Cabinet, c.Chassis, c.BMC)
}

func (c ChassisBMC) Parent() Chassis {
	return Chassis{
		Cabinet: c.Cabinet,
		Chassis: c.Chassis,
	}
}

func (c ChassisBMC) Validate() error {
	xname := c.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid Chassis xname: %s", xname)
	}

	return nil
}

// xXcCwW
type MgmtSwitch struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	Slot    int // W: 1-48
}

func (ms MgmtSwitch) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.MgmtSwitch)
	return fmt.Sprintf(formatStr, ms.Cabinet, ms.Chassis, ms.Slot)
}

func (ms MgmtSwitch) Parent() Chassis {
	return Chassis{
		Cabinet: ms.Cabinet,
		Chassis: ms.Chassis,
	}
}

func (ms MgmtSwitch) Validate() error {
	xname := ms.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid MgmtSwitch xname: %s", xname)
	}

	return nil
}

func (ms MgmtSwitch) MgmtSwitchConnector(switchPort int) MgmtSwitchConnector {
	return MgmtSwitchConnector{
		Cabinet:    ms.Cabinet,
		Chassis:    ms.Chassis,
		Slot:       ms.Slot,
		SwitchPort: switchPort,
	}
}

// xXcCwWjJ
type MgmtSwitchConnector struct {
	Cabinet    int // X: 0-999
	Chassis    int // C: 0-7
	Slot       int // W: 1-48
	SwitchPort int // J: 1-32 // TODO the HSOS page, should allow upto at least 48
}

func (msc MgmtSwitchConnector) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.MgmtSwitchConnector)
	return fmt.Sprintf(formatStr, msc.Cabinet, msc.Chassis, msc.Slot, msc.SwitchPort)
}

func (msc MgmtSwitchConnector) Parent() MgmtSwitch {
	return MgmtSwitch{
		Cabinet: msc.Cabinet,
		Chassis: msc.Chassis,
		Slot:    msc.Slot,
	}
}

func (msc MgmtSwitchConnector) Validate() error {
	xname := msc.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid MgmtSwitchConnector xname: %s", xname)
	}

	return nil
}

// xXcChH
type MgmtHLSwitchEnclosure struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	Slot    int // H: 1-48
}

func (enclosure MgmtHLSwitchEnclosure) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.MgmtHLSwitchEnclosure)
	return fmt.Sprintf(formatStr, enclosure.Cabinet, enclosure.Chassis, enclosure.Slot)
}

func (enclosure MgmtHLSwitchEnclosure) Parent() Chassis {
	return Chassis{
		Cabinet: enclosure.Cabinet,
		Chassis: enclosure.Chassis,
	}
}

func (enclosure MgmtHLSwitchEnclosure) Validate() error {
	xname := enclosure.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid MgmtHLSwitchEnclosure xname: %s", xname)
	}

	return nil
}

func (enclosure MgmtHLSwitchEnclosure) MgmtHLSwitch(space int) MgmtHLSwitch {
	return MgmtHLSwitch{
		Cabinet: enclosure.Cabinet,
		Chassis: enclosure.Chassis,
		Slot:    enclosure.Slot,
		Space:   space,
	}
}

//xXcChHsS
type MgmtHLSwitch struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	Slot    int // H: 1-48
	Space   int // S: 1-4
}

func (mhls MgmtHLSwitch) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.MgmtHLSwitch)
	return fmt.Sprintf(formatStr, mhls.Cabinet, mhls.Chassis, mhls.Slot, mhls.Space)
}

func (mhls MgmtHLSwitch) Parent() MgmtHLSwitchEnclosure {
	return MgmtHLSwitchEnclosure{
		Cabinet: mhls.Cabinet,
		Chassis: mhls.Chassis,
		Slot:    mhls.Slot,
	}
}

func (mhls MgmtHLSwitch) Validate() error {
	xname := mhls.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid MgmtHLSwitch xname: %s", xname)
	}

	return nil
}

// xXcCrR
type RouterModule struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	Slot    int // R: 0-64
}

func (rm RouterModule) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.RouterModule)
	return fmt.Sprintf(formatStr, rm.Cabinet, rm.Chassis, rm.Slot)
}

func (rm RouterModule) Parent() Chassis {
	return Chassis{
		Cabinet: rm.Cabinet,
		Chassis: rm.Chassis,
	}
}

func (rm RouterModule) Validate() error {
	xname := rm.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid RouterModule xname: %s", xname)
	}

	return nil
}

func (rm RouterModule) RouterBMC(bmc int) RouterBMC {
	return RouterBMC{
		Cabinet: rm.Cabinet,
		Chassis: rm.Chassis,
		Slot:    rm.Slot,
		BMC:     bmc,
	}
}

// xXcCrRbB
type RouterBMC struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	Slot    int // R: 0-64
	BMC     int // B: 0
}

func (bmc RouterBMC) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.RouterBMC)
	return fmt.Sprintf(formatStr, bmc.Cabinet, bmc.Chassis, bmc.Slot, bmc.BMC)
}

func (bmc RouterBMC) Parent() RouterModule {
	return RouterModule{
		Cabinet: bmc.Cabinet,
		Chassis: bmc.Chassis,
		Slot:    bmc.Slot,
	}
}

func (bmc RouterBMC) Validate() error {
	xname := bmc.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid RouterBMC xname: %s", xname)
	}

	return nil
}

// xXcCsS
type ComputeModule struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	Slot    int // S: 1-63
}

func (cm ComputeModule) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.ComputeModule)
	return fmt.Sprintf(formatStr, cm.Cabinet, cm.Chassis, cm.Slot)
}

func (cm ComputeModule) Parent() Chassis {
	return Chassis{
		Cabinet: cm.Cabinet,
		Chassis: cm.Chassis,
	}
}

func (cm ComputeModule) Validate() error {
	xname := cm.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid ComputeModule xname: %s", xname)
	}

	return nil
}

func (cm ComputeModule) NodeBMC(bmc int) NodeBMC {
	return NodeBMC{
		Cabinet: cm.Cabinet,
		Chassis: cm.Chassis,
		Slot:    cm.Slot,
		BMC:     bmc,
	}
}

// xXcCsSbB
// Node Card/Node BMC
type NodeBMC struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	Slot    int // S: 1-63
	BMC     int // B: 0-1
}

func (bmc NodeBMC) Parent() ComputeModule {
	return ComputeModule{
		Cabinet: bmc.Cabinet,
		Chassis: bmc.Chassis,
		Slot:    bmc.Slot,
	}
}

func (bmc NodeBMC) Validate() error {
	xname := bmc.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid NodeBMC xname: %s", xname)
	}

	return nil
}

func (bmc NodeBMC) Node(node int) Node {
	return Node{
		Cabinet: bmc.Cabinet,
		Chassis: bmc.Chassis,
		Slot:    bmc.Slot,
		BMC:     bmc.BMC,
		Node:    node,
	}
}

func (bmc NodeBMC) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.NodeBMC)
	return fmt.Sprintf(formatStr, bmc.Cabinet, bmc.Chassis, bmc.Slot, bmc.BMC)
}

// xCcCsSbBnN
type Node struct {
	Cabinet int // X: 0-999
	Chassis int // C: 0-7
	Slot    int // S: 1-63
	BMC     int // B: 0-1 - TODO the HSOS document is wrong here. as we do actually use greater than 1
	Node    int // N: 0-7
}

func (n Node) String() string {
	formatStr, _, _ := xnametypes.GetHMSTypeFormatString(xnametypes.Node)
	return fmt.Sprintf(formatStr, n.Cabinet, n.Chassis, n.Slot, n.BMC, n.Node)
}

func (n Node) Validate() error {
	xname := n.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid node xname: %s", xname)
	}

	return nil
}

func (n Node) Parent() NodeBMC {
	return NodeBMC{
		Cabinet: n.Cabinet,
		Chassis: n.Chassis,
		Slot:    n.Slot,
		BMC:     n.BMC,
	}
}
