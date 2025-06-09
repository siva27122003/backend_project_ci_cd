package Server

import (
	"GRPC/Config"
	"GRPC/model"
	"GRPC/pb"
	"context"
	"database/sql"
	"errors"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommodityServer struct {
	pb.UnimplementedCommodityServiceServer
	db *sql.DB
}

// CreateCommodity creates a new commodity and returns it with its category
func (s *CommodityServer) CreateCommodity(ctx context.Context, in *pb.Commodity) (*pb.CommodityResponse, error) {
	commodity := &model.Commodity{
		CommodityID:  int32(in.Id),
		ProductName:  in.ProductName,
		FarmerID:     int32(in.FarmerId),
		Quantity:     int32(in.Quantity),
		BasePrice:    float64(in.BasePrice),
		Availability: in.Availability,
		CategoryID:   int32(in.CategoryId),
	}

	if err := Config.DB.Create(&commodity).Error; err != nil {
		log.Printf("Failed to create commodity: %v", err)
		return nil, errors.New("error saving commodity details")
	}

	// Load the category details
	if err := Config.DB.Preload("Category").First(&commodity, commodity.CommodityID).Error; err != nil {
		log.Printf("Failed to fetch created commodity with category: %v", err)
		return nil, err
	}

	log.Printf("Commodity created with ID: %v", commodity.CommodityID)

	return &pb.CommodityResponse{
		Commodity: &pb.Commodity{
			Id:           commodity.CommodityID,
			ProductName:  commodity.ProductName,
			FarmerId:     commodity.FarmerID,
			Quantity:     commodity.Quantity,
			BasePrice:    float32(commodity.BasePrice),
			Availability: commodity.Availability,
			CategoryId:   commodity.CategoryID,
		},
		Category: &pb.Category{
			Id:           commodity.CategoryID,
			CategoryName: commodity.Category.CategoryName,
		},
	}, nil
}

// GetCommodities returns all commodities
func (s *CommodityServer) GetCommodities(ctx context.Context, in *pb.Empty) (*pb.CommodityList, error) {
	var commodities []model.Commodity

	if err := Config.DB.Find(&commodities).Error; err != nil {
		log.Printf("Failed to fetch commodities: %v", err)
		return nil, err
	}

	var pbCommodities []*pb.Commodity
	for _, com := range commodities {
		pbCommodities = append(pbCommodities, &pb.Commodity{
			Id:           com.CommodityID,
			ProductName:  com.ProductName,
			FarmerId:     com.FarmerID,
			Quantity:     com.Quantity,
			BasePrice:    float32(com.BasePrice),
			Availability: com.Availability,
			CategoryId:   com.CategoryID,
		})
	}

	return &pb.CommodityList{Commodities: pbCommodities}, nil
}

// DeleteCommodity deletes a commodity by ID
func (s *CommodityServer) DeleteCommodity(ctx context.Context, in *pb.CommodityRequest) (*pb.DeleteResponse, error) {
	var commodity model.Commodity

	if err := Config.DB.First(&commodity, in.Id).Error; err != nil {
		log.Println("Commodity not found:", err)
		return nil, errors.New("commodity not found")
	}

	if err := Config.DB.Delete(&commodity).Error; err != nil {
		log.Println("Failed to delete commodity:", err)
		return nil, err
	}

	return &pb.DeleteResponse{Message: "Commodity deleted successfully"}, nil
}

func (s *CommodityServer) UpdateCommodity(ctx context.Context, in *pb.UpdateCommodityReq) (*pb.CommodityResponse, error) {
	var commodity model.Commodity

	if err := Config.DB.Preload("Category").First(&commodity, in.Id).Error; err != nil {
		log.Println("Commodity not found:", err)
		return nil, errors.New("commodity not found")
	}

	commodity.ProductName = in.ProductName
	commodity.Quantity = in.Quantity
	commodity.BasePrice = float64(in.BasePrice)
	commodity.Availability = in.Availability

	if in.CategoryName != "" {
		commodity.Category.CategoryName = in.CategoryName
	}

	if err := Config.DB.Save(&commodity).Error; err != nil {
		return nil, err
	}

	if err := Config.DB.Save(&commodity.Category).Error; err != nil {
		return nil, err
	}

	return &pb.CommodityResponse{
		Commodity: &pb.Commodity{
			Id:           commodity.CommodityID,
			ProductName:  commodity.ProductName,
			FarmerId:     commodity.FarmerID,
			Quantity:     commodity.Quantity,
			BasePrice:    float32(commodity.BasePrice),
			Availability: commodity.Availability,
			CategoryId:   commodity.CategoryID,
		},
		Category: &pb.Category{
			Id:           commodity.CategoryID,
			CategoryName: commodity.Category.CategoryName,
		},
	}, nil
}

func (s *CommodityServer) GetCommodityByID(ctx context.Context, in *pb.CommodityRequest) (*pb.CommodityResponse, error) {
	var commodity model.Commodity

	if err := Config.DB.Preload("Category").First(&commodity, in.Id).Error; err != nil {
		log.Printf("Commodity not found: %v", err)
		return nil, status.Errorf(codes.NotFound, "Commodity not found")
	}

	return &pb.CommodityResponse{
		Commodity: &pb.Commodity{
			Id:           commodity.CommodityID,
			ProductName:  commodity.ProductName,
			FarmerId:     commodity.FarmerID,
			Quantity:     commodity.Quantity,
			BasePrice:    float32(commodity.BasePrice),
			Availability: commodity.Availability,
			CategoryId:   commodity.CategoryID,
		},
		Category: &pb.Category{
			Id:           commodity.CategoryID,
			CategoryName: commodity.Category.CategoryName,
		},
	}, nil
}
