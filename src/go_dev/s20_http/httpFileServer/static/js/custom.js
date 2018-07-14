$('.navbar-toggler-icon').click( function() {
			$(this).toggleClass('fa-times');
			$(this).toggleClass('fa-bars');
			$('.main-navbar').toggleClass('show');
});

$(window).scroll(function() {

	if ($(window).scrollTop() > 40) {
				$('header').addClass("active");
	}
	else{
		$('header').removeClass("active");
	}
});

 if (screen.width > 1024) {
if ( $('.portipolio-sec').length > 0 ) {

   AOS.init({
        easing: 'ease-in-out-sine'
      });
}
 }