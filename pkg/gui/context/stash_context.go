package context

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type StashContext struct {
	*BasicViewModel[*models.StashEntry]
	*ListContextTrait
}

var _ types.IListContext = (*StashContext)(nil)

func NewStashContext(
	getModel func() []*models.StashEntry,
	view *gocui.View,
	getDisplayStrings func(startIdx int, length int) [][]string,

	onFocus func(...types.OnFocusOpts) error,
	onRenderToMain func(...types.OnFocusOpts) error,
	onFocusLost func() error,

	c *types.HelperCommon,
) *StashContext {
	viewModel := NewBasicViewModel(getModel)

	return &StashContext{
		BasicViewModel: viewModel,
		ListContextTrait: &ListContextTrait{
			Context: NewSimpleContext(NewBaseContext(NewBaseContextOpts{
				ViewName:   "stash",
				WindowName: "stash",
				Key:        STASH_CONTEXT_KEY,
				Kind:       types.SIDE_CONTEXT,
				Focusable:  true,
			}), ContextCallbackOpts{
				OnFocus:        onFocus,
				OnFocusLost:    onFocusLost,
				OnRenderToMain: onRenderToMain,
			}),
			list:              viewModel,
			viewTrait:         NewViewTrait(view),
			getDisplayStrings: getDisplayStrings,
			c:                 c,
		},
	}
}

func (self *StashContext) GetSelectedItemId() string {
	item := self.GetSelected()
	if item == nil {
		return ""
	}

	return item.ID()
}

func (self *StashContext) CanRebase() bool {
	return false
}

func (self *StashContext) GetSelectedRefName() string {
	item := self.GetSelected()

	if item == nil {
		return ""
	}

	return item.RefName()
}

func (self *StashContext) GetSelectedDescription() string {
	item := self.GetSelected()

	if item == nil {
		return ""
	}

	return item.Description()
}
