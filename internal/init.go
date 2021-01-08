package internal

func Init(api SystemInfoAPI, listener Listener) *Core {
	core := &Core{
		systemInfoAPI: api,
		listener:      listener,
	}

	return core
}

func (c *Core) Start() error {
	err := c.listener.RegisterReceiver("SystemInfoAPI", c.systemInfoAPI)
	if err != nil {
		return err
	}

	return c.listener.Handle()
}
