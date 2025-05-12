package schemas

type FollowAndUnfollowUserSchema struct {
	FollowerID  int64 `json:"username" validate:"required"`
	FollowingID int64 `json:"password" validate:"required"`
}
