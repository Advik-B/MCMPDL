package curseforge

import (
	badger "github.com/dgraph-io/badger/v4"
	"io"
	"net/http"
	"net/url"
)

type API struct {
	apiKey      string
	user_agent  string
	baseUrl     string
	using_cache bool
	cache       *badger.DB
	memcache    bool
}

func NewAPI(apiKey string, use_cache bool, use_memory_cache bool) (*API, error) {
	newAPI := &API{
		apiKey:      apiKey,
		user_agent:  "",
		baseUrl:     "https://api.curseforge.com/",
		using_cache: use_cache,
		cache:       nil,
		memcache:    use_memory_cache,
	}
	err := newAPI.init()
	if err != nil {
		return nil, err
	}
	return newAPI, nil
}

func (self *API) init() error {
	if self.using_cache {
		if self.memcache {
			return self.initMemCache()
		} else {
			return self.initDiskCache()
		}
	}
	return nil
}

func (self *API) initMemCache() error {
	options := badger.DefaultOptions("").WithInMemory(true)
	db, err := badger.Open(options)
	if err != nil {
		return err
	}
	self.cache = db
	return nil
}

func (self *API) initDiskCache() error {
	options := badger.DefaultOptions("curseforge_cache")
	db, err := badger.Open(options)
	if err != nil {
		return err
	}
	self.cache = db
	return nil
}

func (self *API) SetUserAgent(user_agent string) {
	self.user_agent = user_agent
}

func (self *API) Fetch(path string) (data []byte, err error) {
	url_to_fetch, err := url.JoinPath(self.baseUrl, path)
	if err != nil {
		return nil, err
	}

	// Check cache first
	if self.using_cache {
		cachedData, err := self.getFromCache(url_to_fetch)
		if err == nil {
			return cachedData, nil // Return cached response
		}
	}

	// Fetch from API
	data, err = self.FetchRaw(url_to_fetch)
	if err != nil {
		return nil, err
	}

	// Store response in cache
	if self.using_cache {
		_ = self.saveToCache(url_to_fetch, data)
	}

	return data, nil
}

// Retrieve data from cache
func (self *API) getFromCache(key string) ([]byte, error) {
	var value []byte
	err := self.cache.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = append([]byte{}, val...) // Copy data
			return nil
		})
	})
	return value, err
}

// Save data to cache
func (self *API) saveToCache(key string, value []byte) error {
	return self.cache.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
}

func (self *API) FetchRaw(url string) (data []byte, err error) {
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if self.user_agent != "" {
		res.Header.Set("User-Agent", self.user_agent)
	}
	res.Header.Set("Accept", "application/json")
	res.Header.Set("x-api-key", self.apiKey)

	client := &http.Client{}

	resp, err := client.Do(res)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
