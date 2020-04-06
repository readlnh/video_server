package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
}

func TestMain(m *testing.M) {
	clearTables()
	fmt.Printf("Hello Test\n")
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("test", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("test", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddVideoInfo", testAddVideo)
	t.Run("GetVideoInfo", testGetVideoInfo)
	t.Run("DelVideoInfo", testDeleteVideoInfo)
	t.Run("RegetVideoInfo", testRegetVideoInfo)
}

func testAddVideo(t *testing.T) {
	videoInfo, err := AddNewVideo(1, "test_video")
	if err != nil {
		t.Errorf("Error of AddVideo: %v", err)
	}
	tempvid = videoInfo.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of RegetVideo: %v", err)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddNewComments", testAddNewComments)
	t.Run("ListComments", testListComments)
}

func testAddNewComments(t *testing.T) {
	err := AddNewComments("12345", 1, "Hello Video!!!")

	if err != nil {
		t.Errorf("Failed to add video: %v\n", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	comments, err := ListComments(vid, from, to)

	if err != nil {
		t.Errorf("Error of ListComments: %v\n", err)
	}

	fmt.Printf("dsjfkjkldf")

	fmt.Println(len(comments))

	for i, comment := range comments {
		fmt.Printf("comment: %d, %v \n", i, comment)
	}
}
