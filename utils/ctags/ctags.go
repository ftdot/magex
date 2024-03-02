package ctags

import "github.com/ftdot/magex/utils"

type CTags struct {
	tags []string
}

func New() *CTags {
	return &CTags{
		tags: []string{},
	}
}

func MergeCTags(tags... *CTags) *CTags {
	ct := New()

	for _, t := range tags {
		ct.AddTags(t.tags)
	}

	return ct
}

////

func (t *CTags) HasTag(tag string) bool {
	return utils.SliceContainsItem(t.tags, tag)
}

func (t *CTags) AddTag(tag string) {
	if utils.SliceContainsItem(t.tags, tag) {
		return
	}
	t.tags = append(t.tags, tag)
}

func (t *CTags) AddTags(tags []string) {
	for _, tag := range tags {
		t.AddTag(tag)
	}
}
