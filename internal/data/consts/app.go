package consts

var (
	MirrorApp = 6 // 视图app
)

// 应用类型，1：表单，2：仪表盘，3：文件夹，4：项目, 5: 汇总表
const (
	LcAppTypeUnknown         = 0
	LcAppTypeForForm         = 1
	LcAppTypeForFolder       = 3
	LcAppTypeForPolaris      = 4
	LcAppTypeForSummaryTable = 5
	LcAppTypeForViewMirror   = 6
)
