package marionette_client

type ActionCollection struct {
	idx     int16
	action  map[int16]map[string]interface{}
}

func (ac ActionCollection) Append(command string, value interface{}) {
	ac.idx = ac.idx + 1
	c := map[string]interface{}{command: value }
	ac.action[ac.idx] = c
}

type Action struct {
	el      *WebElement
	chain   ActionCollection
}

func NewAction(e *WebElement) *Action {
	return &Action{el: e, chain: ActionCollection{}}
}


func (a *Action) Pause(milliseconds int) *Action {
	data := map[string]interface{}{}
	if milliseconds == 0  || milliseconds < 0 {
		data["duration"] = nil
	} else if milliseconds > 60000	{
		data["duration"] = 60000
	} else {
		data["duration"] = milliseconds
	}

	a.chain.Append("pause", data)

	return a
}

func (a *Action) KeyUp(milliseconds int) *Action {

	return a
}

