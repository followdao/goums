package aaa

import (
	"strconv"
	"time"

	proto "github.com/liftbridge-io/liftbridge-api/go"
	"github.com/valyala/fastjson"
	"go.uber.org/zap"
)

// CacheTerminalAccess  cache
func (s *Aaa) CacheTerminalAccess(msg *proto.Message, err error) {
	log := s.log.Named("CacheTerminalAccess " + strconv.Itoa(time.Now().Nanosecond()))
	log.Info("----------------------------------------------------")
	if err != nil {
		log.Error("notify error", zap.Error(err))
	}

	var p fastjson.Parser

	v, er1 := p.ParseBytes(msg.Value)
	if err != nil {
		log.Error("json parse error", zap.Error(er1))
		return
	}
	if v != nil {

	}
}
