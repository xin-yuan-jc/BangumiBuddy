package auth

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/MangataL/BangumiBuddy/internal/config"
)

func TestAuthenticator_Authorize(t *testing.T) {
	type args struct {
		username string
		password string
	}
	testCases := []struct {
		name    string
		stub    func(t *testing.T) (Dependencies, func())
		args    args
		want    Credentials
		wantErr bool
	}{
		{
			name: "success-normal",
			stub: func(t *testing.T) (Dependencies, func()) {
				ctrl := gomock.NewController(t)
				mockConfig := config.NewMockConfig(ctrl)
				mockConfig.EXPECT().GetUsername().Return("user", nil).AnyTimes()
				mockConfig.EXPECT().GetPassword().Return("password", nil).AnyTimes()
				mockConfig.EXPECT().GetToken().Return("token", nil).AnyTimes()
				mockCipher := NewMockCipher(ctrl)
				mockCipher.EXPECT().Check(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				mockTokenOperator := NewMockTokenOperator(ctrl)
				gomock.InOrder(
					mockTokenOperator.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("access", nil),
					mockTokenOperator.EXPECT().Generate(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return("refresh", nil),
				)
				return Dependencies{
					Config:        mockConfig,
					Cipher:        mockCipher,
					TokenOperator: mockTokenOperator,
				}, ctrl.Finish
			},
			args: args{
				username: "user",
				password: "password",
			},
			want: Credentials{
				AccessToken:  "access",
				TokenType:    "bearer",
				RefreshToken: "refresh",
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dep, cleanup := tc.stub(t)
			defer cleanup()
			a := New(dep)

			got, err := a.Authorize(nil, tc.args.username, tc.args.password)
			t.Log(err)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.want, got)
		})
	}
}
