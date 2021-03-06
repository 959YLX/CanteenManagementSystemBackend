package route

import (
	"geekylx.com/CanteenManagementSystemBackend/src/service"
)

type addGoodsRequest struct {
	Token   string
	Species uint64
	Price   float64
	Name    string
}

type addGoodsResponse struct {
	Success bool `json:"success"`
}

type goodsListRequest struct {
	Token string
}

type goodsListResponse struct {
	Goods []*service.GoodsInfo
}

func addGoods(req interface{}) responseWrapper {
	request, ok := req.(addGoodsRequest)
	if !ok {
		return GenerateErrorResponse(PARAM_TYPE_ERROR_CODE, PARAM_TYPE_ERROR_MESSAGE)
	}
	success, err := service.AddGoods(request.Token, request.Species, request.Price, request.Name)
	if err != nil {
		return GenerateErrorResponse(2, err.Error())
	}
	return GenerateSuccessResponse(addGoodsResponse{
		Success: success,
	})
}

func goodsList(req interface{}) responseWrapper {
	request, ok := req.(goodsListRequest)
	if !ok {
		return GenerateErrorResponse(PARAM_TYPE_ERROR_CODE, PARAM_TYPE_ERROR_MESSAGE)
	}
	serviceGoodsList, err := service.GetGoodsList(request.Token)
	if serviceGoodsList == nil || err != nil {
		return GenerateErrorResponse(2, err.Error())
	}
	return GenerateSuccessResponse(goodsListResponse{
		Goods: serviceGoodsList,
	})
}
