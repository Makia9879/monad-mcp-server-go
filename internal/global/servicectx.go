package global

import "context"

var SvcCtx *ServiceContext

type ServiceContext struct {
	RunningPathCtx   string
	ChromeCtx        context.Context
	ChromeCancelFunc context.CancelFunc
}

func NewServiceContext() *ServiceContext {
	return &ServiceContext{}
}

func (s *ServiceContext) Close() {
	if s.ChromeCancelFunc != nil {
		s.ChromeCancelFunc()
	}
}
