package service

type ServiceIf interface {
	Beat(name string)
	SetWarningTimeout(name string, timeout int)
	SetErrorTimeout(name string, timeout int)

	CreateOrUpdateView(name string, regexp string)
}

const (
	ACTION_SET_WARN = "set-warn"
	ACTION_SET_ERR  = "set-err"
	ACTION_BEAT     = "beat"
)

const (
	ACTION_REFRESH_VIEW = "refresh-view"
	ACTION_UPSERT_VIEW  = "upsert-view"
)

type serviceCmd struct {
	Action  string
	Service string
	Value   int
}

type viewCmd struct {
	Action string
	View   string
	Regexp string
}

type client struct {
	svcs *Services
}

func (c *client) Beat(name string) {
	c.svcs.serviceCmdChan <- &serviceCmd{
		Action:  ACTION_BEAT,
		Service: name,
		Value:   1,
	}
}

func (c *client) SetWarningTimeout(name string, timeout int) {
	c.svcs.serviceCmdChan <- &serviceCmd{
		Action:  ACTION_SET_WARN,
		Service: name,
		Value:   timeout,
	}
}
func (c *client) SetErrorTimeout(name string, timeout int) {
	c.svcs.serviceCmdChan <- &serviceCmd{
		Action:  ACTION_SET_ERR,
		Service: name,
		Value:   timeout,
	}
}

func (c *client) CreateOrUpdateView(name string, regexp string) {
	c.svcs.viewCmdChan <- &viewCmd{
		Action: ACTION_UPSERT_VIEW,
		View:   name,
		Regexp: regexp,
	}
}

func (svcs *Services) GetClient() ServiceIf {
	return &client{svcs: svcs}
}
