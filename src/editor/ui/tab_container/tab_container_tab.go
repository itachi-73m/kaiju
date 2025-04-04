package tab_container

import (
	"kaiju/engine/ui/markup/document"
	"kaiju/engine/ui"
	"weak"
)

type TabContainerTab struct {
	Id      string
	Label   string
	parent  weak.Pointer[TabContainer]
	content TabContent
}

func NewTab(content TabContent) TabContainerTab {
	return TabContainerTab{
		Label:   content.TabTitle(),
		content: content,
	}
}

func (t *TabContainerTab) DragUpdate() {}

func (t *TabContainerTab) Reload(uiMan *ui.Manager, root *document.Element) {
	t.parent.Value().host.CreatingEditorEntities()
	t.content.Reload(uiMan, root)
	t.parent.Value().host.DoneCreatingEditorEntities()
}

func (t *TabContainerTab) Destroy() { t.content.Destroy() }
