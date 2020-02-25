/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

export const CommonAcronyms = [
  'Acl',
  'Api',
  'Ascii',
  'Cpu',
  'Css',
  'Csv',
  'Dns',
  'Eof',
  'Guid',
  'Html',
  'Http',
  'Https',
  'Id',
  'Ip',
  'Json',
  'Lhs',
  'Mbcs',
  'Qps',
  'Ram',
  'Rfc',
  'Rhs',
  'Rpc',
  'Sla',
  'Smtp',
  'Sql',
  'Ssh',
  'Tcp',
  'Tls',
  'Ttl',
  'Udp',
  'Ui',
  'Uid',
  'Uuid',
  'Uri',
  'Url',
  'Utf8',
  'Vm',
  'Xml',
  'Xsrf',
  'Xss'
]

export const ReservedWords = [
  // Reserved keywords -- list retrieved from http://golang.org/ref/spec#Keywords
  'break', 'default', 'func', 'interface', 'select',
  'case', 'defer', 'go', 'map', 'struct',
  'chan', 'else', 'goto', 'package', 'switch',
  'const', 'fallthrough', 'if', 'range', 'type',
  'continue', 'for', 'import', 'return', 'var',

  // Reserved predeclared identifiers -- list retrieved from http://golang.org/ref/spec#Predeclared_identifiers
  'bool', 'byte',
  'complex64', 'complex128',
  'error',
  'float32', 'float64',
  'int', 'int8', 'int16', 'int32', 'int64',
  'rune', 'string',
  'uint', 'uint8', 'uint16', 'uint32', 'uint64',
  'uintptr',

  'true', 'false', 'iota',

  'nil',

  'append', 'cap', 'close', 'complex', 'copy', 'delete', 'imag', 'len', 'make', 'new', 'panic', 'print', 'println', 'real', 'recover',


  // Reserved packages -- list retrieved from http://golang.org/pkg/
  // -- Since package names serve as partial identifiers, exclude the standard library
  'archive', 'tar', 'zip',
  'bufio',
  'builtin',
  'bytes',
  'compress', 'bzip2', 'flate', 'gzip', 'lzw', 'zlib',
  'container', 'heap', 'list', 'ring',
  'context',
  'crypto', 'aes', 'cipher', 'des', 'dsa', 'ecdsa', 'ed25519', 'elliptic', 'hmac', 'md5', 'rand', 'rc4', 'rsa', 'sha1', 'sha256', 'sha512', 'subtle', 'tls', 'x509', 'pkix',
  'database', 'sql', 'driver',
  'debug', 'dwarf', 'elf', 'gosym', 'macho', 'pe', 'plan9obj',
  'encoding', 'ascii85', 'asn1', 'base32', 'base64', 'binary', 'csv', 'gob', 'hex', 'json', 'pem', 'xml',
  'errors',
  'expvar',
  'flag',
  'fmt',
  'go', 'ast', 'build', 'constant', 'doc', 'format', 'importer', 'parser', 'printer', 'scanner', 'token', 'types',
  'hash', 'adler32', 'crc32', 'crc64', 'fnv',
  'html', 'template',
  'image', 'color', 'palette', 'draw', 'gif', 'jpeg', 'png',
  'index', 'suffixarray',
  'io', 'ioutil',
  'log', 'syslog',
  'math', 'big', 'bits', 'cmplx', 'rand',
  'mime', 'multipart', 'quotedprintable',
  'net', 'http', 'cgi', 'cookiejar', 'fcgi', 'httptest', 'httptrace', 'httputil', 'pprof', 'mail', 'rpc', 'jsonrpc', 'smtp', 'textproto', 'url',
  'os', 'exec', 'signal', 'user',
  'path', 'filepath',
  'plugin',
  'reflect',
  'regexp', 'syntax',
  'runtime', 'cgo', 'debug', 'msan', 'pprof', 'race', 'trace',
  'sort',
  'strconv',
  'strings',
  'sync', 'atomic',
  'syscall', 'js',
  'testing', 'iotest', 'quick',
  'text', 'scanner', 'tabwriter', 'template', 'parse',
  'time',
  'unicode', 'utf16', 'utf8',
  'unsafe'
]

