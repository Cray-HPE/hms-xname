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
)

// System - sS
type System struct {
}

// String will stringify System into the format of sS
func (x System) String() string {
	return fmt.Sprintf(
		"s0",
	)
}

// CDU will get a child component with the specified ordinal
func (x System) CDU(cDU int) CDU {
	return CDU{
		CDU: cDU,
	}
}

// Cabinet will get a child component with the specified ordinal
func (x System) Cabinet(cabinet int) Cabinet {
	return Cabinet{
		Cabinet: cabinet,
	}
}

// CDU - dD
type CDU struct {
	CDU int
}

// String will stringify CDU into the format of dD
func (x CDU) String() string {
	return fmt.Sprintf(
		"d%d",
		x.CDU,
	)
}

// Parent will determine the parent of this CDU
func (x CDU) Parent() System {
	return System{}
}

// CDUMgmtSwitch will get a child component with the specified ordinal
func (x CDU) CDUMgmtSwitch(cDUMgmtSwitch int) CDUMgmtSwitch {
	return CDUMgmtSwitch{
		CDU:           x.CDU,
		CDUMgmtSwitch: cDUMgmtSwitch,
	}
}

// CDUMgmtSwitch - dDwW
type CDUMgmtSwitch struct {
	CDU           int
	CDUMgmtSwitch int
}

// String will stringify CDUMgmtSwitch into the format of dDwW
func (x CDUMgmtSwitch) String() string {
	return fmt.Sprintf(
		"d%dw%d",
		x.CDU,
		x.CDUMgmtSwitch,
	)
}

// Parent will determine the parent of this CDUMgmtSwitch
func (x CDUMgmtSwitch) Parent() CDU {
	return CDU{
		CDU: x.CDU,
	}
}

// Cabinet - xX
type Cabinet struct {
	Cabinet int
}

// String will stringify Cabinet into the format of xX
func (x Cabinet) String() string {
	return fmt.Sprintf(
		"x%d",
		x.Cabinet,
	)
}

// Parent will determine the parent of this Cabinet
func (x Cabinet) Parent() System {
	return System{}
}

// CEC will get a child component with the specified ordinal
func (x Cabinet) CEC(cEC int) CEC {
	return CEC{
		Cabinet: x.Cabinet,
		CEC:     cEC,
	}
}

// CabinetBMC will get a child component with the specified ordinal
func (x Cabinet) CabinetBMC(cabinetBMC int) CabinetBMC {
	return CabinetBMC{
		Cabinet:    x.Cabinet,
		CabinetBMC: cabinetBMC,
	}
}

// CabinetCDU will get a child component with the specified ordinal
func (x Cabinet) CabinetCDU(cabinetCDU int) CabinetCDU {
	return CabinetCDU{
		Cabinet:    x.Cabinet,
		CabinetCDU: cabinetCDU,
	}
}

// CabinetPDUController will get a child component with the specified ordinal
func (x Cabinet) CabinetPDUController(cabinetPDUController int) CabinetPDUController {
	return CabinetPDUController{
		Cabinet:              x.Cabinet,
		CabinetPDUController: cabinetPDUController,
	}
}

// Chassis will get a child component with the specified ordinal
func (x Cabinet) Chassis(chassis int) Chassis {
	return Chassis{
		Cabinet: x.Cabinet,
		Chassis: chassis,
	}
}

// CEC - xXeE
type CEC struct {
	Cabinet int
	CEC     int
}

// String will stringify CEC into the format of xXeE
func (x CEC) String() string {
	return fmt.Sprintf(
		"x%de%d",
		x.Cabinet,
		x.CEC,
	)
}

// Parent will determine the parent of this CEC
func (x CEC) Parent() Cabinet {
	return Cabinet{
		Cabinet: x.Cabinet,
	}
}

// CabinetBMC - xXbB
type CabinetBMC struct {
	Cabinet    int
	CabinetBMC int
}

// String will stringify CabinetBMC into the format of xXbB
func (x CabinetBMC) String() string {
	return fmt.Sprintf(
		"x%db%d",
		x.Cabinet,
		x.CabinetBMC,
	)
}

// Parent will determine the parent of this CabinetBMC
func (x CabinetBMC) Parent() Cabinet {
	return Cabinet{
		Cabinet: x.Cabinet,
	}
}

// CabinetCDU - xXdD
type CabinetCDU struct {
	Cabinet    int
	CabinetCDU int
}

// String will stringify CabinetCDU into the format of xXdD
func (x CabinetCDU) String() string {
	return fmt.Sprintf(
		"x%dd%d",
		x.Cabinet,
		x.CabinetCDU,
	)
}

// Parent will determine the parent of this CabinetCDU
func (x CabinetCDU) Parent() Cabinet {
	return Cabinet{
		Cabinet: x.Cabinet,
	}
}

// CabinetPDUController - xXmM
type CabinetPDUController struct {
	Cabinet              int
	CabinetPDUController int
}

// String will stringify CabinetPDUController into the format of xXmM
func (x CabinetPDUController) String() string {
	return fmt.Sprintf(
		"x%dm%d",
		x.Cabinet,
		x.CabinetPDUController,
	)
}

// Parent will determine the parent of this CabinetPDUController
func (x CabinetPDUController) Parent() Cabinet {
	return Cabinet{
		Cabinet: x.Cabinet,
	}
}

// CabinetPDU will get a child component with the specified ordinal
func (x CabinetPDUController) CabinetPDU(cabinetPDU int) CabinetPDU {
	return CabinetPDU{
		Cabinet:              x.Cabinet,
		CabinetPDUController: x.CabinetPDUController,
		CabinetPDU:           cabinetPDU,
	}
}

// CabinetPDUNic will get a child component with the specified ordinal
func (x CabinetPDUController) CabinetPDUNic(cabinetPDUNic int) CabinetPDUNic {
	return CabinetPDUNic{
		Cabinet:              x.Cabinet,
		CabinetPDUController: x.CabinetPDUController,
		CabinetPDUNic:        cabinetPDUNic,
	}
}

// CabinetPDU - xXmMpP
type CabinetPDU struct {
	Cabinet              int
	CabinetPDUController int
	CabinetPDU           int
}

// String will stringify CabinetPDU into the format of xXmMpP
func (x CabinetPDU) String() string {
	return fmt.Sprintf(
		"x%dm%dp%d",
		x.Cabinet,
		x.CabinetPDUController,
		x.CabinetPDU,
	)
}

// Parent will determine the parent of this CabinetPDU
func (x CabinetPDU) Parent() CabinetPDUController {
	return CabinetPDUController{
		Cabinet:              x.Cabinet,
		CabinetPDUController: x.CabinetPDUController,
	}
}

// CabinetPDUOutlet will get a child component with the specified ordinal
func (x CabinetPDU) CabinetPDUOutlet(cabinetPDUOutlet int) CabinetPDUOutlet {
	return CabinetPDUOutlet{
		Cabinet:              x.Cabinet,
		CabinetPDUController: x.CabinetPDUController,
		CabinetPDU:           x.CabinetPDU,
		CabinetPDUOutlet:     cabinetPDUOutlet,
	}
}

// CabinetPDUPowerConnector will get a child component with the specified ordinal
func (x CabinetPDU) CabinetPDUPowerConnector(cabinetPDUPowerConnector int) CabinetPDUPowerConnector {
	return CabinetPDUPowerConnector{
		Cabinet:                  x.Cabinet,
		CabinetPDUController:     x.CabinetPDUController,
		CabinetPDU:               x.CabinetPDU,
		CabinetPDUPowerConnector: cabinetPDUPowerConnector,
	}
}

// CabinetPDUOutlet - xXmMpPjJ
type CabinetPDUOutlet struct {
	Cabinet              int
	CabinetPDUController int
	CabinetPDU           int
	CabinetPDUOutlet     int
}

// String will stringify CabinetPDUOutlet into the format of xXmMpPjJ
func (x CabinetPDUOutlet) String() string {
	return fmt.Sprintf(
		"x%dm%dp%dj%d",
		x.Cabinet,
		x.CabinetPDUController,
		x.CabinetPDU,
		x.CabinetPDUOutlet,
	)
}

// Parent will determine the parent of this CabinetPDUOutlet
func (x CabinetPDUOutlet) Parent() CabinetPDU {
	return CabinetPDU{
		Cabinet:              x.Cabinet,
		CabinetPDUController: x.CabinetPDUController,
		CabinetPDU:           x.CabinetPDU,
	}
}

// CabinetPDUPowerConnector - xXmMpPvV
type CabinetPDUPowerConnector struct {
	Cabinet                  int
	CabinetPDUController     int
	CabinetPDU               int
	CabinetPDUPowerConnector int
}

// String will stringify CabinetPDUPowerConnector into the format of xXmMpPvV
func (x CabinetPDUPowerConnector) String() string {
	return fmt.Sprintf(
		"x%dm%dp%dv%d",
		x.Cabinet,
		x.CabinetPDUController,
		x.CabinetPDU,
		x.CabinetPDUPowerConnector,
	)
}

// Parent will determine the parent of this CabinetPDUPowerConnector
func (x CabinetPDUPowerConnector) Parent() CabinetPDU {
	return CabinetPDU{
		Cabinet:              x.Cabinet,
		CabinetPDUController: x.CabinetPDUController,
		CabinetPDU:           x.CabinetPDU,
	}
}

// CabinetPDUNic - xXmMpPiI
type CabinetPDUNic struct {
	Cabinet              int
	CabinetPDUController int
	CabinetPDUNic        int
}

// String will stringify CabinetPDUNic into the format of xXmMpPiI
func (x CabinetPDUNic) String() string {
	return fmt.Sprintf(
		"x%dm%di%d",
		x.Cabinet,
		x.CabinetPDUController,
		x.CabinetPDUNic,
	)
}

// Parent will determine the parent of this CabinetPDUNic
func (x CabinetPDUNic) Parent() CabinetPDUController {
	return CabinetPDUController{
		Cabinet:              x.Cabinet,
		CabinetPDUController: x.CabinetPDUController,
	}
}

// Chassis - xXcC
type Chassis struct {
	Cabinet int
	Chassis int
}

// String will stringify Chassis into the format of xXcC
func (x Chassis) String() string {
	return fmt.Sprintf(
		"x%dc%d",
		x.Cabinet,
		x.Chassis,
	)
}

// Parent will determine the parent of this Chassis
func (x Chassis) Parent() Cabinet {
	return Cabinet{
		Cabinet: x.Cabinet,
	}
}

// CMMFpga will get a child component with the specified ordinal
func (x Chassis) CMMFpga(cMMFpga int) CMMFpga {
	return CMMFpga{
		Cabinet: x.Cabinet,
		Chassis: x.Chassis,
		CMMFpga: cMMFpga,
	}
}

// CMMRectifier will get a child component with the specified ordinal
func (x Chassis) CMMRectifier(cMMRectifier int) CMMRectifier {
	return CMMRectifier{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		CMMRectifier: cMMRectifier,
	}
}

// ChassisBMC will get a child component with the specified ordinal
func (x Chassis) ChassisBMC(chassisBMC int) ChassisBMC {
	return ChassisBMC{
		Cabinet:    x.Cabinet,
		Chassis:    x.Chassis,
		ChassisBMC: chassisBMC,
	}
}

// ComputeModule will get a child component with the specified ordinal
func (x Chassis) ComputeModule(computeModule int) ComputeModule {
	return ComputeModule{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: computeModule,
	}
}

// MgmtHLSwitchEnclosure will get a child component with the specified ordinal
func (x Chassis) MgmtHLSwitchEnclosure(mgmtHLSwitchEnclosure int) MgmtHLSwitchEnclosure {
	return MgmtHLSwitchEnclosure{
		Cabinet:               x.Cabinet,
		Chassis:               x.Chassis,
		MgmtHLSwitchEnclosure: mgmtHLSwitchEnclosure,
	}
}

// MgmtSwitch will get a child component with the specified ordinal
func (x Chassis) MgmtSwitch(mgmtSwitch int) MgmtSwitch {
	return MgmtSwitch{
		Cabinet:    x.Cabinet,
		Chassis:    x.Chassis,
		MgmtSwitch: mgmtSwitch,
	}
}

// RouterModule will get a child component with the specified ordinal
func (x Chassis) RouterModule(routerModule int) RouterModule {
	return RouterModule{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: routerModule,
	}
}

// CMMFpga - xXcCfF
type CMMFpga struct {
	Cabinet int
	Chassis int
	CMMFpga int
}

// String will stringify CMMFpga into the format of xXcCfF
func (x CMMFpga) String() string {
	return fmt.Sprintf(
		"x%dc%df%d",
		x.Cabinet,
		x.Chassis,
		x.CMMFpga,
	)
}

// Parent will determine the parent of this CMMFpga
func (x CMMFpga) Parent() Chassis {
	return Chassis{
		Cabinet: x.Cabinet,
		Chassis: x.Chassis,
	}
}

// CMMRectifier - xXcCtT
type CMMRectifier struct {
	Cabinet      int
	Chassis      int
	CMMRectifier int
}

// String will stringify CMMRectifier into the format of xXcCtT
func (x CMMRectifier) String() string {
	return fmt.Sprintf(
		"x%dc%dt%d",
		x.Cabinet,
		x.Chassis,
		x.CMMRectifier,
	)
}

// Parent will determine the parent of this CMMRectifier
func (x CMMRectifier) Parent() Chassis {
	return Chassis{
		Cabinet: x.Cabinet,
		Chassis: x.Chassis,
	}
}

// ChassisBMC - xXcCbB
type ChassisBMC struct {
	Cabinet    int
	Chassis    int
	ChassisBMC int
}

// String will stringify ChassisBMC into the format of xXcCbB
func (x ChassisBMC) String() string {
	return fmt.Sprintf(
		"x%dc%db%d",
		x.Cabinet,
		x.Chassis,
		x.ChassisBMC,
	)
}

// Parent will determine the parent of this ChassisBMC
func (x ChassisBMC) Parent() Chassis {
	return Chassis{
		Cabinet: x.Cabinet,
		Chassis: x.Chassis,
	}
}

// ChassisBMCNic will get a child component with the specified ordinal
func (x ChassisBMC) ChassisBMCNic(chassisBMCNic int) ChassisBMCNic {
	return ChassisBMCNic{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ChassisBMC:    x.ChassisBMC,
		ChassisBMCNic: chassisBMCNic,
	}
}

// ChassisBMCNic - xXcCbBiI
type ChassisBMCNic struct {
	Cabinet       int
	Chassis       int
	ChassisBMC    int
	ChassisBMCNic int
}

// String will stringify ChassisBMCNic into the format of xXcCbBiI
func (x ChassisBMCNic) String() string {
	return fmt.Sprintf(
		"x%dc%db%di%d",
		x.Cabinet,
		x.Chassis,
		x.ChassisBMC,
		x.ChassisBMCNic,
	)
}

// Parent will determine the parent of this ChassisBMCNic
func (x ChassisBMCNic) Parent() ChassisBMC {
	return ChassisBMC{
		Cabinet:    x.Cabinet,
		Chassis:    x.Chassis,
		ChassisBMC: x.ChassisBMC,
	}
}

// ComputeModule - xXcCsS
type ComputeModule struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
}

// String will stringify ComputeModule into the format of xXcCsS
func (x ComputeModule) String() string {
	return fmt.Sprintf(
		"x%dc%ds%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
	)
}

// Parent will determine the parent of this ComputeModule
func (x ComputeModule) Parent() Chassis {
	return Chassis{
		Cabinet: x.Cabinet,
		Chassis: x.Chassis,
	}
}

// NodeBMC will get a child component with the specified ordinal
func (x ComputeModule) NodeBMC(nodeBMC int) NodeBMC {
	return NodeBMC{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       nodeBMC,
	}
}

// NodeEnclosure will get a child component with the specified ordinal
func (x ComputeModule) NodeEnclosure(nodeEnclosure int) NodeEnclosure {
	return NodeEnclosure{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeEnclosure: nodeEnclosure,
	}
}

// NodePowerConnector will get a child component with the specified ordinal
func (x ComputeModule) NodePowerConnector(nodePowerConnector int) NodePowerConnector {
	return NodePowerConnector{
		Cabinet:            x.Cabinet,
		Chassis:            x.Chassis,
		ComputeModule:      x.ComputeModule,
		NodePowerConnector: nodePowerConnector,
	}
}

// NodeBMC - xXcCsSbB
type NodeBMC struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
}

// String will stringify NodeBMC into the format of xXcCsSbB
func (x NodeBMC) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
	)
}

// Parent will determine the parent of this NodeBMC
func (x NodeBMC) Parent() ComputeModule {
	return ComputeModule{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
	}
}

// Node will get a child component with the specified ordinal
func (x NodeBMC) Node(node int) Node {
	return Node{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          node,
	}
}

// NodeBMCNic will get a child component with the specified ordinal
func (x NodeBMC) NodeBMCNic(nodeBMCNic int) NodeBMCNic {
	return NodeBMCNic{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		NodeBMCNic:    nodeBMCNic,
	}
}

// Node - xXcCsSbBnN
type Node struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	Node          int
}

// String will stringify Node into the format of xXcCsSbBnN
func (x Node) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
	)
}

// Parent will determine the parent of this Node
func (x Node) Parent() NodeBMC {
	return NodeBMC{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
	}
}

// Memory will get a child component with the specified ordinal
func (x Node) Memory(memory int) Memory {
	return Memory{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
		Memory:        memory,
	}
}

// NodeAccel will get a child component with the specified ordinal
func (x Node) NodeAccel(nodeAccel int) NodeAccel {
	return NodeAccel{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
		NodeAccel:     nodeAccel,
	}
}

// NodeAccelRiser will get a child component with the specified ordinal
func (x Node) NodeAccelRiser(nodeAccelRiser int) NodeAccelRiser {
	return NodeAccelRiser{
		Cabinet:        x.Cabinet,
		Chassis:        x.Chassis,
		ComputeModule:  x.ComputeModule,
		NodeBMC:        x.NodeBMC,
		Node:           x.Node,
		NodeAccelRiser: nodeAccelRiser,
	}
}

// NodeHsnNic will get a child component with the specified ordinal
func (x Node) NodeHsnNic(nodeHsnNic int) NodeHsnNic {
	return NodeHsnNic{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
		NodeHsnNic:    nodeHsnNic,
	}
}

// NodeNic will get a child component with the specified ordinal
func (x Node) NodeNic(nodeNic int) NodeNic {
	return NodeNic{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
		NodeNic:       nodeNic,
	}
}

// Processor will get a child component with the specified ordinal
func (x Node) Processor(processor int) Processor {
	return Processor{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
		Processor:     processor,
	}
}

// StorageGroup will get a child component with the specified ordinal
func (x Node) StorageGroup(storageGroup int) StorageGroup {
	return StorageGroup{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
		StorageGroup:  storageGroup,
	}
}

// Memory - xXcCsSbBnNdD
type Memory struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	Node          int
	Memory        int
}

// String will stringify Memory into the format of xXcCsSbBnNdD
func (x Memory) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%dd%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
		x.Memory,
	)
}

// Parent will determine the parent of this Memory
func (x Memory) Parent() Node {
	return Node{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
	}
}

// NodeAccel - xXcCsSbBnNaA
type NodeAccel struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	Node          int
	NodeAccel     int
}

// String will stringify NodeAccel into the format of xXcCsSbBnNaA
func (x NodeAccel) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%da%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
		x.NodeAccel,
	)
}

// Parent will determine the parent of this NodeAccel
func (x NodeAccel) Parent() Node {
	return Node{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
	}
}

// NodeAccelRiser - xXcCsSbBnNrR
type NodeAccelRiser struct {
	Cabinet        int
	Chassis        int
	ComputeModule  int
	NodeBMC        int
	Node           int
	NodeAccelRiser int
}

// String will stringify NodeAccelRiser into the format of xXcCsSbBnNrR
func (x NodeAccelRiser) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%dr%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
		x.NodeAccelRiser,
	)
}

// Parent will determine the parent of this NodeAccelRiser
func (x NodeAccelRiser) Parent() Node {
	return Node{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
	}
}

// NodeHsnNic - xXcCsSbBnNhH
type NodeHsnNic struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	Node          int
	NodeHsnNic    int
}

// String will stringify NodeHsnNic into the format of xXcCsSbBnNhH
func (x NodeHsnNic) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%dh%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
		x.NodeHsnNic,
	)
}

// Parent will determine the parent of this NodeHsnNic
func (x NodeHsnNic) Parent() Node {
	return Node{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
	}
}

// NodeNic - xXcCsSbBnNiI
type NodeNic struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	Node          int
	NodeNic       int
}

// String will stringify NodeNic into the format of xXcCsSbBnNiI
func (x NodeNic) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%di%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
		x.NodeNic,
	)
}

// Parent will determine the parent of this NodeNic
func (x NodeNic) Parent() Node {
	return Node{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
	}
}

// Processor - xXcCsSbBnNpP
type Processor struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	Node          int
	Processor     int
}

// String will stringify Processor into the format of xXcCsSbBnNpP
func (x Processor) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%dp%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
		x.Processor,
	)
}

// Parent will determine the parent of this Processor
func (x Processor) Parent() Node {
	return Node{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
	}
}

// StorageGroup - xXcCsSbBnNgG
type StorageGroup struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	Node          int
	StorageGroup  int
}

// String will stringify StorageGroup into the format of xXcCsSbBnNgG
func (x StorageGroup) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%dg%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
		x.StorageGroup,
	)
}

// Parent will determine the parent of this StorageGroup
func (x StorageGroup) Parent() Node {
	return Node{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
	}
}

// Drive will get a child component with the specified ordinal
func (x StorageGroup) Drive(drive int) Drive {
	return Drive{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
		StorageGroup:  x.StorageGroup,
		Drive:         drive,
	}
}

// Drive - xXcCsSbBnNgGkK
type Drive struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	Node          int
	StorageGroup  int
	Drive         int
}

// String will stringify Drive into the format of xXcCsSbBnNgGkK
func (x Drive) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%dn%dg%dk%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.Node,
		x.StorageGroup,
		x.Drive,
	)
}

// Parent will determine the parent of this Drive
func (x Drive) Parent() StorageGroup {
	return StorageGroup{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
		Node:          x.Node,
		StorageGroup:  x.StorageGroup,
	}
}

// NodeBMCNic - xXcCsSbBiI
type NodeBMCNic struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeBMC       int
	NodeBMCNic    int
}

// String will stringify NodeBMCNic into the format of xXcCsSbBiI
func (x NodeBMCNic) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%di%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeBMC,
		x.NodeBMCNic,
	)
}

// Parent will determine the parent of this NodeBMCNic
func (x NodeBMCNic) Parent() NodeBMC {
	return NodeBMC{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeBMC:       x.NodeBMC,
	}
}

// NodeEnclosure - xXcCsSbBeE
type NodeEnclosure struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeEnclosure int
}

// String will stringify NodeEnclosure into the format of xXcCsSbBeE
func (x NodeEnclosure) String() string {
	return fmt.Sprintf(
		"x%dc%ds%de%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeEnclosure,
	)
}

// Parent will determine the parent of this NodeEnclosure
func (x NodeEnclosure) Parent() ComputeModule {
	return ComputeModule{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
	}
}

// NodeEnclosurePowerSupply will get a child component with the specified ordinal
func (x NodeEnclosure) NodeEnclosurePowerSupply(nodeEnclosurePowerSupply int) NodeEnclosurePowerSupply {
	return NodeEnclosurePowerSupply{
		Cabinet:                  x.Cabinet,
		Chassis:                  x.Chassis,
		ComputeModule:            x.ComputeModule,
		NodeEnclosure:            x.NodeEnclosure,
		NodeEnclosurePowerSupply: nodeEnclosurePowerSupply,
	}
}

// NodeFpga will get a child component with the specified ordinal
func (x NodeEnclosure) NodeFpga(nodeFpga int) NodeFpga {
	return NodeFpga{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeEnclosure: x.NodeEnclosure,
		NodeFpga:      nodeFpga,
	}
}

// NodeEnclosurePowerSupply - xXcCsSbBeEtT
type NodeEnclosurePowerSupply struct {
	Cabinet                  int
	Chassis                  int
	ComputeModule            int
	NodeEnclosure            int
	NodeEnclosurePowerSupply int
}

// String will stringify NodeEnclosurePowerSupply into the format of xXcCsSbBeEtT
func (x NodeEnclosurePowerSupply) String() string {
	return fmt.Sprintf(
		"x%dc%ds%de%dt%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeEnclosure,
		x.NodeEnclosurePowerSupply,
	)
}

// Parent will determine the parent of this NodeEnclosurePowerSupply
func (x NodeEnclosurePowerSupply) Parent() NodeEnclosure {
	return NodeEnclosure{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeEnclosure: x.NodeEnclosure,
	}
}

// NodeFpga - xXcCsSbBfF
type NodeFpga struct {
	Cabinet       int
	Chassis       int
	ComputeModule int
	NodeEnclosure int
	NodeFpga      int
}

// String will stringify NodeFpga into the format of xXcCsSbBfF
func (x NodeFpga) String() string {
	return fmt.Sprintf(
		"x%dc%ds%db%df%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodeEnclosure,
		x.NodeFpga,
	)
}

// Parent will determine the parent of this NodeFpga
func (x NodeFpga) Parent() NodeEnclosure {
	return NodeEnclosure{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
		NodeEnclosure: x.NodeEnclosure,
	}
}

// NodePowerConnector - xXcCsSv
type NodePowerConnector struct {
	Cabinet            int
	Chassis            int
	ComputeModule      int
	NodePowerConnector int
}

// String will stringify NodePowerConnector into the format of xXcCsSv
func (x NodePowerConnector) String() string {
	return fmt.Sprintf(
		"x%dc%ds%dv%d",
		x.Cabinet,
		x.Chassis,
		x.ComputeModule,
		x.NodePowerConnector,
	)
}

// Parent will determine the parent of this NodePowerConnector
func (x NodePowerConnector) Parent() ComputeModule {
	return ComputeModule{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		ComputeModule: x.ComputeModule,
	}
}

// MgmtHLSwitchEnclosure - xXcChH
type MgmtHLSwitchEnclosure struct {
	Cabinet               int
	Chassis               int
	MgmtHLSwitchEnclosure int
}

// String will stringify MgmtHLSwitchEnclosure into the format of xXcChH
func (x MgmtHLSwitchEnclosure) String() string {
	return fmt.Sprintf(
		"x%dc%dh%d",
		x.Cabinet,
		x.Chassis,
		x.MgmtHLSwitchEnclosure,
	)
}

// Parent will determine the parent of this MgmtHLSwitchEnclosure
func (x MgmtHLSwitchEnclosure) Parent() Chassis {
	return Chassis{
		Cabinet: x.Cabinet,
		Chassis: x.Chassis,
	}
}

// MgmtHLSwitch will get a child component with the specified ordinal
func (x MgmtHLSwitchEnclosure) MgmtHLSwitch(mgmtHLSwitch int) MgmtHLSwitch {
	return MgmtHLSwitch{
		Cabinet:               x.Cabinet,
		Chassis:               x.Chassis,
		MgmtHLSwitchEnclosure: x.MgmtHLSwitchEnclosure,
		MgmtHLSwitch:          mgmtHLSwitch,
	}
}

// MgmtHLSwitch - xXcChHsS
type MgmtHLSwitch struct {
	Cabinet               int
	Chassis               int
	MgmtHLSwitchEnclosure int
	MgmtHLSwitch          int
}

// String will stringify MgmtHLSwitch into the format of xXcChHsS
func (x MgmtHLSwitch) String() string {
	return fmt.Sprintf(
		"x%dc%dh%ds%d",
		x.Cabinet,
		x.Chassis,
		x.MgmtHLSwitchEnclosure,
		x.MgmtHLSwitch,
	)
}

// Parent will determine the parent of this MgmtHLSwitch
func (x MgmtHLSwitch) Parent() MgmtHLSwitchEnclosure {
	return MgmtHLSwitchEnclosure{
		Cabinet:               x.Cabinet,
		Chassis:               x.Chassis,
		MgmtHLSwitchEnclosure: x.MgmtHLSwitchEnclosure,
	}
}

// MgmtSwitch - xXcCwW
type MgmtSwitch struct {
	Cabinet    int
	Chassis    int
	MgmtSwitch int
}

// String will stringify MgmtSwitch into the format of xXcCwW
func (x MgmtSwitch) String() string {
	return fmt.Sprintf(
		"x%dc%dw%d",
		x.Cabinet,
		x.Chassis,
		x.MgmtSwitch,
	)
}

// Parent will determine the parent of this MgmtSwitch
func (x MgmtSwitch) Parent() Chassis {
	return Chassis{
		Cabinet: x.Cabinet,
		Chassis: x.Chassis,
	}
}

// MgmtSwitchConnector will get a child component with the specified ordinal
func (x MgmtSwitch) MgmtSwitchConnector(mgmtSwitchConnector int) MgmtSwitchConnector {
	return MgmtSwitchConnector{
		Cabinet:             x.Cabinet,
		Chassis:             x.Chassis,
		MgmtSwitch:          x.MgmtSwitch,
		MgmtSwitchConnector: mgmtSwitchConnector,
	}
}

// MgmtSwitchConnector - xXcCwWjJ
type MgmtSwitchConnector struct {
	Cabinet             int
	Chassis             int
	MgmtSwitch          int
	MgmtSwitchConnector int
}

// String will stringify MgmtSwitchConnector into the format of xXcCwWjJ
func (x MgmtSwitchConnector) String() string {
	return fmt.Sprintf(
		"x%dc%dw%dj%d",
		x.Cabinet,
		x.Chassis,
		x.MgmtSwitch,
		x.MgmtSwitchConnector,
	)
}

// Parent will determine the parent of this MgmtSwitchConnector
func (x MgmtSwitchConnector) Parent() MgmtSwitch {
	return MgmtSwitch{
		Cabinet:    x.Cabinet,
		Chassis:    x.Chassis,
		MgmtSwitch: x.MgmtSwitch,
	}
}

// RouterModule - xXcCrR
type RouterModule struct {
	Cabinet      int
	Chassis      int
	RouterModule int
}

// String will stringify RouterModule into the format of xXcCrR
func (x RouterModule) String() string {
	return fmt.Sprintf(
		"x%dc%dr%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
	)
}

// Parent will determine the parent of this RouterModule
func (x RouterModule) Parent() Chassis {
	return Chassis{
		Cabinet: x.Cabinet,
		Chassis: x.Chassis,
	}
}

// HSNAsic will get a child component with the specified ordinal
func (x RouterModule) HSNAsic(hSNAsic int) HSNAsic {
	return HSNAsic{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		HSNAsic:      hSNAsic,
	}
}

// HSNBoard will get a child component with the specified ordinal
func (x RouterModule) HSNBoard(hSNBoard int) HSNBoard {
	return HSNBoard{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		HSNBoard:     hSNBoard,
	}
}

// HSNConnector will get a child component with the specified ordinal
func (x RouterModule) HSNConnector(hSNConnector int) HSNConnector {
	return HSNConnector{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		HSNConnector: hSNConnector,
	}
}

// RouterBMC will get a child component with the specified ordinal
func (x RouterModule) RouterBMC(routerBMC int) RouterBMC {
	return RouterBMC{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		RouterBMC:    routerBMC,
	}
}

// RouterFpga will get a child component with the specified ordinal
func (x RouterModule) RouterFpga(routerFpga int) RouterFpga {
	return RouterFpga{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		RouterFpga:   routerFpga,
	}
}

// RouterPowerConnector will get a child component with the specified ordinal
func (x RouterModule) RouterPowerConnector(routerPowerConnector int) RouterPowerConnector {
	return RouterPowerConnector{
		Cabinet:              x.Cabinet,
		Chassis:              x.Chassis,
		RouterModule:         x.RouterModule,
		RouterPowerConnector: routerPowerConnector,
	}
}

// RouterTOR will get a child component with the specified ordinal
func (x RouterModule) RouterTOR(routerTOR int) RouterTOR {
	return RouterTOR{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		RouterTOR:    routerTOR,
	}
}

// HSNAsic - xXcCrRaA
type HSNAsic struct {
	Cabinet      int
	Chassis      int
	RouterModule int
	HSNAsic      int
}

// String will stringify HSNAsic into the format of xXcCrRaA
func (x HSNAsic) String() string {
	return fmt.Sprintf(
		"x%dc%dr%da%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.HSNAsic,
	)
}

// Parent will determine the parent of this HSNAsic
func (x HSNAsic) Parent() RouterModule {
	return RouterModule{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
	}
}

// HSNLink will get a child component with the specified ordinal
func (x HSNAsic) HSNLink(hSNLink int) HSNLink {
	return HSNLink{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		HSNAsic:      x.HSNAsic,
		HSNLink:      hSNLink,
	}
}

// HSNLink - xXcCrRaAlL
type HSNLink struct {
	Cabinet      int
	Chassis      int
	RouterModule int
	HSNAsic      int
	HSNLink      int
}

// String will stringify HSNLink into the format of xXcCrRaAlL
func (x HSNLink) String() string {
	return fmt.Sprintf(
		"x%dc%dr%da%dl%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.HSNAsic,
		x.HSNLink,
	)
}

// Parent will determine the parent of this HSNLink
func (x HSNLink) Parent() HSNAsic {
	return HSNAsic{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		HSNAsic:      x.HSNAsic,
	}
}

// HSNBoard - xXcCrReE
type HSNBoard struct {
	Cabinet      int
	Chassis      int
	RouterModule int
	HSNBoard     int
}

// String will stringify HSNBoard into the format of xXcCrReE
func (x HSNBoard) String() string {
	return fmt.Sprintf(
		"x%dc%dr%de%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.HSNBoard,
	)
}

// Parent will determine the parent of this HSNBoard
func (x HSNBoard) Parent() RouterModule {
	return RouterModule{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
	}
}

// HSNConnector - xXcCrRjJ
type HSNConnector struct {
	Cabinet      int
	Chassis      int
	RouterModule int
	HSNConnector int
}

// String will stringify HSNConnector into the format of xXcCrRjJ
func (x HSNConnector) String() string {
	return fmt.Sprintf(
		"x%dc%dr%dj%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.HSNConnector,
	)
}

// Parent will determine the parent of this HSNConnector
func (x HSNConnector) Parent() RouterModule {
	return RouterModule{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
	}
}

// HSNConnectorPort will get a child component with the specified ordinal
func (x HSNConnector) HSNConnectorPort(hSNConnectorPort int) HSNConnectorPort {
	return HSNConnectorPort{
		Cabinet:          x.Cabinet,
		Chassis:          x.Chassis,
		RouterModule:     x.RouterModule,
		HSNConnector:     x.HSNConnector,
		HSNConnectorPort: hSNConnectorPort,
	}
}

// HSNConnectorPort - xXcCrRjJpP
type HSNConnectorPort struct {
	Cabinet          int
	Chassis          int
	RouterModule     int
	HSNConnector     int
	HSNConnectorPort int
}

// String will stringify HSNConnectorPort into the format of xXcCrRjJpP
func (x HSNConnectorPort) String() string {
	return fmt.Sprintf(
		"x%dc%dr%dj%dp%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.HSNConnector,
		x.HSNConnectorPort,
	)
}

// Parent will determine the parent of this HSNConnectorPort
func (x HSNConnectorPort) Parent() HSNConnector {
	return HSNConnector{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		HSNConnector: x.HSNConnector,
	}
}

// RouterBMC - xXcCrRbB
type RouterBMC struct {
	Cabinet      int
	Chassis      int
	RouterModule int
	RouterBMC    int
}

// String will stringify RouterBMC into the format of xXcCrRbB
func (x RouterBMC) String() string {
	return fmt.Sprintf(
		"x%dc%dr%db%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.RouterBMC,
	)
}

// Parent will determine the parent of this RouterBMC
func (x RouterBMC) Parent() RouterModule {
	return RouterModule{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
	}
}

// RouterBMCNic will get a child component with the specified ordinal
func (x RouterBMC) RouterBMCNic(routerBMCNic int) RouterBMCNic {
	return RouterBMCNic{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		RouterBMC:    x.RouterBMC,
		RouterBMCNic: routerBMCNic,
	}
}

// RouterBMCNic - xXcCrRbBiI
type RouterBMCNic struct {
	Cabinet      int
	Chassis      int
	RouterModule int
	RouterBMC    int
	RouterBMCNic int
}

// String will stringify RouterBMCNic into the format of xXcCrRbBiI
func (x RouterBMCNic) String() string {
	return fmt.Sprintf(
		"x%dc%dr%db%di%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.RouterBMC,
		x.RouterBMCNic,
	)
}

// Parent will determine the parent of this RouterBMCNic
func (x RouterBMCNic) Parent() RouterBMC {
	return RouterBMC{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		RouterBMC:    x.RouterBMC,
	}
}

// RouterFpga - xXcCrRfF
type RouterFpga struct {
	Cabinet      int
	Chassis      int
	RouterModule int
	RouterFpga   int
}

// String will stringify RouterFpga into the format of xXcCrRfF
func (x RouterFpga) String() string {
	return fmt.Sprintf(
		"x%dc%dr%df%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.RouterFpga,
	)
}

// Parent will determine the parent of this RouterFpga
func (x RouterFpga) Parent() RouterModule {
	return RouterModule{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
	}
}

// RouterPowerConnector - xXcCrRvV
type RouterPowerConnector struct {
	Cabinet              int
	Chassis              int
	RouterModule         int
	RouterPowerConnector int
}

// String will stringify RouterPowerConnector into the format of xXcCrRvV
func (x RouterPowerConnector) String() string {
	return fmt.Sprintf(
		"x%dc%dr%dv%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.RouterPowerConnector,
	)
}

// Parent will determine the parent of this RouterPowerConnector
func (x RouterPowerConnector) Parent() RouterModule {
	return RouterModule{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
	}
}

// RouterTOR - xXcCrRtT
type RouterTOR struct {
	Cabinet      int
	Chassis      int
	RouterModule int
	RouterTOR    int
}

// String will stringify RouterTOR into the format of xXcCrRtT
func (x RouterTOR) String() string {
	return fmt.Sprintf(
		"x%dc%dr%dt%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.RouterTOR,
	)
}

// Parent will determine the parent of this RouterTOR
func (x RouterTOR) Parent() RouterModule {
	return RouterModule{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
	}
}

// RouterTORFpga will get a child component with the specified ordinal
func (x RouterTOR) RouterTORFpga(routerTORFpga int) RouterTORFpga {
	return RouterTORFpga{
		Cabinet:       x.Cabinet,
		Chassis:       x.Chassis,
		RouterModule:  x.RouterModule,
		RouterTOR:     x.RouterTOR,
		RouterTORFpga: routerTORFpga,
	}
}

// RouterTORFpga - xXcCrRtTfF
type RouterTORFpga struct {
	Cabinet       int
	Chassis       int
	RouterModule  int
	RouterTOR     int
	RouterTORFpga int
}

// String will stringify RouterTORFpga into the format of xXcCrRtTfF
func (x RouterTORFpga) String() string {
	return fmt.Sprintf(
		"x%dc%dr%dt%df%d",
		x.Cabinet,
		x.Chassis,
		x.RouterModule,
		x.RouterTOR,
		x.RouterTORFpga,
	)
}

// Parent will determine the parent of this RouterTORFpga
func (x RouterTORFpga) Parent() RouterTOR {
	return RouterTOR{
		Cabinet:      x.Cabinet,
		Chassis:      x.Chassis,
		RouterModule: x.RouterModule,
		RouterTOR:    x.RouterTOR,
	}
}
