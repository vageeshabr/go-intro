package order

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/vageeshabr/go-intro/10-mocking/services"
	"testing"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSmsSender := services.NewMockSmsSender(ctrl)

	mockEmailSender := services.NewMockEmailSender(ctrl)

	orderSvc := New(mockSmsSender, mockEmailSender)

	type tc struct {
		Input       *OrderCreate
		ExpectedErr error
		Calls       []*gomock.Call
	}

	tcs := []tc{
		{
			Input:       &OrderCreate{},
			ExpectedErr: errors.New("customer details are missing"),
		},
		{
			Input:       &OrderCreate{Customer: &Customer{}},
			ExpectedErr: errors.New("phone and email are missing"),
		},
		{
			Input: &OrderCreate{Customer: &Customer{
				Id:    1,
				Name:  "Vageesha",
				Phone: "+918971469589",
				Email: "",
			}},
			ExpectedErr: nil,
			Calls:       []*gomock.Call{mockSmsSender.EXPECT().Send("+918971469589", "Your order is successfully placed").Return(nil)},
		},
		{
			Input: &OrderCreate{Customer: &Customer{
				Id:    1,
				Name:  "Vageesha",
				Phone: "+918971469589",
				Email: "vageesha@zopsmart.com",
			}},
			ExpectedErr: nil,
			Calls: []*gomock.Call{
				mockSmsSender.EXPECT().Send("+918971469589", "Your order is successfully placed").Return(nil),
				mockEmailSender.EXPECT().Send("vageesha@zopsmart.com", "Your order is successfully placed", "here are the details").Return(nil),
			},
		},
	}
	for i, tc := range tcs {
		actualError := orderSvc.Create(context.Background(), tc.Input)
		if actualError != nil && tc.ExpectedErr != nil && actualError.Error() != tc.ExpectedErr.Error() {
			t.Error(fmt.Sprintf("tc %d has failed", i))
		}
	}

}
