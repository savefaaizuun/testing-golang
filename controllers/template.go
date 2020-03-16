package controllers

import (
	"github.com/astaxie/beego"
)

type TemplateController struct {
	beego.Controller
}

func (c *TemplateController) Get() {
	c.Data["Website"] = "Perputakaan"
	c.Data["Email"] = "savefaaizuun@gmail.com"
	c.TplName = "layout_template.tpl"
	c.Layout = "layout_template.tpl"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "template/header.tpl"
    c.LayoutSections["SidebarMenu"] = "template/sidebar.tpl"
    c.LayoutSections["HtmlFooter"] = "template/footer.tpl"
    c.LayoutSections["Scripts"] = "template/scripts.tpl"
}



