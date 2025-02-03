package services

// NOTE: In GO, all types, their fields and the functions whose names start with capital are automatically exported
// 		Hence, any Service (type) name and its methods that needs to be invocable by RPC,
// 		any Args types and their fields and any Return types and their fields that are to be part of a RPC call
// 		need to have names starting with caps.

type Args struct {
	A, B int
}
