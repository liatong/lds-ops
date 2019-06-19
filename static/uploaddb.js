

$(document).ready(function(){

  $("#uploaddb").bind("click",function () {
    // body...
    console.log("upload db script.")
    $.ajax({
        url: '/dbscript',
        type: 'POST',
        cache: false,
        data: new FormData($('#uploadDBForm')[0]),
        processData: false,
        contentType: false
    }).success(function(data) {
        console.log("success");
        alert("数据库脚本文件上传成功");
        window.location.href = "/dbscript";
    }).error(function() {
        console.log("fail");
        alert("数据库文件上传失败。");
    }); 
    //end 
  });
 

});
