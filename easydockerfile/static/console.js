var socket;

$(document).ready(function () {
    // Create a socket
    //socket = new WebSocket('ws://' + window.location.host + '/console/sync?tutname=' + $('#tutname').text());
    //socket = new WebSocket('ws://' + '127.0.0.1:8089' + '/console/sync?tutname=' + $('#tutname').text());

    socket = new WebSocket('ws://' + 'localhost:8080' + '/v1/testbuild/');

Messenger.options = {
    parentLocations: ['article'],
   // extraClasses: 'messenger-fixed messenger-on-bottom messenger-on-right',
    theme: 'air',

}

	Messenger().post({
	  message: 'There was an explosion while processing your request.',
	  type: 'error',
	  showCloseButton: true
	});

    // Message received on the socket
    socket.onmessage = function (event) {
        //var line = JSON.parse(event.data);
	//$('#console-output').append(data+"<br>")
        //console.log(data);
	//console.log(event);
	console.log(event)
	$('#console-output').append(event.data)
    };

    // Send messages.
    var postCode = function () {
        var uname = $('#tutname').text();
        var content = $('#code').text();
        $.post(
	    '/v1/testbuild/',
	    {data:content},
	    function(data){
		console.log(data);
		});
    }

    $('#submitbtn').click(function () {
        postCode();
    });

    $('#clear-btn').click(function () {
        $('#console-output').html("");
	//$('#console-output').empty();
    });
});
