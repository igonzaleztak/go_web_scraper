package spaceFlight

import (
	"intelygenz/pkg/config"
	"intelygenz/pkg/logs"
	"intelygenz/pkg/schemas"
	"strings"
)

// splitStoriesByLength splits the stories into two categories: short and long. Short stories are those with a title length of less than 5 words. Long stories are those with a title length of more than 5 words.
func splitStoriesByLength(stories []schemas.SpaceFlightNews) (short DataSpaceFlight, long DataSpaceFlight) {
	logs.Logger.Info("[spaceflight.SplitStoriesByLength] splitting stories by length")
	short = make(DataSpaceFlight, 0)
	long = make(DataSpaceFlight, 0)

	for _, story := range stories {
		logs.Logger.Debugf("[spaceflight.SplitStoriesByLengths] checking story %d with title `%s`", story.ID, story.Title)
		if len(strings.Fields(story.Title)) < config.CmdFlags.NumWords {
			logs.Logger.Debugf("[spaceflight.SplitStoriesByLengths] story %d is short", story.ID)
			short = append(short, story)
		} else {
			logs.Logger.Debugf("[spaceflight.SplitStoriesByLengths] story %d is long", story.ID)
			long = append(long, story)
		}
	}
	logs.Logger.Infof("[spaceflight.SplitStoriesByLength] got %d short stories and %d long stories", len(short), len(long))
	return short, long
}
