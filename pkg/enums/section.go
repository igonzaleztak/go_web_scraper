package enums

import "fmt"

// SectionType defines the hacker news sections that are available to scrap
type SectionType string

const (
	SectionTypeNew      SectionType = "new" // newstories
	SectionTypePast     SectionType = "past"
	SectionTypeComments SectionType = "comments"
	SectionTypeAsk      SectionType = "ask"  // askstories
	SectionTypeShow     SectionType = "show" // showstories
	SectionTypeJobs     SectionType = "jobs" // jobstories
	SectionTypeSubmit   SectionType = "submit"
)

func (s *SectionType) String() string {
	return string(*s)
}

func (s *SectionType) ApiString() string {
	switch *s {
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

func (s *SectionType) Set(section string) error {
	switch section {
	case "new":
		*s = SectionTypeNew
		return nil
	case "past":
		*s = SectionTypePast
		return nil
	case "comments":
		*s = SectionTypeComments
		return nil
	case "ask":
		*s = SectionTypeAsk
		return nil
	case "show":
		*s = SectionTypeShow
		return nil
	case "jobs":
		*s = SectionTypeJobs
		return nil
	case "submit":
		*s = SectionTypeSubmit
		return nil
	default:
		return fmt.Errorf("invalid section type: %s. Must be %s, %s, %s, %s, %s, %s or %s", section, SectionTypeNew, SectionTypePast, SectionTypeComments, SectionTypeAsk, SectionTypeShow, SectionTypeJobs, SectionTypeSubmit)
	}
}

func (s *SectionType) Type() string {
	return "section"
}
