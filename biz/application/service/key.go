package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/consts"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/mapper/key"
	"github.com/xh-polaris/openapi-user/biz/infrastructure/util"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type IKeyService interface {
	CreateKey(ctx context.Context, req *user.CreateKeyReq) (*user.CreateKeyResp, error)
	GetKey(ctx context.Context, req *user.GetKeysReq) (*user.GetKeysResp, error)
	UpdateKey(ctx context.Context, req *user.UpdateKeyReq) (*user.UpdateKeyResp, error)
	UpdateHosts(ctx context.Context, req *user.UpdateHostsReq) (*user.UpdateHostsResp, error)
	RefreshKey(ctx context.Context, req *user.RefreshKeyReq) (*user.RefreshKeyResp, error)
	DeleteKey(ctx context.Context, req *user.DeleteKeyReq) (*user.DeleteKeyResp, error)
}

type KeyService struct {
	KeyMongoMapper *key.MongoMapper
}

var KeyServiceSet = wire.NewSet(
	wire.Struct(new(KeyService), "*"),
	wire.Bind(new(IKeyService), new(*KeyService)),
)

func (s *KeyService) CreateKey(ctx context.Context, req *user.CreateKeyReq) (*user.CreateKeyResp, error) {
	userId := req.UserId
	// 名字初始化
	name := req.Name
	if name == "" { // 默认key名字
		name = consts.DefaultKeyName + time.Now().String()
	}
	// 记录域名白名单
	host := req.Hosts
	// 获取现在的时间
	now := time.Now()
	// 生成id
	id := primitive.NewObjectID()
	// 生成密钥内容
	content, err := util.GetKeyManager().IssueKey(id.Hex(), userId, now)

	newKey := &key.Key{
		ID:         id,
		UserId:     userId,
		Name:       name,
		Content:    content,
		Status:     0,
		Hosts:      host,
		Timestamp:  now,
		ExpireTime: now.Add(consts.DefaultExpire),
		CreateTime: now,
		UpdateTime: now,
	}
	err = s.KeyMongoMapper.Insert(ctx, newKey)
	if err != nil {
		return &user.CreateKeyResp{
			Done: false,
			Msg:  "签发密钥失败" + err.Error(),
		}, err
	}
	return &user.CreateKeyResp{
		Done: true,
		Msg:  "签发密钥成功",
	}, nil
}

func (s *KeyService) GetKey(ctx context.Context, req *user.GetKeysReq) (*user.GetKeysResp, error) {
	data, total, err := s.KeyMongoMapper.FindAndCount(ctx, req.UserId, req.PaginationOptions)
	if err != nil {
		return nil, err
	}
	var keys []*user.Key
	for _, val := range data {
		k := &user.Key{}
		err = copier.Copy(k, val)
		if err != nil {
			return nil, err
		}
		k.Id = val.ID.Hex()
		k.CreateTime = val.CreateTime.Unix()
		k.UpdateTime = val.UpdateTime.Unix()
		k.Timestamp = val.Timestamp.Unix()
		k.ExpireTime = val.ExpireTime.Unix()
		k.Status = user.KeyStatus(val.Status)
		keys = append(keys, k)
	}
	return &user.GetKeysResp{
		Keys:  keys,
		Total: total,
	}, nil
}

func (s *KeyService) UpdateKey(ctx context.Context, req *user.UpdateKeyReq) (*user.UpdateKeyResp, error) {
	id := req.Id
	name := req.Name
	status := req.Status
	timestamp := req.Timestamp
	expireTime := req.ExpireTime
	oldKey, err := s.KeyMongoMapper.FindOne(ctx, id)
	if err != nil || oldKey == nil {
		return &user.UpdateKeyResp{
			Done: false,
			Msg:  "key不存在或已删除",
		}, err
	}
	if name != nil {
		oldKey.Name = *name
	}
	if status != nil {
		oldKey.Status = int(status.Number())
	}
	if timestamp != nil {
		oldKey.Timestamp = time.Unix(*timestamp, 0)
	}
	if expireTime != nil {
		oldKey.ExpireTime = oldKey.ExpireTime.Add(time.Duration(*expireTime * 1000000000)) // expireTime以秒为单位
	}
	err = s.KeyMongoMapper.Update(ctx, oldKey)
	if err != nil {
		return &user.UpdateKeyResp{
			Done: false,
			Msg:  "更新失败",
		}, err
	}
	return &user.UpdateKeyResp{
		Done: true,
		Msg:  "success",
	}, nil
}

func (s *KeyService) UpdateHosts(ctx context.Context, req *user.UpdateHostsReq) (*user.UpdateHostsResp, error) {
	id := req.Id
	hosts := req.Hosts
	oldKey, err := s.KeyMongoMapper.FindOne(ctx, id)
	if err != nil || oldKey == nil {
		return &user.UpdateHostsResp{
			Done: false,
			Msg:  "密钥不存在或已删除",
		}, err
	}
	oldKey.Hosts = hosts
	err = s.KeyMongoMapper.Update(ctx, oldKey)
	if err != nil {
		return &user.UpdateHostsResp{
			Done: false,
			Msg:  "更新失败",
		}, err
	}
	return &user.UpdateHostsResp{
		Done: true,
		Msg:  "更新成功",
	}, nil
}

func (s *KeyService) RefreshKey(ctx context.Context, req *user.RefreshKeyReq) (*user.RefreshKeyResp, error) {
	id := req.Id
	oldKey, err := s.KeyMongoMapper.FindOne(ctx, id)
	if err != nil || oldKey == nil {
		return &user.RefreshKeyResp{
			Done: false,
			Msg:  "密钥不存在或已删除",
		}, nil
	}
	updateTime := time.Now()
	oldKey.Content, err = util.GetKeyManager().IssueKey(oldKey.ID.Hex(), oldKey.UserId, updateTime)
	if err != nil {
		return &user.RefreshKeyResp{
			Done: false,
			Msg:  "密钥生成失败",
		}, err
	}
	err = s.KeyMongoMapper.Update(ctx, oldKey)
	if err != nil {
		return &user.RefreshKeyResp{
			Done: false,
			Msg:  "刷新密钥失败",
		}, err
	}
	return &user.RefreshKeyResp{
		Done: true,
		Msg:  "刷新密钥成功",
	}, nil
}

func (s *KeyService) DeleteKey(ctx context.Context, req *user.DeleteKeyReq) (*user.DeleteKeyResp, error) {
	id := req.Id
	err := s.KeyMongoMapper.Delete(ctx, id)
	if err != nil {
		return &user.DeleteKeyResp{
			Done: false,
			Msg:  "删除密钥失败",
		}, err
	}
	return &user.DeleteKeyResp{
		Done: true,
		Msg:  "删除密钥成功",
	}, nil
}

func (s *KeyService) GetKeyForCheck(ctx context.Context, req *user.GetKeyForCheckReq) (*user.GetKeyForCheckResp, error) {
	content := req.Content
	host := req.Host
	timestamp := req.Timestamp

	// 根据密钥内容获取密钥
	k, err := s.KeyMongoMapper.FindOneByContent(ctx, content)
	if err != nil {
		return &user.GetKeyForCheckResp{
			Id:     "",
			UserId: "",
			Check:  false,
			Msg:    "查询密钥失败",
		}, err
	}

	// 解析密钥
	id, userId, freshTime, err := util.GetKeyManager().ParseKey(content)
	if err != nil {
		return &user.GetKeyForCheckResp{
			Id:     "",
			UserId: "",
			Check:  false,
			Msg:    "密钥解析失败",
		}, err
	}

	// 校对密钥正确性
	if k.ID.Hex() != id || k.UserId != userId || k.UpdateTime != freshTime {
		return &user.GetKeyForCheckResp{
			Id:     k.ID.Hex(),
			UserId: k.UserId,
			Check:  false,
			Msg:    "密钥权限校验失败",
		}, consts.ErrParse
	}

	// 校验密钥是否继续生效
	if k.Status != consts.EffectStatus || timestamp > k.ExpireTime.Unix() {
		return &user.GetKeyForCheckResp{
			Id:     k.ID.Hex(),
			UserId: k.UserId,
			Check:  false,
			Msg:    "密钥不可用或已过期",
		}, consts.ErrKeyUnavailable
	}

	// 校验域名是否符合要求
	if !contains(k.Hosts, host) {
		return &user.GetKeyForCheckResp{
			Id:     k.ID.Hex(),
			UserId: k.UserId,
			Check:  false,
			Msg:    "域名未被纳入白名单",
		}, consts.ErrHostUnavailable
	}

	// 校验通过，更新密钥最新的调用时间
	err = s.KeyMongoMapper.UpdateWithTimestamp(ctx, k, timestamp)
	if err != nil {
		return &user.GetKeyForCheckResp{
			Id:     k.ID.Hex(),
			UserId: k.UserId,
			Check:  true,
			Msg:    "最近调用时间更新失败",
		}, consts.ErrUpdate
	}

	return &user.GetKeyForCheckResp{
		Id:     k.ID.Hex(),
		UserId: k.UserId,
		Check:  true,
		Msg:    "校验成功",
	}, nil
}

// contains 检查给定的 host 是否在 hosts 切片中
func contains(hosts []string, host string) bool {
	for _, h := range hosts {
		if h == host {
			return true
		}
	}
	return false
}
