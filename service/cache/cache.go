package cache

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/truecoder34/l0-wb-nats-service/service/models"
	"gorm.io/gorm"
)

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item
}

type Item struct {
	Value      models.Transaction
	Created    time.Time
	Expiration int64
}

/*
	Initialize CACHE
*/
func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)
	cache := Cache{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}
	// If clean up inteval >  0, run garbage collector
	if cleanupInterval > 0 {
		cache.StartGC()
	}
	return &cache
}

/*
	Add data to CACHE
*/
func (c *Cache) Set(key string, value models.Transaction, duration time.Duration) {
	var expiration int64
	// if life time == 0 - use default value
	if duration == 0 {
		duration = c.defaultExpiration
	}
	// Set experation time for cache
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}
	c.Lock()
	defer c.Unlock()
	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}
}

/*
	Load data to cache from dataBase
*/
func (c *Cache) Load(db *gorm.DB) {
	var transaction models.Transaction
	transactions, err := transaction.FindAllTransactions(db)
	if err != nil {
		log.Print(err)
	}
	for _, transaction := range *transactions {
		c.Set(transaction.ID.String(), transaction, 50000*time.Minute)
	}

}

/*
	Get value from CACHE
	TODO : by ID
*/
func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	item, found := c.items[key]
	// key was not found
	if !found {
		return nil, false
	}
	// check fir expiration time , else - it is endless
	if item.Expiration > 0 {
		// if cache is old - return nil
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}
	return item.Value, true
}

/*
	Get all values from cache
*/
func (c *Cache) GetAll() (map[string]Item, error) {
	var err error
	c.RLock()
	defer c.RUnlock()
	items := c.items

	return items, err
}

/*
	Delete from CACHE
*/
func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()
	if _, found := c.items[key]; !found {
		return errors.New("Key not found")
	}
	delete(c.items, key)
	return nil
}
