# Namespace Configuration Language Specification


## Notation

As in the [Go Programming Language Specification](https://go.dev/ref/spec), the syntax is specified using Extended Backus-Naur Form (EBNF):

```ebnf
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions are expressions constructed from terms and the following operators, in increasing precedence:

```ebnf
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

Lower-case production names are used to identify lexical tokens. Non-terminals are in CamelCase. Lexical tokens are enclosed in double quotes "" or back quotes ``.

The form `a … b` represents the set of characters from a through b as alternatives. The horizontal ellipsis `…` is also used elsewhere in the spec to informally denote various enumerations or code snippets that are not further specified.

## Configuraton text representation

The configuration is encoded in UTF-8.

## Lexical elements

### Comments

1. Line comments start with the character sequence // and stop at the end of the line.
2. General comments start with the character sequence /* and stop with the first subsequent character sequence */.

### Identifiers

Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

```ebnf
identifier = letter { letter | unicode_digit } .
```

### Keywords

The configuration language has the following keywords:

* `type`
* `this`

### Operators

The following character sequences represent operators:

```
&&		||		==		!=		(		)		{		}		[		]
=>		.		:
```

## Statements

```ebnf
TypeDefinition  = "type" identifier "{" [ TypeSpec ] "}"
TypeSpec        = { RelationDecl | PermissionDecl }

RelationDecl    = identifier ":" identifier [ "[]" ]

PermissionDecl  = Signature "=>" PermissionBody

Signature       = PermissionName "(" Param ")"
PermissionName  = identifier
Param           = identifier

PermissionBody  = PermissionCheck | { Operator PermissionCheck }
Operator        = "||" | "&&"
PermissionCheck = FunctionCall
FunctionCall    = Callable "(" ( Arguments | LambdaExpr ) ")"
Arguments       = identifier { "," identifier }
LambdaExpr      = identifier "=>" PermissionCheck
Callable        = FunctionCall | Property | BuiltIn
Property        = "this" "." identifier { "." identifier }
BuiltIn         = "any" | "all"
```
