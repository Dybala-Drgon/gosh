package object

type CallableFunc func(args ...Object) (ret []Object, err error)
