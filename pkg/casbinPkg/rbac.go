package casbinpkg

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/diy0663/gohub/pkg/database"
)

var (
	once                 sync.Once
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
)

type PermissionInfo struct {
	Path   string
	Method string
}

// CREATE TABLE `casbin_rule` (
// 	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// 	`ptype` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL, 存 p
// 	`v0` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,    存 角色id
// 	`v1` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,	  存  路由path
// 	`v2` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,    存  路由方法 method
// 	`v3` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`v4` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	`v5` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

func GetCasbinRBAC() *casbin.SyncedCachedEnforcer {
	once.Do(func() {

		adapter, err := gormadapter.NewAdapterByDB(database.DB)
		if err != nil {
			fmt.Println("适配数据库失败请检查casbin表", err)
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			fmt.Println("初始化模型失败", err)
		}
		// 给全局变量赋值
		syncedCachedEnforcer, err = casbin.NewSyncedCachedEnforcer(m, adapter)
		if err != nil {
			fmt.Println("初始化casbin失败", err)
		}
		syncedCachedEnforcer.SetExpireTime(time.Minute * 10)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}

// 移除指定角色的指定权限
func RemovePermissions(roleId int, persions ...string) bool {
	result, err := GetCasbinRBAC().RemoveFilteredPolicy(roleId, persions...)
	if err != nil {
		panic(err)
	} else {
		return result
	}

}

// 根据角色id获取对应的权限
func GetPermissionByRoleId(roleId int) []PermissionInfo {
	var permissions []PermissionInfo
	roleIdString := strconv.Itoa(roleId)
	result := GetCasbinRBAC().GetFilteredPolicy(0, roleIdString)

	for _, v := range result {
		permissions = append(permissions, PermissionInfo{
			Path:   v[0],
			Method: v[1],
		})
	}

	return permissions
}

func UpdatePermission(oldPath, oldMethod, newPath, newMethod string) error {
	err := database.DB.Model(&gormadapter.CasbinRule{}).Where("v1= ? and v2=?", oldPath, oldMethod).Updates(
		map[string]interface{}{"v1": newPath, "v2": newMethod}).Error

	if err != nil {
		return err
	}
	err = GetCasbinRBAC().LoadPolicy()
	if err != nil {
		return err
	}
	return nil
}

func UpdatePermissionByRoleId(roleId int, persions []PermissionInfo) error {
	// 先清空原来的权限
	roleIdStr := strconv.Itoa(roleId)
	_, err := GetCasbinRBAC().RemoveFilteredPolicy(0, roleIdStr)
	if err != nil {
		return errors.New("删除旧权限失败")
	}
	unRepeatPermissions := [][]string{}
	uniqueMap := make(map[string]bool)
	for _, p := range persions {
		key := roleIdStr + ":" + p.Path + ":" + p.Method
		if _, ok := uniqueMap[key]; !ok {
			// 去重
			uniqueMap[key] = true
			// 权限三元组, 在这里设定了顺序, v0是角色id, v1是路由, v2是方法
			unRepeatPermissions = append(unRepeatPermissions, []string{roleIdStr, p.Path, p.Method})
		}
	}
	success, _ := GetCasbinRBAC().AddPolicies(unRepeatPermissions)
	if !success {
		return errors.New("添加权限失败")
	}
	return nil

}
