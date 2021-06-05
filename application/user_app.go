package application

import (
	"forum/domain/entity"
	"forum/domain/repository"
)

type UserApp struct {
	us       repository.UserRepository
}

func NewUserApp(us repository.UserRepository) *UserApp {
	return &UserApp{us: us}
}

type UserAppInterface interface {
	CreateUser(user *entity.User) error
	CheckIfUserExists(nickname string) (string, error)
	GetUserByNickname(nickname string) (*entity.User, error)
	UpdateUser(newUser *entity.User) (*entity.User, error)
	GetUserNicknameWithEmail(email string) (string, error)
	GetUsersWithNicknameAndEmail(nickname, email string) ([]entity.User, error)
}

func (us *UserApp) CreateUser(user *entity.User) error {
	return nil
}

func (us *UserApp) CheckIfUserExists(nickname string) (string, error) {
	return "", nil
}

func (us *UserApp) GetUserByNickname(nickname string) (*entity.User, error) {
	return nil, nil
}

func (us *UserApp) UpdateUser(newUser *entity.User) (*entity.User, error) {
	return nil, nil
}

func (us *UserApp) GetUserNicknameWithEmail(email string) (string, error) {
	return "", nil
}

func (us *UserApp) GetUsersWithNicknameAndEmail(nickname, email string) ([]entity.User, error) {
	return nil, nil
}
//type UserAppInterface interface {
//	CreateUser(*entity.User) (int, error)                      // Create user, returns created user's ID
//	SaveUser(*entity.User) error                               // Save changed user to database
//	DeleteUser(int) error                                      // Delete user with passed userID from database
//	GetUser(int) (*entity.User, error)                         // Get user by his ID
//	GetUsers() ([]entity.User, error)                          // Get all users
//	GetUserByUsername(string) (*entity.User, error)            // Get user by his username
//	CheckUserCredentials(string, string) (*entity.User, error) // Check if passed username and password are correct
//	UpdateAvatar(int, io.Reader) error                         // Replace user's avatar with one passed as second parameter
//	Follow(int, int) error                                     // Make first user follow second
//	Unfollow(int, int) error                                   // Make first user unfollow second
//	CheckIfFollowed(int, int) (bool, error)                    // Check if first user follows second. Err != nil if those users are the same
//}

//// CreateUser add new user to database with passed fields
//// It returns user's assigned ID and nil on success, any number and error on failure
//func (u *UserApp) CreateUser(user *entity.User) (int, error) {
//	userID, err := u.us.CreateUser(user)
//	if err != nil {
//		return -1, err
//	}
//
//	initialBoard := &entity.Board{UserID: userID, Title: "Saved pins", Description: "Fast save"}
//	_, err = u.boardApp.AddBoard(initialBoard)
//	if err != nil {
//
//		_ = u.DeleteUser(user.UserID)
//		return -1, err
//	}
//
//	return userID, nil
//}
//
//// SaveUser saves user to database with passed fields
//// It returns nil on success and error on failure
//func (u *UserApp) SaveUser(user *entity.User) error {
//	return u.us.SaveUser(user)
//}
//
//// SaveUser deletes user with passed ID
//// S3AppInterface is needed for avatar deletion
//// It returns nil on success and error on failure
//func (u *UserApp) DeleteUser(userID int) error {
//	user, err := u.us.GetUser(userID)
//	if err != nil {
//		return err
//	}
//
//	if user.Avatar != entity.AvatarDefaultPath {
//
//
//		if err != nil {
//			return err
//		}
//	}
//
//	return u.us.DeleteUser(userID)
//}
//
//// GetUser fetches user with passed ID from database
//// It returns that user, nil on success and nil, error on failure
//func (u *UserApp) GetUser(userID int) (*entity.User, error) {
//	return u.us.GetUser(userID)
//}
//
//// GetUsers fetches all users from database
//// It returns slice of all users, nil on success and nil, error on failure
//func (u *UserApp) GetUsers() ([]entity.User, error) {
//	return u.us.GetUsers()
//}
//
//// GetUserByUsername fetches user with passed username from database
//// It returns that user, nil on success and nil, error on failure
//func (u *UserApp) GetUserByUsername(username string) (*entity.User, error) {
//	return u.us.GetUserByUsername(username)
//}
//
//// GetUserCredentials check whether there is user with such username/password pair
//// It returns user, nil on success and nil, error on failure
//// Those errors are descriptive and tell what did not match
//func (u *UserApp) CheckUserCredentials(username string, password string) (*entity.User, error) {
//	user, err := u.us.GetUserByUsername(username)
//	if err != nil {
//		return nil, err
//	}
//	if user.Password != password { // TODO: hashing
//		return nil, fmt.Errorf("Password does not match")
//	}
//
//	return user, nil
//}
//
//func (u *UserApp) UpdateAvatar(userID int, file io.Reader) error {
//	user, err := u.GetUser(userID)
//	if err != nil {
//		return fmt.Errorf("Could not find user in database")
//	}
//
//	filenamePrefix, err := GenerateRandomString(40) // generating random image
//	if err != nil {
//		return fmt.Errorf("Could not generate filename")
//	}
//
//	newAvatarPath := "avatars/" + filenamePrefix + ".jpg" // TODO: avatars folder sharding by date
//
//	if err != nil {
//		return fmt.Errorf("File upload failed")
//	}
//
//	oldAvatarPath := user.Avatar
//	user.Avatar = newAvatarPath
//	err = u.SaveUser(user)
//	if err != nil {
//
//		return fmt.Errorf("User saving failed")
//	}
//
//	if oldAvatarPath != entity.AvatarDefaultPath {
//
//		if err != nil {
//			return fmt.Errorf("Old avatar deletion failed")
//		}
//	}
//
//	return nil
//}
//
//func (u *UserApp) Follow(followerID int, followedID int) error {
//	if followerID == followedID {
//		return fmt.Errorf("Users can't follow themselves")
//	}
//	return u.us.Follow(followerID, followedID)
//}
//
//func (u *UserApp) Unfollow(followerID int, followedID int) error {
//	if followerID == followedID {
//		return fmt.Errorf("Users can't unfollow themselves")
//	}
//	return u.us.Unfollow(followerID, followedID)
//}
//
//func (u *UserApp) CheckIfFollowed(followerID int, followedID int) (bool, error) {
//	if followerID == followedID {
//		return false, fmt.Errorf("Users can't follow themselves")
//	}
//	return u.us.CheckIfFollowed(followerID, followedID)
//}
