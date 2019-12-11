import Vue from 'vue'
import Router from 'vue-router'
import login from '@/pages/login.vue'		//登录页面
import register from '@/pages/register.vue'		//注册页面
import retrievePassword from '@/pages/retrievePassword.vue'	//找回密码页面
import main from '@/pages/main.vue'		//主页面
import home from '@/pages/home.vue'		//首页
import mall from '@/pages/mall.vue'		//商城
import merchant from '@/pages/merchant.vue'		//商家
import activity from '@/pages/activity.vue'		//活动
import mine from '@/pages/mine.vue'		//我的
import userinfo from '@/pages/userinfo/userinfo.vue'	//用户信息页
import modifyInfo from '@/pages/userinfo/modifyInfo.vue'	//修改用户信息页
import setting from '@/pages/setting/setting.vue'		//设置
import settingPage from '@/pages/setting/settingPage.vue'	//设置二级页面
import modifyAddress from '@/pages/setting/modifyAddress.vue'	//修改添加收货地址
import loginPassword from '@/pages/setting/loginPassword.vue'	//修改登录密码
import tradingPassword from '@/pages/setting/tradingPassword.vue'	//修改交易密码
import wallet from '@/pages/wallet.vue'		//钱包
import goodsDetail from '@/pages/goodsDetail.vue'		//商品详情页
import allComments from '@/pages/allComments.vue'		//商品评价页
import cart from '@/pages/cart.vue'			//购物车
import confirmOrder from '@/pages/confirmOrder.vue'		//确认订单
import paymentPage from '@/pages/paymentPage.vue'			//订单支付页
import orderPage from '@/pages/orderPage/orderPage.vue'		//订单主页面
import unpaid from '@/pages/orderPage/unpaid.vue'		//待付款页面
import unshipped from '@/pages/orderPage/unshipped.vue'		//待发货页面
import unreceived from '@/pages/orderPage/unreceived.vue'		//待收货页面
import comment from '@/pages/orderPage/comment.vue'		//待评价页面
import allorders from '@/pages/orderPage/allorders.vue'		//全部订单页面
import orderDetail from '@/pages/orderPage/orderDetail.vue'		//订单详情页面
import commentPage from '@/pages/orderPage/commentPage.vue'		//评价页面
import refundReturnApply from '@/pages/orderPage/refundReturnApply.vue'		//退款退货申请页面
import refundAndReturnPage from '@/pages/orderPage/refundAndReturnPage.vue'		//退款退货列表页面
import feedback from '@/pages/feedback.vue'		
import chatPage from '@/pages/chat/chatPage.vue'
import chatRoom from '@/pages/chat/chatRoom.vue'

Vue.use(Router)

// 注册全局函数,路由地址跳转相同时捕获异常
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

export default new Router({
	routes:[
		{
			path:'/',
			redirect:'/main/home'
		},
		{
			path:'/login',
			name:'login',
			component:login
		},
		{
			path:'/register',
			name:'register',
			component:register
		},
		{
			path:'/retrievePassword',
			name:'retrievePassword',
			component:retrievePassword
		},
		{
			path:'/main',
			name:'main',
			component:main,
			meta:{
				requireAuth:true
			},
			children:[
				{
					path:'home',
					name:'home',
					component:home,
					meta:{
						keepAlive:true
					}
				},
				{
					path:'mall',
					name:'mall',
					component:mall,
					meta:{
						keepAlive:true
					}
				},
				{
					path:'merchant',
					name:'merchant',
					component:merchant,
					meta:{
						keepAlive:true
					}
				},
				{
					path:'activity',
					name:'activity',
					component:activity,
					meta:{
						keepAlive:true
					}
				},
				{
					path:'mine',
					name:'mine',
					component:mine,
					meta:{
						keepAlive:true
					}
				}
			]
		},
		{
			path:'/userinfo',
			name:'userinfo',
			component:userinfo,
		},
		{
			path:'/modifyNickname',
			name:'modifyNickName',
			component:modifyInfo,
			meta:{
				title:'昵称'
			}
		},
		{
			path:'/modifyGender',
			name:'modifyGender',
			component:modifyInfo,
			meta:{
				title:'性别'
			}
		},
		{
			path:'/modifyEmail',
			name:'modifyEmail',
			component:modifyInfo,
			meta:{
				title:'邮箱'
			}
		},
		{
			path:'/setting',
			name:'setting',
			component:setting
		},
		{
			path:'/addressSetting',
			name:'addressSetting',
			component:settingPage,
			meta:{
				title:'收货地址'
			}
		},
		{
			path:'/accountAndSecurity',
			name:'accountAndSecurity',
			component:settingPage,
			meta:{
				title:'账户与安全'
			}
		},
		{
			path:'/addAddress',
			name:'addAddress',
			component:modifyAddress,
			meta:{
				title:'添加地址'
			}
		},
		{
			path:'/modifyAddress/:index',
			name:'modifyAddress',
			component:modifyAddress,
			meta:{
				title:'修改地址'
			}
		},
		{
			path:'/loginPassword',
			name:'loginPassword',
			component:loginPassword
		},
		{
			path:'/tradingPassword',
			name:'tradingPassword',
			component:tradingPassword
		},
		{
			path:'/wallet',
			name:'wallet',
			component:wallet
		},
		{
			path:'/goodsDetail',
			name:'goodsDetail',
			component:goodsDetail
		},
		{
			path:'/allComments/:pid',
			name:'allComments',
			component:allComments
		},
		{
			path:'/cart',
			name:'cart',
			component:cart
		},
		{
			path:'/confirmOrder',
			name:'confirmOrder',
			component:confirmOrder
		},
		{
			path:'/paymentPage',
			name:'paymentPage',
			component:paymentPage
		},
		{
			path:'/orderPage',
			name:'orderPage',
			component:orderPage,
			children:[
				{
					path:'unpaid/:name',
					name:'unpaid',
					component:unpaid
				},
				{
					path:'unshipped/:name',
					name:'unshipped',
					component:unshipped
				},
				{
					path:'unreceived/:name',
					name:'unreceived',
					component:unreceived
				},
				{
					path:'comment/:name',
					name:'comment',
					component:comment
				},
				{
					path:'allorders/:name',
					name:'allorders',
					component:allorders
				}
			]
		},
		{
			path:'/orderDetail',
			name:'orderDetail',
			component:orderDetail
		},
		{
			path:'/commentPage',
			name:'commentPage',
			component:commentPage
		},
		{
			path:'/refundApply/:title',
			name:'refundApply',
			component:refundReturnApply
		},
		{
			path:'/returnApply/:title',
			name:'returnApply',
			component:refundReturnApply
		},
		{
			path:'/refundAndReturnPage/:name',
			name:'refundAndReturnPage',
			component:refundAndReturnPage
		},
		{
			path:'/feedback',
			name:'feedback',
			component:feedback
		},
		{
			path:'/chatPage',
			name:'chatPage',
			component:chatPage
		},
		{
			path:'/chatRoom',
			name:'chatRoom',
			component:chatRoom
		}
	]
})