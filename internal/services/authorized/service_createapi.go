package authorized

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo/authorized_api_repo"
	"github.com/xinliangnote/go-gin-api/internal/api/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type CreateAuthorizedAPIData struct {
	BusinessKey string `json:"business_key"` // 调用方key
	Method      string `json:"method"`       // 请求方法
	API         string `json:"api"`          // 请求地址
}

func (s *service) CreateAPI(ctx core.Context, authorizedAPIData *CreateAuthorizedAPIData) (id int32, err error) {
	model := authorized_api_repo.NewModel()
	model.BusinessKey = authorizedAPIData.BusinessKey
	model.Method = authorizedAPIData.Method
	model.Api = authorizedAPIData.API
	model.CreatedUser = ctx.UserName()
	model.IsDeleted = -1

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}

	s.cache.Del(configs.RedisKeyPrefixSignature+authorizedAPIData.BusinessKey, redis.WithTrace(ctx.Trace()))
	return
}