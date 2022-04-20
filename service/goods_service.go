package service

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"grpc-mysql-demo/proto"
)

type ModelAuto struct {
	Model   interface{}
	Comment string
}

type Goods struct {
	Id        int64   `json:"id" gorm:"column:id;type:bigint(20) unsigned not null AUTO_INCREMENT;primary_key"`
	GoodsName string  `json:"goods_name" gorm:"column:goods_name;type:varchar(128) not null;default:''"`
	Price     float32 `json:"price" gorm:"column:price;type:int(11) unsigned not null;default:0"`
}

func init() {
	db, err := InitDB()
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = db.Close()
	}()

	//生成表
	db.AutoMigrate(&proto.Goods{})
	//设置存储以前
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&proto.Goods{})
}

//InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", "root", "root", "127.0.0.1", 3306, "jiang", "10s")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败")
		return nil, err
	}
	return db, nil
}

//GoodsService 定义商品相关的服务接口
type GoodsService interface {
	CreateGoods(ctx context.Context, req *proto.CreateGoodsRequest) (*proto.CreatGoodsResponse, error)
	GetGoodsInfo(ctx context.Context, req *proto.GetInfoRequest) (*proto.GetInfoResponse, error)
	UpdateGoods(ctx context.Context, req *proto.UpdateGoodsRequest) (*proto.UpdateGoodsResponse, error)
	DeleteGoods(ctx context.Context, req *proto.DeleteGoodsRequest) (*proto.DeleteGoodsResponse, error)
	GetListGoods(ctx context.Context, req *proto.GetListGoodsRequest) (*proto.GetListGoodsResponse, error)
}

//定义服务的实现者，首字母小写，不能被其他的包使用，所以会增加一个工厂函数创建实例
type goodsService struct {
}

//NewGoodsService 构建工厂实例
func NewGoodsService() GoodsService {
	return &goodsService{}
}

//CreateGoods 实现创建商品的接口
func (g *goodsService) CreateGoods(ctx context.Context, req *proto.CreateGoodsRequest) (*proto.CreatGoodsResponse, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = db.Close()
	}()

	//组装插入数据库的字段信息
	goods := &proto.Goods{
		GoodsName: req.Goods.GoodsName,
		Price:     req.Goods.Price,
	}

	//执行入库操作
	result := db.Create(&goods)
	if result.Error != nil {
		return nil, errors.New("CreateGoods failed")
	}
	lastId := goods.Id

	//返回信息
	return &proto.CreatGoodsResponse{Version: "v1", Id: lastId}, nil
}

//GetGoodsInfo 实现获取单条数据的接口
func (g *goodsService) GetGoodsInfo(ctx context.Context, req *proto.GetInfoRequest) (*proto.GetInfoResponse, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = db.Close()
	}()

	//定义返回格式结构体
	var goods proto.Goods

	//根据条件查询响应的数据
	db.Where("id = ?", req.Id).First(&goods)
	return &proto.GetInfoResponse{Version: "v1", Goods: &goods}, nil
}

//UpdateGoods 实现更新单条记录的接口
func (g *goodsService) UpdateGoods(ctx context.Context, req *proto.UpdateGoodsRequest) (*proto.UpdateGoodsResponse, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = db.Close()
	}()

	result := db.Exec("UPDATE goods SET price = ? where id = ?", req.Goods.Price, req.Goods.Id)
	fmt.Println(result.Error)
	return &proto.UpdateGoodsResponse{Version: "v1", Updated: result.RowsAffected}, err
}

//DeleteGoods 实现删除单条数据的接口
func (g *goodsService) DeleteGoods(ctx context.Context, req *proto.DeleteGoodsRequest) (*proto.DeleteGoodsResponse, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = db.Close()
	}()

	row := db.Delete(&proto.Goods{}, req.Id).RowsAffected
	return &proto.DeleteGoodsResponse{Version: "v1", Deleted: row}, nil
}

//GetListGoods 实现获取所有数据的接口
func (g *goodsService) GetListGoods(ctx context.Context, req *proto.GetListGoodsRequest) (*proto.GetListGoodsResponse, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = db.Close()
	}()

	//定义查询所有数据格式结构体
	var goods []proto.Goods
	db.Find(&goods)

	//定义返回格式的结构体
	list := []*proto.Goods{}
	for _, item := range goods {
		td := &proto.Goods{
			GoodsName: item.GoodsName,
			Price:     item.Price,
			Id:        item.Id,
		}
		list = append(list, td)
	}

	//数据返回
	return &proto.GetListGoodsResponse{
		Version:       "v1",
		GoodsList: list,
	}, nil
}
