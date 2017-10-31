package agenda

type Meeting struct {
  Sponsor string
  Participators string
  StartDate string
  EndDate string
  Title string
}
/*func (Meeting meeting) isParticipator(username string) bool {
    ifTrue := false
    for i := 0 i < meeting.m_participators.size(); i++
}*/
func (meeting Meeting) GetSponsor() string {
  return meeting.Sponsor
}
func (meeting Meeting) GetParticipator() string {
  return meeting.Participators
}
func (meeting Meeting) GetStartDate() string {
  return meeting.StartDate
}
func (meeting Meeting) GetEndDate() string {
  return meeting.EndDate
}
func (meeting Meeting) GetTitle() string {
  return meeting.Title
}
func (meeting *Meeting) SetSponsor(sponsor string) string {
  meeting.Sponsor = sponsor
}
func (meeting *Meeting) SetParticipator(participators string) string {
  meeting.Participators = participators
}
func (meeting *Meeting) SetStartDate(startDate string) string {
  meeting.StartDate = startDate
}
func (meeting *Meeting) SetEndDate(endDate string) string {
  meeting.EndDate = endDate
}
func (meeting *Meeting) SetTitle(title string) string {
  meeting.Title = title
}
