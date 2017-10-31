package agenda
import(
  //"fmt"
  "sync"
  "os"
  "io/ioutil"
  "encoding/json"
  "strings"
)
type Storage struct {
  users []User
  meetings []Meeting
}

var instance *Storage
var once sync.once

func GetStorageInstance() *Storage {
    once.Do(func() {
      instance = &Storage
    })
    return instance
}

type StorageError string
const (
	// 文件夹--存在/创建
	SucceedCreateDataDir StorageError = "Succeed In Creating Data Dir"
	ExistFileNamedData   StorageError = "Fail To Create Data Dir Because Of Exisiting A File Named \"Data\""
	FailCreateDataDir    StorageError = "Fail To Create Data Dir"
	// 文件--创建
	SucceedCreateDataFile StorageError = "Succeed In Creating Data File"
	FailCreateDataFile    StorageError = "Fail To Create Data File"
	// 文件--读取
	SucceedReadDateFile StorageError = "Succeed In Reading Data File"
	FailReadDataFile    StorageError = "Fail To Read Data File"
	// 文件--获取json数据
	FailGetJsonData StorageError = "Fail To Read Json Data"
	// 文件--写入
	SucceedWriteDataFile StorageError = "Succeed In Writing Data File"
	FailWriteDataFile    StorageError = "Fail To Write Data File"
)

func (storage *Storage) ReadFromUserFile() (bool, StorageError) {
  usersJson, err := ReadFromFile(UserDataPath)
	if err != nil {
		return false, FailReadDataFile
	}
	// 以换行符分隔再逐个解析
	usersList := strings.Split(string(usersJson), "\n")
	for i := 0; i < len(usersList); i++ {
		user := &User{}
		if json.Unmarshal([]byte(usersList[i]), &user) != nil {
			return false, FailGetJsonData
		}
		storage.Users = append(storage.Users, user)
	}
	return true, SucceedReadDateFile
}

func (storage *Storage) ReadFromMeetingFile() (bool, StorageError) {
	meetingsJson, err := ReadFromFile(MeetingDataPath)
	if err != nil {
		return false, FailReadDataFile
	}
	// 以换行符分隔再逐个解析
	meetingsList := strings.Split(string(meetingsJson), "\n")
	for i := 0; i < len(meetingsList); i++ {
		meeting := &Meeting{}
		if json.Unmarshal([]byte(meetingsList[i]), &meeting) != nil {
			return false, FailGetJsonData
		}
		storage.Meetings = append(storage.Meetings, meeting)
	}
	return true, SucceedReadDateFile
}


func WriteToFile(fileName string, content []byte) bool {
	if ioutil.WriteFile(fileName, content, 0644) != nil {
		return false
	}
	return true
}

func (storage *Storage) WriteUserFile() bool {
	var userStringList []string
	for i := 0; i < len(storage.Users); i++ {
		userJson, err := json.Marshal(storage.Users[i])
		if err != nil {
			return false
		}
		userStringList = append(userStringList, string(userJson))
	}
	return WriteToFile(UserDataPath, []byte(strings.Join(userStringList, "\n")))
}

func (storage *Storage) WriteMeetingFile() bool {
	var meetingStringList []string
	for i := 0; i < len(storage.Meetings); i++ {
		meetingJson, err := json.Marshal(storage.Meetings[i])
		if err != nil {
			return false
		}
		meetingStringList = append(meetingStringList, string(meetingJson))
	}
	return WriteToFile(MeetingDataPath, []byte(strings.Join(meetingStringList, "\n")))
}

func (storage *Storage) LogOutStorage() (bool, StorageError) {
	instance = nil
	if !storage.WriteUserFile() {
		return false, FailWriteDataFile
	}
	if !storage.WriteMeetingFile() {
		return false, FailWriteDataFile
	}
	return true, SucceedWriteDataFile
}

//----------------------------------------operate

func (storage *Storage) CreateUser(user User) bool {
	storage.Users = append(storage.Users, user)
	return storage.WriteUserFile()
}
//filter the users
func (storage *Storage) QueryUsers(filter func(user User) bool) []User {
	var users []User
	for _, tUser := range storage.Users {
		if filter(tUser) {
			users = append(users, tUser)
		}
	}
	return users
}


// 更新用户信息，返回是否是否更新成功
func (storage *Storage) UpdateUser(filter func(user User) bool, updatedUser User) bool {
	for index, tUser := range storage.Users {
		if filter(tUser) {
			storage.Users[index] = updatedUser
			return storage.WriteUserFile()
		}
	}
	return false
}

func (storage *Storage) DeleteUser(filter func(user User) bool) bool {
	isDeleted := false // 是否进行过删除
	for index, tUser := range storage.Users {
		if filter(tUser) {
			storage.Users = append(storage.Users[:index], storage.Users[index+1:])
			isDeleted = true
			break
		}
	}
	return isDeleted && storage.WriteUserFile()
}

// ----------- 对会议列表进行操作，需要把改动写入文件，并返回是否成功 ------------
// 创建会议
func (storage *Storage) CreateMeeting(meeting Meeting) bool {
	storage.Meetings = append(storage.Meetings, meeting)
	return storage.WriteMeetingFile()
}

// 根据filter函数查找会议
func (storage *Storage) QueryMeetings(filter func(meeting Meeting) bool) []Meeting {
	var meetings []Meeting
	for _, tMeeting := range storage.Meetings {
		if filter(tMeeting) {
			meetings = append(meetings, tMeeting)
		}
	}
	return meetings
}
// 更新会议信息，返回是否是否更新成功
func (storage *Storage) UpdateMeeting(filter func(meeting Meeting) bool, updatedMeeting Meeting) bool {
	for index, tMeeting := range storage.Meetings {
		if filter(tMeeting) {
			storage.Meetings[index] = updatedMeeting
			return storage.WriteMeetingFile()
		}
	}
	return false
}

// 删除会议
func (storage *Storage) DeleteMeeting(filter func(meeting Meeting) bool) bool {
	isDeleted := false // 是否进行过删除
	for index, tMeeting := range storage.Meetings {
		if filter(tMeeting) {
			storage.Meetings = append(storage.Meetings[:index], storage.Meetings[index+1:])
			isDeleted = true
		}
	}
	return isDeleted && storage.WriteUserFile()
}
