package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
	// "log"
	// "strings"
)

type TemplateController struct {
	beego.Controller
}


func (c *TemplateController) Get() {
	c.Data["Website"] = "Perputakaan"
	c.Data["Email"] = "savefaaizuun@gmail.com"
	c.TplName = "template_layout.tpl"
	c.Layout = "template_layout.tpl"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "template/header.tpl"
    c.LayoutSections["SidebarMenu"] = "template/sidebar.tpl"
    c.LayoutSections["HtmlFooter"] = "template/footer.tpl"
    c.LayoutSections["Scripts"] = "template/script.tpl"
    
}

// func (c *UserController) UserIndex() {
// 	c.TplName = "user/users.tpl"
// 	c.Layout = "user/users.tpl"
// 	c.LayoutSections = make(map[string]string)
// 	c.Data["UserIdentities"] = c.GetSession("user_login")
// 	c.LayoutSections["ScriptPage"] = "scriptPage/users.tpl"
// }

