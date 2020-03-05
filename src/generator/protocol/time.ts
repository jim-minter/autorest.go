/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { comment } from '@azure-tools/codegen'
import { CodeModel } from '@azure-tools/codemodel';

// represents the generated content for an operation group
export class Content {
  readonly name: string;
  readonly content: string;

  constructor(name: string, content: string) {
    this.name = name;
    this.content = content;
  }
}

// Creates the content for required time marshalling helpers.
// Will be empty if no helpers are required.
export async function generateTimeHelpers(session: Session<CodeModel>): Promise<Content[]> {
  const headerText = comment(await session.getValue("header-text", "MISSING LICENSE HEADER"), "// ");
  const namespace = await session.getValue('namespace');
  const content = new Array<Content>();
  if (session.model.language.go!.hasTimeRFC1123) {
    content.push(new Content('time_rfc1123', generateRFC1123Helper(headerText, <string>namespace)));
  }
  if (session.model.language.go!.hasTimeRFC3339) {
    content.push(new Content('time_rfc3339', generateRFC3339Helper(headerText, <string>namespace)));
  }
  return content;
}

function generateRFC1123Helper(header: string, packageName: string): string {
  return `${header}

package ${packageName}

import (
	"strings"
	"time"
)

const (
	rfc1123JSON = \`"\` + time.RFC1123 + \`"\`
)

type timeRFC1123 time.Time

func (t timeRFC1123) MarshalJSON() ([]byte, error) {
	b := []byte(time.Time(t).Format(rfc1123JSON))
	return b, nil
}

func (t timeRFC1123) MarshalText() ([]byte, error) {
	b := []byte(time.Time(t).Format(time.RFC1123))
	return b, nil
}

func (t *timeRFC1123) UnmarshalJSON(data []byte) error {
	p, err := time.Parse(rfc1123JSON, strings.ToUpper(string(data)))
	*t = timeRFC1123(p)
	return err
}

func (t *timeRFC1123) UnmarshalText(data []byte) error {
	p, err := time.Parse(time.RFC1123, string(data))
	*t = timeRFC1123(p)
	return err
}

func (t *timeRFC1123) ToTime() *time.Time {
	return (*time.Time)(t)
}
`;
}

function generateRFC3339Helper(header: string, packageName: string): string {
  return `${header}

package ${packageName}

import (
	"regexp"
	"strings"
	"time"
)

const (
	utcLayoutJSON = \`"2006-01-02T15:04:05.999999999"\`
	utcLayout     = "2006-01-02T15:04:05.999999999"
	rfc3339JSON   = \`"\` + time.RFC3339Nano + \`"\`
)

// Azure reports time in UTC but it doesn't include the 'Z' time zone suffix in some cases.
var tzOffsetRegex = regexp.MustCompile(\`(Z|z|\\+|-)(\\d+:\\d+)*"*$\`)

type timeRFC3339 time.Time

func (t *timeRFC3339) UnmarshalJSON(data []byte) error {
	layout := utcLayoutJSON
	if tzOffsetRegex.Match(data) {
		layout = rfc3339JSON
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) UnmarshalText(data []byte) (err error) {
	layout := utcLayout
	if tzOffsetRegex.Match(data) {
		layout = time.RFC3339Nano
	}
	return t.Parse(layout, string(data))
}

func (t *timeRFC3339) ToTime() *time.Time {
	return (*time.Time)(t)
}

func (t *timeRFC3339) Parse(layout, value string) error {
	p, err := time.Parse(layout, strings.ToUpper(value))
	*t = timeRFC3339(p)
	return err
}
`;
}
