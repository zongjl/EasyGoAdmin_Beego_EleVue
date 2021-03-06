// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package vo

import (
	"easygoadmin/app/models"
	"time"
)

// 用户信息Vo
type UserListVo struct {
	models.User
	Birthday     string      `json:"birthday"`     // 性别
	GenderName   string      `json:"genderName"`   // 性别
	LevelName    string      `json:"levelName"`    // 职级
	PositionName string      `json:"positionName"` // 岗位
	DeptName     string      `json:"deptName"`     // 部门
	RoleIds      interface{} `json:"roleIds"`      // 角色ID
	RoleList     interface{} `json:"roleList"`     // 角色列表
	City         interface{} `json:"city"`         // 省市区
}

// 用户信息Vo
type UserInfoVo struct {
	Id         int       `json:"id"`
	Realname   string    `json:"realname"`
	Nickname   string    `json:"nickname"`
	Gender     int       `json:"gender"`
	Avatar     string    `json:"avatar"`
	Mobile     string    `json:"mobile"`
	Email      string    `json:"email"`
	Birthday   time.Time `json:"birthday"`
	DeptId     int       `json:"deptId"`
	LevelId    int       `json:"levelId"`
	PositionId int       `json:"positionId"`
	Address    string    `json:"address"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Intro      string    `json:"intro"`
	Status     int       `json:"status"`
	Note       string    `json:"note"`
	Sort       int       `json:"sort"`
	RoleIds    []int     `json:"roleIds"` // 角色ID
	City       []string  `json:"city"`    // 省市区
}

// 个人信息Vo
type ProfileInfoVo struct {
	Realname       string        `json:"realname"`       // 真实姓名
	Nickname       string        `json:"nickname"`       // 昵称
	Gender         int           `json:"gender"`         // 性别:1男 2女 3保密
	Avatar         string        `json:"avatar"`         // 头像
	Mobile         string        `json:"mobile"`         // 手机号码
	Email          string        `json:"email"`          // 邮箱地址
	City           []string      `json:"city"`           // 省市区
	Address        string        `json:"address"`        // 详细地址
	Intro          string        `json:"intro"`          // 个人简介
	Roles          []interface{} `json:"roles"`          // 用户角色
	Authorities    []interface{} `json:"authorities"`    // 用户权限
	PermissionList []string      `json:"permissionList"` // 权限列表
}
