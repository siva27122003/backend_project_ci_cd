package Server

import (
	"GRPC/Config"
	"GRPC/model"
	"GRPC/pb"
	"context"
	"errors"
	"log"
)

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
}

// CreateCategory adds a new category to the database
func (s *CategoryServer) CreateCategory(ctx context.Context, in *pb.Category) (*pb.CategoryResponse, error) {
	category := &model.Category{
		ID:           int32(in.Id),
		CategoryName: in.CategoryName,
	}

	err := Config.DB.Create(&category).Error
	if err != nil {
		log.Printf("Error creating category: %v", err)
		return nil, errors.New("error saving category details")
	}

	log.Printf("The Category Created with ID: %d", category.ID)
	return &pb.CategoryResponse{Category: in}, nil
}

// GetCategories retrieves all categories from the database
func (s *CategoryServer) GetCategories(ctx context.Context, in *pb.Empty) (*pb.CategoryList, error) {
	var categories []model.Category
	err := Config.DB.Find(&categories).Error
	if err != nil {
		log.Printf("Error retrieving categories: %v", err)
		return nil, err
	}

	var pbcategory []*pb.Category
	for _, cat := range categories {
		pbcategory = append(pbcategory, &pb.Category{
			Id:           int32(cat.ID),
			CategoryName: cat.CategoryName,
		})
	}

	return &pb.CategoryList{Categories: pbcategory}, nil
}

// DeleteCategory removes a category from the database by ID
func (s *CategoryServer) DeleteCategory(ctx context.Context, in *pb.CategoryRequest) (*pb.DeleteResponse, error) {
	var category model.Category

	err := Config.DB.First(&category, in.Id).Error
	if err != nil {
		log.Printf("Category not found with ID %d: %v", in.Id, err)
		return nil, err
	}

	err = Config.DB.Delete(&category).Error
	if err != nil {
		log.Printf("Error deleting category ID %d: %v", in.Id, err)
		return nil, err
	}

	return &pb.DeleteResponse{Message: "Category record deleted successfully."}, nil
}
