package main

type Syllabus struct {
	ID              string `json:"id"`
	Section         string `json:"section"`
	CRN             string `json:"crn"`
	MeetingDays     string `json:"meetingDays"`
	MeetingTimes    string `json:"meetingTimes"`
	FinalExam       string `json:"finalExam"`
	MeetingLocation string `json:"meetingLocation"`
	Course          Course `json:"course"`
	Instructor      string `json:"instructor"`
	Calendar        string `json:"calendar"`
}

type Course struct {
	Name                         string     `json:"name"`
	ID                           string     `json:"id"`
	Number                       int        `json:"number"`
	CreditHours                  int        `json:"creditHours"`
	Description                  string     `json:"description"`
	Prerequisites                string     `json:"prerequisites"`
	LearningOutcomes             []string   `json:"learningOutcomes"`
	ProgramOutcomes              []Outcomes `json:"programOutcomes"`
	BaccalaureateCharacteristics []Outcomes `json:"baccalaureateCharacteristics"`
	TextBook                     string     `json:"textBook"`
	Modules                      []string
}

type Outcomes struct {
	Value   int    `json:"value"`
	Outcome string `json:"outcome"`
}
