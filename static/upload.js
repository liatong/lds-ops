

$(document).ready(function(){

  $("#upload").bind("click",function () {
    // body...
    console.log("hello world.")
    $.ajax({
        url: '/upload',
        type: 'POST',
        cache: false,
        data: new FormData($('#uploadForm')[0]),
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
