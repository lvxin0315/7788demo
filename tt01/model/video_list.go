package model

type VideoListResponse struct {
	StatusCode int         `json:"statusCode"`
	ItemList   []VideoItem `json:"itemList"`
	Cursor     string      `json:"cursor"`
	HasMore    bool        `json:"hasMore"`
}

type VideoItem struct {
	Id         string `json:"id"`
	Desc       string `json:"desc"`
	CreateTime int    `json:"createTime"`
	Video      struct {
		Id            string   `json:"id"`
		Height        int      `json:"height"`
		Width         int      `json:"width"`
		Duration      int      `json:"duration"`
		Ratio         string   `json:"ratio"`
		Cover         string   `json:"cover"`
		OriginCover   string   `json:"originCover"`
		DynamicCover  string   `json:"dynamicCover"`
		PlayAddr      string   `json:"playAddr"`
		DownloadAddr  string   `json:"downloadAddr"`
		ShareCover    []string `json:"shareCover"`
		ReflowCover   string   `json:"reflowCover"`
		Bitrate       int      `json:"bitrate"`
		EncodedType   string   `json:"encodedType"`
		Format        string   `json:"format"`
		VideoQuality  string   `json:"videoQuality"`
		EncodeUserTag string   `json:"encodeUserTag"`
		CodecType     string   `json:"codecType"`
		Definition    string   `json:"definition"`
		BitrateInfo   []struct {
			GearName    string `json:"GearName"`
			Bitrate     int    `json:"Bitrate"`
			QualityType int    `json:"QualityType"`
			PlayAddr    struct {
				Uri      string   `json:"Uri"`
				UrlList  []string `json:"UrlList"`
				DataSize int      `json:"DataSize"`
				UrlKey   string   `json:"UrlKey"`
				FileHash string   `json:"FileHash"`
				FileCs   string   `json:"FileCs"`
			} `json:"PlayAddr"`
			CodecType string `json:"CodecType"`
		} `json:"bitrateInfo"`
	} `json:"video"`
	Author struct {
		Id             string `json:"id"`
		UniqueId       string `json:"uniqueId"`
		Nickname       string `json:"nickname"`
		AvatarThumb    string `json:"avatarThumb"`
		AvatarMedium   string `json:"avatarMedium"`
		AvatarLarger   string `json:"avatarLarger"`
		Signature      string `json:"signature"`
		Verified       bool   `json:"verified"`
		SecUid         string `json:"secUid"`
		Secret         bool   `json:"secret"`
		Ftc            bool   `json:"ftc"`
		Relation       int    `json:"relation"`
		OpenFavorite   bool   `json:"openFavorite"`
		CommentSetting int    `json:"commentSetting"`
		DuetSetting    int    `json:"duetSetting"`
		StitchSetting  int    `json:"stitchSetting"`
		PrivateAccount bool   `json:"privateAccount"`
		IsADVirtual    bool   `json:"isADVirtual"`
	} `json:"author"`
	Music struct {
		Id          string `json:"id"`
		Title       string `json:"title"`
		PlayUrl     string `json:"playUrl"`
		CoverThumb  string `json:"coverThumb"`
		CoverMedium string `json:"coverMedium"`
		CoverLarge  string `json:"coverLarge"`
		AuthorName  string `json:"authorName"`
		Original    bool   `json:"original"`
		Duration    int    `json:"duration"`
		Album       string `json:"album"`
	} `json:"music"`
	Challenges []struct {
		Id            string `json:"id"`
		Title         string `json:"title"`
		Desc          string `json:"desc"`
		ProfileThumb  string `json:"profileThumb"`
		ProfileMedium string `json:"profileMedium"`
		ProfileLarger string `json:"profileLarger"`
		CoverThumb    string `json:"coverThumb"`
		CoverMedium   string `json:"coverMedium"`
		CoverLarger   string `json:"coverLarger"`
		IsCommerce    bool   `json:"isCommerce"`
	} `json:"challenges"`
	Stats struct {
		DiggCount    int `json:"diggCount"`
		ShareCount   int `json:"shareCount"`
		CommentCount int `json:"commentCount"`
		PlayCount    int `json:"playCount"`
	} `json:"stats"`
	DuetInfo struct {
		DuetFromId string `json:"duetFromId"`
	} `json:"duetInfo"`
	OriginalItem bool `json:"originalItem"`
	OfficalItem  bool `json:"officalItem"`
	TextExtra    []struct {
		AwemeId      string `json:"awemeId"`
		Start        int    `json:"start"`
		End          int    `json:"end"`
		HashtagName  string `json:"hashtagName"`
		HashtagId    string `json:"hashtagId"`
		Type         int    `json:"type"`
		UserId       string `json:"userId"`
		IsCommerce   bool   `json:"isCommerce"`
		UserUniqueId string `json:"userUniqueId"`
		SecUid       string `json:"secUid"`
		SubType      int    `json:"subType"`
	} `json:"textExtra"`
	Secret            bool `json:"secret"`
	ForFriend         bool `json:"forFriend"`
	Digged            bool `json:"digged"`
	ItemCommentStatus int  `json:"itemCommentStatus"`
	ShowNotPass       bool `json:"showNotPass"`
	Vl1               bool `json:"vl1"`
	ItemMute          bool `json:"itemMute"`
	AuthorStats       struct {
		FollowingCount int `json:"followingCount"`
		FollowerCount  int `json:"followerCount"`
		HeartCount     int `json:"heartCount"`
		VideoCount     int `json:"videoCount"`
		DiggCount      int `json:"diggCount"`
		Heart          int `json:"heart"`
	} `json:"authorStats"`
	PrivateItem    bool `json:"privateItem"`
	DuetEnabled    bool `json:"duetEnabled"`
	StitchEnabled  bool `json:"stitchEnabled"`
	ShareEnabled   bool `json:"shareEnabled"`
	IsAd           bool `json:"isAd"`
	DuetDisplay    int  `json:"duetDisplay"`
	StitchDisplay  int  `json:"stitchDisplay"`
	EffectStickers []struct {
		Name string `json:"name"`
		ID   string `json:"ID"`
	} `json:"effectStickers,omitempty"`
	StickersOnItem []struct {
		StickerType int      `json:"stickerType"`
		StickerText []string `json:"stickerText"`
	} `json:"stickersOnItem,omitempty"`
}
