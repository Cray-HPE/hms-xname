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
	"fmt"

	"github.com/Cray-HPE/hms-xname/xnametypes"
)

{{ range $xnameType := . -}}
// {{ $xnameType.Entry.Type }} - {{ $xnameType.Entry.ExampleString }}
type {{ $xnameType.Entry.Type }} struct {
{{- range $i, $field := .Fields }}
	{{ $field }} int // {{ index $xnameType.FieldPlaceHolders $i }}
{{- end }}
}

// Type will return the corresponding HMSType
func (x {{ $xnameType.Entry.Type }}) Type() xnametypes.HMSType  {
	return xnametypes.{{ $xnameType.Entry.Type }}
}

// String will stringify {{ $xnameType.Entry.Type }} into the format of {{ $xnameType.Entry.ExampleString }}
func (x {{ $xnameType.Entry.Type }}) String() string {
	return fmt.Sprintf(
		"{{ $xnameType.Entry.GenStr }}",
		{{- range $field := .Fields }}
		x.{{$field}},
		{{- end }}
	)
}

{{ if $xnameType.Parent -}}
// Parent will determine the parent of this {{ $xnameType.Entry.Type }}
func (x {{ $xnameType.Entry.Type }}) Parent() {{ $xnameType.Entry.ParentType }} {
	return {{ $xnameType.Entry.ParentType }}{
		{{- range $field := $xnameType.Parent.Fields }}
		{{ $field }}: x.{{ $field }},
		{{- end }}
	}
}
{{- end }}

// ParentGeneric will determine the parent of this {{ $xnameType.Entry.Type }}, and return it as a Xname interface
func (x {{ $xnameType.Entry.Type }}) ParentInterface() Xname {
	{{ if $xnameType.Parent -}}
	return x.Parent()
	{{- end }}
	{{ if not $xnameType.Parent -}}
	return nil
	{{- end }}
}

{{ range $child := $xnameType.Children }}

// {{ $child.Entry.Type }} will get a child component with the specified ordinal
func (x {{ $xnameType.Entry.Type }}) {{ $child.Entry.Type }}({{ $child.FunctionParameter }} int) {{ $child.Entry.Type }} {
	return {{ $child.Entry.Type }}{
		{{- range $field := $xnameType.Fields }}
		{{ $field }}: x.{{ $field }},
		{{- end }}
		{{ $child.Entry.Type }}: {{ $child.FunctionParameter }},
	}
}
{{ end -}}

// Validate will validate the string representation of this structure against xnametypes.IsHMSCompIDValid() 
func (x {{ $xnameType.Entry.Type }}) Validate() error {
	xname := x.String()
	if !xnametypes.IsHMSCompIDValid(xname) {
		return fmt.Errorf("invalid {{ $xnameType.Entry.Type }} xname: %s", xname)
	}

	return nil
}

func (x {{ $xnameType.Entry.Type }}) GetCabinet() (int, bool) {

}


func (x {{ $xnameType.Entry.Type }}) GetComputeSlot() (int, bool) {

}

{{ end -}}
