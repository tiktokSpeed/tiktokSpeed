package main

import (
	"context"
	api "github.com/tiktokSpeed/tiktokSpeed/shared/kitex_gen/api"
)

// ApiServiceImpl implements the last service interface defined in the IDL.
type ApiServiceImpl struct{}

// Register implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) Register(ctx context.Context, req *api.DouyinUserRegisterRequest) (resp *api.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) Login(ctx context.Context, req *api.DouyinUserLoginRequest) (resp *api.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) GetUserInfo(ctx context.Context, req *api.DouyinUserRequest) (resp *api.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// Feed implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) Feed(ctx context.Context, req *api.DouyinFeedRequest) (resp *api.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishVideo implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) PublishVideo(ctx context.Context, req *api.DouyinPublishActionRequest) (resp *api.DouyinPublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// VideoList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) VideoList(ctx context.Context, req *api.DouyinPublishListRequest) (resp *api.DouyinPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// Favorite implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) Favorite(ctx context.Context, req *api.DouyinFavoriteActionRequest) (resp *api.DouyinFavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) FavoriteList(ctx context.Context, req *api.DouyinFavoriteListRequest) (resp *api.DouyinFavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// Comment implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) Comment(ctx context.Context, req *api.DouyinCommentActionRequest) (resp *api.DouyinCommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) CommentList(ctx context.Context, req *api.DouyinCommentListRequest) (resp *api.DouyinCommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// Action implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) Action(ctx context.Context, req *api.DouyinRelationActionRequest) (resp *api.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowingList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) FollowingList(ctx context.Context, req *api.DouyinRelationFollowListRequest) (resp *api.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) FollowerList(ctx context.Context, req *api.DouyinRelationFollowerListRequest) (resp *api.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) FriendList(ctx context.Context, req *api.DouyinRelationFriendListRequest) (resp *api.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}

// ChatHistory implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) ChatHistory(ctx context.Context, req *api.DouyinMessageChatRequest) (resp *api.DouyinMessageChatResponse, err error) {
	// TODO: Your code here...
	return
}

// SentMessage implements the ApiServiceImpl interface.
func (s *ApiServiceImpl) SentMessage(ctx context.Context, req *api.DouyinMessageActionRequest) (resp *api.DouyinMessageActionResponse, err error) {
	// TODO: Your code here...
	return
}
