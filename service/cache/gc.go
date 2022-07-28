package cache

import "time"

/*
	GC Launcher
*/
func (c *Cache) StartGC() {
	go c.GC()
}

/*
	Garbage Collector
*/
func (c *Cache) GC() {
	for {
		// wait till time in cleanupInterval
		<-time.After(c.cleanupInterval)
		if c.items == nil {
			return
		}
		// Find elements with expired lifetine and remove from storage
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

/*
	return list of EXPIRED KEYS
*/
func (c *Cache) expiredKeys() (keys []string) {

	c.RLock()

	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}

	return
}

/*
	Remove elements from LIST
*/
func (c *Cache) clearItems(keys []string) {

	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
