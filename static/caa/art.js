var scrollLoadMore;
$(function(){
	// 主菜单
	var $main_nav 	= $("#mainNav>li");
	var $sub_nav 	= $main_nav.find(".sub-nav");
	$main_nav.hover(function(){
		var $t = $(this);
		var $s = $t.siblings('li');
		var $to = $t.find("dl");
		var $so = $s.find("dl");
		$to.addClass('current');
		$so.removeClass('current');
	},function(){
		var $t = $(this);
		var $to = $t.find("dl");
		$to.removeClass('current');
	})

	// 内页菜单
	var $inMenu 	= $("#inMenu");
	var $menu_btn 	= $("#inMenu>.btn-menu");
	var $menu_main 	= $("#inMenu>.menu-wrap");
	var $menu_sub 	= $("#inMenu>.sub-wrap");
	$inMenu.hover(function(){
		var mw = 200;
		var mh = $menu_main.height()+100;
		$inMenu.width(mw).height(mh).addClass('open').show();
		$menu_main.stop().css({left:0});
		$menu_sub.height($menu_main.height());
	},function(){
		var mw = 80;
		var mh = 30;
		$menu_main.stop().css({left:-200});
		$menu_sub.stop().css({left:-360})
		$inMenu.width(mw).height(mh).removeClass("open").show();
	});
	$menu_main.find('li').hover(function(){
		var $t 			= $(this);
		var $sub 		= $($t.children('a').attr('data-id'));
		var ul_length 	= $sub.find("ul").length
		var sub_w 		= ul_length*180;
		$sub.show();
		$sub.siblings('div').hide();
		if(ul_length>0){
			$inMenu.width(sub_w+200);
			$sub.parent().width(sub_w);
			$menu_sub.stop().css({left:200});
		}else{
			$sub.parent().width(180);
			$menu_sub.stop().css({left:-360});
			$inMenu.width(sub_w+200);
		}
	});
	// 首页切换学术报告、通知公告自动切换
	var $tabArr = ["[data-id='#xsyg']","[data-id='#tzgg']"];
	var tabSI = null;
	var interTabsPlay = function(){
		try{
			clearInterval(tabSI);
			tabSI = setInterval(function(){
				$.each($tabArr,function(k,v){
					var $t = $(v);
					var $o = $($t.attr('data-id'));
					var isCurrent = $t.hasClass('current');
					if(isCurrent){
						$t.removeClass('current');
						$o.hide();
					}else{
						$t.addClass('current');
						$o.show();
					}
				});
			},8000);
		}catch(e){}
	}
	$.each($tabArr,function(k,v){
		var $t = $(v);
		$t.hover(function(){
			clearInterval(tabSI);
		},function(){
			interTabsPlay();
		})
	});
	interTabsPlay();


	// 添加国美日新class
	if($("#dailyList")){
		try{
			dailyListAddClass($("#dailyList>li"));
			
		}catch(e){}
	}
	
	// .tabs
	if($("[data-toggle='tabs']")){
		try{
			var $tabs = $("[data-toggle='tabs']");
			$tabs.each(function(){
				var $t = $(this);
				$t.on("click",".tab-nav",function(){
					var $t = $(this);
					var $ts = $t.siblings('.tab-nav');
					var $o = $($t.attr("data-id"));
					var $s = $o.siblings('.tab-plane');
					$t.addClass('current');
					$ts.removeClass('current');
					$o.show();
					$s.hide();
				})
				
			});
		}catch(e){}
	}
	// .tabs
	if($("[data-toggle='tabs-hover']")){
		try{
			var $tabs 	= $("[data-toggle='tabs-hover']");

			$tabs.each(function(){
				var $t = $(this);
				var $t_s = $t.siblings('.tab-plane')
				$t.find(".tab-nav").hover(function(){
					var $t = $(this);
					var $ts = $t.siblings('.tab-nav');
					var $o = $($t.attr("data-id"));
					// var $s = $o.siblings('.tab-plane');
					$t_s.hide();
					$t.addClass('current');
					$ts.removeClass('current');
					$o.show();
				},function(){})
			});
		}catch(e){}
	}

	function bookScroll(){
		try{
			if($(".book-scroll")){
				var $book 		= $(".book-scroll");
				var $bookWrap 	= $(".book-scroll-wrap>ul");
				var $bookPrev	= $book.find(".btn-prev");
				var $bookNext	= $book.find(".btn-next");
				var $bookItem	= $bookWrap.find("li");
				var bookCount	= $bookItem.length;
				var pageNum 	= 4;
				var cur 		= 0;

				var play = function(){
					var $cur = $($bookItem[cur]);
					$cur.animate({width:549},500,function(){
						$cur.addClass('open');
					});
					$cur.siblings('li.open').stop().animate({width:129},500,function(){
						$cur.siblings('li.open').removeClass('open');
					});
					if(cur-pageNum>=-1){
						$bookWrap.stop().animate({"margin-left":(cur-pageNum+2)*-139},500);
					}else{
						$bookWrap.stop().animate({"margin-left":0},500);
					}

				}
				$bookItem.on("click",function(){
					cur = $(this).index();
					play();
				});
				$bookPrev.on("click",function(){
					cur = cur - 1 < 0 ? 0 : cur - 1;
					play();
				});
				$bookNext.on("click",function(){
					cur = cur + 1 >= bookCount ? 0 : cur + 1;
					play();
				});

			}
		}catch(e){}
	}

	bookScroll();

});

function dailyListAddClass($id){
	var winWidth = $(window).width()/2;
	$id.each(function(){
		var $t = $(this);
		var tLeft = winWidth - $t.offset().left;
		if(tLeft<-20){
			$t.addClass('left')
		}
		// if($t)
	});
}
// load more content
var scrollLoadMore = function (obj,url){
	var isLoad = true;
	var page = 0;
	$(window).scroll(function(event) {
		if(isLoad){
			var dTop = $(document).scrollTop();
			var dH = $(document).height();
			var wH = $(window).height();
			if(dH-10 <= dTop+wH){
				$(".loading").fadeIn('fast');
				var tmpUrl = page==0?url+".html":url+"_"+page+".html";
				isLoad = false;
				$.ajax({
					type: "GET",
		            url: tmpUrl,
		            dataType: "html",
					success:function(data){
						page=page+1;
						setTimeout(function(){
							$(".loading").hide();
							$data = $(data);
							$(obj).append($data.find(obj));
							dailyListAddClass($(obj+">li"));
							isLoad=true;
						}, 2000);
					},
					error:function(err){
						if(err.status==404){
							$(".loading").find("span").html("<span style='color:red'>已加载所有内容</span>");
							isLoad = false;
							setTimeout(function(){
								$(".loading").slideUp("slow");
							},3000);
						}
					}
				});
			}
		}
	});
}



$(window).load(function(){
	if($(".slide-items")){
		try{
			// banner切换
			$('.slide-items').anythingSlider({
				theme           : 'metallic',
				mode            : 'f',   // fade mode - new in v1.8!
				resizeContents  : false, // If true, solitary images/objects in the panel will expand to fit the viewport		}
				autoPlay		: true
			});
		}catch(e){}
	}
	// 内页图片
	if($(".fancybox")){
		try{
			$(".fancybox").fancybox({
				'overlayColor'	: '#000',
				'transitionIn'	: 'elastic',
				'transitionOut'	: 'elastic',
				'titlePosition' : 'over',
				'titleFormat'   : function(title, currentArray, currentIndex, currentOpts) {
			    return '<span id="fancybox-title-over">Image ' +  (currentIndex + 1) + ' / ' + currentArray.length + ' ' + title + '</span>';
				}
			});
		}catch(e){}
	}
	// 首页banner
	if($('#dg-container')){
		try{
			$('#dg-container').gallery({
				autoplay	:	true
			});
		}catch(e){}
	}
})