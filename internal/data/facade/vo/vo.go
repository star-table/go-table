package vo

import "time"

// 实际工时详情列表单个对象
type ActualListItem struct {
	// 工时执行人名字
	Name string `json:"name"`
	// 一个日期内的工时信息
	ActualWorkHourDateList []*ActualWorkHourDateItem `json:"actualWorkHourDateList"`
}

type ActualWorkHourDateItem struct {
	// 实际工时的日期，开始日期。
	Date string `json:"date"`
	// 工时，单位：小时。
	WorkHour string `json:"workHour"`
}

type AddIssueAttachmentFsData struct {
	// 标题
	Title string `json:"title"`
	// 链接
	URL string `json:"url"`
}

type AddIssueAttachmentFsReq struct {
	IssueID int64                       `json:"issueId"`
	Data    []*AddIssueAttachmentFsData `json:"data"`
}

type AddIssueAttachmentReq struct {
	IssueID     int64   `json:"issueId"`
	ResourceIds []int64 `json:"resourceIds"`
}

// 将一个/多个成员分配到一个/多个部门中接口入参
type AllocateDepartmentReq struct {
	// 用户id
	UserIds []int64 `json:"userIds"`
	// 部门id
	DepartmentIds []int64 `json:"departmentIds"`
}

// 接入应用信息结构体
type AppInfo struct {
	// 主键
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 应用编号
	Code string `json:"code"`
	// 秘钥1
	Secret1 string `json:"secret1"`
	// 秘钥2
	Secret2 string `json:"secret2"`
	// 负责人
	Owner string `json:"owner"`
	// 审核状态,1待审核,2审核通过,3审核未通过
	CheckStatus int `json:"checkStatus"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 接入应用信息列表响应结构体
type AppInfoList struct {
	Total int64      `json:"total"`
	List  []*AppInfo `json:"list"`
}

// 申请授权请求返回
type ApplyScopesResp struct {
	// 申请时的三方返回 code
	ThirdCode int64 `json:"thirdCode"`
	// 申请时的三方返回 msg
	ThirdMsg string `json:"thirdMsg"`
}

// 批量归档任务响应结构体
type ArchiveIssueBatchResp struct {
	// 成功的id
	SuccessIssues []*Issue `json:"successIssues"`
	// 没有权限的任务id
	NoAuthIssues []*Issue `json:"noAuthIssues"`
}

type AreaLinkageListReq struct {
	// 是否是根
	IsRoot *bool `json:"isRoot"`
	// 大陆板块
	ContinentID *int64 `json:"continentId"`
	// 国家Id
	CountryID *int64 `json:"countryId"`
	// 地区Id
	AreaID *int64 `json:"areaId"`
	// 省/州Id
	StateID *int64 `json:"stateId"`
	// 城市Id
	CityID *int64 `json:"cityId"`
}

type AreaLinkageListResp struct {
	List []*AreaLinkageResp `json:"list"`
}

type AreaLinkageResp struct {
	// 主键
	ID int64 `json:"id"`
	// 名字
	Name string `json:"name"`
	// 中文名
	Cname string `json:"cname"`
	// code
	Code string `json:"code"`
	// 是否默认选择
	IsDefault int `json:"isDefault"`
}

type Attachment struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// host
	Host string `json:"host"`
	// 路径
	Path string `json:"path"`
	// 缩略图路径
	PathCompressed string `json:"pathCompressed"`
	// 文件名
	Name string `json:"name"`
	// 存储类型,1：本地，2：oss,3.钉盘
	Type int `json:"type"`
	// 文件大小
	Size int64 `json:"size"`
	// 创建人姓名
	CreatorName string `json:"creatorName"`
	// 文件后缀
	Suffix string `json:"suffix"`
	// 文件的md5
	Md5 string `json:"md5"`
	// 文件类型
	FileType int `json:"fileType"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
	// 关联任务列表
	IssueList []*Issue `json:"issueList"`
}

type AttachmentList struct {
	Total int64         `json:"total"`
	List  []*Attachment `json:"list"`
}

type AttachmentSimpleInfo struct {
	// url
	URL string `json:"url"`
	// 附件名
	Name string `json:"name"`
	// 后缀
	Suffix string `json:"suffix"`
	// 文件大小
	Size int64 `json:"size"`
}

type AuditIssueReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 审核结果(3确认4驳回)
	Status int `json:"status"`
	// 评论
	Comment *string `json:"comment"`
	// 附件
	Attachments []*AttachmentSimpleInfo `json:"attachments"`
}

// DingTalk第三方扫码登录
type AuthDingCodeReq struct {
	Code string `json:"code"`
}

// 选择组织返回登录信息
type AuthForChosenOrgReq struct {
	OutUserID string `json:"outUserId"`
	OrgID     int64  `json:"orgId"`
}

// DingTalk免登陆 Code 登录验证请求结构体
type AuthReq struct {
	// 免登code
	Code string `json:"code"`
	// 企业id
	CorpID string `json:"corpId"`
}

// DingTalk免登陆 Code 登录验证响应结构体
type AuthResp struct {
	// 持久化登录信息的Token
	Token string `json:"token"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 组织名称
	OrgName string `json:"orgName"`
	// 组织code
	OrgCode string `json:"orgCode"`
	// 用户id
	UserID int64 `json:"userId"`
	// 用户姓名
	Name string `json:"name"`
}

type BasicConfigResp struct {
	// 运行模式，1：Saas集群，2：单机部署，3：私有化集群部署，4：私有化单库部署
	RunMode int `json:"runMode"`
	// 构建信息
	BuildInfo *BuildInfoDefine `json:"buildInfo"`
}

// 展示基础设置
type BasicShowSetting struct {
	// 工作台
	WorkBenchShow bool `json:"workBenchShow"`
	// 侧边栏
	SideBarShow bool `json:"sideBarShow"`
	// 镜像统计
	MirrorStat bool `json:"mirrorStat"`
}

type BeforeAfterIssueListReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
}

type BeforeAfterIssueListResp struct {
	// 等待数
	WaitNum int64 `json:"waitNum"`
	// 阻塞数
	BlockNum int64 `json:"blockNum"`
	// 前置任务列表
	BeforeList []*IssueRestInfo `json:"beforeList"`
	// 后置任务列表
	AfterList []*IssueRestInfo `json:"afterList"`
}

// 绑定手机号或者邮箱请求结构体
type BindLoginNameReq struct {
	// 登录地址，手机号或者邮箱
	Address string `json:"address"`
	// 地址类型: 1：手机号，2：邮箱
	AddressType int `json:"addressType"`
	// 验证码
	AuthCode string `json:"authCode"`
}

type BoolResp struct {
	// 是否符合期望、确定、ok：true 表示成功、是、确定；false 表示否定、异常
	IsTrue bool `json:"isTrue"`
}

// 绑定飞书账号
type BoundFeiShuAccountReq struct {
	// codeToken
	CodeToken string `json:"codeToken"`
}

// 绑定飞书请求结构体
type BoundFeiShuReq struct {
	// orgId
	OrgID int64 `json:"orgId"`
	// codeToken
	CodeToken string `json:"codeToken"`
}

type BuildInfoDefine struct {
	GitCommitLog   string `json:"gitCommitLog"`
	GitStatus      string `json:"gitStatus"`
	BuildTime      string `json:"buildTime"`
	BuildGoVersion string `json:"buildGoVersion"`
}

// 批量归档任务响应结构体
type CancelArchiveIssueBatchResp struct {
	// 成功的id
	SuccessIssues []*Issue `json:"successIssues"`
	// 没有权限的任务id
	NoAuthIssues []*Issue `json:"noAuthIssues"`
	// 父任务已被删除导致无法删除的子任务
	ParentDeletedIssues []*Issue `json:"parentDeletedIssues"`
	// 父任务处于归档中导致无法删除的子任务
	ParentFilingIssues []*Issue `json:"parentFilingIssues"`
}

type CancelArchiveIssueReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 任务id
	IssueIds []int64 `json:"issueIds"`
}

type ChangeList struct {
	// 字段
	Field *string `json:"field"`
	// 字段名
	FieldName *string `json:"fieldName"`
	// 旧值
	OldValue *string `json:"oldValue"`
	// 新值
	NewValue *string `json:"newValue"`
}

type ChangeParentIssueReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 任务所属项目的 id，从该项目中移动任务。
	FromProjectID *int64 `json:"fromProjectId"`
	// 变更的父任务id
	ParentID int64 `json:"ParentId"`
}

type ChangeProjectCustomFieldStatusReq struct {
	// 自定义字段id
	FieldID int64 `json:"fieldId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 任务类型
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 状态(1启用2禁用)
	Status int `json:"status"`
}

type ChatData struct {
	// 群聊id
	OutChatID string `json:"outChatId"`
	// 群聊名称
	Name string `json:"name"`
	// 群聊描述
	Description *string `json:"description"`
	// 关联id(关联列表时有)
	RelationID *int64 `json:"relationId"`
	// 群聊头像
	Avatar string `json:"avatar"`
	// 是否是项目主群
	IsMain bool `json:"isMain"`
}

type ChatListResp struct {
	// 数量
	Total int64 `json:"total"`
	// 列表
	List []*ChatData `json:"list"`
}

// 检查项目的工时是否开启接口参数
type CheckIsEnableWorkHourReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
}

// 检查项目的工时是否开启接口返回
type CheckIsEnableWorkHourResp struct {
	// 该项目是否开启工时功能
	IsEnable bool `json:"isEnable"`
}

// 查询员工是否是任务成员请求参数
type CheckIsIssueMemberReq struct {
	// 查询的任务id
	IssueID int64 `json:"issueId"`
	// 查询该用户是否是任务的成员。成员包括：参与人、负责人
	UserID int64 `json:"userId"`
}

// 检测邮箱、手机号、账户是否存在请求结构体
type CheckLoginNameReq struct {
	// 登录地址，手机号或者邮箱
	Address string `json:"address"`
	// 地址类型: 1：手机号，2：邮箱
	AddressType int `json:"addressType"`
}

// 检查是否有特定的权限，请求参数
type CheckSpecificScopeReq struct {
	// 权限标识。后端提供的枚举值。
	PowerFlag string `json:"powerFlag"`
}

// 检查是否有特定的权限，响应参数
type CheckSpecificScopeResp struct {
	HasPower bool `json:"hasPower"`
}

type CondOrder struct {
	// 是否是正序
	Asc bool `json:"asc"`
	// 字段
	Column int64 `json:"column"`
}

type CondsData struct {
	// 类型(between,equal,gt,gte,in,like,lt,lte,not_in,not_like,not_null,is_null,all_in,values_in)
	Type string `json:"type"`
	// 字段类型
	FieldType *string `json:"fieldType"`
	// 值
	Value interface{} `json:"value"`
	// 字段id
	Column int64 `json:"column"`
	// 左值
	Left interface{} `json:"left"`
	// 右值
	Right interface{} `json:"right"`
}

type ConvertCodeReq struct {
	// 项目名
	Name string `json:"name"`
}

type ConvertCodeResp struct {
	// 项目code
	Code string `json:"code"`
}

// 转化为父任务请求结构体
type ConvertIssueToParentReq struct {
	// 要更新的任务id
	ID int64 `json:"id"`
	// 任务所属项目的 id，从该项目中移动任务。
	FromProjectID *int64 `json:"fromProjectId"`
	// 要更新的projectObjectType
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 状态id
	StatusID *int64 `json:"statusId"`
	// 迭代id
	IterationID *int64 `json:"iterationId"`
}

type CopyIssueBatchReq struct {
	// 原有项目id
	OldProjectID int64 `json:"oldProjectId"`
	// 任务id
	OldIssueIds []int64 `json:"oldIssueIds"`
	// 目标项目id
	ProjectID int64 `json:"projectId"`
	// 迭代id(没有则填0)
	IterationID int64 `json:"iterationId"`
	// 任务对象类型(没有则填0)
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 状态id(没有则填0)
	StatusID int64 `json:"statusId"`
	// 复制内容
	ChooseField []string `json:"chooseField"`
}

type CopyIssueReq struct {
	// 任务id
	OldIssueID int64 `json:"oldIssueId"`
	// 任务标题
	Title string `json:"title"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 迭代id(没有则填0)
	IterationID int64 `json:"iterationId"`
	// 任务对象类型(没有则填0)
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 状态id(没有则填0)
	StatusID int64 `json:"statusId"`
	// 复制内容
	ChooseField []string `json:"chooseField"`
	// 需要复制的子任务id
	ChildrenIds []int64 `json:"childrenIds"`
}

// 创建接入应用信息请求结构体
type CreateAppInfoReq struct {
	// 名称
	Name string `json:"name"`
	// 应用编号
	Code string `json:"code"`
	// 秘钥1
	Secret1 string `json:"secret1"`
	// 秘钥2
	Secret2 string `json:"secret2"`
	// 负责人
	Owner string `json:"owner"`
	// 审核状态,1待审核,2审核通过,3审核未通过
	CheckStatus int `json:"checkStatus"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

type CreateCustomFieldReq struct {
	// 名称
	Name string `json:"name"`
	// 类型(1文本类型2单选框3多选框4日期选框5人员选择6是非选择7数字框)
	FieldType int `json:"fieldType"`
	// 选项值
	FieldValue interface{} `json:"fieldValue"`
	// 是否加入组织字段库(1是2否，不选默认为否)
	IsOrgField *int `json:"isOrgField"`
	// 字段描述
	Remark *string `json:"remark"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 任务类型
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
}

// 通过名称创建新部门入参
type CreateDepartmentForInviteReq struct {
	// 父部门id（根目录为0）
	ParentID *int64 `json:"parentID"`
	// 部门名称
	Name *string `json:"name"`
	// 部门主管(选填)
	LeaderIds []int64 `json:"leaderIds"`
}

// 创建部门请求结构体
type CreateDepartmentReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 部门名称
	Name string `json:"name"`
	// 部门标识
	Code string `json:"code"`
	// 父部门id
	ParentID int64 `json:"parentId"`
	// 排序
	Sort int `json:"sort"`
	// 是否隐藏部门,1隐藏,2不隐藏
	IsHide int `json:"isHide"`
	// 来源渠道,
	SourceChannel string `json:"sourceChannel"`
	// 状态, 1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 添加任务评论请求结构体
type CreateIssueCommentReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 评论信息
	Comment string `json:"comment"`
	// 提及的用户id
	MentionedUserIds []int64 `json:"mentionedUserIds"`
}

// 创建问题对象类型请求结构体
type CreateIssueObjectTypeReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 类型名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 创建问题性质请求结构体
type CreateIssuePropertyReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 类型名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 创建任务请求结构体
type CreateIssueReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 标题
	Title string `json:"title"`
	// 优先级
	PriorityID int64 `json:"priorityId"`
	// 类型id，问题，需求....
	TypeID *int64 `json:"typeId"`
	// 负责人
	OwnerID int64 `json:"ownerId"`
	// 参与人
	ParticipantIds []int64 `json:"participantIds"`
	// 关注人
	FollowerIds []int64 `json:"followerIds"`
	// 关注人部门（后端实际转化为人）
	FollowerDeptIds []int64 `json:"followerDeptIds"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 计划工作时长
	PlanWorkHour *int `json:"planWorkHour"`
	// 所属版本id
	VersionID *int64 `json:"versionId"`
	// 所属模块id
	ModuleID *int64 `json:"moduleId"`
	// 父任务id
	ParentID *int64 `json:"parentId"`
	// 备注
	Remark *string `json:"remark"`
	// 备注详情
	RemarkDetail *string `json:"remarkDetail"`
	// 备注提及人
	MentionedUserIds []int64 `json:"mentionedUserIds"`
	// 所属迭代id
	IterationID *int64 `json:"iterationId"`
	// 问题对象类型id
	IssueObjectID *int64 `json:"issueObjectId"`
	// 来源id
	IssueSourceID *int64 `json:"issueSourceId"`
	// 性质id
	IssuePropertyID *int64 `json:"issuePropertyId"`
	// 状态id
	StatusID *int64 `json:"statusId"`
	// 子任务列表
	Children []*IssueChildren `json:"children"`
	// 关联的标签列表
	Tags []*IssueTagReqInfo `json:"tags"`
	// 关联的附件id列表
	ResourceIds []int64 `json:"resourceIds"`
	// 自定义字段
	CustomField []*UpdateIssueCustionFieldData `json:"customField"`
	// 审批人
	AuditorIds []int64 `json:"auditorIds"`
	// 无码入参
	LessCreateIssueReq map[string]interface{} `json:"lessCreateIssueReq"`
	// 前面的任务id
	BeforeID *int64 `json:"beforeId"`
	// 后面的任务id
	AfterID *int64 `json:"afterId"`
	// 前面的无码数据id
	BeforeDataID *string `json:"beforeDataId"`
	// 后面的无码数据id
	AfterDataID *string `json:"afterDataId"`
	// 排序
	Asc *bool `json:"asc"`
	// 是否可忽略字段准确性（目前主要用于模板生成任务）
	IsImport *bool `json:"isImport"`
}

// 任务添加文件资源
type CreateIssueResourceReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 资源路径
	ResourcePath string `json:"resourcePath"`
	// 资源大小，单位B
	ResourceSize int64 `json:"resourceSize"`
	// 文件名
	FileName string `json:"fileName"`
	// 文件后缀
	FileSuffix string `json:"fileSuffix"`
	// md5
	Md5 *string `json:"md5"`
	// bucketName
	BucketName *string `json:"bucketName"`
}

// 创建问题来源请求结构体
type CreateIssueSourceReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

type CreateIssueViewReq struct {
	// 项目 id
	ProjectID *int64 `json:"projectId"`
	// 视图配置
	Config string `json:"config"`
	// 视图备注
	Remark *string `json:"remark"`
	// 是否私有
	IsPrivate *bool `json:"isPrivate"`
	// 所属任务类型 id：需求、任务、缺陷的 id 值
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
	// 视图名称
	ViewName string `json:"viewName"`
	// 类型，1：表格视图，2：看板视图，3：照片视图
	Type *int `json:"type"`
	// 视图排序
	Sort *int64 `json:"sort"`
}

// 创建工时记录接口请求体
type CreateIssueWorkHoursReq struct {
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 关联的任务id
	IssueID int64 `json:"issueId"`
	// 枚举记录类型：1预估工时记录（总预估工时），2实际工时记录，3详细预估工时（子预估工时）
	Type int64 `json:"type"`
	// 工作者id
	WorkerID int64 `json:"workerId"`
	// 所需工时时间，单位：小时
	NeedTime string `json:"needTime"`
	// 开始时间，时间戳
	StartTime int64 `json:"startTime"`
	// 工时记录的结束时间，时间戳
	EndTime *int64 `json:"endTime"`
	// 工时记录的内容，工作内容
	Desc *string `json:"desc"`
}

// 创建迭代请求结构体
type CreateIterationReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 名称
	Name string `json:"name"`
	// 负责人
	Owner int64 `json:"owner"`
	// 计划开始时间
	PlanStartTime time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime time.Time `json:"planEndTime"`
}

// 创建迭代统计请求结构体
type CreateIterationStatReq struct {
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 迭代id
	IterationID int64 `json:"iterationId"`
	// 问题总数
	IssueCount int `json:"issueCount"`
	// 未开始问题数
	IssueWaitCount int `json:"issueWaitCount"`
	// 进行中问题数
	IssueRunningCount int `json:"issueRunningCount"`
	// 已完成问题数
	IssueEndCount int `json:"issueEndCount"`
	// 需求总数
	DemandCount int `json:"demandCount"`
	// 未开始需求数
	DemandWaitCount int `json:"demandWaitCount"`
	// 进行中需求数
	DemandRunningCount int `json:"demandRunningCount"`
	// 已完成需求数
	DemandEndCount int `json:"demandEndCount"`
	// 故事点总数
	StoryPointCount int `json:"storyPointCount"`
	// 未开始故事点数
	StoryPointWaitCount int `json:"storyPointWaitCount"`
	// 进行中故事点数
	StoryPointRunningCount int `json:"storyPointRunningCount"`
	// 已完成故事点数
	StoryPointEndCount int `json:"storyPointEndCount"`
	// 任务总数
	TaskCount int `json:"taskCount"`
	// 未开始任务数
	TaskWaitCount int `json:"taskWaitCount"`
	// 进行中任务数
	TaskRunningCount int `json:"taskRunningCount"`
	// 已完成任务数
	TaskEndCount int `json:"taskEndCount"`
	// 缺陷总数
	BugCount int `json:"bugCount"`
	// 未开始缺陷数
	BugWaitCount int `json:"bugWaitCount"`
	// 进行中缺陷数
	BugRunningCount int `json:"bugRunningCount"`
	// 已完成缺陷数
	BugEndCount int `json:"bugEndCount"`
	// 测试任务总数
	TesttaskCount int `json:"testtaskCount"`
	// 未开始测试任务数
	TesttaskWaitCount int `json:"testtaskWaitCount"`
	// 进行中测试任务数
	TesttaskRunningCount int `json:"testtaskRunningCount"`
	// 已完成测试任务数
	TesttaskEndCount int `json:"testtaskEndCount"`
	// 扩展
	Ext string `json:"ext"`
	// 统计日期
	StatDate time.Time `json:"statDate"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 新增多个预估工时
type CreateMultiIssueWorkHoursReq struct {
	// 关联的任务id
	IssueID int64 `json:"issueId"`
	// 总预估工时
	TotalIssueWorkHourRecord *NewPredicateWorkHour `json:"totalIssueWorkHourRecord"`
	// 子预估工时列表
	PredictWorkHourList []*NewPredicateWorkHour `json:"predictWorkHourList"`
}

// 创建组织请求结构体
type CreateOrgReq struct {
	// 组织名称
	OrgName string `json:"orgName"`
	// 补全个人姓名
	CreatorName *string `json:"creatorName"`
	// 是否要导入示例数据, 1：导入，2：不导入，默认不导入
	ImportSampleData *int `json:"importSampleData"`
	// 来源平台
	SourcePlatform *string `json:"sourcePlatform"`
	// 来源渠道
	SourceChannel *string `json:"sourceChannel"`
	// 所属行业
	IndustryID *int64 `json:"industryId"`
	// 组织规模
	Scale *string `json:"scale"`
	// codeToken如果是绑定飞书团队就传入
	CodeToken *string `json:"codeToken"`
}

// 创建请求结构体
type CreatePermissionOperationReq struct {
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 权限项id
	PermissionID int64 `json:"permissionId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 操作编号,多个半角逗号分隔
	OperationCodes string `json:"operationCodes"`
	// 描述
	Remark string `json:"remark"`
	// 是否显示,1是,2否
	IsShow int `json:"isShow"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 创建请求结构体
type CreatePermissionReq struct {
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 编号,仅支持大写英文字母
	Code string `json:"code"`
	// 名称
	Name string `json:"name"`
	// 父id
	ParentID int64 `json:"parentId"`
	// 权限项类型,1系统,2组织,3项目
	Type int `json:"type"`
	// 权限路径
	Path string `json:"path"`
	// 是否显示,1是,2否
	IsShow int `json:"isShow"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 创建优先级请求结构体
type CreatePriorityReq struct {
	// 组织id,全局的填0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 类型,1项目优先级,2:需求/任务等优先级
	Type int `json:"type"`
	// 排序
	Sort int `json:"sort"`
	// 背景颜色
	BgStyle string `json:"bgStyle"`
	// 字体颜色
	FontStyle string `json:"fontStyle"`
	// 是否默认,1是,2否
	IsDefault int `json:"isDefault"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 创建流程状态请求结构体
type CreateProcessStatusReq struct {
	// 名称
	Name string `json:"name"`
	// 背景颜色
	BgStyle *string `json:"bgStyle"`
	// 字体颜色
	FontStyle *string `json:"fontStyle"`
	// 状态类型,1未开始,2进行中,3已完成
	Type int `json:"type"`
	// 状态类别,1项目状态,2迭代状态,3任务栏状态,
	Category int `json:"category"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 流程id
	ProcessID int64 `json:"processId"`
	// 前一个状态的id 没有给0
	BeforeID int64 `json:"beforeId"`
}

// 创建项目日统计请求结构体
type CreateProjectDayStatReq struct {
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 问题总数
	IssueCount int `json:"issueCount"`
	// 未开始问题数
	IssueWaitCount int `json:"issueWaitCount"`
	// 进行中问题数
	IssueRunningCount int `json:"issueRunningCount"`
	// 已完成问题数
	IssueEndCount int `json:"issueEndCount"`
	// 需求总数
	DemandCount int `json:"demandCount"`
	// 未开始需求数
	DemandWaitCount int `json:"demandWaitCount"`
	// 进行中需求数
	DemandRunningCount int `json:"demandRunningCount"`
	// 已完成需求数
	DemandEndCount int `json:"demandEndCount"`
	// 故事点总数
	StoryPointCount int `json:"storyPointCount"`
	// 未开始故事点数
	StoryPointWaitCount int `json:"storyPointWaitCount"`
	// 进行中故事点数
	StoryPointRunningCount int `json:"storyPointRunningCount"`
	// 已完成故事点数
	StoryPointEndCount int `json:"storyPointEndCount"`
	// 任务总数
	TaskCount int `json:"taskCount"`
	// 未开始任务数
	TaskWaitCount int `json:"taskWaitCount"`
	// 进行中任务数
	TaskRunningCount int `json:"taskRunningCount"`
	// 已完成任务数
	TaskEndCount int `json:"taskEndCount"`
	// 缺陷总数
	BugCount int `json:"bugCount"`
	// 未开始缺陷数
	BugWaitCount int `json:"bugWaitCount"`
	// 进行中缺陷数
	BugRunningCount int `json:"bugRunningCount"`
	// 已完成缺陷数
	BugEndCount int `json:"bugEndCount"`
	// 测试任务总数
	TesttaskCount int `json:"testtaskCount"`
	// 未开始测试任务数
	TesttaskWaitCount int `json:"testtaskWaitCount"`
	// 进行中测试任务数
	TesttaskRunningCount int `json:"testtaskRunningCount"`
	// 已完成测试任务数
	TesttaskEndCount int `json:"testtaskEndCount"`
	// 扩展
	Ext string `json:"ext"`
	// 统计日期
	StatDate time.Time `json:"statDate"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

type CreateProjectDetailReq struct {
	OrgID             int64  `json:"orgId"`
	ProjectID         int64  `json:"projectId"`
	IsEnableWorkHours *int   `json:"isEnableWorkHours"`
	Notice            string `json:"notice"`
}

// 创建文件夹请求结构体
type CreateProjectFolderReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 文件夹名
	Name string `json:"name"`
	// 父级文件夹id
	ParentID int64 `json:"parentId"`
	// 文件夹类型,0其他,1文档,2图片,3视频,4音频
	FileType int64 `json:"fileType"`
}

// 创建项目对象类型请求结构体
type CreateProjectObjectTypeReq struct {
	// 项目id,用来校验权限
	ProjectID int64 `json:"projectId"`
	// 名称
	Name string `json:"name"`
	// 类型,1迭代，2问题
	ObjectType int `json:"objectType"`
	// 前一个对象类型的id 没有给0
	BeforeID int64 `json:"beforeId"`
}

type CreateProjectReq struct {
	// 编号
	Code *string `json:"code"`
	// 名称
	Name string `json:"name"`
	// 前缀编号
	PreCode *string `json:"preCode"`
	// 负责人id
	Owner int64 `json:"owner"`
	// 负责人id集合
	OwnerIds []int64 `json:"ownerIds"`
	// 项目类型
	ProjectTypeID *int64 `json:"projectTypeId"`
	// 优先级
	PriorityID *int64 `json:"priorityId"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 项目公开性,1公开,2私有
	PublicStatus int `json:"publicStatus"`
	// 资源id
	ResourceID *int64 `json:"resourceId"`
	// 是否归档,1归档,2未归档
	IsFiling *int `json:"isFiling"`
	// 描述
	Remark *string `json:"remark"`
	// 项目状态
	Status *int64 `json:"status"`
	// 创建时间
	CreateTime *time.Time `json:"createTime"`
	// 更新人
	Updator *int64 `json:"updator"`
	// 更新时间
	UpdateTime *time.Time `json:"updateTime"`
	// 资源路径
	ResourcePath string `json:"resourcePath"`
	// 资源类型1本地2oss3钉盘
	ResourceType int `json:"resourceType"`
	// 用户成员id
	MemberIds []int64 `json:"memberIds"`
	// 用户成员部门id
	MemberForDepartmentID []int64 `json:"memberForDepartmentId"`
	// 是否全选（针对于项目成员）
	IsAllMember *bool `json:"isAllMember"`
	// 关注人id
	FollowerIds []int64 `json:"followerIds"`
	// 是否同步到飞书日历(4：负责人，8：关注人，12：关注人+负责人。为了兼容旧版，1包含了关注人和负责人；2表示都不包含。)
	IsSyncOutCalendar *int `json:"isSyncOutCalendar"`
	// 针对哪些群体用户，同步到其飞书日历(4：负责人，8：关注人。往后扩展是基于二进制的位值)
	SyncCalendarStatusList []*int `json:"syncCalendarStatusList"`
	// 是否创建群聊（针对于飞书1是2否默认是）
	IsCreateFsChat *int `json:"isCreateFsChat"`
	// 无码文件夹id
	ParentID *int64 `json:"parentId"`
	// 隐私模式状态。1开启；2不开启；默认2。
	PrivacyStatus *int `json:"privacyStatus"`
}

// 新增项目资源
type CreateProjectResourceReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 文件夹id
	FolderID int64 `json:"folderId"`
	// 资源路径
	ResourcePath string `json:"resourcePath"`
	// 资源大小，单位B
	ResourceSize int64 `json:"resourceSize"`
	// 文件名
	FileName string `json:"fileName"`
	// 文件后缀
	FileSuffix string `json:"fileSuffix"`
	// md5
	Md5 *string `json:"md5"`
	// bucketName
	BucketName *string `json:"bucketName"`
}

// 创建存放各类资源，其他业务表统一关联此表id请求结构体
type CreateResourceReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 路径
	Path string `json:"path"`
	// 文件名
	Name string `json:"name"`
	// 存储类型,1：本地，2：oss,3.钉盘
	Type int `json:"type"`
	// 文件后缀
	Suffix string `json:"suffix"`
	// 文件的md5
	Md5 string `json:"md5"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 创建角色请求结构体
type CreateRoleReq struct {
	// 角色组1组织角色2项目角色
	RoleGroupType int `json:"roleGroupType"`
	// 名称
	Name string `json:"name"`
	// 描述
	Remark *string `json:"remark"`
	// 是否只读 1只读 2可编辑
	IsReadonly *int `json:"isReadonly"`
	// 是否可以变更权限,1可以,2不可以
	IsModifyPermission *int `json:"isModifyPermission"`
	// 是否默认角色,1是,2否
	IsDefault *int `json:"isDefault"`
	// 状态,  1可用,2禁用
	Status *int `json:"status"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
}

// 创建请求结构体
type CreateTagReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 名称
	Name string `json:"name"`
	// 背景颜色
	BgStyle string `json:"bgStyle"`
	// 字体颜色
	FontStyle *string `json:"fontStyle"`
}

// 创建请求结构体
type CreateUserOrganizationReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 用户id
	UserID int64 `json:"userId"`
	// 审核状态,1待审核,2审核通过,3审核不过
	CheckStatus int `json:"checkStatus"`
	// 使用状态,1已使用,2未使用
	UseStatus int `json:"useStatus"`
	// 企业用户状态, 1可用,2禁用
	Status int `json:"status"`
	// 状态变更人id
	StatusChangerID int64 `json:"statusChangerId"`
	// 状态变更时间
	StatusChangeTime time.Time `json:"statusChangeTime"`
	// 审核人id
	AuditorID int64 `json:"auditorId"`
	// 审核时间
	AuditTime time.Time `json:"auditTime"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 新增成员
type CreateUserReq struct {
	// 手机号(必填)
	PhoneNumber string `json:"phoneNumber"`
	// 邮箱(必填)
	Email string `json:"email"`
	// 姓名(必填)
	Name string `json:"name"`
	// 部门id
	DepartmentIds []int64 `json:"departmentIds"`
	// 角色id
	RoleIds []int64 `json:"roleIds"`
	// 状态（1启用2禁用, 默认启用）
	Status int `json:"status"`
}

type CustomField struct {
	// id
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 类型
	FieldType int `json:"fieldType"`
	// 值
	FieldValue []map[string]interface{} `json:"fieldValue"`
	// 是否是组织字段(1组织2项目3系统)
	IsOrgField int `json:"isOrgField"`
	// 启用状态（1启用2禁用）对于项目而言
	Status int `json:"status"`
	// 字段描述
	Remark string `json:"remark"`
	// 使用项目
	ProjectList []*SimpleProjectInfo `json:"projectList"`
	// 编辑人信息
	UpdatorInfo *UserIDInfo `json:"updatorInfo"`
	// 编辑时间
	UpdateTime time.Time `json:"updateTime"`
	// 是否展示(1展示2否，在项目视图里使用该字段)
	IsDisplay int `json:"isDisplay"`
}

type CustomFieldListReq struct {
	// 页码
	Page *int `json:"page"`
	// 每页条数
	Size *int `json:"size"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 任务类型
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
	// 是否是没有被项目(敏捷包括任务类型)的字段（和项目id联合使用，只有传了项目id才生效，1是2否，默认否）
	IsUsedCurrentProject *int `json:"isUsedCurrentProject"`
	// 是否属于组织字段库(1组织2项目3系统)
	IsOrgField []int `json:"isOrgField"`
	// 名称
	Name *string `json:"name"`
	// 排序类型（1创建时间正序2创建时间倒序3添加到项目时间正序4添加到项目时间倒序。默认创建时间正序）
	OrderType *int `json:"orderType"`
}

type CustomFieldListResp struct {
	Total int64          `json:"total"`
	List  []*CustomField `json:"list"`
}

type CustomValue struct {
	// 字段id
	ID int64 `json:"id"`
	// 字段名称
	Name string `json:"name"`
	// 字段值
	Value interface{} `json:"value"`
	// 类型(1文本类型2单选框3多选框4日期选框5人员选择6是非选择7数字框)
	FieldType int `json:"fieldType"`
	// 选项值
	FieldValue []map[string]interface{} `json:"fieldValue"`
	// 是否属于组织字段库(1组织2项目3系统)
	IsOrgField int `json:"isOrgField"`
	// 字段描述
	Remark string `json:"remark"`
	// 字段
	Title string `json:"title"`
	// 启用状态（1启用2禁用）对于项目而言
	Status int `json:"status"`
}

// 删除角色请求结构体
type DelRoleReq struct {
	// 角色ID
	RoleIds []int64 `json:"roleIds"`
	// 如果删除项目角色，projectId必填，如果删除组织角色，projectId可以不传或者传0
	ProjectID *int64 `json:"projectId"`
}

// 删除接入应用信息请求结构体
type DeleteAppInfoReq struct {
	// 主键
	ID int64 `json:"id"`
}

type DeleteCustomFieldReq struct {
	// 自定义字段id
	FieldID int64 `json:"fieldId"`
	// 项目id(传了表示删除项目关联)
	ProjectID *int64 `json:"projectId"`
	// 任务类型
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
}

// 删除部门请求结构体
type DeleteDepartmentReq struct {
	// 主键
	ID int64 `json:"id"`
}

// 批量删除任务请求结构体
type DeleteIssueBatchReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 任务id
	Ids []int64 `json:"ids"`
	// 镜像应用id
	MenuAppID *string `json:"menuAppId"`
}

// 批量删除任务响应结构体
type DeleteIssueBatchResp struct {
	// 删除成功的id
	SuccessIssues []*Issue `json:"successIssues"`
	// 没有权限的任务id
	NoAuthIssues []*Issue `json:"noAuthIssues"`
	// 还有子任务没有选择的父任务id
	RemainChildrenIssues []*Issue `json:"remainChildrenIssues"`
}

// 删除问题对象类型请求结构体
type DeleteIssueObjectTypeReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织Id 暂时不用传 之后用户校验的时候比较是否包含这个orgId 操作的时候是否有当前orgId的权限
	OrgID *int64 `json:"orgId"`
}

// 删除任务响应结构体
type DeleteIssueReq struct {
	// 任务id
	ID int64 `json:"id"`
	// 是否携带子任务(默认带上，兼容以前的)
	TakeChildren *bool `json:"takeChildren"`
}

// 删除子任务请求结构体
type DeleteIssueResourceReq struct {
	// 任务id'
	IssueID int64 `json:"issueId"`
	// 关联资源id列表
	DeletedResourceIds []int64 `json:"deletedResourceIds"`
}

// 删除问题来源请求结构体
type DeleteIssueSourceReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织Id 暂时不用传 之后用户校验的时候比较是否包含这个orgId 操作的时候是否有当前orgId的权限
	OrgID *int64 `json:"orgId"`
}

type DeleteIssueViewReq struct {
	// 主键id，根据主键删除
	ID int64 `json:"id"`
}

// 删除工时记录接口请求体
type DeleteIssueWorkHoursReq struct {
	// 工时记录id
	IssueWorkHoursID int64 `json:"issueWorkHoursId"`
}

// 删除迭代结构体
type DeleteIterationReq struct {
	// 主键
	ID int64 `json:"id"`
}

// 删除迭代统计请求结构体
type DeleteIterationStatReq struct {
	// 主键
	ID int64 `json:"id"`
}

// 删除请求结构体
type DeleteNoticeReq struct {
	// 主键
	ID int64 `json:"id"`
}

// 删除请求结构体
type DeletePermissionOperationReq struct {
	// 主键
	ID int64 `json:"id"`
}

// 删除请求结构体
type DeletePermissionReq struct {
	// 主键
	ID int64 `json:"id"`
}

// 删除优先级请求结构体
type DeletePriorityReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织Id 暂时不用传 之后用户校验的时候比较是否包含这个orgId 操作的时候是否有当前orgId的权限
	OrgID *int64 `json:"orgId"`
}

// 删除流程状态请求结构体
type DeleteProcessStatusReq struct {
	// 主键
	ID int64 `json:"id"`
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type DeleteProjectAttachmentReq struct {
	// 文件id数组
	ResourceIds []int64 `json:"resourceIds"`
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type DeleteProjectAttachmentResp struct {
	// 文件id数组
	ResourceIds []int64 `json:"resourceIds"`
}

// 删除项目日统计请求结构体
type DeleteProjectDayStatReq struct {
	// 主键
	ID int64 `json:"id"`
}

type DeleteProjectDetailReq struct {
	ID int64 `json:"id"`
}

// 删除文件夹请求结构体
type DeleteProjectFolderReq struct {
	// 文件夹id数组
	FolderIds []int64 `json:"folderIds"`
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type DeleteProjectFolderResp struct {
	// 文件夹id数组
	FolderIds []int64 `json:"folderIds"`
}

// 删除项目对象类型请求结构体
type DeleteProjectObjectTypeReq struct {
	// 主键
	ID int64 `json:"id"`
	// 项目iid,用来校验权限
	ProjectID int64 `json:"projectId"`
}

type DeleteProjectReq struct {
	// 项目id
	ID int64 `json:"id"`
}

type DeleteProjectResourceReq struct {
	// 文件id数组
	ResourceIds []int64 `json:"resourceIds"`
	// 文件夹id,只支持相同目录下的批量文件删除
	FolderID int64 `json:"folderId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type DeleteProjectResourceResp struct {
	// 文件id数组
	ResourceIds []int64 `json:"resourceIds"`
}

// 删除存放各类资源，其他业务表统一关联此表id请求结构体
type DeleteResourceReq struct {
	// 主键
	ID int64 `json:"id"`
}

// 删除请求结构体
type DeleteTagReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 标签id集合
	Ids []int64 `json:"ids"`
}

// 删除请求结构体
type DeleteUserOrganizationReq struct {
	// 主键
	ID int64 `json:"id"`
}

// 部门结构体
type Department struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 部门名称
	Name string `json:"name"`
	// 部门标识
	Code string `json:"code"`
	// 父部门id
	ParentID int64 `json:"parentId"`
	// 排序
	Sort int `json:"sort"`
	// 部门状态
	Status int `json:"status"`
	// 是否隐藏部门,1隐藏,2不隐藏
	IsHide int `json:"isHide"`
	// 来源渠道,
	SourceChannel string `json:"sourceChannel"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
}

// 部门列表响应结构体
type DepartmentList struct {
	// 总数
	Total int64 `json:"total"`
	// 列表
	List []*Department `json:"list"`
}

type DepartmentListReq struct {
	// 父部门id
	ParentID *int64 `json:"parentId"`
	// 是否查询最上级部门, 如果是1则为true
	IsTop *int `json:"isTop"`
	// 是否显示隐藏的部门，如果是1则为true，默认不显示
	ShowHiding *int `json:"showHiding"`
	// 部门名称
	Name *string `json:"name"`
	// 部门id
	DepartmentIds []int64 `json:"departmentIds"`
}

// 部门用户信息
type DepartmentMemberInfo struct {
	// id
	UserID int64 `json:"userId"`
	// 姓名
	Name string `json:"name"`
	// 姓名拼音
	NamePy *string `json:"namePy"`
	// 用户头像
	Avatar string `json:"avatar"`
	// 工号：企业下唯一
	EmplID string `json:"emplId"`
	// unionId： 开发者账号下唯一
	UnionID string `json:"unionId"`
	// 用户部门id
	DepartmentID int64 `json:"departmentId"`
	// 用户组织状态
	OrgUserStatus int `json:"orgUserStatus"`
}

type DepartmentMemberListReq struct {
	// 部门id
	DepartmentID *int64 `json:"departmentId"`
}

type DepartmentMembersListReq struct {
	// 名称
	Name *string `json:"name"`
	// 限定人员范围
	UserIds []int64 `json:"userIds"`
	// 需要排除的项目id,取项目之外的组织成员
	ExcludeProjectID *int64 `json:"excludeProjectId"`
	// 关联类型(1负责人2关注人3全部，默认全部,配合项目id使用)
	RelationType *int64 `json:"relationType"`
}

type DepartmentMembersListResp struct {
	// 总数
	Total int64 `json:"total"`
	// 列表
	List []*DepartmentMemberInfo `json:"list"`
}

type DepartmentSimpleInfo struct {
	// 部门id
	ID int64 `json:"id"`
	// 部门名称
	Name string `json:"name"`
	// 部门总人数
	UserCount int64 `json:"userCount"`
}

// 钉钉扫码登录返回结果
type DingAuthCodeResp struct {
	NickName  string              `json:"nickName"`
	OutUserID string              `json:"outUserId"`
	OrgList   []*OrgInfoForChosen `json:"orgList"`
}

// 向钉钉群发送告警日志接口请求参数
type DingTalkInfoReq struct {
	// 日志信息。没有则传空字符串。
	Content string `json:"content"`
	// 其他信息。没有则传空字符串。
	Other string `json:"other"`
}

// 启用/关闭工时记录功能接口请求体
type DisOrEnableIssueWorkHoursReq struct {
	// 项目id，针对这个项目关闭/启用工时功能
	ProjectID int64 `json:"projectId"`
	// 是否关闭工时功能：1启用,2关闭
	Enable int64 `json:"enable"`
}

type DisbandThirdAccountReq struct {
	SourceChannel string `json:"sourceChannel"`
}

// 清空所在组织下某个状态的成员
type EmptyUserReq struct {
	// 状态,  1可用,2禁用
	Status int `json:"status"`
}

type EveryPermission struct {
	// 权限组id
	PermissionID int64 `json:"permissionId"`
	// 修改后的操作项id
	OperationIds []int64 `json:"operationIds"`
}

// 导出通讯录入参
type ExportAddressListReq struct {
	// 搜索字段
	SearchCode *string `json:"searchCode"`
	// 是否已分配部门（1已分配2未分配，默认全部）
	IsAllocate *int `json:"isAllocate"`
	// 是否禁用（1启用2禁用，默认全部）
	Status *int `json:"status"`
	// 角色id
	RoleID *int64 `json:"roleId"`
	// 部门id
	DepartmentID *int64 `json:"departmentId"`
	// 导出字段(name,mobile,email,department,role,isLeader,statusChangeTime,createTime)
	ExportField []string `json:"exportField"`
}

// 导出通讯录响应结果
type ExportAddressListResp struct {
	// 导出文件的下载地址
	URL string `json:"url"`
}

type ExportIssueTemplateResp struct {
	// 模板地址
	URL string `json:"url"`
}

// 工时统计的导出接口请求返回
type ExportWorkHourStatisticResp struct {
	// 导出文件的下载地址。
	URL string `json:"url"`
}

// 获取飞书免登陆Code认证信息
type FeiShuAuthCodeResp struct {
	// 企业ID
	TenantKey string `json:"tenantKey"`
	// 用户OpenID
	OpenID string `json:"openId"`
	// 是否为企业管理
	IsAdmin bool `json:"isAdmin"`
	// 是否被绑定
	Binding bool `json:"binding"`
	// refreshToken
	RefreshToken string `json:"refreshToken"`
	// accessToken
	AccessToken string `json:"accessToken"`
	// token
	Token string `json:"token"`
	// codeToken
	CodeToken string `json:"codeToken"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 组织名称
	OrgName string `json:"orgName"`
	// 外部组织名称
	OutOrgName string `json:"outOrgName"`
	// 组织code
	OrgCode string `json:"orgCode"`
	// 用户id
	UserID int64 `json:"userId"`
	// 用户姓名
	Name string `json:"name"`
}

// 飞书免登陆Code 登录验证请求结构体
type FeiShuAuthReq struct {
	// 免登code
	Code string `json:"code"`
	// 免登code类型，1: code2session, 2: oauth(默认为1)
	CodeType *int `json:"codeType"`
}

// 飞书免登陆Code 登录验证响应结构体
type FeiShuAuthResp struct {
	// 持久化登录信息的Token
	Token string `json:"token"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 组织名称
	OrgName string `json:"orgName"`
	// 组织code
	OrgCode string `json:"orgCode"`
	// 用户id
	UserID int64 `json:"userId"`
	// 用户姓名
	Name string `json:"name"`
	// 企业ID
	TenantKey string `json:"tenantKey"`
	// 用户OpenID
	OpenID string `json:"openId"`
	// 是否为企业管理
	IsAdmin bool `json:"isAdmin"`
}

// 文件夹结构体
type Folder struct {
	// 文件夹id
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 文件夹名
	Name string `json:"name"`
	// 父级文件夹id
	ParentID int64 `json:"parentId"`
	// 文件夹类型,0其他,1文档,2图片,3视频,4音频
	FileType int64 `json:"fileType"`
	// 文件路径
	Path string `json:"path"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建人中文名
	CreatorName string `json:"creatorName"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 文件夹列表相应结构体
type FolderList struct {
	Total int64     `json:"total"`
	List  []*Folder `json:"list"`
}

type FunctionConfigResp struct {
	// 功能
	FunctionCodes []string `json:"functionCodes"`
}

type GetAppTicketResp struct {
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type GetExportFieldsReq struct {
	// 项目id
	ProjectID *int64 `json:"projectId"`
}

type GetExportFieldsResp struct {
	Fields []*GetExportFieldsRespFieldsItem `json:"fields"`
}

type GetExportFieldsRespFieldsItem struct {
	// 字段id
	FieldID int64 `json:"fieldId"`
	// 字段名
	Name string `json:"name"`
	// 是否必须。true 表示必须。如果必须，则必须导出该字段。
	IsMust bool `json:"isMust"`
	// 定义类型。10原生字段，11用户自定义，12系统字段
	DefineType int `json:"defineType"`
}

type GetFsProjectChatPushSettingsReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type GetFsProjectChatPushSettingsResp struct {
	// 添加任务(1开2关)
	CreateIssue int `json:"createIssue"`
	// 任务负责人变更
	UpdateIssueOwner int `json:"updateIssueOwner"`
	// 任务状态变更
	UpdateIssueStatus int `json:"updateIssueStatus"`
	// 任务栏变更
	UpdateIssueProjectObjectType int `json:"updateIssueProjectObjectType"`
	// 任务标题被修改
	UpdateIssueTitle int `json:"updateIssueTitle"`
	// 任务时间被修改
	UpdateIssueTime int `json:"updateIssueTime"`
	// 任务有新的评论
	CreateIssueComment int `json:"createIssueComment"`
	// 任务有新的附件
	UploadNewAttachment int `json:"uploadNewAttachment"`
}

// 获取邀请码请求结构体
type GetInviteCodeReq struct {
	// 平台
	SourcePlatform *string `json:"sourcePlatform"`
}

// 获取邀请码响应结构体
type GetInviteCodeResp struct {
	// 邀请码
	InviteCode string `json:"inviteCode"`
	// 有效时长，单位：秒
	Expire int `json:"expire"`
}

// 获取邀请信息请求结构体
type GetInviteInfoReq struct {
	// 邀请code
	InviteCode string `json:"inviteCode"`
}

// 获取邀请信息响应结构体
type GetInviteInfoResp struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 组织名
	OrgName string `json:"orgName"`
	// 邀请人id
	InviterID int64 `json:"inviterId"`
	// 邀请人姓名
	InviterName string `json:"inviterName"`
}

// 获取任务资源请求结构体
type GetIssueResourcesReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
}

type GetIssueViewListItem struct {
	// 主键id
	ID int64 `json:"id"`
	// 项目 id
	ProjectID int64 `json:"projectId"`
	// 视图配置
	Config string `json:"config"`
	// 视图备注
	Remark string `json:"remark"`
	// 是否私有
	IsPrivate bool `json:"isPrivate"`
	// 视图名称
	ViewName string `json:"viewName"`
	// 类型，1：表格视图，2：看板视图，3：照片视图
	Type int `json:"type"`
	// 视图排序
	Sort int64 `json:"sort"`
	// 所属任务类型 id：需求、任务、缺陷的 id 值
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
}

type GetIssueViewListReq struct {
	// 筛选：视图 id，支持多个 id
	Ids []int64 `json:"ids"`
	// 筛选：项目 id
	ProjectID *int64 `json:"projectId"`
	// 筛选：视图名称。模糊查询
	ViewName *string `json:"viewName"`
	// 筛选：是否私有，true 私有，false 公开
	IsPrivate *bool `json:"isPrivate"`
	// 筛选：类型，1：表格视图，2：看板视图，3：照片视图
	Type *int `json:"type"`
	// 所属任务类型 id：需求、任务、缺陷的 id 值
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
	// 排序类型。1创建时间顺序，2创建时间倒序，3更新时间顺序，4更新时间倒序。默认1。
	SortType *int `json:"sortType"`
	// 页码
	Page *int `json:"page"`
	// 页大小
	Size *int `json:"size"`
}

// 查询任务的工时信息接口参数
type GetIssueWorkHoursInfoReq struct {
	// 关联的任务id
	IssueID int64 `json:"issueId"`
}

// 查询任务的工时信息接口返回值
type GetIssueWorkHoursInfoResp struct {
	// 总预估工时。issue 一旦有工时记录，则一定会有一个总预估工时记录。
	SimplePredictWorkHour *OneWorkHourRecord `json:"simplePredictWorkHour"`
	// 预估工时列表
	PredictWorkHourList []*OneWorkHourRecord `json:"predictWorkHourList"`
	// 实际工时列表
	ActualWorkHourList []*OneActualWorkHourRecord `json:"actualWorkHourList"`
	// 实际总工时时间。单位：小时。
	ActualNeedTimeTotal string `json:"actualNeedTimeTotal"`
}

// 获取工时记录的列表请求参数
type GetIssueWorkHoursListReq struct {
	// 页码
	Page *int64 `json:"page"`
	// 每页条数
	Size *int64 `json:"size"`
	// 关联的任务id
	IssueID int64 `json:"issueId"`
	// 记录类型：1预估工时记录，2实际工时记录，3子预估工时
	Type int64 `json:"type"`
}

// 获取的工时记录列表结果
type GetIssueWorkHoursListResp struct {
	// 总数量
	Total int64 `json:"total"`
	// 任务列表
	List []*IssueWorkHours `json:"list"`
}

// 获取MQTT通道key请求结构体
type GetMQTTChannelKeyReq struct {
	// 通道类型：1、项目（任务，标签，工作栏），2、组织（成员）
	ChannelType int `json:"channelType"`
	// 通道类型为1时必传
	ProjectID *int64 `json:"projectId"`
	// 通道类型为4时必传
	AppID *int64 `json:"appId"`
}

// 获取MQTT通道key响应结构体
type GetMQTTChannelKeyResp struct {
	// 连接地址
	Address string `json:"address"`
	// host
	Host string `json:"host"`
	// port
	Port *int `json:"port"`
	// 通道
	Channel string `json:"channel"`
	// 通道key
	Key string `json:"key"`
}

type GetPayRemindResp struct {
	// 是否需要提示付费信息(为空则表示不需要，否则展示提示信息)
	RemindPayExpireMsg string `json:"remindPayExpireMsg"`
}

type GetPersonalPermissionInfoResp struct {
	Data map[string]interface{} `json:"Data"`
}

type GetProjectMainChatIDReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type GetProjectMainChatIDResp struct {
	// 关联群聊id
	ChatID string `json:"chatId"`
}

type GetShareURLResp struct {
	URL string `json:"url"`
}

// 获取/查询成员列表的参数结构
type GetUserListReq struct {
	// 成员状态
	SearchCode *string `json:"searchCode"`
	// 是否已分配部门（1已分配2未分配，默认全部）
	IsAllocate *int `json:"isAllocate"`
	// 是否禁用（1启用2禁用，默认全部）
	Status *int `json:"status"`
	// 角色 id
	RoleID *int64 `json:"roleId"`
	// 部门 id
	DepartmentID *int64 `json:"departmentId"`
	// 页码（选填，不填为 1）
	Page int `json:"page"`
	// 每页数量（选填，不填为全部）
	Size int `json:"size"`
}

// 获取成员列表响应结构体
type GetUserListResp struct {
	// 总数量
	Total int64 `json:"total"`
	// 成员列表
	List []*UserInfo `json:"list"`
}

// 工时统计查询请求参数
type GetWorkHourStatisticReq struct {
	// 项目id，查询项目下的工时统计。可选
	ProjectIds []*int64 `json:"projectIds"`
	// 任务 id，查询任务下的工时统计。可选
	IssueIds []*int64 `json:"issueIds"`
	// 可多选。任务状态,1:未完成，2：已完成，3：未开始，4：进行中，5: 已逾期
	IssueStatus []*int `json:"issueStatus"`
	// 优先级id，可多选
	IssuePriorities []*int64 `json:"issuePriorities"`
	// 执行者，工时执行人id。可选
	WorkerIds []*int64 `json:"workerIds"`
	// 查询的开始时间。秒级时间戳。可选
	StartTime *int64 `json:"startTime"`
	// 查询的截止时间。秒级时间戳。可选
	EndTime *int64 `json:"endTime"`
	// 是否显示已离职人员。1显示，2不显示。默认不显示。
	ShowResigned *int `json:"showResigned"`
	// 页码
	Page *int64 `json:"page"`
	// 每页条数
	Size *int64 `json:"size"`
}

// 工时统计查询返回参数
type GetWorkHourStatisticResp struct {
	// 多个成员在一段日期内的工时信息列表
	GroupStatisticList []*OnePersonWorkHourStatisticInfo `json:"groupStatisticList"`
	// 数据总数
	Total int64 `json:"total"`
	// 汇总的数据
	Summary *GetWorkHourStatisticSummary `json:"summary"`
}

type GetWorkHourStatisticSummary struct {
	// 筛选条件下的预估工时的总和
	PredictTotal string `json:"predictTotal"`
	// 筛选条件下的实际工时的总和
	ActualTotal string `json:"actualTotal"`
}

type HandleOldIssueToNewReq struct {
	OrgID     int64   `json:"orgId"`
	ProjectID int64   `json:"projectId"`
	IssueIds  []int64 `json:"issueIds"`
}

type HomeIssueGroup struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	// 图片
	Avatar string `json:"avatar"`
	// 背景色
	BgStyle string `json:"bgStyle"`
	// 字体色
	FontStyle string `json:"fontStyle"`
	// 时间跨度
	TimeSpan int64 `json:"timeSpan"`
	// 满足甘特图的任务数量
	FitTotal int64            `json:"fitTotal"`
	List     []*HomeIssueInfo `json:"list"`
}

// 首页任务信息结构体
type HomeIssueInfo struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 父任务id
	ParentID int64 `json:"parentId"`
	// 父任务信息
	ParentInfo []*ParentInfo `json:"parentInfo"`
	// 任务标题
	Title string `json:"title"`
	// 是否是查询结果附带的父任务1是0否
	IsAttach int `json:"isAttach"`
	// 父任务是否是满足条件1是0否
	ParentIsMeetCondition int `json:"parentIsMeetCondition"`
	// 任务信息
	Issue *Issue `json:"issue"`
	// 项目信息
	Project *HomeIssueProjectInfo `json:"project"`
	// 执行人信息
	Owner *HomeIssueOwnerInfo `json:"owner"`
	// 状态信息
	Status *HomeIssueStatusInfo `json:"status"`
	// 优先级信息
	Priority *HomeIssuePriorityInfo `json:"priority"`
	// 标签信息
	Tags []*HomeIssueTagInfo `json:"tags"`
	// 子任务数量
	ChildsNum int64 `json:"childsNum"`
	// 子任务已完成数量
	ChildsFinishedNum int64 `json:"childsFinishedNum"`
	// 任务栏名称
	ProjectObjectTypeName string `json:"projectObjectTypeName"`
	// 状态列表
	AllStatus []*HomeIssueStatusInfo `json:"allStatus"`
	// 来源信息
	SourceInfo *IssueSourceInfo `json:"sourceInfo"`
	// 严重程度信息
	PropertyInfo *IssuePropertyInfo `json:"propertyInfo"`
	// 类型信息
	TypeInfo *IssueObjectTypeInfo `json:"typeInfo"`
	// 迭代名称
	IterationName string `json:"iterationName"`
	// 关注人
	FollowerInfos []*UserIDInfo `json:"followerInfos"`
	// 关联任务数量
	RelateIssueCount int64 `json:"relateIssueCount"`
	// 关联资源数量
	RelateResourceCount int64 `json:"relateResourceCount"`
	// 关联评论数量
	RelateCommentCount int64 `json:"relateCommentCount"`
	// 自定义字段结果
	CustomField []*CustomValue `json:"customField"`
	// 工时信息
	WorkHourInfo *HomeIssueWorkHourInfo `json:"workHourInfo"`
	// 确认人信息
	AuditorsInfo []*UserIDInfoExtraForIssueAudit `json:"auditorsInfo"`
	// 后置任务id集合
	AfterIssueIds []int64 `json:"afterIssueIds"`
	// 无码数据
	LessData map[string]interface{} `json:"lessData"`
}

// 首页任务列表响应结构体
type HomeIssueInfoGroupResp struct {
	// 总数量
	Total int64 `json:"total"`
	// 实际总数量
	ActualTotal int64 `json:"actualTotal"`
	// 时间跨度
	TimeSpan int64 `json:"timeSpan"`
	// 分组列表
	Group []*HomeIssueGroup `json:"group"`
}

// 首页的任务列表请求结构体
type HomeIssueInfoReq struct {
	// 关联类型，1：我发起的，2：我负责的，3：我参与的，4：我关注的，5：我审批的,6:待我审批的（审批人是我，我还没有审批的）
	RelatedType *int `json:"relatedType"`
	// 排序类型，1：项目分组，2：优先级分组，3：创建日期降序，4：最后更新日期降序, 5: 按开始时间最早, 6：按开始时间最晚, 8：按截止时间最近，9：按创建时间最早, 10: sort排序（正序）11：sort排序（倒序）12:截止时间（正序）13：优先级正序14：优先级倒序15：负责人正序16：负责人倒序17：编号正序18：编号倒序19：标题正序20：标题倒序21：状态正序（必须传项目id，敏捷必须指定任务栏）22：状态倒序（必须传项目id，敏捷必须指定任务栏）23:完成时间倒序24:按照传入id排序25:按照父任务正序26：按照父任务倒序
	OrderType *int `json:"orderType"`
	// 状态,1:未完成，2：已完成，3：未开始，4：进行中，5: 已逾期，-1代表待确认，此状态用于审批
	Status     *int  `json:"status"`
	StatusList []int `json:"statusList"`
	// 任务真实状态集合(传入-1代表待确认，此状态用于审批)
	TrulyStatusIds []int64 `json:"trulyStatusIds"`
	// 是否逾期 （1是2否，不传为全部）
	IsOverdue *int `json:"isOverdue"`
	// 流程状态id
	ProcessStatusID *int64 `json:"processStatusId"`
	// 类型，1：主任务,2 子任务
	Type *int `json:"type"`
	// 截止时间开始时间点(若只选择开始时间：表示任务截止日期在这之后的所有任务)
	StartTime *time.Time `json:"startTime"`
	// 截止时间结束时间点(若只选择截止时间：则表示任务截止时间在这之前的所有任务)
	EndTime *time.Time `json:"endTime"`
	// 负责人
	OwnerIds []int64 `json:"ownerIds"`
	// 创建人
	CreatorIds []int64 `json:"creatorIds"`
	// 参与人
	ParticipantIds []int64 `json:"participantIds"`
	// 关注人
	FollowerIds []int64 `json:"followerIds"`
	// 时间范围：本周，全部..
	TimeScope *time.Time `json:"timeScope"`
	// 搜索筛选
	SearchCond *string `json:"searchCond"`
	// code筛选
	Code *string `json:"code"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 迭代id
	IterationID *int64 `json:"iterationId"`
	// 是否已规划(是否关联了迭代)，1: 已规划，2：未规划
	PlanType *int `json:"planType"`
	// 项目对象类型id
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
	// 支持多个项目对象类型同时做筛选
	ProjectObjectTypeIds []int64 `json:"projectObjectTypeIds"`
	// 优先级
	PriorityID *int64 `json:"priorityId"`
	// 优先级集合
	PriorityIds []int64 `json:"priorityIds"`
	// 组合查询类型，1: 今日指派给我，2：最近截止(包括即将逾期和已逾期)，3：今日逾期，4：逾期完成, 5:即将逾期,6:今日创建,7:今日完成,8:今日我关注的，9：今日分配给我审批的，10:今日分配给我审批，待我审批的（审批人是我，我还没有审批的）
	CombinedType *int `json:"combinedType"`
	// 任务标签id
	IssueTagID []int64 `json:"issueTagId"`
	// 项目是否归档(1是2否3所有,不传查询未归档)
	IsFiling *int `json:"isFiling"`
	// 任务是否归档（1是2否3所有，不传默认未归档）
	IssueIsFiling *int `json:"issueIsFiling"`
	// 父任务id
	ParentID *int64 `json:"parentId"`
	// 上次更新时间（会查询这个时间点之后有变动的任务，其中包括被删除的任务）
	LastUpdateTime *time.Time `json:"lastUpdateTime"`
	// 是否查询父任务的子任务, 如果不是空，且值为1，则将子任务的父任务也查询出来
	EnableParentIssues *int `json:"enableParentIssues"`
	// 周期开始时间（只要开始时间和截止时间有一个大于该项则命中条件）
	PeriodStartTime *time.Time `json:"periodStartTime"`
	// 周期结束时间（只要开始时间和截止时间有一个小于该项则命中条件）
	PeriodEndTime *time.Time `json:"periodEndTime"`
	// 附件资源id
	ResourceID *int64 `json:"resourceId"`
	// 父子堆叠列表(1是2否，不传默认为否)
	IsParentBeforeChid *int `json:"isParentBeforeChid"`
	// 任务类型
	IssueObjectID *int64 `json:"issueObjectId"`
	// 任务类型集合
	IssueObjectIds []int64 `json:"issueObjectIds"`
	// 严重程度
	IssuePropertyID *int64 `json:"issuePropertyId"`
	// 严重程度集合
	IssuePropertyIds []int64 `json:"issuePropertyIds"`
	// 需求来源
	IssueSourceID *int64 `json:"issueSourceId"`
	// 需求来源集合
	IssueSourceIds []int64 `json:"issueSourceIds"`
	// 任务id集合
	IssueIds []int64 `json:"issueIds"`
	// 分组类别(仅用于homeIssuesGroup接口:1负责人2状态3优先级4任务栏5迭代6具体状态，其余默认不分组)
	GroupType *int `json:"groupType"`
	// 添加前置任务时传递任务id，排除掉后置任务中已有的任务
	IssueIDForBefore *int64 `json:"issueIdForBefore"`
	// 添加后置任务列表时传递任务id，排除掉前置任务中已有的任务
	IssueIDForAfter *int64 `json:"issueIdForAfter"`
	// 确认人
	AuditorIds []int64 `json:"auditorIds"`
	// 自定义字段(取并集)
	Conds []*CondsData `json:"conds"`
	// 自定义字段排序
	CondOrder []*CondOrder `json:"condOrder"`
	// 无码格式
	LessConds *LessCondsData `json:"lessConds"`
	// 无码格式排序
	LessOrder []*LessOrder `json:"lessOrder"`
	// 仅通过极星查询数据
	IsOnlyPolaris *bool `json:"isOnlyPolaris"`
	// 分配时间开始
	OwnerChangeTimeStart *time.Time `json:"ownerChangeTimeStart"`
	// 分配时间截至
	OwnerChangeTimeEnd *time.Time `json:"ownerChangeTimeEnd"`
	// 当前任务（用于变更父任务时查询任务列表）
	CurrentIssueID *int64 `json:"currentIssueId"`
	// 镜像应用id
	MenuAppID *string `json:"menuAppId"`
}

// 首页任务列表响应结构体
type HomeIssueInfoResp struct {
	// 总数量
	Total int64 `json:"total"`
	// 实际总数量
	ActualTotal int64 `json:"actualTotal"`
	// 首页任务列表
	List []*HomeIssueInfo `json:"list"`
}

// 首页任务-负责人信息结构体
type HomeIssueOwnerInfo struct {
	// 负责人信息
	ID int64 `json:"id"`
	// 负责人id
	UserID int64 `json:"userId"`
	// 负责人名称
	Name string `json:"name"`
	// 负责人头像
	Avatar *string `json:"avatar"`
	// 是否已被删除，为true则代表被组织移除
	IsDeleted bool `json:"isDeleted"`
	// 是否已被禁用, 为true则代表被组织禁用
	IsDisabled bool `json:"isDisabled"`
}

// 首页任务-优先级信息结构体
type HomeIssuePriorityInfo struct {
	// 优先级id
	ID int64 `json:"id"`
	// 优先级名称
	Name string `json:"name"`
	// 背景色
	BgStyle string `json:"bgStyle"`
	// 字体色
	FontStyle string `json:"fontStyle"`
}

// 首页任务-项目信息结构体
type HomeIssueProjectInfo struct {
	// 项目id
	ID int64 `json:"id"`
	// 项目对应的应用 id（无码系统）
	AppID string `json:"appId"`
	// 项目名称
	Name string `json:"name"`
	// 是否归档(1是2否)
	IsFilling int `json:"isFilling"`
	// 项目类型
	ProjectTypeID int64 `json:"projectTypeId"`
	// 项目隐私状态。1开启隐私；2不开启。
	PrivacyStatus int `json:"privacyStatus"`
}

// 首页任务-状态信息结构体
type HomeIssueStatusInfo struct {
	// 状态id
	ID int64 `json:"id"`
	// 状态名
	Name string `json:"name"`
	// 显示名，为空则显示状态名
	DisplayName *string `json:"displayName"`
	// 背景色
	BgStyle string `json:"bgStyle"`
	// 字体色
	FontStyle string `json:"fontStyle"`
	// 状态类型,1未开始,2进行中,3已完成
	Type int `json:"type"`
	// 排序
	Sort int `json:"sort"`
}

// 首页任务tag信息
type HomeIssueTagInfo struct {
	// 标签id
	ID int64 `json:"id"`
	// 标签名
	Name string `json:"name"`
	// 背景颜色
	BgStyle string `json:"bgStyle"`
	// 字体颜色
	FontStyle string `json:"fontStyle"`
}

// 任务面板页展示的工时信息
type HomeIssueWorkHourInfo struct {
	// 任务的预估工时，单位：小时。
	PredictWorkHour string `json:"predictWorkHour"`
	// 任务的实际工时，单位：小时。
	ActualWorkHour string `json:"actualWorkHour"`
	// 预估工时详情列表
	PredictList []*PredictListItem `json:"predictList"`
	// 实际工时详情列表
	ActualList []*ActualListItem `json:"actualList"`
}

// 导入任务
type ImportIssuesReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 项目类型id
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
	// 迭代id
	IterationID *int64 `json:"iterationId"`
	// excel地址
	URL string `json:"url"`
	// url类型, 1 网址，2 本地dist路径
	URLType int `json:"urlType"`
}

type IndustryListResp struct {
	List []*IndustryResp `json:"list"`
}

type IndustryResp struct {
	// 主键
	ID int64 `json:"id"`
	// 名字
	Name string `json:"name"`
	// 中文名
	Cname string `json:"cname"`
}

type InitExistOrgReq struct {
	NeedOcrConfig    int  `json:"needOcrConfig"`
	NeedSummaryTable int  `json:"needSummaryTable"`
	NeedPriority     int  `json:"needPriority"`
	NeedSetToPaid    *int `json:"needSetToPaid"`
	// 登录状态。-1 表示非登录状态下调用接口；1 登录状态。默认 1。
	LoginMode *int   `json:"loginMode"`
	OrgID     *int64 `json:"OrgId"`
	UserID    *int64 `json:"UserId"`
	// 查询 token 对应的一些信息。0不查询，1表示查询。默认 0。
	QueryAuthTokenInfo *int `json:"QueryAuthTokenInfo"`
}

type InitExistOrgResp struct {
	StrInfo string `json:"strInfo"`
}

// 初始化飞书账号
type InitFeiShuAccountReq struct {
	// codeToken
	CodeToken string `json:"codeToken"`
}

type InternalAuthResp struct {
	HasPermission interface{} `json:"hasPermission"`
}

// 邀请成员时，传入的单个成员信息
type InviteUserData struct {
	// 邮箱
	Email string `json:"email"`
	// 姓名（再次邀请时不用传了）
	Name string `json:"name"`
}

// 被邀请的用户信息
type InviteUserInfo struct {
	// 用户id
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 邮箱
	Email string `json:"email"`
	// 邀请时间
	InviteTime time.Time `json:"inviteTime"`
	// 是否24h内已邀请
	IsInvitedRecent bool `json:"isInvitedRecent"`
}

// 邀请的成员列表接口入参
type InviteUserListReq struct {
	// 页码（选填，不填为全部）
	Page int `json:"page"`
	// 数量（选填，不填为全部）
	Size int `json:"size"`
}

// 邀请的成员列表接口返回
type InviteUserListResp struct {
	// 总数
	Total *int64 `json:"total"`
	// 列表
	List []*InviteUserInfo `json:"list"`
}

// 邀请1个或多个成员接口的入参
type InviteUserReq struct {
	// 邀请的成员列表
	Data []*InviteUserData `json:"data"`
}

// 邀请成员的结果数据
type InviteUserResp struct {
	// 成功的邮箱
	SuccessEmail []string `json:"successEmail"`
	// 已邀请的邮箱
	InvitedEmail []string `json:"invitedEmail"`
	// 已经是用户的邮箱
	IsUserEmail []string `json:"isUserEmail"`
	// 不符合规范的邮箱
	InvalidEmail []string `json:"invalidEmail"`
}

// 任务结构体
type Issue struct {
	// 任务id
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 任务code
	Code string `json:"code"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 项目对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 标题
	Title string `json:"title"`
	// 是否归档
	IsFiling int `json:"isFiling"`
	// 负责人id
	Owner int64 `json:"owner"`
	// 优先级id
	PriorityID int64 `json:"priorityId"`
	// 来源
	SourceID int64 `json:"sourceId"`
	// 问题类型id
	IssueObjectTypeID int64 `json:"issueObjectTypeId"`
	// 性质id
	PropertyID int64 `json:"propertyId"`
	// 计划开始时间
	PlanStartTime time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime time.Time `json:"planEndTime"`
	// 实际开始时间
	StartTime time.Time `json:"startTime"`
	// 实际结束时间
	EndTime time.Time `json:"endTime"`
	// 计划工时
	PlanWorkHour int `json:"planWorkHour"`
	// 迭代id
	IterationID int64 `json:"iterationId"`
	// 版本id
	VersionID int64 `json:"versionId"`
	// 模块id
	ModuleID int64 `json:"moduleId"`
	// 父任务id
	ParentID int64 `json:"parentId"`
	// 父任务标题
	ParentTitle string `json:"parentTitle"`
	// 父任务信息
	ParentInfo []*ParentInfo `json:"parentInfo"`
	// 备注
	Remark *string `json:"remark"`
	// 备注详情
	RemarkDetail *string `json:"remarkDetail"`
	// 状态id
	Status int64 `json:"status"`
	// 创建者
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新者
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 排序
	Sort int64 `json:"sort"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
	// 审批状态（1待审批3审批通过）
	AuditStatus int `json:"auditStatus"`
	// 影响的id集合
	IssueIds []int64 `json:"issueIds"`
}

// 项目和任务信息统计
type IssueAndProjectCountStatResp struct {
	// 项目未完成的数量
	ProjectNotCompletedCount int64 `json:"projectNotCompletedCount"`
	// 任务未完成的数量
	IssueNotCompletedCount int64 `json:"issueNotCompletedCount"`
	// 参与项目数
	ParticipantsProjectCount int64 `json:"participantsProjectCount"`
	// 参与归档项目数
	FilingParticipantsProjectCount int64 `json:"filingParticipantsProjectCount"`
}

// 任务分配信息
type IssueAssignRankInfo struct {
	// 姓名
	Name string `json:"name"`
	// 用户头像
	Avatar string `json:"avatar"`
	// 工号：企业下唯一
	EmplID string `json:"emplId"`
	// 分配的未完成的任务数量
	IncompleteissueCount int64 `json:"incompleteissueCount"`
}

// 任务分配排行请求结构体
type IssueAssignRankReq struct {
	// 项目Id
	ProjectID int64 `json:"projectId"`
	// rank数量， 1 <= rankTop <= 100， 默认为5
	RankTop *int `json:"rankTop"`
}

// 子任务创建结构体
type IssueChildren struct {
	// 名称
	Title string `json:"title"`
	// 负责人
	OwnerID int64 `json:"ownerId"`
	// 类型id，问题，需求....
	TypeID *int64 `json:"typeId"`
	// 优先级
	PriorityID int64 `json:"priorityId"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 计划工作时长
	PlanWorkHour *int `json:"planWorkHour"`
	// 备注
	Remark *string `json:"remark"`
	// 状态id
	StatusID *int64 `json:"statusId"`
	// 关联的标签列表
	Tags []*IssueTagReqInfo `json:"tags"`
	// 关注人
	FollowerIds []int64 `json:"followerIds"`
	// 关联的附件id列表
	ResourceIds []int64 `json:"resourceIds"`
	// 问题对象类型id
	IssueObjectID *int64 `json:"issueObjectId"`
	// 来源id
	IssueSourceID *int64 `json:"issueSourceId"`
	// 性质id
	IssuePropertyID *int64 `json:"issuePropertyId"`
	// 子任务
	Children []*IssueChildren `json:"children"`
	// 自定义字段
	CustomField []*UpdateIssueCustionFieldData `json:"customField"`
	// 审批人
	AuditorIds []int64 `json:"auditorIds"`
	// 无码入参
	LessCreateIssueReq map[string]interface{} `json:"lessCreateIssueReq"`
}

// 每日个人完成图数据统计响应数据
type IssueDailyPersonalWorkCompletionStatData struct {
	// 日期
	StatDate string `json:"statDate"`
	// 完成数量
	CompletedCount int64 `json:"completedCount"`
}

// 每日个人完成图数据统计请求结构体
type IssueDailyPersonalWorkCompletionStatReq struct {
	// 开始时间, 开始时间和结束时间可以不传，默认七天
	StartDate *time.Time `json:"startDate"`
	// 结束时间
	EndDate *time.Time `json:"endDate"`
}

// 每日个人完成图数据统计响应结构体
type IssueDailyPersonalWorkCompletionStatResp struct {
	// 数据列表
	List []*IssueDailyPersonalWorkCompletionStatData `json:"list"`
}

// 单个任务信息详情结构体
type IssueInfo struct {
	// 任务信息
	Issue *Issue `json:"issue"`
	// 项目信息
	Project *HomeIssueProjectInfo `json:"project"`
	// 状态信息
	Status *HomeIssueStatusInfo `json:"status"`
	// 优先级信息
	Priority *HomeIssuePriorityInfo `json:"priority"`
	// 执行人信息
	Owner *UserIDInfo `json:"owner"`
	// 执行人信息
	CreatorInfo *UserIDInfo `json:"creatorInfo"`
	// 参与人
	ParticipantInfos []*UserIDInfo `json:"participantInfos"`
	// 关注人
	FollowerInfos []*UserIDInfo `json:"followerInfos"`
	// 下一个状态
	NextStatus []*HomeIssueStatusInfo `json:"nextStatus"`
	// 标签信息
	Tags []*HomeIssueTagInfo `json:"tags"`
	// 来源信息
	SourceInfo *IssueSourceInfo `json:"sourceInfo"`
	// 严重程度信息
	PropertyInfo *IssuePropertyInfo `json:"propertyInfo"`
	// 类型信息
	TypeInfo *IssueObjectTypeInfo `json:"typeInfo"`
	// 迭代名称
	IterationName string `json:"iterationName"`
	// 子任务数量
	ChildsNum int64 `json:"childsNum"`
	// 子任务已完成数量
	ChildsFinishedNum int64 `json:"childsFinishedNum"`
	// 任务类型名称
	ProjectObjectTypeName string `json:"projectObjectTypeName"`
	// 状态列表
	AllStatus []*HomeIssueStatusInfo `json:"allStatus"`
	// 关联任务数量
	RelateIssueCount int64 `json:"relateIssueCount"`
	// 关联资源数量
	RelateResourceCount int64 `json:"relateResourceCount"`
	// 关联评论数量
	RelateCommentCount int64 `json:"relateCommentCount"`
	// 自定义字段结果
	CustomField []*CustomValue `json:"customField"`
	// 工时信息
	WorkHourInfo *HomeIssueWorkHourInfo `json:"workHourInfo"`
	// 确认人信息
	AuditorsInfo []*UserIDInfoExtraForIssueAudit `json:"auditorsInfo"`
	// 上次任务审批催办时间(时间戳0表示最近没有催办(可以催办))
	LastUrgeTime int64 `json:"lastUrgeTime"`
	// 上次**任务**的催办时间(时间戳0表示最近没有催办(可以催办))
	LastUrgeTimeForIssue int64 `json:"lastUrgeTimeForIssue"`
	// 无码数据
	LessData map[string]interface{} `json:"lessData"`
}

type IssueInfoNotDeleteReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
}

// 任务列表响应结构体
type IssueList struct {
	// 总数量
	Total int64 `json:"total"`
	// 任务列表
	List []*Issue `json:"list"`
}

// 问题对象类型结构体
type IssueObjectType struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 类型名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}

// 任务-类型信息结构体
type IssueObjectTypeInfo struct {
	// 类型id
	ID int64 `json:"id"`
	// 类型名
	Name string `json:"name"`
}

// 问题对象类型列表响应结构体
type IssueObjectTypeList struct {
	Total int64              `json:"total"`
	List  []*IssueObjectType `json:"list"`
}

// 获取任务类型列表请求结构体
type IssueObjectTypesReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 项目对象类型id
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
}

// 问题性质结构体
type IssueProperty struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 类型名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}

// 任务-严重程度结构体
type IssuePropertyInfo struct {
	// id
	ID int64 `json:"id"`
	// 来源名称
	Name string `json:"name"`
}

// 问题对象类型列表响应结构体
type IssuePropertyList struct {
	Total int64              `json:"total"`
	List  []*IssueObjectType `json:"list"`
}

// 获取任务性质列表请求结构体
type IssuePropertysReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 项目对象类型id
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
}

// 任务报表响应结构体
type IssueReportResp struct {
	// 总数量
	Total int64 `json:"total"`
	// 分享id
	ShareID string `json:"shareId"`
	// 报表所属者名称
	ReportUserName string `json:"reportUserName"`
	// 开始时间
	StartTime string `json:"startTime"`
	// 结束时间
	EndTime string `json:"endTime"`
	// 任务信息
	List []*HomeIssueInfo `json:"list"`
}

// 任务的简化信息
type IssueRestInfo struct {
	// 任务id
	ID int64 `json:"id"`
	// 任务标题
	Title string `json:"title"`
	// 负责人
	OwnerID int64 `json:"ownerId"`
	// 负责人名称
	OwnerName string `json:"ownerName"`
	// 负责人头像
	OwnerAvatar string `json:"ownerAvatar"`
	// 是否已被删除，为true则代表被组织移除
	OwnerIsDeleted bool `json:"ownerIsDeleted"`
	// 是否已被禁用, 为true则代表被组织禁用
	OwnerIsDisabled bool `json:"ownerIsDisabled"`
	// 是否已完成
	Finished bool `json:"finished"`
	// 状态id
	StatusID int64 `json:"statusId"`
	// 任务栏id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 迭代id
	IterationID int64 `json:"iterationId"`
	// 状态名称
	StatusName string `json:"statusName"`
	// 计划结束时间
	PlanEndTime time.Time `json:"planEndTime"`
	// 计划开始时间
	PlanStartTime time.Time `json:"planStartTime"`
	// 完成时间
	EndTime time.Time `json:"endTime"`
	// 优先级信息
	PriorityInfo *HomeIssuePriorityInfo `json:"priorityInfo"`
	// 状态信息
	StatusInfo *HomeIssueStatusInfo `json:"statusInfo"`
	// 任务栏名称
	ProjectObjectTypeName string `json:"projectObjectTypeName"`
	// 迭代名称
	IterationName string `json:"iterationName"`
	// 状态信息
	AllStatus []*HomeIssueStatusInfo `json:"allStatus"`
	// 关联状态(1关联2被关联)
	Type int `json:"type"`
	// 审批状态（1待审批3审批通过）
	AuditStatus int `json:"auditStatus"`
	// 项目类型
	ProjectTypeID int64 `json:"projectTypeId"`
}

// 任务简单信息请求结构体（任务详情中的子任务信息）
type IssueRestInfoReq struct {
	// 状态,1:未完成，2：已完成
	Status *int `json:"status"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 父任务id
	ParentID *int64 `json:"parentId"`
	// 任务id集合
	IssueIds []*int64 `json:"issueIds"`
}

// 任务简单信息响应结构体（任务详情中的子任务信息）
type IssueRestInfoResp struct {
	// 总数量
	Total int64 `json:"total"`
	// 任务简单信息列表
	List []*IssueRestInfo `json:"list"`
}

// 问题来源结构体
type IssueSource struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}

// 任务-来源信息结构体
type IssueSourceInfo struct {
	// id
	ID int64 `json:"id"`
	// 来源名称
	Name string `json:"name"`
}

// 问题来源列表响应结构体
type IssueSourceList struct {
	Total int64          `json:"total"`
	List  []*IssueSource `json:"list"`
}

// 获取任务来源列表请求结构体
type IssueSourcesReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 项目对象类型id
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
}

type IssueStatByObjectType struct {
	// 对象类型id
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
	// 对象类型名称
	ProjectObjectTypeName *string `json:"projectObjectTypeName"`
	// 数量
	Total int64 `json:"total"`
}

type IssueStatusTypeStatDetailResp struct {
	// 未开始的统计
	NotStart []*IssueStatByObjectType `json:"notStart"`
	// 进行中的统计
	Processing []*IssueStatByObjectType `json:"processing"`
	// 已完成的统计
	Completed []*IssueStatByObjectType `json:"completed"`
}

// 任务状态数量统计请求结构体
type IssueStatusTypeStatReq struct {
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 迭代id
	IterationID *int64 `json:"iterationId"`
	// 关联类型：1我负责的2我参与的3我关注的4我发起的5我确认的
	RelationType *int `json:"relationType"`
}

// 任务状态数量统计响应结构体
type IssueStatusTypeStatResp struct {
	// 状态为未开始的数量
	NotStartTotal int64 `json:"notStartTotal"`
	// 状态为进行中的数量
	ProcessingTotal int64 `json:"processingTotal"`
	// 状态为已完成的数量
	CompletedTotal int64 `json:"completedTotal"`
	// 今日完成数
	CompletedTodayTotal int64 `json:"completedTodayTotal"`
	// 状态为逾期的数量
	OverdueTotal int64 `json:"overdueTotal"`
	// 状态为明日逾期
	OverdueTomorrowTotal int64 `json:"overdueTomorrowTotal"`
	// 逾期完成
	OverdueCompletedTotal int64 `json:"overdueCompletedTotal"`
	// 任务总数
	Total int64 `json:"total"`
	// 今日到期
	OverdueTodayTotal int64 `json:"overdueTodayTotal"`
	// 即将到期
	BeAboutToOverdueSum int64 `json:"beAboutToOverdueSum"`
	// 指派给我的任务
	TodayCount int64 `json:"todayCount"`
	// 今日创建
	TodayCreateCount int64 `json:"todayCreateCount"`
	// @我的数量
	CallMeTotal int64 `json:"callMeTotal"`
	// 待确认的任务数量
	WaitConfirmedTotal int64 `json:"waitConfirmedTotal"`
	// 概览
	List []*StatCommon `json:"list"`
}

// 任务标签结构体
type IssueTagReqInfo struct {
	// 标签id
	ID int64 `json:"id"`
	// 标签名称
	Name string `json:"name"`
}

// 获取任务标签列表请求结构体
type IssueTagsReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
}

// 一条工时记录信息
type IssueWorkHours struct {
	// 工时记录主键
	ID int64 `json:"id"`
	// 记录类型：1预估工时记录，2实际工时记录，3子预估工时
	Type int64 `json:"type"`
	// 工作者id
	WorkerID int64 `json:"workerId"`
	// 所需工时时间，单位：小时
	NeedTime string `json:"needTime"`
	// 开始时间，秒级时间戳。
	StartTime int64 `json:"startTime"`
	// 工时记录的结束时间，秒级时间戳。
	EndTime *int64 `json:"endTime"`
	// 工时记录的内容，工作内容
	Desc *string `json:"desc"`
}

// 迭代结构体
type Iteration struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 名称
	Name string `json:"name"`
	// 负责人
	Owner int64 `json:"owner"`
	// 排序
	Sort int64 `json:"sort"`
	// 版本
	VersionID int64 `json:"versionId"`
	// 计划开始时间
	PlanStartTime time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime time.Time `json:"planEndTime"`
	// 预估工时
	PlanWorkHour int `json:"planWorkHour"`
	// 故事点
	StoryPoint int `json:"storyPoint"`
	// 描述
	Remark *string `json:"remark"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 负责人信息
	OwnerInfo *HomeIssueOwnerInfo `json:"ownerInfo"`
	// 状态信息
	StatusInfo *HomeIssueStatusInfo `json:"statusInfo"`
	// 总任务数量
	AllIssueCount int64 `json:"allIssueCount"`
	// 已完成任务数量
	FinishedIssueCount int64 `json:"finishedIssueCount"`
}

// 迭代信息请求结构体
type IterationInfoReq struct {
	// 迭代id
	ID int64 `json:"id"`
}

// 迭代信息响应结构体
type IterationInfoResp struct {
	// 迭代信息
	Iteration *Iteration `json:"iteration"`
	// 项目信息
	Project *HomeIssueProjectInfo `json:"project"`
	// 状态信息
	Status *HomeIssueStatusInfo `json:"status"`
	// 负责人信息
	Owner *UserIDInfo `json:"owner"`
	// 下一步骤状态列表
	NextStatus []*HomeIssueStatusInfo `json:"nextStatus"`
	// 状态时间信息
	StatusTimeInfo []*StatusTimeInfo `json:"statusTimeInfo"`
}

// 迭代和任务关联请求结构体
type IterationIssueRealtionReq struct {
	// 迭代id
	IterationID int64 `json:"iterationId"`
	// 要添加的任务id列表（除特性任务）
	AddIssueIds []int64 `json:"addIssueIds"`
	// 要移除的任务id列表
	DelIssueIds []int64 `json:"delIssueIds"`
}

// 迭代列表响应结构体
type IterationList struct {
	// 总数量
	Total int64 `json:"total"`
	// 迭代列表
	List []*Iteration `json:"list"`
}

// 迭代列表请求结构体
type IterationListReq struct {
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 名称，模糊搜索
	Name *string `json:"name"`
	// 状态,1:未开始，2：进行中，3：已完成
	StatusType *int `json:"statusType"`
	// 排序（1创建时间正序2创建时间倒序3sort正序4sort倒序,默认4）
	OrderBy *int `json:"orderBy"`
}

// 迭代统计结构体
type IterationStat struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 迭代id
	IterationID int64 `json:"iterationId"`
	// 问题总数
	IssueCount int `json:"issueCount"`
	// 未开始问题数
	IssueWaitCount int `json:"issueWaitCount"`
	// 进行中问题数
	IssueRunningCount int `json:"issueRunningCount"`
	// 已逾期问题数
	IssueOverdueCount int `json:"issueOverdueCount"`
	// 已完成问题数
	IssueEndCount int `json:"issueEndCount"`
	// 需求总数
	DemandCount int `json:"demandCount"`
	// 未开始需求数
	DemandWaitCount int `json:"demandWaitCount"`
	// 进行中需求数
	DemandRunningCount int `json:"demandRunningCount"`
	// 已逾期需求数
	DemandOverdueCount int `json:"demandOverdueCount"`
	// 已完成需求数
	DemandEndCount int `json:"demandEndCount"`
	// 故事点总数
	StoryPointCount int `json:"storyPointCount"`
	// 未开始故事点数
	StoryPointWaitCount int `json:"storyPointWaitCount"`
	// 进行中故事点数
	StoryPointRunningCount int `json:"storyPointRunningCount"`
	// 已逾期故事点数
	StoryPointOverdueCount int `json:"storyPointOverdueCount"`
	// 已完成故事点数
	StoryPointEndCount int `json:"storyPointEndCount"`
	// 任务总数
	TaskCount int `json:"taskCount"`
	// 未开始任务数
	TaskWaitCount int `json:"taskWaitCount"`
	// 进行中任务数
	TaskRunningCount int `json:"taskRunningCount"`
	// 已逾期任务数
	TaskOverdueCount int `json:"taskOverdueCount"`
	// 已完成任务数
	TaskEndCount int `json:"taskEndCount"`
	// 缺陷总数
	BugCount int `json:"bugCount"`
	// 未开始缺陷数
	BugWaitCount int `json:"bugWaitCount"`
	// 进行中缺陷数
	BugRunningCount int `json:"bugRunningCount"`
	// 已逾期缺陷数
	BugOverdueCount int `json:"bugOverdueCount"`
	// 已完成缺陷数
	BugEndCount int `json:"bugEndCount"`
	// 测试任务总数
	TesttaskCount int `json:"testtaskCount"`
	// 未开始测试任务数
	TesttaskWaitCount int `json:"testtaskWaitCount"`
	// 进行中测试任务数
	TesttaskRunningCount int `json:"testtaskRunningCount"`
	// 已逾期测试任务数
	TesttaskOverdueCount int `json:"testtaskOverdueCount"`
	// 已完成测试任务数
	TesttaskEndCount int `json:"testtaskEndCount"`
	// 扩展
	Ext string `json:"ext"`
	// 统计日期
	StatDate time.Time `json:"statDate"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
}

// 迭代统计列表响应结构体
type IterationStatList struct {
	Total int64            `json:"total"`
	List  []*IterationStat `json:"list"`
}

// 迭代统计查询请求
type IterationStatReq struct {
	// 迭代id
	IterationID int64 `json:"iterationId"`
	// 开始时间
	StartDate *time.Time `json:"startDate"`
	// 结束时间
	EndDate *time.Time `json:"endDate"`
}

type IterationStatSimple struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	AllIssues     int64  `json:"allIssues"`
	OverdueIssues int64  `json:"overdueIssues"`
	FinishIssues  int64  `json:"finishIssues"`
}

// 迭代状态类型统计请求结构体
type IterationStatusTypeStatReq struct {
	// 项目id
	ProjectID *int64 `json:"projectId"`
}

// 迭代状态类型统计响应结构体
type IterationStatusTypeStatResp struct {
	// 状态为未开始的数量
	NotStartTotal int64 `json:"notStartTotal"`
	// 状态为进行中的数量
	ProcessingTotal int64 `json:"processingTotal"`
	// 状态为已完成的数量
	CompletedTotal int64 `json:"completedTotal"`
	// 总数量
	Total int64 `json:"total"`
}

type JoinOrgByInviteCodeReq struct {
	InviteCode string `json:"inviteCode"`
}

// 获取JSApi签名请求结构体
type JsAPISignReq struct {
	// 类型:目前只支持:jsapi
	Type string `json:"type"`
	// 路由url
	URL string `json:"url"`
	// dingtalk企业id
	CorpID string `json:"corpId"`
}

// 获取JSApi签名响应结构体
type JsAPISignResp struct {
	// 应用代理id
	AgentID int64 `json:"agentId"`
	// 时间戳
	TimeStamp string `json:"timeStamp"`
	// 随机字符串
	NoceStr string `json:"noceStr"`
	// 签名
	Signature string `json:"signature"`
}

type LessCondsData struct {
	// 类型(between,equal,gt,gte,in,like,lt,lte,not_in,not_like,not_null,is_null,all_in,values_in)
	Type string `json:"type"`
	// 字段类型
	FieldType *string `json:"fieldType"`
	// 值
	Value interface{} `json:"value"`
	// 值（数组）
	Values interface{} `json:"values"`
	// 字段id
	Column string `json:"column"`
	// 左值
	Left interface{} `json:"left"`
	// 右值
	Right interface{} `json:"right"`
	// 嵌套
	Conds []*LessCondsData `json:"conds"`
}

type LessOrder struct {
	// 是否是正序
	Asc bool `json:"asc"`
	// 字段
	Column string `json:"column"`
}

type MemberInfo struct {
	// 成员信息id
	ID *int64 `json:"id"`
	// 成员名称
	Name *string `json:"name"`
	// 成员头像
	Avatar *string `json:"avatar"`
}

type MigrateIssueToLcReq struct {
	OrgIds []int64 `json:"orgIds"`
}

// 新增详细版的预估工时
type NewPredicateWorkHour struct {
	// 工作者id
	WorkerID int64 `json:"workerId"`
	// 预估所需工时时间，单位分钟
	NeedTime string `json:"needTime"`
	// 开始时间，时间戳
	StartTime int64 `json:"startTime"`
	// 工时记录的结束时间，时间戳
	EndTime *int64 `json:"endTime"`
}

// 结构体
type Notice struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 通知类型, 1项目通知,2组织通知,
	Type int `json:"type"`
	// 操作类型
	RelationType string `json:"relationType"`
	// 冗余信息
	Ext string `json:"ext"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// issueId
	IssueID int64 `json:"issueId"`
	// 关联动态id
	TrendsID int64 `json:"trendsId"`
	// 通知内容
	Content string `json:"content"`
	// 被通知人
	Noticer int64 `json:"noticer"`
	// 状态, 1未读,2已读
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 创建人信息
	CreatorInfo *UserIDInfo `json:"creatorInfo"`
	// 项目名称
	ProjectName string `json:"projectName"`
	// 任务名称
	IssueName string `json:"issueName"`
	// 父任务id（没有则为0）
	ParentIssueID int64 `json:"parentIssueId"`
}

// 通知条数
type NoticeCountResp struct {
	Total int64 `json:"total"`
}

// 列表响应结构体
type NoticeList struct {
	Total int64     `json:"total"`
	List  []*Notice `json:"list"`
}

type NoticeListReq struct {
	// 通知类型, 1项目通知,2组织通知,
	Type *int `json:"type"`
	// 是否是@我的相关notice(不传或传1表示普通通知，2表示@我的)
	IsCallMe *int `json:"isCallMe"`
	// 上次分页的最后一条id
	LastID *int64 `json:"lastId"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 任务id
	IssueID *int64 `json:"issueId"`
}

// 查询返回：一条实际工时记录的信息
type OneActualWorkHourRecord struct {
	// 工时记录主键
	ID int64 `json:"id"`
	// 记录类型：1预估工时记录，2实际工时记录，3子预估工时
	Type int64 `json:"type"`
	// 工时记录的工作者、执行者
	Worker *WorkHourWorker `json:"worker"`
	// 所需工时时间，单位：小时
	NeedTime string `json:"needTime"`
	// 开始时间，时间戳
	StartTime int64 `json:"startTime"`
	// 工时记录的结束时间，秒级时间戳。
	EndTime int64 `json:"endTime"`
	// 创建者名称
	CreatorName string `json:"creatorName"`
	// 创建时间
	CreateTime int64 `json:"createTime"`
	// 工时记录的内容，工作内容
	Desc string `json:"desc"`
	// 是否允许当前用户修改。1：可以修改；0：不允许修改。
	IsEnable int `json:"isEnable"`
}

// 一个员工在某一日期的实际工时信息
type OneDateWorkHour struct {
	// 工时日期
	Date string `json:"date"`
	// 星期几。特殊的是：0表示星期天
	WeekDay int64 `json:"weekDay"`
	// 工时时间，单位：小时
	Time string `json:"time"`
}

// 一个员工的在若干个日期内的工时统计信息
type OnePersonWorkHourStatisticInfo struct {
	// 员工id
	WorkerID int64 `json:"workerId"`
	// 员工姓名
	Name string `json:"name"`
	// 预估总工时，单位：小时
	PredictHourTotal string `json:"predictHourTotal"`
	// 实际总工时，单位：小时
	ActualHourTotal string `json:"actualHourTotal"`
	// 在一些日期内的实际工时信息
	DateWorkHourList []*OneDateWorkHour `json:"dateWorkHourList"`
}

// 查询返回：一条预估工时记录的信息
type OneWorkHourRecord struct {
	// 工时记录主键
	ID int64 `json:"id"`
	// 记录类型：1预估工时记录，2实际工时记录，3子预估工时
	Type int64 `json:"type"`
	// 工时记录的工作者、执行者
	Worker *WorkHourWorker `json:"worker"`
	// 所需工时时间，单位：小时
	NeedTime string `json:"needTime"`
	// 开始时间，时间戳
	StartTime int64 `json:"startTime"`
	// 工时记录的结束时间，秒级时间戳。
	EndTime int64 `json:"endTime"`
	// 工时记录的内容，工作内容
	Desc string `json:"desc"`
	// 是否允许当前用户修改。1：可以修改；0：不允许修改。
	IsEnable int `json:"isEnable"`
}

type OperateProjectResp struct {
	// 是否成功
	IsSuccess interface{} `json:"isSuccess"`
}

// 组织配置响应结构体
type OrgConfig struct {
	// id
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 付费级别1通用免费，2标准版
	PayLevel int `json:"payLevel"`
	// 付费开始时间
	PayStartTime time.Time `json:"payStartTime"`
	// 付费结束时间
	PayEndTime time.Time `json:"payEndTime"`
	// 付费级别实际（1免费2标准3试用）
	PayLevelTrue int `json:"payLevelTrue"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 组织人数
	OrgMemberNumber int64 `json:"orgMemberNumber"`
	// 是否是灰度企业
	IsGrayLevel bool `json:"isGrayLevel"`
	// 汇总表id
	SummaryAppID string `json:"summaryAppId"`
	// 展示基础设置
	BasicShowSetting *BasicShowSetting `json:"basicShowSetting"`
	// 企业自定义logo
	Logo string `json:"logo"`
}

type OrgInfoForChosen struct {
	// 组织id
	ID int64 `json:"id"`
	// 组织名称
	Name string `json:"name"`
}

type OrgProjectMemberInfoResp struct {
	// 用户id
	UserID int64 `json:"userId"`
	// 外部用户id
	OutUserID string `json:"outUserId"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 外部组织id
	OutOrgID string `json:"outOrgId"`
	// 姓名
	Name string `json:"name"`
	// 姓名拼音（可能为空）
	NamePy *string `json:"namePy"`
	// 用户头像
	Avatar string `json:"avatar"`
	// 是否有外部信息
	HasOutInfo bool `json:"hasOutInfo"`
	// 是否有组织外部信息
	HasOrgOutInfo bool `json:"hasOrgOutInfo"`
	// 组织用户是否删除
	OrgUserIsDelete int `json:"orgUserIsDelete"`
	// 组织用户状态
	OrgUserStatus int `json:"orgUserStatus"`
	// 组织用户check状态
	OrgUserCheckStatus int `json:"orgUserCheckStatus"`
}

type OrgProjectMemberListResp struct {
	// 总数
	Total int64 `json:"total"`
	// 数据
	List []*OrgProjectMemberInfoResp `json:"list"`
}

type OrgProjectMemberReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type OrgProjectMemberResp struct {
	Owner        *OrgProjectMemberInfoResp   `json:"owner"`
	Participants []*OrgProjectMemberInfoResp `json:"participants"`
	Follower     []*OrgProjectMemberInfoResp `json:"follower"`
	AllMembers   []*OrgProjectMemberInfoResp `json:"allMembers"`
}

// 列表请求结构体
type OrgUserListReq struct {
	// 审核状态,1待审核,2审核通过,3审核不过(成员管理取审核通过的，成员审核取待审核和审核不过的)
	CheckStatus []int `json:"checkStatus"`
	// 使用状态,1已使用,2未使用
	UseStatus *int `json:"useStatus"`
	// 企业用户状态, 1可用,2禁用
	Status *int `json:"status"`
	// 姓名
	Name *string `json:"name"`
	// 邮箱
	Email *string `json:"email"`
	// 手机号
	Mobile *string `json:"mobile"`
}

type OrganizationInfoReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
}

// 组织设置入参
type OrganizationInfoResp struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 组织名称
	OrgName string `json:"orgName"`
	// code
	Code string `json:"code"`
	// 组织网站
	WebSite string `json:"webSite"`
	// 所属行业
	IndustryID int64 `json:"industryId"`
	// 所属行业中文名
	IndustryName string `json:"industryName"`
	// 组织规模
	Scale string `json:"scale"`
	// 所在国家
	CountryID int64 `json:"countryId"`
	// 所在国家中文名
	CountryCname string `json:"countryCname"`
	// 所在省份
	ProvinceID int64 `json:"provinceId"`
	// 所在省份中文名
	ProvinceCname string `json:"provinceCname"`
	// 所在城市
	CityID int64 `json:"cityId"`
	// 所在城市中文名
	CityCname string `json:"cityCname"`
	// 组织地址
	Address string `json:"address"`
	// 组织logo地址
	LogoURL string `json:"logoUrl"`
	// 组织负责人
	Owner int64 `json:"owner"`
	// 负责人信息
	OwnerInfo *UserIDInfo `json:"ownerInfo"`
	// 备注
	Remark string `json:"remark"`
	// 第三方企业编号
	ThirdCode string `json:"thirdCode"`
}

// 结构体
type OrganizationUser struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 用户id
	UserID int64 `json:"userId"`
	// 审核状态,1待审核,2审核通过,3审核不过
	CheckStatus int `json:"checkStatus"`
	// 使用状态,1已使用,2未使用
	UseStatus int `json:"useStatus"`
	// 企业用户状态, 1可用,2禁用
	Status int `json:"status"`
	// 状态变更人id
	StatusChangerID int64 `json:"statusChangerId"`
	// 状态变更时间
	StatusChangeTime time.Time `json:"statusChangeTime"`
	// 审核人id
	AuditorID int64 `json:"auditorId"`
	// 审核时间
	AuditTime time.Time `json:"auditTime"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 用户信息
	UserInfo *PersonalInfo `json:"userInfo"`
	// 审核人信息
	AuditorInfo *PersonalInfo `json:"auditorInfo"`
	// 用户角色
	UserRole *UserRoleInfo `json:"userRole"`
}

// Oss申请signUrl请求结构体
type OssApplySignURLReq struct {
	// 文件url
	URL string `json:"url"`
}

// Oss申请signUrl响应结构体
type OssApplySignURLResp struct {
	// signUrl
	SignURL string `json:"signUrl"`
}

// Oss Post Policy策略请求结构体
type OssPostPolicyReq struct {
	// 策略类型, 1: 项目封面，2：任务资源（需要callback）, 3：导入任务的excel， 4：项目文件（需要callback），5：兼容测试， 6:用户头像，11：excel导入成员
	PolicyType int `json:"policyType"`
	// 如果policyType为1和2和3，那么projectId必传(创建场景传0)
	ProjectID *int64 `json:"projectId"`
	// 如果policyType为2，那么issueId必传
	IssueID *int64 `json:"issueId"`
	// 目录id, policy为4的时候必填
	FolderID *int64 `json:"folderId"`
}

// Oss Post Policy策略响应结构体
type OssPostPolicyResp struct {
	// policy
	Policy string `json:"policy"`
	// 签名
	Signature string `json:"signature"`
	// 文件上传目录
	Dir string `json:"dir"`
	// 有效期
	Expire string `json:"expire"`
	// access Id
	AccessID string `json:"accessId"`
	// Host
	Host string `json:"host"`
	// Region
	Region string `json:"region"`
	// bucket名称
	Bucket string `json:"bucket"`
	// 文件名
	FileName string `json:"fileName"`
	// 文件最大限制
	MaxFileSize int64 `json:"maxFileSize"`
	// callback回调，为空说明不需要回调
	Callback string `json:"callback"`
}

type ParentInfo struct {
	// id
	ID int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// code
	Code string `json:"code"`
}

type PayLimitNumResp struct {
	// 项目数量
	ProjectNum int64 `json:"projectNum"`
	// 任务数量
	IssueNum int64 `json:"issueNum"`
	// 文件大小
	FileSize int64 `json:"fileSize"`
}

// 结构体
type Permission struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 编号,仅支持大写英文字母
	Code string `json:"code"`
	// 名称
	Name string `json:"name"`
	// 父id
	ParentID int64 `json:"parentId"`
	// 权限项类型,1系统,2组织,3项目
	Type int `json:"type"`
	// 权限路径
	Path string `json:"path"`
	// 是否显示,1是,2否
	IsShow int `json:"isShow"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}

// 列表响应结构体
type PermissionList struct {
	Total int64         `json:"total"`
	List  []*Permission `json:"list"`
}

// 结构体
type PermissionOperation struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 权限项id
	PermissionID int64 `json:"permissionId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 操作编号,多个半角逗号分隔
	OperationCodes string `json:"operationCodes"`
	// 描述
	Remark string `json:"remark"`
	// 是否显示,1是,2否
	IsShow int `json:"isShow"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}

// 列表响应结构体
type PermissionOperationList struct {
	Total int64                  `json:"total"`
	List  []*PermissionOperation `json:"list"`
}

type PermissionOperationListResp struct {
	// 权限项信息
	PermissionInfo *Permission `json:"permissionInfo"`
	// 权限操作项信息
	OperationList []*PermissionOperation `json:"operationList"`
	// 角色拥有的操作项权限id
	PermissionHave []int64 `json:"permissionHave"`
}

// 个人信息
type PersonalInfo struct {
	// 主键
	ID int64 `json:"id"`
	// 工号
	EmplID *string `json:"emplId"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 组织名称
	OrgName string `json:"orgName"`
	// 组织code
	OrgCode string `json:"orgCode"`
	// 名称
	Name string `json:"name"`
	// 第三方名称
	ThirdName string `json:"thirdName"`
	// 登录名
	LoginName string `json:"loginName"`
	// 登录名编辑次数
	LoginNameEditCount int `json:"loginNameEditCount"`
	// 邮箱
	Email string `json:"email"`
	// 电话
	Mobile string `json:"mobile"`
	// 生日
	Birthday time.Time `json:"birthday"`
	// 性别
	Sex int `json:"sex"`
	// 剩余使用时长
	Rimanente int `json:"rimanente"`
	// 付费等级
	Level int `json:"level"`
	// 付费等级名
	LevelName string `json:"levelName"`
	// 头像
	Avatar string `json:"avatar"`
	// 来源
	SourceChannel string `json:"sourceChannel"`
	// 语言
	Language string `json:"language"`
	// 座右铭
	Motto string `json:"motto"`
	// 上次登录ip
	LastLoginIP string `json:"lastLoginIp"`
	// 上次登录时间
	LastLoginTime time.Time `json:"lastLoginTime"`
	// 登录失败次数
	LoginFailCount int `json:"loginFailCount"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 密码是否设置过(1已设置0未设置)
	PasswordSet int `json:"passwordSet"`
	// 是否需要提醒（1需要2不需要）
	RemindBindPhone int `json:"remindBindPhone"`
	// 是否是超管
	IsAdmin bool `json:"isAdmin"`
	// 是否是管理员
	IsManager bool `json:"isManager"`
	// 权限
	Functions []string `json:"functions"`
	// 一些额外数据，如：观看新手指引的状态
	ExtraDataMap map[string]interface{} `json:"extraDataMap"`
}

// 预估工时详情列表单个对象
type PredictListItem struct {
	// 工时执行人名字
	Name string `json:"name"`
	// 工时，单位：小时。
	WorkHour string `json:"workHour"`
}

// 优先级结构体
type Priority struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,全局的填0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 类型,1项目优先级,2:需求/任务等优先级
	Type int `json:"type"`
	// 排序
	Sort int `json:"sort"`
	// 背景颜色
	BgStyle string `json:"bgStyle"`
	// 字体颜色
	FontStyle string `json:"fontStyle"`
	// 是否默认,1是,2否
	IsDefault int `json:"isDefault"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 优先级列表响应结构体
type PriorityList struct {
	Total int64       `json:"total"`
	List  []*Priority `json:"list"`
}

// 优先级列表
type PriorityListReq struct {
	// 类型,1项目优先级,2:需求/任务等优先级
	Type *int `json:"type"`
}

// 流程状态结构体
type ProcessStatus struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,全局的填0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 背景颜色
	BgStyle string `json:"bgStyle"`
	// 字体颜色
	FontStyle string `json:"fontStyle"`
	// 状态类型,1未开始,2进行中,3已完成
	Type int `json:"type"`
	// 状态类别,1项目状态,2迭代状态,3问题状态,
	Category int `json:"category"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 流程状态列表响应结构体
type ProcessStatusList struct {
	Total int64            `json:"total"`
	List  []*ProcessStatus `json:"list"`
}

type Project struct {
	// 主键
	ID int64 `json:"id"`
	// 项目对应的应用 id（无码系统）
	AppID string `json:"appId"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 编号
	Code string `json:"code"`
	// 名称
	Name string `json:"name"`
	// 前缀编号
	PreCode string `json:"preCode"`
	// 项目负责人
	Owner int64 `json:"owner"`
	// 项目类型
	ProjectTypeID int64 `json:"projectTypeId"`
	// 项目优先级
	PriorityID int64 `json:"priorityId"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 项目公开性,1公开,2私有
	PublicStatus int `json:"publicStatus"`
	// 项目标识
	ResourceID int64 `json:"resourceId"`
	// 是否归档,1归档,2未归档
	IsFiling int `json:"isFiling"`
	// 描述
	Remark string `json:"remark"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 状态类型,1未开始,2进行中,3已完成
	StatusType int `json:"statusType"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
	// 创建人信息
	CreatorInfo *UserIDInfo `json:"creatorInfo"`
	// 负责人信息
	OwnerInfo *UserIDInfo `json:"ownerInfo"`
	// 负责人信息集合
	OwnersInfo []*UserIDInfo `json:"ownersInfo"`
	// 成员信息
	MemberInfo []*UserIDInfo `json:"memberInfo"`
	// 关注人信息
	FollowerInfo []*UserIDInfo `json:"followerInfo"`
	// 封面地址
	ResourcePath string `json:"resourcePath"`
	// 封面缩略图
	ResourceCompressedPath string `json:"resourceCompressedPath"`
	// 所有任务数量
	AllIssues int64 `json:"allIssues"`
	// 已完成任务数量
	FinishIssues int64 `json:"finishIssues"`
	// 逾期任务数量
	OverdueIssues int64 `json:"overdueIssues"`
	// 最近一次迭代数据
	IterationStat *IterationStatSimple `json:"iterationStat"`
	// 流程状态
	AllStatus []*HomeIssueStatusInfo `json:"allStatus"`
	// 项目类型名称
	ProjectTypeName string `json:"projectTypeName"`
	// 项目类型LangCode，ProjectType.NormalTask  普通任务项目, ProjectType.Agile  敏捷研发项目
	ProjectTypeLangCode string `json:"projectTypeLangCode"`
	// 是否同步到飞书日历(1是2否,默认否)
	IsSyncOutCalendar int `json:"isSyncOutCalendar"`
	// 是否收藏关注(1是0否)
	IsStar int `json:"isStar"`
	// 与我相关的未完成的
	RelateUnfinish int64 `json:"relateUnfinish"`
	// 项目对象类型列表
	ProjectObjectTypeList []*ProjectObjectTypeRestInfo `json:"projectObjectTypeList"`
	// icon
	Icon string `json:"icon"`
}

type ProjectAttachmentReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 文件类型,0其他,1文档,2图片,3视频,4音频
	FileType *int `json:"fileType"`
	// 文件搜索关键字
	KeyWord *string `json:"keyWord"`
}

type ProjectChatListReq struct {
	// 每页数量(不传默认10条)
	PageSize *int64 `json:"pageSize"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 上次分页最后一条关联id（不传默认第一页）
	LastRelationID *int64 `json:"lastRelationId"`
	// 搜索内容
	Name *string `json:"name"`
}

// 项目日统计结构体
type ProjectDayStat struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 问题总数
	IssueCount int `json:"issueCount"`
	// 未开始问题数
	IssueWaitCount int `json:"issueWaitCount"`
	// 进行中问题数
	IssueRunningCount int `json:"issueRunningCount"`
	// 已逾期问题数
	IssueOverdueCount int `json:"issueOverdueCount"`
	// 已完成问题数
	IssueEndCount int `json:"issueEndCount"`
	// 需求总数
	DemandCount int `json:"demandCount"`
	// 未开始需求数
	DemandWaitCount int `json:"demandWaitCount"`
	// 进行中需求数
	DemandRunningCount int `json:"demandRunningCount"`
	// 已逾期需求数
	DemandOverdueCount int `json:"demandOverdueCount"`
	// 已完成需求数
	DemandEndCount int `json:"demandEndCount"`
	// 故事点总数
	StoryPointCount int `json:"storyPointCount"`
	// 未开始故事点数
	StoryPointWaitCount int `json:"storyPointWaitCount"`
	// 进行中故事点数
	StoryPointRunningCount int `json:"storyPointRunningCount"`
	// 已逾期故事点数
	StoryPointOverdueCount int `json:"storyPointOverdueCount"`
	// 已完成故事点数
	StoryPointEndCount int `json:"storyPointEndCount"`
	// 任务总数
	TaskCount int `json:"taskCount"`
	// 未开始任务数
	TaskWaitCount int `json:"taskWaitCount"`
	// 进行中任务数
	TaskRunningCount int `json:"taskRunningCount"`
	// 已逾期任务数
	TaskOverdueCount int `json:"taskOverdueCount"`
	// 已完成任务数
	TaskEndCount int `json:"taskEndCount"`
	// 缺陷总数
	BugCount int `json:"bugCount"`
	// 未开始缺陷数
	BugWaitCount int `json:"bugWaitCount"`
	// 进行中缺陷数
	BugRunningCount int `json:"bugRunningCount"`
	// 已逾期缺陷数
	BugOverdueCount int `json:"bugOverdueCount"`
	// 已完成缺陷数
	BugEndCount int `json:"bugEndCount"`
	// 测试任务总数
	TesttaskCount int `json:"testtaskCount"`
	// 未开始测试任务数
	TesttaskWaitCount int `json:"testtaskWaitCount"`
	// 进行中测试任务数
	TesttaskRunningCount int `json:"testtaskRunningCount"`
	// 已逾期测试任务数
	TesttaskOverdueCount int `json:"testtaskOverdueCount"`
	// 已完成测试任务数
	TesttaskEndCount int `json:"testtaskEndCount"`
	// 扩展
	Ext string `json:"ext"`
	// 统计日期
	StatDate time.Time `json:"statDate"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 项目日统计列表响应结构体
type ProjectDayStatList struct {
	Total int64             `json:"total"`
	List  []*ProjectDayStat `json:"list"`
}

// 迭代统计查询请求
type ProjectDayStatReq struct {
	// 迭代id
	ProjectID int64 `json:"projectId"`
	// 开始时间
	StartDate *time.Time `json:"startDate"`
	// 结束时间
	EndDate *time.Time `json:"endDate"`
}

type ProjectDetail struct {
	// 详情id
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 公告
	Notice string `json:"notice"`
	// 是否启用工时和故事点,1启用,2不启用
	IsEnableWorkHours int `json:"isEnableWorkHours"`
	// 是否同步到飞书日历(1是2否,默认否)
	IsSyncOutCalendar int       `json:"isSyncOutCalendar"`
	Creator           int64     `json:"creator"`
	CreateTime        time.Time `json:"createTime"`
	Updator           int64     `json:"updator"`
	UpdateTime        time.Time `json:"updateTime"`
}

type ProjectDetailList struct {
	Total int64            `json:"total"`
	List  []*ProjectDetail `json:"list"`
}

type ProjectFieldViewReq struct {
	// 视图类型（1看板2列表3表格）
	ViewType int `json:"viewType"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 任务类型
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
}

type ProjectFieldViewResp struct {
	// 关闭的默认的字段code
	ClosedDefaultFields []string `json:"closedDefaultFields"`
	// 自定义字段
	CustomFields []*CustomField `json:"customFields"`
}

type ProjectFolderReq struct {
	// 父文件夹id
	ParentID *int64 `json:"parentId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
}

// 项目信息结构体
type ProjectInfo struct {
	// 主键
	ID int64 `json:"id"`
	// 项目对应的应用 id（无码系统）
	AppID string `json:"appId"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 编号
	Code string `json:"code"`
	// 名称
	Name string `json:"name"`
	// 前缀编号
	PreCode string `json:"preCode"`
	// 项目负责人
	Owner int64 `json:"owner"`
	// 项目类型
	ProjectTypeID int64 `json:"projectTypeId"`
	// 项目优先级
	PriorityID int64 `json:"priorityId"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 项目公开性,1公开,2私有
	PublicStatus int `json:"publicStatus"`
	// 项目标识
	ResourceID int64 `json:"resourceId"`
	// 是否归档,1归档,2未归档
	IsFiling int `json:"isFiling"`
	// 描述
	Remark string `json:"remark"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 创建人信息
	CreatorInfo *UserIDInfo `json:"creatorInfo"`
	// 负责人信息
	OwnerInfo *UserIDInfo `json:"ownerInfo"`
	// 负责人信息集合
	OwnersInfo []*UserIDInfo `json:"ownersInfo"`
	// 成员信息
	MemberInfo []*UserIDInfo `json:"memberInfo"`
	// 成员部门信息
	MemberDepartmentInfo []*DepartmentSimpleInfo `json:"memberDepartmentInfo"`
	// 关注人信息
	FollowerInfo []*UserIDInfo `json:"followerInfo"`
	// 封面地址
	ResourcePath string `json:"resourcePath"`
	// 所有状态
	AllStatus []*HomeIssueStatusInfo `json:"allStatus"`
	// 是否同步到飞书日历(1是2否,默认否)
	IsSyncOutCalendar int `json:"isSyncOutCalendar"`
	// 针对哪些群体用户，同步到其飞书日历(4：负责人，8：关注人。往后扩展是基于二进制的位值)
	SyncCalendarStatusList []*int `json:"syncCalendarStatusList"`
	// 是否创建群聊（针对于飞书1是2否默认是）
	IsCreateFsChat int `json:"isCreateFsChat"`
	// 是否收藏关注(1是0否)
	IsStar int `json:"isStar"`
	// 项目开启隐私模式的状态值：1 开启；2关闭
	PrivacyStatus int `json:"privacyStatus"`
	// icon
	Icon string `json:"icon"`
}

// 项目信息请求结构体
type ProjectInfoReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
}

// 项目任务关联的状态请求结构体
type ProjectIssueRelatedStatusReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 项目对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
}

type ProjectList struct {
	// 项目数量
	Total int64 `json:"total"`
	// 项目列表
	List []*Project `json:"list"`
}

type ProjectMemberIDListReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 返回的项目成员id，知否需要包含额外的管理员 id。可选，1表示需要 0不需要。默认0。
	IncludeAdmin *int `json:"includeAdmin"`
}

type ProjectMemberIDListResp struct {
	// 部门id
	DepartmentIds []int64 `json:"departmentIds"`
	// 人员id
	UserIds []int64 `json:"userIds"`
}

// 项目对象类型结构体
type ProjectObjectType struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,全局的填0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 编号前缀
	PreCode string `json:"preCode"`
	// 名称
	Name string `json:"name"`
	// 类型,1迭代，2问题
	ObjectType int `json:"objectType"`
	// 背景颜色
	BgStyle string `json:"bgStyle"`
	// 字体颜色
	FontStyle string `json:"fontStyle"`
	// icon路径
	Icon string `json:"icon"`
	// 排序
	Sort int `json:"sort"`
	// 描述
	Remark string `json:"remark"`
	// 是否只读,1是 2否
	IsReadonly int `json:"isReadonly"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 项目对象类型列表响应结构体
type ProjectObjectTypeList struct {
	// 总数量
	Total int64 `json:"total"`
	// 列表
	List []*ProjectObjectType `json:"list"`
}

// 项目对象类型简单信息结构体
type ProjectObjectTypeRestInfo struct {
	// 主键
	ID int64 `json:"id"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 类型,1迭代，2问题
	ObjectType int `json:"objectType"`
	// 流程id
	ProcessID int64 `json:"processId"`
}

// 项目对象类型列表响应结构体
type ProjectObjectTypeWithProjectList struct {
	// 列表
	List []*ProjectObjectType `json:"list"`
}

type ProjectResourceInfoReq struct {
	// 资源id
	ResourceID int64 `json:"resourceId"`
	// 应用id
	AppID int64 `json:"appId"`
}

type ProjectResourceInfoResp struct {
	// 主键
	ID int64 `json:"id"`
	// 用户编号
	UserID int64 `json:"userId"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 路径
	Path string `json:"path"`
	// 文件名
	Name string `json:"name"`
	// 存储类型,1：本地，2：oss,3.钉盘
	Type int `json:"type"`
	// 文件后缀
	Suffix string `json:"suffix"`
	// 文件的md5
	Md5 string `json:"md5"`
	// 文件大小
	Size int64 `json:"size"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

type ProjectResourceReq struct {
	// 文件夹id
	FolderID int64 `json:"folderId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type ProjectStatisticsResp struct {
	// 迭代数量
	IterationTotal int64 `json:"iterationTotal"`
	// 任务数量
	TaskTotal int64 `json:"taskTotal"`
	// 成员数量
	MemberTotal int64 `json:"memberTotal"`
}

// 项目支持的对象类型请求结构体
type ProjectSupportObjectTypeListReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
}

// 项目支持的对象类型响应结构体
type ProjectSupportObjectTypeListResp struct {
	// 项目支持的对象类型
	ProjectSupportList []*ProjectObjectTypeRestInfo `json:"projectSupportList"`
	// 迭代支持的对象类型
	IterationSupportList []*ProjectObjectTypeRestInfo `json:"iterationSupportList"`
}

// 结构体
type ProjectType struct {
	// 主键
	ID int64 `json:"id"`
	// 组织编号
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 封面
	Cover string `json:"cover"`
	// 默认流程编号
	DefaultProcessID int64 `json:"defaultProcessId"`
	// 是否只读,2否,1是
	IsReadonly int `json:"isReadonly"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 项目栏目详情
	ProjectObjectTypeList []*ProjectObjectType `json:"projectObjectTypeList"`
}

// 项目类型分类结构体
type ProjectTypeCategory struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,全局的填0
	OrgID int64 `json:"orgId"`
	// 名称
	Name string `json:"name"`
	// icon
	Icon string `json:"icon"`
	// 类型,1迭代，2问题
	ObjectType int `json:"objectType"`
	// 排序
	Sort int `json:"sort"`
	// 描述
	Remark string `json:"remark"`
	// 是否只读,1是 2否
	IsReadonly int `json:"isReadonly"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 模板总数
	ProjectTypeTotal int64 `json:"projectTypeTotal"`
	// 模板列表
	ProjectTypeList []*ProjectType `json:"projectTypeList"`
}

type ProjectTypeCategoryList struct {
	// 总数量
	Total int64 `json:"total"`
	// 列表
	List []*ProjectTypeCategory `json:"list"`
}

// 列表响应结构体
type ProjectTypeList struct {
	Total int64          `json:"total"`
	List  []*ProjectType `json:"list"`
}

type ProjectTypeListReq struct {
	// 分类id
	CategoryID int64 `json:"categoryId"`
}

type ProjectTypeListResp struct {
	// 总数量
	Total int64 `json:"total"`
	// 列表
	List []*ProjectType `json:"list"`
}

type ProjectUserListReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type ProjectUserListResp struct {
	Total int64          `json:"total"`
	List  []*ProjectUser `json:"list"`
}

type ProjectsReq struct {
	// 主键
	ID *int64 `json:"id"`
	// 名称
	Name *string `json:"name"`
	// 项目负责人
	Owner *int64 `json:"owner"`
	// 项目类型
	ProjectTypeID *int64 `json:"projectTypeId"`
	// 项目优先级
	PriorityID *int64 `json:"priorityId"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 是否归档,1归档,2未归档,3全部（不传默认查未归档）
	IsFiling *int `json:"isFiling"`
	// 项目状态,从状态表取
	Status *int64 `json:"status"`
	// 项目状态，通用（1未开始2进行中3已完成4未完成）
	StatusType *int `json:"statusType"`
	// 创建人集合
	CreatorIds []int64 `json:"creatorIds"`
	// 负责人集合
	OwnerIds []int64 `json:"ownerIds"`
	// 关联类型(0所有1我发起的2我负责的3我参与的4我负责的和我参与的5我关注的)
	RelateType *int64 `json:"relateType"`
	// 参与人
	Participants []int64 `json:"participants"`
	// 参与部门
	ParticipantDeptIds []int64 `json:"participantDeptIds"`
	// 关注人
	Followers []int64 `json:"followers"`
	// 与我相关即我是成员（1是2否）
	IsMember *int `json:"isMember"`
	// 项目id集合
	ProjectIds []int64 `json:"projectIds"`
}

type QuitResult struct {
	// 是否退出
	IsQuitted interface{} `json:"isQuitted"`
}

// 阅读通知结构体
type ReadNoticeReq struct {
	// 主键
	ID int64 `json:"id"`
}

type RecoverRecycleBinRecordReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 对应资源的id
	RelationID int64 `json:"relationId"`
	// 类型1：任务2：标签3：文件夹4：文件5：附件
	RelationType int `json:"relationType"`
}

type RecycleBin struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 关联对象id
	RelationID int64 `json:"relationId"`
	// 类型1：任务2：标签3：文件夹4：文件5：附件
	RelationType int `json:"relationType"`
	// 名称
	Name string `json:"name"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 执行人信息
	CreatorInfo *UserIDInfo `json:"creatorInfo"`
	// 是否可操作
	IsCanDo bool `json:"isCanDo"`
	// 关联资源信息
	ResourceInfo *ResourceInfo `json:"resourceInfo"`
	// 标签信息
	TagInfo *Tag `json:"tagInfo"`
}

type RecycleBinList struct {
	Total int64         `json:"total"`
	List  []*RecycleBin `json:"list"`
}

type RecycleBinListReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 类型1：任务2：标签3：文件夹4：文件5：附件 不传或0为全部
	RelationType int `json:"relationType"`
}

// 官网注册联系人请求结构体
type RegisterWebSiteContactReq struct {
	// 性别：1 女，2 男
	Sex *int `json:"sex"`
	// 姓名
	Name *string `json:"name"`
	// 联系信息，可填手机号或者是邮箱（目前仅支持手机号）
	ContactInfo string `json:"contactInfo"`
	// 问题描述
	Remark *string `json:"remark"`
	// 问题图片url
	ResourceUrls []string `json:"resourceUrls"`
	// 来源
	Source *string `json:"source"`
}

// 关联任务列表请求结构体
type RelatedIssueListReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
}

type RelationIssue struct {
	// 关联信息id
	ID *int64 `json:"id"`
	// 关联信息名称
	Title *string `json:"title"`
}

type RelationType struct {
	// 用户id
	UserID *int64 `json:"userId"`
	// 类型id
	RelationType *int `json:"relationType"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 外部组织用户id
	OutOrgUserID *string `json:"outOrgUserId"`
	// 外部用户id
	OutUserID *string `json:"outUserId"`
	// 姓名
	Name *string `json:"name"`
	// 头像
	Avatar *string `json:"avatar"`
}

// 移除一个或多个邀请的成员
type RemoveInviteUserReq struct {
	// 要移除的id
	Ids []int64 `json:"ids"`
	// 是否删除全部(1是，其余为否)
	IsAll int `json:"isAll"`
}

// 移除组织
type RemoveOrgMemberReq struct {
	// 要移除的组织成员列表
	MemberIds []int64 `json:"memberIds"`
}

// 移出项目成员
type RemoveProjectMemberReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 要移除的项目成员列表
	MemberIds []int64 `json:"memberIds"`
	// 要移除的项目部门列表
	MemberForDepartmentID []int64 `json:"memberForDepartmentId"`
}

// 重新设置密码请求结构体
type ResetPasswordReq struct {
	// 当前密码
	CurrentPassword string `json:"currentPassword"`
	// 新密码
	NewPassword string `json:"newPassword"`
}

// 存放各类资源，其他业务表统一关联此表id结构体
type Resource struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// host
	Host string `json:"host"`
	// 路径
	Path string `json:"path"`
	// 缩略图路径
	PathCompressed string `json:"pathCompressed"`
	// 文件名
	Name string `json:"name"`
	// 存储类型,1：本地，2：oss,3.钉盘
	Type int `json:"type"`
	// 文件大小
	Size int64 `json:"size"`
	// 创建人姓名
	CreatorName string `json:"creatorName"`
	// 文件后缀
	Suffix string `json:"suffix"`
	// 文件的md5
	Md5 string `json:"md5"`
	// 文件类型
	FileType int `json:"fileType"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

type ResourceInfo struct {
	// 路径
	URL *string `json:"url"`
	// 资源名称
	Name *string `json:"name"`
	// 资源大小
	Size *int64 `json:"size"`
	// 上传时间
	UploadTime *time.Time `json:"uploadTime"`
	// 文件后缀
	Suffix *string `json:"suffix"`
	// 上传人
	Creator *int64 `json:"creator"`
	// 上传人名
	CreatorName *string `json:"creatorName"`
}

// 存放各类资源，其他业务表统一关联此表id列表响应结构体
type ResourceList struct {
	Total int64       `json:"total"`
	List  []*Resource `json:"list"`
}

// 找回密码请求结构体
type RetrievePasswordReq struct {
	// 账号，可以是邮箱或者手机号
	Username string `json:"username"`
	// 验证码
	AuthCode string `json:"authCode"`
	// 新密码
	NewPassword string `json:"newPassword"`
}

// 结构体
type Role struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,全局为0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 描述
	Remark string `json:"remark"`
	// 是否只读 1只读 2可编辑
	IsReadonly int `json:"isReadonly"`
	// 是否可以变更权限,1可以,2不可以
	IsModifyPermission int `json:"isModifyPermission"`
	// 是否默认角色,1是,2否
	IsDefault int `json:"isDefault"`
	// 角色分组
	RoleGroupID int64 `json:"roleGroupId"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
}

// 搜索用户接口入参
type SearchUserReq struct {
	// 邮箱
	Email string `json:"email"`
}

// 搜索用户返回值结构
type SearchUserResp struct {
	// 搜索到的用户的状态（1可邀请2已邀请3已添加/已注册）
	Status int `json:"status"`
	// 用户信息
	UserInfo *UserInfo `json:"userInfo"`
}

// 发送各种验证码请求结构体
type SendAuthCodeReq struct {
	// 验证方式: 1: 登录验证码，2：注册验证码，3：修改密码验证码，4：找回密码验证码，5：绑定验证码, 6：解绑验证码
	AuthType int `json:"authType"`
	// 地址类型: 1：手机号，2：邮箱
	AddressType int `json:"addressType"`
	// 联系地址，根据地址类型区分手机号或者邮箱
	Address string `json:"address"`
	// 验证码id
	CaptchaID *string `json:"captchaId"`
	// 输入的验证码
	CaptchaPassword *string `json:"captchaPassword"`
	// 易盾验证码
	YidunValidate *string `json:"yidunValidate"`
}

// 发送短信登录验证码请求结构体
type SendSmsLoginCodeReq struct {
	// 手机号
	PhoneNumber string `json:"phoneNumber"`
	// 验证码id
	CaptchaID *string `json:"captchaId"`
	// 输入的验证码
	CaptchaPassword *string `json:"captchaPassword"`
	// 易盾验证码
	YidunValidate *string `json:"yidunValidate"`
}

// 设置登录密码密码请求结构体
type SetPasswordReq struct {
	// 密码
	Password string `json:"password"`
}

type SetShareURLReq struct {
	// 分享的key
	Key string `json:"key"`
	// 分享的value
	URL string `json:"url"`
}

// 设置部门主管/管理员
type SetUserDepartmentLevelReq struct {
	// 用户id
	UserID int64 `json:"userId"`
	// 是否是部门主管(1是2否)
	IsLeader int `json:"isLeader"`
	// 部门id
	DepartmentID int64 `json:"departmentId"`
}

// 将用户变成任务成员请求参数
type SetUserJoinIssueReq struct {
	// 查询的任务id
	IssueID int64 `json:"issueId"`
	// 查询该用户是否是任务的成员。成员包括：参与人、负责人
	UserID int64 `json:"userId"`
}

type SimpleProjectInfo struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	ProjectTypeID int64  `json:"projectTypeId"`
}

type SimpleTagInfo struct {
	ID   *int64  `json:"id"`
	Name *string `json:"name"`
}

type StatCommon struct {
	// 名称
	Name string `json:"name"`
	// 数量
	Count int64 `json:"count"`
}

type StatData struct {
	IssueCount             int `json:"issueCount"`
	IssueWaitCount         int `json:"issueWaitCount"`
	IssueRunningCount      int `json:"issueRunningCount"`
	IssueEndCount          int `json:"issueEndCount"`
	StoryPointCount        int `json:"storyPointCount"`
	StoryPointWaitCount    int `json:"storyPointWaitCount"`
	StoryPointRunningCount int `json:"storyPointRunningCount"`
	StoryPointEndCount     int `json:"storyPointEndCount"`
}

type StatExtResp struct {
	Issue *StatIssueExt `json:"issue"`
}

type StatIssueExt struct {
	Data interface{} `json:"data"`
}

type StatusTimeInfo struct {
	StatusID      int64     `json:"statusId"`
	StatusName    string    `json:"statusName"`
	StatusType    int       `json:"statusType"`
	PlanStartTime time.Time `json:"planStartTime"`
	PlanEndTime   time.Time `json:"planEndTime"`
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
}

type StatusTimeInfoReq struct {
	StatusID      int64      `json:"statusId"`
	PlanStartTime *time.Time `json:"planStartTime"`
	PlanEndTime   *time.Time `json:"planEndTime"`
	StartTime     *time.Time `json:"startTime"`
	EndTime       *time.Time `json:"endTime"`
}

type StopThirdIntegrationReq struct {
	SourceChannel string `json:"sourceChannel"`
}

type StypeList struct {
	List []string `json:"list"`
}

type SwitchUserOrganizationReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
}

type SyncProToLcReq struct {
	// 同步类型：syncProView 创建默认项目视图；syncProject 同步项目到无码
	SyncType string `json:"syncType"`
	// 同步的组织列表
	OrgIds []int64 `json:"orgIds"`
	// OrgIds 为空时， StartPage 起作用。表示从多少页的偏移量查询组织开始跑脚本
	StartPage int `json:"startPage"`
}

// 从飞书方同步成员、部门等信息
type SyncUserInfoFromFeiShuReq struct {
	// 是否需要同步用户姓名
	NeedSyncName bool `json:"needSyncName"`
	// 是否需要同步用户头像
	NeedSyncAvatar bool `json:"needSyncAvatar"`
	// 是否需要同步部门架构信息
	NeedSyncDepartment bool `json:"needSyncDepartment"`
	// 用户来源、渠道。如果是飞书，则传 `fs`
	SourceChannel string `json:"sourceChannel"`
}

// 结构体
type Tag struct {
	// id
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 名称
	Name string `json:"name"`
	// 名称拼音
	NamePinyin string `json:"namePinyin"`
	// 背景色
	BgStyle string `json:"bgStyle"`
	// 字体色
	FontStyle string `json:"fontStyle"`
	// 使用任务数
	UsedNum int64 `json:"usedNum"`
	// 创建人id
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
}

// 列表响应结构体
type TagList struct {
	Total int64  `json:"total"`
	List  []*Tag `json:"list"`
}

// 查询请求
type TagListReq struct {
	// 名称
	Name *string `json:"name"`
	// 名称拼音
	NamePinyin *string `json:"namePinyin"`
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type ThirdAccountBindListResp struct {
	// 平台
	SourceChannel string `json:"sourceChannel"`
	// 姓名
	Name string `json:"name"`
	// 头像
	Avatar string `json:"avatar"`
}

type TransferOrgOwnerReq struct {
	OwnerID int64 `json:"ownerId"`
}

// 动态信息
type Trend struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 一级模块, 系统,组织,项目等
	Module1 string `json:"module1"`
	// 二级模块id
	Module2Id int64 `json:"module2Id"`
	// 二级模块.系统设置,消息设置,项目问题等
	Module2 string `json:"module2"`
	// 三级模块id
	Module3Id int64 `json:"module3Id"`
	// 三级模块,issus,迭代
	Module3 string `json:"module3"`
	// 操作编号
	OperCode string `json:"operCode"`
	// 被操作对象id
	OperObjID int64 `json:"operObjId"`
	// 被操作对象类型
	OperObjType string `json:"operObjType"`
	// 操作对象属性
	OperObjProperty string `json:"operObjProperty"`
	// 主关联对象id
	RelationObjID int64 `json:"relationObjId"`
	// 主关联对象类型
	RelationObjType string `json:"relationObjType"`
	// 关联类型
	RelationType string `json:"relationType"`
	// 新值,json
	NewValue *string `json:"newValue"`
	// 旧值,json
	OldValue *string `json:"oldValue"`
	// 扩展信息
	Ext string `json:"ext"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 被操作对象名称
	OperObjName string `json:"operObjName"`
	// 操作人名称
	CreatorInfo *UserIDInfo `json:"creatorInfo"`
	// 评论内容
	Comment *string `json:"comment"`
	// 扩展信息详情
	Extension *TrendExtension `json:"extension"`
	// 操作对象是否被删除
	ObjIsDelete bool `json:"objIsDelete"`
}

type TrendAuditInfo struct {
	// 确认装填（3通过4驳回）
	Status *int `json:"status"`
	// 确认内容
	Remark *string `json:"remark"`
	// 附件
	Attachments []*ResourceInfo `json:"attachments"`
}

type TrendExtension struct {
	IssueType *string `json:"issueType"`
	// 操作对象名称
	ObjName *string `json:"ObjName"`
	// 变更列表（主要用于更新字段）
	ChangeList []*ChangeList `json:"changeList"`
	// 涉及的变更成员信息（人员更新，关联对象增加/删除）
	MemberInfo []*MemberInfo `json:"memberInfo"`
	// 涉及的标签变更信息
	TagInfo []*SimpleTagInfo `json:"tagInfo"`
	// 关联问题信息
	RelationIssue *RelationIssue `json:"relationIssue"`
	// 关联资源信息
	ResourceInfo []*ResourceInfo `json:"resourceInfo"`
	// 通用变更数组
	CommonChange []*string `json:"commonChange"`
	// 文件夹id
	FolderID *int64 `json:"folderId"`
	// 字段id
	FieldIds []*int64 `json:"fieldIds"`
	// 项目对象类型id
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
	// 项目对象类型名称
	ProjectObjectTypeName *string `json:"projectObjectTypeName"`
	// 审批信息
	AuditInfo *TrendAuditInfo `json:"auditInfo"`
	// 增加的字段
	AddedFormFields []string `json:"addedFormFields"`
	// 删除的字段
	DeletedFormFields []string `json:"deletedFormFields"`
	// 更新的字段
	UpdatedFormFields []string `json:"updatedFormFields"`
}

// 动态列表请求结构体
type TrendReq struct {
	// 上次分页的最后一条动态id
	LastTrendID *int64 `json:"lastTrendId"`
	// 对象类型
	ObjType *string `json:"objType"`
	// 对象id
	ObjID *int64 `json:"objId"`
	// 操作id
	OperID *int64 `json:"operId"`
	// 开始时间
	StartTime *time.Time `json:"startTime"`
	// 结束时间
	EndTime *time.Time `json:"endTime"`
	// 分类（1任务动态2评论3项目动态（仅包括项目）4项目动态（包括项目和任务））5审批
	Type *int `json:"type"`
	// page
	Page *int64 `json:"page"`
	// size
	Size *int64 `json:"size"`
	// 排序（1时间正序2时间倒叙）
	OrderType *int `json:"orderType"`
}

// 动态列表
type TrendsList struct {
	// 总数量
	Total int64 `json:"total"`
	// 页码
	Page int64 `json:"page"`
	// size
	Size int64 `json:"size"`
	// 分页的最后一条动态id
	LastTrendID int64 `json:"lastTrendId"`
	// 列表
	List []*Trend `json:"list"`
}

// 解绑登录方式请求结构体（只剩下一种登录方式的时候不允许解绑）
type UnbindLoginNameReq struct {
	// 地址类型: 1：手机号，2：邮箱
	AddressType int `json:"addressType"`
	// 验证码
	AuthCode string `json:"authCode"`
}

type UnrelatedChatListReq struct {
	// 每页数量(不传默认10条)
	PageSize *int64 `json:"pageSize"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 上次分页最后一条外部群聊id（不传默认第一页）
	LastOutChatID *string `json:"lastOutChatId"`
	// 搜索内容
	Name *string `json:"name"`
}

// 更新接入应用信息请求结构体
type UpdateAppInfoReq struct {
	// 主键
	ID int64 `json:"id"`
	// 名称
	Name string `json:"name"`
	// 应用编号
	Code string `json:"code"`
	// 秘钥1
	Secret1 string `json:"secret1"`
	// 秘钥2
	Secret2 string `json:"secret2"`
	// 负责人
	Owner string `json:"owner"`
	// 审核状态,1待审核,2审核通过,3审核未通过
	CheckStatus int `json:"checkStatus"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

type UpdateCustomFieldReq struct {
	// 自定义字段id
	FieldID int64 `json:"fieldId"`
	// 名称
	Name *string `json:"name"`
	// 值
	FieldValue interface{} `json:"fieldValue"`
	// 是否加入组织字段库(1是2否)
	IsOrgField *int `json:"isOrgField"`
	// 字段描述
	Remark *string `json:"remark"`
	// 更新的字段
	UpdateField []string `json:"updateField"`
}

// 更新部门信息入参
type UpdateDepartmentForInviteReq struct {
	// 部门id
	DepartmentID *int64 `json:"DepartmentId"`
	// 部门名称（选填）
	Name *string `json:"Name"`
	// 部门主管(选填，不传表示不更新，传空数组则表示取消主管)
	LeaderIds []int64 `json:"LeaderIds"`
}

type UpdateDepartmentOrgRoleReq struct {
	// 部门id
	DepartmentID int64 `json:"departmentId"`
	// 修改后的角色id
	RoleID int64 `json:"roleId"`
	// 项目Id
	ProjectID *int64 `json:"projectId"`
}

// 更新部门请求结构体
type UpdateDepartmentReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 部门名称
	Name string `json:"name"`
	// 部门标识
	Code string `json:"code"`
	// 父部门id
	ParentID int64 `json:"parentId"`
	// 排序
	Sort int `json:"sort"`
	// 是否隐藏部门,1隐藏,2不隐藏
	IsHide int `json:"isHide"`
	// 来源渠道,
	SourceChannel string `json:"sourceChannel"`
	// 状态, 1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

type UpdateFsProjectChatPushSettingsReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 添加任务(1开2关)
	CreateIssue int `json:"createIssue"`
	// 任务负责人变更
	UpdateIssueOwner int `json:"updateIssueOwner"`
	// 任务状态变更
	UpdateIssueStatus int `json:"updateIssueStatus"`
	// 任务栏变更
	UpdateIssueProjectObjectType int `json:"updateIssueProjectObjectType"`
	// 任务标题被修改
	UpdateIssueTitle int `json:"updateIssueTitle"`
	// 任务时间被修改
	UpdateIssueTime int `json:"updateIssueTime"`
	// 任务有新的评论
	CreateIssueComment int `json:"createIssueComment"`
	// 任务有新的附件
	UploadNewAttachment int `json:"uploadNewAttachment"`
}

// 任务添加关联任务
type UpdateIssueAndIssueRelateReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 添加的关联任务id集合
	AddRelateIssueIds []int64 `json:"addRelateIssueIds"`
	// 删除的关联任务id集合
	DelRelateIssueIds []int64 `json:"delRelateIssueIds"`
}

type UpdateIssueBeforeAfterIssuesReq struct {
	// 前后置类型 1前置2后置
	Type int `json:"type"`
	// 任务id
	IssueID int64 `json:"issueId"`
	// 添加的关联任务id集合
	AddRelateIssueIds []int64 `json:"addRelateIssueIds"`
	// 删除的关联任务id集合
	DelRelateIssueIds []int64 `json:"delRelateIssueIds"`
}

type UpdateIssueCustionFieldData struct {
	// 字段id
	FieldID int64 `json:"fieldId"`
	// 字段值
	Value interface{} `json:"value"`
	// 名称
	Title string `json:"title"`
}

type UpdateIssueCustomFieldReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 字段参数
	Data []*UpdateIssueCustionFieldData `json:"data"`
}

// 更新问题对象类型请求结构体
type UpdateIssueObjectTypeReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 类型名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新项目对象类型请求结构体
type UpdateIssueProjectObjectTypeBatchReq struct {
	// 要更新的任务id
	Ids []int64 `json:"ids"`
	// 原来项目id
	OldProjectID int64 `json:"oldProjectId"`
	// 要更新的projectObjectType
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 状态id
	StatusID *int64 `json:"statusId"`
	// 迭代id
	IterationID *int64 `json:"iterationId"`
	// 镜像应用id
	MenuAppID *string `json:"menuAppId"`
}

// 批量移动任务响应结构体
type UpdateIssueProjectObjectTypeBatchResp struct {
	// 成功的id
	SuccessIssues []*Issue `json:"successIssues"`
	// 没有权限的任务id
	NoAuthIssues []*Issue `json:"noAuthIssues"`
	// 还有子任务没有选择的父任务id
	RemainChildrenIssues []*Issue `json:"remainChildrenIssues"`
	// 单纯是子任务的任务id
	ChildrenIssues []*Issue `json:"childrenIssues"`
}

// 更新项目对象类型请求结构体
type UpdateIssueProjectObjectTypeReq struct {
	// 要更新的任务id
	ID int64 `json:"id"`
	// 任务当前所属的项目 id
	FromProjectID *int64 `json:"fromProjectId"`
	// 任务当前所属的项目对象类型 id
	FromProjectObjectTypeID *int64 `json:"fromProjectObjectTypeId"`
	// 要更新的projectObjectType
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 状态id
	StatusID *int64 `json:"statusId"`
	// 迭代id
	IterationID *int64 `json:"iterationId"`
	// 是否携带子任务(默认不带)
	TakeChildren *bool `json:"takeChildren"`
}

// 更新问题性质请求结构体
type UpdateIssuePropertyReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 类型名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新任务请求结构体
type UpdateIssueReq struct {
	// 要更新的任务id
	ID int64 `json:"id"`
	// 标题
	Title *string `json:"title"`
	// 负责人
	OwnerID *int64 `json:"ownerId"`
	// 优先级id
	PriorityID *int64 `json:"priorityId"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 计划工作时长
	PlanWorkHour *int `json:"planWorkHour"`
	// 备注
	Remark *string `json:"remark"`
	// 备注详情
	RemarkDetail *string `json:"remarkDetail"`
	// 备注提及人
	MentionedUserIds []int64 `json:"mentionedUserIds"`
	// 迭代
	IterationID *int64 `json:"iterationId"`
	// 来源
	SourceID *int64 `json:"sourceId"`
	// 问题对象类型id
	IssueObjectTypeID *int64 `json:"issueObjectTypeId"`
	// 问题性质id
	IssuePropertyID *int64 `json:"issuePropertyId"`
	// 参与人
	ParticipantIds []int64 `json:"participantIds"`
	// 关注人
	FollowerIds []int64 `json:"followerIds"`
	// 关注人部门（后端实际转化为人）
	FollowerDeptIds []int64 `json:"followerDeptIds"`
	// 审核人
	AuditorIds []int64 `json:"auditorIds"`
	// 变动的字段列表
	UpdateFields []string `json:"updateFields"`
	// 无码更新入参
	LessUpdateIssueReq map[string]interface{} `json:"lessUpdateIssueReq"`
}

// 更新任务响应结构体
type UpdateIssueResp struct {
	// 任务id
	ID int64 `json:"id"`
}

// 更新任务Sort请求结构体
type UpdateIssueSortReq struct {
	// 任务id
	ID int64 `json:"id"`
	// 任务所属项目，从该项目中移动任务顺序。
	FromProjectID *int64 `json:"fromProjectId"`
	// 任务所属的项目类型
	FromProjectObjectTypeID *int64 `json:"fromProjectObjectTypeId"`
	// 排序位置标记，上一个任务id, beforeId和afterId至少传一个，否则不会更新sort
	BeforeID *int64 `json:"beforeId"`
	// 排序位置标记，下一个任务id
	AfterID *int64 `json:"afterId"`
	// 项目对象类型id
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
	// 状态id（主要针对于敏捷任务）
	StatusID *int64 `json:"statusId"`
	// 前置任务数据id
	BeforeDataID *string `json:"beforeDataId"`
	// 后置任务数据id
	AfterDataID *string `json:"afterDataId"`
	// 排序
	Asc *bool `json:"asc"`
}

// 更新问题来源请求结构体
type UpdateIssueSourceReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新任务状态请求结构体
type UpdateIssueStatusReq struct {
	// 任务id
	ID int64 `json:"id"`
	// 要更新的状态id
	NextStatusID *int64 `json:"nextStatusId"`
	// 要更新的状态类型,1: 未开始，2：进行中，3：已完成
	NextStatusType *int `json:"nextStatusType"`
	// 完成父任务时，是否同步更新子任务的状态,1:是，2：否，默认为否
	NeedModifyChildStatus *int `json:"needModifyChildStatus"`
}

// 更新任务标签关联请求结构体
type UpdateIssueTagsReq struct {
	// 任务id
	ID int64 `json:"id"`
	// 新关联的标签列表，addTags和delTags可以同时存在
	AddTags []*IssueTagReqInfo `json:"addTags"`
	// 要取消关联的标签列表
	DelTags []*IssueTagReqInfo `json:"delTags"`
}

type UpdateIssueViewReq struct {
	// 主键id，根据主键更新
	ID int64 `json:"id"`
	// 更新值：视图配置
	Config *string `json:"config"`
	// 更新值：视图备注
	Remark *string `json:"remark"`
	// 更新值：是否私有，true 私有，false 公开
	IsPrivate *bool `json:"isPrivate"`
	// 更新值：视图名称
	ViewName *string `json:"viewName"`
	// 更新值：类型，1：表格视图，2：看板视图，3：照片视图
	Type *int `json:"type"`
	// 视图排序
	Sort *int64 `json:"sort"`
	// 所属任务类型 id：需求、任务、缺陷的 id 值
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
}

// 编辑工时记录接口请求体
type UpdateIssueWorkHoursReq struct {
	// 工时记录id
	IssueWorkHoursID int64 `json:"issueWorkHoursId"`
	// 所需工时时间，单位：小时
	NeedTime string `json:"needTime"`
	// 工时执行者id
	WorkerID int64 `json:"workerId"`
	// 预留，剩余工时计算方式：1动态计算；2手动填写。没有则传 1
	RemainTimeCalType int64 `json:"remainTimeCalType"`
	// 预留，手动填写的剩余工时的值。没有则传 0
	RemainTime int64 `json:"remainTime"`
	// 工时的开始时间，**秒**级时间戳，没有则传 0
	StartTime int64 `json:"startTime"`
	// 工时的截止时间，**秒**级时间戳，没有则传 0
	EndTime int64 `json:"endTime"`
	// 工时记录的内容，工作内容
	Desc *string `json:"desc"`
}

// 更新迭代请求结构体
type UpdateIterationReq struct {
	// 主键
	ID int64 `json:"id"`
	// 名称
	Name *string `json:"name"`
	// 负责人
	Owner *int64 `json:"owner"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 变动的字段列表
	UpdateFields []string `json:"updateFields"`
}

type UpdateIterationSortReq struct {
	// 迭代id
	IterationID int64 `json:"iterationId"`
	// 移动位置的前一个迭代id,移到最前面传0
	BeforeID int64 `json:"beforeId"`
	// 后一个迭代，分页导致找不到前一个目标迭代
	AfterID *int64 `json:"afterId"`
}

// 更新迭代统计请求结构体
type UpdateIterationStatReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 迭代id
	IterationID int64 `json:"iterationId"`
	// 问题总数
	IssueCount int `json:"issueCount"`
	// 未开始问题数
	IssueWaitCount int `json:"issueWaitCount"`
	// 进行中问题数
	IssueRunningCount int `json:"issueRunningCount"`
	// 已完成问题数
	IssueEndCount int `json:"issueEndCount"`
	// 需求总数
	DemandCount int `json:"demandCount"`
	// 未开始需求数
	DemandWaitCount int `json:"demandWaitCount"`
	// 进行中需求数
	DemandRunningCount int `json:"demandRunningCount"`
	// 已完成需求数
	DemandEndCount int `json:"demandEndCount"`
	// 故事点总数
	StoryPointCount int `json:"storyPointCount"`
	// 未开始故事点数
	StoryPointWaitCount int `json:"storyPointWaitCount"`
	// 进行中故事点数
	StoryPointRunningCount int `json:"storyPointRunningCount"`
	// 已完成故事点数
	StoryPointEndCount int `json:"storyPointEndCount"`
	// 任务总数
	TaskCount int `json:"taskCount"`
	// 未开始任务数
	TaskWaitCount int `json:"taskWaitCount"`
	// 进行中任务数
	TaskRunningCount int `json:"taskRunningCount"`
	// 已完成任务数
	TaskEndCount int `json:"taskEndCount"`
	// 缺陷总数
	BugCount int `json:"bugCount"`
	// 未开始缺陷数
	BugWaitCount int `json:"bugWaitCount"`
	// 进行中缺陷数
	BugRunningCount int `json:"bugRunningCount"`
	// 已完成缺陷数
	BugEndCount int `json:"bugEndCount"`
	// 测试任务总数
	TesttaskCount int `json:"testtaskCount"`
	// 未开始测试任务数
	TesttaskWaitCount int `json:"testtaskWaitCount"`
	// 进行中测试任务数
	TesttaskRunningCount int `json:"testtaskRunningCount"`
	// 已完成测试任务数
	TesttaskEndCount int `json:"testtaskEndCount"`
	// 扩展
	Ext string `json:"ext"`
	// 统计日期
	StatDate time.Time `json:"statDate"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新迭代状态请求结构体
type UpdateIterationStatusReq struct {
	// 迭代id
	ID int64 `json:"id"`
	// 要更新的状态id
	NextStatusID int64 `json:"nextStatusId"`
	// 上个状态的实际结束时间
	BeforeStatusEndTime time.Time `json:"beforeStatusEndTime"`
	// 下个状态的实际开始时间
	NextStatusStartTime time.Time `json:"nextStatusStartTime"`
}

type UpdateIterationStatusTimeReq struct {
	IterationID  int64                `json:"iterationId"`
	StatusUpdate []*StatusTimeInfoReq `json:"statusUpdate"`
}

// 编辑：编辑详细预估工时
type UpdateMultiIssueWorkHoursReq struct {
	// 关联的任务id
	IssueID int64 `json:"issueId"`
	// 总预估工时记录
	TotalIssueWorkHourRecord *UpdateOneMultiWorkHourRecord `json:"totalIssueWorkHourRecord"`
	// 子预估工时列表
	IssueWorkHourRecords []*UpdateOneMultiWorkHourRecord `json:"issueWorkHourRecords"`
}

// 编辑：详细预估工时中的某个单独工时
type UpdateOneMultiWorkHourRecord struct {
	// 工时记录主键
	ID int64 `json:"id"`
	// 记录类型：1预估工时记录，2实际工时记录，3子预估工时
	Type int64 `json:"type"`
	// 工时记录的工作者、执行者id
	WorkerID int64 `json:"workerId"`
	// 所需工时时间，单位：小时
	NeedTime string `json:"needTime"`
	// 开始时间，秒级时间戳。
	StartTime int64 `json:"startTime"`
	// 截止时间，秒级时间戳。
	EndTime int64 `json:"endTime"`
	// 工时记录的内容，工作内容
	Desc *string `json:"desc"`
}

type UpdateOrgBasicShowSettingReq struct {
	// 工作台
	WorkBenchShow bool `json:"workBenchShow"`
	// 侧边栏
	SideBarShow bool `json:"sideBarShow"`
	// 镜像统计
	MirrorStat bool `json:"mirrorStat"`
}

// 修改组织成员审核状态请求结构体
type UpdateOrgMemberCheckStatusReq struct {
	// 要修改的组织成员列表
	MemberIds []int64 `json:"memberIds"`
	// 审核状态, 1待审核,2审核通过,3审核不过
	CheckStatus int `json:"checkStatus"`
}

// 修改组织成员状态请求结构体
type UpdateOrgMemberStatusReq struct {
	// 要修改的组织成员列表
	MemberIds []int64 `json:"memberIds"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
}

// 组织设置入参
type UpdateOrganizationSettingsReq struct {
	// 组织id
	OrgID int64 `json:"orgId"`
	// 组织名称
	OrgName string `json:"orgName"`
	// 组织code
	Code *string `json:"code"`
	// 所属行业
	IndustryID *int64 `json:"industryId"`
	// 组织规模
	Scale *string `json:"scale"`
	// 所在国家
	CountryID *int64 `json:"countryId"`
	// 所在省份
	ProvinceID *int64 `json:"provinceId"`
	// 所在城市
	CityID *int64 `json:"cityId"`
	// 组织地址
	Address *string `json:"address"`
	// 组织logo地址
	LogoURL *string `json:"logoUrl"`
	// 组织负责人
	Owner *int64 `json:"owner"`
	// 变动的字段列表
	UpdateFields []string `json:"updateFields"`
}

// 更新请求结构体
type UpdatePermissionOperationReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 权限项id
	PermissionID int64 `json:"permissionId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 操作编号,多个半角逗号分隔
	OperationCodes string `json:"operationCodes"`
	// 描述
	Remark string `json:"remark"`
	// 是否显示,1是,2否
	IsShow int `json:"isShow"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新请求结构体
type UpdatePermissionReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 编号,仅支持大写英文字母
	Code string `json:"code"`
	// 名称
	Name string `json:"name"`
	// 父id
	ParentID int64 `json:"parentId"`
	// 权限项类型,1系统,2组织,3项目
	Type int `json:"type"`
	// 权限路径
	Path string `json:"path"`
	// 是否显示,1是,2否
	IsShow int `json:"isShow"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新优先级请求结构体
type UpdatePriorityReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,全局的填0
	OrgID int64 `json:"orgId"`
	// 语言编号
	LangCode string `json:"langCode"`
	// 名称
	Name string `json:"name"`
	// 类型,1项目优先级,2:需求/任务等优先级
	Type int `json:"type"`
	// 排序
	Sort int `json:"sort"`
	// 背景颜色
	BgStyle string `json:"bgStyle"`
	// 字体颜色
	FontStyle string `json:"fontStyle"`
	// 是否默认,1是,2否
	IsDefault int `json:"isDefault"`
	// 描述
	Remark string `json:"remark"`
	// 状态,  1可用,2禁用
	Status int `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新流程状态请求结构体
type UpdateProcessStatusReq struct {
	// 主键
	ID int64 `json:"id"`
	// 流程id
	ProcessID int64 `json:"processId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 背景颜色
	BgStyle *string `json:"bgStyle"`
	// 字体颜色
	FontStyle *string `json:"fontStyle"`
	// 名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 后一系列对象类型的id
	AfterID *int64 `json:"afterId"`
	// 往前挪的一些列对象类型的id
	BeforeID *int64 `json:"beforeId"`
	// 状态类型,1未开始,2进行中,3已完成
	Type int `json:"type"`
}

// 更新项目日统计请求结构体
type UpdateProjectDayStatReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id,应该是全局的,因此填0
	OrgID int64 `json:"orgId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 问题总数
	IssueCount int `json:"issueCount"`
	// 未开始问题数
	IssueWaitCount int `json:"issueWaitCount"`
	// 进行中问题数
	IssueRunningCount int `json:"issueRunningCount"`
	// 已完成问题数
	IssueEndCount int `json:"issueEndCount"`
	// 需求总数
	DemandCount int `json:"demandCount"`
	// 未开始需求数
	DemandWaitCount int `json:"demandWaitCount"`
	// 进行中需求数
	DemandRunningCount int `json:"demandRunningCount"`
	// 已完成需求数
	DemandEndCount int `json:"demandEndCount"`
	// 故事点总数
	StoryPointCount int `json:"storyPointCount"`
	// 未开始故事点数
	StoryPointWaitCount int `json:"storyPointWaitCount"`
	// 进行中故事点数
	StoryPointRunningCount int `json:"storyPointRunningCount"`
	// 已完成故事点数
	StoryPointEndCount int `json:"storyPointEndCount"`
	// 任务总数
	TaskCount int `json:"taskCount"`
	// 未开始任务数
	TaskWaitCount int `json:"taskWaitCount"`
	// 进行中任务数
	TaskRunningCount int `json:"taskRunningCount"`
	// 已完成任务数
	TaskEndCount int `json:"taskEndCount"`
	// 缺陷总数
	BugCount int `json:"bugCount"`
	// 未开始缺陷数
	BugWaitCount int `json:"bugWaitCount"`
	// 进行中缺陷数
	BugRunningCount int `json:"bugRunningCount"`
	// 已完成缺陷数
	BugEndCount int `json:"bugEndCount"`
	// 测试任务总数
	TesttaskCount int `json:"testtaskCount"`
	// 未开始测试任务数
	TesttaskWaitCount int `json:"testtaskWaitCount"`
	// 进行中测试任务数
	TesttaskRunningCount int `json:"testtaskRunningCount"`
	// 已完成测试任务数
	TesttaskEndCount int `json:"testtaskEndCount"`
	// 扩展
	Ext string `json:"ext"`
	// 统计日期
	StatDate time.Time `json:"statDate"`
	// 项目状态,从状态表取
	Status int64 `json:"status"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

type UpdateProjectDetailReq struct {
	// 详情id
	ID int64 `json:"id"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 公告
	Notice *string `json:"notice"`
	// 是否启用工时和故事点,1启用,2不启用
	IsEnableWorkHours *int `json:"isEnableWorkHours"`
	// 是否同步到飞书日历(1是2否,默认否)
	IsSyncOutCalendar *int `json:"isSyncOutCalendar"`
}

type UpdateProjectFieldViewReq struct {
	// 视图类型（1看板2列表3表格）
	ViewType int `json:"viewType"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 任务类型
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 关闭的默认字段
	ClosedDefaultFields []string `json:"closedDefaultFields"`
	// 关闭的自定义字段
	ClosedCustomFields []int64 `json:"closedCustomFields"`
}

// 更新文件夹请求结构体
type UpdateProjectFolderReq struct {
	// 文件夹id
	FolderID int64 `json:"folderId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 文件夹名
	Name *string `json:"name"`
	// 父级文件夹id
	ParentID *int64 `json:"parentId"`
	// 文件夹类型,0其他,1文档,2图片,3视频,4音频
	FileType *int64 `json:"fileType"`
	// 变动的字段列表
	UpdateFields []string `json:"updateFields"`
}

type UpdateProjectFolderResp struct {
	// 文件夹id
	FolderID int64 `json:"folderId"`
}

// 更新项目对象类型请求结构体
type UpdateProjectObjectTypeReq struct {
	// 主键
	ID int64 `json:"id"`
	// 项目iid,用来校验权限
	ProjectID int64 `json:"projectId"`
	// 名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
	// 后一系列对象类型的id
	AfterID *int64 `json:"afterId"`
	// 往前挪的一些列对象类型的id
	BeforeID *int64 `json:"beforeId"`
}

type UpdateProjectReq struct {
	// 项目id
	ID int64 `json:"id"`
	// 编号
	Code *string `json:"code"`
	// 名称
	Name *string `json:"name"`
	// 前缀编号
	PreCode *string `json:"preCode"`
	// 负责人id
	Owner *int64 `json:"owner"`
	// 负责人id集合
	OwnerIds []int64 `json:"ownerIds"`
	// 优先级
	PriorityID *int64 `json:"priorityId"`
	// 计划开始时间
	PlanStartTime *time.Time `json:"planStartTime"`
	// 计划结束时间
	PlanEndTime *time.Time `json:"planEndTime"`
	// 项目公开性,1公开,2私有
	PublicStatus *int `json:"publicStatus"`
	// 资源id
	ResourceID *int64 `json:"resourceId"`
	// 描述
	Remark *string `json:"remark"`
	// 项目状态
	Status *int64 `json:"status"`
	// 资源路径
	ResourcePath *string `json:"resourcePath"`
	// 资源类型1本地2oss3钉盘
	ResourceType *int `json:"resourceType"`
	// 用户成员id
	MemberIds []int64 `json:"memberIds"`
	// 用户成员部门id
	MemberForDepartmentID []int64 `json:"memberForDepartmentId"`
	// 是否全选（针对于项目成员）
	IsAllMember *bool `json:"isAllMember"`
	// 关注人id
	FollowerIds []int64 `json:"followerIds"`
	// 针对哪些群体用户，同步到其飞书日历(4：同步给负责人，8：同步给关注人。16:同步到订阅日历。往后扩展是基于二进制的位值)。该值是所有状态的算术总和。
	IsSyncOutCalendar *int `json:"isSyncOutCalendar"`
	// 变动的字段列表
	UpdateFields []string `json:"updateFields"`
	// 针对哪些群体用户，同步到其飞书日历(4：同步给负责人，8：同步给关注人。16:同步到订阅日历。往后扩展是基于二进制的位值)
	SyncCalendarStatusList []*int `json:"syncCalendarStatusList"`
	// 是否创建群聊（针对于飞书1是2否默认是）
	IsCreateFsChat *int `json:"isCreateFsChat"`
	// 隐私模式状态。1开启；2不开启；默认2。
	PrivacyStatus *int `json:"privacyStatus"`
}

type UpdateProjectResourceFolderReq struct {
	// 当前文件夹id
	CurrentFolderID int64 `json:"currentFolderId"`
	// 目标文件夹id
	TargetFolderID int64 `json:"targetFolderId"`
	// 文件id数组
	ResourceIds []int64 `json:"resourceIds"`
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type UpdateProjectResourceFolderResp struct {
	// 文件id数组
	ResourceIds []int64 `json:"resourceIds"`
}

type UpdateProjectResourceNameReq struct {
	// 文件id
	ResourceID int64 `json:"resourceId"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 文件名
	FileName *string `json:"fileName"`
	// 文件后缀
	FileSuffix *string `json:"fileSuffix"`
	// 文件大小
	FileSize int64 `json:"fileSize"`
	// 文件Md5
	FileMd5 *string `json:"fileMd5"`
	// 修改者UserId
	UpdaterID int64 `json:"updaterId"`
	// 修改项
	UpdateFields []string `json:"updateFields"`
}

type UpdateProjectStatusReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 下一个状态
	NextStatusID int64 `json:"nextStatusId"`
}

type UpdateRelateChat struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 群聊外部id
	OutChatIds []string `json:"outChatIds"`
}

// 更新存放各类资源，其他业务表统一关联此表id请求结构体
type UpdateResourceReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 路径
	Path string `json:"path"`
	// 文件名
	Name string `json:"name"`
	// 存储类型,1：本地，2：oss,3.钉盘
	Type int `json:"type"`
	// 文件后缀
	Suffix string `json:"suffix"`
	// 文件的md5
	Md5 string `json:"md5"`
	// 文件大小
	Size int64 `json:"size"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新角色操作权限
type UpdateRolePermissionOperationReq struct {
	// 角色id
	RoleID int64 `json:"roleId"`
	// 项目id
	ProjectID *int64 `json:"projectId"`
	// 所有涉及到的更改的权限组
	UpdatePermissions []*EveryPermission `json:"updatePermissions"`
}

// 更新角色请求结构体
type UpdateRoleReq struct {
	// 角色ID
	RoleID int64 `json:"roleId"`
	// 角色名称, 非必填，为空则不更新
	Name *string `json:"name"`
}

// 更新请求结构体
type UpdateTagReq struct {
	ID int64 `json:"id"`
	// 名称
	Name *string `json:"name"`
	// 背景颜色
	BgStyle *string `json:"bgStyle"`
}

// 更新用户配置请求结构体
type UpdateUserConfigReq struct {
	// 主键
	ID int64 `json:"id"`
	// 个人日报开启状态, 2否, 1是
	DailyReportMessageStatus int `json:"dailyReportMessageStatus"`
	// 我负责的 2否, 1是
	OwnerRangeStatus int `json:"ownerRangeStatus"`
	// 我参与的, 2否, 1是
	ParticipantRangeStatus int `json:"participantRangeStatus"`
	// 我关注的, 2否, 1是
	AttentionRangeStatus int `json:"attentionRangeStatus"`
	// 我创建的, 2否, 1是
	CreateRangeStatus int `json:"createRangeStatus"`
	// 任务提醒状态 2否, 1是
	RemindMessageStatus int `json:"remindMessageStatus"`
	// 评论和at我的通知
	CommentAtMessageStatus int `json:"commentAtMessageStatus"`
	// 任务更新状态, 2否, 1是
	ModifyMessageStatus int `json:"modifyMessageStatus"`
	// 任务关联动态, 2否, 1是
	RelationMessageStatus int `json:"relationMessageStatus"`
	// 项目日报开启状态, 2否, 1是
	DailyProjectReportMessageStatus int `json:"dailyProjectReportMessageStatus"`
}

// 更新用户配置响应结构体
type UpdateUserConfigResp struct {
	// 主键
	ID int64 `json:"id"`
}

// 更新用户默认项目配置请求结构体
type UpdateUserDefaultProjectConfigReq struct {
	// 默认项目id, 机器人创建项目的时候会选用这个项目
	DefaultProjectID int64 `json:"defaultProjectId"`
	// 默认工作栏id
	DefaultProjectObjectTypeID *int64 `json:"defaultProjectObjectTypeId"`
}

// 更改用户个人信息
type UpdateUserInfoReq struct {
	// 姓名
	Name *string `json:"name"`
	// 性别
	Sex *int `json:"sex"`
	// 用户头像
	Avatar *string `json:"avatar"`
	// 生日
	Birthday *time.Time `json:"birthday"`
	// 是否需要提醒绑定手机号
	RemindBindPhone *int `json:"remindBindPhone"`
	// 变动的字段列表
	UpdateFields []string `json:"updateFields"`
}

type UpdateUserOrgRoleBatchReq struct {
	// 用户id
	UserIds []int64 `json:"userIds"`
	// 修改后的角色id
	RoleID int64 `json:"roleId"`
}

type UpdateUserOrgRoleReq struct {
	// 用户id
	UserID int64 `json:"userId"`
	// 修改后的角色id
	RoleID int64 `json:"roleId"`
	// 项目Id
	ProjectID *int64 `json:"projectId"`
}

// 更新请求结构体
type UpdateUserOrganizationReq struct {
	// 主键
	ID int64 `json:"id"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 用户id
	UserID int64 `json:"userId"`
	// 审核状态,1待审核,2审核通过,3审核不过
	CheckStatus int `json:"checkStatus"`
	// 使用状态,1已使用,2未使用
	UseStatus int `json:"useStatus"`
	// 企业用户状态, 1可用,2禁用
	Status int `json:"status"`
	// 状态变更人id
	StatusChangerID int64 `json:"statusChangerId"`
	// 状态变更时间
	StatusChangeTime time.Time `json:"statusChangeTime"`
	// 审核人id
	AuditorID int64 `json:"auditorId"`
	// 审核时间
	AuditTime time.Time `json:"auditTime"`
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 更新人
	Updator int64 `json:"updator"`
	// 更新时间
	UpdateTime time.Time `json:"updateTime"`
	// 乐观锁
	Version int `json:"version"`
	// 是否删除,1是,2否
	IsDelete int `json:"isDelete"`
}

// 更新用户pc配置请求结构体
type UpdateUserPcConfigReq struct {
	// pc桌面通知开关状态, 2否, 1是
	PcNoticeOpenStatus *int `json:"pcNoticeOpenStatus"`
	// pc任务提醒状态, 2否, 1是
	PcIssueRemindMessageStatus *int `json:"pcIssueRemindMessageStatus"`
	// pc组织相关推送状态, 2否, 1是
	PcOrgMessageStatus *int `json:"pcOrgMessageStatus"`
	// pc项目相关推送状态, 2否, 1是
	PcProjectMessageStatus *int `json:"pcProjectMessageStatus"`
	// pc评论相关推送状态, 2否, 1是
	PcCommentAtMessageStatus *int `json:"pcCommentAtMessageStatus"`
	// 变动的字段列表
	UpdateFields []string `json:"updateFields"`
}

// 编辑/更新成员信息
type UpdateUserReq struct {
	// 用户id(必填)
	UserID int64 `json:"userId"`
	// 手机号(选填)
	PhoneNumber *string `json:"phoneNumber"`
	// 邮箱(选填)
	Email *string `json:"email"`
	// 姓名(选填)
	Name *string `json:"name"`
	// 部门id（选填）
	DepartmentIds []int64 `json:"departmentIds"`
	// 角色id（选填）
	RoleIds []int64 `json:"roleIds"`
	// 状态（1启用2禁用 选填）
	Status *int `json:"status"`
}

type UploadOssByFsImageKeyReq struct {
	IamgeKey string `json:"iamgeKey"`
	IsApp    bool   `json:"isApp"`
}

type UploadOssByFsImageKeyResp struct {
	URL string `json:"url"`
}

type UrgeAuditIssueReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 是否在任务群聊中 at 负责人
	IsNeedAtIssueOwner *bool `json:"isNeedAtIssueOwner"`
	// 催促内容
	UrgeText string `json:"urgeText"`
}

type UrgeIssueReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
	// 是否在任务群聊中 at 负责人
	IsNeedAtIssueOwner bool `json:"isNeedAtIssueOwner"`
	// 催促内容。（可选）
	UrgeText *string `json:"urgeText"`
}

type UseOrgCustomFieldReq struct {
	// 自定义字段id集合
	FieldIds []int64 `json:"fieldIds"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 任务类型
	ProjectObjectTypeID *int64 `json:"projectObjectTypeId"`
}

// 用户配置信息结构体
type UserConfig struct {
	// 主键
	ID int64 `json:"id"`
	// 个人日报开启状态, 2否, 1是
	DailyReportMessageStatus int `json:"dailyReportMessageStatus"`
	// 我负责的 2否, 1是
	OwnerRangeStatus int `json:"ownerRangeStatus"`
	// 我参与的, 2否, 1是
	ParticipantRangeStatus int `json:"participantRangeStatus"`
	// 我关注的, 2否, 1是
	AttentionRangeStatus int `json:"attentionRangeStatus"`
	// 我创建的, 2否, 1是
	CreateRangeStatus int `json:"createRangeStatus"`
	// 任务提醒状态 2否, 1是
	RemindMessageStatus int `json:"remindMessageStatus"`
	// 评论和at我的通知
	CommentAtMessageStatus int `json:"commentAtMessageStatus"`
	// 任务更新状态, 2否, 1是
	ModifyMessageStatus int `json:"modifyMessageStatus"`
	// 任务关联动态, 2否, 1是
	RelationMessageStatus int `json:"relationMessageStatus"`
	// 项目日报开启状态, 2  否 1:是
	DailyProjectReportMessageStatus int `json:"dailyProjectReportMessageStatus"`
	// 默认项目id, 机器人创建项目的时候会选用这个项目
	DefaultProjectID int64 `json:"defaultProjectId"`
	// 默认工作栏
	DefaultProjectObjectTypeID int64 `json:"defaultProjectObjectTypeId"`
	// pc桌面通知开关状态, 2否, 1是
	PcNoticeOpenStatus int `json:"pcNoticeOpenStatus"`
	// pc任务提醒状态, 2否, 1是
	PcIssueRemindMessageStatus int `json:"pcIssueRemindMessageStatus"`
	// pc组织相关推送状态, 2否, 1是
	PcOrgMessageStatus int `json:"pcOrgMessageStatus"`
	// pc项目相关推送状态, 2否, 1是
	PcProjectMessageStatus int `json:"pcProjectMessageStatus"`
	// pc评论相关推送状态, 2否, 1是
	PcCommentAtMessageStatus int `json:"pcCommentAtMessageStatus"`
}

// 部门信息
type UserDepartmentData struct {
	// 部门id
	DepartmentID *int64 `json:"departmentId"`
	// 是否是主管：1是2否
	IsLeader *int `json:"isLeader"`
	// 部门名称
	DeparmentName *string `json:"deparmentName"`
}

// 用户id信息
type UserIDInfo struct {
	// 用户id
	ID int64 `json:"id"`
	// 用户id
	UserID int64 `json:"userId"`
	// 用户名称
	Name string `json:"name"`
	// 用户拼音
	NamePy string `json:"namePy"`
	// 用户头像
	Avatar string `json:"avatar"`
	// 工号：企业下唯一
	EmplID string `json:"emplId"`
	// unionId： 开发者账号下唯一(暂时废弃，返回空)
	UnionID string `json:"unionId"`
	// 是否已被删除，为true则代表被组织移除
	IsDeleted bool `json:"isDeleted"`
	// 是否已被禁用, 为true则代表被组织禁用
	IsDisabled bool `json:"isDisabled"`
}

// 用户id信息
type UserIDInfoExtraForIssueAudit struct {
	// 用户id
	ID int64 `json:"id"`
	// 用户id
	UserID int64 `json:"userId"`
	// 用户名称
	Name string `json:"name"`
	// 用户拼音
	NamePy string `json:"namePy"`
	// 用户头像
	Avatar string `json:"avatar"`
	// 工号：企业下唯一
	EmplID string `json:"emplId"`
	// unionId： 开发者账号下唯一(暂时废弃，返回空)
	UnionID string `json:"unionId"`
	// 是否已被删除，为true则代表被组织移除
	IsDeleted bool `json:"isDeleted"`
	// 是否已被禁用, 为true则代表被组织禁用
	IsDisabled bool `json:"isDisabled"`
	// 状态(1未查看2已查看未审核3审核通过4驳回)
	AuditStatus int `json:"auditStatus"`
}

// 成员信息结构体
type UserInfo struct {
	// 成员 id
	UserID int64 `json:"userID"`
	// 姓名
	Name string `json:"name"`
	// 姓名拼音
	NamePy string `json:"namePy"`
	// 用户头像
	Avatar string `json:"avatar"`
	// 邮箱
	Email string `json:"email"`
	// 手机
	PhoneNumber string `json:"phoneNumber"`
	// 用户部门信息
	DepartmentList []*UserDepartmentData `json:"departmentList"`
	// 角色信息
	RoleList []*UserRoleData `json:"roleList"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 状态：1启用，2禁用
	Status int `json:"status"`
	// 禁用时间
	StatusChangeTime time.Time `json:"statusChangeTime"`
	// 是否是组织创建人
	IsCreator bool `json:"isCreator"`
}

// 用户登录请求结构体
type UserLoginReq struct {
	// 登录类型: 1、短信验证码登录，2、账号密码登录，3、邮箱验证码登录
	LoginType int `json:"loginType"`
	// 登录类型为1时，loginName为手机号； 登录类型为3时，loginName为邮箱
	LoginName string `json:"loginName"`
	// 登录类型为2时，密码必传
	Password *string `json:"password"`
	// 验证码
	AuthCode *string `json:"authCode"`
	// 注册时可以带上名字
	Name *string `json:"name"`
	// 邀请码, 邀请注册时必填
	InviteCode *string `json:"inviteCode"`
	// 来源通道
	SourceChannel string `json:"sourceChannel"`
	// 平台
	SourcePlatform string `json:"sourcePlatform"`
	// codeToken
	CodeToken *string `json:"codeToken"`
}

// 用户登录响应结构体
type UserLoginResp struct {
	// 用户token
	Token string `json:"token"`
	// 用户id
	UserID int64 `json:"userId"`
	// 组织id
	OrgID int64 `json:"orgId"`
	// 组织名称
	OrgName string `json:"orgName"`
	// 组织code
	OrgCode string `json:"orgCode"`
	// 用户名称
	Name string `json:"name"`
	// 头像
	Avatar string `json:"avatar"`
	// 是否需要创建组织
	NeedInitOrg bool `json:"needInitOrg"`
	// 是否不是飞书手机号
	NotFsMobile bool `json:"notFsMobile"`
}

// 用户组织列表响应结构体
type UserOrganization struct {
	// 组织id
	ID int64 `json:"id"`
	// 组织名称
	Name string `json:"name"`
	// 组织code
	Code string `json:"code"`
	// 组织网站
	WebSite string `json:"webSite"`
	// 所属行业
	IndustryID int64 `json:"industryId"`
	// 组织规模
	Scale string `json:"scale"`
	// 来源平台
	SourcePlatform string `json:"sourcePlatform"`
	// 来源渠道
	SourceChannel string `json:"sourceChannel"`
	// 所在国家
	CountryID int64 `json:"countryId"`
	// 所在省份
	ProvinceID int64 `json:"provinceId"`
	// 所在城市
	CityID int64 `json:"cityId"`
	// 组织地址
	Address string `json:"address"`
	// 组织logo地址
	LogoURL string `json:"logoUrl"`
	// 组织标识
	ResorceID int64 `json:"resorceId"`
	// 组织所有人,创建时默认为创建人
	Owner int64 `json:"owner"`
	// 企业是否认证
	IsAuthenticated int `json:"IsAuthenticated"`
	// 是否为企业管理员
	IsAdmin bool `json:"isAdmin"`
	// 描述
	Remark string `json:"remark"`
	// 是否展示
	IsShow int `json:"isShow"`
	// 是否删除,1是,2否
	IsDelete *int `json:"isDelete"`
	// 对于该用户组织是否可用（1是2否）
	OrgIsEnabled *int `json:"OrgIsEnabled"`
	// 组织可用功能
	Functions []string `json:"functions"`
}

// 列表响应结构体
type UserOrganizationList struct {
	Total int64               `json:"total"`
	List  []*OrganizationUser `json:"list"`
}

type UserOrganizationListResp struct {
	// 用户组织列表
	List []*UserOrganization `json:"list"`
}

// 用户注册请求结构体
type UserRegisterReq struct {
	// 注册用户名（邮箱，手机号，账号等等）
	UserName string `json:"userName"`
	// 注册类型(1,手机号，2，账号，3，邮箱)(暂时只支持邮箱/手机)
	RegisterType int `json:"registerType"`
	// 姓名
	Name *string `json:"name"`
	// 密码，只有注册类型为2时必填
	Password *string `json:"password"`
	// 短信或者邮箱验证码，当注册类型为1和3时必填
	AuthCode *string `json:"authCode"`
	// 来源通道
	SourceChannel string `json:"sourceChannel"`
	// 平台
	SourcePlatform string `json:"sourcePlatform"`
}

// 用户注册响应结构体
type UserRegisterResp struct {
	// 用户token
	Token string `json:"token"`
}

// 角色信息
type UserRoleData struct {
	// 角色id
	RoleID *int64 `json:"RoleId"`
	// 角色名称
	RoleName *string `json:"RoleName"`
}

type UserRoleInfo struct {
	// 角色id
	ID int64 `json:"id"`
	// 角色名称
	Name string `json:"name"`
	// 角色lang_code
	LangCode string `json:"langCode"`
}

// 当前所在的部门成员状态信息
type UserStatResp struct {
	// 所有成员数量
	AllCount int64 `json:"allCount"`
	// 未分配成员数量
	UnallocatedCount int64 `json:"unallocatedCount"`
	// 未接受邀请成员数量
	UnreceivedCount int64 `json:"unreceivedCount"`
	// 已禁用成员数量
	ForbiddenCount int64 `json:"forbiddenCount"`
}

type ViewAuditIssueReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
}

type WithdrawIssueReq struct {
	// 任务id
	IssueID int64 `json:"issueId"`
}

// 工时记录的工作者、执行者
type WorkHourWorker struct {
	// 执行人的id
	UserID int64 `json:"userId"`
	// 执行人的名字
	Name string `json:"name"`
	// 执行人的头像
	Avatar string `json:"avatar"`
}

type FsDocumentData struct {
	// 标题
	Title string `json:"title"`
	// 文档类型
	DocsType string `json:"docsType"`
	// token
	DocsToken string `json:"docsToken"`
	// url
	DocsURL string `json:"docsUrl"`
	// 创建人
	OwnerName string `json:"ownerName"`
	// 创建人id
	OwnerID string `json:"ownerId"`
}

type FsDocumentListReq struct {
	// 关键字
	SearchKey *string `json:"searchKey"`
}

type FsDocumentListResp struct {
	Total int64             `json:"total"`
	List  []*FsDocumentData `json:"list"`
}

type GetIssueViewListResp struct {
	// 总数量
	Total int64 `json:"total"`
	// 查询到的视图列表
	List []*GetIssueViewListItem `json:"list"`
}

type IssueListStatData struct {
	// 对象类型id
	ProjectObjectTypeID int64 `json:"projectObjectTypeId"`
	// 对象类型名称
	ProjectObjectTypeName string `json:"projectObjectTypeName"`
	// 数量
	Total int64 `json:"total"`
	// 已完成数量
	FinishedCount int64 `json:"finishedCount"`
	// 逾期数量
	OverdueCount int64 `json:"overdueCount"`
}

type IssueListStatReq struct {
	// 项目id
	ProjectID int64 `json:"projectId"`
}

type IssueListStatResp struct {
	List []*IssueListStatData `json:"list"`
}

type OrgProjectMemberListReq struct {
	// 关联类型(1负责人2关注人3全部，默认全部)
	RelationType *int64 `json:"relationType"`
	// 名字
	Name *string `json:"name"`
	// 项目id
	ProjectID int64 `json:"projectId"`
	// 是否忽略离职人员(默认否)
	IgnoreDelete *bool `json:"ignoreDelete"`
}

type ProjectObjectTypesReq struct {
	// 类型,1迭代，2问题
	ObjectType int `json:"objectType"`
	// 任务栏id集合
	Ids []int64 `json:"Ids"`
}

// 项目类型入参
type ProjectTypesReq struct {
	// 主键
	ID *int64 `json:"id"`
	// 组织编号
	OrgID *int64 `json:"orgId"`
	// 语言编号
	LangCode *string `json:"langCode"`
	// 名称
	Name *string `json:"name"`
}

type ProjectUser struct {
	// 创建人
	Creator int64 `json:"creator"`
	// 创建时间
	CreateTime time.Time `json:"createTime"`
	// 用户信息
	UserInfo *PersonalInfo `json:"userInfo"`
	// 创建人信息（添加人）
	CreatorInfo *PersonalInfo `json:"creatorInfo"`
	// 用户角色
	UserRole *UserRoleInfo `json:"userRole"`
	// 类型（1用户2部门）
	Type int `json:"type"`
	// 部门信息
	DepartmentInfo *DepartmentSimpleInfo `json:"departmentInfo"`
}
