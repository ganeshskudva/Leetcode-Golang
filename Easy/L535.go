package Easy

type Codec struct {
	mp map[string]string
}

func Constructor() Codec {
	return Codec{
		mp: make(map[string]string),
	}
}

func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	res := hex.EncodeToString(NewSHA256([]byte(longUrl)))
	this.mp[res] = longUrl
    
    return res
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.mp[shortUrl]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
