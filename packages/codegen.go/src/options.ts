/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { capitalize, comment } from '@azure-tools/codegen';
import { values } from '@azure-tools/linq';
import * as go from '../../codemodel.go/src/gocodemodel.js';
import { commentLength, contentPreamble } from './helpers.js';
import { ImportManager } from './imports.js';

// Creates the content in options.go
export async function generateOptions(codeModel: go.CodeModel): Promise<string> {
  if (codeModel.paramGroups.length === 0) {
    return '';
  }

  const imports = new ImportManager();
  let optionsText = contentPreamble(codeModel);
  let content = '';

  for (const paramGroup of values(codeModel.paramGroups)) {
    content += emit(paramGroup, imports);
  }

  optionsText += imports.text();
  optionsText += content;
  return optionsText;
}

function emit(struct: go.StructType, imports: ImportManager): string {
  let text = '';
  if (struct.description) {
    text += `${comment(struct.description, '// ', undefined, commentLength)}\n`;
  }
  text += `type ${struct.name} struct {\n`;

  if (struct.fields.length === 0) {
    // this is an optional params placeholder struct
    text += '\t// placeholder for future optional parameters\n';
  } else {
    // used to track when to add an extra \n between fields that have comments
    let first = true;

    for (const field of values(struct.fields)) {
      imports.addImportForType(field.type);
      if (field.description) {
        if (!first) {
          // add an extra new-line between fields IFF the field
          // has a comment and it's not the very first one.
          text += '\n';
        }
        text += `\t${comment(field.description, '// ', undefined, commentLength)}\n`;
      }

      let typeName = go.getTypeDeclaration(field.type);
      if (go.isLiteralValue(field.type)) {
        // for constants we use the underlying type name
        typeName = go.getLiteralValueTypeName(field.type.type);
      }

      let pointer = '*';
      if (field.byValue) {
        pointer = '';
      }
      text += `\t${capitalize(field.fieldName)} ${pointer}${typeName}\n`;
      first = false;
    }
  }

  text += '}\n\n';
  return text;
}
