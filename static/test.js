
$(document).ready(function(){
	$("#upload").bind("click",function () {
		// body...
		console.log("hello world.")
	})

	getPackageList()

});

function getPackageList(){
	var gurl ='/package'
	var page = $("#pageNum").text();
	var pagesize = $("#pageSize").text();
	var postdata = {};
	postdata["page"]=page;
	postdata['pagesize']=pagesize;
	postdata['application']="admin-service";
	postdata['version']="20190502";
	postdata['enviroment']="as";
	console.log("test");

	console.log(gurl,page,pagesize,postdata);

	$.ajax(gurl,{dataType:'json',type:'POST',data:postdata}).done(function(data){
            
            console.log(data['data']);
            let tb = $('#showproject tbody');
            //组成data数据
            var  tableData = {}

            for (key in data['data']){
                tableData[key]=[];
                tableData[key].push(data['data'][key]['application_name']);
                tableData[key].push(data['data'][key]['env']);
                tableData[key].push(data['data'][key]['filename']);
                tableData[key].push(data['data'][key]['version']);
                //tableData[key].push(data['data'][key]['mocode']);
                //tableData[key].push(data['data'][key]['filepath']);
                tableData[key].push(data['data'][key]['create_time']);
                tableData[key].push(data['data'][key]['upload_time']);
            }
            showDataTable(tb,tableData,null,null,null);
        }).fail(function(data){
            console.log("error")
        })
}
