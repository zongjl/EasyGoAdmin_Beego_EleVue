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

package controllers

import (
	"easygoadmin/app/models"
	"easygoadmin/conf"
	"easygoadmin/utils"
	"easygoadmin/utils/common"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gstr"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"regexp"
	"time"
)

var ConfigWeb = new(ConfigWebController)

type ConfigWebController struct {
	BaseController
}

func (ctl *ConfigWebController) Index() {
	// 获取配置列表
	configData := make([]models.Config, 0)
	orm.NewOrm().QueryTable(new(models.Config)).Filter("mark", 1).All(&configData)
	configList := make([]map[string]interface{}, 0)
	for _, v := range configData {
		item := make(map[string]interface{})
		item["config_id"] = v.Id
		item["config_name"] = v.Name

		// 查询配置项列表
		itemData := make([]models.ConfigData, 0)
		orm.NewOrm().QueryTable(new(models.ConfigData)).Filter("config_id", v.Id).Filter("status", 1).Filter("mark", 1).OrderBy("sort").All(&itemData)
		itemList := make([]map[string]interface{}, 0)
		for _, v := range itemData {
			item := make(map[string]interface{})
			item["id"] = v.Id
			item["title"] = v.Title
			item["code"] = v.Code
			item["value"] = v.Value
			item["type"] = v.Type

			if v.Type == "checkbox" {
				// 复选框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, val := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(val, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				// 选择项
				item["optionsList"] = optionsList
				// 选择值
				item["value"] = gstr.Split(v.Value, ",")

			} else if v.Type == "radio" {
				// 单选框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, v := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				item["optionsList"] = optionsList

			} else if v.Type == "select" {
				// 下拉选择框
				re := regexp.MustCompile(`\r?\n`)
				list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
				optionsList := make(map[string]string)
				for _, v := range list {
					re2 := regexp.MustCompile(`:|：|\s+`)
					item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
					optionsList[item[0]] = item[1]
				}
				item["optionsList"] = optionsList
			} else if v.Type == "image" {
				// 单图片
				item["value"] = utils.GetImageUrl(v.Value)
			} else if v.Type == "images" {
				// 多图片
				list := gstr.Split(v.Value, ",")
				itemList := make([]string, 0)
				for _, v := range list {
					// 图片地址
					item := utils.GetImageUrl(v)
					itemList = append(itemList, item)
				}
				item["value"] = itemList
			}
			itemList = append(itemList, item)
		}
		item["itemList"] = itemList
		// 加入数组
		configList = append(configList, item)
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: configList,
	})
}

func (ctl *ConfigWebController) Save() {
	if utils.AppDebug() {
		// 返回结果
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "演示环境，暂无权限操作",
		})
		return
	}
	// key：string类型，value：interface{}  类型能存任何数据类型
	var jsonObj map[string]interface{}
	data := ctl.Ctx.Input.RequestBody
	json.Unmarshal(data, &jsonObj)
	for key, val := range jsonObj {
		// 数组处理
		if reflect.ValueOf(val).Kind() == reflect.Slice {
			if reflect.TypeOf(val).String() == "[]interface {}" {
				// 初始化URL数组
				item := make([]string, 0)
				for _, v := range val.([]interface{}) {
					value := gconv.String(v)
					// 判断是否http(s)开头
					if gstr.SubStr(value, 0, 4) == "http" ||
						gstr.SubStr(value, 0, 5) == "https" {
						// 图片地址
						if gstr.Contains(value, conf.CONFIG.EGAdmin.Image) {
							url, _ := utils.SaveImage(value, "config")
							item = append(item, url)
						}
					} else {
						// 复选框处理
						item = append(item, value)
					}
				}
				// 逗号拼接
				val = gstr.Join(item, ",")
			}
		} else {
			// 图片处理
			if gstr.Contains(gconv.String(val), "http://") ||
				gstr.Contains(gconv.String(val), "https://") {
				// 图片地址
				if gstr.Contains(gconv.String(val), conf.CONFIG.EGAdmin.Image) {
					val, _ = utils.SaveImage(gconv.String(val), "config")
				}
			}
		}

		// 查询记录
		var entity models.ConfigData
		err := orm.NewOrm().QueryTable(new(models.ConfigData)).Filter("code", key).One(&entity)
		if err != nil {
			continue
		}
		entity.Value = gconv.String(val)
		entity.UpdateUser = utils.Uid(ctl.Ctx)
		entity.UpdateTime = time.Now()
		entity.Update()
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "保存成功",
	})
}
