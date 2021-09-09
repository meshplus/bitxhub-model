package pb

import "fmt"

func Int32(i int32) *Arg {
	return generateArg(Arg_I32, []byte(fmt.Sprintf("%d", i)))
}

func Int64(i int64) *Arg {
	return generateArg(Arg_I64, []byte(fmt.Sprintf("%d", i)))
}

func Uint32(i uint32) *Arg {
	return generateArg(Arg_U32, []byte(fmt.Sprintf("%d", i)))
}

func Uint64(i uint64) *Arg {
	return generateArg(Arg_U64, []byte(fmt.Sprintf("%d", i)))
}

func Float64(i float64) *Arg {
	return generateArg(Arg_F64, []byte(fmt.Sprintf("%f", i)))
}

func String(content string) *Arg {
	return generateArg(Arg_String, []byte(content))
}

func Bytes(content []byte) *Arg {
	return generateArg(Arg_Bytes, content)
}

func generateArg(typ Arg_Type, content []byte) *Arg {
	return &Arg{
		Type:  typ,
		Value: content,
	}
}
