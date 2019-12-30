package aaa

import (
	"strconv"

	"github.com/valyala/bytebufferpool"

	"github.com/tsingson/goums/pkg/vtils"
)

var (
	liveKeyByte = []byte("-gslb") // 直播 gslb
	authKeyByte = []byte("-jwt")
	regKeyByte  = []byte("-reg")
)

// sn 禁用与解禁

// userID 缓存主键
func idCacheKey(userID []byte) []byte {
	bb := bytebufferpool.Get()
	_, _ = bb.Write(userID)
	_, _ = bb.Write(authKeyByte)
	idKey := bb.Bytes()
	bytebufferpool.Put(bb)
	return idKey
}

func idCacheKeyString(id string) []byte {
	bb := bytebufferpool.Get()
	_, _ = bb.WriteString(id)
	_, _ = bb.Write(authKeyByte)
	idKey := bb.Bytes()
	bytebufferpool.Put(bb)
	return idKey
}

// 用户访问权限主键, 如果存在缓存, 则不允许访问( 禁用) , 如果不存在缓存, 则可以访问( 解禁)
func idAccessKeyString(userID string) []byte {
	return vtils.S2B(userID)
}

func snCacheKey(sn []byte) []byte {
	bb := bytebufferpool.Get()
	_, _ = bb.Write(sn)
	_, _ = bb.Write(authKeyByte)
	snKey := bb.Bytes()
	bytebufferpool.Put(bb)
	return snKey
}

// SnCacheKeyString  key
func snCacheKeyString(sn string) []byte {
	bb := bytebufferpool.Get()
	_, _ = bb.WriteString(sn)
	_, _ = bb.Write(authKeyByte)
	snKey := bb.Bytes()
	bytebufferpool.Put(bb)
	return snKey
}

// SnCacheRegKeyString key
func snCacheRegKeyString(sn string) []byte {
	bb := bytebufferpool.Get()
	_, _ = bb.WriteString(sn)
	_, _ = bb.Write(regKeyByte)
	snKey := bb.Bytes()
	bytebufferpool.Put(bb)
	return snKey
}

// SerialCacheDel del
func (s *Aaa) serialCacheDel(serial string) {
	s.cache.Del(snCacheKeyString(serial))
}

// SerialCacheSet del
func (s *Aaa) serialCacheSet(serial string, in []byte) bool {
	return s.cache.Set(snCacheKeyString(serial), in)
}

// SerialCacheGet del
func (s *Aaa) serialCacheGet(serial string) ([]byte, bool) {
	return s.cache.Get(snCacheKeyString(serial))
}

// SerialCacheRegDel del
func (s *Aaa) serialCacheRegDel(serial string) {
	s.cache.Del(snCacheRegKeyString(serial))
}

// SerialCacheRegSet del
func (s *Aaa) serialCacheRegSet(serial string, in []byte) bool {
	return s.cache.Set(snCacheRegKeyString(serial), in)
}

// SerialCacheRegGet del
func (s *Aaa) serialCacheRegGet(serial string) (v []byte, exists bool) {
	return s.cache.Get(snCacheRegKeyString(serial))
}

// SnAccessKeyString key
func snAccessKeyString(sn string) []byte {
	return vtils.S2B(sn)
}

// SnAccessKeyString key
func snAccessKey(sn []byte) []byte {
	return sn
}

// SerialAccessGet get
func (s *Aaa) serialAccessGet(serial string) bool {
	_, ck := s.cache.Get(snAccessKeyString(serial))
	return ck
}

// SerialAccessSet set
func (s *Aaa) serialAccessSet(serial []byte) bool {
	return s.cache.Set(snAccessKey(serial), []byte("1"))
}

// SerialAccessDel set
func (s *Aaa) serialAccessDel(serial []byte) {
	s.cache.Del(snAccessKey(serial))
}

func int32String(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}
