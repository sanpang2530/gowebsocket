package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/youthlin/t"
	"golang.org/x/text/language"
)

const (
	CookieKeyLang        = "lang"
	HeaderKeyLang        = "lang"
	HeaderAcceptLanguage = "Accept-Language" // 浏览器语言
)

func T(ctx *gin.Context) {
	if ctx.GetHeader("target") == "wakeUpNeo" {
		// 后台只用中文
		// model.SetTs(ctx, "zh_CN")
		ctx.Next()
		return
	}
	// 非后台不能出现中文

	// 请求头设置了语言
	lang := ctx.GetHeader(HeaderKeyLang)

	//fmt.Println("GetHeader(HeaderKeyLang) >>> ", lang)
	if lang != "" {
		//lang = model.Match(lang)
		////fmt.Println("lang1 > ", lang)
		//lang = locale.Normalize(lang)
		////fmt.Println("lang2 > ", lang)
		//model.SetTs(ctx, lang)
		ctx.Next()
		return
	}

	// 2 cookie 中设置了语言
	lang, err := ctx.Cookie(CookieKeyLang)
	if err == nil {
		//lang = model.Match(lang)
		//lang = locale.Normalize(lang)
		//model.SetTs(ctx, lang)
		ctx.Next()
		return
	}

	// 3 浏览器标头获取语言
	accept := ctx.GetHeader(HeaderAcceptLanguage)
	if accept == "" {
		//lang = t.SourceCodeLocale()
		//model.SetTs(ctx, lang)
		ctx.Next()
		return
	}

	supported := t.Locales()
	var supportedTags []language.Tag
	for _, lang := range supported {
		supportedTags = append(supportedTags, language.Make(lang))
	}
	matcher := language.NewMatcher(supportedTags)
	userPref, _, err := language.ParseAcceptLanguage(accept)
	if err == nil {
		_, index, _ := matcher.Match(userPref...)
		lang = supported[index]
		//lang = model.Match(lang)
		//model.SetTs(ctx, lang)
	}
	ctx.Next()
}
