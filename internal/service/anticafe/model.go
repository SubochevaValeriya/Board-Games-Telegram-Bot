package anticafe

type anticafeConf struct {
	anticafeName     string
	link             string
	googleSheetsInfo *googleSheetsInfo
	vkTopicInfo      *vkTopicInfo
}

type AnticafeGameInStock struct {
	AnticafeName string
	IsInStock    bool
	Link         string
}

type googleSheetsInfo struct {
	id         string
	sheetTitle string
	column     int
}

type vkTopicInfo struct {
	groupID string
	topicID string
}
