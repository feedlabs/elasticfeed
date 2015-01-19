package stream

import "github.com/feedlabs/feedify"

type baseController struct {
	feedify.Controller // Embed struct that has stub implementation of the interface.
//	i18n.Locale      // For i18n usage when process data and render template.
}

// Prepare implemented Prepare() method for baseController.
// It's used for language option check and setting.
//func (this *baseController) Prepare() {
//	// Reset language option.
//	this.Lang = "" // This field is from i18n.Locale.
//
//	// 1. Get language information from 'Accept-Language'.
//	al := this.Ctx.Request.Header.Get("Accept-Language")
//	if len(al) > 4 {
//		al = al[:5] // Only compare first 5 letters.
////		if i18n.IsExist(al) {
//			this.Lang = al
////		}
//	}
//
//	// 2. Default language is English.
//	if len(this.Lang) == 0 {
//		this.Lang = "en-US"
//	}
//
//	// Set template level language option.
//	this.Data["Lang"] = this.Lang
//}
