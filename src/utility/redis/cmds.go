package redis

/**************** connections ******************/
func (c *Conn) Auth(password string) (bool, error) {
	v, e := c.Call("AUTH", password)
	if e != nil {
		Debug("AUTH failed:"+e.Error)
	}
}
