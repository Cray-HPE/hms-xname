// MIT License
//
// (C) Copyright [2021-2022] Hewlett Packard Enterprise Development LP
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

package xnames

import (
	"errors"
	"strconv"

	"github.com/Cray-HPE/hms-xname/xnametypes"
)

var ErrUnknownStruct = errors.New("unable to determine HMS Type from struct")

// GetHMSType for a given xname structure will return its HMSType
// If the given object is not a structure from the xnames package, 
// then the ErrUnknownStruct will be returned along with HMSTypeInvalid
func GetHMSType(obj interface{}) (xnametypes.HMSType, error) {
    switch obj.(type) {
{{ range $xnameType := . }}
	case {{$xnameType.Entry.Type}}, *{{$xnameType.Entry.Type}}:
		return xnametypes.{{$xnameType.Entry.Type}}, nil
{{- end }}
    }
    return xnametypes.HMSTypeInvalid, ErrUnknownStruct
}

// FromString will convert the string representation of a xname into a xname structure
// If the string is not a valid xname, then nil and HMSTypeInvalid will be returned. 
func FromString(xname string) GenericXname {
	hmsType := xnametypes.GetHMSType(xname)
	if hmsType == xnametypes.HMSTypeInvalid {
		return nil
	}

	re, err := xnametypes.GetHMSTypeRegex(hmsType)
	if err != nil {
		return nil
	}

	_, argCount, err := xnametypes.GetHMSTypeFormatString(hmsType)
	if err != nil {
		return nil
	}

	matchesRaw := re.FindStringSubmatch(xname)
	if (argCount + 1) != len(matchesRaw) {
		return nil
	}

	// If we have gotten to this point these matches should be integers, so we can safely convert them
	// to integers from strings.
	matches := []int{}
	for _, matchRaw := range matchesRaw[1:] {
		match, err := strconv.Atoi(matchRaw)
		if err != nil {
			return nil
		}

		matches = append(matches, match)
	}

	var component GenericXname

	switch hmsType {
{{- range $xnameType := . }}
	case xnametypes.{{$xnameType.Entry.Type}}:
		component = {{$xnameType.Entry.Type}}{
			{{- range $i, $field := $xnameType.Fields }}
			{{ $field }}: matches[{{$i}}],
			{{- end }}
		}
{{- end }}
	default:
		return nil
	}
	return component
}
