$("#menu-toggle").click(function(e) {
	e.preventDefault();
	$("#wrapper").toggleClass("toggled");
});

$('#sidebar').affix({
	offset: {
		top: 0
	}
});

var $body   = $(document.body);
var navHeight = $('.navbar').outerHeight(true) + 10;

$body.scrollspy({
	target: '#leftCol',
	offset: navHeight
});

$(document).ready(function(){
	$('.sidebar ul li a').click(function(){
		$('html, body').animate({
			scrollTop: $( $.attr(this, 'href') ).offset().top - 60
		}, 500);

		if( $(window).width() < 979 ){
			$('.navbar-toggle').trigger('click');
		}
		return false;
	});
	
	$(window).resize(function(){
		if( $(window).width() > 979 && ( $('#leftCol').attr('aria-expanded') == 'undefined' || $('#leftCol').attr('aria-expanded') == 'false' ) ){
			$('.navbar-toggle').trigger('click');
		}
	});
});