package enums

import "fmt"

// SectionType defines the hacker news sections that are available to scrap
type SectionType string

const (
	SectionInvalid      SectionType = "invalid"
	SectionTypeNew      SectionType = "new" // newstories
	SectionTypePast     SectionType = "past"
	SectionTypeComments SectionType = "comments"
	SectionTypeAsk      SectionType = "ask"  // askstories
	SectionTypeShow     SectionType = "show" // showstories
	SectionTypeJobs     SectionType = "jobs" // jobstories
	SectionTypeSubmit   SectionType = "submit"
)

func (s SectionType) String() string {
	return string(s)
}

func (s SectionType) IsValid() bool {
	switch s {
	case SectionTypeNew, SectionTypePast, SectionTypeComments, SectionTypeAsk, SectionTypeShow, SectionTypeJobs, SectionTypeSubmit:
		return true
	default:
		return false
	}
}

func (s SectionType) ApiString() string {
	switch s {
	case SectionTypeNew:
		return "newstories"
	case SectionTypePast:
		return "paststories"
	case SectionTypeComments:
		return "commentstories"
	case SectionTypeAsk:
		return "askstories"
	case SectionTypeShow:
		return "showstories"
	case SectionTypeJobs:
		return "jobstories"
	case SectionTypeSubmit:
		return "submitstories"
	default:
		return ""
	}
}

func ParseSectionType(s string) (SectionType, error) {
	switch s {
	case "new":
		return SectionTypeNew, nil
	case "past":
		return SectionTypePast, nil
	case "comments":
		return SectionTypeComments, nil
	case "ask":
		return SectionTypeAsk, nil
	case "show":
		return SectionTypeShow, nil
	case "jobs":
		return SectionTypeJobs, nil
	case "submit":
		return SectionTypeSubmit, nil
	default:
		return SectionInvalid, fmt.Errorf("invalid section type: %s", s)
	}
}
