package service

import "github.com/Duvewo/banquend/handler"

type Service interface {
	REGISTER(handler.Handler)
}

func REGISTER(h handler.Handler) {

}
