/**
 * 图片管理类
 */

 export default{

  //Common 
  common:{
    ic_header_bg:require('../assets/Images/Home/header_bg.png'),
    ic_up_arrow:require('../assets/Images/Shop/up.png'),
    
    //轮播图
    ic_banner_data:[
      {image:require('../assets/Images/Ads/ad-01.jpg')},
      {image:require('../assets/Images/Ads/ad-02.jpg')},
      {image:require('../assets/Images/Ads/ad-03.jpg')},
      {image:require('../assets/Images/Ads/ad-04.jpg')},
      {image:require('../assets/Images/Ads/ad-05.jpg')},
      {image:require('../assets/Images/Ads/ad-06.jpg')},
      {image:require('../assets/Images/Ads/ad-07.jpg')},
      {image:require('../assets/Images/Ads/ad-08.jpg')}
    ],

    ic_no_network:require('../assets/Images/Shop/noNetWork.png'),
    ic_cart_null:require('../assets/Images/Shop/NULL.png'),
    ic_kong:require('../assets/Images/Mine/kong.png'),
    ic_good:require('../assets/Images/Mine/Good.png'),
    ic_good_on:require('../assets/Images/Mine/Good1.png'),
    ic_middle:require('../assets/Images/Mine/Middle.png'),
    ic_middle_on:require('../assets/Images/Mine/Middle1.png'),
    ic_bad:require('../assets/Images/Mine/Bad.png'),
    ic_bad_on:require('../assets/Images/Mine/Bad1.png')
  },

  //首页
  home:{
    ic_gift:require('../assets/Images/tabbar/lipin.png'),
    ic_input:require('../assets/Images/Home/input.png'),
    ic_jinbi2:require('../assets/Images/Home/jinbi2.png'),
    ic_integral_data:[
      require('../assets/Images/Shop/time_1_1.png'),
      require('../assets/Images/Shop/time_2_1.png'),
      require('../assets/Images/Shop/time_3_1.png'),
      require('../assets/Images/Shop/time_4_1.png'),
      require('../assets/Images/Shop/time_1_1.png'),
      require('../assets/Images/Shop/time_2_1.png'),
      require('../assets/Images/Shop/time_3_1.png'),
      require('../assets/Images/Shop/time_4_1.png'),
    ]
  },

  //Shop
  shop:{
    //外层
    ic_cart:require('../assets/Images/Shop/cart.png'),
    ic_category_else:require('../assets/Images/tabbar/else.png'),
    ic_timer_data:[
      {image:require('../assets/Images/Shop/time_1.png')},
      {image:require('../assets/Images/Shop/time_2.png')},
      {image:require('../assets/Images/Shop/time_3.png')},
      {image:require('../assets/Images/Shop/time_4.png')}
    ],
    ic_category_icon:[
      {normal:require('../assets/Images/Shop/shop_class_all_icon.png'),selected:require('../assets/Images/Shop/shop_class_all_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_cf_icon.png'),selected:require('../assets/Images/Shop/shop_class_cf_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_dq_icon.png'),selected:require('../assets/Images/Shop/shop_class_dq_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_office_icon.png'),selected:require('../assets/Images/Shop/shop_class_office_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_qj_icon.png'),selected:require('../assets/Images/Shop/shop_class_qj_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_ryp_icon.png'),selected:require('../assets/Images/Shop/shop_class_ryp_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_cf_icon.png'),selected:require('../assets/Images/Shop/shop_class_cf_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_dq_icon.png'),selected:require('../assets/Images/Shop/shop_class_dq_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_office_icon.png'),selected:require('../assets/Images/Shop/shop_class_office_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_qj_icon.png'),selected:require('../assets/Images/Shop/shop_class_qj_icon_on.png')},
      {normal:require('../assets/Images/Shop/shop_class_ryp_icon.png'),selected:require('../assets/Images/Shop/shop_class_ryp_icon_on.png')}
    ],
    //商品详情
    ic_back:require('../assets/Images/Shop/back.png'),
    ic_item:require('../assets/Images/Shop/item10.png'),
    ic_bg:require('../assets/Images/Mine/bg.png'),
    ic_btn_bg:require('../assets/Images/Mine/Set/exit.png'),

    //确认订单
    ic_address_icon:require('../assets/Images/Shop/addr.png'),
    ic_line:require('../assets/Images/Shop/orderLine.png'),
    ic_no_address:require('../assets/Images/Shop/addressEmpt.png'),
    ic_cancel:require('../assets/Images/Shop/cancle.png'),
    ic_sure:require('../assets/Images/Shop/sure.png'),

    //支付
    ic_up:require('../assets/Images/Shop/up1.png'),
    ic_down:require('../assets/Images/Shop/down.png'),
    ic_cancelPay:require('../assets/Images/Shop/canclePay.png'),
    ic_radio_normal:require('../assets/Images/Shop/radio_normal.png'),
    ic_radio_selected:require('../assets/Images/Shop/radio_selected.png'),

    //选择收货地址
    ic_addr_bg:require('../assets/Images/Mine/background.png'),
    ic_more:require('../assets/Images/Mine/Set/moreFeedback.png'),
    ic_delete:require('../assets/Images/Mine/Set/delete.png'),
    

  },

  //Business
  business:{
    ic_item_bg:[
      {image:require('../assets/Images/Shop/time_1_1.png')},
      {image:require('../assets/Images/Shop/time_2_1.png')},
      {image:require('../assets/Images/Shop/time_3_1.png')},
      {image:require('../assets/Images/Shop/time_4_1.png')},
    ]
  },

  //Activity
  activity:{
    
  },

  //MineView
  mineView:{
    ic_set:require('../assets/Images/Mine/set.png'),
    ic_info:require('../assets/Images/Mine/information.png'),
    ic_default_head1:require('../assets/Images/Mine/head1.png'),
    ic_default_head2:require('../assets/Images/Mine/head2.png'),
    ic_go:require('../assets/Images/Mine/go.png'),
    ic_btn_bg:require('../assets/Images/Mine/Set/exit.png'),
    ic_bg:require('../assets/Images/Mine/background.png'),
    ic_label_data:[
      require('../assets/Images/Mine/balance.png'),
      require('../assets/Images/Mine/pay.png'),
      require('../assets/Images/Mine/deliver.png'),
      require('../assets/Images/Mine/stay.png'),
      require('../assets/Images/Mine/comment.png'),
      require('../assets/Images/Mine/all.png'),
      require('../assets/Images/Mine/balance.png'),
      require('../assets/Images/Mine/issue.png'),
      require('../assets/Images/Mine/ourselves.png')
    ],
    ic_info_data:[
      require('../assets/Images/Mine/SeltInfo/nickName.png'),
      require('../assets/Images/Mine/SeltInfo/sex.png'),
      require('../assets/Images/Mine/SeltInfo/date.png'),
      require('../assets/Images/Mine/SeltInfo/email.png'),
      require('../assets/Images/Mine/SeltInfo/realName.png')
    ],
    ic_select:require('../assets/Images/Mine/Autonym/select.png'),
    ic_just:require('../assets/Images/Mine/Autonym/just.png'),
    ic_against:require('../assets/Images/Mine/Autonym/against.png')
    
  },

  //登录/注册/找回密码
  login:{
    ic_bg: require('../assets/Images/Login/bg.png'),
    ic_logo: require('../assets/Images/Login/logo.png'),
    ic_account: require('../assets/Images/Login/Account.png'),
    ic_lock: require('../assets/Images/Login/PassWordAgion.png'),
    ic_unlock: require('../assets/Images/Login/password.png'),
    ic_login_btn: require('../assets/Images/Login/Signinbutton.png'),
    ic_border:require('../assets/Images/Register/reactangle3.png'),
    ic_code_btn:require('../assets/Images/Register/reactangle2.png'),
    ic_register_btn:require('../assets/Images/Register/register.png'),
    ic_qq: require('../assets/Images/Login/qqSignIn.png'),
    ic_wx: require('../assets/Images/Login/WeChatSignIn.png'),
    ic_wb: require('../assets/Images/Login/MicroBlogSignIn.png')
  },

  //底部标签
  tabController:{
    ic_home_normal:require('../assets/Images/tabbar/home.png'),
    ic_home_selected:require('../assets/Images/tabbar/home_1.png'),
    ic_shop_normal:require('../assets/Images/tabbar/shangcheng.png'),
    ic_shop_selected:require('../assets/Images/tabbar/shangcheng_1.png'),
    ic_business_normal:require('../assets/Images/tabbar/shangjia.png'),
    ic_business_selected:require('../assets/Images/tabbar/shangjia_1.png'),
    ic_activity_normal:require('../assets/Images/tabbar/huodong.png'),
    ic_activity_selected:require('../assets/Images/tabbar/huodong_1.png'),
    ic_mine_normal:require('../assets/Images/tabbar/my.png'),
    ic_mine_selected:require('../assets/Images/tabbar/my_1.png')
  },

  //钱包
  wallet:{
    ic_bg:require('../assets/Images/Mine/background.png'),
    ic_back:require('../assets/Images/Mine/back.png'),
    ic_block:require('../assets/Images/Wallet/block.png')
  },

  //设置
  setting:{
    ic_go:require('../assets/Images/Mine/go.png'),
    ic_btn_bg:require('../assets/Images/Mine/Set/exit.png'),
    ic_more:require('../assets/Images/Mine/more.png'),
    ic_omit:require('../assets/Images/Mine/Set/omit.png'),
    ic_omit_data:[
      require('../assets/Images/Mine/Set/moreInformation.png'),
      require('../assets/Images/Mine/Set/moreHomePage.png'),
      require('../assets/Images/Mine/Set/moreFeedback.png'),
      require('../assets/Images/Mine/Set/moreService.png')
    ],
    ic_bg:require('../assets/Images/Mine/background.png'),
    ic_radio_normal:require('../assets/Images/Shop/radio_normal.png'),
    ic_radio_selected:require('../assets/Images/Shop/radio_selected.png'),
    ic_edit:require('../assets/Images/Mine/Set/moreFeedback.png'),
    ic_delete:require('../assets/Images/Mine/Set/delete.png'),
    ic_arrow:require('../assets/Images/Mine/up.png'),
    ic_confirm:require('../assets/Images/Mine/Set/confirm.png')
  },

  //我的订单
  indent:{
    ic_bg:require('../assets/Images/Mine/background.png'),
    ic_back:require('../assets/Images/Mine/back.png'),
    ic_info:require('../assets/Images/Mine/information.png'),
    ic_wuli:require('../assets/Images/Shop/wuliu.png'),
    ic_wallet:require('../assets/Images/Mine/wallet.png'),
    ic_addr:require('../assets/Images/Mine/addr.png'),
    ic_line:require('../assets/Images/Shop/orderLine.png'),
    ic_no_addr:require('../assets/Images/Shop/addressEmpt.png'),
    ic_cancel:require('../assets/Images/Shop/cancle.png'),
    ic_sure:require('../assets/Images/Shop/sure.png'),
    ic_tab_normal:[
      require('../assets/Images/Mine/pay.png'),
      require('../assets/Images/Mine/deliver.png'),
      require('../assets/Images/Mine/stay.png'),
      require('../assets/Images/Mine/comment.png'),
      require('../assets/Images/Mine/all.png')
    ],
    ic_tab_selected:[
      require('../assets/Images/Mine/pay.png'),
      require('../assets/Images/Mine/deliver.png'),
      require('../assets/Images/Mine/stay.png'),
      require('../assets/Images/Mine/comment.png'),
      require('../assets/Images/Mine/all.png')
    ],
    
  },

  //组合组件
  myintergral:{
    ic_info_bg:require('../assets/Images/Home/user_info_bg.png'),
    ic_default_head:require('../assets/Images/Mine/head1.png'),
    ic_jinbi1:require('../assets/Images/Home/jinbi1.png')
  },
	
	//聊天
	chat:{
		ic_audiom:require('../assets/Images/chat/audiom.png'),
		ic_audio:require('../assets/Images/chat/audio.png'),
		ic_kbord:require('../assets/Images/chat/kbord.png'),
		ic_smile:require('../assets/Images/chat/smile.png'),
		ic_more:require('../assets/Images/chat/more.png'),
		ic_send:require('../assets/Images/chat/send.png'),
		ic_upload:require('../assets/Images/chat/upload.png'),
		ic_camera:require('../assets/Images/chat/camera.png'),
		ic_audiocall:require('../assets/Images/chat/audiocall.png'),
		ic_videocall:require('../assets/Images/chat/videocall.png'),
		ic_redpackage:require('../assets/Images/chat/redpackage.png'),
		ic_exchange:require('../assets/Images/chat/exchange.png'),
		ic_address:require('../assets/Images/chat/address.png'),
		ic_person:require('../assets/Images/chat/person.png')
	}
 }
