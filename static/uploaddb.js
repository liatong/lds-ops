

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
        alert("文件上传成功");
    }).error(function() {
        console.log("fail");
        alert("文件上传失败。");
    }); 
    //end 
  });
 

});
