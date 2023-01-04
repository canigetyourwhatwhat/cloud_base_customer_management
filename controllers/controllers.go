package controllers

import "erply/service"

type Controller struct {
	sessionKey *string
	clientCode *string
	service    *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{service: s}
}
