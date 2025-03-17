package product

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pujidjayanto/goginboilerplate/internal/dto"
	mockProductService "github.com/pujidjayanto/goginboilerplate/mocks/internal_/service/product"
	"github.com/pujidjayanto/goginboilerplate/pkg/pagination"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTest(t *testing.T) (*gin.Engine, *mockProductService.Service) {
	productService := mockProductService.NewService(t)
	controller := NewController(productService)

	gin.SetMode(gin.TestMode)
	r := gin.New()

	r.GET("/products", controller.Index)
	r.GET("/products/paginated", controller.IndexPaginated)

	return r, productService
}

func TestIndex(t *testing.T) {
	t.Run("success get all products", func(t *testing.T) {
		r, mockService := setupTest(t)

		mockService.EXPECT().GetAll(mock.Anything).Return(&dto.GetAllProductResponse{
			Products: []*dto.ProductItem{
				{
					Id:   1,
					Name: "Avanza",
				},
				{
					Id:   2,
					Name: "Kijang",
				},
			},
		}, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response dto.GetAllProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		mockService.AssertExpectations(t)
	})
}

func TestIndexPaginated(t *testing.T) {
	t.Run("success get paginated products", func(t *testing.T) {
		r, mockService := setupTest(t)

		expectedResponse := &dto.GetAllProductPaginatedResponse{
			Products: []*dto.ProductItem{
				{
					Id:        1,
					Name:      "Avanza",
					Price:     "150000.00",
					Quantity:  5,
					CreatedAt: "2024-03-17T15:57:27Z",
				},
				{
					Id:        2,
					Name:      "Kijang",
					Price:     "200000.00",
					Quantity:  3,
					CreatedAt: "2024-03-17T15:57:27Z",
				},
			},
			PaginationResponse: pagination.PaginationResponse{
				CurrentPage: 1,
				CurrentSize: 10,
				TotalPages:  1,
				TotalData:   2,
			},
		}

		// Mock the service call with any request parameters
		mockService.EXPECT().
			GetAllPaginated(mock.Anything, mock.Anything).
			Return(expectedResponse, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/products/paginated?page=1&pageSize=10", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		type productData struct {
			ID        int       `json:"id"`
			Name      string    `json:"name"`
			Price     string    `json:"price"`
			Quantity  int       `json:"quantity"`
			CreatedAt time.Time `json:"createdAt"`
		}

		var responseWrapper struct {
			Data     []productData `json:"data"`
			Message  string        `json:"message"`
			Metadata struct {
				CurrentPage int `json:"currentPage"`
				CurrentSize int `json:"currentSize"`
				TotalData   int `json:"totalData"`
				TotalPage   int `json:"totalPage"`
			} `json:"metadata"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &responseWrapper)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(responseWrapper.Data))

		mockService.AssertExpectations(t)
	})
}
