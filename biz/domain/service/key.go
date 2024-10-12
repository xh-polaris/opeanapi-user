package service

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/openapi/user"
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/config"
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/consts"
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/mapper/key"
	"github.com/xhpolaris/opeanapi-user/biz/infrastructure/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type KeyService struct {
	Config         config.Config
	KeyMongoMapper key.IMongoMapper
}

func (s *KeyService) CreateKey(ctx context.Context, req *user.CreateKeyReq) (*user.CreateKeyResp, error) {
	userId := req.User.UserId
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
	data, total, err := s.KeyMongoMapper.FindAndCount(ctx, req.User.UserId, req.PaginationOptions)
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
		k.DeleteTime = val.DeleteTime.Unix()
		keys = append(keys, k)
	}
	return &user.GetKeysResp{
		Keys:  keys,
		Total: total,
	}, nil
}

func (s *KeyService) UpdateKey(ctx context.Context, req *user.UpdateKeyReq) (*user.UpdateKeyResp, error) {
	id := req.Id
	name := *req.Name
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
	if name != "" {
		oldKey.Name = name
	}
	if status != nil {
		oldKey.Status = int(status.Number())
	}
	if timestamp != nil {
		oldKey.Timestamp = time.Unix(*timestamp, 0)
	}
	if expireTime != nil {
		oldKey.ExpireTime = oldKey.ExpireTime.Add(time.Duration(*expireTime))
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
