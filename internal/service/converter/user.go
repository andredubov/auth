package converter

import (
	"github.com/andredubov/auth/internal/service/model"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToUserFromCreateRequest converts grpc request to chat service layer model
func ToUserFromCreateRequest(r *auth_v1.CreateRequest) model.User {
	return model.User{
		Name:            r.GetInfo().GetName(),
		Email:           r.GetInfo().GetEmail(),
		Role:            model.Role(r.GetInfo().GetRole()),
		Password:        r.GetPassword(),
		PasswordConfirm: r.GetPasswordConfirm(),
	}
}

// ToUserUpdateInfoFromUpdateRequest converts grpc request to chat service layer model
func ToUserUpdateInfoFromUpdateRequest(r *auth_v1.UpdateRequest) model.UpdateUserInfo {
	updateUserInfo := model.UpdateUserInfo{ID: r.GetId()}

	if r.GetInfo().GetName() != nil {
		updateUserInfo.Name = &r.GetInfo().GetName().Value
	}

	if r.GetInfo().GetEmail() != nil {
		updateUserInfo.Email = &r.GetInfo().GetEmail().Value
	}

	if r.GetInfo().GetRole() != auth_v1.UserRole_UNKNOWN {
		role := model.Role(r.GetInfo().GetRole())
		updateUserInfo.Role = &role
	}

	return updateUserInfo
}

// ToUserFromService converts service layer model to grpc model
func ToUserFromService(user *model.User) *auth_v1.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &auth_v1.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserInfoFromService converts service layer model to grpc server model
func ToUserInfoFromService(user *model.User) *auth_v1.UserInfo {
	return &auth_v1.UserInfo{
		Name:  user.Name,
		Email: user.Email,
		Role:  auth_v1.UserRole(user.Role),
	}
}
