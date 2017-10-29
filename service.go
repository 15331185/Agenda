package agenda
import(
  "fmt"
)

type Service struct {
    AgendaStorage *storage.Storage
}

func StartAgenda(service *Service) (bool, storage.StorageError) {
  service.AgendaStorage = storage.GetStorageInstance()
return service.AgendaStorage.ReadFromDataFile()
}

func (service *Service) QuitAgenda() {
	// return false
}

func (service *Service) UserLogin(userName string, password string) bool {
	// 根据获得的用户名和密码判断是否存在该用户名
	return len(service.AgendaStorage.QueryUsers(func(user User) bool {
		return user.GetUserName() == userName && user.GetPassword() == password
	})) == 1
}

func (service *Service) DeleteUser(userName string, password string) bool {
	if len(service.AgendaStorage.QueryUsers(func(user model.User) bool {
		return user.GetUserName() == userName && user.GetPassword() == password
	})) != 1 {
		return false		// 不存在同名用户则删除失败
	}
	// 存在同名用户则进行删除操作
	return service.AgendaStorage.DeleteUser(func(user model.User) bool {
		return user.GetUserName() == userName && user.GetPassword() == password
	})
}

func (service *Service) CreateMeeting(sponsor string, title string,
				startDate string, endDate string, participators []model.User) bool {
	// 判断title是否已存在
	if len(service.AgendaStorage.QueryMeetings(func(meeting Meeting) bool {
		return meeting.GetTitle() == title
	})) > 0 {
		return false
	}
	// 判断时间合法
	// 判断时间字符串是否符合格式要求：2012-2-2/11:23并解析字符串为Int数组
	var startDateIntArray [5]int
	var endDateIntArray [5]int
	if !changetoint(startDate, &startDateIntArray) && !changetoint(endDate, &endDateIntArray) {
		return false
	}
	// 判断时间数字是否合法
	// if !model.IsValidDateTime(startDateIntArray)
	// 判断开始和结束时间是否大小不对
	// 判断参与者是否有其他同时段的会议
	// 创建会议
	return true
}

func (service *Service) MeetingQueryByTitle(userName string, title string) []Meeting {
	return []Meeting{}
}

// 列出该用户发起或参与的所有会议

func (service *Service) ListAllMeetings(userName string) []Meeting {

	return []Meeting{}

}



// 列出该用户发起的所有会议

func (service *Service) ListAllSponsorMeetings(userName string, password string) bool {

	return false

}

// 列出该用户参加的所有会议

func (service *Service) ListAllParticipateMeetings(userName string, password string) bool {
	return false
}

// 删除发起者sponsor题目title会议

func (service *Service) DeleteMeeting(sponsor string, title string) bool {
	return false
}

// 删除sponsor所有会议
func (service *Service) DeleteAllMeetings(sponsor string) bool {
	return false
}
