package model 

import (
	"github.com/liatong/lds-ops/sql"
	"fmt"
)

type Dbscript struct{
	Name string `json:"scriptname" form:"scriptname"`
	App string `json:"app" form:"app"`
	AppVersion string `json:"appversion" form:"appversion"`
	Branch string `json:"branch" form:"branch"`
	FilePath string `json:"filepath" form:"filepath"`
	CreateTime string `json:"createtime" form:"createtime"`
	Comment string `json:"comment" form:"comment"`
}

func(p *Dbscript)UploadDbcheck()error{

	//res,err := sql.Pool.Exec("insert into ops_package (`filename`,`application_name`,`env`,`version`,`mdcode`,`filepath`,`create_time`,`upload_time`) values ('admin-service.jar','admin-service','eu','2019050200','333333','/tmp/upload/admin-service/5.0/DSC_4205.JPG','2019-05-02 00:00:00','2019-05-02 00:00:00') ")
	res,err := sql.Pool.Exec("insert into ops_dbscript (`scriptname`,`app`,`appversion`,`branch`,`filepath`,`createtime`,`comment`) values (?,?,?,?,?,?,?) on duplicate key update filepath=?",p.Name,p.App,p.AppVersion,p.Branch,p.FilePath,p.CreateTime,p.Comment,p.FilePath)
	if err != nil {
		return err 
	}
	fmt.Print(res)
	return nil
}

