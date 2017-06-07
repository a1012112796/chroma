package lexers

import (
    . "github.com/alecthomas/chroma" // nolint
)

// Thrift lexer.
var Thrift = Register(MustNewLexer(
    &Config{
        Name:      "Thrift",
        Aliases:   []string{"thrift"},
        Filenames: []string{"*.thrift"},
        MimeTypes: []string{"application/x-thrift"},
    },
    Rules{
        "root": {
            Include("whitespace"),
            Include("comments"),
            {`"`, LiteralStringDouble, Combined("stringescape", "dqs")},
            {`\'`, LiteralStringSingle, Combined("stringescape", "sqs")},
            {`(namespace)(\s+)`, ByGroups(KeywordNamespace, TextWhitespace), Push("namespace")},
            {`(enum|union|struct|service|exception)(\s+)`, ByGroups(KeywordDeclaration, TextWhitespace), Push("class")},
            {`((?:(?:[^\W\d]|\$)[\w.\[\]$<>]*\s+)+?)((?:[^\W\d]|\$)[\w$]*)(\s*)(\()`, ByGroups(UsingSelf("root"), NameFunction, Text, Operator), nil},
            Include("keywords"),
            Include("numbers"),
            {`[&=]`, Operator, nil},
            {`[:;,{}()<>\[\]]`, Punctuation, nil},
            {`[a-zA-Z_](\.\w|\w)*`, Name, nil},
        },
        "whitespace": {
            {`\n`, TextWhitespace, nil},
            {`\s+`, TextWhitespace, nil},
        },
        "comments": {
            {`#.*$`, Comment, nil},
            {`//.*?\n`, Comment, nil},
            {`/\*[\w\W]*?\*/`, CommentMultiline, nil},
        },
        "stringescape": {
            {`\\([\\nrt"\'])`, LiteralStringEscape, nil},
        },
        "dqs": {
            {`"`, LiteralStringDouble, Pop(1)},
            {`[^\\"\n]+`, LiteralStringDouble, nil},
        },
        "sqs": {
            {`'`, LiteralStringSingle, Pop(1)},
            {`[^\\\'\n]+`, LiteralStringSingle, nil},
        },
        "namespace": {
            {`[a-z*](\.\w|\w)*`, NameNamespace, Pop(1)},
            Default(Pop(1)),
        },
        "class": {
            {`[a-zA-Z_]\w*`, NameClass, Pop(1)},
            Default(Pop(1)),
        },
        "keywords": {
            {`(async|oneway|extends|throws|required|optional)\b`, Keyword, nil},
            {`(true|false)\b`, KeywordConstant, nil},
            {`(const|typedef)\b`, KeywordDeclaration, nil},
            {`(?:cpp_namespace|cpp_include|cpp_type|java_package|cocoa_prefix|csharp_namespace|delphi_namespace|php_namespace|py_module|perl_package|ruby_namespace|smalltalk_category|smalltalk_prefix|xsd_all|xsd_optional|xsd_nillable|xsd_namespace|xsd_attrs|include)\b`, KeywordNamespace, nil},
            {`(?:void|bool|byte|i16|i32|i64|double|string|binary|map|list|set|slist|senum)\b`, KeywordType, nil},
            {`\b(?:BEGIN|END|__CLASS__|__DIR__|__FILE__|__FUNCTION__|__LINE__|__METHOD__|__NAMESPACE__|abstract|alias|and|args|as|assert|begin|break|case|catch|class|clone|continue|declare|def|default|del|delete|do|dynamic|elif|else|elseif|elsif|end|enddeclare|endfor|endforeach|endif|endswitch|endwhile|ensure|except|exec|finally|float|for|foreach|function|global|goto|if|implements|import|in|inline|instanceof|interface|is|lambda|module|native|new|next|nil|not|or|pass|public|print|private|protected|raise|redo|rescue|retry|register|return|self|sizeof|static|super|switch|synchronized|then|this|throw|transient|try|undef|unless|unsigned|until|use|var|virtual|volatile|when|while|with|xor|yield)\b`, KeywordReserved, nil},
        },
        "numbers": {
            {`[+-]?(\d+\.\d+([eE][+-]?\d+)?|\.?\d+[eE][+-]?\d+)`, LiteralNumberFloat, nil},
            {`[+-]?0x[0-9A-Fa-f]+`, LiteralNumberHex, nil},
            {`[+-]?[0-9]+`, LiteralNumberInteger, nil},
        },
    },
))
