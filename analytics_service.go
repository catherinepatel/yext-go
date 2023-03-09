package yext

const analyticsPath = "analytics/reports"

type AnalyticsService struct {
	client *Client
}

type AnalyticsFilters struct {
	StartDate                        *string   `json:"startDate"`
	EndDate                          *string   `json:"endDate"`
	LocationIds                      *[]string `json:"locationIds"`
	EntityIds                        *[]string `json:"entityIds"`
	EntityTypes                      *[]string `json:"entityType"`
	FolderId                         *int      `json:"folderId"`
	Countries                        *[]string `json:"countries"`
	LocationLabels                   *[]string `json:"locationLabels"`
	Platforms                        *[]string `json:"platforms"`
	GoogleActionType                 *[]string `json:"googleActionType"`
	CustomerActionType               *[]string `json:"customerActionType"`
	GoogleQueryType                  *[]string `json:"googleQueryType"`
	Hours                            *[]int    `json:"hours"`
	Ratings                          *[]int    `json:"ratings"`
	FrequentWords                    *[]string `json:"frequentWords"`
	Partners                         *[]int    `json:"partners"`
	ReviewLabels                     *[]int    `json:"reviewLabels"`
	PageTypes                        *[]string `json:"pageTypes"`
	ListingsLiveType                 *string   `json:"listingsLiveType"`
	QueryTemplate                    *[]string `json:"queryTemplate"`
	SearchEngine                     *[]string `json:"searchEngine"`
	Keyword                          *[]string `json:"keyword"`
	Competitor                       *[]string `json:"competitor"`
	MatchPosition                    *[]string `json:"matchPosition"`
	SearchResultType                 *[]string `json:"searchResultType"`
	MatchType                        *[]string `json:"matchType"`
	MinSearchFrequency               *int      `json:"minSearchFrequency"`
	MaxSearchFrequency               *int      `json:"maxSearchFrequency"`
	FoursquareCheckinType            *string   `json:"foursquareCheckinType"`
	FoursquareCheckinAge             *string   `json:"foursquareCheckinAge"`
	FoursquareCheckinGender          *string   `json:"foursquareCheckinGender"`
	InstagramContentType             *string   `json:"instagramContentType"`
	Age                              *[]string `json:"age"`
	Gender                           *string   `json:"gender"`
	FacebookImpressionType           *[]string `json:"facebookImpressionType"`
	FacebookStoryType                *[]string `json:"facebookStoryType"`
	AnswersExperience                *[]string `json:"ANSWERS_EXPERIENCE"`
	AnswersConfigurationVersionLabel *string   `json:"ANSWERS_CONFIGURATION_VERSION_LABEL"`
	AnswersTrafficType               *[]string `json:"ANSWERS_TRAFFIC_TYPE"`
}

type AnalyticsReportRequest struct {
	Metrics    []string          `json:"metrics"`
	Dimensions []string          `json:"dimensions"`
	Filters    *AnalyticsFilters `json:"filters"`
}

type AnalyticsReportResponse struct {
	Data []*AnalyticsData `json:"data"`
	Id   int              `json:"id"`
}

func (a *AnalyticsService) Create(req *AnalyticsReportRequest) (*AnalyticsReportResponse, *Response, error) {
	arr := &AnalyticsReportResponse{}
	r, err := a.client.DoRequestJSON("POST", analyticsPath, req, arr)
	return arr, r, err
}
