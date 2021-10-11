package services

//
//import (
//	"github.com/Summrry-top/gin-glory-api/models"
//	"github.com/Summrry-top/gin-glory-api/utils"
//)
//
//func GetMusic(g *models.GetParam) *models.Response {
//	if utils.Zero(g.Id) {
//		return GetMusicAll(g)
//	}
//	return GetMusicById(g.Id)
//}
//
//func GetMusicById(id int64) *models.Response {
//	var music models.Music
//	if ok := GetOne("id", id, &music); !ok {
//		return Err400("Music不存在！")
//	}
//	return Success("Music单条数据", music.GetMusicJson())
//}
//
//func GetMusicAll(p *models.GetParam) *models.Response {
//	var musics []models.Music
//	var pageData models.PageData
//	if ok := GetPagination(p, &pageData, &musics); !ok {
//		return Err400("Music没有数据！")
//	}
//	return Success("Music多条数据", pageData)
//}
//
//func PostMusic() *models.Response {
//
//	return nil
//}
//
//func PostMusicAdd() *models.Response {
//
//	return nil
//}
//
//func PostMusicEdit() *models.Response {
//
//	return nil
//}
//
//func PostMusicDelete() *models.Response {
//
//	return nil
//}
