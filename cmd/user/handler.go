package main

import (
	"context"

	"github.com/tiktokSpeed/tiktokSpeed/cmd/user/dao"
	"github.com/tiktokSpeed/tiktokSpeed/pkg/jwt"
	"github.com/tiktokSpeed/tiktokSpeed/shared/consts"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
	"github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/base"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
}

// Feed implements the UserService interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *api.DouyinUserRegisterRequest) (resp *api.DouyinUserRegisterResponse, err error) {

	resp = new(api.DouyinUserRegisterResponse)

	user, err := dao.NewUser(req.Username, req.Password)
	if err != nil {
		resp.StatusCode = int32(consts.ErrCode)
		resp.StatusMsg = "Failed to create user"
		return resp, err
	}

	// generate token by Jason Web Token
	token, err := jwt.CreateToken(jwt.CustomClaims{ID: user.Id})

	if err != nil {
		resp.StatusCode = int32(consts.ErrCode)
		resp.StatusMsg = "Failed to generate token"
		return resp, err
	}

	// generate response
	resp.Token = token
	resp.UserId = user.Id
	resp.StatusCode = int32(consts.CorrectCode)
	resp.StatusMsg = "Register successfully"

	return resp, nil
}

// Login implements the UserService interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *api.DouyinUserLoginRequest) (r *api.DouyinUserLoginResponse, err error) {
	resp := new(api.DouyinUserLoginResponse)

	user, err := dao.GetUserByUsername(req.Username)
	if err != nil {
		resp.StatusCode = int32(consts.ErrCode)
		resp.StatusMsg = "The user does not exist"
		return resp, err
	}

	// use bcrypt to hash password for security, and compare the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		resp.StatusCode = int32(consts.ErrCode)
		resp.StatusMsg = "The password is wrong"
		return resp, err
	}

	// generate token by Jason Web Token
	token, err := jwt.CreateToken(jwt.CustomClaims{ID: user.ID})

	if err != nil {
		resp.StatusCode = int32(consts.ErrCode)
		resp.StatusMsg = "Failed to generate token"
		return resp, err
	}

	resp.StatusCode = int32(consts.CorrectCode)
	resp.StatusMsg = "Login successfully"
	resp.UserId = user.ID
	resp.Token = token
	return resp, nil
}

// GetUserInfo implements the UserService interface.
func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *api.DouyinUserRequest) (r *api.DouyinUserResponse, err error) {
	resp := new(api.DouyinUserResponse)

	// TODO: Requirement clarification: what is token used for?

	user, err := dao.GetUserById(req.UserId)
	if err != nil {
		resp.StatusCode = int32(consts.ErrCode)
		resp.StatusMsg = "Failed to get user info"
		return resp, err
	}

	resp.StatusCode = int32(consts.CorrectCode)
	resp.StatusMsg = "Get user info successfully"

	// fill user info with user, social, interact info
	// TODO: fill user info with user, social, interact info
	resp.User = &base.User{
		Id:   user.ID,
		Name: user.Username,
	}

	return resp, nil
}

// No need those functions below
func (s *UserServiceImpl) Feed(ctx context.Context, req *api.DouyinFeedRequest) (r *api.DouyinFeedResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) PublishVideo(ctx context.Context, req *api.DouyinPublishActionRequest) (r *api.DouyinPublishActionResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) VideoList(ctx context.Context, req *api.DouyinPublishListRequest) (r *api.DouyinPublishListResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) Favorite(ctx context.Context, req *api.DouyinFavoriteActionRequest) (r *api.DouyinFavoriteActionResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) FavoriteList(ctx context.Context, req *api.DouyinFavoriteListRequest) (r *api.DouyinFavoriteListResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) Comment(ctx context.Context, req *api.DouyinCommentActionRequest) (r *api.DouyinCommentActionResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) CommentList(ctx context.Context, req *api.DouyinCommentListRequest) (r *api.DouyinCommentListResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) Action(ctx context.Context, req *api.DouyinRelationActionRequest) (r *api.DouyinRelationActionResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) FollowingList(ctx context.Context, req *api.DouyinRelationFollowListRequest) (r *api.DouyinRelationFollowListResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) FollowerList(ctx context.Context, req *api.DouyinRelationFollowerListRequest) (r *api.DouyinRelationFollowerListResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) FriendList(ctx context.Context, req *api.DouyinRelationFriendListRequest) (r *api.DouyinRelationFriendListResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) ChatHistory(ctx context.Context, req *api.DouyinMessageChatRequest) (r *api.DouyinMessageChatResponse, err error) {
	// TODO: implement this method
	return nil, nil
}

func (s *UserServiceImpl) SentMessage(ctx context.Context, req *api.DouyinMessageActionRequest) (r *api.DouyinMessageActionResponse, err error) {
	// TODO: implement this method
	return nil, nil
}
