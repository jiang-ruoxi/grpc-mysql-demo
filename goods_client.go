package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-mysql-demo/proto"
	"math/rand"
	"strconv"
	"time"
)

//InitMakeClient 初始grpc客户端
func InitMakeClient() (proto.GoodsServiceClient, *grpc.ClientConn) {
	//打开grpc服务端连接
	conn, err := grpc.Dial("127.0.0.1:8899", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("order client grpc.Dial failed,err:%v\n", err)
	}

	//创建grpc客户端连接
	client := proto.NewGoodsServiceClient(conn)
	return client, conn
}

func main() {
	//创建
	HandleCreateGoods()

	//获取
	HandleGetGoodsInfo()

	//更新
	HandleUpdateGoods()

	//删除
	HandleDeleteGoods()

	//获取所有
	HandleGetGoodsList()
}

//HandleCreateGoods 创建数据的操作
func HandleCreateGoods() {
	client, conn := InitMakeClient()
	defer func() {
		conn.Close()
	}()

	goodsName := "iphone " + strconv.Itoa(MakeRandNumber(20))
	//构造请求参数
	createParams := &proto.CreateGoodsRequest{
		Version: "v1",
		Goods: &proto.Goods{
			GoodsName: goodsName,
			Price:     4699,
		},
	}

	//发起调用操作
	res, err := client.CreateGoods(context.Background(), createParams)
	if err != nil {
		fmt.Printf("order client CreateGoods failed,err:%v\n", err)
		return
	}
	fmt.Println("CreateGoods response：", res)
}

//HandleGetGoodsInfo 获取操作
func HandleGetGoodsInfo() {
	client, conn := InitMakeClient()
	defer func() {
		conn.Close()
	}()

	//获取一条数据
	req2 := &proto.GetInfoRequest{
		Version: "v1",
		Id:  2,
	}
	//发起调用
	res, err := client.GetGoodsInfo(context.Background(), req2)
	if err != nil {
		fmt.Printf("order client GetGoodsInfo failed,err:%v\n", err)
		return
	}
	fmt.Println("GetGoodsInfo response：", res)
}

//HandleUpdateGoods 更新数据
func HandleUpdateGoods() {
	client, conn := InitMakeClient()
	defer func() {
		conn.Close()
	}()
	//更新一条数据
	req3 := &proto.UpdateGoodsRequest{
		Version: "v1",
		Goods: &proto.Goods{
			Id:    2,
			Price: int64(MakeRandNumber(10000)),
		},
	}
	//发起调用
	res, err := client.UpdateGoods(context.Background(), req3)
	if err != nil {
		fmt.Printf("order client UpdateGoods failed,err:%v\n", err)
		return
	}
	fmt.Println("UpdateGoods response：", res)
}

//HandleDeleteGoods 删除数据
func HandleDeleteGoods() {
	client, conn := InitMakeClient()
	defer func() {
		conn.Close()
	}()
	//删除一条数据
	req4 := &proto.DeleteGoodsRequest{
		Version: "v1",
		Id:  1,
	}
	//发起调用
	res, err := client.DeleteGoods(context.Background(), req4)
	if err != nil {
		fmt.Printf("order client DeleteGoods failed,err:%v\n", err)
		return
	}
	fmt.Println("DeleteGoods response：", res)
}

//HandleGetGoodsList 获取所有的数据
func HandleGetGoodsList() {
	client, conn := InitMakeClient()
	defer func() {
		conn.Close()
	}()
	//获取所有数据的参数
	req4 := &proto.GetListGoodsRequest{
		Version: "v1",
	}
	//发起调用
	res, err := client.GetListGoods(context.Background(), req4)
	if err != nil {
		fmt.Printf("order client GetListGoods failed,err:%v\n", err)
		return
	}
	fmt.Println("GetListGoods response：", res)
}

//MakeRandNumber 获取1000之内的随机数
func MakeRandNumber(num int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(num)
	return r
}
