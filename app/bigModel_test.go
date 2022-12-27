package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
	"reflect"
	"testing"
)

func TestRemoveRepeatedElement(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name       string
		args       args
		wantNewArr []string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				arr: []string{
					"victor",
					"abc",
					"ccc",
					"victor",
					"abc",
				},
			},
			wantNewArr: []string{
				"ccc",
				"victor",
				"abc",
			},
		},
		{
			name: "case2",
			args: args{
				arr: []string{
					"a",
					"123",
					"123",
					"bbc",
					"c",
					"a",
				},
			},
			wantNewArr: []string{
				"123",
				"bbc",
				"c",
				"a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewArr := RemoveRepeatedElement(tt.args.arr); !reflect.DeepEqual(gotNewArr, tt.wantNewArr) {
				t.Errorf("RemoveRepeatedElement() = %v, want %v", gotNewArr, tt.wantNewArr)
			}
		})
	}
}

func Test_bigModelRecordService_GetBigModelRecordsByType(t *testing.T) {
	type fields struct {
		ub repository.UserWithBigModel
	}
	type args struct {
		d domain.BigModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantDto BigModelDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bigModelRecordService{
				ub: tt.fields.ub,
			}
			gotDto, err := b.GetBigModelRecordsByType(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("bigModelRecordService.GetBigModelRecordsByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDto, tt.wantDto) {
				t.Errorf("bigModelRecordService.GetBigModelRecordsByType() = %v, want %v", gotDto, tt.wantDto)
			}
		})
	}
}
