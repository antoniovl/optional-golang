package optional

type Optional struct {
	value *interface{}
}


func (opt *Optional) Get() interface{} {
	return opt.value
}

func (opt *Optional) IsPresent() bool {
	return opt.value != nil
}

func Empty() Optional {
	return Optional{value: nil}
}

func Of(v interface{}) Optional {
	if v == nil {
		return Empty()
	}
	return Optional{v}
}
