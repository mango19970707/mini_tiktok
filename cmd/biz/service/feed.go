package service

import (
	"context"
	"time"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
)

const maxCount int = 30 //一次获取视频的最大数量

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (s *FeedService) Feed(req *biz.FeedRequest) ([]*biz.Video, int64, error) {
	latestTime := req.LatestTime
	if latestTime == 0 {
		latestTime = time.Now().Unix()
	}

	//从数据库查找数据
	t:=TimeStampToTime(latestTime)
	vdos, err := dal.QueryVideosByTime(t)
	if err!=nil{
		return []*biz.Video{},time.Now().Unix(),err
	}

	//获取下次的最新时间
	nextTime := vdos[len(vdos)-1].PublishDate.Unix()
	videos := []*biz.Video{}
	for i:=0;i<len(vdos);i++{
		videos=append(videos,MongoVdoToBizVdo(vdos[i]))
	}

	return videos, nextTime, nil
}

//将video.go中的Video转化为biz.pb.go中的video类型
func MongoVdoToBizVdo(vdo *dal.Video) *biz.Video {
	res:=&biz.Video{}
	res.Id=vdo.VideoId
	//从mongodb查询对应用户信息
	f := true
	u1:=biz.User{
		Id:            1,
		Name:          "TestUser",
		FollowCount:   10,
		FollowerCount: 110,
		IsFollow:      &f,
	}
	res.Author=&u1
	res.PlayUrl=vdo.PlayUrl
	res.CoverUrl=vdo.CoverUrl
	res.FavoriteCount=vdo.FavoriteCount
	res.CommentCount=int64(len(vdo.Comments))
	//如果用户登录，根据token判断当前用户是否点赞
	res.IsFavorite=&f
	return res
}

func TimeStampToTime(stamp int64)time.Time{
	tm := time.Unix(stamp, 0)
	t := tm.Format("2006-01-02 15:04:05")
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, t, loc) //使用模板在对应时区转化为time.time类型
	return theTime
}
