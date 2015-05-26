$(document).ready(function () {
    var postConecnt = function () {
        var code = $('#code').text();
        $.post("", 
        {
	   code:code
        },
	function(data,status){
        console.log("Data: " + data + "\nStatus: " + status);
    	});
    }
//    var postAjax = function () {
//		var code = $('#code').text();
//		$.ajax({
//		    type: 'POST',
//		    url: "v1/testbuild" ,
//		    data: code,
//		    dataType:'jsonp',
//	            jsonp: "jsonpcallback",
//		    success: function(data){
//		    	console.log(data)
//			$("console").append(data+"<br>")
//	 	    }
			
//		});

//    }

    $('#debugbtn').click(function () {

	console.log("asdadasd");
        postConecnt();
//		postAjax();
    });
});
