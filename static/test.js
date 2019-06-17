
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

function showDataTable(tb,obj,stdhtml,endtdhtml,bindfunc){
        
        var rt = 1;
        tb.children().remove();
        if(typeof obj != "undefined" && typeof obj == "object" && typeof tb == "object"){
            $.each(obj,function(key,value){
                var itr = document.createElement('tr');
                //create tr fist td,if have stdhtml
                if( typeof stdhtml == 'string'){
                    var std = document.createElement('td');
                    std.innerHTML = stdhtml;
                    itr.appendChild(std);
                }
                //console.log(value);
                /*
                for ( key in value){
                	console.log(key);
                	var btd = document.createElement('td');
                    btd.innerText= value[key];
                    itr.appendChild(btd);

                };
                */
                
                value.map(function(text){
                    var btd = document.createElement('td');
                    btd.innerText= text;
                    itr.appendChild(btd);
                });
                
                // create define td when have stdhtml
                if( typeof endtdhtml == 'string'){
                    var endtd = document.createElement('td');
                    endtd.innerHTML = endtdhtml;
                    itr.appendChild(endtd);
                }
                
                //append tr to tbody
                tb.append(itr);
            });
            
            //  callback append html even function
            if( typeof bindfunc == "function" ){
                bindfunc();
            };
            rt = 1;
            
        }else{
            console.log('show table is obj is undefine.');
            rt = 0;
        }
        console.log('debug');
        return rt;
        
    }
// 翻页功能
   $("#upPage").click(function(){
        
        if($("#pageNum").text() != '1'){
            $("#pageNum").text(parseInt($("#pageNum").text())-1);
            console.log($("#pageNum").text());
            getPackageList();
        }
    })

    $("#downPage").click(function(){
        $("#pageNum").text(parseInt($("#pageNum").text())+1);
        getPackageList();
    })
