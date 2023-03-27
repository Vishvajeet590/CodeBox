package utils

import "strings"

const CodeDIR = "./judge/code"

const CPP = "CPP"
const JAVA = "JAVA"
const PYTHON = "PYTHON"
const INVALID = "invalid"

func GetLanguageCode(lang string) string {
	lang = strings.ToLower(lang)
	switch lang {
	case "java":
		return JAVA
	case "cpp":
		return CPP
	case "python":
		return PYTHON
	default:
		return INVALID
	}
}

const CppExt = ".cpp"
const JavaExt = ".java"
const PythonExt = ".py"

func GetExtension(languageCode string) string {
	switch languageCode {
	case CPP:
		return CppExt
	case JAVA:
		return JavaExt
	case PYTHON:
		return PythonExt
	default:
		return INVALID
	}
}

//Compilation

const JAVA_COMPILE_CMD = "javac"
const JAVA_CLASS_NAME = "Codebox.class"
const JAVA_RUN_CMD = "java"
const JAVA_RUN_FILE = "Codebox"

const PYTHON_RUN_CMD = "python"

// Status Codes
const RESULT_NIL = "NIL"
const RESULT_COMPILE_ERROR = "COMPILE_ERROR"
const RESULT_RUNTIME_ERROR = "RUNTIME_ERROR"
const RESULT_PASS = "PASS"
const RESULT_FAIL = "FAIL"
