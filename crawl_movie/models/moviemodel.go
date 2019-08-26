package models

import (
	"regexp"

	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MovieInfo struct {
	Id                   int
	Movie_id             int
	Movie_name           string
	Movie_picter         string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
	Remark               string
	_Create_time         string
}

var db orm.Ormer

func init() {
	// 需要在init中注册定义的model
	orm.Debug = true //开启调试模式，
	orm.RegisterDataBase("default", "mysql", "root:123@tcp(10.104.118.44:6603)/movie_data?charset=utf8", 30)
	orm.RegisterModel(new(MovieInfo))
	// 自动建表
	orm.RunSyncdb("default", false, true)

}
func Add_data(MovieInfo *MovieInfo) {

	//orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
	db.Insert(MovieInfo)

}

//单个匹配
//电影导演
func GetMovieDirctor(moviehtml string) string {

	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*)</a>`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	return strings.Trim(character, "/")
}

//电影名字
func GetMovieName(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`<span property="v:itemreviewed">(.*)</span>`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

//电影id

//多个匹配

//主演
func GetMovieCharacter(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	return strings.Trim(character, "/")
}

//编剧
func GetMovieWrite(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}

	req := regexp.MustCompile(`<span ><span class='pl'>编剧</span>: <span class='attrs'>(<a.*?>(.*?)</a>)*?</span></span><br/>`)
	//req := regexp.MustCompile(`(?s:<span class="attrs"><a.*?>(.*?)</a></span>)`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	req1 := regexp.MustCompile(`<a href=".*?">(.*?)</a>`)
	result1 := req1.FindAllStringSubmatch(string(character), -1)
	character1 := ""
	for _, value := range result1 {
		character1 += value[1] + "/"
	}
	return strings.Trim(character1, "/")
}

//时长
func GetMovieLongTime(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`<span property="v:runtime".*?>(.*?)</span>`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	return strings.Trim(character, "/")
}

//类型
func GetMovieGenre(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`<span property="v:genre">(.*?)</span>`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	return strings.Trim(character, "/")
}

//语言
func GetMovieLanguage(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`(?s:<span class="pl">语言:</span>(.*?)<br/>)`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	return strings.Trim(character, "/")
}

//评分
func GetMovieGrade(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`<.*?property="v:average">(.*?)</strong>`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	return strings.Trim(character, "/")
}

//上映时间
func GetMovieOnTime(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`<span property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	return strings.Trim(character, "/")
}

//制片国家
func GetMovieCountry(moviehtml string) string {
	if moviehtml == "" {
		return "err"
	}
	req := regexp.MustCompile(`(?s:<span class="pl">制片国家/地区:</span>(.*?)<br/>)`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	if len(result) == 0 {
		return ""
	}
	character := ""
	for _, value := range result {
		character += value[1] + "/"
	}
	return strings.Trim(character, "/")
}

// func GetMovieId(moviehtml string) string {
// 	if moviehtml == "" {
// 		return ""
// 	}
// 	req := regexp.MustCompile(`https://movie.douban.com/subject/(.*?)/.*?`)
// 	result := req.FindAllStringSubmatch(moviehtml, -1)
// 	if len(result) != 0 {
// 		var character string
// 		character = ""
// 		for _, value := range result {
// 			character += value[1] + "/"
// 		}

// 		return strings.Trim(character, "/")
// 	}
// 	return ""
// }

//获取其他链接
func GetMovieUrl(moviehtml string) []string {

	req := regexp.MustCompile(` <a.*?href="(https://movie.douban.com/.*?)"`)
	result := req.FindAllStringSubmatch(moviehtml, -1)
	var movieSets []string
	for _, v := range result {

		movieSets = append(movieSets, v[1])
	}

	return movieSets
}
