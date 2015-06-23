package model

type ResourceApi struct {
	resourceManager interface{}
}

func (this *ResourceApi) CreateEntryMetric() {}
func (this *ResourceApi) DeleteEntryMetric() {}
func (this *ResourceApi) UpdateEntryMetric() {}
func (this *ResourceApi) GetEntryMetric() {}

func (this *ResourceApi) ClearFeed() {}
func (this *ResourceApi) ReorderFeed() {}
func (this *ResourceApi) CreateFeedEntry() {}
func (this *ResourceApi) DeleteFeedEntry() {}
func (this *ResourceApi) UpdateFeedEntry() {}

func NewResourceApi(resourceManager interface{}) *ResourceApi {
	return &ResourceApi{resourceManager}
}
