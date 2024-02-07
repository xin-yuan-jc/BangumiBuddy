package auth

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticator_CheckCookie(t *testing.T) {
	testCases := []struct {
		name      string
		r         *http.Request
		stub      func(t *testing.T) (UserRepository, func())
		newCookie bool
		wantErr   bool
	}{
		{
			name: "no session",
			r: func() *http.Request {
				r, _ := http.NewRequest("GET", "/home", nil)
				return r
			}(),
			stub: func(t *testing.T) (UserRepository, func()) {
				return nil, func() {}
			},
			wantErr: true,
		},
		{
			name: "cookie not right",
			r: func() *http.Request {
				r, _ := http.NewRequest("GET", "/home", nil)
				r.AddCookie(&http.Cookie{Name: "bangumi_session", Value: "wrong"})
				return r
			}(),
			stub: func(t *testing.T) (UserRepository, func()) {
				ctrl := gomock.NewController(t)
				ur := NewMockUserRepository(ctrl)
				ur.EXPECT().GetUserCookie(gomock.Any(), gomock.Any()).Return(UserCookie{}, assert.AnError).AnyTimes()
				return ur, ctrl.Finish
			},
			wantErr: true,
		},
		{
			name: "cookie expired",
			r: func() *http.Request {
				r, _ := http.NewRequest("GET", "/home", nil)
				r.AddCookie(&http.Cookie{Name: "bangumi_session", Value: "wrong"})
				return r
			}(),
			stub: func(t *testing.T) (UserRepository, func()) {
				ctrl := gomock.NewController(t)
				ur := NewMockUserRepository(ctrl)
				ur.EXPECT().GetUserCookie(gomock.Any(), gomock.Any()).Return(UserCookie{
					ExpireAt: time.Now().Add(-time.Second),
				}, nil).AnyTimes()
				return ur, ctrl.Finish
			},
			wantErr: true,
		},
		{
			name: "refresh cookie",
			r: func() *http.Request {
				r, _ := http.NewRequest("GET", "/home", nil)
				r.AddCookie(&http.Cookie{Name: "bangumi_session", Value: "wrong"})
				return r
			}(),
			stub: func(t *testing.T) (UserRepository, func()) {
				ctrl := gomock.NewController(t)
				ur := NewMockUserRepository(ctrl)
				ur.EXPECT().GetUserCookie(gomock.Any(), gomock.Any()).Return(UserCookie{
					Username: "user",
					Cookie:   "cookie",
					ExpireAt: time.Now().Add(time.Second),
				}, nil).AnyTimes()
				ur.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				return ur, ctrl.Finish
			},
			wantErr:   false,
			newCookie: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ur, clo := tc.stub(t)
			defer clo()
			auth := NewAuth(ur)

			got, err := auth.CheckCookie(tc.r)
			t.Log(err)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.newCookie, !reflect.DeepEqual(got, http.Cookie{}))
		})
	}
}

func TestAuthenticator_Login(t *testing.T) {
	type args struct {
		username string
		password string
	}
	testCases := []struct {
		name    string
		args    args
		stub    func(t *testing.T) (UserRepository, func())
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				username: "user",
				password: "password",
			},
			stub: func(t *testing.T) (UserRepository, func()) {
				ctrl := gomock.NewController(t)
				ur := NewMockUserRepository(ctrl)
				ur.EXPECT().GetUser(gomock.Any()).Return(User{
					Username: "user",
					Password: "password",
				}, nil).AnyTimes()
				ur.EXPECT().SetCookie(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				return ur, ctrl.Finish
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ur, clo := tc.stub(t)
			defer clo()
			a := NewAuth(ur)

			cookie, err := a.Login(context.Background(), tc.args.username, tc.args.password)

			assert.Equal(t, tc.wantErr, err != nil)
			if !tc.wantErr {
				assert.NotEqual(t, http.Cookie{}, cookie)
			}
		})
	}
}

func TestAuthenticator_Logout(t *testing.T) {
	type args struct {
		username string
	}
	testCases := []struct {
		name    string
		args    args
		stub    func(t *testing.T) (UserRepository, func())
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				username: "user",
			},
			stub: func(t *testing.T) (UserRepository, func()) {
				ctrl := gomock.NewController(t)
				ur := NewMockUserRepository(ctrl)
				ur.EXPECT().DeleteCookie(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				return ur, ctrl.Finish
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ur, clo := tc.stub(t)
			defer clo()
			a := NewAuth(ur)

			err := a.Logout(context.Background(), tc.args.username)
			t.Log(err)

			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestAuthenticator_UpdateUser(t *testing.T) {
	type args struct {
		username string
		password string
	}
	testCases := []struct {
		name    string
		args    args
		stub    func(t *testing.T) (UserRepository, func())
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				username: "user",
				password: "password",
			},
			stub: func(t *testing.T) (UserRepository, func()) {
				ctrl := gomock.NewController(t)
				ur := NewMockUserRepository(ctrl)
				ur.EXPECT().ApplyUser(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				return ur, ctrl.Finish
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ur, clo := tc.stub(t)
			defer clo()
			a := NewAuth(ur)

			err := a.UpdateUser(context.Background(), tc.args.username, tc.args.password)

			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
