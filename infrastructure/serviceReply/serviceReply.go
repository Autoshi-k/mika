package serviceReply

type Reply interface {
}

type Response struct {
	status string
	data   interface{}
	errId  string
}

func NewErr() Reply { // TODO
	return nil
}

func NewInternalError(err error) Reply {
	return &Response{
		status: "failure",
		errId:  "internalServiceErr",
		data:   err.Error(),
	}
}

func NewBadRequest(errId string, err error) Reply {
	return &Response{
		status: "failure",
		errId:  errId,
		data:   err.Error(),
	}
}

func NewDbError(err error) Reply {
	return &Response{
		status: "failure",
		errId:  "dbError",
		data:   err.Error(),
	}
}

func (r Response) Error() string {
	return r.errId
}
