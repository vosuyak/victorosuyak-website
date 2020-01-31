package models

type Education struct {
	School     string
	StartDate  string
	EndDate    string
	Degree     string
	City       string
	State      string
	Online     bool
	Website    string
	WebsiteUrl string
	Price      int
	Type       string
	Courses    []Course
}
type Course struct {
	Name        string
	Title       string
	Description string
	Topic       string
	Rating      string
}
