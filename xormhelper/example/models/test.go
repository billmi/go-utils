package models

type Test struct {
	Id                 int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	DeviceType         int    `json:"device_type" xorm:"not null default 0 comment('设备类型：1-') unique INT(11)"`
	DeviceTypeName     string `json:"device_type_name" xorm:"not null VARCHAR(64)"`
	DeviceTypeTitle    string `json:"device_type_title" xorm:"not null default '' comment('设备类型标题') VARCHAR(64)"`
	RId                int    `json:"r_id" xorm:"not null INT(11)"`
	DeviceScreenWidth  int    `json:"device_screen_width" xorm:"not null default 0 comment('屏宽') INT(11)"`
	DeviceScreenHeight int    `json:"device_screen_height" xorm:"not null default 0 comment('屏高') INT(11)"`
	UpdateTime         int    `json:"update_time" xorm:"not null default 0 INT(11)"`
	CreateTime         int    `json:"create_time" xorm:"not null default 0 INT(11)"`
}
