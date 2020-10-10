package daos

import (
	"order-ops/dtos"
	"order-ops/models"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pborman/uuid"
)

func getDBCLocal() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1@tcp(localhost:3306)/testdb?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return db
}

func Test_orderDaoImpl_Create(t *testing.T) {
	db := getDBCLocal()
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		record *models.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "normal case2",
			fields: fields{
				db: db,
			},
			args: args{
				record: &models.Order{
					OrderNumber:  uuid.New(),
					Quantiny:     1,
					CustomerName: "minh",
					Phone:        "011",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &orderDaoImpl{
				db: tt.fields.db,
			}
			if err := dao.Create(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("orderDaoImpl.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_orderDaoImpl_Updates(t *testing.T) {
	db := getDBCLocal()

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		record *models.Order
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "update only name",
			fields: fields{
				db: db,
			},
			args: args{
				record: &models.Order{
					OrderNumber:  "791e5c75-5417-47e2-908b-001a534e1db4",
					CustomerName: "minh123",
				},
			},
			wantErr: false,
		},
		{
			name: "update only phone",
			fields: fields{
				db: db,
			},
			args: args{
				record: &models.Order{
					OrderNumber: "791e5c75-5417-47e2-908b-001a534e1db4",
					Phone:       "022",
				},
			},
			wantErr: false,
		},
		{
			name: "update label",
			fields: fields{
				db: db,
			},
			args: args{
				record: &models.Order{
					OrderNumber:           "791e5c75-5417-47e2-908b-001a534e1db4",
					TrackingNumber:        "123",
					URL:                   "https://google.com",
					PartnerTrackingNumber: "321",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &orderDaoImpl{
				db: tt.fields.db,
			}
			if err := dao.Updates(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("orderDaoImpl.Updates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_orderDaoImpl_Search(t *testing.T) {
	db := getDBCLocal()
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		queries []dtos.SearchQuery
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Order
		wantErr bool
	}{
		{
			name: "normal case",
			fields: fields{
				db: db,
			},
			args: args{
				queries: []dtos.SearchQuery{
					{
						Key:   "customer_name=?",
						Value: "minh",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &orderDaoImpl{
				db: tt.fields.db,
			}
			_, err := dao.Search(tt.args.queries)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderDaoImpl.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
