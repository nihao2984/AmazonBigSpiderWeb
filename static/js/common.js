$(function(){ 

	 /*banner*/
    banner('.bannerNew_btn li', 'click', '.bannerNew_f .bannerNew_s', '.bannerNew', 5000, '.bannerNew_l', '.bannerNew_r');
	 function banner(btn,className,son,box,time,left,right){
			var now_page = 0;
			var page_num = $(son).length;
			$(son).eq(now_page).css({zIndex:5}).show();
			$(son).eq(0).siblings().css({zIndex:3}).hide();
			var timer = null;
			$(btn).eq(now_page).addClass(className);
			$(btn).click(function(){
					if($(this).index()!=now_page&&!$(son).is(":animated")){
							$(son).eq(now_page).fadeOut(500,function(){
									$(this).css({zIndex:3});
								})
							$(son).eq($(this).index()).css({zIndex:5}).fadeIn(500);
							$(btn).removeClass(className);
							$(btn).eq($(this).index()).addClass(className);
							now_page = $(this).index();
						}
				})
			function bannerAuto(){
						if(!$(son).is(":animated")){
							if(now_page<page_num-1){
								$(son).eq(now_page).fadeOut(500,function(){
										$(this).css({zIndex:3});
									})
								$(son).eq(now_page+1).css({zIndex:5}).fadeIn(500);
								$(btn).removeClass(className);
								$(btn).eq(now_page+1).addClass(className);
								now_page++
							}else{							
								$(son).eq(page_num-1).fadeOut(500,function(){
										$(this).css({zIndex:3});
									})
								$(son).eq(0).css({zIndex:5}).fadeIn(500);
								$(btn).removeClass(className);
								$(btn).eq(0).addClass(className);
								now_page=0;
							}
						}
				}

			$(left).click(function(){
						if(!$(son).is(":animated")){
							if(now_page==0){
								$(son).eq(0).fadeOut(500,function(){
										$(this).css({zIndex:3});
									})
								$(son).eq(page_num-1).css({zIndex:5}).fadeIn(500);
								$(btn).removeClass(className);
								$(btn).eq(page_num-1).addClass(className);
								now_page=page_num-1;
							}else{							
								$(son).eq(now_page).fadeOut(500,function(){
										$(this).css({zIndex:3});
									})
								$(son).eq(now_page-1).css({zIndex:5}).fadeIn(500);
								$(btn).removeClass(className);
								$(btn).eq(now_page-1).addClass(className);
								now_page--;
							}
						}

			})
			$(right).click(function(){
						if(!$(son).is(":animated")){
							if(now_page<page_num-1){
								$(son).eq(now_page).fadeOut(500,function(){
										$(this).css({zIndex:3});
									})
								$(son).eq(now_page+1).css({zIndex:5}).fadeIn(500);
								$(btn).removeClass(className);
								$(btn).eq(now_page+1).addClass(className);
								now_page++
							}else{							
								$(son).eq(page_num-1).fadeOut(500,function(){
										$(this).css({zIndex:3});
									})
								$(son).eq(0).css({zIndex:5}).fadeIn(500);
								$(btn).removeClass(className);
								$(btn).eq(0).addClass(className);
								now_page=0;
							}
						}

				})

			timer= setInterval(bannerAuto,time);
			$(box).hover(function(){
					clearInterval(timer);
				},function(){
					timer = setInterval(bannerAuto,time);
			})
		 }
	
	
	
	
	
	


})
