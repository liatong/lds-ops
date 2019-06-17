package model 

import (
	"github.com/liatong/lds-ops/sql"
	"fmt"
)

type Package struct{
	Name string `json:"filename" form:"filename"`
	Application string `json:"application_name" form:"application_name"`
	Enviroment string `json:"env" form:"env"`
	Version string `json:"version" form:"version"`
	Mdcode string `json:"mocode" form:"mdcode"`
	FilePath string `json:"filepath" form:"filepath"`
	CreateTime string `json:"create_time" form:"create_time"`
	UploadTime string `json:"upload_time" form:"upload_time"`

}

func(p *Package)UploadPackage()error{

	//res,err := sql.Pool.Exec("insert into ops_package (`filename`,`application_name`,`env`,`version`,`mdcode`,`filepath`,`create_time`,`upload_time`) values ('admin-service.jar','admin-service','eu','2019050200','333333','/tmp/upload/admin-service/5.0/DSC_4205.JPG','2019-05-02 00:00:00','2019-05-02 00:00:00') ")
	res,err := sql.Pool.Exec("insert into ops_package (`filename`,`application_name`,`env`,`version`,`mdcode`,`filepath`,`create_time`,`upload_time`) values (?,?,?,?,?,?,?,?) on duplicate key update mdcode=?,upload_time=?,filepath=?",p.Name,p.Application,p.Enviroment,p.Version,p.Mdcode,p.FilePath,p.CreateTime,p.UploadTime,p.Mdcode,p.UploadTime,p.FilePath)
	if err != nil {
		return err 
	}
	fmt.Print(res)
	return nil
}

