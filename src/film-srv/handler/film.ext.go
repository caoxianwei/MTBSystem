package handler

import (
	"context"
	"film-srv/db"
	"go.uber.org/zap"
	"share/pb"
	"share/utils/log"
	errors "share/errors"
	"strconv"
)

type FilmServiceExtHandler struct {
	logger *zap.Logger
}

func NewFilmServiceExtHandler() *FilmServiceExtHandler {
	return &FilmServiceExtHandler{
		logger: log.Instance(),
	}
}

// 正在售票列表
func (f *FilmServiceExtHandler) HotPlayMovies(ctx context.Context, req *pb.HotPlayMoviesReq, rsp *pb.HotPlayMoviesRep) error {

	films, err := db.SelectTickingFilims()
	if err != nil {
		f.logger.Error("err", zap.Any("films", err))
		return err
	}
	MoviesPB := []*pb.Movie{}
	for _, film := range films {
		// 处理影片演员信息
		filmActors,err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			return err
		}
		for _,filmActor := range filmActors {
			film.ActorName =  append(film.ActorName,filmActor.ActorName)
		}
		// 处理影片种类信息

		filmPB := film.ToProtoDBMovies()
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
	return nil
}

// 电影详情
func (f *FilmServiceExtHandler) MovieDetail(ctx context.Context, req *pb.MovieDetailReq, rsp *pb.MovieDetailRep) error {

	movieId := req.MovieId
	film,err := db.SelectFilmDetail(movieId)
	if err != nil {
		f.logger.Error("error",zap.Error(err))
		return errors.ErrorFilmFailed
	}
	rsp.TitleEn = film.TitleEn
	rsp.TitleEn = film.TitleEn
	rsp.CommonSpecial = film.CommonSpecial
	rsp.Content = film.CommonSpecial
	rsp.Image = film.Img
	rsp.Rating = film.RatingFinal
	rsp.RunTime = film.Length
	rsp.Year = film.RYear
	str1 := strconv.Itoa(int(film.RYear))+"-"+strconv.Itoa(int(film.RMonth))+"-"+strconv.Itoa(int(film.RDay))
	str2 := film.Country
	pb := pb.Release{
		Date:str1,
		Location:str2,
	}
	rsp.Release = &pb
	return nil
}

// 获取演员和导演信息
func (f *FilmServiceExtHandler) MovieCreditsWithTypes(ctx context.Context, req *pb.MovieCreditsWithTypesReq, rsp *pb.MovieCreditsWithTypesRep) error {

	movieId := req.MovieId
	persons := []*pb.Persons{}
	directors,err := db.SelectActors(movieId,2)
	if err != nil {
		f.logger.Error("error",zap.Error(err))
		return errors.ErrorFilmFailed
	}

	for _,director := range directors {
		directorPB := pb.Persons{
			Name: director.NameCN,
			NameEn: director.NameEN,
			Image: director.ActorPhoto,
		}
		persons = append(persons, &directorPB)
	}

	typE := pb.Type{
		Persons:persons,
		TypeName:"导演",
		TypeNameEc:"Director",
	}

	actors,err := db.SelectActors(movieId,1)
	persons1 := []*pb.Persons{}

	for _,actor := range actors {
		actor := pb.Persons{
			Name: actor.NameCN,
			NameEn: actor.NameEN,
			Image: actor.ActorPhoto,
		}
		persons1 = append(persons1, &actor)
	}

	typE1 := pb.Type{
		Persons:persons1,
		TypeName:"演员",
		TypeNameEc:"Director",
	}
	types := []*pb.Type{}
	types = append(types, &typE)
	types = append(types, &typE1)
	rsp.Types = types
	return nil
}

// 获取剧照信息
func (f *FilmServiceExtHandler) ImageAll(ctx context.Context, req *pb.ImageAllReq, rsp *pb.ImageAllRep) error {

	movieId := req.MovieId
	images,err := db.SelectFilimImages(movieId)
	imagesPB := []*pb.Image{}
	if err != nil {
		f.logger.Error("error",zap.Error(err))
		return errors.ErrorFilmFailed
	}
	for _,image := range images {
		imagePB := pb.Image{
			Image:image.ImageUrl,
		}
		imagesPB = append(imagesPB, &imagePB)
	}
	rsp.Images = imagesPB
	return nil
}

// 即将上映的影片
func (f *FilmServiceExtHandler) LocationMovies(ctx context.Context, req *pb.LocationMoviesReq, rsp *pb.LocationMoviesRep) error {

	films, err := db.SelectTickingFilims()
	if err != nil {
		f.logger.Error("err", zap.Any("films", err))
		return err
	}
	MoviesPB := []*pb.Movie{}
	for _, film := range films {
		// 处理影片演员信息
		filmActors,err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			return err
		}
		for _,filmActor := range filmActors {
			film.ActorName =  append(film.ActorName,filmActor.ActorName)
		}
		// 处理影片种类信息

		filmPB := film.ToProtoDBMovies()
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
	return nil
}

// 即将上映的影片
func (f *FilmServiceExtHandler) MovieComingNew(ctx context.Context, req *pb.MovieComingNewReq, rsp *pb.MovieComingNewRep) error {

	films, err := db.SelectTickingFilims()
	if err != nil {
		f.logger.Error("err", zap.Any("films", err))
		return err
	}
	MoviesPB := []*pb.Movie{}
	for _, film := range films {
		// 处理影片演员信息
		filmActors,err := db.SelectFilmActorByMid(film.MovieId)
		if err != nil {
			return err
		}
		for _,filmActor := range filmActors {
			film.ActorName =  append(film.ActorName,filmActor.ActorName)
		}
		filmPB := film.ToProtoDBMovies()
		MoviesPB = append(MoviesPB, filmPB)
	}
	rsp.Movies = MoviesPB
	return nil
}

// 搜搜出来的影片信息
func (f *FilmServiceExtHandler) Search(ctx context.Context, req *pb.SearchReq, rsp *pb.SearchRep) error {

	rsp.Total = 2
	genres1 := pb.Genres{
		Type:"喜剧",
	}
	genres2 := pb.Genres{
		Type:"悲剧",
	}
	rating := pb.Rating{
		Average:3.7,
	}
	genres := []*pb.Genres{}
	genres = append(genres,&genres1)
	genres = append(genres,&genres2)
	image1 := pb.Images{
		Small:"https://img3.doubanio.com/view/celebrity/s_ratio_celebrity/public/p54250.jpg",
	}
	movie1 := pb.SearchMovie{
		Image:&image1,
		Genres:genres,
		Title:"金钱掌控",
		Id:"1",
		Year:"2017",
		Rating:&rating,
	}
	movie2 := pb.SearchMovie{
		Image:&image1,
		Genres:genres,
		Title:"金钱掌控",
		Id:"1",
		Year:"2017",
		Rating:&rating,
	}
	searchMovies := []*pb.SearchMovie{}
	searchMovies = append(searchMovies, &movie1)
	searchMovies = append(searchMovies, &movie2)
	rsp.Subjects = searchMovies
	return nil
}


func (f *FilmServiceExtHandler) GetFilmsByCid(ctx context.Context, req *pb.GetFilmsByCidReq, rsp *pb.GetFilmsByCidRsp) error {

	return nil
}