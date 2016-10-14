package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Redirect("backstage/sign", 302);
}
