package sms

import (
	errors2 "errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	"iris-antd-admin/libs"
	"iris-antd-admin/models"
	"strings"
	"time"
)

var TencentSmsErrorInfo = map[string]string{
	"FailedOperation.ContainSensitiveWord":                            "短信内容中含有敏感词，请联系 sms helper。",
	"FailedOperation.FailResolvePacket":                               "请求包解析失败，通常情况下是由于没有遵守 API 接口说明规范导致的，请参考 请求包体解析1004错误详解。",
	"FailedOperation.InsufficientBalanceInSmsPackage":                 "套餐包余量不足，请 购买套餐包。",
	"FailedOperation.JsonParseFail":                                   "解析请求包体时候失败。",
	"FailedOperation.MarketingSendTimeConstraint":                     "营销短信发送时间限制，为避免骚扰用户，营销短信只允许在8点到22点发送。",
	"FailedOperation.PhoneNumberInBlacklist":                          "手机号在黑名单库中，通常是用户退订或者命中运营商黑名单导致的，可联系 sms helper 解决。",
	"FailedOperation.SignatureIncorrectOrUnapproved":                  "签名格式错误或者签名未审批，签名只能由中英文、数字组成，要求2 - 12个字,如果符合签名格式规范，请核查签名是否已审批。",
	"FailedOperation.TemplateIncorrectOrUnapproved":                   "模版未审批或请求的内容与审核通过的模版内容不匹配，请参考 1014错误详解。",
	"InternalError.OtherError":                                        "其他错误，请联系 sms helper 并提供失败手机号。",
	"InternalError.RequestTimeException":                              "请求发起时间不正常，通常是由于您的服务器时间与腾讯云服务器时间差异超过10分钟导致的，请核对服务器时间及 API 接口中的时间字段是否正常。",
	"InternalError.RestApiInterfaceNotExist":                          "不存在该 RESTAPI 接口，请核查 REST API 接口说明。",
	"InternalError.SendAndRecvFail":                                   "接口超时或后短信收发包超时，请检查您的网络是否有波动，或联系 sms helper 解决。",
	"InternalError.SigFieldMissing":                                   "后端包体中请求包体没有 Sig 字段或 Sig 为空。",
	"InternalError.SigVerificationFail":                               "后端校验 Sig 失败。",
	"InternalError.Timeout":                                           "请求下发短信超时，请参考 60008错误详解。",
	"InternalError.UnknownError":                                      "未知错误类型。",
	"InvalidParameterValue.ContentLengthLimit":                        "请求的短信内容太长，短信长度规则请参考 国内短信内容长度计算规则。",
	"InvalidParameterValue.IncorrectPhoneNumber":                      "手机号格式错误，请参考 1016错误详解。",
	"InvalidParameterValue.ProhibitedUseUrlInTemplateParameter":       "禁止在模板变量中使用 URL。",
	"InvalidParameterValue.SdkAppidNotExist":                          "SdkAppid 不存在。",
	"InvalidParameterValue.TemplateParameterFormatError":              "验证码模板参数格式错误，验证码类模版，模版变量只能传入0 - 6位（包括6位）纯数字。",
	"InvalidParameterValue.TemplateParameterLengthLimit":              "单个模板变量字符数超过12个，企业认证用户不限制单个变量值字数，您可以 变更实名认证模式，变更为企业认证用户后，该限制变更约1小时左右生效。",
	"LimitExceeded.AppDailyLimit":                                     "业务短信日下发条数超过设定的上限 ，可自行到控制台调整短信频率限制策略。",
	"LimitExceeded.DailyLimit":                                        "短信日下发条数超过设定的上限 (国际/港澳台)，如需调整限制，可联系 sms helper。",
	"LimitExceeded.DeliveryFrequencyLimit":                            "下发短信命中了频率限制策略，可自行到控制台调整短信频率限制策略，如有其他需求请联系 sms helper。",
	"LimitExceeded.PhoneNumberCountLimit":                             "调用短信发送 API 接口单次提交的手机号个数超过200个，请遵守 API 接口说明。",
	"LimitExceeded.PhoneNumberDailyLimit":                             "单个手机号日下发短信条数超过设定的上限，可自行到控制台调整短信频率限制策略。",
	"LimitExceeded.PhoneNumberOneHourLimit":                           "单个手机号1小时内下发短信条数超过设定的上限，可自行到控制台调整短信频率限制策略。",
	"LimitExceeded.PhoneNumberSameContentDailyLimit":                  "单个手机号下发相同内容超过设定的上限，可自行到控制台调整短信频率限制策略。",
	"LimitExceeded.PhoneNumberThirtySecondLimit":                      "单个手机号30秒内下发短信条数超过设定的上限，可自行到控制台调整短信频率限制策略。",
	"MissingParameter.EmptyPhoneNumberSet":                            "传入的号码列表为空，请确认您的参数中是否传入号码。",
	"UnauthorizedOperation.IndividualUserMarketingSmsPermissionDeny":  "个人用户没有发营销短信的权限，请参考 权益区别。",
	"UnauthorizedOperation.RequestIpNotInWhitelist":                   "请求 IP 不在白名单中，您配置了校验请求来源 IP，但是检测到当前请求 IP 不在配置列表中，如有需要请联系 sms helper。",
	"UnauthorizedOperation.RequestPermissionDeny":                     "请求没有权限，请联系 sms helper。",
	"UnauthorizedOperation.SdkAppidIsDisabled":                        "此 sdkappid 禁止提供服务，如有需要请联系 sms helper。",
	"UnauthorizedOperation.SerivceSuspendDueToArrears":                "欠费被停止服务，可自行登录腾讯云充值来缴清欠款。",
	"UnauthorizedOperation.SmsSdkAppidVerifyFail":                     "SmsSdkAppid 校验失败。",
	"UnsupportedOperation.":                                           "不支持该请求。",
	"UnsupportedOperation.ContainDomesticAndInternationalPhoneNumber": "群发请求里既有国内手机号也有国际手机号。",
	"UnsupportedOperation.UnsuportedRegion":                           "不支持该地区短信下发。",
}

func RegisterSmsSend(mobile string, code string) (bool, error) {
	return SendSms(libs.Config.SMS.SecretId, libs.Config.SMS.SecretKey, libs.Config.SMS.SdkAppid, libs.Config.SMS.Sign, "743900", []string{"+86" + mobile}, []string{code}, []string{})
}

func CheckRegisterSmsCode(mobile string, code string) (bool, error) {
	tempSms := models.NewSms()
	tempSms.Mobile = "+86" + mobile
	tempSms.TemplateID = "743900"
	s := &models.Search{
		Fields: []*models.Filed{
			{
				Key:       "mobile",
				Condition: "=",
				Value:     tempSms.Mobile,
			},
			{
				Key:       "template_id",
				Condition: "=",
				Value:     tempSms.TemplateID,
			},
		},
		OrderBy: "id",
		Sort:    "desc",
	}
	tempSms, _ = models.GetSms(s)
	if tempSms.ID <= 0 {
		return false, errors2.New("验证码验证失败")
	}

	if tempSms.Status > 1 {
		return false, errors2.New("验证码已失效,请重新获取")
	}
	if tempSms.CreatedAt.Add(3*time.Minute).Unix() <= time.Now().Unix() {
		return false, errors2.New("验证码超时")
	}

	if !strings.Contains(tempSms.TemplateParamSet, code) {
		return false, errors2.New("验证码验证失败")
	}
	tempSms.Status = 2
	models.UpdateSmsById(tempSms.ID, tempSms)

	return true, nil
}

func SendSms(secretId string, secretKey string, sdkAppid string, sign string, templateID string, mobileArr []string, templateParamSet []string, sessionContext []string) (bool, error) {
	smsDBs := make([]models.Sms, 0)
	for _, value := range mobileArr {
		smsDb := models.Sms{
			SecretId:         secretId,
			Sign:             sign,
			SdkAppid:         sdkAppid,
			Mobile:           value,
			TemplateID:       templateID,
			TemplateParamSet: libs.StructToString(templateParamSet),
			SessionContext:   libs.StructToString(sessionContext),
			Status:           1,
		}
		smsDBs = append(smsDBs, smsDb)
	}

	credential := common.NewCredential(secretId, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "", cpf)

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(mobileArr)
	request.TemplateParamSet = common.StringPtrs(templateParamSet)
	request.TemplateID = common.StringPtr(templateID)
	request.SmsSdkAppid = common.StringPtr(sdkAppid)
	request.Sign = common.StringPtr(sign)
	request.SessionContext = common.StringPtr(libs.StructToString(sessionContext))

	response, err := client.SendSms(request)

	if err != nil {
		for _, tempSmsDb := range smsDBs {
			tempSmsDb.Error = err.Error()
			tempSmsDb.CreateSms()
		}
		return false, err
	}
	for _, value := range response.Response.SendStatusSet {
		if *value.Code != "Ok" {
			var tempErrorinfo string
			if _, ok := TencentSmsErrorInfo[*value.Code]; ok {
				tempErrorinfo = TencentSmsErrorInfo[*value.Code]
			}
			if len(tempErrorinfo) > 0 {
				for _, tempSmsDb := range smsDBs {
					tempSmsDb.Error = tempErrorinfo
					tempSmsDb.CreateSms()
				}
				return false, errors2.New(tempErrorinfo)
			} else {
				for _, tempSmsDb := range smsDBs {
					tempSmsDb.Error = *value.Code
					tempSmsDb.CreateSms()
				}
				return false, errors.NewTencentCloudSDKError(*value.Code, *value.Message, "")
			}

		}

	}
	for _, tempSmsDb := range smsDBs {
		tempSmsDb.Response = response.ToJsonString()
		tempSmsDb.CreateSms()
	}
	return true, nil
}
