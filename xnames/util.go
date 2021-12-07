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
	"errors"
	"strconv"

	"github.com/Cray-HPE/hms-xname/xnametypes"
)

var ErrUnknownStruct = errors.New("unable to determine HMS Type from struct")

// GetHMSType for a given xname structure will return its HMSType
func GetHMSType(obj interface{}) (xnametypes.HMSType, error) {
	// Handy bash fragment to generate the type switch below
	// for hms_type in $(cat ./xname/types.go | grep '^type' | awk '{print $2}'); do
	// echo "	case $hms_type, *$hms_type:"
	// echo "		return xnametypes.$hms_type, nil"
	// done
	switch obj.(type) {
	case System, *System:
		return xnametypes.System, nil
	case CDU, *CDU:
		return xnametypes.CDU, nil
	case CDUMgmtSwitch, *CDUMgmtSwitch:
		return xnametypes.CDUMgmtSwitch, nil
	case Cabinet, *Cabinet:
		return xnametypes.Cabinet, nil
	case CabinetPDUController, *CabinetPDUController:
		return xnametypes.CabinetPDUController, nil
	case Chassis, *Chassis:
		return xnametypes.Chassis, nil
	case ChassisBMC, *ChassisBMC:
		return xnametypes.ChassisBMC, nil
	case MgmtSwitch, *MgmtSwitch:
		return xnametypes.MgmtSwitch, nil
	case MgmtSwitchConnector, *MgmtSwitchConnector:
		return xnametypes.MgmtSwitchConnector, nil
	case MgmtHLSwitchEnclosure, *MgmtHLSwitchEnclosure:
		return xnametypes.MgmtHLSwitchEnclosure, nil
	case MgmtHLSwitch, *MgmtHLSwitch:
		return xnametypes.MgmtHLSwitch, nil
	case RouterModule, *RouterModule:
		return xnametypes.RouterModule, nil
	case RouterBMC, *RouterBMC:
		return xnametypes.RouterBMC, nil
	case ComputeModule, *ComputeModule:
		return xnametypes.ComputeModule, nil
	case NodeBMC, *NodeBMC:
		return xnametypes.NodeBMC, nil
	case Node, *Node:
		return xnametypes.Node, nil
	}

	return xnametypes.HMSTypeInvalid, ErrUnknownStruct
}

// FromString will convert the string representation of a xname into a xname structure
func FromString(xname string) (interface{}, xnametypes.HMSType) {
	hmsType := xnametypes.GetHMSType(xname)
	if hmsType == xnametypes.HMSTypeInvalid {
		return nil, hmsType
	}

	re, err := xnametypes.GetHMSTypeRegex(hmsType)
	if err != nil {
		return nil, xnametypes.HMSTypeInvalid
	}

	_, argCount, err := xnametypes.GetHMSTypeFormatString(hmsType)
	if err != nil {
		return nil, xnametypes.HMSTypeInvalid
	}

	matchesRaw := re.FindStringSubmatch(xname)
	if (argCount + 1) != len(matchesRaw) {
		return nil, xnametypes.HMSTypeInvalid
	}

	// If we have gotten to this point these matches should be integers, so we can safely convert them
	// to integers from strings.
	matches := []int{}
	for _, matchRaw := range matchesRaw[1:] {
		match, err := strconv.Atoi(matchRaw)
		if err != nil {
			return nil, xnametypes.HMSTypeInvalid
		}

		matches = append(matches, match)
	}

	var component interface{}

	switch hmsType {
	case xnametypes.System:
		component = System{}
	case xnametypes.CDU:
		component = CDU{
			CoolingGroup: matches[0],
		}
	case xnametypes.CDUMgmtSwitch:
		component = CDUMgmtSwitch{
			CoolingGroup: matches[0],
			Slot:         matches[1],
		}
	case xnametypes.Cabinet:
		component = Cabinet{
			Cabinet: matches[0],
		}
	case xnametypes.CabinetPDUController:
		component = CabinetPDUController{
			Cabinet:       matches[0],
			PDUController: matches[1],
		}
	case xnametypes.Chassis:
		component = Chassis{
			Cabinet: matches[0],
			Chassis: matches[1],
		}
	case xnametypes.ChassisBMC:
		component = ChassisBMC{
			Cabinet: matches[0],
			Chassis: matches[1],
			BMC:     matches[2],
		}
	case xnametypes.MgmtSwitch:
		component = MgmtSwitch{
			Cabinet: matches[0],
			Chassis: matches[1],
			Slot:    matches[2],
		}
	case xnametypes.MgmtSwitchConnector:
		component = MgmtSwitchConnector{
			Cabinet:    matches[0],
			Chassis:    matches[1],
			Slot:       matches[2],
			SwitchPort: matches[3],
		}
	case xnametypes.MgmtHLSwitchEnclosure:
		component = MgmtHLSwitchEnclosure{
			Cabinet: matches[0],
			Chassis: matches[1],
			Slot:    matches[2],
		}
	case xnametypes.MgmtHLSwitch:
		component = MgmtHLSwitch{
			Cabinet: matches[0],
			Chassis: matches[1],
			Slot:    matches[2],
			Space:   matches[3],
		}
	case xnametypes.RouterModule:
		component = RouterModule{
			Cabinet: matches[0],
			Chassis: matches[1],
			Slot:    matches[2],
		}
	case xnametypes.RouterBMC:
		component = RouterBMC{
			Cabinet: matches[0],
			Chassis: matches[1],
			Slot:    matches[2],
			BMC:     matches[3],
		}
	case xnametypes.ComputeModule:
		component = ComputeModule{
			Cabinet: matches[0],
			Chassis: matches[1],
			Slot:    matches[2],
		}
	case xnametypes.NodeBMC:
		component = NodeBMC{
			Cabinet: matches[0],
			Chassis: matches[1],
			Slot:    matches[2],
			BMC:     matches[3],
		}
	case xnametypes.Node:
		component = Node{
			Cabinet: matches[0],
			Chassis: matches[1],
			Slot:    matches[2],
			BMC:     matches[3],
			Node:    matches[4],
		}
	default:
		return nil, xnametypes.HMSTypeInvalid
	}
	return component, hmsType
}
