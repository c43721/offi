package cache

type Cache interface {
	GetLogs(matchId int) (LogSet, error)
	SetLogs(matchId int, match *LogSet) error

	SetLogError(matchId int, err error) error
	CheckLogError(matchId int) error

	DeleteLogs(matchId int) (*LogSet, error)

	GetPlayer(playerID string) (string, error)
	SetPlayer(playerID, steamID string) error

	GetMatch(logId int) (MatchPage, error)
	SetMatch(logIds []int, matchPage *MatchPage) error

	GetAllKeys(hashKey string) ([]string, error)
}
