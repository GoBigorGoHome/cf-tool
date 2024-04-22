package client

// Langs generated by
// ^[\s\S]*?value="(.+?)"[\s\S]*?>([\s\S]+?)<[\s\S]*?$
//     "\1": "\2",

// Regular expressions:
// Groups are marked by the '(', ')' metacharacters.

var Langs = map[string]string{
	"43": "GNU GCC C11 5.1.0",
	"50": "GNU G++14 6.4.0",
	"54": "GNU G++17 7.3.0",
	"89": "GNU G++20 13.2 (64 bit, winlibs)",
	"65": "C# 8, .NET Core 3.1",
	"79": "C# 10, .NET SDK 6.0",
	"9": "C# Mono 6.8",
	"28": "D DMD32 v2.105.0",
	"32": "Go 1.22.2",
	"12": "Haskell GHC 8.10.1",
	"87": "Java 21 64bit",
	"36": "Java 8 32bit",
	"83": "Kotlin 1.7.20",
	"88": "Kotlin 1.9.21",
	"19": "OCaml 4.02.1",
	"3": "Delphi 7",
	"4": "Free Pascal 3.2.2",
	"51": "PascalABC.NET 3.8.3",
	"13": "Perl 5.20.1",
	"6": "PHP 8.1.7",
	"7": "Python 2.7.18",
	"31": "Python 3.8.10",
	"40": "PyPy 2.7.13 (7.3.0)",
	"41": "PyPy 3.6.9 (7.3.0)",
	"70": "PyPy 3.10 (7.3.15, 64bit)",
	"67": "Ruby 3.2.2",
	"75": "Rust 1.75.0 (2021)",
	"20": "Scala 2.12.8",
	"34": "JavaScript V8 4.8.0",
	"55": "Node.js 15.8.0 (64bit)",
}

/* todo: 
 1. 弄清楚LangsExt是干什么用的。
 2. 更新LangsExt。
*/

// LangsExt language's ext
var LangsExt = map[string]string{
	"GNU C11":               "c",
	"GNU C++20 (64)":        "cpp",
	"GNU C++17 (64)":        "cpp",
	"Clang++17 Diagnostics": "cpp",
	"GNU C++0x":             "cpp",
	"GNU C++":               "cpp",
	"GNU C++11":             "cpp",
	"GNU C++14":             "cpp",
	"GNU C++17":             "cpp",
	"MS C++":                "cpp",
	"MS C++ 2017":           "cpp",
	"Mono C#":               "cs",
	"D":                     "d",
	"Go":                    "go",
	"Haskell":               "hs",
	"Kotlin":                "kt",
	"Ocaml":                 "ml",
	"Delphi":                "pas",
	"FPC":                   "pas",
	"PascalABC.NET":         "pas",
	"Perl":                  "pl",
	"PHP":                   "php",
	"Python 2":              "py",
	"Python 3":              "py",
	"PyPy 2":                "py",
	"PyPy 3":                "py",
	"Ruby":                  "rb",
	"Rust":                  "rs",
	"JavaScript":            "js",
	"Node.js":               "js",
	"Q#":                    "qs",
	"Java":                  "java",
	"Java 6":                "java",
	"Java 7":                "java",
	"Java 8":                "java",
	"Java 9":                "java",
	"Java 10":               "java",
	"Java 11":               "java",
	"Tcl":                   "tcl",
	"F#":                    "fs",
	"Befunge":               "bf",
	"Pike":                  "pike",
	"Io":                    "io",
	"Factor":                "factor",
	"Cobol":                 "cbl",
	"Secret_171":            "secret_171",
	"Ada":                   "adb",
	"FALSE":                 "f",
	"":                      "txt",
}
