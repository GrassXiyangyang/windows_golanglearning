package controllers

import (
	"crawl_movie/models"

	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {

	models.ConnectRedis("127.0.0.1:6379")
	//根url
	sUrl := "https://movie.douban.com/top250"
	models.PutinQueue(sUrl)
	var movieinfo models.MovieInfo
	for {
		//var movieinfo models.MovieInfo
		length := models.GetQueueLength()
		if length == 0 {
			break
		}
		sUrl = models.PopfromQueue()
		if models.IsVisit(sUrl) {
			continue
		}
		req := httplib.Get(sUrl)
		smoviehtml, err := req.String()
		if err != nil {
			panic(err)
		}
		movieinfo.Movie_name = models.GetMovieName(smoviehtml)
		if movieinfo.Movie_name != "" {

			if !models.NameIsVisit(movieinfo.Movie_name) {
				models.AddnameToSet(movieinfo.Movie_name)
				movieinfo.Movie_main_character = models.GetMovieCharacter(smoviehtml)
				movieinfo.Movie_director = models.GetMovieDirctor(smoviehtml)
				movieinfo.Movie_country = models.GetMovieCountry(smoviehtml)
				movieinfo.Movie_grade = models.GetMovieGrade(smoviehtml)
				movieinfo.Movie_language = models.GetMovieLanguage(smoviehtml)
				movieinfo.Movie_on_time = models.GetMovieOnTime(smoviehtml)
				movieinfo.Movie_type = models.GetMovieGenre(smoviehtml)
				movieinfo.Movie_span = models.GetMovieLongTime(smoviehtml)
				movieinfo.Movie_writer = models.GetMovieWrite(smoviehtml)

				//movieinfo.Movie_id = models.GetMovieId(smoviehtml)
				c.Ctx.WriteString(string(movieinfo.Id))
				models.Add_data(&movieinfo)
				movieinfo.Id = 0
			}
		}
		urls := models.GetMovieUrl(smoviehtml)
		for _, url := range urls {
			models.PutinQueue(url)
			//c.Ctx.WriteString(url)
		}
		models.AddToSet(sUrl)
		time.Sleep(time.Second)
		//c.Ctx.WriteString("  end !  ")

	}

	//输入数据库
	// movieinfo.Movie_name = models.GetMovieName(smoviehtml)
	// movieinfo.Movie_main_character = models.GetMovieCharacter(smoviehtml)
	// movieinfo.Movie_director = models.GetMovieDirctor(smoviehtml)
	// movieinfo.Movie_country = models.GetMovieCountry(smoviehtml)
	// movieinfo.Movie_grade = models.GetMovieGrade(smoviehtml)
	// movieinfo.Movie_language = models.GetMovieLanguage(smoviehtml)
	// movieinfo.Movie_on_time = models.GetMovieOnTime(smoviehtml)
	// movieinfo.Movie_type = models.GetMovieGenre(smoviehtml)
	// movieinfo.Movie_span = models.GetMovieLongTime(smoviehtml)
	// // movieinfo.Movie_writer = models.GetMovieWrite(smoviehtml)
	// models.Add_data(&movieinfo)
	//输出屏幕
	// //c.Ctx.WriteString(models.GetMovieDirctor(smoviehtml) + "\n")
	//c.Ctx.WriteString(models.GetMovieName(smoviehtml) + "\n")
	// // c.Ctx.WriteString(models.GetMovieCharacter(smoviehtml) + "\n")
	// c.Ctx.WriteString("write:" + models.GetMovieWrite(smoviehtml) + "\n")
	// // c.Ctx.WriteString("time:" + models.GetMovieLongTime(smoviehtml) + "\n")
	// // c.Ctx.WriteString(models.GetMovieGenre(smoviehtml) + "\n")
	// // c.Ctx.WriteString(models.GetMovieLanguage(smoviehtml) + "\n")
	// // c.Ctx.WriteString(models.GetMovieGrade(smoviehtml) + "\n")
	// // c.Ctx.WriteString(models.GetMovieOnTime(smoviehtml) + "\n")
	// // c.Ctx.WriteString(models.GetMovieCountry(smoviehtml) + "\n")
	//c.Ctx.WriteString(models.GetMovieId(smoviehtml))
	// models.ConnectRedis("127.0.0.1:6379")
	// urls := models.GetMovieUrl(smoviehtml)
	// for _, url := range urls {
	// 	models.PutinQueue(url)
	// }
	// c.Ctx.WriteString(fmt.Sprintf("%v", urls))
}
